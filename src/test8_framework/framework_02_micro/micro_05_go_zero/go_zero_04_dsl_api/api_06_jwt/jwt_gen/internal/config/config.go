package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 读取 JWT 认证需要的密钥和过期时间配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

}
