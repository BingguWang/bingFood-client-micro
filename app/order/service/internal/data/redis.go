package data

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/order/service/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis/v8"
    "time"
)

var settleTimout = 10 * time.Minute

func (o *orderRepo) RedisSetKV(ctx context.Context, key, val string) (string, error) {
    log.Infof("insert to redis , key :%v, val :%v", key, val)
    retData, err := o.data.redisCli.Set(ctx, key, val, settleTimout).Result() // 停留在结算页面没操作超过10分钟结算就作废
    if err != nil {
        return "", err
    }
    return retData, nil
}

func (o *orderRepo) RedisGetVal(ctx context.Context, key string) (string, error) {
    log.Infof("get val from redis , key :%v", key)
    retData, err := o.data.redisCli.Get(ctx, key).Result() // 停留在结算页面没操作超过10分钟结算就作废
    if err != nil {
        return "", err
    }
    return retData, nil
}

func (o *orderRepo) RedisDelKey(ctx context.Context, key string) (int64, error) {
    log.Infof("del key from redis , key :%v", key)
    retData, err := o.data.redisCli.Del(ctx, key).Result() // 停留在结算页面没操作超过10分钟结算就作废
    if err != nil {
        return 0, err
    }
    return retData, nil
}

// *conf.Data就是数据的配置结构体
func NewRedis(c *conf.Data) (*redis.Client, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     c.Redis.Addr,
        Password: c.Redis.Password,  // no password set
        DB:       (int)(c.Redis.Db), // 0 means to use default DB
    })
    pong, err := client.Ping(context.Background()).Result()
    if err != nil {
        return nil, v1.ErrorInternal("redis connect ping failed, err: %v", err)
    } else {
        log.Infof("redis connect ping response:%v \"pong\"", pong)
        return client, nil
    }
}
