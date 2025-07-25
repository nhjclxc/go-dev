package main

import (
	"gin_generator/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// gin 代码生成

// 代码生成项目地址：https://github.com/nhjclxc/generator

// @title 示例 API
// @version 1.0
// @description 使用 Swagger 演示 token 请求头设置
// @BasePath /
// @host localhost:8080
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// 启用跨域支持
	r.Use(cors.Default())

	// 路由分组，所有需要鉴权的接口用 AuthMiddleware 包裹
	privateGroup := r.Group("/private")
	publicGroup := r.Group("/public")

	(&router.GenTableRouter{}).InitGenTableRouter(privateGroup, publicGroup)
	(&router.SysJobRouter{}).InitSysJobRouter(privateGroup, publicGroup)
	(&router.SysRoleRouter{}).InitSysRoleRouter(privateGroup, publicGroup)

	r.Run(":8080")
}
