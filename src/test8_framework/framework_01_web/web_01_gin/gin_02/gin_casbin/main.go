package main

import (
	"fmt"
	"gin_casbin/config"
	"gin_casbin/handler"
	"gin_casbin/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 DB
	config.InitDB()

	// GORM Adapter for MySQL
	a, err := gormadapter.NewAdapterByDB(config.DB)
	if err != nil {
		fmt.Println("NewAdapterByDB", err)
		return
	}

	e, err := casbin.NewEnforcer("model.conf", a)
	if err != nil {
		fmt.Println("NewEnforcer", err)
		return
	}
	if err := e.LoadPolicy(); err != nil {
		fmt.Println("LoadPolicy failed:", err)
		return
	}
	fmt.Println("LoadPolicy success")

	r := gin.Default()

	// 登录接口
	r.POST("/login", handler.Login)

	// 受保护的接口
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware(), middleware.CasbinMiddleware(e))
	{
		// common
		auth.GET("/home", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /home \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})
		auth.GET("/home2", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /home2 \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})

		// admin
		auth.GET("/user", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /user \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})

		// superadmin
		auth.GET("/role", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /role \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})
	}

	open := r.Group("/openapi")
	{
		open.GET("/hi", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			str := fmt.Sprintf("%s 访问了 /hi \n", username)
			fmt.Printf(str)
			ctx.JSON(200, str)
		})
	}

	r.Run(":8080")
}
