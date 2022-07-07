package biz

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v12 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    jwt2 "github.com/golang-jwt/jwt/v4"
    "github.com/jinzhu/copier"
)

type OrderCase struct {
    oc v12.OrderServiceClient

    log *log.Helper
}
type OrderSrvInterface interface {
    SettleOrder(ctx context.Context, req *v1.SettleOrderRequest) (ret *v1.SettleOrderReply_Data, err error)
    SubmitOrder(ctx context.Context, req *v1.SubmitOrderRequest) (string, error)
    PayOrder(ctx context.Context, req *v1.PayOrderRequest) (*v1.WxPayMpOrderResult, error)
}

func NewOrderCase(oc v12.OrderServiceClient, logger log.Logger) *OrderCase {
    return &OrderCase{oc: oc, log: log.NewHelper(logger)}
}

func (oc *OrderCase) SettleOrder(ctx context.Context, req *v1.SettleOrderRequest) (*v1.SettleOrderReply_Data, error) {
    oc.log.WithContext(ctx).Infof("SettleOrder args: %v", utils.ToJsonString(req))

    // todo 假设一下在这里加入context参数，其实是要在其他 地方塞入
    valCtx := context.WithValue(ctx, "userMobile", "15759216850")

    // 取出ctx里的userMobile
    userMobile := valCtx.Value("userMobile").(string)

    // 调用order service 结算订单
    result, err := oc.oc.SettleOrder(valCtx, &v12.SettleOrderRequest{CartIds: req.CartIds, UserMobile: userMobile})
    if err != nil {
        return nil, err
    }
    log.Infof("调用服务bingfood.order.service/SettleOrder, 得到结果: %v ", utils.ToJsonString(result))
    var ret v1.SettleOrderReply_Data
    copier.CopyWithOption(&ret, result.Data, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    return &ret, err
}

func (oc *OrderCase) SubmitOrder(ctx context.Context, req *v1.SubmitOrderRequest) (string, error) {
    oc.log.WithContext(ctx).Infof("SubmitOrder args: %v", utils.ToJsonString(req))

    var r v12.SubmitOrderRequest
    copier.CopyWithOption(&r, req, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    ret, err := oc.oc.SubmitOrder(ctx, &r)
    if err != nil {
        return "", err
    }
    log.Infof("调用服务bingfood.order.service/SubmitOrder, 得到结果: %v ", utils.ToJsonString(ret))
    return ret.OrderNumber, nil
}

func (oc *OrderCase) PayOrder(ctx context.Context, req *v1.PayOrderRequest) (*v1.WxPayMpOrderResult, string, error) {
    oc.log.WithContext(ctx).Infof("PayOrder args: %v", utils.ToJsonString(req))
    r := &v12.PayOrderRequest{OrderNumber: req.OrderNumber}
    ret, err := oc.oc.PayOrder(ctx, r)
    if err != nil {
        return nil, "", err
    }
    log.Infof("调用服务bingfood.order.service/PayOrder, 得到结果: %v ", utils.ToJsonString(ret))
    var result v1.WxPayMpOrderResult
    copier.Copy(&result, &ret.WxPayMpOrderResult)
    return &result, ret.PayNo, err
}

func (oc *OrderCase) NoticePayOrderHandler(ctx context.Context, req *v1.NoticePayOrderRequest) error {
    oc.log.WithContext(ctx).Infof("NoticePayOrderHandler args: %v", utils.ToJsonString(req))
    r := &v12.PayOrderSuccessRequest{OrderNumber: req.PayNo}

    if _, err := oc.oc.PayOrderSuccess(ctx, r); err != nil {
        return err
    }
    log.Infof("调用服务bingfood.order.service/PayOrder")
    return nil
}

func NewOrderServiceClient(r registry.Discovery, ac *conf.JWT) v12.OrderServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.order.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
            jwt.Client(func(token *jwt2.Token) (interface{}, error) {
                return []byte(ac.ServiceSecretKey), nil
            }, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v12.NewOrderServiceClient(conn)
    return c
}
