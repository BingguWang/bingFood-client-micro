package service

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/biz"
)

type ProdServiceImpl struct {
    v1.UnimplementedProdServiceServer

    pc *biz.ProdUseCase
}

func NewProdServiceImpl(pc *biz.ProdUseCase) *ProdServiceImpl {
    return &ProdServiceImpl{pc: pc}
}
func (s *ProdServiceImpl) GetSkuByCond(ctx context.Context, in *v1.GetSkuByCondRequest) (*v1.GetSkuByCondReply, error) {
    list, total, err := s.pc.GetSkuByCondHandler(ctx, in)
    if err != nil {
        return &v1.GetSkuByCondReply{
            RetMsg: err.Error(),
            Data:   nil,
        }, err
    }
    return &v1.GetSkuByCondReply{
        RetCode: 200,
        RetMsg:  "成功获取购物车",
        Data: &v1.CartPagination{
            List:     list,
            Total:    total,
            Page:     in.PageInfo.Page,
            PageSize: in.PageInfo.PageSize,
        },
    }, nil
}
