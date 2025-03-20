package main

import "fmt"

// 结构体模拟实现其他语言的“继承”
// https://www.cnblogs.com/jinyanshenxing/p/15939484.html
func main() {

	animal := &Animal{"zhangsan"}
	fmt.Println(animal)

	dog := &Dog{
		age:    18,
		Animal: Animal{"lisi"},
	}
	dog.feed("aaa")
	dog.wangwangwang("qqq")
	fmt.Println(dog)

}

type Animal struct {
	name string
}
type Dog struct {
	age int
	// 使用匿名字段来模拟属性继承和方法继承
	Animal
}

// 用来模型方法继承
func (self *Animal) feed(food string) string {
	fmt.Printf("feed, name = %s, feed = %s \n", self.name, food)

	return "返回数据"
}

func (self *Dog) wangwangwang(food string) string {
	fmt.Printf("dog, name = %s, wangwangwang !!!\n", self.name, food)

	return "wangwangwang 返回数据"
}
