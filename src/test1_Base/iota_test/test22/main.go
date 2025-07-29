package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("hello world")


	r := gin.Default()
	r.GET("hello/gk", helloGK)
	panic(r.Run(":8080"))

}

func helloGK(ctx *gin.Context) {

	// gorm自动创建数据表，db.autoMigrate



	ctx.JSON(200,gin.H{
		"data": "欢迎加入吉快科技！！！",
	})

}
