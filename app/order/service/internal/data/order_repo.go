package data

import (
    "context"
    "fmt"
    v12 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v11 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "strconv"
    "time"
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
func (o *orderRepo) FindByOrderNumber(ctx context.Context, orderNumber string) (*entity.Order, error) {
    db := o.data.db
    var od []*entity.Order
    if err := db.Where(&entity.Order{OrderNumber: orderNumber}).
        Where("order_status = 0").
        Find(&od).Error; err != nil {
        return nil, err
    }
    if len(od) == 0 {
        return nil, v12.ErrorInternal(fmt.Sprintf("传入订单号不正确，不存在此未支付的订单:%v", orderNumber))
    }
    return od[0], nil
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
        if e := o.changeSkuStock(order.OrderItems, ctx, -1); e != nil {
            return e
        }

        // todo 把订单号存入到MQ里,超时未支付则状态设为取消, 因为没有用分布式事务，所以操作放在这里，用了dtm后应该放在bingfood服务里
        //if err = PubOrderNumberToMQ(od.OrderNumber); err != nil {
        //    return
        //}

        return
    }); err != nil {
        // todo 回滚库存, 暂时使用手动回滚，后面可以结合dtm分布式事务
        if e := o.changeSkuStock(order.OrderItems, ctx, 1); e != nil {
            err = e
        }
        return "", err
    }
    return order.OrderNumber, nil
}

func (o *orderRepo) changeSkuStock(items []entity.OrderItem, ctx context.Context, isAdd int64) error {
    for _, item := range items {
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

func (o *orderRepo) InsertOrderPay(ctx context.Context, orderPay *entity.OrderPay) error {
    log.Infof("InsertOrderPay , req is : %v", utils.ToJsonString(orderPay))
    db := o.data.db

    // 生成orderPay的支付号PayNo,这个值作为外部号码传给微信支付接口
    number := o.data.node.Generate()
    orderPay.PayNo = strconv.FormatInt(number.Int64(), 10)

    if err := db.Transaction(func(tx *gorm.DB) (err error) {
        if err = tx.Create(&orderPay).Error; err != nil {
            log.Errorf("insert orderPay failed : %v", err.Error())
            return
        }
        return
    }); err != nil {
        return err
    }
    return nil
}

func (o *orderRepo) AfterPaySuccess(ctx context.Context, payNo string) error {
    log.Infof("AfterPaySuccess , req is : %v", utils.ToJsonString(payNo))

    db := o.data.db
    if err := db.Transaction(func(tx *gorm.DB) error {
        var orderPay []entity.OrderPay
        if err := db.Where(&entity.OrderPay{PayNo: payNo}).Find(&orderPay).Error; err != nil {
            return err
        }
        if len(orderPay) == 0 {
            return v11.ErrorInternal(fmt.Sprintf("支付信息有误,payNo:%v", payNo))
        }

        // 修改订单支付表的支付状态
        log.Infof("修改订单支付表信息...")
        tx2 := tx.Model(&entity.OrderPay{}).Where("pay_no = ? AND pay_status = 0", payNo).
            Update("pay_status", 1)
        if rows := tx2.RowsAffected; rows == 0 {
            return v11.ErrorInternal("the orderPay has been paid , payNo : %v", payNo)
        }

        // 修改订单状态为已支付
        log.Infof("修改订单表信息...")
        txx := tx.Model(&entity.Order{}).Where("order_number = ? AND order_status = 0", orderPay[0].OrderNumber).
            Select("order_status", "pay_type", "pay_at").
            Updates(map[string]interface{}{"order_status": 1, "pay_type": orderPay[0].PayType, "pay_at": time.Now()})
        if rows := txx.RowsAffected; rows == 0 {
            return v11.ErrorInternal("order has been paid , orderNumber : %v", orderPay[0].OrderNumber)
        }
        return nil
    }); err != nil {
        return err
    }

    return nil
}

func (o *orderRepo) AfterPayTimeout(ctx context.Context, orderNumber string) error {
    log.Infof("AfterPayTimeout , req is : %v", utils.ToJsonString(orderNumber))

    db := o.data.db
    var order entity.Order
    if err := db.Transaction(func(tx *gorm.DB) error {
        // TODO 状态转换用用状态机
        txx := tx.Model(&entity.Order{}).Where("order_number = ? AND order_status = 0", orderNumber).
            Update("order_status", 4)
        if rows := txx.RowsAffected; rows == 0 {
            log.Infof("has been paid ,orderNumber:%v", orderNumber)
            return nil
        }
        if err := txx.Error; err != nil {
            log.Errorf("orderNumber:%v, update order_status failed : %v", orderNumber, err)
            return err
        }

        if err := tx.Preload("OrderItems").
            Where(&entity.Order{OrderNumber: orderNumber}).
            Find(&order).Error; err != nil {
            log.Errorf("select order failed : %v", err)
            return err
        }
        // 恢复库存,调用prod服务
        if e := o.changeSkuStock(order.OrderItems, ctx, 1); e != nil {
            return e
        }

        return nil
    }); err != nil {
        // 回滚 恢复库存操作
        if e := o.changeSkuStock(order.OrderItems, ctx, -1); e != nil {
            err = e
        }
        log.Errorf("transaction failed : %v", err)
        return err
    }
    return nil
}
