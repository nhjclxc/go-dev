package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 使用中间件实现简单限流（基于 token bucket / leaky bucket）

// go get golang.org/x/time/rate

// 实现limit.go漏桶算法

// 注册限流中间件

// 接口测试


func main() {

	router := gin.Default()

	// 启用跨域支持
	router.Use(cors.Default())

	// 注册限流中间件
	router.Use(RateLimitMiddleware())

	// http://127.0.0.1:8080/ping
	// for /L %i in (1,1,20) do curl -s -w "%%{http_code}\n" -o NUL http://localhost:8080/ping
	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	log.Println("服务启动于 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
