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
	adminRouter "web_05_router/routers/admin"
	apiRouter "web_05_router/routers/api"
)

func init() {
	ns := beego.NewNamespace("/v1",
		apiRouter.APIRouterLinkNamespace,
		adminRouter.AadminRouterLinkNamespace,
	)

	// http://localhost:8080/v1/api/user/getMyInfo/user_11111
	// http://localhost:8080/v1/api/goods/getNewstGoodsList
	// http://localhost:8080/v1/admin/user/insertUser
	// http://localhost:8080/v1/admin/user/updateUser
	// http://localhost:8080/v1/admin/goods/insertGoods

	// v2
	// v3
	// ...

	beego.AddNamespace(ns)
}
