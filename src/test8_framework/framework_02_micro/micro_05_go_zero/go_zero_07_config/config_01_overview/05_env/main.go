package main
//
//import (
//	"flag"
//	"fmt"
//	"github.com/zeromicro/go-zero/core/conf"
//)
//
//var env = flag.String("env", "dev", "environment: dev | test | prod")
//
//type Config struct {
//	AppName string `json:",default=env-practice-111"`
//	Env struct {
//		Avtive string `json:",default=dev"`
//	}
//
//	Host string
//	Port int
//}
//
//// go run main.go -env=dev
//// go run main.go -env dev
//func main01() {
//	flag.Parse()
//
//	if env == nil {
//		*env = "dev"
//	}
//
//	confFile := fmt.Sprintf("etc/config-%s.yaml", *env)
//
//	var c Config
//	conf.MustLoad(confFile, &c)
//
//	fmt.Printf("当前环境：%s\n", *env)
//	fmt.Printf("配置：%+v\n", c)
//}