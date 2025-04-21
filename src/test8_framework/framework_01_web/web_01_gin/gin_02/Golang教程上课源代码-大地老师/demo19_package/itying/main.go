/*
1、go mod init 项目名称             表示：初始化项目
2、配置第三方包
3、go mod tidy                      表示：下载当前项目缺少的依赖
4、运行项目
*/
package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("你好go")

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

	value := gjson.Get(json, "name.last")
	println(value.String()) //Prichard
}
