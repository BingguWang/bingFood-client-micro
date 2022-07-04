package middleware

import (
    "context"
    "fmt"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
)

func AuthMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            if tr, ok := transport.FromServerContext(ctx); ok {
                // Do something on entering
                defer func() {
                    // Do something on exiting
                }()
                tokenString := tr.RequestHeader().Get("token")

                if tokenString == "" {
                    return nil, errors.New(400, "未登录或非法访问", "failed")
                }
                fmt.Println(tokenString)

                claims, err := utils.ParseToken(tokenString)
                fmt.Println(utils.ToJsonString(claims))
                if err != nil {
                    if err == utils.TokenIsExpired {
                        return nil, errors.New(400, "token授权已过期", "failed")
                    }
                    return nil, err
                }
            }
            return handler(ctx, req)
        }
    }
}
