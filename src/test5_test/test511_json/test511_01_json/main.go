package main

import (
	"encoding/json"
	"fmt"
)

// https://denganliang.github.io/the-way-to-go_ZH_CN/12.9.html
/*
   	数据结构 –> 指定格式 = 序列化 或 编码（传输之前）
    指定格式 –> 数据结构 = 反序列化 或 解码（传输之后）
      序列化是在内存中把数据转换成指定格式（数据 -> 字符串），反之亦然（字符串 -> 数据）。
*/

type Address struct {
	Type    string
	City    string
	Country string
	Num     int
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {

	add1 := &Address{
		Type:    "aaa",
		City:    "东京都",
		Country: "小日子",
		Num:     666,
	}
	add2 := &Address{
		Type:    "bbb",
		City:    "拉斯维加斯",
		Country: "漂亮国",
		Num:     888,
	}

	vc := VCard{
		FirstName: "Luo",
		LastName:  "Chao",
		Addresses: []*Address{add1, add2},
		Remark:    "不知道中国十八于是备注",
	}

	// 序列化 = 对象 -> 字符串
	vcJsonByte, err := json.Marshal(vc)
	vcJsonStr := string(vcJsonByte)
	fmt.Println(err)
	fmt.Println(vcJsonByte)
	fmt.Println(vcJsonStr)

	//// 将 vcJson 写入文件
	//file, _ := os.OpenFile("vcJson.json", os.O_CREATE|os.O_RDWR, 0666)
	//defer file.Close()
	//
	//file.Write(vcJsonByte)

	// 反序列化 = 字符串 -> 对象
	// 情况1，已知某个json字符串的对象了，先创建对象，后将json字符串序列化到中国对象里面
	unVC := VCard{} // 先声明一个对象来接收反序列的数据
	fmt.Println(unVC)
	err2 := json.Unmarshal(vcJsonByte, &unVC)
	fmt.Println(err2)
	fmt.Println(unVC)

	// 情况2：还不知道json字符串是哪个对象，这个时候就需要使用一个接口对象（类似Java的Object）来接收这个序列化后的对象
	var f interface{}
	fmt.Println(f)
	err3 := json.Unmarshal(vcJsonByte, &f)
	fmt.Println(err3)
	fmt.Println(f) // map[Addresses:[map[City:东京都 Country:小日子 Type:aaa] map[City:拉斯维加斯 Country:漂亮国 Type:bbb]] FirstName:Luo LastName:Chao Remark:不知道中国十八于是备注]

	// f.() ，这个是类型断言的用法
	//src/test5_test/test512_oop/test512_02/test512_02_04.go
	//src/test5_test/test512_oop/test512_02/test513_02_05.go
	jsonMap := f.(map[string]interface{})

	fmt.Println(jsonMap["FirstName"])
	fmt.Println(jsonMap["LastName"])
	fmt.Println(jsonMap["Addresses"])

	println("---------------------------------")
	// 我们可以通过 for range 语法和 type switch 来访问其实际类型：
	for key, val := range jsonMap {
		fmt.Println(key)
		if intVal, ok := val.(int); ok {
			fmt.Println("intVal = ", intVal)
		}
		if floatVal, ok := val.(float64); ok {
			fmt.Println("floatVal = ", floatVal)
		}
		if stringVal, ok := val.(string); ok {
			fmt.Println("stringVal = ", stringVal)
		}
		if AddressesVal, ok := val.([]Address); ok {
			fmt.Println("AddressesVal = ", AddressesVal)
		}
		if AddressesVal, ok := val.([]*Address); ok {
			fmt.Println("AddressesVal = ", AddressesVal)
		}

	}

}
