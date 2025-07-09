package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初识 gin 框架
func main() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "我是第 %d 个 %s 应用哦 \n", 1, "gin")
	})
	router.GET("/news", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")
		context.String(http.StatusOK, "我是 news 页面，现在请求的文章id = %v, name = %v \n", id, name)
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

	//启动端口监听
	// 默认是：0.0.0.0:8080
	router.Run(":8090")
	//router.Run("localhost:8090")
}
