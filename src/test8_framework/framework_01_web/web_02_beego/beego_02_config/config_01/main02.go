package main

import (
	_ "config_01/routers"
	"fmt"
	_ "github.com/beego/beego/v2/core/config/yaml"

	beego "github.com/beego/beego/v2/server/web"
)

// 如何读取不同环境的配置文件
func main02() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 读取不同环境的配置文件

	// 而 beego/v2 推荐使用下面的方法

	// 1. 先读取主配置文件的运行模式 runmode
	//mainConfig, err := config.NewConfig("yaml", "conf/config.yaml")
	err := beego.LoadAppConfig("yaml", "conf/config.yaml")
	if err != nil {
		fmt.Printf("配置文件读取失败:config.NewConfig 配置文件读取出错 yaml： %v\n", err)
	}
	//runmode, _ := mainConfig.String("runmode")
	runmode, err := beego.AppConfig.String("runmode")
	fmt.Println("runmode = ", runmode)
	appname, err := beego.AppConfig.String("appname")
	fmt.Println("appname = ", appname)

	var configFile string
	switch runmode {
	case "dev":
		// go run main.go -env=dev
		configFile = "conf/config-dev.yaml"
	case "test":
		// go run main.go -env=test
		configFile = "conf/config-test.yaml"
	case "prod":
		// go run main.go -env=prod
		configFile = "conf/config-prod.yaml"
	default:
		panic("未知环境参数！必须是 dev/test/prod")
	}

	// 2. 接着从不同环境中获取具体配置详细
	// 从新加载对应环境的配置，注意：本操作会将著配置 "conf/config.yaml" 里面的所有配置覆盖
	// 在程序运行中，如需使用主配置里面的相关配置信息，应当在内存中做配置信息备份处理，防止重新加载配置文件之后无法读取主配置文件里面的配置信息
	err2 := beego.LoadAppConfig("yaml", configFile)
	if err2 != nil {
		// config.NewConfig 配置文件读取出错 yaml：config: unknown adaptername "yaml" (forgotten import?)
		fmt.Printf("配置文件读取失败:config.NewConfig 配置文件读取出错 yaml-runmode： %v\n", err2)
		return
	}

	// 主配置里面，独有的配置信息被环境配置覆盖了
	appname2, err := beego.AppConfig.String("appname")
	fmt.Println("appname2 = ", appname2)

	// 读取加载的配置文件
	mysqlConfig, err3 := beego.AppConfig.GetSection("mysql")
	if err3 != nil {
		fmt.Printf("配置文件读取失败:beego.AppConfig.GetSection(\"mysql\")： %v\n", err3)
		return
	}
	fmt.Println("dbname = ", mysqlConfig["dbname"])

	// 直接读取多级的配置文件
	dbname2, err := beego.AppConfig.String("mysql.dbname")
	if err != nil {
		return
	}
	fmt.Println("dbname2 = ", dbname2)

	//beego.Run()
}
