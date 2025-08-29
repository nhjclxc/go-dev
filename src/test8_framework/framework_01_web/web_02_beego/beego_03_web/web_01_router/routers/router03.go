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

	usercontroller := controllers.UserController{}
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/anonymous_user",
			beego.NSInclude(
				&usercontroller,
			),
		),
		// http://localhost:8090/api/v1/logoutHandler222
		beego.NSGet("/logoutHandler222", usercontroller.LogoutHandler222),
	)

	beego.AddNamespace(ns)
}
