package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutesInit(router *gin.Engine) {
	userRouter := router.Group("/anonymous_user")
	{
		userRouter.GET("/getById", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
		userRouter.GET("/getPageList", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
	}
}
