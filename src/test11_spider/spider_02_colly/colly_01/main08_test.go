package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/gocolly/redisstorage"
	"testing"
)

// 分布式爬虫
// https://m.w3cschool.cn/colly/colly-distributed.html

// 方案 1：代理池（轻量级分布式）
// 只有一台机器，但想“假装”成很多机器。
func TestMain0801(t *testing.T) {

	c := colly.NewCollector()

	// 把代理地址换成你自己的
	p, _ := proxy.RoundRobinProxySwitcher(
		"http://127.0.0.1:8080",
		"socks5://127.0.0.1:1080",
		"http://user:pass@ip:port", // 带账号密码示例
	)
	c.SetProxyFunc(p)

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("用代理", r.Request.ProxyURL, "抓到", len(r.Body), "字节")
	})

	c.Visit("https://www.w3cschool.cn/")
}

// 方案 2：多节点爬虫（真·分布式）
// 多台云服务器 / 本地多台电脑，需要统一调度。
func TestMain0802(t *testing.T) {

	// redisstorage不是colly库实现的，必须在安装一个库：go get github.com/gocolly/redisstorage
	// 1. 连接 Redis
	s := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "httpbin_test",
	}
	s.Init()
	defer s.Client.Close()

	// 2、创建收集器
	c := colly.NewCollector()

	// 会自动将cookie数据和已访问的链接自动存入reids
	// 因此，要想实现多节点分布式爬取，只需将当前爬虫在多个节点启动就可以做到了
	// colly会通过这个redis去检查是否已访问
	// 重启程序 → 自动跳过已爬 URL。
	// 开 10 台云主机 → 共享同一份 Redis，0 重复抓取。
	c.SetStorage(s)

	// 3、请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r)
	})

	// 3. 正常写业务逻辑
	c.OnHTML("title", func(e *colly.HTMLElement) {
		println("标题：", e.Text)
	})

	c.Visit("https://www.w3cschool.cn/")

}
