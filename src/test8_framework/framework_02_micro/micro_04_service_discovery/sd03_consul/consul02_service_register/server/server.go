package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"golang.org/x/sync/errgroup"
	"log"
	"sd03_consul"
	"time"
)

type Server struct {
	ctx context.Context
}

func (s *Server) Start() {

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	group, _ := errgroup.WithContext(ctx)
	defer cancel()

	group.Go(func() error {
		// 使用ConsulClient.Health().Service做服务发现

		for range time.Tick(3 * time.Second) {
			//| 参数         | 当前值   | 含义                                        |
			//| ----------- | ------- | ----------------------------------------- |
			//| service     | `"web"` | 查询服务名为 `"web"` 的节点                        |
			//| tag         | `""`    | 不按 tag 过滤，获取所有 `"web"` 服务                 |
			//| passingOnly | `true`  | 只返回 **健康节点**，自动过滤掉 unhealthy 节点           |
			//| q           | `nil`   | 使用默认查询参数，没有 blocking query 或特殊 datacenter |
			services, _, err := sd03_consul.ConsulClient.Health().Service("web", "", true, nil)
			if err != nil {
				log.Fatal(err)
			}

			if len(services) == 0 {
				fmt.Println("服务 web 没有健康节点")
			} else {
				for _, s := range services {
					fmt.Printf("Node: %s, Address: %s:%d\n",
						s.Node.Node, s.Service.Address, s.Service.Port)
				}
				fmt.Println()
			}
		}

		return nil
	})

	group.Go(func() error {
		// 使用ConsulClient.Health().Service做服务监听

		// 当监听的服务name里面有实例上线或下线时就会触发

		var lastIndex uint64 = 0
		name := "web"
		for {
			services, meta, err := sd03_consul.ConsulClient.Health().Service(name, "", true, &api.QueryOptions{
				WaitIndex: lastIndex,
				WaitTime:  2 * time.Minute,
			})
			if err != nil {
				fmt.Println("watch error:", err)
				continue
			}

			if meta.LastIndex == lastIndex {
				continue
			}

			lastIndex = meta.LastIndex
			fmt.Println("service changed! new list:", services, time.Now().Format("15:04:05"))
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("服务遇到错误退出", err)
		return
	}

}
