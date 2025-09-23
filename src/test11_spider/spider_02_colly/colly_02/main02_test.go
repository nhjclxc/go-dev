package colly_02

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"testing"
)

// Colly 错误处理：让爬虫永不“崩溃”
// 网络不稳定、404、503 是常态。给爬虫加 3 行“保险丝”，出错也能优雅记录、自动重试。
func TestMain02(t *testing.T) {

	c := colly.NewCollector()

	c.OnError(func(resp *colly.Response, err error) {
		fmt.Printf("❌ 爬取失败：%s\n状态码：%d\n错误信息：%v\n",
			resp.Request.URL, resp.StatusCode, err)
		if resp.StatusCode == 404 {
			log.Printf("url[%s] 404 \n", resp.Request.URL)
		}

		if resp.StatusCode >= 500 {
			// 服务器错误，重试
			err := resp.Request.Retry()
			if err != nil {
				fmt.Println("重试失败:", err)
				return
			}

			//retries := resp.Request.Ctx.GetAny("retries")
			//count := 0
			//if retries != nil {
			//	count = retries.(int)
			//}
			//if count < 3 { // 最多重试3次
			//	fmt.Println("重试次数:", count+1)
			//	resp.Request.Ctx.Put("retries", count+1)
			//
			//	// 延迟重试（可选）
			//	time.Sleep(2 * time.Second)
			//
			//	// 重新发起请求
			//	resp.Request.Retry()
			//} else {
			//	fmt.Println("已达到最大重试次数:", resp.Request.URL)
			//}
		}
	})

	c.Visit("https://www.w3cschool.cn/notf111ound")

}
