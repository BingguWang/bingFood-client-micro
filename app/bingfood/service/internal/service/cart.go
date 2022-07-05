package service

import (
    "context"
    v1 "github.com/go-kratos/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

func (s *BingfoodServiceImpl) AddCartItem(ctx context.Context, in *v1.AddCartItemRequest) (*v1.AddCartItemReply, error) {
    reply, err := s.cc.AddCartItem(ctx, in)
    if err != nil {
        return &v1.AddCartItemReply{RetMsg: err.Error()}, err
    }
    log.Infof(utils.ToJsonString(reply))
    return &v1.AddCartItemReply{RetCode: 200, RetMsg: "成功新增购物车"}, nil
}
