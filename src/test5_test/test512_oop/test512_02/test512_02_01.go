package main

import "fmt"

func main1() {

	stu := Student1{
		age: 18,
		Person1: Person1{
			name: "你的名字",
		},
	}

	fmt.Println(stu)

	stu.sayHello()

}

type Person1 struct {
	name string
}

type Student1 struct {
	age int
	Person1
}

func (this *Person1) sayHello() {
	fmt.Println("Person.sayHello ", this.name)
}

// 模拟方法继承
// 当结构体和匿名结构体有相同的字段或者方法时，编译器采用 就近原则 访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
func (this *Student1) sayHello() {
	fmt.Println("Student.sayHello ", this.name, this.age)
}
