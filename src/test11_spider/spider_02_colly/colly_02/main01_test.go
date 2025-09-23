package colly_02

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
)

// 访问 【https://m.w3cschool.cn/】发现链接就进去接着访问
func TestMain01(t *testing.T) {

	c := colly.NewCollector(
		colly.AllowedDomains("m.w3cschool.cn"),
		colly.Async(true),
	)
	c.MaxDepth = 2

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// link有可能是没有域名的地址，所以经过下面方法之后把域名加上返回
		// 自动补全绝对路径后再访问
		absoluteURL := e.Request.AbsoluteURL(link) // 把相对 URL 转换成绝对 URL。
		fmt.Printf("发现链接：%s → %s， absoluteURL = %s \n", e.Text, link, absoluteURL)
		c.Visit(absoluteURL)
	})

	c.Visit("https://m.w3cschool.cn/")
	c.Wait()

}
