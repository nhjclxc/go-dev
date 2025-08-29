package user

import "fmt"

type User struct {
	ID   int
	Name string
}

// 以下是User这个结构体的实例方法，
func (u *User) List() []User {
	return []User{{ID: 1, Name: "test"}}
}
func (u *User) PrintUser() {
	fmt.Printf("User{id = %d, name = %s} \n", u.ID, u.Name)
}

func init() {
	fmt.Println("anonymous_user.init")
}

// 以下称为一个函数
func List2(u *User) []User {
	return []User{{ID: 1, Name: "test"}}
}
