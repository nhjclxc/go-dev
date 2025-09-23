package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// 爬取w3c网页里面的网页，深度=2
// https://m.w3cschool.cn/colly/colly-vofy30nk.html
func TestMain06(t *testing.T) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		// e.Request.Depth 获取当前请求的深度：
		fmt.Printf("Depth=%d, 发现链接🔗=%s\n", e.Request.Depth, href)
		// 这里 e.Request.Visit(href) 会把链接交给 同一个 Collector (c) 再去请求。
		// Collector 的所有 OnHTML/OnResponse 回调都会对新页面生效。
		// 因此，href访问到到连接仍然会调用当前这个c.OnHTML的回调方法
		e.Request.Visit(href)
	})

	// 爬取深度设置为2，避免无限爬取
	c.MaxDepth = 2
	c.Visit("https://www.w3cschool.cn/")

}

// colly访问api数据
func TestMain0602(t *testing.T) {

	// 创建默认收集器
	c := colly.NewCollector(
		// 不同的网站对访问者的身份有不同的要求。有时候，我们需要让爬虫机器人伪装成不同的浏览器，这就需要修改用户代理（User-Agent）。代码示例：
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36"),
	)

	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)

		// 二）设置请求头
		//有时候，目标网站会检查请求头来判断是否是真实的浏览器访问。为了更好地模拟浏览器行为，我们可以设置请求头。代码示例：
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)") // 修改value的值可以做到动态修改代理值的目的
		r.Headers.Set("Referer", "https://example.com")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	})

	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("收到响应：", string(r.Body))
	})

	// （一）设置请求延迟
	//在爬取多个页面时，频繁的请求可能会对目标网站的服务器造成压力。
	//为了避免这种情况，我们可以设置请求延迟，让 Colly 在发送请求之间等待一段时间。代码示例：
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",             // 对所有域名生效
		Parallelism: 2,               // 同时发送 2 个请求
		Delay:       1 * time.Second, // 每个请求之间间隔 1 秒
	})

	// 这样设置后，目标网站就会认为是常见的浏览器在访问，而不是一个简单的爬虫程序。

	// 访问 API
	c.Visit("https://api.example.com/data")
}

// （一）设置 HTTP 超时时间
func TestMain0603(t *testing.T) {
	c := colly.NewCollector()

	// colly内部有一个默认的http请求配置，
	// 但是也可以通过一下方法来修改默认的http请求配置

	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 设置连接超时时间为 30 秒
			KeepAlive: 30 * time.Second, // 设置连接保持活动时间为 30 秒
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second, // 设置 TLS 握手超时时间为 10 秒
	})
}

// 给colly的日志输出
func TestMain0605(t *testing.T) {

	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
	)

	// 创建一个日志文件
	file, _ := os.Create("TestMain0605.log")

	log.SetOutput(file)

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
			fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
			fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
			fmt.Printf("\tNumGC = %v\n", m.NumGC)
			time.Sleep(3 * time.Second)
		}
	}()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//log.Printf("[%s]链接🔗：%s \n", strings.ReplaceAll(e.Text, "\\s+", ""), e.Attr("href"))
		log.Printf("[%s]链接🔗：%s \n", strings.TrimSpace(e.Text), e.Attr("href"))

		time.Sleep(1 * time.Second)
	})

	c.Visit("https://m.w3cschool.cn/")

}
