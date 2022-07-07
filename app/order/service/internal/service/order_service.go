package service

import (
    "context"
    "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/utils"
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
    return &v1.SettleOrderReply{RetCode: 200, RetMsg: "call SettleOrder successfully : " + "useMobile : " + in.UserMobile, Data: data}, nil
}

func (s *OrderServiceImpl) SubmitOrder(ctx context.Context, in *v1.SubmitOrderRequest) (*v1.SubmitOrderReply, error) {
    orderNumber, err := s.od.SubmitOrderHandler(ctx, in)
    if err != nil {
        return nil, err
    }
    return &v1.SubmitOrderReply{RetCode: 200, RetMsg: "call SubmitOrder successfully", OrderNumber: orderNumber}, nil
}

func (s *OrderServiceImpl) PayOrder(ctx context.Context, in *v1.PayOrderRequest) (*v1.PayOrderReply, error) {
    ret, payNo, err := s.od.PayOrderHandler(ctx, in)
    if err != nil {
        return nil, err
    }
    return &v1.PayOrderReply{RetCode: 200, RetMsg: "call PayOrder successfully", WxPayMpOrderResult: ret, PayNo: payNo}, nil
}

func (s *OrderServiceImpl) PayOrderSuccess(ctx context.Context, in *v1.PayOrderSuccessRequest) (*v1.PayOrderSuccessReply, error) {
    if err := s.od.PayOrderSuccessHandler(ctx, in); err != nil {
        return nil, err
    }
    return &v1.PayOrderSuccessReply{RetCode: 200, RetMsg: "call PayOrderSuccess successfully"}, nil
}
