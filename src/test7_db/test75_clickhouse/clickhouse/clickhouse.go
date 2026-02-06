package ch

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// ClickHouseClient ClickHouse客户端
type ClickHouseClient struct {
	DB   *sql.DB         // database/sql（查询）
	Conn clickhouse.Conn // 原生接口（批量写）
}

// New 创建新的 ClickHouse 连接
func New(cfg *ClickHouseConfig) (*ClickHouseClient, error) {
	opts := clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
			"send_logs_level":    "trace",
		},
		DialTimeout: cfg.DialTimeout,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			//logger.Info("SQL: %s, Duration: %s", logEvent.Text, logEvent.Time)
			//fmt.Println(fmt.Sprintf(format, v...))
		},
		Logger:               NewSlog(),
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
	}

	// 1、原生 Conn（用于批量写入）
	conn, _ := clickhouse.Open(&opts)

	// 2、database/sql（用于查询）
	db := clickhouse.OpenDB(&opts)

	// 设置连接池参数
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// 输出执行的sql
	ctx := clickhouse.Context(context.Background(), clickhouse.WithLogs(func(log *clickhouse.Log) {
		fmt.Printf("SQL: %s, Duration: %s\n", log.Text, log.Time)
	}))

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("连接 ClickHouse 失败: %w", err)
	}

	client := ClickHouseClient{
		Conn: conn,
		DB:   db,
	}

	return &client, nil
}

// Close 关闭 ClickHouse 连接
func (cli *ClickHouseClient) Close() error {
	if cli.DB == nil {
		return nil
	}
	return cli.DB.Close()
}

func NewSlog() *slog.Logger {

	var writer io.Writer = os.Stdout
	level := slog.LevelInfo
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

	var handler slog.Handler = slog.NewJSONHandler(writer, opts)

	globalLogger := slog.New(handler)
	slog.SetDefault(globalLogger)

	return globalLogger
}

var ClickHouseCli *ClickHouseClient

func init() {

	cfg := ClickHouseConfig{
		Host:            "localhost",
		Port:            9000,
		Database:        "default",
		Username:        "default",
		Password:        "111111",
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: 1 * time.Hour,
		DialTimeout:     10 * time.Second,
		ReadTimeout:     10 * time.Second,
	}

	var err error
	ClickHouseCli, err = New(&cfg)
	if err != nil {
		fmt.Printf("初始化 ClickHouse 失败 : %s", err.Error())
	}

}
