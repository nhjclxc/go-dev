package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// startCmd 启动所有服务命令
var start2Cmd = &cobra.Command{
	Use:   "start2",
	Short: "启动所有服务",
	Long:  `启动所有服务，包括 HTTP API、gRPC 调度服务和心跳检测定时任务`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// 1️⃣ 创建可取消的 context

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		_ = ctx

		// 2️⃣ 捕获系统信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// 3️⃣ 启动所有服务
		var wg sync.WaitGroup // 使用 WaitGroup 等待所有服务启动

		// 启动 HTTP 服务
		wg.Add(1)
		s1, err := NewServer(ctx, "协程1")
		_ = err
		go func() {
			defer wg.Done()
			if err := s1.Start(); err != nil {
				return
			}
		}()

		// 启动 HTTP 服务
		wg.Add(1)
		s2, err := NewServer(ctx, "协程22")
		go func() {
			defer wg.Done()
			if err := s2.Start(); err != nil {
				return
			}
		}()
		// 启动 HTTP 服务
		wg.Add(1)
		s3, err := NewServer(ctx, "协程333")
		go func() {
			defer wg.Done()
			if err := s3.Start(); err != nil {
				return
			}
		}()

		// 3️⃣ 捕获系统信号
		sig := <-quit
		cancel()
		fmt.Printf("收到退出信号: %v\n", sig)

		slog.Info("正在优雅关闭服务...")

		stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// 关闭所有服务
		if err := s1.Stop(stopCtx); err != nil {
			slog.Error("关闭 s1 服务失败", "error", err)
		}
		if err := s2.Stop(stopCtx); err != nil {
			slog.Error("关闭 s2 服务失败", "error", err)
		}
		if err := s3.Stop(stopCtx); err != nil {
			slog.Error("关闭 s3 服务失败", "error", err)
		}

		// 5️⃣ 阻塞等待所有 worker 完成
		wg.Wait()
		fmt.Println("所有 worker 已优雅退出，程序结束")
		slog.Info("程序停止！！！")

		return nil
	},
}

// go run main.go start2 -c ./config/admin.yaml

type Server struct {
	serverName string
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	startOnce  sync.Once
	stopOnce   sync.Once
}

func NewServer(ctx0 context.Context, serverName string) (*Server, error) {
	slog.Info("New: %v", serverName)
	// 这个ctx是用于在Server内部进行传递通知的
	ctx, cancel := context.WithCancel(ctx0)
	return &Server{
		ctx:        ctx,
		cancel:     cancel,
		serverName: serverName,
	}, nil
}

func (s *Server) Stop(ctx context.Context) error {
	// 入参的ctx是用于控制 stop方法的

	var err error

	s.stopOnce.Do(func() {
		s.cancel()

		done := make(chan struct{})
		go func() {
			s.wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			slog.Info("server shutdown complete", "name", s.serverName)
		case <-ctx.Done():
			err = ctx.Err()
		}
	})

	return err
}

func (s *Server) Start() error {
	s.startOnce.Do(func() {
		s.wg.Add(1)

		go func() {
			defer s.wg.Done()

			for {
				select {
				case <-s.ctx.Done():
					slog.Info("server stopped", "name", s.serverName)
					return
				default:
					s.run()
				}
			}
		}()
	})

	return nil
}

func (s *Server) run() {
	// 这里一般在实现一个内部的run方法，业务就是在run里面跑
	fmt.Printf("%s running ...\n", s.serverName)
	time.Sleep(500 * time.Millisecond)
}
