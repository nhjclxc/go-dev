package controllers

import (
	"encoding/json"
	"fmt"
	"time"
	"web_01_router/models"

	beego "github.com/beego/beego/v2/server/web"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

// Operations about Users
type UserController struct {
	beego.Controller
	uuid time.Time
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

func (u *UserController) DoLogout() {
	u.Data["json"] = "MyDoLogout success"
	u.ServeJSON()
}

// 如果使用函数注册的话，这个方法就不是由beego注入的了，因此，这个方法实际是静态方法，没有绑定实例
func LogoutHandler(ctx *beecontext.Context) {

	fmt.Println("LogoutHandler success")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "LogoutHandler success",
	})
}
func (u *UserController) LogoutHandler222(ctx *beecontext.Context) {

	fmt.Println("LogoutHandler222 success")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "LogoutHandler222 success",
	})

	fmt.Println(u)
	fmt.Println(u.Data)
	fmt.Println(u.ViewPath)

	// u 对象不能使用
	//u.Data["json"] = "MyDoLogout success"
	//u.ServeJSON()
}

func (u *UserController) HelloWorld() {
	u.Ctx.WriteString("Hello, world")
}

// 不建议使用没有指针接收器的控制器,应当写为(u *UserController)
// 但是这样beego也能执行
func (u UserController) HelloWorldNonePoint() {
	u.Ctx.WriteString("Hello, world HelloWorldNonePoint")
}

// 只会在调用 UserController 的接口之前被执行,其他的控制器不受这个影响
func (u *UserController) Prepare() {
	fmt.Println("u.uuid = ", u.uuid)
	fmt.Println("u.ViewPath = ", u.ViewPath)

	// 1. 拿请求头
	token := u.Ctx.Input.Header("Authorization")
	cookie := u.Ctx.Input.Header("cookie")
	fmt.Println("Authorization Token:", token)
	fmt.Println("cookie cookie:", cookie)

	// 2. 拿Query参数（URL ?xxx=xxx 这种）
	username := u.GetString("username")
	fmt.Println("Query参数 username:", username)

	// 3. 拿Path参数（比如 /user/123，这种需要路由里定义了:param）
	id := u.Ctx.Input.Param(":id")
	fmt.Println("路径参数 id:", id)

	// 4. 拿POST表单参数
	password := u.GetString("password")
	fmt.Println("表单参数 password:", password)

	// 5. 拿整个请求Body（比如JSON提交）
	bodyBytes := u.Ctx.Input.RequestBody
	fmt.Println("原始Body内容:", string(bodyBytes))

	// 6. 拿客户端IP
	clientIP := u.Ctx.Input.IP()
	fmt.Println("客户端IP:", clientIP)

	fmt.Println("UserController Prepare... 每次请求都会执行")
}
