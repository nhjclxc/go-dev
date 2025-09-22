package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// colly快速开始案例
func main01() {
	//if true {
	//	saveImage("https://unsplash-assets.imgix.net/unsplashplus/module-01-v2.jpg?fm=jpg&q=60&w=3000", "/Users/lxc20250729/data/go/spider/colly/01")
	//	return
	//}

	//链接：https://juejin.cn/post/7231130096337207353

	// 创建一个新的收集对象
	c := colly.NewCollector()

	// 在访问页面之前执行的回调函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// 在访问页面之后执行的回调函数
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL.String())
	})

	// 在访问页面时发生错误时执行的回调函数
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	//// 在访问页面时发生重定向时执行的回调函数
	//c.OnRedirect(func(r *colly.Response) {
	//	fmt.Println("Redirected to", r.Request.URL.String())
	//})

	var wg sync.WaitGroup

	//OnHTML 用于注册回调函数，只在元素出现时处理数据，当爬虫抓取的页面中匹配某个 CSS 选择器的元素出现时，就会触发这个回调。
	c.OnHTML("img", func(e *colly.HTMLElement) {
		// 在访问页面时发现图片时执行的回调函数 方法注册了一个回调函数，用于处理页面中的图片元素。
		url := e.Attr("src")
		if url != "" {
			//fmt.Println("Found image:", url)
			//wg.Add(1)
			//go saveImage(&wg, url, "/Users/lxc20250729/data/go/spider/colly/01")
		}
	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		//fmt.Println("script = ", e.Attr("src"))
	})

	// 使用css的属性选择器查找数据
	c.OnHTML("div.container-CgvcQa div", func(e *colly.HTMLElement) {
		fmt.Println("script = ", e.Attr("src"))
	})
	// 不监测某些标签
	//c.OnHTMLDetach("img")

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("OnScraped to", r.Request.URL.String())
	})

	// 发起访问  输入你要访问的网址
	c.Visit("https://unsplash.com/")

	wg.Wait()

}

func saveImage(wg *sync.WaitGroup, url string, path string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("download err ", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("创建目录失败: %v \n", err)
			return
		}
	}

	fp := filepath.Join(path, getFilename(url))
	out, err := os.Create(fp)
	if err != nil {
		fmt.Printf("创建文件失败: %v \n", err)
		return
	}
	defer out.Close()

	// 拷贝数据
	_, err = io.Copy(out, resp.Body)

	fmt.Printf("saveIamge successful [%s] .\n", fp)
}

func getFilename(url string) string {
	if url == "" {
		return string(rand.Int32()) + ".jpg"
	}

	sss := strings.Split(strings.Split(url, "?")[0], "/")
	a := sss[len(sss)-1]
	return a + ".jpg"
}
