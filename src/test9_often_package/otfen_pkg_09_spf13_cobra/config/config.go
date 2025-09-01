package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
)

type Config struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

func InitConfig(cfg *Config, cfgFile string) error {
	// default 用于打印默认配置

	if cfgFile != "" {
		// 显式指定配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 默认路径: ./config/config.yaml
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		slog.Info("未指定配置文件路径，使用默认配置文件！！！", slog.String("config", "./config/config.yaml"))
	}

	// 环境变量支持（可选）
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件出错:", err)
		return err
	}

	// 绑定到结构体
	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("解析配置文件出错:", err)
		return err
	}

	fmt.Println("使用配置文件:", viper.ConfigFileUsed())
	return nil
}
