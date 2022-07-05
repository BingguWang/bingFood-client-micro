package biz

import (
    "context"
    v1 "github.com/go-kratos/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v13 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
)

type CartCase struct {
    cc v13.CartServiceClient

    log *log.Helper
}

func NewCartCase(cc v13.CartServiceClient, logger log.Logger) *CartCase {
    return &CartCase{cc: cc, log: log.NewHelper(logger)}
}

type CartSrvInterface interface {
    AddCartItem(ctx context.Context, req *v1.AddCartItemRequest) (ret *v13.AddCartItemReply, err error)
}

func (uc *CartCase) AddCartItem(ctx context.Context, req *v1.AddCartItemRequest) (ret *v13.AddCartItemReply, err error) {
    uc.log.WithContext(ctx).Infof("AddCartItem args: %v", utils.ToJsonString(req))

    // todo 假设一下在这里加入context参数，其实是要在其他 地方塞入
    valCtx := context.WithValue(ctx, "userMobile", "15759216850")

    // 取出ctx里的userMobile
    _ = valCtx.Value("userMobile").(string)

    // 调用order service 结算订单

    rq := &v13.AddCartItemRequest{}
    copier.CopyWithOption(rq, req, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    ret, err = uc.cc.AddCartItem(ctx, rq)
    log.Infof("调用服务bingfood.cart.service/AddCartItem, 得到结果: %v ", utils.ToJsonString(ret))
    if err != nil {
        return nil, err
    }
    return ret, err
}
