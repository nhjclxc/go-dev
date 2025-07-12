package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func IdempotencyMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUUid := c.PostForm("reqUUid") // 或 c.Query("reqUUid")，按你实际方式选择
		if reqUUid == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "缺少 reqUUid 参数",
			})
			return
		}

		key := "form_token:" + reqUUid

		// 原子删除并判断是否存在
		deleted, err := rdb.Del(context.Background(), key).Result()
		if err != nil || deleted == 0 {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"code": 409,
				"msg":  "请勿重复提交",
			})
			return
		}

		c.Next() // 放行
	}
}
