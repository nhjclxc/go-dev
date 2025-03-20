package main

import (
	"encoding/json"
	"fmt"
)

// 实则为 json 转 对象obj 练习
func main3() {
	/*
			GO 服务器：
				接收自客户端数据：结构体对象 ===>>> json字符串
				给客户端返回数据：json字符串 ===>>> 结构体对象

		要使用一个包 import "encoding/json"
	*/

	p2 := Person3{
		Id:   666,
		Name: "利好世界",
		Arr:  []int{1, 2, 3, 4, 5, 6},
		Dict: map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},
	}

	fmt.Println(p2)

	//// 将person变量序列化为 json 格式字串
	//    // json.Marshal 函数中使用反射
	jsonByte, err := json.Marshal(p2) // 第一个返回值为json字符串的byte字节数组，第二个为序列化时出错返回的异常信息
	if err != nil {
		fmt.Println("json处理错误，", err)
	}
	fmt.Println(jsonByte)
	// 使用 string() 直接将字节数组转化为字符串
	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

	// 将json字符串 反序列化回对象
	p21 := Person3{}
	fmt.Println(p21)
	//jsonStr2 := "{\"Id\":888,\"name\":\"我就是女王\",\"arr2\":[1,2,3,4,5,6,7,8,9],\"Dict\":{\"key1111\":1,\"key2\":2,\"key3\":3}}"
	//jsonStr2 := "{\"I111d\":888,\"name\":\"我就是女王\",\"arr2\":[1,2,3,4,5,6,7,8,9],\"Dict\":{\"key1111\":1,\"key2\":2,\"key3\":3}}"
	jsonStr2 := "{\"I111d888,\"name\":\"我就是女王\",\"arr2\":[1,2,3,4,5,6,7,8,9],\"Dict\":{\"key1111\":1,\"key2\":2,\"key3\":3}}"
	err2 := json.Unmarshal([]byte(jsonStr2), &p21)
	if err2 != nil {
		fmt.Println("json反序列化处理错误，", err2)
	}
	fmt.Printf("反序列化后结果：%p, %v \n", &p21, p21)
	p21.Id = 123
	fmt.Println(p21)

	fmt.Println(&p2) // 输出String方法

}

// 4、struct 的每个字段上，可以写上一个 tag, 该 tag 可以通过反射机制获取，常见的使用场景就是序列化和反序列化。
type Person3 struct {
	Id int
	// 序列化为json的时候的字段名为name，
	// 通过json反序列化为对象的时候的名称为Name
	Name string `json:"name"` // `json:"name"`  就是 struct tag
	Arr  []int  `json:"arr2"`
	Dict map[string]int
	// 注意：如果第一个字母是小写的，即这个字段是外部不可见的，序列化的时候也无法导出
	// 即，要想把某个属性导出，那么就必须设置为可见，即第一个字母为大写
	phone string
}

// 5、如果一个类型实现了 String()这个方法，那么 fmt.Println 默认会调用这个变量的 String()进行输出：
func (this *Person3) String() string {
	// 格式化这个字符串，模拟Java的 toString() 方法
	return fmt.Sprintf("Person3.String{Id=%d, Name=%s, Arr=%v, Dict=%v, phone=%s}\n",
		this.Id, this.Name, this.Arr, this.Dict, this.phone)
}
