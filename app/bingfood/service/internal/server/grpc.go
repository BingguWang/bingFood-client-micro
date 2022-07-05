package server

import (
    "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/service"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC configs.
func NewGRPCServer(c *conf.Server, svc *service.BingfoodServiceImpl, logger log.Logger) *grpc.Server {
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            recovery.Recovery(),
        ),
    }
    if c.Grpc.Network != "" {
        opts = append(opts, grpc.Network(c.Grpc.Network))
    }
    if c.Grpc.Addr != "" {
        opts = append(opts, grpc.Address(c.Grpc.Addr))
    }
    if c.Grpc.Timeout != nil {
        opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
    }
    srv := grpc.NewServer(opts...)
    v1.RegisterBingfoodServiceServer(srv, svc)
    return srv
}
