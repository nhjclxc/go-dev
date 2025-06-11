package message

import (
	"fmt"
)

func HandleMessage(topic string, payload []byte) {
	fmt.Printf("接收到来自 [%s] 的消息：%s", topic, string(payload))

	// 可以根据 topic 决定做什么，比如：
	if topic == "sensor/data" {

		//var data SensorData
		//json.Unmarshal(payload, &data)
		// 保存到数据库、处理告警等...
	}
}
