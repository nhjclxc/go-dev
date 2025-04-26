package main

import (
	"io"
	"net/http"
	"os"
)

// 普通文件下载
// 本示例说明如何从网上将文件下载到本地计算机。通过io.Copy()直接使用并传递响应主体，我们将数据流式传输到文件中，
// 而不必将其全部加载到内存中-小文件不是问题，但下载大文件时会有所不同。
func main1() {
	fileUrl := "https://dldir1v6.qq.com/weixin/Universal/Windows/WeChatWin.exe"
	if err := DownloadFile1("WeChatWin1.exe", fileUrl); err != nil {
		panic(err)
	}
}

// download file会将url下载到本地文件，它会在下载时写入，而不是将整个文件加载到内存中。
func DownloadFile1(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
