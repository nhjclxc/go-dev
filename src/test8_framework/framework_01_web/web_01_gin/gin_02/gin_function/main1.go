package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 初识 gin 框架
func main1() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		fmt.Println(c.ClientIP())

		ctx := c.Copy()
		ctx.Set("key1", "val123")

		go func() {

			value, exists := ctx.Get("key1")
			if !exists {
				fmt.Println("ctx not exists key1")
			}
			fmt.Println("goroutine exists key1 = ", value)
		}()
	})
	// http://127.0.0.1:8090
	router.Run(":8090")
}
