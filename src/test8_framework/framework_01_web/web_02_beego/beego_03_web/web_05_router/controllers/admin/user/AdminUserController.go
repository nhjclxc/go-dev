package user

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// Operations about goods
type AdminUserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Success 200 {object} models.User
// @router /v1/api/goods/getNewstGoodsList [get]
func (u *AdminUserController) InsertUser(ctx *beecontext.Context) {

	fmt.Println("InsertUser")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "InsertUser",
	})

}

func (u *AdminUserController) UpdateUser(ctx *beecontext.Context) {

	fmt.Println("UpdateUser")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "UpdateUser",
	})

}
