package main

import (
	"fmt"
	"testing"
)

func TestStringSet(t *testing.T) {

	set := NewStringSet()
	set.Add("a")
	set.Add("b")
	set.Add("a") // 不重复

	fmt.Println(set.Contains("a")) // true
	fmt.Println(set.Contains("c")) // false

	set.Remove("a")
	fmt.Println(set.Contains("a")) // false

	fmt.Println(set.Values()) // [b]
	fmt.Println(set.Len())    // 1
}