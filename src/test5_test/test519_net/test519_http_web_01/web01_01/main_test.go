package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

// 压测参数
const (
	targetURL   = "http://localhost:8080/world" // 测试接口
	concurrency = 50                            // 并发数
	duration    = 10 * time.Second              // 测试总时间
)

func TestName(t *testing.T) {

	var wg sync.WaitGroup
	var successCount int64

	stop := make(chan struct{})

	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := http.Client{Timeout: 5 * time.Second}

			for {
				select {
				case <-stop:
					return
				default:
					resp, err := client.Get(targetURL)
					if err == nil && resp.StatusCode == 200 {
						successCount++
					}
					if resp != nil {
						err := resp.Body.Close()
						if err != nil {
							return
						}
					}
				}
			}
		}()
	}

	// 每秒打印当前 QPS
	ticker := time.NewTicker(time.Second)
	go func() {
		var lastCount int64
		for range ticker.C {
			currentCount := successCount
			qps := currentCount - lastCount
			fmt.Printf("当前 QPS: %d\n", qps)
			lastCount = currentCount
		}
	}()

	// 等待测试时间结束
	time.Sleep(duration)
	close(stop)
	wg.Wait()
	ticker.Stop()

	totalTime := time.Since(start).Seconds()
	fmt.Printf("总请求数: %d, 平均 QPS: %.2f\n", successCount, float64(successCount)/totalTime)
}
