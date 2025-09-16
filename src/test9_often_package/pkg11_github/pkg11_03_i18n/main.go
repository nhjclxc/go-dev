package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// go get github.com/nicksnyder/go-i18n/v2/goi18n
// go get github.com/nicksnyder/go-i18n/v2/i18n
func main() {

	// 创建本地化包
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	path := "/Users/lxc20250729/lxc/code/go-dev/src/test9_often_package/pkg11_github/pkg11_03_i18n/active"
	// 加载翻译文件
	bundle.MustLoadMessageFile(path + "/active.en.json")
	bundle.MustLoadMessageFile(path + "/active.zh.json")

	r := gin.Default()

	// http://127.0.0.1:8090/sayHello?lang=zh&name=%E5%BC%A0%E4%B8%89
	// http://127.0.0.1:8090/sayHello?lang=en&name=%E5%BC%A0%E4%B8%89
	r.GET("/sayHello", func(ctx *gin.Context) {

		lang := ctx.Query("lang")
		name := ctx.Query("name")
		var localizer *i18n.Localizer
		switch lang {
		case "en":
			localizer = i18n.NewLocalizer(bundle, "en")
		case "zh":
			localizer = i18n.NewLocalizer(bundle, "zh")
		default:
			ctx.JSON(500, gin.H{"msg": "暂不支持的语言"})
			return
		}
		// 选择语言（实际项目里可根据 Accept-Language 或用户设置）
		//localizer := i18n.NewLocalizer(bundle, "en")
		//localizer := i18n.NewLocalizer(bundle, "zh")

		Addr := "北京"
		Year := "2025"

		// 翻译
		msg, _ := localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "hello",
			TemplateData: map[string]string{
				"Name": name,
				"Addr": Addr,
				"Year": Year,
			},
		})
		fmt.Println(msg) // 输出: 你好，小明！

		ctx.JSON(200, gin.H{
			"msg":  "操作成功！！！",
			"data": msg,
		})

	})

	r.Run(":8090")
}
