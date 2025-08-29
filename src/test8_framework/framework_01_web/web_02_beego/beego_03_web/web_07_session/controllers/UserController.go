package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"web_07_session/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
	UserMap map[string]*models.User
}

// @Title GetById
// @Description get anonymous_user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getById/:uid [get]
func (this *UserController) GetById(ctx *beecontext.Context) {

	// http://localhost:8080/v1/user/getById/123?username=root&password=root@123

	// 获取路径参数
	// 在参数前面加一个冒号 :
	// 注意这个路径参数一定要和 rotuer 里面定义的路径参数一致，否则无法获取
	userId := ctx.Input.Param(":userId")
	user := this.UserMap[userId]

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": user,
	})

	fmt.Println("this.Data = ", this.Data)
	fmt.Println("this.Ctx = ", this.Ctx)

	// session 有几个方便的方法：
	//
	//SetSession(name string, value interface{})
	//GetSession(name string) interface{}
	//DelSession(name string)
	//SessionRegenerateID()
	//DestroySession()

	// 普通 Cookie 处理
	//Beego 通过Context直接封装了对普通 Cookie 的处理方法，可以直接使用：
	//
	//GetCookie(key string)
	//SetCookie(name string, value string, others ...interface{})

	Goland := ctx.GetCookie("Goland-7113b5cb")
	fmt.Println("Goland = ", Goland)

	// put something into cookie,set Expires time
	ctx.SetCookie("cookieName", "web cookie", 10)

}
