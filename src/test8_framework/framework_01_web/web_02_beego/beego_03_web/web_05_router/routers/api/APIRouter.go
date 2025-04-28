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
	apiGoods "web_05_router/routers/api/goods"
	apiUser "web_05_router/routers/api/user"
)

var (
	APIRouterLinkNamespace beego.LinkNamespace
)

func init() {

	APIRouterLinkNamespace = beego.NSNamespace("/api",
		apiUser.APIUserRouterLinkNamespace,
		apiGoods.APIGoodsRouterLinkNamespace,
	)
}
