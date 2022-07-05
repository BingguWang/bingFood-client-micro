// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/biz"
	"github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/conf"
	"github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/server"
	"github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, jwt *conf.JWT, logger log.Logger, registry *conf.Registry) (*kratos.App, func(), error) {
	discovery := biz.NewDiscovery(registry)
	orderServiceClient := biz.NewOrderServiceClient(discovery, jwt)
	orderCase := biz.NewOrderCase(orderServiceClient, logger)
	cartServiceClient := biz.NewCartServiceClient(discovery, jwt)
	cartCase := biz.NewCartCase(cartServiceClient, logger)
	bingfoodServiceImpl := service.NewBingfoodService(orderCase, cartCase, logger)
	grpcServer := server.NewGRPCServer(confServer, bingfoodServiceImpl, logger)
	httpServer := server.NewHTTPServer(confServer, jwt, bingfoodServiceImpl, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
	}, nil
}
