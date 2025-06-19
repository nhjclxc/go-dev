package config

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql MysqlConf `json:"mysql"`      // 数据库配置
}

type MysqlConf struct {
	DataSource string
}

func (m MysqlConf) NewSqlConn() sqlx.SqlConn {
	return sqlx.NewMysql(m.DataSource)
}