// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"web_06_input/controllers"
	"web_06_input/models"

	beego "github.com/beego/beego/v2/server/web"
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
		beego.NSNamespace("/user",
			// http://localhost:8080/v1/user/insertUser
			beego.NSPost("insertUser", userController.InsertUser),
			// http://localhost:8080/v1/user/deleteUser/123
			beego.NSDelete("deleteUser/:userId", userController.DeleteUser),
			// http://localhost:8080/v1/user/updateUser/123
			beego.NSPut("updateUser/:userId", userController.UpdateUser),
			// http://localhost:8080/v1/user/getAll
			beego.NSGet("getAll", userController.GetAll),
			// http://localhost:8080/v1/user/getList?xxx=x
			beego.NSGet("getList", userController.GetList),
			// http://localhost:8080/v1/user/getListPage?xxx=x
			beego.NSGet("getListPage", userController.GetListPage),
			// http://localhost:8080/v1/user/getById/123
			beego.NSGet("getById/:userId", userController.GetById),
			// http://localhost:8080/v1/user/upload1/123
			beego.NSPost("upload1/:userId", userController.Upload1),
			// http://localhost:8080/v1/user/upload1/123
			beego.NSPost("upload2/:userId", userController.Upload2),
			// http://localhost:8080/v1/user/downloadFile
			beego.NSGet("downloadFile", userController.DownloadFile),
		),
	)
	beego.AddNamespace(ns)
}
