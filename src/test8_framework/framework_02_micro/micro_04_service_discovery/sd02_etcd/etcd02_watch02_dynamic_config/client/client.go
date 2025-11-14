package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sd02_etcd"
	"sd02_etcd/etcd02_watch02_dynamic_config/config"
	"strconv"
)

// 使用etcd实现动态配置
// 1️⃣ 思路解释
// Server端下发配置
//
//	Server 将配置写入 etcd（PUT key value 或 DELETE key）
//	可以通过 key 命名规范（如 /config/serviceA/db_host）组织不同服务或模块的配置
//
// Client端监听配置
//
//	Client 启动时，先 GET 当前配置，初始化本地缓存
//	使用 Watch 监听对应 key 或 key 前缀
//	当配置被修改或删除时，Watch 会收到事件
//	Client 根据事件更新本地内存配置 → 实现 动态下发
//
// 业务层使用本地缓存
//
//	Client 的业务逻辑直接读取内存中的配置
//	不需要重新访问 etcd
//	当配置更新，业务立即生效

var ErrConfigNotFound = errors.New("配置中心不存在该配置")

type Service struct {
	serviceId           string
	keyPrefix           string
	config              *config.Config
	initConfigReversion int64
}

func (s *Service) initConfig() error {
	databaseHost, err := s.getConfig("databaseHost")
	if err != nil {
		if !errors.Is(err, ErrConfigNotFound) {
			return fmt.Errorf("配置获取失败：%w", err)
		}
		databaseHost = "default DatabaseHost"
	}
	s.config.DatabaseHost = databaseHost

	databasePort, err := s.getConfig("databasePort")
	if err != nil {
		if !errors.Is(err, ErrConfigNotFound) {
			return fmt.Errorf("配置获取失败：%w", err)
		}
		databaseHost = "default DatabasePort"
	}
	s.config.DatabasePort, _ = strconv.Atoi(databasePort)
	return nil
}

func (s *Service) getConfig(serviceName string) (string, error) {
	getResp, err := sd02_etcd.EtcdClient.Get(context.Background(), "config/"+s.serviceId+"/"+serviceName)
	if err != nil {
		fmt.Println("etcd get err", err)
		return "", err
	}
	for _, kv := range getResp.Kvs {
		return string(kv.Value), nil
	}
	return "", ErrConfigNotFound
}

func (s *Service) initConfig2() error {
	getResp, err := sd02_etcd.EtcdClient.Get(context.Background(), s.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("etcd get err", err)
		return err
	}

	s.initConfigReversion = getResp.Header.Revision
	for _, keyValue := range getResp.Kvs {
		err := json.Unmarshal(keyValue.Value, &s.config)
		if err != nil {
			fmt.Printf("json序列化错误：%s \n", err)
			return err
		}
	}

	return nil
}

func (s *Service) printConfig() {
	fmt.Println("配置更新：")
	fmt.Printf("最新配置：%v \n", s.config)
}

// 客户端监听配置变化
func main() {
	serviceId := "order-service"
	s := Service{
		serviceId: serviceId,
		config:    &config.Config{},
		keyPrefix: "config/" + serviceId,
	}

	// 服务启动之后先获取一个配置（初始化配置）
	//err := s.initConfig()
	err := s.initConfig2()
	if err != nil {
		fmt.Println("初始化配置失败：", err)
		return
	}
	s.printConfig()

	// 创建watch
	// "config/"+serviceId
	// , clientv3.WithRev(s.initConfigReversion+1) 确保在初始化到watch这段时间内配置没有被更新过
	watchChan := sd02_etcd.EtcdClient.Watch(context.Background(), s.keyPrefix, clientv3.WithPrefix(), clientv3.WithRev(s.initConfigReversion+1))

	fmt.Println("listening config change ...")
	for watchResponse := range watchChan {
		for _, event := range watchResponse.Events {
			fmt.Printf("配置中心操作类型：%d, 操作配置：%s \n", event.Type, string(event.Kv.Key))
			switch event.Type {
			case clientv3.EventTypePut:
				err := json.Unmarshal(event.Kv.Value, &s.config)
				if err != nil {
					fmt.Printf("序列化错误：%s \n", err)
					return
				}
			case clientv3.EventTypeDelete:
				// 执行一些默认操作
			}
			s.printConfig()
		}
	}

}
