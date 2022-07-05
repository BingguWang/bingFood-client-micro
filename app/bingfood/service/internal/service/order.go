package service

import (
    "context"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"

    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
)

func (s *BingfoodServiceImpl) OrderSettle(ctx context.Context, in *v1.SettleOrderRequest) (*v1.SettleOrderReply, error) {
    reply, err := s.oc.SettleOrder(ctx, in)
    if err != nil {
        return &v1.SettleOrderReply{RetMsg: err.Error()}, err
    }
    log.Infof(utils.ToJsonString(reply.Data))
    var data v1.SettleOrderReply_Data
    copier.CopyWithOption(&data, reply.Data, copier.Option{
        DeepCopy:    true,
        IgnoreEmpty: true,
    })
    log.Infof(utils.ToJsonString(data))

    return &v1.SettleOrderReply{RetCode: 200, RetMsg: "成功结算 : ", Data: &data}, nil
}
