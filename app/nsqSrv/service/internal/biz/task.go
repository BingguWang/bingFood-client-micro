package biz

import (
    "context"
    "fmt"
    v1 "github.com/BingguWang/bingfood-client-micro/api/order/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/nsqio/go-nsq"
    "os"
    "os/signal"
    "syscall"
)

func UnPayTask() {
    config := nsq.NewConfig()

    consumer, err := nsq.NewConsumer("unPayOrder", "ch1", config)
    if err != nil {
        fmt.Println("create consumer failed :", err.Error())
        return
    }
    // 开启10个协程去处理这个消费者的消费工作
    consumer.AddConcurrentHandlers(&UnPayTaskHandler{internal.GlobalOc}, 20)
    // consumer.AddHandler(&myMessageHandler{})

    // ConnectToNSQLookupd会循环监听一直消费
    if err := consumer.ConnectToNSQLookupd("1.14.163.5:4161"); err != nil {
        log.Fatal(err)
    }

    // wait for signal to exit
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan

    // Gracefully stop the consumer.
    consumer.Stop()
}

type UnPayTaskHandler struct {
    oc v1.OrderServiceClient
}

func NewUnPayTaskHandler(oc v1.OrderServiceClient) *UnPayTaskHandler {
    return &UnPayTaskHandler{oc: oc}
}

// HandleMessage implements the Handler interface.
func (h *UnPayTaskHandler) HandleMessage(m *nsq.Message) error {
    if len(m.Body) == 0 {
        // Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
        // In this case, a message with an empty body is simply ignored/discarded.
        return nil
    }

    // do whatever actual message processing is desired
    if err := h.processMessage(m.Body); err != nil {
        log.Errorf("handler err : %v", err)
        return err
    }

    // Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
    return nil
}
func (h *UnPayTaskHandler) processMessage(msg []byte) error {
    log.Infof("msg is : %v", string(msg))

    orderNumber := string(msg)

    // 调用order 服务执行取消订单逻辑
    fmt.Println("oc === nil ? ", h.oc == nil)
    log.Infof("开始调用服务bingfood.order.service/PayOrderTimeout")
    if _, err := h.oc.PayOrderTimeout(context.TODO(), &v1.PayOrderTimeoutRequest{OrderNumber: orderNumber}); err != nil {
        return err
    }
    log.Infof("成功调用服务bingfood.order.service/PayOrderTimeout")
    return nil
}
