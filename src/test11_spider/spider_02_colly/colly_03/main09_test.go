package colly_03

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"testing"
)

// 将【https://ssr1.scrape.center/】里面的电影数据保存为json

func Test09(t *testing.T) {

	c := colly.NewCollector()

	c.OnHTML("div.el-row", func(e *colly.HTMLElement) {
		e.ForEach("div", func(i int, e2 *colly.HTMLElement) {
			if i == 0 {
				return
			}
			var title, categories string
			if i == 1 {
				title = e2.ChildText("h2")
				fmt.Println("title", title)
				cs := make([]string, 0)
				e2.ForEach("button span", func(i int, e3 *colly.HTMLElement) {
					cs = append(cs, e2.ChildText("button span"))
					fmt.Println("11", e2.ChildText("button span"))
				})
				categories = strings.Join(cs, ",")
				fmt.Println("cs", cs)
				fmt.Println("categories", categories)
			}

		})
	})

	c.Visit("https://ssr1.scrape.center/")

}
