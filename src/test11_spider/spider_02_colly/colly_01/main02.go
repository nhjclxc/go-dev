package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main02() {

	// 把百度首页最上面一行按钮对应的链接输出出来

	c := colly.NewCollector()

	c.OnHTML("div.s-top-left", func(e *colly.HTMLElement) {
		fmt.Println("div.s-top-left ", e.Name)
		e.ForEach("a", func(i int, ee *colly.HTMLElement) {
			href := ee.Attr("href")
			fmt.Println(i, ee.Text, href)
		})
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("a[href] ", e.Name, e.Text, e.Attr("href"))
	})

	c.Visit("https://baidu.com")

}
