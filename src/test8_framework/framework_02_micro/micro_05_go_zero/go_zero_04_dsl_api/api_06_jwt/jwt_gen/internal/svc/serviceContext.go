package svc

import (
	"jwt_gen/internal/config"
	"jwt_gen/internal/middleware"
	"jwt_gen/internal/middleware/jwt"
)

type ServiceContext struct {
	Config config.Config
	// 注册自定义 jwt 中间件
	JwtAuthMiddleware  *jwt.JwtAuthMiddleware
	UserAgentMiddleware  *middleware.UserAgentMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 注入 jwt 中间件，要将配置文件夹传入，读取token密钥
		JwtAuthMiddleware: jwt.NewAuthMiddleware(c),
		UserAgentMiddleware: middleware.NewUserAgentMiddleware(),
	}
}
