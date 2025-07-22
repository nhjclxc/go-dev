package global

import (
	localConfig "gin_data_encryption/config"
	"github.com/spf13/viper"
)

// 全局变量放在这个文件里面

// 读取配置的
var GlobalViper *viper.Viper

// 全局的配置结构体对象
var GlobalConfig *localConfig.Config
