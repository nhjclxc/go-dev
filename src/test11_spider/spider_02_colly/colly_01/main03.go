package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// https://github.com/trending

	// 拉取 github trending 页面每一项的数据

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		// e.ChildAttr 是用来获取子元素的属性值的，比如 href 或 src。
		repoName := e.ChildAttr("h2 a", "href")
		// e.ChildText返回的是该子元素的文本内容
		shortIntroduction := e.ChildText("p")
		fmt.Println("article.Box-row ", repoName, shortIntroduction)
	})

	c.Visit("https://github.com/trending")
}
