package service

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
)

func (s *BingfoodServiceImpl) AddCartItem(ctx context.Context, in *v1.AddCartItemRequest) (*v1.AddCartItemReply, error) {
    err := s.cc.AddCartItem(ctx, in)
    if err != nil {
        return &v1.AddCartItemReply{RetMsg: err.Error()}, err
    }
    return &v1.AddCartItemReply{RetCode: 200, RetMsg: "成功新增购物车"}, nil
}

func (s *BingfoodServiceImpl) GetCartDetail(ctx context.Context, in *v1.GetCartByCondRequest) (*v1.GetCartByCondReply, error) {
    data, err := s.cc.GetCartDetail(ctx, in)
    if err != nil {
        return &v1.GetCartByCondReply{
            RetMsg: err.Error(), Data: nil}, err
    }
    return &v1.GetCartByCondReply{RetCode: 200, RetMsg: "成功获取购物车详情", Data: data}, nil
}
