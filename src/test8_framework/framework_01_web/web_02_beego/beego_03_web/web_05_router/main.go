package main

import (
	_ "web_05_router/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// web_05_router 将前面router 相关的所有知识点做一个最佳实践,
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
