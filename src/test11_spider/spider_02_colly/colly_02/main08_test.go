package colly_02

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"testing"
)

// Colly URL 过滤器：精准控制爬取路径
// 只想爬指定路径？用正则表达式给爬虫戴上“紧箍咒”，多余链接统统跳过！
func TestMain08(t *testing.T) {
	// 1. 创建收集器，只匹配课程相关 URL
	c := colly.NewCollector(
		colly.URLFilters(
			// 允许：首页 或 /course 开头
			regexp.MustCompile(`https://www\.w3cschool\.cn/?$`),
			regexp.MustCompile(`https://www\.w3cschool\.cn/course.*`),
		),
	)

	// 2. 发现链接就打印并继续访问
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("发现课程链接：%q → %s\n", e.Text, link)
		// 自动补全绝对路径
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// 3. 打印访问日志
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("正在访问：", r.URL.String())
	})

	// 4. 从课程首页开始
	c.Visit("https://www.w3cschool.cn/course")
}
