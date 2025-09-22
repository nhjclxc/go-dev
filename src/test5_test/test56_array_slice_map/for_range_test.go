package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSlice(t *testing.T) {
	nums := []int{66, 88, 99}
	for i, num := range nums {
		fmt.Printf("i = %d, val = %d \n", i, num)
	}

	fmt.Println()
	for i := range nums {
		fmt.Printf("i = %d \n", i)
	}

	// 随机生成[0,n)范围内的随机数，每个数字只出现一次
	for _, value := range rand.Perm(5) {
		fmt.Println(value)
	}
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"a": 11,
		"b": 22,
		"c": 33,
	}
	for key, val := range m {
		fmt.Printf("key = %s, val = %d \n", key, val)
	}

	fmt.Println()
	for key := range m {
		fmt.Printf("key = %s \n", key)
	}

	val, ok := m["a"]
	fmt.Printf("val = %d, ok = %t \n", val, ok)
	val, ok = m["d"]
	fmt.Printf("val = %d, ok = %t \n", val, ok)
	val = m["a"]
	fmt.Printf("val = %d \n", val)
	val = m["d"]
	fmt.Printf("val = %d \n", val)

}
