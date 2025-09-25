package colly_03

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
)

// 爬取中国大学MOOC下面的所有AI课程及其讲师名字
//AI课程页面课程数据接口【https://www.icourse163.org/channel/125001.htm】
// 定义两个爬取器，一个爬取列表，一个爬取详细

func TestMain02(t *testing.T) {
	listC := colly.NewCollector()

	listC.OnHTML("div._2Mzxu", func(classE *colly.HTMLElement) {
		classType := classE.ChildText("p._1vJfX")
		fmt.Println("classType ", classType)
	})

	listC.Visit("https://www.icourse163.org/channel/125001.htm")

}
