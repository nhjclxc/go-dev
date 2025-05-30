package main

import (
	"fmt"
	"time"
)

// MyTicker 结构体，类似 time.Ticker
type MyTicker struct {
	C      chan time.Time // 用于传出周期时间点
	stopCh chan struct{}  // 用于停止 ticker
}

// NewMyTicker 创建一个自定义 ticker
func NewMyTicker(d time.Duration) *MyTicker {
	t := &MyTicker{
		C:      make(chan time.Time),
		stopCh: make(chan struct{}),
	}

	// 启动后台 goroutine
	go func() {
		for {
			select {
			case <-time.After(d):
				t.C <- time.Now()
			case <-t.stopCh:
				close(t.C)
				return
			}
		}
	}()

	return t
}

// Stop 方法用于停止 ticker
func (t *MyTicker) Stop() {
	close(t.stopCh)
}


func main() {
	ticker := NewMyTicker(2 * time.Second)
	defer ticker.Stop()

	//for i := 0; i < 5; i++ {
	//	t := <-ticker.C
	//	fmt.Println("Tick at", t)
	//}

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}


	fmt.Println("Ticker stopped.")
}

