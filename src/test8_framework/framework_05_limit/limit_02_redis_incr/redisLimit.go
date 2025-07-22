package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // 根据你的配置修改
})

func RedisRateLimiter(maxRequests int, windowSeconds int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 以 userId 或 IP 作为限流 key
		clientID := c.ClientIP() // 可换成 userID、token 等
		key := fmt.Sprintf("rate_limit:%s", clientID)

		// 使用 INCR 并设置过期时间（固定窗口计数法）
		//count, err := rdb.Incr(ctx, key).Result()
		count, err := rdb.Incr(key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "redis error"})
			return
		}

		if count == 1 {
			// 第一次请求，设置过期时间
			//rdb.Expire(ctx, key, time.Duration(windowSeconds)*time.Second)
			rdb.Expire(key, time.Duration(windowSeconds)*time.Second)
		}

		if count > int64(maxRequests) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"msg": "Too many requests, please try again later",
			})
			return
		}

		c.Next()
	}
}
