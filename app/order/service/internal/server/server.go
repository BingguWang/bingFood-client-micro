package server

import (
	"fmt"
	"github.com/go-kratos/bingfood-client-micro/app/order/service/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is configs providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(cf *conf.Registry) registry.Registrar {
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
