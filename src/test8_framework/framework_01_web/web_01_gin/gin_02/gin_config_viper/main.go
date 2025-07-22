package main

import (
	"github.com/gin-gonic/gin"
	"gin_config_viper/core"
	"net/http"
	"strconv"

	localConfig "gin_config_viper/global"
)

func main() {

	localConfig.GlobalViper = core.Viper() // 初始化Viper

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/logs", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")

		context.String(http.StatusOK, "我是 news 页面，当前环境是：%s，当前配置为：%#v，现在请求的文章id = %v, name = %v \n", localConfig.GlobalConfig.Name, localConfig.GlobalConfig, id, name)
	})

	// go run main.go
	// go run main.go -c config-dev.yaml
	// go run main.go -c config-test.yaml
	// go run main.go -c config-prod.yaml

	// http://localhost:8080/logs?id=666&name=HelloGolang

	router.Run(":" + strconv.Itoa(localConfig.GlobalConfig.Port))

}
