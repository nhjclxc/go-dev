package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutesInit(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "admin.user")
		})
		adminRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "admin.news")
		})
	}
}
