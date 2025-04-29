package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	_ "web_08_painc/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// web_08_painc 学习 Beego 框架的 错误处理 机制
// https://beegodoc.com/zh/developing/web/error/
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 从 panic 中恢复
	//如果你希望用户在服务器处理请求过程中，即便发生了 panic 依旧能够返回响应，那么可以使用 Beego 的恢复机制。该机制是默认开启的。依赖于配置项：
	beego.BConfig.RecoverPanic = true

	// 自定义panic之后的处理行为，那么可以重新设置web.BConfig.RecoverFunc
	// 类似于Java中的全局异常处理
	beego.BConfig.RecoverFunc = func(context *context.Context, config *beego.Config) {
		if err := recover(); err != nil {
			context.JSONResp(map[string]any{
				"code": 500,
				"success": false,
				"msg": fmt.Sprintf("全局异常处理 you panic, err: %v", err),
			})
		}
	}


	// Admin 管理后台
	// https://beegodoc.com/zh/developing/web/admin/
	beego.BConfig.Listen.EnableAdmin = true



	beego.Run()
}
