package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	RabbitMQConf RabbitMQ `json:"RabbitMQConf"` // RabbitMQ 配置
}

// RabbitMQ 配置
type RabbitMQ struct {
	Url string `json:"url"`
}
