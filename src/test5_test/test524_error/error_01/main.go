package main

import "fmt"

/*
Go 使用 error 接口表示错误，这是 Go 处理异常的标准机制。
type error interface {
    Error() string
}
只要一个类型实现了 Error() string 方法，就实现了 error 接口。
 */


type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func main() {
	var err error = MyError{"something went wrong"}
	fmt.Println(err.Error()) // 输出: something went wrong
}


