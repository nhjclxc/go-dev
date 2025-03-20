package main

import "fmt"

func main3() {

	stu := Studnet{name: "zxcvb"}

	fmt.Println(stu)
	stu.SayHi()

	//  MyInterface 变量
	var mi MyInterface = &stu // 接口变量必须赋值指针
	fmt.Println(mi)
	mi.SayHi()

}

type MyInterface interface {
	SayHi()
}

type Studnet struct {
	name string
}

// Student 实现 MyInterface 的 SayHi() 方法
func (this *Studnet) SayHi() {
	fmt.Printf("Studnet.SayHi.name = %s \n", this.name)
}

func (this *Studnet) DoWork() {
}

// go里面没有方法重载
//func (this *Studnet) DoWork(str string)  {
//}
