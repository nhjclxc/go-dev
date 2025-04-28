package main

import (
	_ "web_03_router/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// [web模块-路由-Web 命名空间](https://beegodoc.com/zh/developing/web/router/namespace.html)
func main() {
	/*

		启动步骤
		1. 创建项目：`bee api web_03_router`
		2. 进入 `cd web_03_router` 目录
		3. 更新包：`go mod tidy`
		5. 运行项目执行 `bee run`
	*/

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
