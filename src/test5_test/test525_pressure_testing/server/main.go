package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// 模拟处理延迟
		time.Sleep(50 * time.Millisecond)
		c.String(http.StatusOK, "pong" + time.Now().String())
	})

	r.Run(":8080")
}
