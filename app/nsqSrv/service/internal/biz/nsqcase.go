package biz

import (
    "context"
    "fmt"
    v1 "github.com/BingguWang/bingfood-client-micro/api/nsqSrv/service/v1/pbgo/v1"
    v3 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/nsqio/go-nsq"
    "time"
)

type NsqCase struct {
    pc  *nsq.Producer
    log *log.Helper
}

func NewNsqCase(pc *nsq.Producer, oc v3.OrderServiceClient, logger log.Logger) *NsqCase {
    internal.GlobalOc = oc
    return &NsqCase{pc: pc, log: log.NewHelper(logger)}
}

func (oc *NsqCase) PubUnPayOrderToMQHandler(ctx context.Context, req *v1.PubUnPayOrderToMQRequest) error {
    oc.log.WithContext(ctx).Infof("PubUnPayOrderToMQHandler args: %v", utils.ToJsonString(req))
    log.Infof("未支付的订单号存入MQ : %v", req.OrderNumber)

    // 3分钟未支付的订单就消费(视为订单取消)掉
    fmt.Println("[[[[[", oc.pc.String())

    if err := oc.pc.DeferredPublish("unPayOrder", 20*time.Second, []byte(req.OrderNumber)); err != nil {
        log.Errorf("推送未支付订单消息到MQ失败 : %v", err.Error())
        return err
    }
    return nil
}
