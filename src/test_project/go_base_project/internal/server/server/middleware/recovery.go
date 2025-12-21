package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_base_project/pkg/logger"
	"net/http"
	"runtime/debug"
	"strings"
)

// Recovery 恢复中间件，捕获 panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("发生 panic", "error", err)
				logger.Error("stack trace:\n\n", projectStack(fmt.Sprintf("%v", err)))

				c.JSON(http.StatusOK, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "服务器内部错误",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func projectStack(err string) string {
	fullStack := debug.Stack()
	lines := strings.Split(string(fullStack), "\n")

	var filtered []string
	filtered = append(filtered, "panic reason:"+err)
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// 只保留包含你项目路径的堆栈
		if strings.Contains(line, "android-traffic-monitor") && strings.Contains(line, ".go:") {
			// 上一行一般是函数签名，保留
			if i > 0 {
				filtered = append(filtered, lines[i-1])
			}
			filtered = append(filtered, line)
		}
	}

	return strings.Join(filtered, "\n")
}
