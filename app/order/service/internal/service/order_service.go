package service

import (
    "context"
    "github.com/go-kratos/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

type OrderServiceImpl struct {
    v1.UnimplementedOrderServiceServer

    od *biz.Ordercase
}

func NewOrderService(od *biz.Ordercase) *OrderServiceImpl {
    return &OrderServiceImpl{od: od}
}

func (s *OrderServiceImpl) SettleOrder(ctx context.Context, in *v1.SettleOrderRequest) (*v1.SettleOrderReply, error) {
    retData, err := s.od.SettleOrderHandler(ctx, in)
    if err != nil {
        return &v1.SettleOrderReply{RetMsg: err.Error()}, err
    }
    data := retData.(*v1.SettleOrderReply_Data)
    log.Infof("data : %v", utils.ToJsonString(retData))
    return &v1.SettleOrderReply{RetCode: 200, RetMsg: "settle successfully : " + "useMobile : " + in.UserMobile, Data: data}, nil
}
