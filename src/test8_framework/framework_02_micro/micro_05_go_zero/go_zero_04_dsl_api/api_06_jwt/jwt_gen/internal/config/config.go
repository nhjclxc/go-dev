package config

import (
	"github.com/zeromicro/go-zero/rest"
)

// Config 结构体读取的是 anonymous_user-api.yaml 配置文件里面的配置
type Config struct {
	// rest.RestConf 通过匿名字段嵌入
	// rest.RestConf 定义了go-zero 里面一些默认的配置
	rest.RestConf

	// 读取 JWT 认证需要的密钥和过期时间配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
