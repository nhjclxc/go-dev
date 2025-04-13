package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 判断1-2000中哪些数是素数
func main14() {

	start := time.Now().UnixNano()

	var wg sync.WaitGroup
	wg.Add(3)
	var asyncCount int32 = 0

	go do(&wg, &asyncCount, 1, 2000)
	go do(&wg, &asyncCount, 2000, 4000)
	go do(&wg, &asyncCount, 4000, 6000)

	wg.Wait()
	end := time.Now().UnixNano()

	start2 := time.Now().UnixNano()
	var syncCount int = 0
	do(nil, &syncCount, 0, 6000)
	end2 := time.Now().UnixNano()

	fmt.Println("----------------------------------")
	fmt.Println("async = ", asyncCount, (end - start))
	fmt.Println("sync = ", syncCount, (end2 - start2))

}
func do(wg *sync.WaitGroup, counter interface{}, start, end int) {
	var count int = 0
	for i := start; i < end; i++ {
		if isPrime(i) {
			//fmt.Println(i)
			count++
		}
	}
	if wg != nil {
		wg.Done()
	}

	switch c := counter.(type) {
	case *int32: // 并发用 atomic
		atomic.AddInt32(c, int32(count))
	case *int: // 串行直接累加
		*c += count
	}
}

// 判断 n 是不是素数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	//for i := 2; i < n; i++ {
	for i := 2; i*i <= n; i++ { // 🚀 优化点：只判断到 sqrt(n)
		if n%i == 0 {
			return false
		}
	}
	return true
}
