package biz

import (
    "context"
    "fmt"
    v12 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
    v1 "github.com/go-kratos/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
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

func (oc *Ordercase) SettleOrderHandler(ctx context.Context, req *v1.SettleOrderRequest) (ret interface{}, err error) {
    oc.log.WithContext(ctx).Infof("SettleOrderHandler args: %v", utils.ToJsonString(req))

    if req.CartIds == nil || len(req.CartIds) == 0 {
        return nil, v1.ErrorInvalidArgument("传入的购物车id不能为空")
    }
    ret, err = oc.SettleOrder(ctx, req)
    if err != nil {
        return nil, err
    }
    return ret, nil
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
    log.Infof("调用服务bingfood.cart.service/GetCartByCartIds, 得到结果: %v ", utils.ToJsonString(reply))
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
        itemList        []*entity.OrderItem
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
        itemList = append(itemList, &item)
        shopId = v.ShopId
        prodName += sku.ProdName + sku.SkuName + ";"
    }

    // TODO 配送费应该从配送系统计算得到，这里只是用个数值替一下
    deliverFeeTotal = 5 * 100 // 假设是固定的配送费

    discountTotal = (oriPriceTotal - priceTotal) + redPacket
    finalTotal = packingFeeTotal + priceTotal + deliverFeeTotal - discountTotal

    var items []*v1.OrderItem
    copier.CopyWithOption(&items, &itemList, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })

    retData := &v1.SettleOrderReply_Data{
        ShopId:         shopId,
        UserMobile:     req.UserMobile,
        ProdNums:       prodNums,
        PackingAmount:  packingFeeTotal,
        DeliverAmount:  deliverFeeTotal,
        ProdAmount:     priceTotal,
        DiscountAmount: discountTotal,
        FinalAmount:    finalTotal,
        //OrderItems:     items,
        ProdName: prodName,
    }

    // 返回的结算内容存到redis里,后面的提交订单时不需要前端再传过来了,提交订单的时候删掉
    key := "settledOrder_" + strconv.FormatUint(shopId, 10) + "_" + req.UserMobile // TODO 规范,常数写到其他地方去
    if _, err = oc.repo.AddSettleToRedis(ctx, key, utils.ToJsonString(retData)); err != nil {
        return nil, err
    }
    fmt.Println(utils.ToJsonString(retData))
    return retData, nil
}
