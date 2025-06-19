package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_19_db_mysql/internal/config"
)

type ServiceContext struct {
	Config config.Config

	// 持有数据库句柄
	SqlConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		SqlConn: sqlx.NewMysql(c.Mysql.DataSource),
	}

}
