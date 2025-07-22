下面给你一个基于 **Redis + Lua 脚本实现令牌桶限流** 的完整示例，包括：

* **Lua 脚本**（原子操作，支持令牌桶限流）
* **Go 代码**调用 Lua 脚本实现分布式限流中间件（基于 Gin）

---

# 1. Lua 脚本（令牌桶限流）

这个脚本实现的逻辑是：

* 根据 key 获取当前令牌数量和时间戳
* 计算新增令牌数，令牌数量不能超过桶容量
* 判断是否有令牌可用（令牌数 > 0 则允许请求并扣除一个令牌）
* 返回是否允许

```lua
-- KEYS[1] - 限流 key，比如用户IP或ID
-- ARGV[1] - 令牌桶容量（最大令牌数）
-- ARGV[2] - 令牌生成速率（每秒产生多少令牌，float）
-- ARGV[3] - 当前时间戳（秒）
-- 返回 1 允许访问，0 拒绝访问

local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- redis中存储结构：hash表，field "tokens" 存令牌数，field "timestamp" 存上次更新时间
local data = redis.call("HMGET", key, "tokens", "timestamp")
local tokens = tonumber(data[1])
local last_timestamp = tonumber(data[2])

if tokens == nil then
    tokens = capacity
    last_timestamp = now
end

-- 计算时间差，计算增加令牌数
local delta = math.max(0, now - last_timestamp)
local add_tokens = delta * rate
tokens = math.min(capacity, tokens + add_tokens)

local allowed = 0
if tokens >= 1 then
    tokens = tokens - 1
    allowed = 1
end

-- 更新 redis
redis.call("HMSET", key, "tokens", tokens, "timestamp", now)
redis.call("EXPIRE", key, 3600) -- 设置1小时过期，防止key无限增长

return allowed
```

---

# 2. Go 代码调用 Lua 脚本限流

```go
package main

import (
	"context"
	"fmt"
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
func RedisRateLimiterMiddleware(capacity int, rate float64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate_limiter:" + c.ClientIP()
		now := time.Now().Unix()

		// 执行 Lua 脚本
		allowed, err := rdb.Eval(ctx, luaScript, []string{key},
			capacity, rate, now).Int()
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

func main() {
	r := gin.Default()

	// 令牌桶容量10，生成速率5/s
	r.Use(RedisRateLimiterMiddleware(10, 5))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("server start error:", err)
	}
}
```

---

# 3. 使用说明

* **capacity=10** 表示桶最大容量10个令牌（支持最多10次突发请求）
* **rate=5** 表示每秒产生5个令牌，持续补充桶内令牌
* 如果令牌不够了，接口返回 HTTP 429

---

# 4. 总结

* 该方案适合**多实例分布式环境**，共享 Redis 实现全局限流；
* 利用 Lua 脚本保证 Redis 操作原子性，防止竞态条件；
* 可根据业务调整 `capacity` 和 `rate` 参数，灵活控制限流策略。

---

如果你需要，我也能帮你写一个 **完整测试脚本** 或者 **docker-compose 集成 Redis + Gin 服务示例**。
需要告诉我！
