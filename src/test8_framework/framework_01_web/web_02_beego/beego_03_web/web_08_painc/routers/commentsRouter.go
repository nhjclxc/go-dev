package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["web_08_painc/controllers:UserController"] = append(beego.GlobalControllerRouter["web_08_painc/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/getById/:userId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
