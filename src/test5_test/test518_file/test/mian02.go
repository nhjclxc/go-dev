package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"strings"
)

// go get -u github.com/dustin/go-humanize

// 带进度条的大文件下载
func main2() {

	// 下面的示例是带有进度条的大文件下载，我们将响应主体传递到其中，io.Copy()但是如果使用a，TeeReader则可以传递计数器来跟踪进度。
	//在下载时，我们还将文件另存为临时文件，因此在完全下载文件之前，我们不会覆盖有效文件。

	fileUrl := "https://dldir1v6.qq.com/weixin/Universal/Windows/WeChatWin.exe"

	fmt.Println("Download Started")

	err := DownloadFile("WeChatWin2.exe", fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
