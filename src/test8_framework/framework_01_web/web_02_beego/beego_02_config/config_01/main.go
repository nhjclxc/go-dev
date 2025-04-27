package main

import (
	_ "config_01/routers"
	"fmt"
	"github.com/beego/beego/v2/core/config"

	// 读取json配置的驱动实现
	_ "github.com/beego/beego/v2/core/config/json"

	beego "github.com/beego/beego/v2/server/web"
)

// 读取json配置
// 详细看：https://beegodoc.com/zh/developing/config/#%E6%94%AF%E6%8C%81%E7%9A%84%E6%A0%BC%E5%BC%8F
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 读取不同环境的配置文件

	// 而 beego/v2 推荐使用下面的方法

	// 1. 先读取主配置文件的运行模式 runmode
	err := beego.LoadAppConfig("json", "conf/config.json")
	if err != nil {
		fmt.Printf("配置文件读取失败:config.NewConfig 配置文件读取出错 yaml： %v\n", err)
	}
	runmode, err := beego.AppConfig.String("runmode")
	fmt.Println("runmode = ", runmode)

	// 方式2
	err2 := config.InitGlobalInstance("json", "conf/config.json")
	if err2 != nil {
		fmt.Printf("配置文件读取失败:config.NewConfig 配置文件读取出错 yaml： %v\n", err2)
		panic(err2)
	}

	runmode2, _ := config.String("runmode")
	fmt.Println("runmode2 = ", runmode2)

	//beego.Run()
}
