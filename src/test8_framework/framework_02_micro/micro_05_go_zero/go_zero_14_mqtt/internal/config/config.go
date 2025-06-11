package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MqttConf MqttConfig `json:"MqttConf"` // Mqtt 配置
}

// Mqtt 配置类
type MqttConfig struct {
	Broker    string   `json:"Broker"`
	ClientId  string   `json:"ClientId"`
	Username  string   `json:"Username"`
	Password  string   `json:"Password"`
	SubTopics []string `json:"SubTopics"`
}
