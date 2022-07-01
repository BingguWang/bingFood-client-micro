package biz

import (
    "github.com/go-kratos/bingfood-client-micro/app/user/service/internal/biz/user"
    "github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(user.NewUserUsecase)
