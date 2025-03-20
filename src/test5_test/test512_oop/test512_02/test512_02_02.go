package main

import "fmt"

// 面向对象编程-多重继承
// 如一个 struct 嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。
func main2() {

	t2 := Teacher2{
		ClassNo: 666,
		Person2: Person2{
			Name: "张三",
		},
		Work2: Work2{
			Wages: 123456,
		},
	}

	fmt.Println(t2)

	t2.SayHello()
	t2.DoWork()
	t2.Lessons()

}

type Person2 struct {
	Name string
}

func (this *Person2) SayHello() {
	fmt.Println("Person2.sayHello ", this.Name)
}

type Work2 struct {
	Wages int
}

func (this *Work2) DoWork() {
	fmt.Println("Work2.doWork ", this.Wages)
}

type Teacher2 struct {
	ClassNo int
	// 以下嵌入多重继承
	Person2
	Work2
}

func (this *Teacher2) Lessons() {
	fmt.Println("Teacher.Lessons ", this.ClassNo)
}
