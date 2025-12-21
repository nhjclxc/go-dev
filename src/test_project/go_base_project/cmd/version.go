package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go_base_project/pkg/version"
)

// versionCmd 版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示 go_base_project 的版本信息，包括版本号、构建时间、Git 提交哈希等`,
	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := version.GetVersion()
		fmt.Println(versionInfo.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
