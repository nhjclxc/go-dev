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
	"web_02_router/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//注册函数式风格路由注册

	userController := controllers.UserController{}
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/anonymous_user",
			beego.NSInclude(
				&userController,
			),

			// 在某个命名空间里面注册路由写 NSxxx ,其中 xxx就是对应的请求方法,(NS表示 Namespace)
			// http://localhost:8080/v1/user/testLogout
			beego.NSGet("/testLogout", userController.TestLogout),
			beego.NSGet("/testLogout222", userController.TestLogout222),
			// http://localhost:8080/v1/user/getUserById/666
			beego.NSGet("/getUserById/:userId", userController.GetUserById),
		),
	)
	// 不在某个命名空间内,则直接使用 xxx 注册路由
	beego.Get("/v1/anonymous_user/testLogout333", userController.TestLogout333)

	//Get(rootpath string, f HandleFunc)
	//Post(rootpath string, f HandleFunc)
	//Delete(rootpath string, f HandleFunc)
	//Put(rootpath string, f HandleFunc)
	//Head(rootpath string, f HandleFunc)
	//Options(rootpath string, f HandleFunc)
	//Patch(rootpath string, f HandleFunc)
	//Any(rootpath string, f HandleFunc)

	// 过多与路由规则相关的知识,请看:https://beegodoc.com/zh/developing/web/router/router_rule.html
	// * 匹配 , 如: /api/anonymous_user/name/*
	// /api/anonymous_user/name
	// /api/anonymous_user/name/tom
	// /api/anonymous_user/name/jerry/home
	// 即，只要前缀符合/api/anonymous_user/name，那么就会命中。

	// /api/*/name , 则匹配：/api/tom/name 、 /api/tom/jerry/name

	// * 匹配可以类似于 Java 中的切面编程（AOP）

	// 在 Beego 里，有一个叫 过滤器（Filter） 的东西，
	//它和 Java AOP 的 "切面拦截器" 十分相似，本质就是：
	//在请求处理前/后
	//在某些URL匹配规则下
	//执行一段统一的逻辑

	/*
			在Beego里实现类似切面的拦截, 推荐使用 beego.InsertFilter

			// 在所有 /api/ 开头的请求前执行
			beego.InsertFilter("/api/*", beego.BeforeRouter, func(ctx *context.Context) {
			    // 这里可以做鉴权、日志等切面逻辑
			    fmt.Println("拦截器：请求路径是", ctx.Input.URL())
			})

		解释一下：

		/api/*：就是你说的路由通配符（星号 * 匹配所有子路径）

		beego.BeforeRouter：表示在路由匹配之前就执行

		func(ctx *context.Context)：拦截处理函数

		是不是跟 Java 的 @Around / @Before advice 很像？ 几乎就是 Go版 AOP了。

	*/

	beego.AddNamespace(ns)

	// 查看所有已注册的路由

	beego.BConfig.RouterCaseSensitive = false
	beego.AutoRouter(&userController)
	tree := beego.PrintTree()
	methods := tree["Data"].(beego.M)
	fmt.Println("num = ", len(methods))
	for k, v := range methods {
		fmt.Printf("%s => %v\n", k, v)
	}

	// GET => [/anonymous_user/getuserbyid/* map[*:GetUserById] controllers.UserController]
	// GET 方法访问接口
	// 接口路径是:/anonymous_user/getuserbyid/*, * 表示任意的参数
	// 执行 controllers.UserController 里面的 GetUserById 方法
}
