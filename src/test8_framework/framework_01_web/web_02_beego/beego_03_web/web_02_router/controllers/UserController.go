package controllers

import (
	"encoding/json"
	"fmt"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"time"
	"web_02_router/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller

	uuid string
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// beecontext "github.com/beego/beego/v2/server/web/context"

// 要使用 注册函数式风格路由注册 ,则必须在方法的参数里面加入 ctx *beecontext.Context, 且不能由其他的参数
// 同时不能定义返回值
func (u *UserController) TestLogout(ctx *beecontext.Context) {

	fmt.Println("TestLogout")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "TestLogout success",
	})

	u.uuid = time.Now().String()

	fmt.Println("u.uuid = ", u.uuid)
}
func (u *UserController) TestLogout222(ctx *beecontext.Context) {

	fmt.Println("TestLogout222")

	err := ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "TestLogout222 success",
	})
	if err != nil {
		return
	}

	// 验证在 router.go 里面,注册 TestLogout 和 TestLogout222 的时候是不是使用了同一个实例
	fmt.Println("u.uuid = ", u.uuid)
}
func (u *UserController) TestLogout333(ctx *beecontext.Context) {

	fmt.Println("TestLogout333")

	err := ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "TestLogout333 success",
	})
	if err != nil {
		return
	}

	// 验证在 router.go 里面,注册 TestLogout 和 TestLogout222 的时候是不是使用了同一个实例
	fmt.Println("u.uuid = ", u.uuid)
}

func (u *UserController) GetUserById(ctx *beecontext.Context) {
	fmt.Println("GetUserById")

	userId := ctx.Input.Param(":userId")
	fmt.Println("路径参数 userId:", userId)

	err := ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "路径参数 userId:" + userId,
	})
	if err != nil {
		return
	}
}
