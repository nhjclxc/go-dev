package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"go_base_project/pkg/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		clientIP := c.ClientIP()

		logger.Info("HTTP 请求",
			"status", status,
			"method", method,
			"path", path,
			"query", query,
			"ip", clientIP,
			"latency", latency,
		)
	}
}
