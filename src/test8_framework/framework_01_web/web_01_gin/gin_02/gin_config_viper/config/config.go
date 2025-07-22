package config

// 读取 config.yaml 的配置
type Server struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"` // 级别
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 日志保留天数
}
