package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"os"
)

// Gin 中使用go-ini 加载.ini 配置文件
// go-ini Github地址：https://github.com/go-ini/ini
// go-ini 官方文档：https://ini.unknwon.io/， https://gowalker.org/gopkg.in/ini.v1

// go get -u github.com/go-ini/ini
func main() {

	//读取配置
	config, err := ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Name:", config.Section("app").Key("app_name").String())
	fmt.Println("App http_port:", config.Section("app").Key("http_port").String())
	fmt.Println("log_path:", config.Section("log").Key("log_path").String())

	// 我们可以做一些候选值限制的操作
	fmt.Println("Server Protocol:",
		config.Section("server").Key("protocol").In("http", []string{"http", "https"}))

	// 差不多了，修改某个值然后进行保存
	config.Section("log").Key("log_level").SetValue("INFO")
	config.SaveTo("./config/app.ini.local")

	// 创建路由器
	router := gin.Default()

	// go-ini 入门操作
	// https://ini.unknwon.io/docs/intro/getting_started

	// 注册路由
	router.GET("/ini", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"data": "Gin 中使用go-ini 加载.ini 配置文件",
		})
	})

	// 启动
	router.Run(":8090")

}
