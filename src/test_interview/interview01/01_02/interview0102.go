package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint = 1
	var b uint = 2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b - a)
	fmt.Println(a - b) // 18446744073709551615
	var max uint64 = math.MaxUint64
	fmt.Println(max)
}


func init(){
	fmt.Println("init 3")
}

func init(){
	fmt.Println("init 1")
}


func init(){
	fmt.Println("init 2")
}
