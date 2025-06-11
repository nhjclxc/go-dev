package svc

import (
	"fmt"
	"go_zero_14_mqtt/common/mqtt"
	"go_zero_14_mqtt/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	MqttClient *mqtt.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Printf("%#v", c)
	return &ServiceContext{
		Config:     c,
		MqttClient: mqtt.NewMqttClient(c.MqttConf),
	}
}
