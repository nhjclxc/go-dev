package example

import (
	"fmt"
	"message_02_mqtt/mqttcore"
	"testing"
)

func TestName(t *testing.T) {

	// Broker 地址（tcp://host:port）
	mqttCfg := mqttcore.MqttConfig{
		Broker:    "tcp://localhost:1883",
		ClientId:  "go-client-",
		Username:  "admin",
		Password:  "public",
		SubTopics: []string{"/test/sub"},
	}

	client := mqttcore.NewMqttClient(&mqttCfg)

	err := client.Publish("/test/send", 0, "who are you?")
	if err != nil {
		fmt.Println("q", err)
		return
	}

}
