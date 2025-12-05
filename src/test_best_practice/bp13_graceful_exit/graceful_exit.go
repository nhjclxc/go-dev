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

// 模拟程序优雅退出
// graceful_exit.go

type Server struct {
	count int
	ctx   context.Context
	wg    *sync.WaitGroup
}

func NewServer(ctx context.Context, count int) *Server {
	return &Server{
		count: count,
		ctx:   ctx,
		wg:    &sync.WaitGroup{},
	}
}
func (s *Server) Start() error {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				fmt.Println("程序被关闭，退出服务！！！")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("count = ", s.count)
				s.count++
			}
		}
	}()
	return nil
}
func (s *Server) Shutdown() error {
	fmt.Println("等待后台任务退出...")
	s.wg.Wait()
	fmt.Println("Server Closed")
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	s := NewServer(ctx, 123)
	s.Start()

	// 优雅退出  等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down servers...")

	// 通知 goroutine 收到退出信号
	cancel()

	// 通知服务关闭
	s.Shutdown()

	fmt.Println("Server stopped")
}

/*
lxc20250729@lxc20250729deMacBook-Pro bp13_graceful_exit % go run graceful_exit.go
count =  123
count =  124
count =  125
^CShutting down servers...
等待后台任务退出...
count =  126
程序被关闭，退出服务！！！
Server Closed
Server stopped
*/
