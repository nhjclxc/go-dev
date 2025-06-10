package svc

import (
	"go_zero_09_http/internal/config"
	"go_zero_09_http/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	// 注册自定义中间件
	CustomMiddleware  *middleware.CustomMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 注入自定义中间件
		CustomMiddleware: middleware.NewCustomMiddleware(),
	}
}


