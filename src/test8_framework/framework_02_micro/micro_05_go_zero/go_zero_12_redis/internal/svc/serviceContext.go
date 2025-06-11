package svc

import (
	"go_zero_12_redis/common/utils/redisutils"
	"go_zero_12_redis/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	RedisCache    *redisutils.RedisCache // 私有 redis 实例
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisCache: redisutils.NewRedisCache(c.Redis),
	}
}
