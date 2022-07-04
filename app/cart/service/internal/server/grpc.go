package server

import (
    v12 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/conf"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/service"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC configs.
func NewGRPCServer(c *conf.Server, svc *service.CartServiceImpl, logger log.Logger) *grpc.Server {
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
    v12.RegisterCartServiceServer(srv, svc)
    return srv
}
