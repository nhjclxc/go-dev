package cmd

import (
	"base_project/pkg/client"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"

	"base_project/config"
	"base_project/pkg/logger"
	"base_project/pkg/version"
)

var (
	showVersion bool
)

// clientCmd 客户端命令
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "启动客户端",
	Long:  `启动客户端，负责为 ... 能力。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 如果指定了 --version 标志，显示版本信息后退出
		if showVersion {
			fmt.Println(version.GetVersion().String())
			return nil
		}

		// 1️⃣ 创建可取消的 context
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		// 2️⃣ 捕获系统信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// 加载客户端专用配置
		cfg, err := config.LoadClientConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 初始化日志
		log, logCloser, err := logger.New(cfg.Log)
		if err != nil {
			return fmt.Errorf("初始化日志失败: %w", err)
		}
		// 设置为全局默认日志（便于 slog 直接使用）
		slog.SetDefault(log)

		defer func() {
			// 关闭日志写入器
			if logCloser != nil {
				if err := logCloser.Close(); err != nil {
				}
			}
		}()

		slog.Info("启动客户端...")
		slog.Info("版本信息: " + version.Version)

		// 3️⃣ 启动 worker
		var wg sync.WaitGroup
		wg.Add(1)

		clientRunner := client.NewClientRunner(cfg)
		go func() {
			defer wg.Done()

			err := clientRunner.Run(ctx)
			if err != nil {
				return
			}
		}()

		// 4️⃣ 等待系统信号
		//go func() {
		sig := <-quit
		fmt.Printf("收到退出信号: %v\n", sig)
		cancel() // 通知所有 worker 退出

		// 关闭服务
		if err := clientRunner.Stop(); err != nil {
			slog.Error("关闭 Server 服务失败", "error", err)
		}
		//}()

		// 5️⃣ 阻塞等待所有 worker 完成
		wg.Wait()
		fmt.Println("所有 worker 已优雅退出，程序结束")

		return nil
	},
}

func init() {
	clientCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "显示版本信息")
}

// go run main.go client -c ./config/client.yaml
