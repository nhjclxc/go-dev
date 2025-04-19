package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://www.bilibili.com",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}

	// 使用http.Get 向某个网站发送get请求获取数据
	res, _ := http.Get("https://www.bilibili.com")
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("Got: %q", string(data))
}
