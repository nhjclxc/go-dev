package user

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// Operations about goods
type AdminGoodsController struct {
	beego.Controller
}

// @Title Get
// @Description get anonymous_user by uid
// @Success 200 {object} models.User
// @router /v1/api/goods/getNewstGoodsList [get]
func (u *AdminGoodsController) InsertGoods(ctx *beecontext.Context) {

	fmt.Println("InsertGoods")

	ctx.JSONResp(map[string]any{
		"code": 200,
		"data": "InsertGoods",
	})

}
