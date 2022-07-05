package biz

import (
    "context"
    "fmt"
    v13 "github.com/go-kratos/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    v12 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    jwt2 "github.com/golang-jwt/jwt/v4"
    "github.com/google/wire"
    clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBingfoodCase, NewDiscovery, NewOrderServiceClient, NewCartServiceClient)

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

func NewOrderServiceClient(r registry.Discovery, ac *conf.JWT) v13.OrderServiceClient {
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
    c := v13.NewOrderServiceClient(conn)
    return c
}
func NewCartServiceClient(r registry.Discovery, ac *conf.JWT) v12.CartServiceClient {
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
    c := v12.NewCartServiceClient(conn)
    return c
}
