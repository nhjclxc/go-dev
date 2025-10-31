package mqttcore

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// MqttClient mqtt客户端
type MqttClient struct {
	client mqtt.Client
	cfg    *MqttConfig
	mu     sync.Mutex // 确保并发安全（发布时）
}

// MqttConfig 配置类
type MqttConfig struct {
	Broker          string            `json:"broker"`
	ClientId        string            `json:"clientId"`
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	CleanSession    bool              `json:"cleanSession"`
	QoS             byte              `json:"QoS"`
	SubscribeTopics []*SubscribeTopic `json:"subTopics"`
	EnableTLS       bool              `json:"enableTLS"`
	TLSCaFilePath   string            `json:"tLSCaFilePath"`
}

// SubscribeTopic topic订阅
type SubscribeTopic struct {
	Topic    string              `json:"broker"`
	Callback mqtt.MessageHandler `json:"-"`
}

func NewMqttClient(cfg *MqttConfig) *MqttClient {
	mc := MqttClient{
		cfg: cfg,
	}

	opts := mc.setClientOptions(cfg)

	// 启动客户端
	client := mqtt.NewClient(opts)
	token := client.Connect()
	// 尝试ping
	if !token.WaitTimeout(3 * time.Second) {
		log.Println("订阅超时")
	} else if token.Error() != nil {
		log.Printf("订阅失败: %v", token.Error())
	}

	mc.client = client
	return &mc
}

// setClientOptions 配置连接参数
func (mc *MqttClient) setClientOptions(cfg *MqttConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(fmt.Sprintf("%s-%d", cfg.ClientId, time.Now().UnixNano()))
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetCleanSession(cfg.CleanSession)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(3 * time.Second)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetWriteTimeout(5 * time.Second)
	// 离线重连后会自动重新发布未确认的 QoS 1/2 消息。
	//opts.SetStore(mqtt.NewFileStore("/tmp/mqtt_store"))

	// 定义收到消息的回调函数
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("DefaultPublishHandler Received message on [%s]: %s\n", msg.Topic(), msg.Payload())
		mc.handleMessage(msg.Topic(), msg)
	})
	// 链接成功之后会执行的回调
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("✅ 已连接到 EMQX")
		// 有时连接刚建立还没准备好，可以在 OnConnect 中加一点延迟：
		time.Sleep(1 * time.Second)
		mc.subscribeTopics()
	}
	// 链接断开后执行的回调
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("❌ 连接断开: %v", err)
		// 可以根据错误类型判断是网络问题、认证失败、超时等
		if err != nil {
			if strings.Contains(err.Error(), "connection refused") {
				log.Println("连接被拒绝，请检查 Broker 地址或认证信息")
			} else if strings.Contains(err.Error(), "network error") {
				log.Println("网络问题，请检查网络连接")
			} else {
				log.Println("其他原因导致断开:", err)
			}
		}
	}
	// 断开重连
	opts.OnReconnecting = func(c mqtt.Client, opts *mqtt.ClientOptions) {
		log.Println("🔄 正在重连...")
	}
	//// 设置TLS证书
	//opts.SetTLSConfig()
	return opts
}

func (mc *MqttClient) Run(ctx context.Context) {
	// 负责连接和订阅循环（可重连）
	<-ctx.Done()
	mc.Close()
}

// Close 断开mqtt连接
func (mc *MqttClient) Close() {
	if mc.client != nil && mc.client.IsConnected() {
		log.Println("🚪 MQTT 客户端断开连接")
		mc.client.Disconnect(250) // 参数为等待毫秒数
	}
}

// IsConnected 判断mqtt是否已连接
func (mc *MqttClient) IsConnected() bool {
	if mc.client == nil {
		return false
	}
	return mc.client.IsConnected()
}

// subscribeTopics 订阅配置中的所有主题
func (mc *MqttClient) subscribeTopics() {
	for _, topic := range mc.cfg.SubscribeTopics {
		mc.Subscribe(topic)
	}
}

