package biz

import (
    v13 "github.com/go-kratos/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    "github.com/go-kratos/kratos/v2/log"
)

// 专门用于处理user相关的逻辑
type UserCase struct {
    cc v13.UserServiceClient

    log *log.Helper
}

func NewUserCase(cc v13.UserServiceClient, logger log.Logger) *UserCase {
    return &UserCase{cc: cc, log: log.NewHelper(logger)}
}

type UserSrvInterface interface {
}
