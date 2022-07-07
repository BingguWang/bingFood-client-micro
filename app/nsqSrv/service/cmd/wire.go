//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/data"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/server"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/service"
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.JWT, log.Logger, *conf.NSQ, *conf.Registry) (*kratos.App, func(), error) {
    panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
