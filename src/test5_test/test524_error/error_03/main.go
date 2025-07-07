package main

import (
	"errors"
	"fmt"
)

// ✅ 四、错误链工具函数
//Go 标准库 errors 包中提供了几个方法来操作链式错误。
func main() {

	// 1. errors.Unwrap(err error) error
	// 获取被包裹的底层错误。
	err0 := errors.New("inner 0 error \n")
	err1 := fmt.Errorf("inner 1 error: %w \n", err0)
	err := fmt.Errorf("outer error: %w \n", err1)
	fmt.Println(errors.Unwrap(err)) // 输出: inner error


	// 2. errors.Is(err, target error) bool
	//判断某个错误链中是否存在特定错误（根据 == 判断）。
	fmt.Println(errors.Is(err, err0))
	errTemp := errors.New("inner errTemp error \n")
	fmt.Println(errors.Is(err, errTemp))


	// 3. errors.As(err, target interface{}) bool
	//判断并提取链中具体某种类型的错误。
	err11 := doSomething()
	var myErr *MyError
	if errors.As(err11, &myErr) {
		fmt.Printf("捕获到自定义错误: Code=%d, Msg=%s\n", myErr.Code, myErr.Msg)
	}





}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return e.Msg
}

func doSomething() error {
	return fmt.Errorf("wrap: %w", &MyError{Code: 500, Msg: "Internal Error"})
}