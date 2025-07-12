package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// 防止前端重复提交
func main() {

	// 创建redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认使用DB 0
	})


	// 创建一个默认的路由引擎
	router := gin.Default()

	// 获取uuid接口
	router.GET("/getUuid", func(ctx *gin.Context) {
		// 在 getUuid 接口里面返回一个uuid给前端，
		// 当前端点击post或put类型接口要求防止重复提交的接口的时候带上这个uuid

		reqUUid := uuid.New().String()
		// 5分钟过期，key 不存在时才设置（SETNX）
		nx := redisClient.SetNX(context.Background(), reqUUid, 1, time.Minute*5)

		if nx.Err() != nil || !nx.Val() {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "获取UUID失败",
			})
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": nil,
				"msg": "操作成功",
				"data": reqUUid,
			},
		)
	})

	router.GET("/insert", func(ctx *gin.Context) {
		reqUUid := ctx.Query("reqUUid")
		data := ctx.Query("data")


		// 拿到 p.ReqUUid 去缓存里面看看是否存在，存在则说明是第一次提交该请求，请求放行
		// 若不存在，则说明已经被前面的请求删除了，当前请求拦截

		exists, err := redisClient.Del(context.Background(), reqUUid).Result()
		if err != nil || exists == 0 {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"code":  http.StatusOK,
					"error": "state == 0 请勿重复提交",
					"msg": "state == 0 请勿重复提交",
					"data": nil,
				},
			)
			return
		}

		// 执行业务逻辑


		// 处理请求
		fmt.Println("Goroutine 1: 获取到锁，开始执行" + data)
		time.Sleep(2 * time.Second) // 模拟长时间持有锁

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "处理完成",
				"msg": "处理完成",
				"data": "处理完成",
			},
		)

	})

	fmt.Println("服务启动！！！")
	//启动端口监听
	router.Run(":8090")
}
