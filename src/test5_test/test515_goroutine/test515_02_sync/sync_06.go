package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map
func main06() {

	// Go语言中内置的map不是并发安全的。
	// 要想使用线程安全的映射，则必须使用 sync.Map

	// map 不安全的使用示例
	unsafeMapUse()

	safeMapUse()
}

func safeMapUse() {

	// map 加锁安全使用示例

	// 使用的是go内置的sync.Map包

	var syncMap = sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap.Store(key, n)         // 存
			value, _ := syncMap.Load(key) // 取
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func unsafeMapUse() {
	// map 不安全使用示例

	var m = make(map[string]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m[key] = n
			fmt.Printf("k=:%v,v:=%v\n", key, m[key])
			wg.Done()
		}(i)
	}
	wg.Wait()
	/*
		fatal error: concurrent map writes

			goroutine 20 [running]:
			internal/runtime/maps.fatal({0xe647d4?, 0x0?})
			        D:/develop/go/src/runtime/panic.go:1053 +0x18
			main.unsafeMapUse.func1(0x1)
			        D:/code/go/go-dev/src/test5_test/test515_goroutine/test515_02_sync/sync_06.go:26 +0x58
			created by main.unsafeMapUse in goroutine 1

			        D:/code/go/go-dev/src/test5_test/test515_goroutine/test515_02_sync/sync_06.go:24 +0x45

	*/
}
