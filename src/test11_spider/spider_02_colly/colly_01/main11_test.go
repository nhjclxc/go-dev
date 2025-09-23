package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"testing"
	"time"
)

// https://m.w3cschool.cn/colly/colly-extensions.html
//
// 官方扩展 = 一行代码 + 零配置，立刻拥有随机 UA、自动 Referer、限速等实用功能。本文带你 3 分钟全部学会！
// 目前常用 4 件套：
// 扩展名				作用						一行代码
// RandomUserAgent	每次请求随机 UA，防封			extensions.RandomUserAgent(c)
// Referer			自动把上一页 URL 设为 Referer	extensions.Referer(c)
// URLLengthFilter	过滤超长 URL					extensions.URLLengthFilter(c, 2083)
// MaxDepth			限制爬取深度					extensions.MaxDepth(c, 3)
func TestMain11(t *testing.T) {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	extensions.URLLengthFilter(c, 2038)
	c.MaxDepth = 3

	// 使用
	TimerExtension(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(111)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(222)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(333)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("UA=%s  Referer=%s\n",
			r.Request.Headers.Get("User-Agent"),
			r.Request.Headers.Get("Referer"))
	})

	c.OnHTML("", func(e *colly.HTMLElement) {

	})

	c.Visit("https://www.w3cschool.cn/")

}

// 示例：打印每次请求耗时
func TimerExtension(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("start", time.Now())
	})
	c.OnResponse(func(r *colly.Response) {
		start := r.Ctx.GetAny("start").(time.Time)
		fmt.Printf("耗时 %v → %s \n", time.Since(start), r.Request.URL)
	})
}
