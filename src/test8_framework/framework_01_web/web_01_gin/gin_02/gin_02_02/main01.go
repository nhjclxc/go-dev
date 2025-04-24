package main

import (
	"gin_02_02/middleware"
	myRouter "gin_02_02/router"
	"github.com/gin-gonic/gin"
)

func main() {

	// 创建 gin 路由器
	router := gin.Default()

	// main01将学习 gin 里面的中间件
	// 中间件一、路由中间件
	// 中间件二、路由分组中间件
	// 中间件三、全局中间件
	// 中间件之间的数据共享

	//可以使用 gin 里面的中间件实现Java里面的拦截器（配合Abort()方法）、过滤器（配合Next()方法）和切面配合Next()方法等等

	// 注册全局中间件
	router.Use(middleware.GlobalPainc, middleware.Authentication)

	// 注册路由
	myRouter.UserRouterInit(router)

	// 启动 gin 项目
	router.Run(":8090")
}