func (mc *MqttClient) Subscribe(subscribe *SubscribeTopic) {
	token := mc.client.Subscribe(subscribe.Topic, mc.cfg.QoS, subscribe.Callback)
	//token := c.Subscribe(topic.Topic, mc.cfg.QoS, func(client mqtt.Client, msg mqtt.Message) {
	//	log.Printf("[Message] topic=%s payload=%s", msg.Topic(), msg.Payload())
	//	mc.handleMessage(msg.Topic(), msg)
	//})
	token.Wait()
	if token.Error() != nil {
		log.Printf("❌ 订阅 [%s] 失败: %v", subscribe.Topic, token.Error())
	} else {
		log.Printf("📡 已订阅主题: %s", subscribe.Topic)
	}
}

// Publish 发布消息
func (mc *MqttClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	if token := mc.client.Publish(topic, qos, retained, payload); token.WaitTimeout(3*time.Second) && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// PublishQos0 发布 Qos0 消息（快捷方法）
func (mc *MqttClient) PublishQos0(topic string, payload interface{}) error {
	return mc.Publish(topic, 0, false, payload)
}

func (mc *MqttClient) PublishWithAck(topic string, qos byte, retained bool, payload interface{}) error {
	token := mc.client.Publish(topic, qos, retained, payload)
	if token.WaitTimeout(3 * time.Second) {
		if err := token.Error(); err != nil {
			log.Printf("❌ 发布失败: %v", err)
			return err
		}
		log.Printf("📨 发布成功: topic=%s", topic)
		return nil
	}
	return fmt.Errorf("publish timeout")
}

// handleMessage 处理接收到的消息
func (mc *MqttClient) handleMessage(topic string, msg mqtt.Message) {
	fmt.Println("------ 收到消息 ------")
	fmt.Printf("主题: %s\n", msg.Topic())
	fmt.Printf("QoS: %d\n", msg.Qos())
	fmt.Printf("消息ID: %d\n", msg.MessageID())
	fmt.Printf("是否保留: %v\n", msg.Retained())
	fmt.Printf("是否重发: %v\n", msg.Duplicate())
	fmt.Printf("内容: %s\n", string(msg.Payload()))

	// 业务逻辑示例
	switch topic {
	case "sensor/data":
		// TODO: JSON 解析、数据库保存、处理逻辑等
	default:
		log.Printf("[未匹配主题] %s -> %s", topic, msg.Payload())
	}
}

// 1) 使用系统证书池（适合 broker 使用公网 CA 签名证书）
func NewTLSConfigSystemCA(serverName string, insecureSkipVerify bool) (*tls.Config, error) {
	// 从系统证书池加载 Root CAs（某些平台需要额外处理）
	rootCAs, err := x509.SystemCertPool()
	if err != nil || rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: insecureSkipVerify, // 测试时可 true，但生产不要
		ServerName:         serverName,         // 用于 SNI / 验证证书 CN/ SAN
		MinVersion:         tls.VersionTLS12,
	}
	return tlsConfig, nil
}

// NewTLSConfigWithCA 2) 使用自签 CA（需要提供 CA 文件）
func NewTLSConfigWithCA(caFile string, serverName string, insecureSkipVerify bool) (*tls.Config, error) {
	caPem, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA file: %w", err)
	}
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(caPem); !ok {
		// 如果 CA 文件不是 PEM 格式或追加失败，需要检查文件
		log.Println("warning: no certs appended, check caFile")
	}
	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: insecureSkipVerify,
		ServerName:         serverName,
		MinVersion:         tls.VersionTLS12,
	}
	return tlsConfig, nil
}

// 3) 双向（客户端证书 + CA）
func NewTLSConfigWithClientCert(caFile, certFile, keyFile, serverName string, insecureSkipVerify bool) (*tls.Config, error) {
	// load CA
	caPem, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA file: %w", err)
	}
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(caPem); !ok {
		log.Println("warning: no certs appended from CA pem")
	}

	// load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("load x509 key pair: %w", err)
	}

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: insecureSkipVerify,
		ServerName:         serverName, // 必填以校验 Broker 证书（除非你明确跳过）
		MinVersion:         tls.VersionTLS12,
	}
	return tlsConfig, nil
}
