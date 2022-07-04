package biz

import (
    "context"
    "fmt"
    v12 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1"
    v1 "github.com/go-kratos/bingfood-client-micro/api/order/service/v1"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/conf"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/jinzhu/copier"
    clientv3 "go.etcd.io/etcd/client/v3"
    "strconv"
)

type Ordercase struct {
    repo OrderRepo
    cc   v12.CartServiceClient

    log *log.Helper
}

func NewOrdercase(repo OrderRepo, cc v12.CartServiceClient, logger log.Logger) *Ordercase {
    return &Ordercase{repo: repo, cc: cc, log: log.NewHelper(logger)}
}

type OrderRepo interface {
    //Save(context.Context, *entity.Order) error
    //Update(context.Context, *entity.Order) error
    FindByOrderNumber(context.Context, int64) (*entity.Order, error)
    //ListByOrderNumber(context.Context, string) ([]*entity.Order, error)
    //ListAll(context.Context) ([]*entity.Order, error)

    AddSettleToRedis(context.Context, string, string) (string, error)
}

func NewDiscovery(cf *conf.Registry) registry.Discovery {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   cf.Etcd.Endpoints,
        DialTimeout: cf.Etcd.DialTimeout.AsDuration(),
    })
    fmt.Println(cf.Etcd.Endpoints)
    if err != nil {
        panic(err)
    }
    r := etcd.New(client)
    return r
}

func NewCartServiceClient(r registry.Discovery) v12.CartServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.cart.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v12.NewCartServiceClient(conn)
    return c
}

func (oc *Ordercase) SettleOrderHandler(ctx context.Context, req *v1.SettleOrderRequest) error {
    oc.log.WithContext(ctx).Infof("SettleOrderHandler args: %v", utils.ToJsonString(req))

    if req.CartIds == nil || len(req.CartIds) == 0 {
        return v1.ErrorInvalidArgument("传入的购物车id不能为空")
    }
    if _, err := oc.SettleOrder(ctx, req); err != nil {
        return err
    }
    return nil
}

func (oc *Ordercase) SettleOrder(ctx context.Context, req *v1.SettleOrderRequest) (res interface{}, err error) {
    oc.log.WithContext(ctx).Infof("request args are : %v", utils.ToJsonString(req))

    // todo 调用服务获取购物车
    cartServiceClient := oc.cc
    rq := &v12.GetCartByCartIdsRequest{Ids: req.CartIds, PageInfo: &v12.PageInfo{
        Page:     0,
        PageSize: 10,
    }}
    reply, err := cartServiceClient.GetCartByCartIds(ctx, rq)
    if err != nil {
        return nil, err
    }
    cartRows := reply.Data.List

    var (
        oriPriceTotal   int32 // 原价总和
        packingFeeTotal int32 // 打包费
        priceTotal      int32 // 现价总和
        finalTotal      int32 // 最后需支付金额
        discountTotal   int32 // 总共优惠的金额
        deliverFeeTotal int32 // 配送费
        redPacket       int32 // 红包
        itemList        []entity.OrderItem
        prodNums        int32 // 总商品个数
        shopId          uint64
        prodName        string // 商品名，用分号连接
    )

    for _, v := range cartRows {
        sku := v.Sku
        item := entity.OrderItem{
            UserId:     v.UserId,
            ShopId:     v.ShopId,
            ProdId:     0,
            ProdName:   sku.ProdName,
            ProdNums:   (int)(v.ProdNums),
            Pic:        sku.Pic,
            Price:      (int)(sku.Price),
            ProdAmount: (int)(sku.Price * v.ProdNums),
            OriPrice:   (int)(sku.OriPrice),
            SkuId:      sku.SkuId,
            SkuName:    sku.SkuName,
            PropId:     sku.ProdId,
            PropName:   sku.ProdName + sku.SkuName,
        }
        oriPriceTotal += sku.OriPrice
        priceTotal += sku.Price
        packingFeeTotal += sku.PackingFee

        prodNums += v.ProdNums
        itemList = append(itemList, item)
        shopId = v.ShopId
        prodName += sku.ProdName + sku.SkuName + ";"
    }

    // TODO 配送费应该从配送系统计算得到，这里只是用个数值替一下
    deliverFeeTotal = 5 * 100 // 假设是固定的配送费

    discountTotal = (oriPriceTotal - priceTotal) + redPacket
    finalTotal = packingFeeTotal + priceTotal + deliverFeeTotal - discountTotal

    claims := ctx.Value("claims")
    userClaims := claims.(*utils.UserClaims)
    //fmt.Println(itemList)

    var items *v1.OrderItem
    copier.CopyWithOption(items, itemList, copier.Option{
        IgnoreEmpty: true,
        DeepCopy:    true,
    })

    retData := v1.SettleOrderReply_Data{
        ShopId:         shopId,
        UserMobile:     userClaims.UserMobile,
        ProdNums:       prodNums,
        PackingAmount:  packingFeeTotal,
        DeliverAmount:  deliverFeeTotal,
        ProdAmount:     priceTotal,
        DiscountAmount: discountTotal,
        FinalAmount:    finalTotal,
        OrderItems:     items,
        ProdName:       prodName,
    }
    fmt.Println(utils.ToJsonString(retData))

    // 返回的结算内容存到redis里,后面的提交订单时不需要前端再传过来了,提交订单的时候删掉
    key := "settledOrder_" + strconv.FormatUint(shopId, 10) + "_" + userClaims.UserMobile // TODO 规范,常数写到其他地方去
    oc.repo.AddSettleToRedis(ctx, key, utils.ToJsonString(retData))
    if err != nil {
        return
    }
    return retData, nil
}
