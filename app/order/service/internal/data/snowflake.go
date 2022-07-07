package data

import (
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/conf"
    "github.com/bwmarrin/snowflake"
    "github.com/go-kratos/kratos/v2/log"
)

func NewSnowFlakeNode(c *conf.Data) (*snowflake.Node, error) {
    var node *snowflake.Node
    node, err := snowflake.NewNode(1) // 新建一个节点号为1的node
    if err != nil {
        log.Errorf("make snowflake node failed, err: %v", err)
        return nil, err
    }
    return node, nil
}
