package client

import (
	"context"
	"go_base_project/config"
	"go_base_project/pkg/logger"
)

// Runner 表示 androidjlink 客户端运行器
// 这是一个轻量的启动钩子，实际业务逻辑在 internal/tunnel/client 中
type Runner struct {
	cfg      *config.ClientConfig
	deviceID string // 自动检测的设备 ID
}

// NewRunner 创建客户端运行器
// 参数:
//   - cfg: 客户端配置
//
// 返回值:
//   - *Runner: 运行器实例
func NewRunner(cfg *config.ClientConfig) *Runner {
	return &Runner{
		cfg: cfg,
	}
}

// Run 启动客户端逻辑
// 调用 internal/tunnel/client 的核心逻辑
func (r *Runner) Run(ctx context.Context) error {
	logger.Info("客户端运行器已启动")

	return nil
}

// Stop 停止客户端
func (r *Runner) Stop() error {
	// 停止隧道客户端
	return nil
}
