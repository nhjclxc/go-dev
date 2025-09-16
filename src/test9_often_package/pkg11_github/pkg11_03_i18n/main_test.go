package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"testing"
)

func TestName(t *testing.T) {

	//bundle := i18n.NewBundle(language.English)
	//// 使用标准库的 json.Unmarshal 作为反序列化器
	//bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	//
	//// 如果你用文件，可以这样加载（文件需符合 go-i18n 支持的消息文件格式）
	//// bundle.MustLoadMessageFile("active.en.json")
	//// bundle.MustLoadMessageFile("active.zh.json")
	//
	//// 或者直接在代码里添加消息，避免文件格式困惑
	//bundle.AddMessages(language.Make("en"), &i18n.Message{
	//	ID:    "hello",
	//	Other: "Hello, {{.Name}}!",
	//})
	//bundle.AddMessages(language.Make("zh"), &i18n.Message{
	//	ID:    "hello",
	//	Other: "你好，{{.Name}}！",
	//})
	//
	//localizer := i18n.NewLocalizer(bundle, "zh")
	//msg, _ := localizer.Localize(&i18n.LocalizeConfig{
	//	MessageID:    "hello",
	//	TemplateData: map[string]string{"Name": "小明"},
	//})
	//fmt.Println(msg) // 你好，小明！

	// 创建本地化包
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	path := "/Users/lxc20250729/lxc/code/go-dev/src/test9_often_package/pkg11_github/pkg11_03_i18n/active"
	// 加载翻译文件
	bundle.MustLoadMessageFile(path + "/active.en.json")
	bundle.MustLoadMessageFile(path + "/active.zh.json")

	// 选择语言（实际项目里可根据 Accept-Language 或用户设置）
	localizer := i18n.NewLocalizer(bundle, "en")
	//localizer := i18n.NewLocalizer(bundle, "zh")

	// 翻译
	msg, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "hello",
		TemplateData: map[string]string{
			"Name": "小明",
		},
	})
	fmt.Println(msg) // 输出: 你好，小明！
}
