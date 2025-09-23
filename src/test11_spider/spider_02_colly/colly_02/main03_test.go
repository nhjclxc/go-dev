package colly_02

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"testing"
)

// Colly 登录爬取：轻松搞定需要登录的网站
// 需要登录才能爬数据？用 Colly 的 Post 方法模拟登录，像逛自家网站一样轻松抓取！
// https://m.w3cschool.cn/colly/colly-examples-login.html
func TestMain03(t *testing.T) {

	// 创建一个收集器
	c := colly.NewCollector()
	// 4994649,lxc123456

	// 模拟登录（替换为实际登录 URL 和参数）
	err := c.Post("https://www.w3cschool.cn/login", map[string]string{
		"username": "4994649",   // 替换为你的用户名
		"password": "lxc123456", // 替换为你的密码
	})
	if err != nil {
		log.Fatal("登录失败：", err)
	}

	// 登录成功后，添加回调函数
	c.OnResponse(func(r *colly.Response) {
		log.Println("收到响应，状态码：", r.StatusCode)
	})

	c.OnHTML("div.project-list a.project-a", func(e *colly.HTMLElement) {
		fmt.Printf("链接文字: %s \n", e.Attr("title"))
	})

	c.OnHTML("div.project-list", func(e *colly.HTMLElement) {
		// 课程标题（优先取属性）
		title := e.ChildAttr("a.project-a", "title")
		if title == "" {
			title = strings.TrimSpace(e.ChildText("h5.bkname a"))
		}

		// 课程链接（绝对路径）
		link := e.Request.AbsoluteURL(e.ChildAttr("a.project-a", "href"))

		// 封面图
		cover := e.Request.AbsoluteURL(e.ChildAttr("a.project-a img", "src"))

		// 观看人数
		watchNum := e.ChildText("div.watch-number span")

		// VIP 标签（如果存在）
		vip := e.ChildText("a.portlet-vip-mark")

		fmt.Printf("标题: %s\n链接: %s\n封面: %s\n观看数: %s\nVIP: %s\n\n",
			title, link, cover, watchNum, vip)
	})

	// 开始爬取（访问需要登录才能看到的页面）
	c.Visit("https://www.w3cschool.cn/my#mycollection")
}
