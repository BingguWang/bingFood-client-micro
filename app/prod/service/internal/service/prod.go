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
    var limit, offset int64
    if in.PageInfo != nil {
        limit = in.PageInfo.PageSize
        offset = in.PageInfo.Page
    }
    in.PageInfo = &v1.PageInfo{Page: offset, PageSize: limit}
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
            Page:     limit,
            PageSize: offset,
        },
    }, nil
}

func (s *ProdServiceImpl) UpdateSkuStock(ctx context.Context, in *v1.UpdateSkuStockRequest) (*v1.UpdateSkuStockReply, error) {
    if err := s.pc.UpdateSkuStockHandler(ctx, in); err != nil {
        return &v1.UpdateSkuStockReply{
            RetMsg: "更新库存失败 : " + err.Error(),
        }, err
    }
    return &v1.UpdateSkuStockReply{
        RetCode: 200,
        RetMsg:  "成功更新库存",
    }, nil
}
