// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	":/code/go/go-dev/src/test8_framework/framework_01_web/web_02_beego/beego_01_hello/hello_api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tab_user",
			beego.NSInclude(
				&controllers.TabUserController{},
			),
		),

		beego.NSNamespace("/tab_user_card",
			beego.NSInclude(
				&controllers.TabUserCardController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
