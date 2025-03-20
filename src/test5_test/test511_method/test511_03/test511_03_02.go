package main

import "fmt"

func main() {
	/*
	   	Golang 表达式 ：根据调用者不同，方法分为两种表现形式:

	       instance.method(args...) ---> <type>.func(instance, args...)
	      	前者称为 method value，后者 method expression。

	      两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
	*/

	u := User{1, fmt.Sprintf("Tom")}
	u.Test()

	// 需要注意，method value 会复制 receiver。
	mValue := u.Test // // 立即复制 receiver，因为不是指针类型，不受后续修改影响。
	mValue()         // 隐式传递 receiver

	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver

	// 测试 ，method value 会复制 receiver。特性
	u.id = 666
	u.Test()
	mValue()

	u2 := User{1, "Tom"}
	mValue2 := u2.Test // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u2.id, u2.name = 2, "Jack"
	u2.Test()

	mValue2()

}

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}
