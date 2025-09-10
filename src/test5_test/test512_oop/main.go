package main

import (
	"fmt"
	"reflect"
)

// 两个结构体可以等值比较吗

type User1 struct {
	Name string
}
type User2 struct {
	Name string
}

func main() {

	u1 := User1{Name: "zhangsan"}
	u11 := User1{Name: "zhangsan"}
	u2 := User2{Name: "zhangsan"}

	fmt.Printf("u = %#v \n", u1)
	fmt.Printf("u = %#v \n", u11)
	fmt.Printf("u = %#v \n", u2)

	fmt.Println(u1 == u11)
	//fmt.Println(u1 == u2) // cannot compare u1 == u2 (mismatched types User1 and User2) 无效运算: u1 == u2(类型 User1 和 User2 不匹配)

	fmt.Println(reflect.DeepEqual(u1, u11))
	fmt.Println(reflect.DeepEqual(u1, u2))

}
