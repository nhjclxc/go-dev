package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/redisstorage"
	"net/http"
	"testing"
	"time"
)

// colly 高并发配置

func TestMain10(t *testing.T) {
	//一、默认配置的 3 个瓶颈
	//瓶颈	  		 	说明							风险
	//纯内存存储			Cookie + 已爬 URL 全放内存	进程挂掉 = 数据清零
	//同步阻塞			一次只能发一个请求				速度瓶颈
	//长连接 Keep-Alive	打开文件句柄不释放				句柄耗尽，程序崩溃

	//1️⃣ 换持久化存储：断电也不怕
	//把默认内存存储换成 Redis（或 BoltDB/SQLite 等），已爬 URL 实时落盘。

	rs := redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "colly:xxx",
	}
	rs.Init()
	defer rs.Client.Close()

	// 2️⃣ 开启异步：并发飞起
	//默认是同步阻塞，改异步后回调不再堆栈爆炸。
	c := colly.NewCollector(
		colly.Async(true), // 关键开关
	)
	c.SetStorage(&rs)
	//并发数控制：加 c.Limit(&colly.LimitRule{Parallelism: 16}) 防止把目标站打挂。
	//内存保护：c.SetRequestTimeout(30 * time.Second) 避免慢请求堆积。
	c.Limit(&colly.LimitRule{
		Parallelism: 32,
		Delay:       200 * time.Millisecond,
	})

	//3️
	//3️⃣ 关闭 K⃣ee⃣p⃣-⃣Alive：句柄不再爆表
	//长连接会占用大量文件描述符，百万级任务建议关闭或缩短。
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true, // 短连接，用完即关
		MaxIdleConns:      100,  // 如仍需复用，可限制数量
	})

	c.Visit("")

	// 开启Async之后，一定要等待所有gor关闭，不然有些gor会没有执行完
	c.Wait() // 等待所有 goroutine 结束

}
