package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	target := "http://localhost:8080/ping"
	totalRequests := 3
	concurrency := 50

	var wg sync.WaitGroup
	var mutex sync.Mutex

	successCount := 0
	failCount := 0

	var totalTime time.Duration
	var maxTime time.Duration
	minTime := time.Hour // 初始化为 1 小时，便于比较

	startAll := time.Now()
	sem := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(target)
			duration := time.Since(start)

			mutex.Lock()
			totalTime += duration
			if duration > maxTime {
				maxTime = duration
			}
			if duration < minTime {
				minTime = duration
			}

			if err != nil || resp.StatusCode != http.StatusOK {
				failCount++
			} else {
				successCount++
			}
			mutex.Unlock()

			if resp != nil {
				resp.Body.Close()
			}
			<-sem
		}()
	}

	wg.Wait()
	elapsed := time.Since(startAll)

	avgQPS := float64(totalRequests) / elapsed.Seconds()
	avgResponseTime := totalTime / time.Duration(totalRequests)

	fmt.Println("========== 压力测试结果 ==========")
	fmt.Printf("目标地址     ：%s\n", target)
	fmt.Printf("总请求数     ：%d\n", totalRequests)
	fmt.Printf("并发数       ：%d\n", concurrency)
	fmt.Printf("成功请求数   ：%d\n", successCount)
	fmt.Printf("失败请求数   ：%d\n", failCount)
	fmt.Printf("总耗时       ：%.3fs\n", elapsed.Seconds())
	fmt.Printf("平均 QPS     ：%.2f\n", avgQPS)
	fmt.Printf("平均响应时间 ：%v\n", avgResponseTime)
	fmt.Printf("最小响应时间 ：%v\n", minTime)
	fmt.Printf("最大响应时间 ：%v\n", maxTime)
}
