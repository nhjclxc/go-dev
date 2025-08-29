package main

import (
	_ "config_01/routers"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/beego/beego/v2/core/config/yaml"

	beego "github.com/beego/beego/v2/server/web"
)

// 如何读取配置文件
func main01() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 读取配置文件数据

	// 1、使用全局实例
	// Beego 默认会解析当前应用下的 conf/app.conf 文件，后面就可以通过config包名直接使用
	// 使用 config.String 读取的默认是 conf/app.conf 文件

	// 没有Section的配置，直接 String、Int等方法获取
	appname, err := config.String("appname")
	if err != nil {
		fmt.Println("appname 配置文件读取出错：" + err.Error())
		return
	}
	fmt.Println("appname = ", appname)

	// 有Section的配置，先获取整个 section 的配置，再读取里面的详细配置
	section, err := config.GetSection("mysql")
	if err != nil {
		return
	}
	fmt.Println("mysql.ip: ", section["ip"])
	fmt.Println("mysql.port: ", section["port"])
	fmt.Println("mysql.anonymous_user: ", section["anonymous_user"])

	fmt.Println("-----------------------")

	// 2、使用使用Configer实例
	// 如果要从多个源头读取配置，或者说自己不想依赖于全局配置，那么可以自己初始化一个配置实例：
	cfgIni, err := config.NewConfig("ini", "conf/local.ini")
	if err != nil {
		fmt.Println("config.NewConfig 配置文件读取出错：" + err.Error())
	}
	password, _ := cfgIni.String("password")
	fmt.Println("auto load config name is", password)

	// 读取 yaml 配置
	// 由于 Beego 的 config 模块里，支持不同格式的配置文件，ini、json、yaml,但是！YAML 支持是独立插件，需要显式导入对应的适配器包。
	// 因此,要想 beego 能够读取到 yaml 类型的配置文件,则必须导入其包,如下
	//  _ "github.com/beego/beego/v2/core/config/yaml"
	// 注意前面有个下划线 _，表示只导入初始化注册，不直接用函数。因为适配器注册是靠 init() 函数自动注册的，需要你把包引入，才能把 "YAML" 注册到适配器工厂里。
	cfgYaml, err := config.NewConfig("yaml", "conf/config-dev.yaml")
	if err != nil {
		// config.NewConfig 配置文件读取出错 yaml：config: unknown adaptername "yaml" (forgotten import?),,,就是没导入yaml适配器导致的
		fmt.Println("config.NewConfig 配置文件读取出错 yaml：" + err.Error())
	}
	redis, _ := cfgYaml.GetSection("redis")
	fmt.Println("host = ", redis["host"])
	fmt.Println("port = ", redis["port"])
	fmt.Println("auth = ", redis["auth"])

	// 以上方法适用于 beego/v2 以前的版本,

	// 而 beego/v2 推荐使用下面的方法

	// 先把配置文件注册加载进来
	err = beego.LoadAppConfig("yaml", "conf/config-dev.yaml")
	if err != nil {
		// config.NewConfig 配置文件读取出错 yaml：config: unknown adaptername "yaml" (forgotten import?)
		fmt.Println("config.NewConfig 配置文件读取出错 yaml：" + err.Error())
		return
	}
	// 读取加载的配置文件
	fmt.Println("beego.AppConfig[appname] = ", beego.AppConfig.DefaultString("appname", "aaa"))

	redis2, err := beego.AppConfig.GetSection("redis")
	if err != nil {
		fmt.Println("beego.AppConfig.GetSection(\"redis\")：" + err.Error())
		return
	}
	fmt.Println("host = ", redis2["host"])
	fmt.Println("port = ", redis2["port"])
	fmt.Println("auth = ", redis2["auth"])

	//beego.Run()
}
