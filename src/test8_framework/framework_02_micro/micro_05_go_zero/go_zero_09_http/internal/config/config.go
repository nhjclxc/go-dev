package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 中间件的定义，这个是指go-zero系统自带的中间件
	// 要使用这些中间件的话要去配置文件里面开启，true标识开启这些对于的中间件
	// 注意在rest.RestConf里面已经自带了Middlewares MiddlewaresConf，因此下面这条语句别写写了回属性冲突
	//Middlewares MiddlewaresConf

}

//type MiddlewaresConf struct {
//	Trace      bool `json:",default=true"`
//	Log        bool `json:",default=true"`
//	Prometheus bool `json:",default=true"`
//	MaxConns   bool `json:",default=true"`
//	Breaker    bool `json:",default=true"`
//	Shedding   bool `json:",default=true"`
//	Timeout    bool `json:",default=true"`
//	Recover    bool `json:",default=true"`
//	Metrics    bool `json:",default=true"`
//	MaxBytes   bool `json:",default=true"`
//	Gunzip     bool `json:",default=true"`
//}