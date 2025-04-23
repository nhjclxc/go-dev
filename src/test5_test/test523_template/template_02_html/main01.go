package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	/*
				html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用text/template包。

			在基于MVC的Web架构中，我们通常需要在后端渲染一些数据到HTML文件中，从而实现动态的网页效果。

		通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。模板中的注释引用数据接口的元素（一般如结构体的字段或者字典的键）来控制执行过程和获取需要呈现的值。模板执行时会遍历结构并将指针表示为’.‘（称之为”dot”）指向运行过程中数据结构的当前位置的值。

		用作模板的输入文本必须是utf-8编码的文本。”Action”—数据运算和控制单位—由”“界定；在Action之外的所有文本都不做修改的拷贝到输出中。Action内部不能有换行，但注释可以有换行。

	*/

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("src/test5_test/test523_template/template_02_html/html01.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "baidu.com")
}
