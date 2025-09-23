package colly_02

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
	"testing"
	"time"
)

// Colly 代理轮询：自动切换 IP 防封实战
// 单 IP 被封？用 Colly 官方 proxy.RoundRobinProxySwitcher，一行代码实现多代理轮询，轻松伪装成“千军万马”。

func TestMain05(t *testing.T) {
	// 创建收集器，允许重复访问同一 URL
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.AllowedDomains("www.w3cschool.cn"),
	)

	// 1. 设置代理池（socks5/http 均可）
	rp, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
		"http://127.0.0.1:8080", // 也可混用 HTTP 代理
	)
	if err != nil {
		log.Fatal("代理设置失败：", err)
	}
	c.SetProxyFunc(rp)

	// 2. 打印每次使用的代理和返回内容
	c.OnResponse(func(r *colly.Response) {
		log.Printf("代理：%s | 返回长度：%d 字节 | URL：%s",
			r.Request.ProxyURL, len(r.Body), r.Request.URL)
	})

	// 3. 连续访问 5 次，观察 IP 轮换
	for i := 0; i < 5; i++ {
		c.Visit("https://www.w3cschool.cn/")
		time.Sleep(1 * time.Second) // 避免太快
	}

}
