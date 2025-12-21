package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"go_base_project/config"
	"go_base_project/internal/controller"
	"go_base_project/internal/repository"
	"go_base_project/internal/service"
	"go_base_project/pkg/database"
	"go_base_project/pkg/logger"
	"go_base_project/server/admin/middleware"
)

// Server 管理服务器
type Server struct {
	server *http.Server
	cfg    *config.AdminConfig
	ctx    context.Context
}

// NewServer 创建管理服务器实例
func NewServer(ctx context.Context, cfg *config.AdminConfig) *Server {
	return &Server{cfg: cfg, ctx: ctx}
}

// Start 启动管理服务器
func (s *Server) Start() error {
	// 设置 Gin 模式
	if !s.cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建 Gin 引擎
	r := gin.New()

	// 注册中间件
	r.Use(middleware.CORS()) // CORS 中间件（允许跨域）
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	// 启用 pprof
	if s.cfg.App.Debug {
		pprof.Register(r)
	}

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 注册路由
	s.registerRoutes(r)

	// 创建 HTTP 服务器
	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.HTTP.Port),
		Handler:      r,
		ReadTimeout:  s.cfg.HTTP.ReadTimeout,
		WriteTimeout: s.cfg.HTTP.WriteTimeout,
		IdleTimeout:  s.cfg.HTTP.IdleTimeout,
	}

	logger.Info("管理服务器启动", "port", s.cfg.HTTP.Port)

	// 启动服务器
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("管理服务器启动失败: %w", err)
	}

	return nil
}

// Stop 停止管理服务器
func (s *Server) Stop(ctx context.Context) error {
	logger.Info("正在关闭管理服务器...")
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// registerRoutes 注册路由
func (s *Server) registerRoutes(r *gin.Engine) {
	// 初始化依赖
	db := database.Get()

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
