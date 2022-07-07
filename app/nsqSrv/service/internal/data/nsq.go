package data

import (
    "github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/conf"
    "github.com/nsqio/go-nsq"
)

// *conf.Data就是数据的配置结构体
func NewNsqProducer(c *conf.NSQ) (*nsq.Producer, error) {
    config := nsq.NewConfig()
    producer, err := nsq.NewProducer(c.Addr, config)
    if err != nil {
        return nil, err
    }
    return producer, nil
}
