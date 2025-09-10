package main

import (
	"github.com/jordan-wright/email" // go get github.com/jordan-wright/email
	"log"
	"net/smtp"
)

// email测试
// https://topgoer.com/其他/发邮件.html
func main() {
	// "github.com/jordan-wright/email"
	// go get github.com/jordan-wright/email
	// https://pkg.go.dev/github.com/jordan-wright/email

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <XXX@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"XXX@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "你的邮箱账号", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
