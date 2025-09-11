package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ConnPool 模拟连接池 ConnPool ：用一个 buffered channel 实现
type ConnPool struct {
	pool chan struct{}
}

// NewConnPool 初始化连接池，size 表示最大连接数
func NewConnPool(size int) *ConnPool {
	return &ConnPool{
		pool: make(chan struct{}, size),
	}
}

// Acquire 获取连接（占用资源）
func (c *ConnPool) Acquire() bool {
	select {
	case c.pool <- struct{}{}:
		return true
	default:
		return false // 池子满了
	}
}

// Release 释放连接（归还资源）
func (c *ConnPool) Release() {
	select {
	case <-c.pool:
	default:
	}
}

// 用gin来模拟一个链接池，没发送一次请求，链接池资源减1，接口执行完毕之后接口恢复
func main() {
	r := gin.Default()

	// 初始化一个最大 3 个连接的池子
	connPool := NewConnPool(3)

	r.GET("/do", func(c *gin.Context) {
		if !connPool.Acquire() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "连接池已满，请稍后再试",
			})
			return
		}
		defer connPool.Release() // 请求结束后归还资源

		// 模拟耗时操作
		time.Sleep(2 * time.Second)

		c.JSON(http.StatusOK, gin.H{
			"message": "处理成功",
			"time":    fmt.Sprintf("%s", time.Now().Format("15:04:05")),
		})
	})

	r.Run(":8080")
}
