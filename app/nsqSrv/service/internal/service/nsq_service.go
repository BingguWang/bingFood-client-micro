package service

import (
    "context"
    "github.com/BingguWang/bingfood-client-micro/api/nsqSrv/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/biz"
)

type NsqServiceImpl struct {
    v1.UnimplementedNsqServiceServer

    nc *biz.NsqCase
}

func NewNsqService(nc *biz.NsqCase) *NsqServiceImpl {
    return &NsqServiceImpl{nc: nc}
}
func (s *NsqServiceImpl) PubUnPayOrderToMQ(ctx context.Context, in *v1.PubUnPayOrderToMQRequest) (*v1.PubUnPayOrderToMQReply, error) {
    if err := s.nc.PubUnPayOrderToMQHandler(ctx, in); err != nil {
        return nil, err
    }
    return &v1.PubUnPayOrderToMQReply{RetCode: 200, RetMsg: "call PubUnPayOrderToMQ successfully"}, nil
}
