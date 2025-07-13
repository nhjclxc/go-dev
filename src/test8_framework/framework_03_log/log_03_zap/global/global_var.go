package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	localConfig "log_03_zap/config"
)

// 全局变量放在这个文件里面


// zap 日志工具，用这个打日志
var GlobalZapLog *zap.Logger

// 读取配置的
var GlobalViper *viper.Viper


var GlobalConfig *localConfig.Server
