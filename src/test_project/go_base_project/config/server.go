package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// ServerConfig 服务端配置
type ServerConfig struct {
	App         *ServerAppConfig `mapstructure:"app"`
	Log         *LogConfig       `mapstructure:"log"`
	HTTP        *HTTPConfig      `mapstructure:"http"`
	Database    *DatabaseConfig  `mapstructure:"database"`
	RedisConfig *RedisConfig     `mapstructure:"redis"`
	Cron        *CronConfig      `mapstructure:"cron"`
	Login       *LoginConfig     `mapstructure:"login"`
}

// ServerAppConfig 应用配置
type ServerAppConfig struct {
	Name   string `mapstructure:"name"`
	Env    string `mapstructure:"env"`
	Debug  bool   `mapstructure:"debug"`
	Ipdata string `mapstructure:"ipdata"`
}

// HTTPConfig HTTP 服务配置
type HTTPConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// CronTask 单个定时任务配置
type CronTask struct {
	Name    string `mapstructure:"name"`
	Spec    string `mapstructure:"spec"`
	Enabled bool   `mapstructure:"enabled"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// LoginConfig 登录相关配置
type LoginConfig struct {
	JWT *JWTConfig `mapstructure:"jwt"`
}

// JWTConfig JWT相关配置
type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key"` // JWT签名密钥
	ExpiresIn int    `mapstructure:"expires_in"` // 过期时间(小时)
}

// LoadServerConfig 加载管理服务配置
// 参数:
//   - configPath: 配置文件路径
//
// 返回值:
//   - *ServerConfig: 管理服务配置实例
//   - error: 错误信息
func LoadServerConfig(configPath string) (*ServerConfig, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置
	var cfg ServerConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &cfg, nil
}
