package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

const MAX_CONNECT = 3

type DBConnection struct {
	ID int
}

func main() {
	r := gin.Default()

	// 用 channel 模拟连接池
	pool := make(chan *DBConnection, MAX_CONNECT)
	for i := 1; i <= MAX_CONNECT; i++ {
		pool <- &DBConnection{ID: i}
	}

	r.GET("/conn", func(ctx *gin.Context) {
		select {
		case conn := <-pool: // 获取连接
			defer func() { pool <- conn }() // 处理完归还

			fmt.Printf("使用连接 %d\n", conn.ID)
			time.Sleep(5 * time.Second) // 模拟操作

			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  fmt.Sprintf("操作完成，连接 %d 已归还", conn.ID),
			})
		default: // 没有空闲连接
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "连接池已满，请稍后再试",
			})
		}
	})

	r.Run(":8080")
}
