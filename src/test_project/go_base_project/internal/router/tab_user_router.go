package router

import (
	"github.com/gin-gonic/gin"
	"go_base_project/internal/controller"
)

// TabUserRouter 用户  路由Router层
type TabUserRouter struct {
	tabUserController *controller.TabUserController
}

// NewTabUserRouter 创建 TabUser 用户  路由Router层
func NewTabUserRouter(tabUserController *controller.TabUserController) *TabUserRouter {
	return &TabUserRouter{
		tabUserController: tabUserController,
	}
}

// InitTabUserRouter 初始化 TabUserRouter 路由
func (tur *TabUserRouter) InitTabUserRouter(privateRouterOrigin *gin.RouterGroup, publicRouterOrigin *gin.RouterGroup) {
	privateRouter := privateRouterOrigin.Group("/tab/user")
	{
		// PrivateRouter 下是一些必须进行登录的接口
		// http://localhost:8080/private

		privateRouter.POST("", tur.tabUserController.InsertTabUser)              // 新增用户
		privateRouter.PUT("", tur.tabUserController.UpdateTabUser)               // 修改用户
		privateRouter.DELETE("/:idList", tur.tabUserController.DeleteTabUser)    // 删除用户
		privateRouter.GET("/:id", tur.tabUserController.GetTabUserById)          // 获取用户 详细信息
		privateRouter.GET("/list", tur.tabUserController.GetTabUserList)         // 查询用户 列表
		privateRouter.GET("/pageList", tur.tabUserController.GetTabUserPageList) // 分页查询用户 列表
		privateRouter.GET("/export", tur.tabUserController.ExportTabUser)        // 导出用户 列表
	}

	publicRouter := publicRouterOrigin.Group("/tab/user")
	{
		// PublicRouter 下是一些无需登录的接口，可以直接访问，无须经过授权操作
		// http://localhost:8080/public

		publicRouter.POST("", tur.tabUserController.InsertTabUser)              // 新增用户
		publicRouter.PUT("", tur.tabUserController.UpdateTabUser)               // 修改用户
		publicRouter.DELETE("/:id", tur.tabUserController.DeleteTabUser)        // 删除用户
		publicRouter.GET("/:id", tur.tabUserController.GetTabUserById)          // 获取用户 详细信息
		publicRouter.GET("/list", tur.tabUserController.GetTabUserList)         // 查询用户 列表
		publicRouter.GET("/pageList", tur.tabUserController.GetTabUserPageList) // 分页查询用户 列表
		publicRouter.GET("/export", tur.tabUserController.ExportTabUser)        // 导出用户 列表
	}
}
