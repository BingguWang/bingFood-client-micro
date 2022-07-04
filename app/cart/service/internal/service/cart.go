package service

import (
    "context"
    v1 "github.com/go-kratos/bingfood-client-micro/api/cart/service/v1"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/biz"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/utils"
    "github.com/jinzhu/copier"
    "github.com/prometheus/common/log"
)

type CartServiceImpl struct {
    v1.UnimplementedCartServiceServer

    cc *biz.CartUseCase
}

func NewCartService(uc *biz.CartUseCase) *CartServiceImpl {
    return &CartServiceImpl{cc: uc}
}

func (s *CartServiceImpl) AddCartItem(ctx context.Context, in *v1.AddCartItemRequest) (*v1.AddCartItemReply, error) {
    if err := s.cc.AddCartItem(ctx, in); err != nil {
        return &v1.AddCartItemReply{RetMsg: err.Error()}, err
    }
    return &v1.AddCartItemReply{RetCode: 200, RetMsg: "成功加入购物车 : " + "userMOBILE from ctx"}, nil
}

func (s *CartServiceImpl) GetCartByCond(ctx context.Context, in *v1.GetCartByCondRequest) (*v1.GetCartByCondReply, error) {
    list, total, err := s.cc.GetCartHandler(ctx, in)
    if err != nil {
        return &v1.GetCartByCondReply{
            RetMsg: err.Error(),
            Data:   nil,
        }, err
    }
    var ret []*v1.Cart
    copier.CopyWithOption(ret, list, copier.Option{
        IgnoreEmpty: true,
        DeepCopy:    true,
    })
    return &v1.GetCartByCondReply{
        RetCode: 200,
        RetMsg:  "成功获取购物车 : " + "userMOBILE from ctx",
        Data: &v1.CartPagination{
            List:     ret,
            Total:    total,
            Page:     in.PageInfo.Page,
            PageSize: in.PageInfo.PageSize,
        },
    }, nil
}

func (s *CartServiceImpl) GetCartByCartIds(ctx context.Context, in *v1.GetCartByCartIdsRequest) (*v1.GetCartByCartIdsReply, error) {
    list, total, err := s.cc.GetCartByIdsHandler(ctx, in)
    log.Info("GetCartByIdsHandler ret is : %v", utils.ToJsonString(list))
    if err != nil {
        return &v1.GetCartByCartIdsReply{
            RetMsg: err.Error(),
            Data:   nil,
        }, err
    }
    var ret []*v1.Cart
    copier.CopyWithOption(&ret, &list, copier.Option{
        IgnoreEmpty: false,
    })
    log.Info("000000000000000 ret is : %v", utils.ToJsonString(ret))

    return &v1.GetCartByCartIdsReply{
        RetCode: 200,
        RetMsg:  "成功获取购物车 ",
        Data: &v1.CartPagination{
            List:     ret,
            Total:    total,
            Page:     in.PageInfo.Page,
            PageSize: in.PageInfo.PageSize,
        },
    }, nil
}
