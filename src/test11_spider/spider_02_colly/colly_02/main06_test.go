package colly_02

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"testing"
)

//Colly 队列爬取：让爬虫任务排队执行
//一键把 1000 个 URL 交给“小管家”排队，自动限速、并发可控，复制即可跑！

func TestMain06(t *testing.T) {

	q, err := queue.New(
		3,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)
	if err != nil {
		return
	}

	c := colly.NewCollector()

	// 3. 打印每次访问的 URL
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("正在访问：", r.URL)
	})

	// 4. 把 URL 压进队列（模拟 5 个课程页）
	for i := 1; i <= 5; i++ {
		q.AddURL(fmt.Sprintf("https://www.w3cschool.cn/course/%d", i))
	}

	// 5. 启动队列，自动消费所有 URL
	q.Run(c)

}
