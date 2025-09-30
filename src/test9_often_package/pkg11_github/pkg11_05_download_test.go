package pkg11_github

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// 带进度条的大文件下载
// go get github.com/dustin/go-humanize
func Test1105(t *testing.T) {
	fmt.Println("Download Started")

	fileUrl := "https://s3plus.meituan.net/v1/mss_7e425c4d9dcb4bb4918bbfa2779e6de1/mpack/default/demo.flv"
	err := DownloadFile("demo.flv", fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
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
