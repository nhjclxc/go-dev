package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 使用 Redis 实现分布式限流（适合多实例部署）
///     具体实现方式：固定窗口计数法（Fixed Window Counter），原理：以 Redis 的 key 作为窗口标记，在单位时间内请求次数超限就拒绝。
//       也就是redis的自增Incr命令

// go get github.com/redis/go-redis/v9





func main() {

	router := gin.Default()

	// 启用跨域支持
	router.Use(cors.Default())

	// 每个 IP 每分钟最多 10 次
	router.Use(RedisRateLimiter(10, 60))



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
