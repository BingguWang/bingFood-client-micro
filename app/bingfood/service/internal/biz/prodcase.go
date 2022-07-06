package biz

import (
    "context"
    v13 "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    jwt2 "github.com/golang-jwt/jwt/v4"
)

type ProdCase struct {
    psc v13.ProdServiceClient

    log *log.Helper
}
type ProdSrvInterface interface {
    GetSkuByCondFunc(ctx context.Context, in *v13.GetSkuByCondRequest) (*v13.GetSkuByCondReply, error)
}

func NewProdCase(psc v13.ProdServiceClient, logger log.Logger) *ProdCase {
    return &ProdCase{psc: psc, log: log.NewHelper(logger)}
}
func (pc *ProdCase) GetSkuByCondFunc(ctx context.Context, req *v13.GetSkuByCondRequest) (*v13.GetSkuByCondReply, error) {
    pc.log.WithContext(ctx).Infof("GetSkuByCond args: %v", utils.ToJsonString(req))

    // 调用order service 获取sku
    ret, err := pc.psc.GetSkuByCond(ctx, req)
    log.Infof("调用服务bingfood.prod.service/GetSkuByCond, 得到结果: %v ", utils.ToJsonString(ret))
    if err != nil {
        return nil, err
    }
    return ret, err
}

func NewProdServiceClient(r registry.Discovery, ac *conf.JWT) v13.ProdServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.prod.service"),
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
    c := v13.NewProdServiceClient(conn)
    return c

}
