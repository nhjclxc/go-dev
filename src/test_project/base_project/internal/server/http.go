package server

import (
	"base_project/internal/middleware"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"base_project/pkg/app"
	"base_project/pkg/health"
)

// HTTPServer HTTP 服务器
type HTTPServer struct {
	server *http.Server
	app    *app.App
	health *health.Health
}

// NewHTTPServer 创建 HTTP 服务器实例
func NewHTTPServer(application *app.App) *HTTPServer {
	// 初始化健康检查器
	h := health.New()

	// 注册数据库检查器
	if application.DB != nil {
		h.Register(health.NewDatabaseChecker(application.DB))
	}

	// 注册 Redis 检查器
	if application.Redis != nil {
		h.Register(health.NewRedisChecker(application.Redis))
	}

	return &HTTPServer{
		app:    application,
		health: h,
	}
}

// Start 启动 HTTP 服务器
func (s *HTTPServer) Start() error {
	cfg := s.app.Config

	// 设置 Gin 模式
	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建 Gin 引擎
	r := gin.New()

	// 注册中间件
	r.Use(middleware.Recovery())

	// 添加 OpenTelemetry 中间件（在 Logger 之前，以便 Logger 能获取 trace_id）
	if s.app.Telemetry != nil && s.app.Telemetry.IsEnabled() {
		r.Use(otelgin.Middleware(cfg.App.Name))
	}

	r.Use(middleware.Logger())

	// 启用 pprof
	if cfg.App.Debug {
		pprof.Register(r)
	}

	// 注册健康检查路由
	s.registerHealthRoutes(r)

	// 注册业务路由
	s.registerRoutes(r)

	// 创建 HTTP 服务器
	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.HTTP.Port),
		Handler:      r,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	slog.Info("HTTP 服务器启动", "port", cfg.HTTP.Port)

	// 启动服务器
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP 服务器启动失败: %w", err)
	}

	return nil
}

// Stop 停止 HTTP 服务器
func (s *HTTPServer) Stop(ctx context.Context) error {
	slog.Info("正在关闭 HTTP 服务器...")
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// registerHealthRoutes 注册健康检查路由
func (s *HTTPServer) registerHealthRoutes(r *gin.Engine) {
	// 简单存活检查（K8s liveness probe）
	// 只检查服务是否运行，不检查依赖
	r.GET("/health/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// 就绪检查（K8s readiness probe）
	// 检查所有依赖是否就绪
	r.GET("/health/ready", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		result := s.health.Check(ctx)

		status := http.StatusOK
		if result.Status == health.StatusDown {
			status = http.StatusServiceUnavailable
		}

		c.JSON(status, result)
	})

	// 完整健康检查（包含详细信息）
	r.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		result := s.health.Check(ctx)

		status := http.StatusOK
		if result.Status == health.StatusDown {
			status = http.StatusServiceUnavailable
		}

		c.JSON(status, result)
	})
}

// registerRoutes 注册业务路由
func (s *HTTPServer) registerRoutes(r *gin.Engine) {
	db := s.app.DB
	_ = db
}
