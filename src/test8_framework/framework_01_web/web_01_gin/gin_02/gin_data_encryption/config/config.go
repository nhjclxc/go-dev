package config

// 配置结构体
type Config struct {
	Name    string  `mapstructure:"name"`
	Port    int     `mapstructure:"port"`
	Request Request `mapstructure:"request"`
}

type Request struct {
	Encrypt Encrypt `mapstructure:"encrypt"`
}

type Encrypt struct {
	Key           string `mapstructure:"key"`
	EncryptHeader string `mapstructure:"encryptHeader"`
}
