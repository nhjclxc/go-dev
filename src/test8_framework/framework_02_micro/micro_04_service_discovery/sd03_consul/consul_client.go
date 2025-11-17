package sd03_consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// go get github.com/hashicorp/consul/api
// docs: https://developer.hashicorp.com/consul/docs

var ConsulClient *api.Client

func init() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		fmt.Println("创建consul失败：", err)
		return
	}
	ConsulClient = client
}
