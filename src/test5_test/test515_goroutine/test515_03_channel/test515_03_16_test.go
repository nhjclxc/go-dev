package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 使用chan来传递数据，
// 实现功能实现三个方法并且使用chan传递数据来实现数据的相加
//实现三个方法：1、随机生成一个slice。2、将1中生成的随机数数组传入3中进行计算返回。3、计算2中传入数组的和

const chanLength = 5

var random *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func genSlice(min, max int, inCh chan<- int) {
	defer close(inCh) // 关闭

	for i := 0; i < chanLength; i++ {
		num := random.Intn(max-min+1) + min
		fmt.Printf("%d, ", num)
		inCh <- num
	}
	fmt.Println()
}

func sum(out <-chan int) int {
	res := 0
	// 读取通道数据的时候，如果使用for，那么通道必须现实关闭
	for val := range out {
		res += val
	}
	return res
}

func sum2(out <-chan int) int {
	res := 0
	flag := true
	for flag {
		select {
		case val, ok := <-out:
			if ok {
				res += val
			} else {
				flag = false
			}
		}
	}
	return res
}

func clu(min, max int) int {
	ch := make(chan int, chanLength)

	go genSlice(min, max, ch)
	return sum2(ch)
}

func Test161(t *testing.T) {

	sum := clu(1, 10)

	fmt.Println("sum = ", sum)

}
