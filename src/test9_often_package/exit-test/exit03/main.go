package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 使用信号进行优雅关闭
// sigCh := make(chan os.Signal, 1) signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
// worker 模拟任务执行
func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d 启动\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d 接收到退出信号，停止工作\n", id)
			return
		default:
			fmt.Printf("Worker %d 正在工作...\n", id)
			time.Sleep(1 * time.Second) // 模拟任务执行
		}
	}
}

func main() {
	// 1️⃣ 创建可取消的 context
	ctx, cancel := context.WithCancel(context.Background())

	// 2️⃣ 捕获系统信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 3️⃣ 启动 worker
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, &wg, i)
	}

	// 4️⃣ 等待系统信号
	go func() {
		sig := <-sigCh
		fmt.Printf("收到退出信号: %v\n", sig)
		cancel() // 通知所有 worker 退出
	}()

	// 5️⃣ 阻塞等待所有 worker 完成
	wg.Wait()
	fmt.Println("所有 worker 已优雅退出，程序结束")

	/*

		// 等待中断信号
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		// 优雅关闭
		manager.StopAll()
	*/
}
