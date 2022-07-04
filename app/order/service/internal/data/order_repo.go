package data

import (
    "context"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/data/entity"
    "github.com/go-kratos/kratos/v2/log"
)

type orderRepo struct {
    data *Data
    log  *log.Helper
}

// NewGreeterRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
    return &orderRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}
func (o *orderRepo) FindByOrderNumber(context.Context, int64) (*entity.Order, error) {
    return nil, nil
}

func (o *orderRepo) AddSettleToRedis(ctx context.Context, k, v string) (string, error) {
    return o.RedisSetKV(ctx, k, v)
}
