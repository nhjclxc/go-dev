package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"testing"
)

// Colly 多收集器：让爬虫分工协作

// 目标：获取编程狮[https://m.w3cschool.cn/]，每一个课程的收藏数量【先访问列表，接着要访问列表详细数据】

// 不使用多收集器的实现
func TestMain0901(t *testing.T) {

	c := colly.NewCollector()

	baseUrl := "https://m.w3cschool.cn"
	titleMap := make(map[string]string)

	// 匹配 <a> 标签，并且 class 包含 item-content
	// 匹配 a.item-content 内部任意层级的 class 为 item-title 的元素
	//c.OnHTML("a.item-content .item-title", func(e *colly.HTMLElement) {
	c.OnHTML("a.item-content", func(e *colly.HTMLElement) {
		// 1、获取列表
		href := e.Attr("href")
		title := e.ChildText(".item-title")
		titleMap[href] = title
		//fmt.Printf("a.item-content : %s -> [%s]\n", title, href)

		// 通过e.DOM遍历出 a.item-content 下的所有 .item-title 类的元素
		//e.DOM.Find(".item-title").Each(func(_ int, s *goquery.Selection) {
		//	text := s.Text()
		//	fmt.Println("标题:", text)
		//})

		// 2、访问详细获取数量
		reqUrl := baseUrl + href
		e.Request.Visit(reqUrl)

	})
	c.OnHTML("div.card-footer", func(e *colly.HTMLElement) {
		uri := strings.ReplaceAll(e.Request.URL.Path, baseUrl, "")

		//fmt.Println(e.Request.URL.Path)
		fmt.Printf("title: %s, %s \n", titleMap[uri], e.Text)
	})

	c.Visit("https://m.w3cschool.cn/")
}

// 使用多收集器的实现
func TestMain0902(t *testing.T) {

	baseUrl := "https://m.w3cschool.cn"
	listC := colly.NewCollector(colly.AllowedDomains("m.w3cschool.cn"))

	detailC := colly.NewCollector(colly.AllowedDomains("m.w3cschool.cn"))

	listC.OnHTML("a.item-content", func(e *colly.HTMLElement) {
		// 1、获取列表
		href := e.Attr("href")
		title := e.ChildText(".item-title")

		// 使用 Context 传递上下文
		collyCtx := colly.NewContext()
		collyCtx.Put("title", title)

		// 2、访问详细获取数量
		reqUrl := baseUrl + href
		// 通知detailC去向指定链接发起请求
		detailC.Request("GET", reqUrl, nil, collyCtx, nil)
	})

	detailC.OnHTML("div.card-footer", func(e *colly.HTMLElement) {
		fmt.Printf("title: %s, %s \n", e.Request.Ctx.Get("title"), e.Text)
	})

	listC.Visit("https://m.w3cschool.cn/")
}
