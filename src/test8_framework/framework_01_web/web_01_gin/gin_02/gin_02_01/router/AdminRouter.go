package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutesInit(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/anonymous_user", func(c *gin.Context) {
			c.String(http.StatusOK, "admin.anonymous_user")
		})
		adminRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "admin.news")
		})
	}
}
