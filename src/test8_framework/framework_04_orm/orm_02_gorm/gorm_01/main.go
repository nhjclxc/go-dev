package main

import (
	"github.com/gin-gonic/gin"
	"gorm_01/middleware"
	myRouter "gorm_01/router"
)

// GORM 增删改查 的使用
func main() {

	// 创建路由器
	router := gin.Default()

	// 注册全局中间件
	router.Use(middleware.GlobalPainc)

	// 注册路由
	myRouter.GenTable2RoutesInit(router)

	// 启动
	router.Run(":8090")

}
