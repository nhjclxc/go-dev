package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
)

var env = flag.String("env", "dev", "environment: dev | test | prod")

type MainConfig struct {
	AppName string `json:",default=env-practice-111"`
	Env struct {
		Avtive string `json:",default=dev"`
	}
}
type EnvConfig struct {
	Host string
	Port int
}

// go run main02.go
// go run main02.go -env=dev
// go run main02.go -env=test
// go run main02.go -env=prod
func main() {
	flag.Parse()

	// 先加载主配置文件
	mainConfFile := fmt.Sprintf("etc/config.yaml")
	var mainConfig MainConfig
	conf.MustLoad(mainConfFile, &mainConfig)

	fmt.Printf("主配置：%+v\n", mainConfig)
	//fmt.Printf("环境变量：%v\n", mainConfig.Env.Avtive)
	fmt.Printf("环境变量：%v\n", *env)

	if env == nil && mainConfig.Env.Avtive != "" {
		*env = mainConfig.Env.Avtive
	}
	if env == nil || *env == "" {
		*env = "dev"
	}

	// 再加载环境文件
	envConfFile := fmt.Sprintf("etc/config-%s.yaml", *env)
	var envConfig EnvConfig
	conf.MustLoad(envConfFile, &envConfig)

	fmt.Printf("%s 环境配置：%+v\n", *env, envConfig)


}