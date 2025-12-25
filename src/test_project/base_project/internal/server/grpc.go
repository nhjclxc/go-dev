package server

import (
	grpcmw "base_project/internal/middleware"
	"context"
	"fmt"
	"log/slog"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"base_project/pkg/app"
)

// GRPCServer gRPC 服务器
type GRPCServer struct {
	server *grpc.Server
	app    *app.App
}

// NewGRPCServer 创建 gRPC 服务器实例
func NewGRPCServer(application *app.App) *GRPCServer {
	return &GRPCServer{
		app: application,
	}
}

// Start 启动 gRPC 服务器
func (s *GRPCServer) Start() error {
	cfg := s.app.Config

	lis, err := net.Listen(cfg.GRPC.Network, fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		return fmt.Errorf("gRPC 监听失败: %w", err)
	}

	// 创建 gRPC 服务器选项
	opts := []grpc.ServerOption{}

	// 添加鉴权拦截器
	authInterceptor := grpcmw.NewAuthInterceptor(&cfg.GRPC.Auth)
	opts = append(opts,
		grpc.ChainUnaryInterceptor(authInterceptor.Unary()),
		grpc.ChainStreamInterceptor(authInterceptor.Stream()),
	)

	if cfg.GRPC.Auth.Enabled {
		slog.Info("gRPC 鉴权已启用")
	}

	// 添加 OpenTelemetry 拦截器
	if s.app.Telemetry != nil && s.app.Telemetry.IsEnabled() {
		opts = append(opts,
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
		)
	}

	// 创建 gRPC 服务器
	s.server = grpc.NewServer(opts...)

	// 注册服务
	s.registerServices()

	slog.Info("gRPC 服务器启动", "port", cfg.GRPC.Port)

	// 启动服务器
	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("gRPC 服务器启动失败: %w", err)
	}

	return nil
}

// Stop 停止 gRPC 服务器
func (s *GRPCServer) Stop(ctx context.Context) error {
	slog.Info("正在关闭 gRPC 服务器...")
	if s.server != nil {
		// 使用 goroutine 尝试优雅关闭
		stopped := make(chan struct{})
		go func() {
			s.server.GracefulStop()
			close(stopped)
		}()

		// 等待优雅关闭或超时
		select {
		case <-stopped:
			slog.Info("gRPC 服务器已优雅关闭")
		case <-ctx.Done():
			slog.Warn("gRPC 优雅关闭超时，强制关闭")
			s.server.Stop() // 强制关闭
		}
	}
	return nil
}

// registerServices 注册 gRPC 服务
func (s *GRPCServer) registerServices() {
	// ====================== ... ==========================

}
