package main

import (
	"github.com/gin-gonic/gin"
	"gorm_01/middleware"
	myRouter "gorm_01/router"
)

// GORM 增删改查 的使用
func main() {

	// 创建路由器
	router := gin.Default()

	// 注册全局中间件
	router.Use(middleware.GlobalPainc)

	// 注册路由
	myRouter.GenTable2RoutesInit(router)
	myRouter.GenTableColumn2RoutesInit(router)

	// 连表查询
	// 一、一对一：在子表中写：GenTable2 GenTable2 `gorm:"foreignKey:TableId;references:TableId"`
	// 二、一对多：在主表中写：GenTableColumn2 []GenTableColumn2 `gorm:"foreignKey:TableId"`
	// 三、多对多：

	// 启动
	router.Run(":8090")

}
