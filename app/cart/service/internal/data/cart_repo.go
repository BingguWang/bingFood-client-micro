package data

import (
    "context"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/biz"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/data/entity"
    "github.com/go-kratos/bingfood-client-micro/app/cart/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type cartRepo struct {
    data *Data
    log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
    return &cartRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
func (c *cartRepo) AddCartRepo(ctx context.Context, cartRow *entity.Cart) error {
    c.log.WithContext(ctx).Infof("cart to insert  : %v", utils.ToJsonString(cartRow))

    db := c.data.db
    if err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Clauses(clause.OnConflict{
            Columns:   []clause.Column{{Name: "shop_id"}, {Name: "user_id"}, {Name: "sku_id"}}, // key column，如果id已存在则变为更新操作
            DoUpdates: clause.AssignmentColumns([]string{"sku_id", "prod_nums"}),               // 更新操作要更新的字段,更新为新值
        }).Create(&cartRow).Error
        return err
    }); err != nil {
        c.log.WithContext(ctx).Infof("insert prod failed : %v", err.Error())
        return err
    }
    return nil
}

func (c *cartRepo) GetCart(ctx context.Context, cart *entity.Cart, limit, offset int) (ret []*entity.Cart, total int64, err error) {
    c.log.WithContext(ctx).Infof("get cart condition is : %v", utils.ToJsonString(cart))

    db := c.data.db

    db.Where(&cart).Count(&total)
    err = db.Limit(limit).Offset(offset).Where(&cart).Find(&ret).Error
    if err != nil {
        return
    }
    return ret, total, nil
}

func (c *cartRepo) GetCartByIds(ctx context.Context, ids []uint64, limit int, offset int) (ret []*entity.Cart, total int64, err error) {
    c.log.WithContext(ctx).Infof("get cart by ids ,ids  is : %v", utils.ToJsonString(ids))

    db := c.data.db

    db.Where("cart_id in ?", ids).Count(&total)
    err = db.Limit(limit).Offset(offset).Where("cart_id in ?", ids).Find(&ret).Error
    if err != nil {
        return
    }
    return ret, total, nil
}
