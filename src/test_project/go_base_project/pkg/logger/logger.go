package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"go_base_project/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	globalLogger     *slog.Logger
	lumberjackLogger *lumberjack.Logger
)

// Init 初始化日志
// 使用 lumberjack 库实现日志自动切割和清理功能
// 参数:
//   - cfg: 日志配置
//
// 返回值:
//   - error: 错误信息
func Init(cfg *config.LogConfig) error {
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 将所有日志时间转换为东八区（Asia/Shanghai, UTC+8）
			// 这样无论系统时区是什么，日志中显示的都是东八区时间
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					// 加载东八区时区
					loc, err := time.LoadLocation("Asia/Shanghai")
					if err != nil {
						// 如果加载失败，使用 FixedZone 创建 UTC+8
						loc = time.FixedZone("CST", 8*60*60)
					}
					return slog.Time(slog.TimeKey, t.In(loc))
				}
			}
			return a
		},
	}

	var writer io.Writer

	// 根据配置决定输出位置
	if cfg.Output == "file" {
		// 使用 lumberjack 实现日志自动切割和清理
		if cfg.FilePath == "" {
			cfg.FilePath = "logs"
		}

		// 确保日志目录存在
		if err := os.MkdirAll(cfg.FilePath, 0755); err != nil {
			return fmt.Errorf("创建日志目录失败: %w", err)
		}

		// 配置 lumberjack
		// 日志文件名格式: logs/go_base_project.log
		// 切割后格式: logs/go_base_project-2025-10-31T08-15-30.123.log
		lumberjackLogger = &lumberjack.Logger{
			Filename:   filepath.Join(cfg.FilePath, "go_base_project.log"), // 日志文件路径
			MaxSize:    100,                                                // 单个日志文件最大尺寸（MB），超过后自动切割
			MaxBackups: 0,                                                  // 保留的旧日志文件个数（0表示不限制个数，由MaxAge控制）
			MaxAge:     7,                                                  // 保留的旧日志文件天数（超过7天的日志会被自动删除）
			Compress:   true,                                               // 是否压缩旧日志文件（gzip格式，节省磁盘空间）
			LocalTime:  true,                                               // 使用本地时间（东八区）而非UTC时间
		}
		writer = lumberjackLogger
	} else {
		// 输出到标准输出
		writer = os.Stdout
	}

	var handler slog.Handler
	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(writer, opts)
	} else {
		handler = slog.NewTextHandler(writer, opts)
	}

	globalLogger = slog.New(handler)
	slog.SetDefault(globalLogger)

	return nil
}

// Get 获取全局日志实例
// 返回值:
//   - *slog.Logger: 日志实例
func Get() *slog.Logger {
	if globalLogger == nil {
		globalLogger = slog.Default()
	}
	return globalLogger
}

// Close 关闭日志（用于程序退出时）
// 返回值:
//   - error: 错误信息
func Close() error {
	if lumberjackLogger != nil {
		return lumberjackLogger.Close()
	}
	return nil
}

// Debug 调试日志
// 参数:
//   - msg: 日志消息
//   - args: 附加参数（键值对）
func Debug(msg string, args ...any) {
	Get().Debug(msg, args...)
}

// Info 信息日志
// 参数:
//   - msg: 日志消息
//   - args: 附加参数（键值对）
func Info(msg string, args ...any) {
	Get().Info(msg, args...)
}

// Warn 警告日志
// 参数:
//   - msg: 日志消息
//   - args: 附加参数（键值对）
func Warn(msg string, args ...any) {
	Get().Warn(msg, args...)
}

// Error 错误日志
// 参数:
//   - msg: 日志消息
//   - args: 附加参数（键值对）
func Error(msg string, args ...any) {
	Get().Error(msg, args...)
}
