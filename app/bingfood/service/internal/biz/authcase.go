package biz

import (
    "github.com/go-kratos/kratos/v2/log"
)

// AuthCase用于解处理认证逻辑
type AuthCase struct {
    key       string
    userInter UserSrvInterface
    log       *log.Helper
}

func NewAuthCase(userInter UserSrvInterface, logger log.Logger) *AuthCase {
    return &AuthCase{userInter: userInter, log: log.NewHelper(logger)}
}

type AuthSrvInterface interface {
}
