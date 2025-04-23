package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {

	/*
		   在 Go 语言中，text/template 是标准库提供的一个强大的模板引擎，主要用于生成纯文本（如配置文件、Markdown 文档、纯文本报告等）。
		   它通过定义模板语法并将数据渲染进去，从而生成最终文本内容。

		   下面是对 text/template 的详解，包括基本用法、常见语法、函数扩展和嵌套模板等方面。

		   import "text/template"


		   模板的使用一般分为三步：
			   定义模板（使用字符串或文件）
			   解析模板（template.New + Parse / ParseFiles / ParseGlob）
			   执行模板（Execute 或 ExecuteTemplate）

		   三、模板语法
		   1. 变量
			   {{.FieldName}}       // 访问字段
			   {{.}}                // 代表当前对象

		   2. 条件判断
			   {{if .IsAdmin}}Welcome Admin{{end}}
			   {{if .Name}}{{.Name}}{{else}}Anonymous{{end}}

		   3. 循环
			   {{range .Users}}
			   Name: {{.Name}}, Age: {{.Age}}
			   {{end}}

		   4. 变量定义
			   {{ $x := .Value }}
			   {{$x}}


	*/

	//text01()

	text02()

	/*
		嵌套模板（template 定义）
			const tmpl = `
			{{define "header"}}Header Part{{end}}
			{{define "footer"}}Footer Part{{end}}
			{{define "body"}}
			  {{template "header"}}
			  Hello, {{.Name}}
			  {{template "footer"}}
			{{end}}
			`

			t := template.Must(template.New("main").Parse(tmpl))
			t.ExecuteTemplate(os.Stdout, "body", map[string]string{"Name": "Bob"})


	*/

	/*
		从文件读取模板
			t := template.Must(template.ParseFiles("template.txt"))
			t.Execute(os.Stdout, data)

		t := template.Must(template.ParseGlob("templates/*.tmpl"))


	*/
}

// 函数模板
func text02() {

	// 定义要执行的数据
	funcMap := map[string]any{
		"toUpper": strings.ToUpper,
		"add":     func(a, b int) int { return a + b },
	}

	// 1、定义模板
	const tmplStr = `{{ toUpper .Name}} is {{ add .Age 1}} years old.`

	// 2、创建模板
	tmpl := template.Must(template.New("tmplName").Funcs(funcMap).Parse(tmplStr))

	// 定义要传入的数据
	data := map[string]any{
		"Name": "zhangsan",
		"Age":  18,
	}

	// 3、执行模板
	tmpl.Execute(os.Stdout, data)

}

// 字符串模板
func text01() {
	// 1、定义模板 ， 注意：使用的是模板字符串符号``，并非引号''
	const tmpl = `Hello, {{.Name}}! You are {{.Age}} years old.   {{.}}`

	// 2、创建模板并解析
	t := template.Must(template.New("example").Parse(tmpl))

	// 要传入的数据
	data := map[string]interface{}{
		"Name": "Alice",
		"Age":  30,
	}

	// 3、执行模板，输出到标准输出
	t.Execute(os.Stdout, data)
}
