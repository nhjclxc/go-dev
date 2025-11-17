package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"sd03_consul"
	"strconv"
)

func main() {

	// go run client.go -id=web1 -port=8091

	id := flag.String("id", "web", "服务名称")
	addr := flag.String("addr", "192.168.201.167", "服务所在的ip")
	portStr := flag.String("port", "8091", "服务所在的端口")

	// 解析命令行参数
	flag.Parse()

	port, _ := strconv.Atoi(*portStr)

	e := gin.Default()

	e.GET("/health", func(c *gin.Context) {
		fmt.Printf("%s心跳检测接口被执行 \n", *id)
		c.JSON(200, gin.H{"msg": "successful"})
	})

	e.GET("/down", func(c *gin.Context) {
		// 服务下线
		err := sd03_consul.ConsulClient.Agent().ServiceDeregister(*id)
		if err != nil {
			fmt.Println("服务下线失败：", err)
			return
		}
		fmt.Println("服务下线成功！")
	})

	// 将这个服务注册到conslu

	// 构造服务信息
	reg := &api.AgentServiceRegistration{
		ID:      "web:" + *id,
		Name:    "web",
		Address: *addr,
		Port:    port,
		Tags:    []string{"v1"},
		// api.AgentServiceCheck 服务的健康检测接口，这个接口返回的数据格式无要求
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + *addr + ":" + *portStr + "/health",
			Interval: "5s",
			Timeout:  "3s",
		},
	}

	// 注册
	err := sd03_consul.ConsulClient.Agent().ServiceRegister(reg)
	if err != nil {
		fmt.Printf("%s 注册服务失败：%s", *id, err)
		return
	}

	// 启动gin的http服务
	fmt.Printf("%s:%s 启动成功\n", *addr, *portStr)
	e.Run(":" + *portStr)

}
