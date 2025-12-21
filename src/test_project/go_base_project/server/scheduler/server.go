package scheduler

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go_base_project/config"
	"go_base_project/pkg/logger"
	"go_base_project/server/scheduler/tasks"
)

// Server 定时任务服务器
type Server struct {
	cron *cron.Cron
	cfg  *config.AdminConfig
}

// NewServer 创建定时任务服务器实例
// 参数:
//   - cfg: 管理服务配置
//
// 返回值:
//   - *Server: 服务器实例
func NewServer(cfg *config.AdminConfig) *Server {
	// 创建 cron 实例
	c := cron.New(
		cron.WithSeconds(), // 支持秒级定时
		cron.WithChain( // 添加中间件
			cron.Recover(cron.DefaultLogger), // 恢复 panic
		),
	)

	return &Server{
		cron: c,
		cfg:  cfg,
	}
}

// Start 启动定时任务服务器
func (s *Server) Start() error {
	// 注册所有任务
	for _, taskCfg := range s.cfg.Cron.Tasks {
		if !taskCfg.Enabled {
			logger.Info("跳过未启用的任务", "task", taskCfg.Name)
			continue
		}

		task := tasks.GetTask(taskCfg.Name)
		if task == nil {
			logger.Warn("未找到任务", "task", taskCfg.Name)
			continue
		}

		// 添加任务
		_, err := s.cron.AddFunc(taskCfg.Spec, func() {
			ctx := context.Background()
			logger.Info("开始执行定时任务", "task", task.Name())
			if err := task.Execute(ctx); err != nil {
				logger.Error("定时任务执行失败", "task", task.Name(), "error", err)
			} else {
				logger.Info("定时任务执行成功", "task", task.Name())
			}
		})

		if err != nil {
			return fmt.Errorf("添加定时任务失败: %s, %w", taskCfg.Name, err)
		}

		logger.Info("注册定时任务", "task", taskCfg.Name, "spec", taskCfg.Spec)
	}

	logger.Info("定时任务服务器启动")
	s.cron.Start()

	// 阻塞等待
	select {}
}

// Stop 停止定时任务服务器
func (s *Server) Stop(ctx context.Context) error {
	logger.Info("正在关闭定时任务服务器...")
	if s.cron != nil {
		stopCtx := s.cron.Stop()
		<-stopCtx.Done()
	}
	return nil
}

// RunTask 立即执行指定任务（用于测试）
func (s *Server) RunTask(ctx context.Context, taskName string, args []string) error {
	task := tasks.GetTask(taskName)
	if task == nil {
		return fmt.Errorf("未找到任务: %s", taskName)
	}

	logger.Info("立即执行任务", "task", taskName, "args", args)
	if err := task.Execute(ctx); err != nil {
		logger.Error("任务执行失败", "task", taskName, "error", err)
		return err
	}

	logger.Info("任务执行成功", "task", taskName)
	return nil
}
