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

	// 创建一个路由组
	v1 := router.Group("/v1")
	{
		v1.GET("/getInfo", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
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
