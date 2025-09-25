package colly_03

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
)

// 爬取 Instagram上【https://www.instagram.com/tan__suan/】的所有开放照片

func TestMain06(t *testing.T) {

	c := colly.NewCollector()

	//c.OnHTML("div.x1lliihq.x1n2onr6.xh8yej3.x4gyw5p.x1mpyi22.x1j53mea", func(e *colly.HTMLElement) {
	c.OnHTML("div.x1lliihq", func(e *colly.HTMLElement) {
		fmt.Println("找到了")
	})

	c.Visit("https://www.instagram.com/tan__suan/")

}
