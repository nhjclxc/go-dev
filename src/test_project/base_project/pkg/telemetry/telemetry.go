package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"

	"base_project/config"
)

// Telemetry 遥测组件，包含 Tracing 和 Metrics
type Telemetry struct {
	TracerProvider *sdktrace.TracerProvider
	MeterProvider  *sdkmetric.MeterProvider
	enabled        bool
}

// New 创建新的遥测实例
func New(cfg *config.TelemetryConfig) (*Telemetry, error) {
	tel := &Telemetry{
		enabled: cfg.Enabled,
	}

	if !cfg.Enabled {
		// 如果未启用，返回空的 Telemetry 实例
		return tel, nil
	}

	// 创建资源描述
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("创建资源描述失败: %w", err)
	}

	ctx := context.Background()

	// 初始化 Tracing
	if cfg.Tracing.Enabled {
		tp, err := initTracing(ctx, cfg, res)
		if err != nil {
			return nil, fmt.Errorf("初始化 Tracing 失败: %w", err)
		}
		tel.TracerProvider = tp
		otel.SetTracerProvider(tp)
	}

	// 初始化 Metrics
	if cfg.Metrics.Enabled {
		mp, err := initMetrics(ctx, cfg, res)
		if err != nil {
			// 如果 Tracing 已初始化，先关闭它
			if tel.TracerProvider != nil {
				_ = tel.TracerProvider.Shutdown(ctx)
			}
			return nil, fmt.Errorf("初始化 Metrics 失败: %w", err)
		}
		tel.MeterProvider = mp
		otel.SetMeterProvider(mp)
	}

	// 设置全局 propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tel, nil
}

// initTracing 初始化链路追踪
func initTracing(ctx context.Context, cfg *config.TelemetryConfig, res *resource.Resource) (*sdktrace.TracerProvider, error) {
	var exporter sdktrace.SpanExporter
	var err error

	timeout := cfg.Exporter.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	switch cfg.Exporter.Type {
	case "stdout":
		exporter, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
	case "otlp":
		fallthrough
	default:
		exporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithEndpoint(cfg.Exporter.Endpoint),
			otlptracehttp.WithInsecure(), // 开发环境使用，生产环境应配置 TLS
			otlptracehttp.WithTimeout(timeout),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("创建 trace exporter 失败: %w", err)
	}

	// 配置采样率
	sampler := sdktrace.AlwaysSample()
	if cfg.Tracing.SampleRate < 1.0 {
		sampler = sdktrace.TraceIDRatioBased(cfg.Tracing.SampleRate)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sampler),
	)

	return tp, nil
}

// initMetrics 初始化指标监控
func initMetrics(ctx context.Context, cfg *config.TelemetryConfig, res *resource.Resource) (*sdkmetric.MeterProvider, error) {
	var exporter sdkmetric.Exporter
	var err error

	timeout := cfg.Exporter.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	switch cfg.Exporter.Type {
	case "stdout":
		exporter, err = stdoutmetric.New()
	case "otlp":
		fallthrough
	default:
		exporter, err = otlpmetrichttp.New(ctx,
			otlpmetrichttp.WithEndpoint(cfg.Exporter.Endpoint),
			otlpmetrichttp.WithInsecure(), // 开发环境使用，生产环境应配置 TLS
			otlpmetrichttp.WithTimeout(timeout),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("创建 metric exporter 失败: %w", err)
	}

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter,
			sdkmetric.WithInterval(15*time.Second), // 每 15 秒导出一次指标
		)),
		sdkmetric.WithResource(res),
	)

	return mp, nil
}

// Shutdown 关闭遥测组件
func (t *Telemetry) Shutdown(ctx context.Context) error {
	if !t.enabled {
		return nil
	}

	var errs []error

	if t.TracerProvider != nil {
		if err := t.TracerProvider.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("关闭 TracerProvider 失败: %w", err))
		}
	}

	if t.MeterProvider != nil {
		if err := t.MeterProvider.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("关闭 MeterProvider 失败: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("关闭遥测组件时发生错误: %v", errs)
	}

	return nil
}

// IsEnabled 返回遥测是否启用
func (t *Telemetry) IsEnabled() bool {
	return t.enabled
}

// Tracer 返回一个命名的 tracer
func Tracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
