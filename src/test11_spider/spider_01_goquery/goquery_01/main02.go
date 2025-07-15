package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func main() {

	// 爬取 https://www.jianshu.com/ 的首页文章标题

	// 1. 发起 HTTP 请求
	res, err := http.Get("https://www.jianshu.com/")
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
	/*
		| 操作       | 示例                                                |
		| --------        | ------------------------------------------------- |
		| 查找标签         | `doc.Find("div")`                                 |
		| 查找 class      | `doc.Find(".className")`                          |
		| 查找 ID         | `doc.Find("#idName")`                             |
		| 获取文本        | `s.Text()`                                        |
		| 获取属性        | `s.Attr("href")`                                  |
		| 遍历结果        | `Each(func(i int, s *goquery.Selection) { ... })` |
		| 获取子节点      | `s.Find("img")`                                   |
		| 父节点/兄弟节点  | `s.Parent()`, `s.Next()`                          |
		| 筛选           | `s.Filter(".active")`, `s.HasClass("xx")`         |

	*/
	// doc.Find(传入css的选择器)
	nodes := doc.Find("#list-container").Nodes
	for _, node := range nodes {

		fmt.Println("网页标题:", node.Data)

		childNodes := node.ChildNodes()

		childNodes(func(child *html.Node) bool {
			fmt.Println("childNodes: " + child.Data) // 或 child.SomeField

			if "ul" == child.Data {
				ulChildNodes := child.ChildNodes()
				ulChildNodes(func(child *html.Node) bool {
					fmt.Println("ulChildNodes: " + child.Data) // 或 child.SomeField
					attr := child.Attr
					liNode := false
					for _, attribute := range attr {
						liNode = strings.HasPrefix(attribute.Val, "note-")
						if liNode {
							break
						}
						fmt.Println("ulChildNodes.child.Attr: " + attribute.Key + ", val = " + attribute.Val) // 或 child.SomeField
					}

					if liNode {
						liChildNodes := child.ChildNodes()
						liChildNodes(func(child *html.Node) bool {
							fmt.Println("ulChildNodes.liChildNodes: " + child.Data) // 或 child.SomeField

							if "a" == child.Data {

							}


							return true        // 继续遍历；如果返回 false 则中断遍历
						})
					}


					return true        // 继续遍历；如果返回 false 则中断遍历
				})

			}


			return true        // 继续遍历；如果返回 false 则中断遍历
		})





	}
	//
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//// 4. 抓取所有超链接
	//doc.Find("a").Each(func(i int, s *goquery.Selection) {
	//	href, exists := s.Attr("href")
	//	text := s.Text()
	//	if exists {
	//		fmt.Printf("链接：%s (%s)\n", text, href)
	//	}
	//})
}
