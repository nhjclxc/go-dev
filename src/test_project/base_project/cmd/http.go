package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"base_project/internal/server"
)

// httpCmd HTTP 服务命令
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "启动 HTTP 服务",
	Long:  `启动 HTTP 服务，提供 RESTful API 接口`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化应用
		application, err := initApp()
		if err != nil {
			return err
		}
		defer application.Close()

		slog.Info("启动 HTTP 服务...")

		// 创建并启动 HTTP 服务器
		httpServer := server.NewHTTPServer(application)
		return httpServer.Start()
	},
}
