package main

import (
	myRouter "gin_02_01/router"
	"github.com/gin-gonic/gin"
)

// gin 框架的路由文件分组
// 所有的 路由文件都在 /router 文件夹下面
func main07() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 注册分文件路由
	myRouter.DefaultgoRoutesInit(router)
	myRouter.AdminRoutesInit(router)
	myRouter.UserRoutesInit(router)
	myRouter.RoleRoutesInit(router)

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run("localhost:8090")
}
