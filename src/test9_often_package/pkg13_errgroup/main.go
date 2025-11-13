package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

// import "golang.org/x/sync/errgroup"
// go get golang.org/x/sync/errgroup
// 管理并发 goroutine 并收集错误的工具，
func main() {

	// WaitGroup 和 errgroup 的对比使用

	usewaitgroup()

	useergroup()

}

func usewaitgroup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 使用sync.WaitGroup有两个点，第一：在协程外部增加计数wg.Add(1)，第二：在协程退出前减少计数wg.Done()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("grouting1 ctx被关闭！！！")
				// 协程执行完毕，wg计数减1
				wg.Done()
				return
			default:
				if time.Now().Second()%2 == 0 {
					fmt.Println("第一个协程")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("grouting2 ctx被关闭！！！")
				// 协程执行完毕，wg计数减1
				wg.Done()
				return
			default:
				if time.Now().Second()%2 != 0 {
					fmt.Println("第2个协程")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	// 测试关闭 goroutine
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
		fmt.Println("协程关闭！！")
	}()
	wg.Wait()
	fmt.Println("所有 goroutine 已退出")

}

func useergroup() {
	ctx, cancel := context.WithCancel(context.Background())
	group, egCtx := errgroup.WithContext(ctx)
	//var group errgroup.Group

	group.Go(func() error {
		// 开启一个grouting
		for {
			select {
			case <-egCtx.Done():
				fmt.Println("grouting1 ctx被关闭！！！")
				return fmt.Errorf("ctx被关闭")
			default:
				if time.Now().Second()%2 == 0 {
					fmt.Println("第一个协程")
					time.Sleep(1 * time.Second)
				}
			}
		}
	})
	group.Go(func() error {
		// 开启er个grouting
		for {
			select {
			case <-egCtx.Done():
				fmt.Println("ctx被关闭！！！")
				return fmt.Errorf("grouting2 ctx被关闭")
			default:
				if time.Now().Second()%2 != 0 {
					fmt.Println("第二个协程")
					time.Sleep(1 * time.Second)

				}
			}
		}
	})

	fmt.Println("grouting running ... ")

	// 测试关闭协程
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
		fmt.Println("协程关闭！！")
	}()

	// 主线程等待
	if err := group.Wait(); err != nil {
		fmt.Println("errgroup err ", err)
		return
	}
}
