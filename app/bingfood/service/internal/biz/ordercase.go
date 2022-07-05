package biz

import (
    "context"
    v1 "github.com/go-kratos/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v12 "github.com/go-kratos/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

type OrderCase struct {
    oc v12.OrderServiceClient

    log *log.Helper
}
type OrderSrvInterface interface {
    SettleOrder(ctx context.Context, req *v1.SettleOrderRequest) (ret *v12.SettleOrderReply, err error)
}

func NewOrderCase(oc v12.OrderServiceClient, logger log.Logger) *OrderCase {
    return &OrderCase{oc: oc, log: log.NewHelper(logger)}
}

func (oc *OrderCase) SettleOrder(ctx context.Context, req *v1.SettleOrderRequest) (ret *v12.SettleOrderReply, err error) {
    oc.log.WithContext(ctx).Infof("SettleOrder args: %v", utils.ToJsonString(req))

    // todo 假设一下在这里加入context参数，其实是要在其他 地方塞入
    valCtx := context.WithValue(ctx, "userMobile", "15759216850")

    // 取出ctx里的userMobile
    userMobile := valCtx.Value("userMobile").(string)

    // 调用order service 结算订单
    ret, err = oc.oc.SettleOrder(valCtx, &v12.SettleOrderRequest{CartIds: req.CartIds, UserMobile: userMobile})
    log.Infof("调用服务bingfood.order.service/SettleOrder, 得到结果: %v ", utils.ToJsonString(ret))
    if err != nil {
        return nil, err
    }
    return ret, err
}
