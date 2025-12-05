package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// 网页地址：https://mirrors.tuna.tsinghua.edu.cn/centos-stream/9-stream/BaseOS/ppc64le/iso/
// 源站文件地址：https://mirrors.tuna.tsinghua.edu.cn/centos-stream/9-stream/BaseOS/ppc64le/iso/CentOS-Stream-9-20251124.0-ppc64le-boot.iso

// 模拟代理回源，且当前服务去源站文件地址下载文件到当前服务器，同时向客户端提供服务

// var origin = "https://mirrors.tuna.tsinghua.edu.cn/centos-stream/9-stream/BaseOS/ppc64le/iso/CentOS-Stream-9-20251124.0-ppc64le-boot.iso"？
// var origin = "https://software.download.prss.microsoft.com/dbazure/Win10_22H2_Chinese_Simplified_x32v1.iso?t=d579e8d9-31b8-4375-bade-ddadacec93ad&P1=1764145045&P2=601&P3=2&P4=wzokwwFOktNrGvi4rAdcs433L%2fh9wpX2Mozrj0SCyoWR%2bFmA5bMCbRjS7vgwykRtJ3UJIUw5VDySgugvJWSZ3cc%2b1E4Lf%2fon%2bwBBMWlP1BxhnGl%2bBtF905Yo9TJ3b9SsFXFHwcF%2fzdBLkokdBEeudBVrwUP5tE9kTOp8oenc2gEiHRk2Pft8yN7C8ThETFbc%2fGjB1ccer6iLWyUE%2fHjqsoStUTzU5mUm6CR0XqegewFIQZutu0Tz3enp%2fNCOoONSUTjEio%2f%2bywoWUdgJDxaPnX6vT%2bncwJx5oHMtT2d9WMqf%2fQQudbv9hAimW6uraL5bcMvRwWljoFwjnnW8tsnHhg%3d%3d"
var origin = "https://dl.google.com/dl/android/aosp/google_devices-tokay-12990991-8abda025.tgz?hl=zh-cn"
var cacheDir = "/Users/lxc20250729/cdn-cache" // 本地缓存目录

func proxyHandler(w http.ResponseWriter, r *http.Request) {

	// 解析查询参数
	filename := r.URL.Query().Get("filename") // 获取 filename 参数的值
	if filename == "" {
		http.Error(w, "filename parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Requested filename: %s\n", filename)

	// 先检查本地是否有该文件
	localPath := filepath.Join(cacheDir, filename)
	fmt.Println("localPath", localPath)

	// 文件存在，提供服务
	if _, err := os.Stat(localPath); err == nil {
		log.Printf("Cache hit: %s", localPath)
		http.ServeFile(w, r, localPath)
		return
	}

	log.Printf("Cache miss, fetching from origin: %s", origin)

	// 目标文件不存在，去源站拉，同时缓存到本地和给客户端提供服务
	originUrl, _ := url.Parse(origin)
	client := &http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Get(originUrl.String())
	if err != nil {
		http.Error(w, "Failed to fetch from origin", http.StatusBadGateway)
		log.Printf("Error fetching origin: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Origin return non-200", http.StatusBadGateway)
		log.Printf("Origin return non-200, now status: %d", resp.StatusCode)
		return
	}

	// 创建缓存目录
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		http.Error(w, "MkdirAll cacheDir err: "+err.Error(), http.StatusBadGateway)
		log.Printf("MkdirAll cacheDir err: %v", err)
		return
	}

	// 创建本地文件用于缓存
	createFile, err := os.Create(localPath)
	if err != nil {
		http.Error(w, "os.Create err: "+err.Error(), http.StatusBadGateway)
		log.Printf("os.Create err: %v", err)
		return
	}
	defer createFile.Close()

	// 边缓存边写
	// 边读取边写入本地 + 返回客户端（流式传输，大文件友好）
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// 创建一个 MultiWriter，写入到MultiWriter的数据会同时写入到MultiWriter的多个目标。
	// mw里面的数据同时会写入w 和 createFile
	mw := io.MultiWriter(w, createFile)
	// 将源站返回的数据写入mw
	if _, err := io.Copy(mw, resp.Body); err != nil {
		log.Printf("error while copying to client & cache: %v", err)
		return
	}

	log.Printf("Finished caching and serving: %s", filename)

}

func main() {

	// http://127.0.0.1:8081/?filename=google_devices-tokay-12990991-8abda025.tgz
	http.HandleFunc("/", proxyHandler)
	log.Println("CDN Proxy on :8081")
	http.ListenAndServe(":8081", nil)
}
