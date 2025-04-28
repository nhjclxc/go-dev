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

// 使用 AutoPrefix 来注册路由
// https://beegodoc.com/zh/developing/web/router/ctrl_style/#autoprefix
func init() {

	// 注意：通过 AutoPrefix 方式注册的路由对应的句柄方法不能由参数，如	Logout，DoLogout，HelloWorld等

	// get http://localhost:8080/api/user/logout
	// get http://localhost:8080/api/user/dologout
	// get http://localhost:8080/api/user/helloworld
	// you will see return "Hello, world"
	ctrl := &controllers.UserController{}
	beego.AutoPrefix("api", ctrl)
	//beego.Run()
}
