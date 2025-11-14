package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sd02_etcd"
	"time"
)

type Service struct {
	serviceId string
	grant     *clientv3.LeaseGrantResponse
	cancel    context.CancelFunc
}

func (s *Service) KeepAlive() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	fmt.Printf("%s 向服务端发送心跳 \n", s.serviceId)
	aliveCh, err := sd02_etcd.EtcdClient.KeepAlive(ctx, s.grant.ID)
	if err != nil {
		fmt.Printf("%s 发送心跳失败: %s", s.serviceId, err)
		return
	}

	go func() {
		for ch := range aliveCh {
			fmt.Printf("%v 续租成功：%d \n", time.Now(), ch.ID)
		}

	}()
}

type ServiceManagement struct {
	serviceMap map[string]*Service
}

func main() {

	e := gin.Default()

	sm := ServiceManagement{serviceMap: make(map[string]*Service)}

	// http://127.0.0.1:8090/online/service1
	e.GET("/online/:serviceId", func(ginCtx *gin.Context) {
		serviceId := ginCtx.Param("serviceId")

		// 1、创建租约
		grant, err := sd02_etcd.EtcdClient.Grant(context.Background(), 3)
		if err != nil {
			fmt.Println("创建租约失败：", err)
			ginCtx.JSON(200, gin.H{
				"message": "创建租约失败：" + err.Error(),
			})
			return
		}
		// 2、服务上线绑定租约
		putResp, err := sd02_etcd.EtcdClient.Put(context.Background(), serviceId, serviceId, clientv3.WithLease(grant.ID))
		if err != nil {
			fmt.Println("服务上线失败：", err)
			ginCtx.JSON(200, gin.H{
				"message": "服务上线失败：" + err.Error(),
			})
			return
		}
		fmt.Println("putResp.Header.Revision", putResp.Header.Revision)

		s := Service{
			serviceId: serviceId,
			grant:     grant,
		}
		// 3、向服务端发送心跳
		s.KeepAlive()
		sm.serviceMap[serviceId] = &s

		ginCtx.JSON(200, gin.H{"message": serviceId + " 上线成功"})

	})

	e.GET("offline/:serviceId", func(ginCtx *gin.Context) {
		serviceId := ginCtx.Param("serviceId")
		service, ok := sm.serviceMap[serviceId]
		if !ok {
			ginCtx.JSON(200, gin.H{"message": serviceId + " 未上线"})
			return
		}

		// 停止续租
		if service.cancel != nil {
			service.cancel()
		}

		// 看看key在不在
		getResp, _ := sd02_etcd.EtcdClient.Get(context.Background(), serviceId)
		fmt.Println("取消租约前：", getResp.Count)
		for index, keyValue := range getResp.Kvs {
			fmt.Println(index, string(keyValue.Value))
		}

		// 取消租约，同时删除对应的key
		_, err := sd02_etcd.EtcdClient.Revoke(context.Background(), service.grant.ID)
		if err != nil {
			ginCtx.JSON(500, gin.H{"message": "下线失败: " + err.Error()})
			return
		}

		// 看看key还在不在
		getResp2, _ := sd02_etcd.EtcdClient.Get(context.Background(), serviceId)
		fmt.Println("取消租约前：", getResp2.Count)
		for index, keyValue := range getResp2.Kvs {
			fmt.Println(index, string(keyValue.Value))
		}

		delete(sm.serviceMap, serviceId)
		ginCtx.JSON(200, gin.H{"message": serviceId + " 下线成功"})
	})

	e.Run(":8090")

}
