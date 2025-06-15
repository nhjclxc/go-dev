package svc

import (
	"context"
	redisutils "go_zero_18_delay_queue_redis/common/utils/redisutils"
	"go_zero_18_delay_queue_redis/internal/config"
	"go_zero_18_delay_queue_redis/internal/delayQueue"
)

type ServiceContext struct {
	Config     config.Config
	RedisCache *redisutils.RedisCache // 私有 redis 实例
	DelayQueueRedisStream *delayQueue.DelayQueueRedisStream
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := redisutils.NewRedisCache(c.Redis)
	delayQueueRedisStream := delayQueue.NewDelayQueueRedisStream(redisCache)
	// 开始消费
	delayQueueRedisStream.ConsumeTasks(context.Background())
	return &ServiceContext{
		Config:     c,
		RedisCache: redisCache,
		DelayQueueRedisStream: delayQueueRedisStream,
	}
}