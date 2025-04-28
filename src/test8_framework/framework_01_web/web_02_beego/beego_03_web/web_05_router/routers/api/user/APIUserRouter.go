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
	controllers "web_05_router/controllers/api/user"
)

var (
	APIUserRouterLinkNamespace beego.LinkNamespace
)

func init() {

	userController := controllers.ApiUserController{}
	APIUserRouterLinkNamespace = beego.NSNamespace("/user",
		// /v1/api/user/getMyInfo
		beego.NSGet("/getMyInfo/:uid", userController.GetMyInfo),
		// 后续还有更多的 API 要被注册 ...
	)

}
