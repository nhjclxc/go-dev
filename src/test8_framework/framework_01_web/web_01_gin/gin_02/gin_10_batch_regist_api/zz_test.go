package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	// curl -v http://test20251222.ajyun.cc:81/path1893761722.js -x 192.168.207.197:80

	cacheEndpoint := "http://192.168.207.197:80"
	targetUrl := "http://test20251222.ajyun.cc:81/path1893761722.js"

	transport := &http.Transport{}
	proxyURL, err := url.Parse(cacheEndpoint)
	if err == nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	timeout := 30 * time.Second
	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, targetUrl, nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}

	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
	errMsg := fmt.Sprintf("fetch failed: status=%d", resp.StatusCode)
	fmt.Println(errMsg)

}
