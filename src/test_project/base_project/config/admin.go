package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config 全局配置结构
type Config struct {
	App        AppConfig        `mapstructure:"app"`
	HTTP       HTTPConfig       `mapstructure:"http"`
	GRPC       GRPCConfig       `mapstructure:"grpc"`
	Login      LoginConfig      `mapstructure:"login"`
	Database   DatabaseConfig   `mapstructure:"database"`
	ClickHouse ClickHouseConfig `mapstructure:"clickhouse"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Log        LogConfig        `mapstructure:"log"`
	Cron       CronConfig       `mapstructure:"cron"`
	Telemetry  TelemetryConfig  `mapstructure:"telemetry"`
}

// HTTPConfig HTTP 服务配置
type HTTPConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// GRPCConfig gRPC 服务配置
type GRPCConfig struct {
	Port    int            `mapstructure:"port"`
	Network string         `mapstructure:"network"`
	Auth    GRPCAuthConfig `mapstructure:"auth"`
}

// GRPCAuthConfig gRPC 鉴权配置
type GRPCAuthConfig struct {
	Enabled bool   `mapstructure:"enabled"` // 是否启用鉴权
	Token   string `mapstructure:"token"`   // 预共享密钥
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// ClickHouseConfig ClickHouse 配置
type ClickHouseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Database        string        `mapstructure:"database"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	DialTimeout     time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// CronTask 单个定时任务配置
type CronTask struct {
	Name    string `mapstructure:"name"`
	Spec    string `mapstructure:"spec"`
	Enabled bool   `mapstructure:"enabled"`
}

// TelemetryConfig 遥测配置（Tracing + Metrics）
type TelemetryConfig struct {
	Enabled     bool           `mapstructure:"enabled"`
	ServiceName string         `mapstructure:"service_name"`
	Tracing     TracingConfig  `mapstructure:"tracing"`
	Metrics     MetricsConfig  `mapstructure:"metrics"`
	Exporter    ExporterConfig `mapstructure:"exporter"`
}

// TracingConfig 链路追踪配置
type TracingConfig struct {
	Enabled    bool    `mapstructure:"enabled"`
	SampleRate float64 `mapstructure:"sample_rate"`
}

// MetricsConfig 指标监控配置
type MetricsConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ExporterConfig 导出器配置
type ExporterConfig struct {
	Type     string        `mapstructure:"type"`     // otlp, stdout
	Endpoint string        `mapstructure:"endpoint"` // OTLP endpoint
	Timeout  time.Duration `mapstructure:"timeout"`
}

// LoginConfig 登录相关配置
type LoginConfig struct {
	JWT *JWTConfig `mapstructure:"jwt"`
}

// JWTConfig JWT相关配置
type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key"` // JWT签名密钥
	ExpiresIn int    `mapstructure:"expires_in"` // 过期时间(小时)
}

// LoadAdmin 加载配置文件
func LoadAdmin(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// 设置默认值
	setDefaults(v)

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &cfg, nil
}

// setDefaults 设置默认配置值
func setDefaults(v *viper.Viper) {
	// App 默认值
	v.SetDefault("app.name", "base_project")
	v.SetDefault("app.env", "development")
	v.SetDefault("app.debug", true)

	// HTTP 默认值
	v.SetDefault("http.port", 8080)
	v.SetDefault("http.read_timeout", "30s")
	v.SetDefault("http.write_timeout", "30s")
	v.SetDefault("http.idle_timeout", "60s")

	// gRPC 默认值
	v.SetDefault("grpc.port", 9090)
	v.SetDefault("grpc.network", "tcp")
	v.SetDefault("grpc.auth.enabled", false)
	v.SetDefault("grpc.auth.token", "")

	// ClickHouse 默认值
	v.SetDefault("clickhouse.host", "localhost")
	v.SetDefault("clickhouse.port", 9000)
	v.SetDefault("clickhouse.database", "base_project")
	v.SetDefault("clickhouse.username", "default")
	v.SetDefault("clickhouse.password", "")
	v.SetDefault("clickhouse.max_open_conns", 10)
	v.SetDefault("clickhouse.max_idle_conns", 5)
	v.SetDefault("clickhouse.conn_max_lifetime", "1h")
	v.SetDefault("clickhouse.dial_timeout", "10s")
	v.SetDefault("clickhouse.read_timeout", "30s")

	// 日志默认值
	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "json")
	v.SetDefault("log.output", "stdout")
	v.SetDefault("log.file_path", "logs/app.log")
	v.SetDefault("log.max_size", 100)   // 100MB
	v.SetDefault("log.max_backups", 10) // 保留 10 个文件
	v.SetDefault("log.max_age", 30)     // 保留 30 天
	v.SetDefault("log.compress", true)  // 压缩旧文件

	// Cron 默认值
	v.SetDefault("cron.timezone", "Asia/Shanghai")

	// Telemetry 默认值
	v.SetDefault("telemetry.enabled", false)
	v.SetDefault("telemetry.service_name", "base_project")
	v.SetDefault("telemetry.tracing.enabled", true)
	v.SetDefault("telemetry.tracing.sample_rate", 1.0)
	v.SetDefault("telemetry.metrics.enabled", true)
	v.SetDefault("telemetry.exporter.type", "otlp")
	v.SetDefault("telemetry.exporter.endpoint", "localhost:4318")
	v.SetDefault("telemetry.exporter.timeout", "10s")

	// HashRing 默认值
	v.SetDefault("scheduler.hash_ring.virtual_nodes", 150)
	v.SetDefault("scheduler.hash_ring.replica_count", 2)

	// Selector 默认值
	v.SetDefault("scheduler.selector.enable_region_match", true)
	v.SetDefault("scheduler.selector.enable_isp_match", true)
	v.SetDefault("scheduler.selector.fallback_to_any", true)
	v.SetDefault("scheduler.selector.max_bandwidth_usage", 0.95)
	v.SetDefault("scheduler.selector.max_error_rate", 0.05)

	// Sync 默认值
	v.SetDefault("scheduler.sync.refresh_interval", "5s")
	v.SetDefault("scheduler.sync.enable_pubsub", true)
}
