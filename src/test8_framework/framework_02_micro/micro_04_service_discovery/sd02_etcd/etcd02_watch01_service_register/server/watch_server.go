package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sd02_etcd"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//defer sd02_etcd.EtcdClient.Close()

	watchKey := "micro-service-" // micro-service-1,micro-service-2,micro-service-3 ...
	watchChan := sd02_etcd.EtcdClient.Watch(ctx, watchKey, clientv3.WithPrefix())

	// 监听服务是否在线
	for watchResponse := range watchChan {
		for _, event := range watchResponse.Events {
			if event.Type == mvccpb.PUT {
				// 微服务上线
				fmt.Printf("[Service] %s is alive \n", event.Kv.Key)
			} else {
				//} else if event.Type == mvccpb.DELETE {
				fmt.Printf("[Service] %s is offline \n", event.Kv.Key)
			}
		}
	}

}
