package svc

import (
	"github.com/robfig/cron/v3"
	"go_zero_11_cron/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cron   *cron.Cron
}

func NewServiceContext(c config.Config) *ServiceContext {
	cronJob := cron.New() // 默认支持秒级调度从 v3 开始 cron.New(cron.WithSeconds())

	return &ServiceContext{
		Config: c,
		Cron:   cronJob,
	}
}
