package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	e := gin.Default()
	//api := e.Group("/api/v1")
	api := e.Group("/api/v1", OpenApiAuthMiddleware())
	api.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
			"code":  200,
		})
	})
	e.Run(":8080")
}
