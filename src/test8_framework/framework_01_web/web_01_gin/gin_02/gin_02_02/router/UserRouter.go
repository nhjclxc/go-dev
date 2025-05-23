package router

import (
	middleware2 "gin_02_02/middleware"
	"github.com/gin-gonic/gin"

	"gin_02_02/controller"
)

func UserRouterInit(router *gin.Engine) {

	// 注册路由分组中间件，方式一
	//userGroup := router.Group("/user", middleware.UserMiddleware)

	// 注册路由分组中间件，方式二
	userGroup := router.Group("/user")
	userGroup.Use(middleware2.UserMiddleware)
	{
		userController := controller.UserController{}

		// Gin 中的中间件必须是一个gin.HandlerFunc类型，配置路由的时候可以传递多个func回调函数，最后一个func回调函数前面触发的方法都可以称为中间件。

		// func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
		// func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
		// 根据以上POST、GET等方法的定义，我们可以知道
		// 每一个接口都可以有多个 句柄handler 来处理每隔请求，
		// 那么，我们就可以通过这个特性来实现接口鉴权、接口参数日志，以及返回值的记录等等操作，
		// 其中，除了业务接口以外的 句柄handler 都被称之为接口中间件 Middleware。
		// 如果一个接口对应多个 句柄handler ，那么 gin 将安装 句柄handler顺序依次执行，
		// 此外，如果某个 句柄handler 里面执行了 context.Abort() 方法，那么剩下的 句柄handler 将不会被执行
		// 会在执行 context.Abort() 方法的位置终止该请求
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/getById",
			middleware2.Authentication,  // 先鉴权
			middleware2.RequestParamLog, // 接着打印请求日志
			userController.GetById,      // 执行接口业务
			middleware2.ResponseDataLog, // 响应数据
		)
		userGroup.GET("/pageList", userController.PageList)
	}

}
