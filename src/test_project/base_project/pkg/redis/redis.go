package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"

	"base_project/config"
)

// New 创建新的 Redis 客户端
func New(cfg *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	// 测试连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("连接 Redis 失败: %w", err)
	}

	// 启用 OpenTelemetry tracing
	if err := redisotel.InstrumentTracing(client); err != nil {
		return nil, fmt.Errorf("启用 Redis tracing 失败: %w", err)
	}

	// 启用 OpenTelemetry metrics
	if err := redisotel.InstrumentMetrics(client); err != nil {
		return nil, fmt.Errorf("启用 Redis metrics 失败: %w", err)
	}

	return client, nil
}
