package test06

import (
	"fmt"
	"testing"
)

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func Test111(t *testing.T) {

	// 用闭包实现计数器

	c := counter()
	println(c())
	println(c())
	println(c())
	println(c())
	println(c())

	fmt.Println("===========================")
	c2 := counter()
	println(c2())
	println(c2())
	println(c2())

}
