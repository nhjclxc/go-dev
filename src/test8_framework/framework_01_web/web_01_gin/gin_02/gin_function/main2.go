package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 初识 gin 框架
func main() {

	router := gin.Default()
	router.GET("/2", func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		f := c.Query("f")

		flag := true
		for f == "f" && flag {
			select {
			case <-c.Done():
				fmt.Println("请求被关闭")
				flag = false
				break
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("处理中...")
			}
		}
		fmt.Println("请求结束")

	})

	router.GET("/", func(c *gin.Context) {
		fmt.Println("client IP:", c.ClientIP())
		f := c.Query("f")

		for f == "f" {
			select {
			case <-c.Done():
				fmt.Println("请求被关闭") // 只有真正感知到 TCP 关闭才会打印
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("处理中22...")
			}
		}

		fmt.Println("请求结束")
		c.String(200, "done")
	})
	// http://127.0.0.1:8090
	router.Run(":5001")
}
