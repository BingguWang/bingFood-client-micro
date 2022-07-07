package biz

import (
    "context"
    "encoding/json"
    "fmt"
    v12 "github.com/BingguWang/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
    v1 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
    "strconv"
    "time"
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
    FindByOrderNumber(context.Context, string) (*entity.Order, error)
    //ListByOrderNumber(context.Context, string) ([]*entity.Order, error)
    //ListAll(context.Context) ([]*entity.Order, error)
    AddSettleToRedis(context.Context, string, string) (string, error)
    GetSettleFromRedis(context.Context, string) (string, error)
    DelSettleFromRedis(context.Context, string) (int64, error)
    InsertOrder(context.Context, *entity.Order) (string, error)
    InsertOrderPay(context.Context, *entity.OrderPay) error
    AfterPaySuccess(context.Context, string) error
    AfterPayTimeout(context.Context, string) error
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
        OrderItems:     items,
        ProdName:       prodName,
    }

    // 返回的结算内容存到redis里,后面的提交订单时不需要前端再传过来了,提交订单的时候删掉
    key := "settledOrder_" + strconv.FormatUint(shopId, 10) + "_" + req.UserMobile // TODO 规范,常数写到其他地方去
    if _, err = oc.repo.AddSettleToRedis(ctx, key, utils.ToJsonString(retData)); err != nil {
        return nil, err
    }
    fmt.Println(utils.ToJsonString(retData))
    return retData, nil
}

func (oc *Ordercase) SubmitOrderHandler(ctx context.Context, req *v1.SubmitOrderRequest) (orderNumber string, err error) {
    oc.log.WithContext(ctx).Infof("SubmitOrderHandler args: %v", utils.ToJsonString(req))
    // 去redis取出结算好的订单信息
    key := "settledOrder_" + strconv.FormatUint(req.ShopId, 10) + "_" + req.UserMobile
    var data string
    data, err = oc.repo.GetSettleFromRedis(ctx, key)
    if err != nil {
        if data == "" {
            err = v1.ErrorInternal("表单已过期，请重新结算")
        }
        return
    }

    var settledOrder v1.SettleOrderReply_Data
    _ = json.Unmarshal([]byte(data), &settledOrder)
    fmt.Println(utils.ToJsonString(settledOrder))

    var od entity.Order
    _ = copier.CopyWithOption(&od, &settledOrder, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    _ = copier.CopyWithOption(&od.ReceiveAddr, req.ReceiveAddr, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })

    for i := 0; i < len(od.OrderItems); i++ {
        od.OrderItems[i].OrderNumber = od.OrderNumber
    }
    od.Remark = req.Remarks

    // 插入order及order_item,更新库存
    if _, err := oc.repo.InsertOrder(ctx, &od); err != nil {
        return "", err
    }

    // redis删除之前保存的结算信息
    if _, err := oc.repo.DelSettleFromRedis(ctx, key); err != nil {
        return "", err
    }

    return od.OrderNumber, nil
}

func (oc *Ordercase) PayOrderHandler(ctx context.Context, req *v1.PayOrderRequest) (result *v1.WxPayMpOrderResult, payNo string, err error) {
    oc.log.WithContext(ctx).Infof("PayOrderHandler args: %v", utils.ToJsonString(req))
    order, err := oc.repo.FindByOrderNumber(ctx, req.OrderNumber)
    if err != nil {
        return nil, "", err
    }
    var (
        totalFee int     // 一共支付多少钱
        openid   = "123" // 用户唯一标识
    )

    orderPay := entity.OrderPay{
        OrderNumber: order.OrderNumber,
        ShopId:      order.ShopId,
        UserId:      order.UserId,
        UserMobile:  order.UserMobile,
        PayAmount:   totalFee,
        PayTypeName: "微信支付",
        PayType:     1,
        PayStatus:   0, // 回调成功后才修改为1
    }
    if err := oc.repo.InsertOrderPay(ctx, &orderPay); err != nil {
        return nil, "", err
    }

    // 传入微信支付SDK需要的参数
    // 模拟调用wx支付接口...
    return MockWxPay(totalFee, orderPay.PayNo, openid), orderPay.PayNo, nil
}
func MockWxPay(totalFee int, outTrade, userId string) *v1.WxPayMpOrderResult {
    time.Sleep(500 * time.Millisecond)
    return &v1.WxPayMpOrderResult{}
}

func (oc *Ordercase) PayOrderSuccessHandler(ctx context.Context, req *v1.PayOrderSuccessRequest) (err error) {
    oc.log.WithContext(ctx).Infof("PayOrderSuccessHandler args: %v", utils.ToJsonString(req))

    if err := oc.repo.AfterPaySuccess(ctx, req.OrderNumber); err != nil {
        return err
    }
    return nil
}

func (oc *Ordercase) PayOrderTimeoutHandler(ctx context.Context, req *v1.PayOrderTimeoutRequest) (err error) {
    oc.log.WithContext(ctx).Infof("PayOrderTimeoutHandler args: %v", utils.ToJsonString(req))

    if err := oc.repo.AfterPayTimeout(ctx, req.OrderNumber); err != nil {
        return err
    }
    return nil
}
