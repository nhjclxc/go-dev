// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"web_01_router/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// 创建一个新的路由命名空间，前缀是 /v1。
	// 命名空间（Namespace）是 Beego 提供的一种分组机制，方便组织 API 版本化、分类管理路由。 类似于 gin 的routerGroup
	ns := beego.NewNamespace("/v1",
		// 在 /v1 命名空间下，再创建一个子命名空间 /object。
		// 这个子命名空间下的路由路径是 /v1/object/xxx
		beego.NSNamespace("/object",
			beego.NSInclude(
				// NSInclude 动读取 Controller 里用 @router 标注的方法，并批量注册成路由。
				// 将 ObjectController 中通过注释 @router 声明的路由绑定进来。
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

	// 把上面构建的命名空间 ns 注册到 Beego 的全局路由表中。
	beego.AddNamespace(ns)
}
