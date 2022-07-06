package biz

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v13 "github.com/BingguWang/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
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

type CartCase struct {
    cc v13.CartServiceClient

    log *log.Helper
}

func NewCartCase(cc v13.CartServiceClient, logger log.Logger) *CartCase {
    return &CartCase{cc: cc, log: log.NewHelper(logger)}
}

type CartSrvInterface interface {
    AddCartItem(ctx context.Context, req *v1.AddCartItemRequest) (err error)
    GetCartDetail(ctx context.Context, in *v1.GetCartByCondRequest) (*v1.CartPagination, error)
}

func (uc *CartCase) AddCartItem(ctx context.Context, req *v1.AddCartItemRequest) (err error) {
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
    ret, err := uc.cc.AddCartItem(ctx, rq)
    log.Infof("调用服务bingfood.cart.service/AddCartItem, 得到结果: %v ", utils.ToJsonString(ret))
    if err != nil {
        return err
    }
    return err
}

func (uc *CartCase) GetCartDetail(ctx context.Context, req *v1.GetCartByCondRequest) (*v1.CartPagination, error) {
    uc.log.WithContext(ctx).Infof("GetCartDetail args: %v", utils.ToJsonString(req))
    var rq v13.GetCartByCondRequest
    copier.Copy(&rq, req)
    log.Infof("参数：%s", utils.ToJsonString(rq))

    ret, err := uc.cc.GetCartByCond(ctx, &rq)
    if err != nil {
        return nil, err
    }
    log.Infof("调用服务bingfood.cart.service/GetCartByCond, 得到结果: %v ", utils.ToJsonString(ret.Data))
    var reply v1.CartPagination
    copier.CopyWithOption(&reply, &ret.Data, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    return &reply, nil
}

func NewCartServiceClient(r registry.Discovery, ac *conf.JWT) v13.CartServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.cart.service"),
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
    c := v13.NewCartServiceClient(conn)
    return c
}
