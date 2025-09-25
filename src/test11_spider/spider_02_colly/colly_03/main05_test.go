package colly_03

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"testing"
)

// 爬取【https://groups.google.com/g/bojug】里面的所有邮件
func TestMain05(t *testing.T) {

	listC := colly.NewCollector()
	detailC := colly.NewCollector(
		colly.Async(true))

	mails := make([]Mail, 0)

	listC.OnHTML("div.yhgbKd", func(e *colly.HTMLElement) {
		//author := e.ChildText("span.z0zUgf")
		//title := e.ChildText("span.o1DPKc")
		linkTemp := e.ChildAttr("a[href]", "href")
		link := e.Request.AbsoluteURL(linkTemp)
		//date := e.ChildText("div.kOkyJc div.tRlaM")
		//fmt.Println("找到了 yhgbKd", author, title, link, date)

		ctx := colly.NewContext()
		ctx.Put("link", link)
		detailC.Request("GET", link, nil, ctx, nil)

	})

	// 翻页
	listC.OnHTML("body > a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		fmt.Println("翻页", e.Attr("href"))
	})

	detailC.OnHTML("div.D1OdOb", func(e *colly.HTMLElement) {
		author := e.ChildText("div.LgTNRd h3.s1f8Zd")
		title := e.ChildText("html-blob")
		date := e.ChildText("div.ELCJ4d span.zX2W9c")
		message := e.ChildText("div.ptW7te")
		link := e.Request.Ctx.Get("link")
		//fmt.Printf("找到了 author = %s, title = %s, link = %s, date = %v, message = %v \n", author, title, link, date, message)

		mails = append(mails, Mail{
			Title:   title,
			Link:    link,
			Author:  author,
			Date:    date,
			Message: message,
		})
	})

	listC.Visit("https://groups.google.com/g/bojug")
	detailC.Wait()

	data, err := json.MarshalIndent(mails, "", "")
	if err != nil {
		return
	}
	err = os.WriteFile("mails.json", data, 0644)
	if err != nil {
		return
	}

}

// Mail 单封邮件结构
type Mail struct {
	Title   string `json:"标题"`
	Link    string `json:"链接"`
	Author  string `json:"作者"`
	Date    string `json:"日期"`
	Message string `json:"正文"`
}
