package main

import "fmt"

// 匿名字段
// https://topgoer.com/方法/匿名字段.html
func main() {
	user := User{id: 666, name: "zhangsan"}
	m := Manager{user}

	fmt.Printf("user: %p\n", &user)
	user.ToString()
	fmt.Printf("Manager: %p\n", &m)
	m.ToString()
	m.user.ToString()
}

type User struct {
	id   int
	name string
}

type Manager struct {
	user User
}

func (self *User) ToString() {
	fmt.Printf("User: %p, %v\n", self, self)
}

func (self *Manager) ToString() {
	fmt.Printf("Manager: %p, %v\n", self, self)
}
