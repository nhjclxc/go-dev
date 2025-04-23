package main

import (
	myRouter "gin_02_01/router2"
	"github.com/gin-gonic/gin"
)

// 八、Gin中自定义控制器
func main() {

	// 8.1、控制器分组
	// 当我们的项目比较大的时候有必要对我们的控制器进行分组
	//新建 controller/admin/NewsController.go

	// 创建路由器
	router := gin.Default()

	// 注册路由
	myRouter.NewsRoutesInit(router)
	myRouter.UserRoutesInit(router)

	// 启动
	router.Run(":8090")

}
