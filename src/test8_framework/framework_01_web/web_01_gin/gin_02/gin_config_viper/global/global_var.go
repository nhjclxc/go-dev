package global

import (
	"github.com/spf13/viper"
	localConfig "gin_config_viper/config"
)

// 全局变量放在这个文件里面


// 读取配置的
var GlobalViper *viper.Viper

// 全局的配置结构体对象
var GlobalConfig *localConfig.Server
