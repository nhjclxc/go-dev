package router

import (
	"github.com/gin-gonic/gin"

	"gin_02_03/controller"
)

func UserRoutesInit(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userController := controller.UserController{}
		// 路由绑定 handler 处理方法
		userRouter.GET("/getById", userController.GetById)
		userRouter.GET("/getCaptcha", userController.GetCaptcha)
		userRouter.GET("/validateCaptcha", userController.ValidateCaptcha)
	}
}
