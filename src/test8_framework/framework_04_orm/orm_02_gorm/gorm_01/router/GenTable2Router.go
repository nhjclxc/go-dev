package router

import (
	"github.com/gin-gonic/gin"

	"gorm_01/controller"
)

func GenTable2RoutesInit(router *gin.Engine) {
	genTable2Router := router.Group("/genTable2")
	{
		genTable2Controller := controller.GenTable2Controller{}
		// 路由绑定 handler 处理方法

		// http://localhost:8090/genTable2/insert
		genTable2Router.POST("/insert", genTable2Controller.Insert)
		// http://localhost:8090/genTable2/delete
		genTable2Router.DELETE("/delete", genTable2Controller.Delete)
		// http://localhost:8090/genTable2/update
		genTable2Router.PUT("/update", genTable2Controller.Update)
		// http://localhost:8090/genTable2/getById
		genTable2Router.GET("/getById", genTable2Controller.GetById)
		// http://localhost:8090/genTable2/getAll
		genTable2Router.GET("/getAll", genTable2Controller.GetAll)
		// http://localhost:8090/genTable2/getList
		genTable2Router.GET("/getList", genTable2Controller.GetList)
		// http://localhost:8090/genTable2/getPageList
		genTable2Router.GET("/getPageList", genTable2Controller.GetPageList)
	}
}
