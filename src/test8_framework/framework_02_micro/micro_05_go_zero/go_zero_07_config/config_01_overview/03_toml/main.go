package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
)


type Config struct {
	Host   string `json:"Host"`
	Port   int    `json:"Port"`
	Device struct {
		Face struct {
			Yufan struct {
				AppKey    string `json:"appKey"`
				AppSecret string `json:"appSecret"`
			} `json:"Yufan"`
			Haik struct {
				BaseUrl string `json:"baseUrl"`
			} `json:"Haik"`
		} `json:"Face"`
		Vehicle struct {
			Palt string `json:"palt"`
		} `json:"vehicle"`
	} `json:"Device"`
}


var f = flag.String("f", "config.toml", "config file")


func main() {

	// go run main.go -f config.toml

	flag.Parse()
	var c Config
	conf.MustLoad(*f, &c)
	fmt.Printf("%#v", c)
	println()
	println()
	fmt.Printf("%v", c)

}
