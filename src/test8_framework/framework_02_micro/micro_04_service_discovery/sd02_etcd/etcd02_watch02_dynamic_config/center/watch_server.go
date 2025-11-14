package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sd02_etcd/etcd02_watch02_dynamic_config/config"

	"sd02_etcd"
)

func main() {
	// 配置中心修改配置，以便客户端发现配置的变化

	e := gin.Default()

	// http://127.0.0.1:8090/user-service/databaseHost/127.0.0.1
	// http://127.0.0.1:8090/user-service/databasePort/3306
	e.GET("/:serviceId/:configKey/:configValue", func(c *gin.Context) {
		serviceId := c.Param("serviceId")
		configKey := c.Param("configKey")
		configValue := c.Param("configValue")
		putResp, err := sd02_etcd.EtcdClient.Put(context.Background(), "config/"+serviceId+"/"+configKey, configValue)
		if err != nil {
			fmt.Printf("配置下发失败：%s \n", err.Error())
			c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("配置下发失败：%s", err.Error())})
			return
		}
		fmt.Println("操作成功：", putResp.Header.Revision)
		c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("操作成功：%d", putResp.Header.Revision)})
	})

	// http://127.0.0.1:8090/order-service
	e.POST("/:serviceId", func(c *gin.Context) {
		serviceId := c.Param("serviceId")

		var config config.Config
		err := c.ShouldBindBodyWithJSON(&config)
		if err != nil {
			fmt.Println("json序列化错误：", err)
			return
		}
		jsonByte, _ := json.Marshal(config)

		putResp, err := sd02_etcd.EtcdClient.Put(context.Background(), "config/"+serviceId, string(jsonByte))
		if err != nil {
			fmt.Printf("配置下发失败：%s \n", err.Error())
			c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("配置下发失败：%s", err.Error())})
			return
		}
		fmt.Println("操作成功：", putResp.Header.Revision)

		c.JSON(200, gin.H{"code": 200, "msg": fmt.Sprintf("操作成功：%d", putResp.Header.Revision)})
	})

	e.Run(":8090")
}
