在 Go 项目中使用 `Gin` + `Zipkin` 实现链路追踪，你可以使用 `OpenTelemetry` 或者直接使用 `OpenZipkin` 的 Go 库实现 Zipkin 的追踪功能。以下是推荐的方案：

---

## ✅ 推荐方式：`Gin` + `OpenTelemetry` + `Zipkin`

OpenTelemetry 是统一标准，未来趋势，支持导出到 Zipkin、Jaeger、Prometheus 等后端。

---

### 🔧 一、依赖安装

```bash
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/sdk
go get go.opentelemetry.io/otel/exporters/zipkin
go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
```

---

### 🧱 二、初始化 OpenTelemetry + Zipkin

```go
package main

import (
    "context"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/zipkin"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitTracer(serviceName string, zipkinURL string) func(context.Context) error {
    // 创建 Zipkin exporter
    exporter, err := zipkin.New(
        zipkinURL,
    )
    if err != nil {
        log.Fatalf("failed to create zipkin exporter: %v", err)
    }

    // 创建 Tracer Provider
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String(serviceName),
        )),
    )

    otel.SetTracerProvider(tp)

    return tp.Shutdown
}
```

---

### 🚀 三、Gin 中集成链路追踪中间件

```go
package main

import (
    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
    shutdown := InitTracer("gin-service", "http://localhost:9411/api/v2/spans")
    defer shutdown(context.Background())

    r := gin.Default()

    // 添加 OpenTelemetry 中间件（自动注入 traceId/spanId）
    r.Use(otelgin.Middleware("gin-service"))

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"msg": "Hello, trace!"})
    })

    r.Run(":8080")
}
```

---

### 📦 四、运行 Zipkin

你可以用 Docker 快速启动一个 Zipkin 实例：

```bash
docker run -d -p 9411:9411 openzipkin/zipkin
```

打开浏览器访问：`http://localhost:9411` 可查看追踪数据。

---

### 🧪 五、效果演示

1. 启动 Go 服务。
2. 请求一次：`curl http://localhost:8080/hello`
3. 打开 Zipkin UI，在 trace 页面看到请求链路信息（traceId、duration、spanName 等）。

---

## 📌 关键点说明

| 名称                         | 说明                                        |
| -------------------------- | ----------------------------------------- |
| `otelgin.Middleware`       | 自动为每个 HTTP 请求创建 trace span                |
| `traceId`                  | 全局唯一 ID，表示一条完整的调用链                        |
| `spanId`                   | 当前调用的唯一 ID（一次 HTTP 请求）                    |
| `otel.SetTracerProvider()` | 设置全局的 Tracer 实例                           |
| `zipkin.New()`             | 创建一个 Zipkin exporter，将 trace 数据发送给 Zipkin |

---

## ✅ 附加功能（可选）

* 使用 `otel.Tracer().Start()` 创建自定义 span
* 把 `traceId` 打印到日志中（如 zap logger 中注入）
* 支持跨服务链路追踪（通过请求头 `traceparent`）

---

需要我补充 **调用其他微服务时透传 traceId** 或 **集成 zap 日志打印 traceId** 的实现吗？
