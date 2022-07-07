package service

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

func (s *BingfoodServiceImpl) OrderSettle(ctx context.Context, in *v1.SettleOrderRequest) (*v1.SettleOrderReply, error) {
    data, err := s.oc.SettleOrder(ctx, in)
    if err != nil {
        return &v1.SettleOrderReply{RetMsg: err.Error()}, err
    }
    log.Infof(utils.ToJsonString(data))
    return &v1.SettleOrderReply{RetCode: 200, RetMsg: "成功结算 : ", Data: data}, nil
}
func (s *BingfoodServiceImpl) OrderSubmit(ctx context.Context, in *v1.SubmitOrderRequest) (*v1.SubmitOrderReply, error) {
    claims := ctx.Value("claims").(*utils.UserClaims)
    in.UserMobile = claims.UserMobile
    orderNumber, err := s.oc.SubmitOrder(ctx, in)
    if err != nil {
        return nil, err
    }
    return &v1.SubmitOrderReply{
        RetCode:     200,
        RetMsg:      "订单提交成功",
        OrderNumber: orderNumber,
    }, nil
}

func (s *BingfoodServiceImpl) OrderPay(cxt context.Context, in *v1.PayOrderRequest) (*v1.PayOrderReply, error) {
    ret, payNo, err := s.oc.PayOrder(cxt, in)
    if err != nil {
        return nil, err
    }
    return &v1.PayOrderReply{RetCode: 200, RetMsg: "支付成功", WxPayMpOrderResult: ret, PayNo: payNo}, nil
}

func (s *BingfoodServiceImpl) NoticePayOrder(ctx context.Context, in *v1.NoticePayOrderRequest) (*v1.NoticePayOrderReply, error) {
    if err := s.oc.NoticePayOrderHandler(ctx, in); err != nil {
        return nil, err
    }
    return &v1.NoticePayOrderReply{RetCode: 200, RetMsg: "支付成功回调操作"}, nil
}
