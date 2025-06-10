package main

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"log"
	"net"
)

type Config struct {
	service.ServiceConf
	MaxUsers int
}

// 实现 Validator 接口，
// 那么go-zero就会自动调用这个方法
func (c Config) Validate() error {

	fmt.Printf("Config.Validate \n")

	if len(c.Name) == 0 {
		return errors.New("name 不能为空")
	}
	if c.MaxUsers <= 0 {
		return errors.New("max users 必须为正数")
	}

	// 验证ipv4地址是否正确 MetricsUrl
	parsedIP := net.ParseIP(c.MetricsUrl)
	if !(parsedIP != nil && parsedIP.To4() != nil) {
		return errors.New(c.MetricsUrl + " IPv4 地址不合法！！！")
	}

	return nil
}

// Auto Validation，自动验证配置是否正确
// https://go-zero.dev/docs/tutorials/go-zero/configuration/auto-validation
func main() {

	var config Config

	// 方式一
	//conf.MustLoad("etc/config.yaml", &config)
	//config.MustSetUp()

	// 方式二
	err := conf.Load("etc/config.yaml", &config)
	if err != nil {
		// 这里会捕获加载错误和验证错误
		log.Fatal(err)
	}

	// do your job

	fmt.Printf("config = %v \n", config)

	// 这项新功能引入了一个在加载后自动检查配置的验证机制。以下是使用方法：
}
