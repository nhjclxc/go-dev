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

// 手动注册路由注册路由
// https://beegodoc.com/zh/developing/web/router/ctrl_style/#%E6%89%8B%E5%8A%A8%E8%B7%AF%E7%94%B1
func init() {

	// 如果我们并不想利用AutoRoute或者AutoPrefix来注册路由，因为这两个都依赖于Controller的名字，也依赖于方法的名字。某些时候我们可能期望在路由上，有更强的灵活性。

	// 在 v2.0.2 我们引入了全新的注册方式。下面我们来看一个完整的例子
	// 需要注意的是，我们新的注册方法，要求我们传入方法的时候，传入的是(*YourController).MethodName。
	//这是因为 Go 语言特性，要求在接收器是指针的时候，如果你希望拿到这个方法，那么应该用(*YourController)的形式。

	// get http://localhost:8080/api/user/helloworld
	beego.CtrlGet("/api/user/helloworld22", (*controllers.UserController).HelloWorld)
	beego.CtrlGet("/api/user/helloworld33", (*controllers.UserController).HelloWorld)
	beego.CtrlGet("/api/user/helloworldNonePoint", (*controllers.UserController).HelloWorldNonePoint)

	/*
		不是同一个实例！
		每次请求都会新建一个新的 UserController 实例，分别处理各自的请求。
		所以它们是不同的对象，内存地址也不同。

			beego.CtrlGet 注册的是：
			Controller 的方法指针（*UserController).HelloWorld），不是Controller实例本身。
			Beego在收到HTTP请求时，每次都会创建一个新的Controller实例。
				这样可以保证每个请求是线程安全的。
				不然如果所有请求都用同一个对象，多个请求一来，数据就混乱了（并发问题）。

			大概内部逻辑是这样的：
				每次收到请求时，Beego大概做了这些事：
				反射出 Controller 类型（UserController）
				new(UserController) 创建一个新的对象
				初始化一些字段（比如 Ctx、Data、Input、Output）
				调用你指定的方法（比如 HelloWorld）
			所以 helloworld22 和 helloworld33 注册的是同一个方法引用，
			但是处理请求时，拿到的是不同的实例！所以是不同实例，互不干扰！

	*/

	// get http://localhost:8080/api/user/123
	//beego.CtrlGet("api/user/:id", usercontroller.HelloWorld)
	//
	//// post http://localhost:8080/api/user/update
	//beego.CtrlPost("api/user/update", (*UserController).UpdateUser)
	//
	//// http://localhost:8080/api/user/home
	//beego.CtrlAny("api/user/home", (*UserController).UserHome)
	//
	//// delete http://localhost:8080/api/user/delete
	//beego.CtrlDelete("api/user/delete", (*UserController).DeleteUser)
	//
	//// head http://localhost:8080/api/user/head
	//beego.CtrlHead("api/user/head", (*UserController).HeadUser)
	//
	//// patch http://localhost:8080/api/user/options
	//beego.CtrlOptions("api/user/options", (*UserController).OptionUsers)
	//
	//// patch http://localhost:8080/api/user/patch
	//beego.CtrlPatch("api/user/patch", (*UserController).PatchUsers)
	//
	//// put http://localhost:8080/api/user/put
	//beego.CtrlPut("api/user/put", (*UserController).PutUsers)

}
