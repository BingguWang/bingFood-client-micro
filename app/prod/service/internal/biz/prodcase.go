package biz

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
)

type ProdRepo interface {
    GetSkuByCond(ctx context.Context, sku *entity.Sku, limit, offset int) (ret []*entity.Sku, total int64, err error)
    GetSkuBySkuIds(ctx context.Context, ids []uint64, limit, offset int) (ret []*entity.Sku, total int64, err error)
}
type ProdUseCase struct {
    repo ProdRepo

    log *log.Helper
}

func NewProdUseCase(repo ProdRepo, logger log.Logger) *ProdUseCase {
    return &ProdUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (pc *ProdUseCase) GetSkuByCondHandler(ctx context.Context, req *v1.GetSkuByCondRequest) (ret []*v1.Sku, total int64, err error) {
    pc.log.WithContext(ctx).Infof("GetSkuByCond args: %v", utils.ToJsonString(req))
    limit := (int)(req.PageInfo.PageSize)
    offset := (int)(req.PageInfo.PageSize)
    var r entity.Sku
    copier.CopyWithOption(&r, &req.SkuCond, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    skuList, total, err := pc.repo.GetSkuByCond(ctx, &r, limit, offset)
    if err != nil {
        return nil, 0, err
    }
    copier.CopyWithOption(&ret, &skuList, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    return ret, total, err
}
