package router2

import (
	"github.com/gin-gonic/gin"

	"gin_02_01/controller/admin"
)

// 初始化 News 路由
func NewsRoutesInit(router *gin.Engine) {
	newsRouter := router.Group("/news")
	{
		newsController := admin.NewsController{}
		// 绑定 news 路由的 句柄处理方法
		newsRouter.GET("/index", newsController.Index)
		newsRouter.GET("/getById", newsController.GetNewsById)
	}
}
