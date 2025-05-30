package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
)

// json
//yaml | yml
//toml

// 定义读取配置的结构体
// config.yaml文件里面的配置是怎么样的，这里的Config结构体就定义成什么样
type Config struct {
	Host string `json:",default=0.0.0.0"`
	Port int

	Device struct {
		Face struct {
			Yufan struct {
				// 大小写不敏感
				AppKey string
				AppSecret string
			}
			Haik struct {
				// 大小写不敏感
				BaseUrl string
			}
		}
		Vehicle struct {
			Palt string
		}
	}

}



var f = flag.String("f", "config.yaml", "config file")

func main() {
	flag.Parse()
	var c Config
	// 使用 func MustLoad(path string, v interface{}, opts ...Option) 进行加载配置，path 为配置的路径，v 为结构体。 这个方法会完成配置的加载，如果配置加载失败，整个程序会 fatal 停止掉。
	conf.MustLoad(*f, &c)
	println(c.Host)
	fmt.Printf("%#v", c)
	println()
	println()
	fmt.Printf("%v", c)
	fmt.Println(c.Device)
	fmt.Println(c.Device.Face)
	fmt.Println(c.Device.Face.Yufan)
	fmt.Println(c.Device.Face.Yufan.AppKey)
	fmt.Println(c.Device.Face.Yufan.AppSecret)
	fmt.Println(c.Device.Face.Haik)
	fmt.Println(c.Device.Face.Haik.BaseUrl)
	fmt.Println(c.Device.Vehicle)
	fmt.Println(c.Device.Vehicle.Palt)



}