package database

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"base_project/config"
)

// New 创建新的数据库连接
func New(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Second)

	// 注册 OpenTelemetry tracing callbacks
	registerTracingCallbacks(db)

	return db, nil
}

// Close 关闭数据库连接
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// registerTracingCallbacks 注册 GORM tracing 回调
func registerTracingCallbacks(db *gorm.DB) {
	tracer := otel.Tracer("gorm")

	// Query
	_ = db.Callback().Query().Before("gorm:query").Register("otel:before_query", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.query")
	})
	_ = db.Callback().Query().After("gorm:query").Register("otel:after_query", endSpan)

	// Create
	_ = db.Callback().Create().Before("gorm:create").Register("otel:before_create", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.create")
	})
	_ = db.Callback().Create().After("gorm:create").Register("otel:after_create", endSpan)

	// Update
	_ = db.Callback().Update().Before("gorm:update").Register("otel:before_update", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.update")
	})
	_ = db.Callback().Update().After("gorm:update").Register("otel:after_update", endSpan)

	// Delete
	_ = db.Callback().Delete().Before("gorm:delete").Register("otel:before_delete", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.delete")
	})
	_ = db.Callback().Delete().After("gorm:delete").Register("otel:after_delete", endSpan)

	// Raw
	_ = db.Callback().Raw().Before("gorm:raw").Register("otel:before_raw", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.raw")
	})
	_ = db.Callback().Raw().After("gorm:raw").Register("otel:after_raw", endSpan)

	// Row
	_ = db.Callback().Row().Before("gorm:row").Register("otel:before_row", func(db *gorm.DB) {
		startSpan(db, tracer, "gorm.row")
	})
	_ = db.Callback().Row().After("gorm:row").Register("otel:after_row", endSpan)
}

func startSpan(db *gorm.DB, tracer trace.Tracer, spanName string) {
	if db.Statement == nil || db.Statement.Context == nil {
		return
	}

	ctx, span := tracer.Start(db.Statement.Context, spanName)
	db.Statement.Context = ctx
	db.InstanceSet("otel:span", span)
	db.InstanceSet("otel:start_time", time.Now())
}

func endSpan(db *gorm.DB) {
	val, ok := db.InstanceGet("otel:span")
	if !ok {
		return
	}

	span, ok := val.(trace.Span)
	if !ok {
		return
	}

	defer span.End()

	// 设置 SQL 语句属性
	if db.Statement != nil {
		span.SetAttributes(
			attribute.String("db.system", "mysql"),
			attribute.String("db.statement", db.Statement.SQL.String()),
			attribute.String("db.table", db.Statement.Table),
		)

		// 设置影响行数
		span.SetAttributes(attribute.Int64("db.rows_affected", db.RowsAffected))
	}

	// 记录错误
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		span.RecordError(db.Error)
		span.SetStatus(codes.Error, db.Error.Error())
	}

	// 计算耗时
	if startTime, ok := db.InstanceGet("otel:start_time"); ok {
		if t, ok := startTime.(time.Time); ok {
			span.SetAttributes(attribute.Int64("db.duration_ms", time.Since(t).Milliseconds()))
		}
	}
}

// WithContext 为数据库操作添加 context（便捷方法）
func WithContext(db *gorm.DB, ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
