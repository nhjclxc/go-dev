package core

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log_03_zap/global"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	if ok, _ := PathExists(global.GlobalConfig.Zap.Director); !ok { // 判断是否有Director文件夹，不存在则创建改文件夹
		fmt.Printf("create %v directory\n", global.GlobalConfig.Zap.Director)
		_ = os.Mkdir(global.GlobalConfig.Zap.Director, os.ModePerm)
	}
	// 获取日志级别
	levels := global.GlobalConfig.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	// 遍历出所有小于配置文件里面level的日志级别
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.GlobalConfig.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}



func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}