package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"web_08_painc/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /getById/:userId [get]
func (this *UserController) GetById() {

	userId := this.GetString(":userId")

	fmt.Println("GetById.userId = ", userId)

	// 测试全局异常处理是否失效，beego.BConfig.RecoverFunc
	if userId == strconv.Itoa(111) {
		zero := 0
		fmt.Println(111 / zero)
	}


	if userId != "" {
		user, err := models.GetUser(userId)
		if err != nil {
			this.Data["json"] = err.Error()
			// 错误处理？？？

			// 我们在做 Web 开发的时候，经常需要页面跳转和错误处理，Beego 这方面也进行了考虑，通过 Redirect 方法来进行跳转：
			//this.Redirect("https://www.baidu.com/", 302)

			fmt.Println("终止请求前 111")
			// 终止请求
			this.Abort("401")

			// 这样 this.Abort("401") 之后的代码不会再执行，
			fmt.Println("终止请求后 222")

		} else {
			this.Data["json"] = user
		}
	}
	this.ServeJSON()
}
