

https://github.com/PuerkitoBio/goquery




Go 语言的 [`goquery`](https://github.com/PuerkitoBio/goquery) 是一个非常流行的 HTML 解析库，它的语法与 jQuery 非常相似，适合用于网页爬虫、数据抽取、结构化 HTML 处理等场景。

---

## 🧠 一、goquery 是什么？

`goquery` 是 Go 语言对网页 DOM 的操作库，它基于 `golang.org/x/net/html`，封装了很多类 jQuery 的方法，比如：

```go
doc.Find("div.title").Text()
doc.Find("a").Each(...)
doc.Find("img").Attr("src")
```

---

## 🚀 二、安装方式

```bash
go get github.com/PuerkitoBio/goquery
```

---

## ✅ 三、最小可运行示例：抓取网页标题和所有链接

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 1. 发起 HTTP 请求
	res, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("状态码错误: %d", res.StatusCode)
	}

	// 2. 加载 HTML 文档
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 抓取 <title> 标签内容
	title := doc.Find("title").Text()
	fmt.Println("网页标题:", title)

	// 4. 抓取所有超链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		text := s.Text()
		if exists {
			fmt.Printf("链接：%s (%s)\n", text, href)
		}
	})
}
```

---

## 🔧 四、常见 goquery 操作技巧

| 操作       | 示例                                                |
| -------- | ------------------------------------------------- |
| 查找标签     | `doc.Find("div")`                                 |
| 查找 class | `doc.Find(".className")`                          |
| 查找 ID    | `doc.Find("#idName")`                             |
| 获取文本     | `s.Text()`                                        |
| 获取属性     | `s.Attr("href")`                                  |
| 遍历结果     | `Each(func(i int, s *goquery.Selection) { ... })` |
| 获取子节点    | `s.Find("img")`                                   |
| 父节点/兄弟节点 | `s.Parent()`, `s.Next()`                          |
| 筛选       | `s.Filter(".active")`, `s.HasClass("xx")`         |

---

## 🧪 五、实战：爬取简书首页文章标题

```go
res, _ := http.Get("https://www.jianshu.com/")
doc, _ := goquery.NewDocumentFromReader(res.Body)

doc.Find(".title a").Each(func(i int, s *goquery.Selection) {
	title := s.Text()
	link, _ := s.Attr("href")
	fmt.Printf("[%d] %s - https://www.jianshu.com%s\n", i+1, title, link)
})
```

---

## 🧰 六、配合使用的库建议

| 任务          | 推荐库                                              |
| ----------- | ------------------------------------------------ |
| 发请求         | `net/http` 或 `github.com/go-resty/resty/v2`（更强大） |
| 解析 HTML     | `github.com/PuerkitoBio/goquery`                 |
| 自动处理 Cookie | `http.CookieJar`                                 |
| 自动重试        | `github.com/hashicorp/go-retryablehttp`          |
| 并发控制        | goroutine + `sync.WaitGroup` + `chan`            |

---

## 🧱 七、进阶建议

* ✅ 使用 `User-Agent` 模拟浏览器防止被拦截
* ✅ 支持代理 IP，避免被封
* ✅ 设置超时，防止阻塞
* ✅ 抽象数据结构和通用爬虫框架
* ✅ 可配合 `colly` 框架用于更复杂的爬虫需求（见下）

---

## 🔁 goquery 与 colly 的关系？

| 特性   | goquery   | colly           |
| ---- | --------- | --------------- |
| 定位   | HTML 解析器  | 爬虫框架，封装 goquery |
| 适合用途 | HTML 数据提取 | 自动爬取网页、分页、抓图等   |
| 并发   | 需自己控制     | 内置并发抓取支持        |
| 简洁程度 | 更灵活、偏底层   | 更自动化、封装性更强      |

> 🔸**建议：**

* 如果只是提取页面数据，`goquery` 足够；
* 如果要构建爬虫系统（自动翻页、限速、并发等），建议用 `colly`。

---

## 📌 示例项目练习推荐

1. 爬取豆瓣电影 Top250
2. 抓取知乎热榜标题及链接
3. 提取掘金首页文章信息
4. 抓取并下载网页图片

---

## ✅ 总结

| 学习目标  | 建议操作                      |
| ----- | ------------------------- |
| 快速入门  | 用 `goquery` 抓几个简单网站的标题、链接 |
| 熟悉选择器 | 多练习 class、id、子节点查找        |
| 提升能力  | 配合并发、限速、失败重试、代理使用等        |
| 项目实战  | 可做一个新闻聚合爬虫、文章摘要生成工具等      |

---

如果你有想抓取的网址，我可以手把手为你写一段 `goquery` 爬虫代码。你是否有目标网站或需求？
