package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Redis + Lua 脚本实现令牌桶限流
//		Lua 脚本（原子操作，支持令牌桶限流）
// 		Go 代码调用 Lua 脚本实现分布式限流中间件（基于 Gin）

// go get github.com/redis/go-redis/v9


/*
1. Lua 脚本（令牌桶限流）
这个脚本实现的逻辑是：
	根据 key 获取当前令牌数量和时间戳
	计算新增令牌数，令牌数量不能超过桶容量
	判断是否有令牌可用（令牌数 > 0 则允许请求并扣除一个令牌）
	返回是否允许


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

 */



// 当前方案适合多实例分布式环境，共享 Redis 实现全局限流；
// 利用 Lua 脚本保证 Redis 操作原子性，防止竞态条件；
// 可根据业务调整 capacity 和 rate 参数，灵活控制限流策略。
func main() {

	router := gin.Default()

	// 启用跨域支持
	router.Use(cors.Default())

	// 每个 IP 每分钟最多 10 次
	router.Use(RedisLuaRateLimiterMiddleware(10, 60))



	// http://127.0.0.1:8080/ping
	// for /L %i in (1,1,20) do curl -s -w "%%{http_code}\n" -o NUL http://localhost:8080/ping
	// 测试接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	log.Println("服务启动于 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
