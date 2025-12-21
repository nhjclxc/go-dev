package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go_base_project/config"
	"go_base_project/internal/controller"
	"go_base_project/internal/server/server/middleware"
	"go_base_project/pkg/database"
	"go_base_project/pkg/logger"
	redisPkg "go_base_project/pkg/redis"
	"gorm.io/gorm"
	"log/slog"
	"net/http"

	"go_base_project/internal/repository"
	"go_base_project/internal/service"
)

// Server HTTP 服务器
type Server struct {
	server   *http.Server
	cfg      *config.ServerConfig
	ctx      context.Context
	db       *gorm.DB
	redisCli *redis.Client
}

// NewHTTPServer 创建 HTTP 服务器实例
func NewServer(cfg *config.ServerConfig) (*Server, error) {
	s := Server{
		cfg: cfg,
	}

	// 注册数据库检查器
	if cfg.Database != nil {
		db0, err := database.NewMySQL(cfg.Database)
		if err != nil {
			logger.Info("Server.NewServer.NewMySQL err:", err)
			return nil, err
		}
		s.db = db0
	}

	// 注册 Redis 检查器
	if cfg.RedisConfig != nil {
		redis0, err := redisPkg.NewRedis(cfg.RedisConfig)
		if err != nil {
			logger.Info("Server.NewServer.NewRedis err:", err)
			return nil, err
		}
		s.redisCli = redis0
	}

	return &s, nil
}

// Start 启动 HTTP 服务器
func (s *Server) Start(ctx context.Context) error {
	s.ctx = ctx
	cfg := s.cfg

	// 设置 Gin 模式
	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建 Gin 引擎
	r := gin.New()

	// 注册中间件
	r.Use(middleware.CORS()) // CORS 中间件（允许跨域）
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

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
func (s *Server) Stop(ctx context.Context) error {
	slog.Info("正在关闭 HTTP 服务器...")
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// registerRoutes 注册路由
func (s *Server) registerRoutes(r *gin.Engine) {
	db := s.db

	// API 路由组
	api := r.Group("/api/v1", middleware.AdminAuthMiddleware(s.cfg.Login.JWT))
	openApi := r.Group("/openapi/v1", middleware.OpenApiAuthMiddleware()) // 不鉴权的

	// 初始化 User
	tabUserRepo := repository.NewTabUserRepository(db)
	tabUserService := service.NewTabUserService(db, tabUserRepo)
	s.registerabUserRoutes(api, openApi, tabUserService)

	// ...
}

func (s *Server) registerabUserRoutes(api *gin.RouterGroup, openApi *gin.RouterGroup, userService *service.TabUserService) {
	// 初始化 Controller
	tabUserController := controller.NewTabUserController(userService)

	// 用户路由
	userApi := api.Group("/user")
	{
		userApi.POST("", tabUserController.InsertTabUser)
		userApi.GET("/:id", tabUserController.GetTabUserById)
		userApi.GET("", tabUserController.GetTabUserList)
		userApi.PUT("/:id", tabUserController.UpdateTabUser)
		userApi.DELETE("/:id", tabUserController.DeleteTabUser)
	}

	// 用户路由
	userOpenApi := openApi.Group("/openapi/user")
	_ = userOpenApi
}
