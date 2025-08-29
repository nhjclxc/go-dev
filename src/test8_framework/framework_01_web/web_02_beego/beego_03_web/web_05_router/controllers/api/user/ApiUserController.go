package user

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// Operations about Users
type ApiUserController struct {
	beego.Controller
}

/*
// @Title Get
// @Description get anonymous_user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
*/
func (u *ApiUserController) GetMyInfo(ctx *beecontext.Context) {
	//uid := u.GetString(":uid")
	//fmt.Println("uid = ", uid)

	uid := ctx.Input.Param(":uid")
	fmt.Println("路径参数 uid:", uid)

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "路径参数 uid:" + uid,
	})
}
