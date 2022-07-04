package data

import (
    biz2 "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/biz"
    "github.com/go-kratos/kratos/v2/log"
)

type bingfoodRepo struct {
    data *Data
    log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz2.BingfoodRepo {
    return &bingfoodRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
