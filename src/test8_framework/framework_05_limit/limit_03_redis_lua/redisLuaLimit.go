package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

var ctx = context.Background()

// 初始化 Redis 客户端
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Lua 脚本字符串
const luaScript = `
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local data = redis.call("HMGET", key, "tokens", "timestamp")
local tokens = tonumber(data[1])
local last_timestamp = tonumber(data[2])

if tokens == nil then
    tokens = capacity
    last_timestamp = now
end

local delta = math.max(0, now - last_timestamp)
local add_tokens = delta * rate
tokens = math.min(capacity, tokens + add_tokens)

local allowed = 0
if tokens >= 1 then
    tokens = tokens - 1
    allowed = 1
end

redis.call("HMSET", key, "tokens", tokens, "timestamp", now)
redis.call("EXPIRE", key, 3600)

return allowed
`

// RateLimiter 中间件参数
// capacity=10 表示桶最大容量10个令牌（支持最多10次突发请求）
// rate=5 表示每秒产生5个令牌，持续补充桶内令牌
// 如果令牌不够了，接口返回 HTTP 429
func RedisLuaRateLimiterMiddleware(capacity int, rate float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate_limiter:" + c.ClientIP()
		now := time.Now().Unix()

		// 执行 Lua 脚本
		allowed, err := rdb.Eval(ctx, luaScript, []string{key}, capacity, rate, now).Int()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
			return
		}

		if allowed == 1 {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try later.",
			})
		}
	}
}
