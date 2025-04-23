package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	/*
		pipeline
		pipeline是指产生数据的操作。比如{{.}}、{{.Name}}等。Go的模板语法中支持使用管道符号|链接多个命令，
		用法和unix下的管道类似：|前面的命令会将运算结果(或返回值)传递给后一个命令的最后一个位置。

		注意 : 并不是只有使用了|才是pipeline。Go的模板语法中，pipeline的概念是传递数据，只要能产生数据的，都是pipeline。
	*/

	http.HandleFunc("/", sayHello03)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func sayHello03(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("src/test5_test/test523_template/template_02_html/html03.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入w
	user := map[string]any{
		"Name":   "枯藤",
		"Gender": "男",
		"Age":    18,
	}

	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, user)
}
