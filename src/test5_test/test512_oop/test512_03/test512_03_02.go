package main

import "fmt"

func main2() {
	// 无法将 'Student{}' (类型 Student) 用作类型 People 类型未实现 'People'，因为 'Speak' 方法有指针接收器
	// 接口变量只能保存实例的地址，不能直接赋值为实例对象的变量
	//var peo People = Student{}

	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

type People interface {
	Speak(string) string
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
