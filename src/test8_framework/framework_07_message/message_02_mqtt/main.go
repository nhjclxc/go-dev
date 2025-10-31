package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初识 gin 框架
func main() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	//// Broker 地址（tcp://host:port）
	//mqttCfg := mqttcore.MqttConfig{
	//	Broker:    "tcp://localhost:1883",
	//	ClientId:  "go-client-",
	//	Username:  "admin",
	//	Password:  "public",
	//	SubTopics: []string{"/test/sub"},
	//}
	//
	//client := mqttcore.NewMqttClient(&mqttCfg)
	//
	//client.Publish("/test/send", 0, "who are you?")

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "我是第 %d 个 %s 应用哦 \n", 1, "gin")
	})

	// go run main.go

	//启动端口监听
	// 默认是：0.0.0.0:8080
	router.Run(":8090")
}
