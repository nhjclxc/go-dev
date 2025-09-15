package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reflect"
)

// gin实现多次bind读取数据

func BindAndValidateMiddleware(objType reflect.Type) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "读取请求体失败"})
			c.Abort()
			return
		}

		// 缓存 body，供 Handler 或其他中间件使用
		c.Set("rawBody", bodyBytes)

		// 重新设置 c.Request.Body，让 Handler 可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 创建结构体实例
		obj := reflect.New(objType).Interface()

		// 绑定 JSON
		if err := c.ShouldBindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

type UserRequest struct {
	Username string `json:"username" from:"username" `
	Password string `json:"password" from:"password" `
	Phone    string `json:"Phone" from:"Phone" `
	Email    string `json:"email" from:"email" `
}

func main() {

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var req UserRequest
		//err := c.ShouldBindJSON(&req)
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("val %#v", req)

		c.JSON(200, gin.H{"data": req})
	})

	r.Run(":8080")

}
