// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"web_03_router/controllers"

	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

func init() {

	// [web模块-路由-Web 命名空间](https://beegodoc.com/zh/developing/web/router/namespace.html)

	// namespace，也叫做命名空间，是 Beego 提供的一种逻辑上的组织 API 的手段。

	userController := controllers.UserController{}

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSCtrlGet("/home", (*controllers.UserController).Logout),
		beego.NSRouter("/user", &controllers.UserController{}),

		// 最建议的方法,函数式路由注册方法
		beego.NSGet("/health", userController.Health),
	)

	beego.AddNamespace(ns)

	//
	//

	// namespace 的条件执行
	// Beego 的namespace提供了一种条件判断机制。只有在符合条件的情况下，注册在该namespace下的路由才会被执行。本质上，这只是一个filter的应用。

	ns2 := beego.NewNamespace("/v2",
		// 例如，我们希望用户的请求的头部里面一定要带上一个x-trace-id才会被后续的请求处理：
		beego.NSCond(func(b *beecontext.Context) bool {
			return b.Request.Header["x-trace-id"][0] != ""
		}),

		// 最建议的方法,函数式路由注册方法
		beego.NSGet("/healthV2", userController.HealthV2),
	)
	// 一般来说，我们现在也不推荐使用这个特性，因为它的功能和filter存在重合，
	//我们建议大家如果有需要，应该考虑自己正常实现一个filter，代码可理解性会更高。
	//该特性会考虑在未来的版本用一个filter来替代，而后移除该方法。

	beego.AddNamespace(ns2)

	//
	//
	//

	// Namespace Filter
	// namespace同样支持filter。该filter只会作用于这个namespace之下注册的路由，而对别的路由没有影响。
	// 我们有两种方式添加Filter，
	//	第一种方式: 一个是在NewNamespace中，调用web.NSBefore或者web.NSAfter，
	ns3 := beego.NewNamespace("/v3",
		beego.NSBefore(func(ctx *beecontext.Context) {
			fmt.Println("beego.NSBefore: ", ctx.Request.URL)
		}),
		// 最建议的方法,函数式路由注册方法
		beego.NSGet("/healthV3", userController.HealthV3),
		// 只要中间异常返回、直接写了响应、或者路由不通，NSAfter就不会触发
		// 如果中途直接写了响应（比如ctx.ResponseWriter.Write())，或者panic了，或者路由不匹配，NSAfter是不会被正常调用的。
		// 即要想 NSAfter 被执行,则在 HealthV3 中不能使用 ctx.JSONResp 写入响应
		beego.NSAfter(func(ctx *beecontext.Context) {
			fmt.Println("beego.NSAfter: ", ctx.Request.URL)
		}),
	)

	beego.AddNamespace(ns3)

	//

	//	第二种方式: 也可以调用ns.Filter()
	ns5 := beego.NewNamespace("/v5",
		// 最建议的方法,函数式路由注册方法
		beego.NSGet("/healthV5", userController.HealthV5),
	)
	ns5.Filter("before", func(ctx *beecontext.Context) {
		fmt.Println("this is filter for health before ", ctx.Request.URL)
	})
	ns5.Filter("after", func(ctx *beecontext.Context) {
		fmt.Println("this is filter for health after ", ctx.Request.URL)
	})

	beego.AddNamespace(ns5)

}
