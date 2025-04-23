package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin 渲染网页模板
func main03() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 注意：要想使用网页模板，必须在创建完 路由器 之后加载模板页面数据
	router.LoadHTMLGlob("templates/*")

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "首页")
	})

	// 响应 html 模板数据
	router.GET("/html", func(context *gin.Context) {

		// func (c *Context) HTML(code int, name string, obj any)
		// 第一个参数：响应状态码
		// 第二个参数：模板页面名称
		// 第三个参数：模板页面的数据

		context.HTML(http.StatusOK, "index.html", gin.H{
			"code":      200,
			"timestamp": time.Now(),
			"success":   true,
			"msg":       "操作成功",
			"data":      []int{1, 2, 3, 4, 5, 6},
		})
	})

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run("localhost:8090")
}
