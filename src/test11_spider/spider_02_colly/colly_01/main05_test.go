package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/proxy"
	"github.com/gocolly/colly/storage"
	"testing"
)

// Distributed scraping 分布式爬取
// https://go-colly.org/docs/best_practices/distributed/
func TestMain05(t *testing.T) {

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
		colly.Debugger(&debug.LogDebugger{}),
	)

	// 加入代理实现分布式
	p, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
		"http://127.0.0.1:8080",
	)
	if err == nil {
		c.SetProxyFunc(p)
	}
	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		// e.ChildAttr 是用来获取子元素的属性值的，比如 href 或 src。
		repoName := e.ChildAttr("h2 a", "href")
		fmt.Println("article.Box-row ", repoName)
	})

	c.Visit("https://github.com/trending")

	s := storage.InMemoryStorage{}
	s.Init()

	c.SetStorage(&s)

}
