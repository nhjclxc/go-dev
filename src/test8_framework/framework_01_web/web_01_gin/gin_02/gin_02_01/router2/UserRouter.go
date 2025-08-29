package router2

import (
	"github.com/gin-gonic/gin"

	"gin_02_01/controller/admin"
)

func UserRoutesInit(router *gin.Engine) {
	userRouter := router.Group("/anonymous_user")
	{
		userController := admin.UserController{}
		// 路由绑定 handler 处理方法
		userRouter.GET("/getById", userController.GetUserById)
		userRouter.GET("/getPageList", userController.GetUserPageList)
	}
}
