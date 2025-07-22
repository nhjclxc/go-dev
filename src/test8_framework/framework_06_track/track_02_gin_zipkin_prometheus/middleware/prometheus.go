package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zsais/go-gin-prometheus"
)

func PrometheusMiddleware(router *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	// 也可以手动设置指标端点路径（默认是 /metrics）
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}