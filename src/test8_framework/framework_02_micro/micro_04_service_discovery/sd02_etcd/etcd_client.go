package sd02_etcd

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// clientv3.Client内部有状态，因此clientv3.Client可以被多次使用，也可以被多个go协程并发使用
var etcdClient *clientv3.Client
var EtcdClient *clientv3.Client

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
		if errors.Is(err, context.DeadlineExceeded) {
			// handle errors
			fmt.Printf("handle errors, err:%v\n", err)
			return
		}
		return
	}
	fmt.Println("connect to etcd success")

	etcdClient = cli
	EtcdClient = cli
}
