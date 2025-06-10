package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	LogConf logx.LogConf `json:"log"` // 对应配置文件中的 log 节点
}


func main() {

	// 定义配置结构体
	var c Config

	// 读取配置文件
	conf.MustLoad("etc/config.yaml", &c)

	// 加载配置文件的log部分
	logx.MustSetup(c.LogConf)

	logx.Debug(context.Background(), "logx的 Debug 日志输出")
	logx.Info(context.Background(), "logx的 Info 日志输出")
	logx.Error(context.Background(), "logx的 Error 日志输出")

	// do your job

	fmt.Printf("c = %v \n", c)
}
