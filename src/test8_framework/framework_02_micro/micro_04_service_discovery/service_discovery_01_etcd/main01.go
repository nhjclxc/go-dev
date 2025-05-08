package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

// go 操作 etcd
// https://topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cetcd/%E6%93%8D%E4%BD%9Cetcd.html

// https://pkg.go.dev/go.etcd.io/etcd/client/v3?utm_source=chatgpt.com

// go get go.etcd.io/etcd/clientv3
func main01() {


	// etcd client put/get demo
	// use etcd/clientv3

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.106.59.225:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
