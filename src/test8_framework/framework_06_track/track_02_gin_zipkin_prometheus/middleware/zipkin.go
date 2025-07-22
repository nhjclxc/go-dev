// middleware/tracing.go
package middleware

import (
	"github.com/gin-gonic/gin"
	zipkin "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/propagation/b3"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const zipkinUrl = "http://39.106.59.225:9411/api/v2/spans"


func InitZipkinTracer() (*zipkin.Tracer, error) {
	// reporter 向 Zipkin 后端上报数据
	reporter := zipkinhttp.NewReporter(zipkinUrl)

	endpoint, err := zipkin.NewEndpoint("gin-server", "localhost:8080")
	if err != nil {
		return nil, err
	}

	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	return tracer, nil
}


func TracingMiddleware(zipkinTracer *zipkin.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中提取 trace 信息
		sc := zipkinTracer.Extract(b3.ExtractHTTP(c.Request))

		// 为当前请求创建新的 span
		span := zipkinTracer.StartSpan(
			c.FullPath(),
			zipkin.Parent(sc),
			zipkin.Kind(model.Server),
		)
		defer span.Finish()

		// 添加 traceId 到 context 中，供业务日志使用
		c.Set("traceId", span.Context().TraceID.String())
		c.Set("spanId", span.Context().ID.String())

		c.Next()
	}
}
