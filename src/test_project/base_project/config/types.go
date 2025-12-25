package config

// 通用的结构体写在这里

// AppConfig 应用配置
type AppConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`       // 日志级别: debug, info, warn, error
	Format     string `mapstructure:"format"`      // 日志格式: json, text
	Output     string `mapstructure:"output"`      // 输出方式: stdout, file
	FilePath   string `mapstructure:"file_path"`   // 日志文件路径
	MaxSize    int    `mapstructure:"max_size"`    // 单个文件最大大小（MB）
	MaxBackups int    `mapstructure:"max_backups"` // 保留的旧文件数量
	MaxAge     int    `mapstructure:"max_age"`     // 保留的最大天数
	Compress   bool   `mapstructure:"compress"`    // 是否压缩旧文件
}

// CronConfig 定时任务配置
type CronConfig struct {
	Timezone string     `mapstructure:"timezone"`
	Tasks    []CronTask `mapstructure:"tasks"`
}
