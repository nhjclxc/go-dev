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


// 使用gin的中间件实现通用的防止重复提交
func main() {
	// 初始化 Redis 客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 测试 Redis 连接
	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic("Redis连接失败: " + err.Error())
	}

	// Gin 路由
	r := gin.Default()

	// 获取 UUID 的接口
	r.GET("/getUuid", func(c *gin.Context) {
		reqUUid := fmt.Sprintf("form_token:%d", time.Now().UnixNano())
		redisClient.SetNX(ctx, reqUUid, 1, time.Minute*5)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": reqUUid,
			"msg":  "获取成功",
		})
	})

	// 应用中间件，仅用于 insert 路由
	r.GET("/insert", IdempotencyMiddleware(redisClient), func(c *gin.Context) {
		data := c.PostForm("data")

		// 模拟业务处理
		fmt.Println("处理数据:", data)
		time.Sleep(2 * time.Second)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "处理完成",
		})
	})




	fmt.Println("服务启动！！！")
	r.Run(":8090")
}
