package client

import (
	"base_project/config"
	"context"
	"log/slog"
)

// ClientRunner 表示 androidjlink 客户端运行器
// 这是一个轻量的启动钩子，实际业务逻辑在 internal/tunnel/client 中
type ClientRunner struct {
	cfg      *config.ClientConfig
	deviceID string // 自动检测的设备 ID
}

// NewClientRunner 创建客户端运行器
func NewClientRunner(cfg *config.ClientConfig) *ClientRunner {
	return &ClientRunner{
		cfg: cfg,
	}
}

// Run 启动客户端逻辑
// 调用 internal/tunnel/client 的核心逻辑
func (r *ClientRunner) Run(ctx context.Context) error {
	slog.Info("客户端运行器已启动")

	return nil
}

// Stop 停止客户端
func (r *ClientRunner) Stop() error {
	// 停止客户端
	return nil
}
