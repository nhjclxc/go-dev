package cmd

import (
	"fmt"
	"go_base_project/internal/client"

	"github.com/spf13/cobra"

	"go_base_project/config"
	"go_base_project/pkg/logger"
	"go_base_project/pkg/version"
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

		// 加载客户端专用配置
		cfg, err := config.LoadClientConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}

		// 初始化日志
		if err := logger.Init(cfg.Log); err != nil {
			return fmt.Errorf("初始化日志失败: %w", err)
		}

		logger.Info("启动客户端...")
		logger.Info("版本信息: " + version.Version)

		runner := client.NewRunner(cfg)
		return runner.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "显示版本信息")
}
