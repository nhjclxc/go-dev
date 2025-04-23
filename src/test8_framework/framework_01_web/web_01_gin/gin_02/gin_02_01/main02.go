package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin 框架各种数据类型的响应，string、json、jsonp、html、xml...
func main02() {

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "首页")
	})
	router.GET("/json", func(context *gin.Context) {
		// func (c *Context) JSON(code int, obj any)
		// type H map[string]any

		//context.JSON(http.StatusOK, map[string]interface{}{
		context.JSON(http.StatusOK, map[string]any{
			"code":      200,
			"timestamp": time.Now(),
			"success":   true,
			"msg":       "操作成功",
			"data":      []int{1, 2, 3, 4, 5, 6},
		})
	})
	router.GET("/json2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":      200,
			"timestamp": time.Now(),
			"success":   true,
			"msg":       "操作成功",
			"data":      []int{1, 2, 3, 4, 5, 6},
		})
	})
	router.GET("/json3", func(context *gin.Context) {
		context.JSON(http.StatusOK, Student{
			Id:   666,
			Name: "zhangsan",
			Age:  18,
			addr: "中国-北京",
		})
	})
	router.GET("/json4", func(context *gin.Context) {
		context.JSON(http.StatusOK, &Student{
			Id:   666,
			Name: "zhangsan",
			Age:  18,
			addr: "中国-北京",
		})
	})

	// 响应 jsonp 数据
	router.GET("/jsonp", func(context *gin.Context) {
		context.JSONP(http.StatusOK, &Student{
			Id:   666,
			Name: "zhangsan",
			Age:  18,
			addr: "中国-北京",
		})
	})

	//启动端口监听
	// 默认是：0.0.0.0:8080
	//router.Run(":8090")
	router.Run("localhost:8090")
}

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"my_name"`
	Age  int
	// 要注意可见性，addr前端不可见
	addr string `json:"addr"`
}
