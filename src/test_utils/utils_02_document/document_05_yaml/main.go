package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// YAML Ain’t Markup Language，一种非常简洁的非标记语言，可以快速的对Yaml进行编码和解码。
//
// 官网地址：https://gopkg.in/yaml.v2
//
// GoDoc：https://godoc.org/gopkg.in/yaml.v2

// go get -u gopkg.in/yaml.v2
func main() {
	/*
		基本规则
			大小写敏感、易于使用，容易阅读
			使用缩进表示层级关系
			只能使用空格键
			适合表示程序语言的数据结构
			缩进长度没有限制，只要元素对齐就表示这些元素属于一个层级
			使用#表示注释
			字符串可以不用引号标注
			可用于不同程序间交换数据
			支持泛型工具
			丰富的表达能力和可扩展性
	*/

	filename := "src/test_utils/utils_02_document/document_05_yaml/application.yaml"
	y := new(ServiceYaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file err %v\n", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("yaml unmarshal: %v\n", err)
		return
	}

	fmt.Printf("MySQL host : %s, port : %d, anonymous_user : %s, password : %s, db_name : %s \n",
		y.MySQL.Host, y.MySQL.Port, y.MySQL.User, y.MySQL.Password, y.MySQL.DbName)

	fmt.Printf("Redis host : %s, port %d, auth : %s\n",
		y.Redis.Host, y.Redis.Port, y.Redis.Auth)

	fmt.Printf("Vip Counter: %d, Vip List : %v\n", y.NginxProxy.Counter, y.NginxProxy.NginxList)

	n, err := yaml.Marshal(y)
	if err != nil {
		log.Fatalf("marshal err : %v\n", err)
		return
	}
	fmt.Printf("yaml marshal : %v\n", string(n))

}

type MySQLConfig struct {
	User     string `yaml:"anonymous_user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
}
type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth string `yaml:"auth"`
}
type NginxProxyConfig struct {
	Counter   int      `yaml:"counter"`
	NginxList []string `yaml:"nginx_list"`
}

type ServiceYaml struct {
	MySQL      MySQLConfig      `yaml:"mysql"`
	Redis      RedisConfig      `yaml:"redis"`
	NginxProxy NginxProxyConfig `yaml:"nginx_proxy"`
}
