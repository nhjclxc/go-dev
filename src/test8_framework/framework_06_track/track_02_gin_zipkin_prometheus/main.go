package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"track_02_gin_zipkin_prometheus/middleware"
)

// Gin + OpenTelemetry + Zipkin 实现链路追踪
//go get go.opentelemetry.io/otel
//go get go.opentelemetry.io/otel/sdk
//go get go.opentelemetry.io/otel/exporters/zipkin
//go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin


// 启动一个zipkin实例
// docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/openzipkin/zipkin:latest
// docker run -d -p 9411:9411 swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/openzipkin/zipkin:latest
// http://39.106.59.225:9411/zipkin/

// go run main.go

// 发起一个请求，看zipkin UI的请求记录

// Prometheus 配置



// 继续增加zap日志配合
// go get go.uber.org/zap




func main() {

	router := gin.Default()

	// 启用跨域支持
	router.Use(cors.Default())


	// 初始化 Zipkin Tracer
	tracer, err := middleware.InitZipkinTracer()
	if err != nil {
		log.Fatalf("无法初始化Zipkin: %v", err)
	}

	// 启用中间件
	middleware.PrometheusMiddleware(router)
	router.Use(middleware.TracingMiddleware(tracer))

	// 示例接口
	router.GET("/hello", func(c *gin.Context) {
		traceId, _ := c.Get("traceId")
		spanId, _ := c.Get("spanId")

		log.Printf("处理请求：traceId=%v, spanId=%v", traceId, spanId)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
			"traceId": traceId,
			"spanId":  spanId,
		})
	})



	log.Println("服务启动于 :8866")
	if err := router.Run(":8866"); err != nil {
		log.Fatal(err)
	}


}
