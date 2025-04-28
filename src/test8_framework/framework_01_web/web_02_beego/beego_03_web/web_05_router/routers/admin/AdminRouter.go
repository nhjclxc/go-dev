// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package admin

import (
	beego "github.com/beego/beego/v2/server/web"
	apiGoods "web_05_router/routers/admin/goods"
	apiUser "web_05_router/routers/admin/user"
)

var (
	AadminRouterLinkNamespace beego.LinkNamespace
)

func init() {

	AadminRouterLinkNamespace = beego.NSNamespace("/admin",
		apiUser.AdminUserRouterLinkNamespace,
		apiGoods.AdminGoodsRouterLinkNamespace,
	)
}
