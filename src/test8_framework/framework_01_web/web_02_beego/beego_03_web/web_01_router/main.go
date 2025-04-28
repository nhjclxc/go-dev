package main

import (
	_ "web_01_router/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// web_01_router 主要学习的是 https://beegodoc.com/zh/developing/web/router/ctrl_style/ 链接里面的内容

// beego 项目的入口函数
func main() {

	// 检查当前 Beego 应用的运行模式（RunMode）是不是 "dev"，也就是开发模式。
	if beego.BConfig.RunMode == "dev" {
		// 开启静态资源映射
		// 允许目录索引浏览。
		// 当访问一个 URL 映射到服务器上的一个文件夹，而不是具体文件时，如果 DirectoryIndex = true，服务器会自动列出该目录下的所有文件。
		beego.BConfig.WebConfig.DirectoryIndex = true
		// 将静态资源 swagger 文件夹 映射到 /swagger 这个url路径下
		// 把 URL 路径 /swagger 映射到项目目录下的 swagger 这个文件夹。
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 启动 Beego 框架的 Web Server。
	//beego.Run()
	beego.Run(":8080")

}
