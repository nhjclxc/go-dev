package router

import (
	"github.com/gin-gonic/gin"

	"gorm_01/controller"
)

func GenTableColumn2RoutesInit(router *gin.Engine) {
	genTable2Router := router.Group("/genTableColumn2")
	{
		genTableColumn2Controller := controller.GenTableColumn2Controller{}
		// 路由绑定 handler 处理方法

		// http://localhost:8090/genTableColumn2/getById
		genTable2Router.GET("/getById", genTableColumn2Controller.GetById)
	}
}
