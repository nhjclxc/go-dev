package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// "math/rand" 的使用

func Test1(t *testing.T) {

	rand.Seed(time.Now().Unix())

	x := rand.Int()
	fmt.Println("x = ", x)

	xn := rand.Intn(10)
	fmt.Println("xn = ", xn)
	xn2 := xn + 100
	fmt.Println("xn2 = ", xn2)

}

// 官方建议显式创建一个 rand.Rand 对象，绑定自己的 Source，这样可以独立使用，避免全局依赖。
func Test2(t *testing.T) {
	// 创建一个新的随机源
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// 使用 r 来生成随机数
	fmt.Println(r.Intn(100)) // [0,100)
	fmt.Println(r.Float64()) // [0.0,1.0)
	fmt.Println(r.Perm(5))   // 随机排列

	//rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)

}

func Test3(t *testing.T) {
	for i := 48; i <= 57; i++ {
		fmt.Printf("rune = %c, i = %d \n", rune(i), i)
	}
	fmt.Println()

	for i := 65; i <= 90; i++ {
		fmt.Printf("rune = %c, i = %d \n", rune(i), i)
	}
	fmt.Println()

	for i := 97; i <= 122; i++ {
		fmt.Printf("rune = %c, i = %d \n", rune(i), i)
	}
}

func Test5(t *testing.T) {
	fmt.Println(genRandPassword(5))
	fmt.Println(genRandPassword(5))
	fmt.Println(genRandPassword(5))
	fmt.Println(genRandPassword(15))
	fmt.Println(genRandPassword(25))
}
