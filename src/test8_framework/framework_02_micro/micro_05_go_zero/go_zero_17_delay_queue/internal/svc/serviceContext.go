package svc

import (
	"github.com/zeromicro/go-queue/dq"
	"go_zero_17_delay_queue/common/utils/redisutils"
	"go_zero_17_delay_queue/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	RedisCache *redisutils.RedisCache // 私有 redis 实例
	DqProducer dq.Producer
	DqConsumer dq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RedisCache: redisutils.NewRedisCache(c.Redis),
		DqProducer: dq.NewProducer(c.DqConf.Beanstalks),
		DqConsumer: dq.NewConsumer(c.DqConf),
	}
}
