package biz

import (
    "context"
    "fmt"
    v13 "github.com/go-kratos/bingfood-client-micro/api/order/service/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/google/wire"
    clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBingfoodCase, NewDiscovery, NewOrderServiceClient)

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

func NewOrderServiceClient(r registry.Discovery) v13.OrderServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.order.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v13.NewOrderServiceClient(conn)
    return c
}
