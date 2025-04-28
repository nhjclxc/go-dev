package user

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// Operations about goods
type ApiUserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Success 200 {object} models.User
// @router /v1/api/goods/getNewstGoodsList [get]
func (u *ApiUserController) GetNewstGoodsList(ctx *beecontext.Context) {

	fmt.Println("GetNewstGoodsList")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "GetNewstGoodsList",
	})

}
