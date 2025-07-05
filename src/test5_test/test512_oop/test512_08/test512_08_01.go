package main

import "fmt"

type Person81 struct {
	Name string
}

// 值接收者
func (p Person81) SayHello81() {
	fmt.Println("Hello,", p.Name)
}

// 指针接收者
func (p *Person81) SetName81(name string) {
	p.Name = name
}

// 测试对象的方法集
//| 类型       | 方法集                         |
//| --------- | ----------------------------- |
//| `Person`  | `SayHello()`                  |
//| `*Person` | `SayHello()`, `SetName(name)` |
func main() {

	//p := Person81{Name: "张三"}
	//fmt.Printf("值： %#v \n", p)
	//p.SayHello81()
	//p.SetName81("我是张三")
	//fmt.Printf("值： %#v \n", p)
	//
	//
	//fmt.Println()
	//
	//pp := &Person81{Name: "里斯"}
	//fmt.Printf("值： %#v \n", pp)
	//pp.SayHello81()
	//pp.SetName81("我是里斯")
	//fmt.Printf("值： %#v \n", pp)


	var p Person81 = Person81{Name: "张三"}
	var pp *Person81 = &Person81{Name: "里斯"}

	// 都可以赋值给 Helloer，因为 SayHello 是值接收者方法
	var h1 Helloer = p
	var h2 Helloer = pp

	h1.SayHello81()
	h2.SayHello81()


	// ❌ 编译错误：Person（值类型）没有实现 SetName（是指针接收者）
	//var s1 Setter = p // 无法将 'p' (类型 Person81) 用作类型 Setter 类型未实现 'Setter'，因为 'SetName81' 方法有指针接收器

	// ✅ 正确：*Person 包含指针接收者方法
	var s2 Setter = &p
	s2.SetName81("你好 Golang")


}

//1. 实现接口时要注意方法集匹配
type Helloer interface {
	SayHello81()
}


type Setter interface {
	SetName81(string)
}