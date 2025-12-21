//go:build server
// +build server

package cmd

import (
	"context"
	"fmt"
	"go_base_project/internal/server/scheduler"

	"github.com/spf13/cobra"

	"go_base_project/config"
	"go_base_project/pkg/database"
	"go_base_project/pkg/logger"
)

var (
	taskName string
	taskArgs []string
)

// cronCmd 定时任务服务命令
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "启动定时任务服务",
	Long:  `启动定时任务服务，管理和执行所有定时任务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载管理服务配置（定时任务使用与 admin 相同的配置）
		cfg, err := config.LoadServerConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 初始化日志
		if err := logger.Init(&cfg.Log); err != nil {
			return fmt.Errorf("初始化日志失败: %w", err)
		}

		// 初始化数据库
		if err := database.Init(&cfg.Database); err != nil {
			return fmt.Errorf("初始化数据库失败: %w", err)
		}

		// 如果指定了 task 参数，则立即执行该任务
		if taskName != "" {
			logger.Info("立即执行任务", "task", taskName, "args", taskArgs)
			schedulerServer := scheduler.NewServer(cfg)
			return schedulerServer.RunTask(context.Background(), taskName, taskArgs)
		}

		logger.Info("启动定时任务服务...")

		// 创建并启动定时任务服务器
		schedulerServer := scheduler.NewServer(cfg)
		return schedulerServer.Start()
	},
}

func init() {
	// go run main.go cron --task=port_recycle --config=./config.yaml
	cronCmd.Flags().StringVar(&taskName, "task", "", "立即执行指定的任务（用于测试）")
	cronCmd.Flags().StringArrayVar(&taskArgs, "args", []string{}, "任务参数")
	rootCmd.AddCommand(cronCmd)
}
