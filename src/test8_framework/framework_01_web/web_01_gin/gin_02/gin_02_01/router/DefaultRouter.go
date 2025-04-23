package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultgoRoutesInit(router *gin.Engine) {
	defaultRouter := router.Group("/")
	{
		defaultRouter.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "项目首页")
		})
	}
}
