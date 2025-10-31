package main

import (
	"android06_kill_app/kill_app_mqtt/dto"
	"android06_kill_app/kill_app_mqtt/mqttcore"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	ctx        context.Context
	mqttClient *mqttcore.MqttClient
	ginEngine  *gin.Engine
	server     *http.Server
	// 用于实时的将执行结果回传给前端
	respWaitPool *sync.Map // deviceId -> chan dto.KillAppCmdResp
}

func NewServer(ctx context.Context, mqttCfg *mqttcore.MqttConfig, addr string) *Server {

	var mqttClient *mqttcore.MqttClient = mqttcore.NewMqttClient(mqttCfg)

	ginEngine := gin.Default()
	s := Server{
		ctx:          ctx,
		mqttClient:   mqttClient,
		ginEngine:    ginEngine,
		respWaitPool: &sync.Map{},
		server: &http.Server{
			Addr:    addr,
			Handler: ginEngine,
		},
	}

	s.setRouter()
	return &s
}

func (s Server) setRouter() {
	// 路由绑定
	// http://127.0.0.1:8090/cmd/CZ2504020139/kill/com.feedying.live.mix,cn.miguvideo.migutv,cn.juqing.cesuwang_tv
	// http://127.0.0.1:8090/deviceId/kill/pkgs
	s.ginEngine.GET("cmd/:deviceId/kill/:pkgs", func(c *gin.Context) {
		deviceId := c.Param("deviceId")
		pkgs := c.Param("pkgs")
		if deviceId == "" {
			c.String(400, "deviceId 不能为空")
		}
		if pkgs == "" {
			c.String(400, "pkgs 不能为空")
		}

		// 1️⃣ 创建一个等待响应的通道
		respChan := make(chan dto.KillAppCmdResp, 1)
		s.respWaitPool.Store(deviceId, respChan)
		defer s.respWaitPool.Delete(deviceId)

		// 2️⃣ 发布MQTT命令
		var req = dto.KillAppCmdReq{PackageList: strings.Split(pkgs, ",")}
		msgByte, _ := json.Marshal(&req)
		err := s.mqttClient.PublishWithAck("cmd/"+deviceId+"/kill", 1, false, msgByte)
		if err != nil {
			c.String(500, fmt.Sprintf("发送消息出现错误: %s", err))
			return
		}

		// 3️⃣ 阻塞等待回调或超时
		select {
		case resp := <-respChan:
			c.JSON(http.StatusOK, resp)
		case <-time.After(5 * time.Second):
			c.String(504, "等待设备响应超时")
		}

		// todo 如何获取到mqtt到回调消息，之后实时返回给前端

		//c.String(http.StatusOK, "我是第 %d 个 %s 应用哦 \n", 1, "gin")
	})
}

func (s Server) StartAll() error {

	// 监听mqtt的消息
	go s.mqttClient.Run(s.ctx)

	//启动端口监听
	return s.server.ListenAndServe()
}

func (s Server) StopAll() {

	// 优雅关闭
	s.mqttClient.Close()

	if s.server != nil {
		s.server.Shutdown(context.Background())
	}

	log.Println("✅ 优雅退出")
}

func main() {
	// go run main.go

	// 服务器下发哪些app要被kill
	// +表示deviceId
	// cmd/+/kill，服务端通知#指定的设备进行杀进程，对应的回调topic：cmd/+/kill/callback
	// cmd/+/process，服务端请求#指定的设备的当前运行的进程信息，对应的回调topic：cmd/+/process/callback
	// cmd/+/flow，服务端通知#指定的设备进行流量统计，对应的回调topic：cmd/+/flow/callback
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var s *Server

	callBackTopic := "cmd/+/kill/callback"
	subscribeTopics := make([]*mqttcore.SubscribeTopic, 0)
	subscribeTopics = append(subscribeTopics, &mqttcore.SubscribeTopic{
		Topic: callBackTopic,
		Callback: func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("[Message - server - 消息回调处理器正在处理消息] topic=%s payload=%s", msg.Topic(), msg.Payload())
			var resp = dto.KillAppCmdResp{}
			json.Unmarshal(msg.Payload(), &resp)
			fmt.Printf("json序列化消息：%#v \n", resp)

			slist := strings.Split(msg.Topic(), "/")
			deviceId := slist[1]
			if chVal, ok := s.respWaitPool.Load(deviceId); ok {
				ch := chVal.(chan dto.KillAppCmdResp)
				select {
				case ch <- resp:
				default:
					log.Printf("⚠️ 设备[%s]的回调已过期或已消费", deviceId)
				}
			}
		},
	})

	// Broker 地址（tcp://host:port）
	mqttCfg := mqttcore.MqttConfig{
		Broker:          "tcp://localhost:1883",
		ClientId:        "killapp-server-",
		Username:        "admin",
		Password:        "public",
		SubscribeTopics: subscribeTopics,
	}

	s = NewServer(ctx, &mqttCfg, ":8090")

	// ✅ 1️⃣ 在独立 goroutine 启动服务
	go func() {
		if err := s.StartAll(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ HTTP 服务异常退出: %v", err)
		}
	}()

	// ✅ 2️⃣ 主 goroutine 等待信号
	// 捕获退出信号（Ctrl+C 或 docker stop）
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("🛑 收到退出信号，正在关闭...")

	// ✅ 3️⃣ 触发取消并优雅退出
	cancel()
	s.StopAll()

}
