package main

import (
	myRouter "gin_02_03/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 八、Gin中自定义控制器
func main01() {

	// 创建路由器
	router := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")

	// 仅输出日志文件
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 注册路由
	myRouter.UserRoutesInit(router)

	// 启动
	router.Run(":8090")

}
