package model

import (
	"fmt"
	"testing"
)

// 使用go里面内置的test模块，对每一个单元进行单元测试

// go专门的test文件内，不得出现main方法

func TestNewInstance(t *testing.T) {
	c := Customer{
		Id:    1,
		Name:  "aaa",
		Sex:   "zz",
		Age:   0,
		Phone: "aaaa",
		Email: "aaaaaaa",
		Money: 666,
	}
	fmt.Println(c.String())

	cus := NewCustomer(111, "张安", "buzd ", 16, "123", "aaa", 123)
	fmt.Println(cus)
}
