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

	// router01 的方式写路由，造成了过多的文本硬编码

	usercontroller := controllers.UserController{}

	// beego/v2 支持更少的硬编码注册路由
	// http://localhost:8090/doLogout
	beego.Router("/doLogout", &usercontroller, "get:DoLogout")


	// 以下两个方式注意：：：LogoutHandler是一个静态方法了，不是UserController的实例，不能再用实例去注册了

	// 再现代一点：直接函数指针绑定（beego/v2 支持），直接使用 句柄Handler 引用来注册接口
	beego.Get("/logoutHandler", controllers.LogoutHandler)

	// 支持命名空间的绑定
	ns := beego.NewNamespace("/api/v1",
		beego.NSGet("/logoutHandler", controllers.LogoutHandler),
		// web.NSGet(path, handlerFunc)
		// web.NSPost(path, handlerFunc)
		// web.NSPut(path, handlerFunc)
		// web.NSDelete(path, handlerFunc)
		// web.NSPatch(path, handlerFunc)
		// web.NSOptions(path, handlerFunc)
	)
	beego.AddNamespace(ns)
}
