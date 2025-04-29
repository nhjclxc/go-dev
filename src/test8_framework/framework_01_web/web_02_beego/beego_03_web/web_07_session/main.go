package main

import (
	_ "web_07_session/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// beego - Session 的学习，https://beegodoc.com/zh/developing/web/session/
// beego - Cookie 的学习，https://beegodoc.com/zh/developing/web/cookie/
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 在 web模块中使用 session 相当方便，只要在 main 入口函数中设置如下：
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
