package main

import "fmt"

// 匿名字段
// https://topgoer.com/方法/匿名字段.html
// https://www.cnblogs.com/jinyanshenxing/p/15939484.html
func main1() {
	user := User{id: 666, name: "zhangsan"}
	m := Manager{user}

	fmt.Printf("user: %p\n", &user)
	user.ToString()
	fmt.Printf("Manager: %p\n", &m)
	m.ToString()
}

type User struct {
	id   int
	name string
}

type Manager struct {
	// 以下的User就是匿名字段
	// 不使用一个变量来表示的数据类型，那么这个数据类型在这个结构体里面就是一个匿名字段
	User
}

func (self *User) ToString() {
	fmt.Printf("User: %p, %v\n", self, self)
}

//func (self *Manager) ToString() {
//	fmt.Printf("Manager: %p, %v\n", self, self)
//}
