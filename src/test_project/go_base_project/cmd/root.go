package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Use:   "go_base_project",
	Short: "go_base_project",
	Long:  `go_base_project 提供 ... 等功能`,
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "执行命令失败: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	// 持久化标志（所有子命令都可以使用）
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "配置文件路径")

	// 子命令在各自的文件中通过 init() 函数注册到 rootCmd
}
