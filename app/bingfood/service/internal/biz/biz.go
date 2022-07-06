package biz

import (
    "fmt"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/google/wire"
    clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
    NewOrderServiceClient,
    NewCartServiceClient,
    NewUserServiceClient,
    NewProdServiceClient,
    NewAuthCase,
    NewOrderCase,
    NewProdCase,
    NewCartCase,
    NewUserCase,
    NewDiscovery,

    NewUserSrvInterface,
)

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
