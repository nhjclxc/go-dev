package main

import (
	"flag"
	"fmt"
	"gin_docker/config"
	"gin_docker/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// 初识 gin 框架
func main() {

	var configFile string

	// 允许传入配置文件路径
	flag.StringVar(&configFile, "c", "", "choose config file")
	flag.Parse()

	if configFile == "" {
		// 尝试从环境变量读取
		configFile = os.Getenv("APP_CONFIG")
		if configFile == "" {
			// 默认使用 dev 配置
			configFile = "config/config-dev.yaml"
		}
	}

	var config *config.Config = config.InitConfig(configFile)

	fmt.Printf("读取到的配置文件： %#v \n", config)

	// 初始化日志
	core.InitLogger()


	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "我是第 %d 个 %s 应用哦 \n", 1, "gin")
	})
	router.GET("/news", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")
		context.String(http.StatusOK, "我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v，使用makefile构建2213 \n", config.Env, id, name)
	})
	router.POST("/insert", func(context *gin.Context) {
		context.String(http.StatusOK, "我是 insert \n")
	})
	router.PUT("/update", func(context *gin.Context) {
		context.String(http.StatusOK, "我是 update \n")
	})
	router.DELETE("/delete", func(context *gin.Context) {
		context.String(http.StatusOK, "我是 delete \n")
	})

	// go run main.go
	// go run main.go -c config/config-dev.yaml
	// go run main.go -c config/config-test.yaml
	// go run main.go -c config/config-prod.yaml

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run(":" + strconv.Itoa(config.Port))
	//router.Run("localhost:8090")
}

/*
docker run -d \
  -v $(pwd)/config:/app/config \
  -v /home/logs/gin-app:/app/logs \
  -p 8081:8081 \
  --name gin-dev \
  gin-app ./server -c config/config-dev.yaml

 */