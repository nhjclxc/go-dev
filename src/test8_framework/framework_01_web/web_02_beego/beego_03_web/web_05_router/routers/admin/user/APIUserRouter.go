// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	controllers "web_05_router/controllers/admin/user"
)

var (
	AdminUserRouterLinkNamespace beego.LinkNamespace
)

func init() {

	adminUserController := controllers.AdminUserController{}
	AdminUserRouterLinkNamespace = beego.NSNamespace("/user",
		// /v1/admin/user/insertUser
		beego.NSPost("/insertUser", adminUserController.InsertUser),
		beego.NSPut("/insertUser", adminUserController.UpdateUser),
		// 后续还有更多的 API 要被注册 ...
	)

}
