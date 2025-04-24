package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 框架的 路由组
// https://gin-gonic.com/zh-cn/docs/examples/grouping-routes/
func main06() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由组将所有的 API 按照一颗前缀树来构建
	// 树根是 IP:Port
	// 接着子树是一级路由组
	// 下级还可以是路由组（路由组下面可以有路由组），或是具体路由

	// 创建一个路由组
	v1 := router.Group("/v1")
	{
		v1.GET("/getInfo", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})

		// 路由组里面的路由组
		v12 := v1.Group("/v12")
		{
			v12.GET("/getInfo", func(context *gin.Context) {
				context.String(http.StatusOK, context.FullPath())
			})
		}
	}
	v1.GET("/getInfo22", func(context *gin.Context) {
		context.String(http.StatusOK, context.FullPath())
	})

	v2 := router.Group("/v2")
	{

		v2.GET("/getInfo222", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})

		v2.GET("/getInfo333", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
	}

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run("localhost:8090")

}
