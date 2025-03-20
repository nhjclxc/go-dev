package main

import "fmt"

//工厂模式
//Golang 的结构体没有构造函数，通常可以使用工厂模式来解决这个问题。
// 因为这里的 Student 的首字母 S 是大写的，如果我们想在其它包创建 Student 的实例(比如 main 包)，引入 model 包后，就可以直接创建 Student 结构体的变量(实例)。
//但是问题来了，如果首字母是小写的，比如 是 type student struct {....} 就不不行了，怎么办---> 工厂模式来解决.

// 工厂模式为了解决 结构体私有化问题，即私有化的结构体对象在外部包不可创建对象，那么就通过工厂模式。
func main() {

	// 使用工厂模式创建一个 person4 对象
	p4 := GetPerson4Instance(666, "你的名字！！！")

	fmt.Println(p4)
	fmt.Println(*p4)
}

type person4 struct {
	Id   int
	Name string
}

// 使用工厂模型创建一个私有化person4的实例对象
func GetPerson4Instance(id int, name string) *person4 {
	return &person4{
		Id:   id,
		Name: name,
	}
}
