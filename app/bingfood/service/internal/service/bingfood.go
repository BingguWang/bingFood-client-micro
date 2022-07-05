package service

import (
    v1 "github.com/go-kratos/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/biz"
    "github.com/go-kratos/kratos/v2/log"
)

type BingfoodServiceImpl struct {
    v1.UnimplementedBingfoodServiceServer
    oc *biz.OrderCase
    cc *biz.CartCase
    //ac  *biz.AuthCase
    log *log.Helper
}

func NewBingfoodService(
    oc *biz.OrderCase,
    cc *biz.CartCase,
//ac *biz.AuthCase,
    logger log.Logger,
) *BingfoodServiceImpl {
    return &BingfoodServiceImpl{
        cc: cc,
        oc: oc,
        //ac:  ac,
        log: log.NewHelper(log.With(logger, "module", "service/interface")),
    }
}
