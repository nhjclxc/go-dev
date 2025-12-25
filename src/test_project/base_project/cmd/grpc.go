package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"base_project/internal/server"
)

// grpcCmd gRPC 服务命令
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "启动 gRPC 服务",
	Long:  `启动 gRPC 服务，接收边缘节点心跳上报和 PCDN 节点状态上报`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化应用
		application, err := initApp()
		if err != nil {
			return err
		}
		defer application.Close()

		slog.Info("启动 gRPC 服务...")

		// 创建并启动 gRPC 服务器
		grpcServer := server.NewGRPCServer(application)
		return grpcServer.Start()
	},
}
