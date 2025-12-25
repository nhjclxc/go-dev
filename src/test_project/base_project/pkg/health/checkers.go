package health

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DatabaseChecker 数据库健康检查
type DatabaseChecker struct {
	db *gorm.DB
}

// NewDatabaseChecker 创建数据库检查器
func NewDatabaseChecker(db *gorm.DB) *DatabaseChecker {
	return &DatabaseChecker{db: db}
}

// Name 返回检查器名称
func (c *DatabaseChecker) Name() string {
	return "database"
}

// Check 执行数据库健康检查
func (c *DatabaseChecker) Check(ctx context.Context) CheckDetail {
	start := time.Now()

	sqlDB, err := c.db.DB()
	if err != nil {
		return CheckDetail{
			Status:  StatusDown,
			Message: fmt.Sprintf("获取数据库连接失败: %v", err),
			Latency: time.Since(start).String(),
		}
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return CheckDetail{
			Status:  StatusDown,
			Message: fmt.Sprintf("数据库连接失败: %v", err),
			Latency: time.Since(start).String(),
		}
	}

	return CheckDetail{
		Status:  StatusUp,
		Latency: time.Since(start).String(),
	}
}

// RedisChecker Redis 健康检查
type RedisChecker struct {
	client *redis.Client
}

// NewRedisChecker 创建 Redis 检查器
func NewRedisChecker(client *redis.Client) *RedisChecker {
	return &RedisChecker{client: client}
}

// Name 返回检查器名称
func (c *RedisChecker) Name() string {
	return "redis"
}

// Check 执行 Redis 健康检查
func (c *RedisChecker) Check(ctx context.Context) CheckDetail {
	start := time.Now()

	if err := c.client.Ping(ctx).Err(); err != nil {
		return CheckDetail{
			Status:  StatusDown,
			Message: fmt.Sprintf("Redis 连接失败: %v", err),
			Latency: time.Since(start).String(),
		}
	}

	return CheckDetail{
		Status:  StatusUp,
		Latency: time.Since(start).String(),
	}
}

// CustomChecker 自定义健康检查
type CustomChecker struct {
	name    string
	checkFn func(ctx context.Context) CheckDetail
}

// NewCustomChecker 创建自定义检查器
func NewCustomChecker(name string, checkFn func(ctx context.Context) CheckDetail) *CustomChecker {
	return &CustomChecker{
		name:    name,
		checkFn: checkFn,
	}
}

// Name 返回检查器名称
func (c *CustomChecker) Name() string {
	return c.name
}

// Check 执行自定义健康检查
func (c *CustomChecker) Check(ctx context.Context) CheckDetail {
	return c.checkFn(ctx)
}
