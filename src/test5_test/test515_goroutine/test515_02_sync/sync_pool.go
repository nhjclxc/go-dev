package main

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"sync"
//	"time"
//)
//
//// "sync.Pool"的使用
//
///*
//1️⃣ 什么是 sync.Pool
//sync.Pool 是 Go 标准库 sync 包提供的一个 对象池，用于 复用临时对象，减少内存分配和垃圾回收压力。
//
//它的特点：
//
//	缓存临时对象，避免频繁创建和销毁。
//	并发安全，多 goroutine 可以同时 Get / Put。
//	非严格保证：对象池里的对象可能随时被垃圾回收清掉，不保证一定能复用。
//
//2️⃣ 典型用途
//sync.Pool 常用于：
//
//	减少频繁分配的临时对象：例如缓冲区、随机数生成器、临时 struct 等。
//	并发场景：多 goroutine 需要用同一种对象，但每次都新建太慢。
//	提升性能：避免 GC 压力，特别是短生命周期、大量临时对象。
//*/
//
////sync.Pool模拟一个数据库连接池的功能，不需要真正的链接数据库，模拟链接操作
//
//const MAX_CONNECT = 3
//
//var conns = make([]bool, MAX_CONNECT+1)
//
//type DBConnection struct {
//	ID int
//}
//
//// dbPool 用 sync.Pool 模拟连接池
//var dbPool = sync.Pool{
//	New: func() any {
//		var connId int = 0
//		for i, flag := range conns {
//			if i > 0 && !flag {
//				connId = i
//				conns[i] = true
//				break
//			}
//		}
//		if connId == 0 {
//			return nil
//		}
//		fmt.Println("connId", connId)
//
//		// 每次创建新的连接（模拟连接ID自增）
//		return &DBConnection{ID: connId}
//	},
//}
//
//func conn() (*DBConnection, error) {
//	conn := dbPool.Get()
//	if conn == nil {
//		return nil, fmt.Errorf("连接数已满！！！")
//	}
//	return conn.(*DBConnection), nil
//}
//
//func main() {
//
//	r := gin.Default()
//
//	r.GET("/conn", func(ctx *gin.Context) {
//
//		connection, err := conn()
//		if err != nil {
//			ctx.JSON(200, gin.H{
//				"code": 500,
//				"msg":  err.Error(),
//			})
//			return
//		}
//
//		go func() {
//			defer func() {
//				dbPool.Put(connection)
//				conns[connection.ID] = false
//			}()
//			time.Sleep(10 * time.Second)
//		}()
//
//		ctx.JSON(200, gin.H{
//			"code": 200,
//			"msg":  "操作成功",
//		})
//
//	})
//
//	r.Run(":8080")
//}
