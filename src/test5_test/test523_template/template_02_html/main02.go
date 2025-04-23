package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", sayHello02)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func sayHello02(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("src/test5_test/test523_template/template_02_html/html02.html")
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
