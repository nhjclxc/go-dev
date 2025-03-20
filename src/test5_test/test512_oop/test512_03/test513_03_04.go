package main

import "fmt"

func main4() {

	p := Phone{
		Name: "IPhone",
	}
	p.Run()

	// func (this *Phone) Run()
	var usbPhone USB = &p
	usbPhone.Run()

	c := Computer{
		Name: "HP Computer",
	}
	c.Run()
	// func (self Computer) Run()
	var usbComputer USB = c
	usbComputer.Run()

	// 从以下的两条赋值语句我们可以看出
	// var usbPhone USB = &p：
	//		如果结构体在实现一个接口的时候，它的接收者是 指针类型 ，那么将其实例对象赋值给接口变量的时候也必须是 指针类型 的地址
	// var usbComputer USB = c：
	//		如果一个结构体在实现一个接口的时候，它的接收者是 变量类型 ，那么将其实例对象赋值给接口变量的时候也必须是 变量类型

}

type USB interface {
	Run()
}
type Phone struct {
	Name string
}

func (this *Phone) Run() {
	fmt.Printf("Phone.Run.Name = %s \n", this.Name)
}

type Computer struct {
	Name string
}

func (self Computer) Run() {
	fmt.Printf("Computer.Run.Name = %s \n", self.Name)
}
