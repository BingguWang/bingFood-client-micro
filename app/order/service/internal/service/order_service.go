package service

import (
    "context"
    "github.com/go-kratos/bingfood-client-micro/api/order/service/v1"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/biz"
)

type OrderServiceImpl struct {
    v1.UnimplementedOrderServiceServer

    od *biz.Ordercase
}

func NewOrderService(od *biz.Ordercase) *OrderServiceImpl {
    return &OrderServiceImpl{od: od}
}

func (s *OrderServiceImpl) SettleOrder(ctx context.Context, in *v1.SettleOrderRequest) (*v1.SettleOrderReply, error) {
    err := s.od.SettleOrderHandler(ctx, in)
    if err != nil {
        return &v1.SettleOrderReply{RetMsg: err.Error()}, err
    }
    return &v1.SettleOrderReply{RetCode: 200, RetMsg: "settle successfully : " + "useMobile from ctx", Data: nil}, nil
}
