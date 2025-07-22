package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strconv"
)

// gin实现全局panic异常处理



func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {


		defer func() {
			if err := recover(); err != nil {
				// 打印错误和堆栈信息
				fmt.Printf("panic recovered: %v\n", err)
				fmt.Printf("stack trace:\n%s\n", string(debug.Stack()))

				// 返回统一格式的 JSON 响应
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":  500,
					"msg":   "服务器内部错误：" + fmt.Sprintf("%v", err),
					"error": fmt.Sprintf("%v", err),
					"data": nil,
				})
			}
		}()

		c.Next()

	}
}

func main() {

	r := gin.Default()

	// 启用跨域支持
	r.Use(cors.Default())
	r.Use(RecoveryMiddleware())

	//// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	//r.Use(gin.Recovery())

	r.GET("/cul", cul)

	r.Run(":8080")
}

func cul(c *gin.Context) {
	num1Str := c.Query("num1")
	num2Str := c.Query("num2")

	num1, _ := strconv.ParseInt(num1Str, 10, 64)
	num2, _ := strconv.ParseInt(num2Str, 10, 64)
	res := strconv.FormatInt(num1/num2, 10)
	fmt.Printf("cul = " + res)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": nil,
			"msg":   "操作成功",
			"data":  "res = " + res,
		},
	)
}
