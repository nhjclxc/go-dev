package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

type RestConf struct {
	service.ServiceConf
	Host     string `json:",default=0.0.0.0"`
	Port     int
	CertFile string `json:",optional"`
	KeyFile  string `json:",optional"`
	Verbose  bool   `json:",optional"`
	MaxConns int    `json:",default=10000"`
	MaxBytes int64  `json:",default=1048576"`
	// milliseconds
	Timeout      int64 `json:",default=3000"`
	CpuThreshold int64 `json:",default=900,range=[0:1000]"`
}

func main() {
	var restConf RestConf
	conf.MustLoad("etc/config.yaml", &restConf)

	restConf.MustSetUp()
	// do your job

	fmt.Printf("restConf = %v \n", restConf)
}
