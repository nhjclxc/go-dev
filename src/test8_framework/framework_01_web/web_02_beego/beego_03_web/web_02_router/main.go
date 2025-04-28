package main

import (
	_ "web_02_router/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// web_02_router 主要学习的是  注册函数式风格路由注册(https://beegodoc.com/zh/developing/web/router/functional_style/)
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
