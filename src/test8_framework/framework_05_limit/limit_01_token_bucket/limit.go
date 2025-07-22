package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)


//🌟 示例代码：每个 IP 每秒最多 5 个请求，最多突发 10 个

// 存储每个IP对应的限流器
var limiterMap = make(map[string]*rate.Limiter)
var mutex sync.Mutex

// 获取某个IP对应的限流器（如果没有就新建一个）
func getLimiterForIP(ip string) *rate.Limiter {
	mutex.Lock()
	defer mutex.Unlock()

	limiter, exists := limiterMap[ip]
	if !exists {
		// 每秒产生5个令牌，最多允许突发10个
		limiter = rate.NewLimiter(5, 10)
		limiterMap[ip] = limiter
	}
	return limiter
}

// Gin中间件：限流器
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiterForIP(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": 429,
				"msg":  "Too many requests. Please try again later.",
			})
			return
		}

		c.Next()
	}
}

/*
Limiter 的“请求次数 +1”是通过 Allow() 成功获取令牌来 隐式完成的，而非显式统计请求次数。

limiter.Allow() 该方法的工作流程如下：
1. 检查当前时间是否允许“领取”一个令牌。
2. 如果桶中有令牌（不小于 1 个）：
	允许请求（返回 true）。
	同时扣除一个令牌。
3. 如果桶空了（令牌不足）：
	拒绝请求（返回 false）。


*/