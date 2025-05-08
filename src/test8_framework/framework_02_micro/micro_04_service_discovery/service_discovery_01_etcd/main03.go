package main

import (
	"fmt"
	"time"
)

// lease租约

import (
	"context"
	"log"

	"go.etcd.io/etcd/client/v3"
)

func main03() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.106.59.225:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd success.")
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	//resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /lmh/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/lmh/", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}