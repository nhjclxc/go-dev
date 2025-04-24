package main

import (
	myRouter "gin_02_03/router"
	"github.com/gin-gonic/gin"
)

// 八、Gin中自定义控制器
func main() {

	// 创建路由器
	router := gin.Default()

	// 注册路由
	myRouter.UserRoutesInit(router)

	// 启动
	router.Run(":8090")

}
