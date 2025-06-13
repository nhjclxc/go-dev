package svc

import (
	"go_zero_15_rabbitmq/common/rabbitmq"
	"go_zero_15_rabbitmq/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMQ *rabbitmq.RabbitMQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		RabbitMQ: rabbitmq.NewRabbitMQ(c.RabbitMQConf.Url),
	}
}
