package main

import (
	"fmt"
	"testing"
	"time"
)

// main01.go 的压力测试

func Benchmark_Division(b *testing.B) {
	fmt.Println("b.N = ", b.N)
	// 当前时间
	start := time.Now()
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
	// 当前时间
	end := time.Now()
	// 计算时间差
	duration := end.Sub(start)

	// 输出时间差
	fmt.Println("时间差为:", duration)
	fmt.Printf("时间差为.Seconds(): = %v \n\n", duration.Seconds())

}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
