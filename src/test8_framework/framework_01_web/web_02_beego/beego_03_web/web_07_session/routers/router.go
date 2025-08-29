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
	"web_07_session/controllers"
	"web_07_session/models"
)

func init() {

	userController := &controllers.UserController{}
	userController.UserMap = make(map[string]*models.User, 8)
	userController.UserMap["123"] = &models.User{
		UserId:   "123",
		Username: "root",
		Password: "root@123",
	}

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/anonymous_user",
			// http://localhost:8080/v1/user/getById/123
			beego.NSGet("getById/:userId", userController.GetById),
		),
	)

	beego.AddNamespace(ns)
}
