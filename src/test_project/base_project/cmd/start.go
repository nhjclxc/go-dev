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

	"base_project/internal/server"
)

// startCmd 启动所有服务命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动所有服务",
	Long:  `启动所有服务，包括 HTTP API、gRPC 调度服务和心跳检测定时任务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化应用
		application, err := initApp()
		if err != nil {
			return err
		}

		slog.Info("启动所有服务...")

		// 创建服务器实例
		httpServer := server.NewHTTPServer(application)
		grpcServer := server.NewGRPCServer(application)
		cronServer := server.NewCronServer(application)

		// 使用 WaitGroup 等待所有服务启动
		var wg sync.WaitGroup

		// 启动 HTTP 服务
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpServer.Start(); err != nil {
				slog.Error("HTTP 服务启动失败", "error", err)
			}
		}()

		// 启动 gRPC 服务
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := grpcServer.Start(); err != nil {
				slog.Error("gRPC 服务启动失败", "error", err)
			}
		}()

		// 启动定时任务服务
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := cronServer.Start(); err != nil {
				slog.Error("定时任务服务启动失败", "error", err)
			}
		}()

		// 监听退出信号
		//quit := make(chan os.Signal, 1)
		//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		//<-quit

		// 1️⃣ 创建可取消的 context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// 2️⃣ 捕获系统信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		sig := <-quit
		cancel()
		fmt.Printf("收到退出信号: %v\n", sig)

		slog.Info("正在优雅关闭服务...")

		// 关闭所有服务
		if err := httpServer.Stop(ctx); err != nil {
			slog.Error("关闭 HTTP 服务失败", "error", err)
		}

		if err := grpcServer.Stop(ctx); err != nil {
			slog.Error("关闭 gRPC 服务失败", "error", err)
		}

		if err := cronServer.Stop(ctx); err != nil {
			slog.Error("关闭定时任务服务失败", "error", err)
		}

		slog.Info("所有服务已关闭")

		// 关闭应用（包括数据库、Redis、Telemetry、日志等）
		if err := application.Close(); err != nil {
			slog.Error("关闭应用资源失败", "error", err)
		}
		slog.Info("所有资源已关闭")

		// 5️⃣ 阻塞等待所有 worker 完成
		wg.Wait()
		fmt.Println("所有 worker 已优雅退出，程序结束")
		slog.Info("程序停止！！！")

		return nil
	},
}

// go run main.go start -c ./config/admin.yaml
