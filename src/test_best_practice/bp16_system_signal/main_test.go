package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

// 如何拦截系统信号
// https://labex.io/zh/tutorials/go-how-to-intercept-system-signals-450907
func Test1(t *testing.T) {

	/*
		常见系统信号
		信号		编号	描述
		SIGINT	2	来自键盘的中断 (Ctrl+C)
		SIGTERM	15	终止信号
		SIGKILL	9	立即终止进程
		SIGHUP	1	在控制终端上检测到挂起
		SIGALRM	14	闹钟信号
	*/
	fmt.Printf("当前 PID: %d\n", os.Getpid()) // <- 打印进程id

	// 默认信号处理
	signals := make(chan os.Signal, 1)
	//signal.Notify(signals)
	signal.Notify(signals, syscall.SIGALRM)
	//signal.Notify(signals, syscall.SIGINT)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("等待信号...")
	sig := <-signals
	fmt.Printf("接收到信号: %v\n", sig)
}

// Test2 特定信号拦截
func Test2(t *testing.T) {
	fmt.Printf("当前 PID: %d\n", os.Getpid()) // <- 打印进程id

	signals := make(chan os.Signal, 1)
	signal.Notify(signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		sig := <-signals
		switch sig {
		case syscall.SIGINT:
			fmt.Println("接收到SIGINT")
		case syscall.SIGTERM:
			fmt.Println("接收到SIGTERM")
		}
		os.Exit(0)
	}()

	select {}
}

// Test3 复杂信号处理示例
func Test3(t *testing.T) {
	fmt.Printf("当前 PID: %d\n", os.Getpid()) // <- 打印进程id

	sigChan := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGKILL,
	)

	go func() {
		for {
			select {
			case sig := <-sigChan:
				switch sig {
				case syscall.SIGINT:
					fmt.Println("开始优雅关闭")
					time.Sleep(2 * time.Second)
					done <- true
				case syscall.SIGKILL:
					fmt.Println("接收到终止信号")
					os.Exit(1)
				}
			}
		}
	}()

	<-done
	fmt.Println("关闭完成")
}

func Test5(t *testing.T) {

	now := time.Now()
	now = now.Add(-10 * time.Hour)
	// 以不同布局格式化时间
	formattedTimeLayout := now.Format("3:04 PM, January 2, 2006")
	fmt.Println("格式化后的时间（自定义布局）:", formattedTimeLayout)
	fmt.Println("<UNK>:", time.Since(now))

}

func Test6(t *testing.T) {
	deadline := time.Now().Add(5 * time.Second)
	remaining := deadline.Sub(time.Now())
	fmt.Println("<remaining>:", remaining)

	deadline2 := time.Now().Add(5 * time.Second)
	remaining2 := time.Until(deadline2)
	fmt.Println("<remaining2>:", remaining2)
}

func Test7(t *testing.T) {
	header := "user=a&name=zhangsan&age=11"

	before, after, found := strings.Cut(header, "&")
	fmt.Printf("header: %s, before: %s, after: %s\n", header, before, after)
	fmt.Println(found)

}

func Test8(t *testing.T) {

	var after string
	var before string
	var found bool
	after = "user=a&name=zhangsan&age=11"

	m := make(map[string]string)
	for {
		before, after, found = strings.Cut(after, "&")
		if key, val, ok := strings.Cut(before, "="); ok {
			m[key] = val
		}
		if !found {
			break
		}
	}

	fmt.Println(m)

}

func Test9(t *testing.T) {
	// 高性能字符串追加 (fmt.Appendf)Go 1.19 引入了直接向字节切片追加格式化字符串的能力，
	// 避免了 fmt.Sprintf 带来的隐式内存分配。

	var originStr []byte = []byte("num str: ")

	for i := range 10 {
		originStr = fmt.Appendf(originStr, "%d,", i)
	}

	fmt.Println(string(originStr))

}

func Test10(t *testing.T) {

	layout := "2006-01-02 15:04:05"

	startStr := "2026-02-01 00:00:00"
	endStr := "2026-03-01 00:00:00"

	start, _ := time.Parse(layout, startStr)
	end, _ := time.Parse(layout, endStr)

	diff := end.Sub(start)

	minutes := diff.Minutes()

	fmt.Printf("相差分钟数: %.0f\n", minutes)
}

/*
$upstream_addr - $upstream_status '

192.168.207.196:33366 - 502, 200
192.168.207.197:33366 - 502, -

'$sent_http_x_upstream_type $sent_http_x_upstream_addr $sent_http_X_Cache '
- - "-"
- - \"-\"

'$server_addr $server_type $body_bytes_sent';
127.0.0.1 0 10240
127.0.0.1 0 0

192.168.207.196:33366,
*/
