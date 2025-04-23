package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleRoutesInit(router *gin.Engine) {
	roleRouter := router.Group("/role")
	{
		roleRouter.GET("/getById", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
		roleRouter.GET("/getPageList", func(context *gin.Context) {
			context.String(http.StatusOK, context.FullPath())
		})
	}
}
