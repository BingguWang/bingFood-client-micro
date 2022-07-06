package data

import (
    "context"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/prod/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

type prodRepo struct {
    data *Data
    log  *log.Helper
}

func NewProdRepo(data *Data, logger log.Logger) biz.ProdRepo {
    return &prodRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
func (c *prodRepo) GetSkuByCond(ctx context.Context, sku *entity.Sku, limit, offset int) (ret []*entity.Sku, total int64, err error) {
    c.log.WithContext(ctx).Infof("get sku condition is : %v", utils.ToJsonString(sku))
    db := c.data.db

    db.Where(&sku).Count(&total)
    err = db.Limit(limit).Offset(offset).Where(&sku).Find(&ret).Error
    if err != nil {
        return
    }
    return ret, total, nil
}

func (c *prodRepo) GetSkuBySkuIds(ctx context.Context, ids []uint64, limit, offset int) (ret []*entity.Sku, total int64, err error) {
    c.log.WithContext(ctx).Infof("get sku by ids ,ids  is : %v", utils.ToJsonString(ids))
    db := c.data.db

    db.Where("sku_id in ?", ids).Count(&total)
    err = db.Limit(limit).Offset(offset).Where("sku_id in ?", ids).Find(&ret).Error
    if err != nil {
        return
    }
    return ret, total, nil
}
