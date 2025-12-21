package config

// 通用配置类

// LogConfig 日志配置
type LogConfig struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Output   string `mapstructure:"output"`
	FilePath string `mapstructure:"file_path"`
}

// CronConfig 定时任务配置
type CronConfig struct {
	Timezone string     `mapstructure:"timezone"`
	Tasks    []CronTask `mapstructure:"tasks"`
}
