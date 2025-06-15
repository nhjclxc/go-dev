package config

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis  redis.RedisConf
	DqConf dq.DqConf
	//DqConf struct {
	//	// Beanstalks: 多个 Beanstalk 节点配置
	//	Beanstalks []Beanstalk `json:"beanstalks"`
	//
	//	// Redis：redis 配置，主要在这里面使用 Setnx 去重
	//	Redis redis.RedisConf `json:"redis"`
	//} `json:"dqConf"`
}

type Beanstalk struct {
	Endpoint string `json:"endpoint"`
	Tube     string `json:"tube"`
}
