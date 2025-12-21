package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go_base_project/internal/server/server"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go_base_project/config"
	"go_base_project/pkg/logger"
)

// serverCmd 服务端命令
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务端",
	Long:  `启动服务端，负责为 ... 能力。`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// 1️⃣ 创建可取消的 context
		ctx, cancel := context.WithCancel(cmd.Context())

		// 2️⃣ 捕获系统信号
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		// 加载服务端专用配置
		cfg, err := config.LoadServerConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 初始化日志
		if err := logger.Init(cfg.Log); err != nil {
			return fmt.Errorf("初始化日志失败: %w", err)
		}

		logger.Info("启动 服务端...")

		server, err := server.NewServer(cfg)
		if err != nil {
			logger.Info("启动服务失败：", "err", err)
		}

		// 3️⃣ 启动 worker
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := server.Start(ctx); err != nil {
				logger.Error("启动服务失败：", "err", err)
				return
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("服务2")
		}()

		// 4️⃣ 等待系统信号
		go func() {
			sig := <-sigCh
			fmt.Printf("收到退出信号: %v\n", sig)
			cancel() // 通知所有 worker 退出

			// 关闭所有服务
			if err := server.Stop(ctx); err != nil {
				logger.Error("关闭 Server 服务失败", "error", err)
			}
		}()

		// 5️⃣ 阻塞等待所有 worker 完成
		wg.Wait()
		fmt.Println("所有 worker 已优雅退出，程序结束")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

// go run .\main.go server -c ./config/server.yaml
