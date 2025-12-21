//go:build server
// +build server

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go_base_project/config"
	"go_base_project/pkg/logger"
	tunnelserver "go_base_project/server/tunnel/server"
)

// tunnelServerCmd SSH 隧道服务端命令
var tunnelServerCmd = &cobra.Command{
	Use:   "admin",
	Short: "启动服务端",
	Long:  `启动服务端，负责为 ... 能力。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载服务端专用配置
		cfg, err := config.LoadServerConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 初始化日志
		if err := logger.Init(&cfg.Log); err != nil {
			return fmt.Errorf("初始化日志失败: %w", err)
		}

		logger.Info("启动 SSH 隧道服务端...")

		runner := tunnelserver.NewRunner(cfg)
		return runner.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(tunnelServerCmd)
}
