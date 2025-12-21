package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	App  *ClientAppConfig `mapstructure:"app"`
	Log  *LogConfig       `mapstructure:"log"`
	Cron CronConfig       `mapstructure:"cron"`
}

// ClientAppConfig 应用配置
type ClientAppConfig struct {
	Name   string `mapstructure:"name"`
	Env    string `mapstructure:"env"`
	Debug  bool   `mapstructure:"debug"`
	Ipdata string `mapstructure:"ipdata"`
}

// LoadClientConfig 加载客户端配置
// 参数:
//   - configPath: 配置文件路径
//
// 返回值:
//   - *ClientConfig: 客户端配置实例
//   - error: 错误信息
func LoadClientConfig(configPath string) (*ClientConfig, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置
	var cfg ClientConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &cfg, nil
}
