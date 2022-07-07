package data

import (
    "context"
    "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "strconv"
)

type orderRepo struct {
    data *Data
    pc   v1.ProdServiceClient

    log *log.Helper
}

// NewGreeterRepo .
func NewOrderRepo(data *Data, pc v1.ProdServiceClient, logger log.Logger) biz.OrderRepo {
    return &orderRepo{
        data: data,
        pc:   pc,
        log:  log.NewHelper(logger),
    }
}
func (o *orderRepo) FindByOrderNumber(context.Context, int64) (*entity.Order, error) {
    return nil, nil
}

func (o *orderRepo) AddSettleToRedis(ctx context.Context, k, v string) (string, error) {
    return o.RedisSetKV(ctx, k, v)
}
func (o *orderRepo) GetSettleFromRedis(ctx context.Context, k string) (string, error) {
    return o.RedisGetVal(ctx, k)
}
func (o *orderRepo) DelSettleFromRedis(ctx context.Context, k string) (int64, error) {
    return o.RedisDelKey(ctx, k)
}
func (o *orderRepo) InsertOrder(ctx context.Context, order *entity.Order) (string, error) {
    log.Infof("InsertOrder , req is : %v", utils.ToJsonString(order))
    db := o.data.db

    // 生成orderNumber
    number := o.data.node.Generate()
    order.OrderNumber = strconv.FormatInt(number.Int64(), 10)

    if err := db.Transaction(func(tx *gorm.DB) (err error) {
        if err = tx.Omit(clause.Associations).Create(&order.OrderItems).Error; err != nil {
            log.Errorf("insert orderItem failed : %v", err.Error())
            return
        }
        if err = tx.Omit(clause.Associations).Create(&order).Error; err != nil {
            log.Errorf("insert order failed : %v", err.Error())
            return err
        }
        // 更新库存
        if e := o.changeSkuStock(order, ctx, -1); e != nil {
            return e
        }

        // todo 把订单号存入到MQ里,超时未支付则状态设为取消, 因为没有用分布式事务，所以操作放在这里，用了dtm后应该放在bingfood服务里
        //if err = PubOrderNumberToMQ(od.OrderNumber); err != nil {
        //    return
        //}

        return
    }); err != nil {
        // todo 回滚库存, 暂时使用手动回滚，后面可以结合dtm分布式事务
        if e := o.changeSkuStock(order, ctx, 1); e != nil {
            err = e
        }
        return "", err
    }
    return order.OrderNumber, nil
}

func (o *orderRepo) changeSkuStock(order *entity.Order, ctx context.Context, isAdd int64) error {
    for _, item := range order.OrderItems {
        // 更新库存
        r := v1.UpdateSkuStockRequest{
            SkuId:     item.SkuId,
            ChangeVal: int64(item.ProdNums) * isAdd,
        }
        if _, err := o.pc.UpdateSkuStock(ctx, &r); err != nil {
            return err
        }
        log.Infof("调用服务bingfood.order.service/UpdateSkuStock")
    }
    return nil
}
