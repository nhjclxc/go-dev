package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	go func() {
		time.Sleep(12 * time.Second)
		fmt.Println("cancel()取消", time.Now().Format("15:04:05"))
		cancel()
	}()

	retryInterval := 5 * time.Second

	// 测试1
	// time.Sleep的阻塞了检测机制：time.Sleep() 是一个阻塞调用。也就是说，在 Sleep 期间，goroutine 什么都不干，也不会检测到 ctx 是否被取消。
	//	进入time.Sleep后，如果此时其他通道已经有数据可以读取了，也不会立即去执行其他协程，而会等待time.Sleep的结束后才能去执行其他协程
	/*
		time.Sleep(5 * time.Second)
		它的行为是：
			当前 goroutine 彻底阻塞 5 秒；
			Go 调度器会挂起这个 goroutine；
			它不会执行任何 select、channel、ctx.Done() 检测；
			直到 5 秒后，这个 goroutine 被唤醒；
			才能继续执行下一行代码。
	*/
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("检测设备 ID 被取消:", time.Now().Format("15:04:05"))
	//		return
	//	default:
	//		time.Sleep(retryInterval)
	//		//	fmt.Println("执行default")
	//	}
	//	fmt.Println("执行外部", time.Now().Format("15:04:05"))
	//}
	// 输出
	// 执行外部 2025-11-12 09:51:02
	//执行外部 2025-11-12 09:51:07
	//cancel()取消 2025-11-12 09:51:09
	//检测设备 ID 被取消: 2025-11-12 09:51:12
	// 从上面的输出可以看出time.Sleep不可以立即响应<-ctx.Done()，只有当前goroutine被唤醒后才能触发<-ctx.Done()

	//测试2
	// time.After的阻塞了检测机制：非阻塞，time.After只触发一次
	// 	如果检测到time.After分支
	/*
		<-time.After(5 * time.Second)会在内部启动一个 定时器 goroutine，返回一个 channel：
		它的行为是：
			它返回的 channel 不会阻塞当前 goroutine；
			当前 goroutine 在 select 中等待多个事件；
			一旦 ctx.Done() 先触发，select 会立即执行该分支；
			如果 time.After 到期了而 ctx 没取消，则执行超时分支。
	*/
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("检测设备 ID 被取消:", time.Now().Format("15:04:05"))
	//		return
	//	case <-time.After(retryInterval):
	//		fmt.Println("执行外部", time.Now().Format("15:04:05"))
	//	}
	//}
	// 输出
	// 执行外部 10:04:14
	//执行外部 10:04:19
	//cancel()取消 10:04:21
	//检测设备 ID 被取消: 10:04:21
	// 从上面的输出可以看出time.After可以立即响应<-ctx.Done()

	//测试3
	// time.Tick的阻塞了检测机制：非阻塞，time.Tick每 5 秒触发一次
	// 	如果检测到time.After分支
	/*
		time.Tick(5 * time.Second)，创建一个定时器Ticker 对象返回
		time.Tick() 直接返回了 ticker.C，你没法调用 Stop() 去释放定时器。
		它的行为是：
			启动一个内部 goroutine；
			每隔 5 秒向 channel 发送当前时间；
			一直循环，直到 ticker.Stop() 被调用。
	*/
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("检测设备 ID 被取消:", time.Now().Format("15:04:05"))
	//		return
	//	case <-time.Tick(retryInterval):
	//		fmt.Println("执行外部", time.Now().Format("15:04:05"))
	//	}
	//}
	// 执行外部 10:24:29
	//执行外部 10:24:34
	//cancel()取消 10:24:36
	//检测设备 ID 被取消: 10:24:36

	// 测试4
	ticker := time.NewTicker(retryInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx被取消", time.Now().Format("15:04:05"))
			return
		case <-ticker.C:
			fmt.Println("执行ticker", time.Now().Format("15:04:05"))
		}
	}
	// 执行ticker 10:33:00
	//执行ticker 10:33:05
	//cancel()取消 10:33:07
	//ctx被取消 10:33:07

	/*
	   | 功能点    | `time.After`       | `time.Tick`        | `time.NewTicker`  |
	   | ------ | ------------------ | ------------------ | ----------------- |
	   | 用途     | 等待一次超时             | 周期性触发              | 周期性触发（可停止）        |
	   | 返回类型   | `<-chan time.Time` | `<-chan time.Time` | `*time.Ticker`    |
	   | 是否循环   | ❌ 否                | ✅ 是                | ✅ 是               |
	   | 能否手动停止 | ❌ 否                | ❌ 否                | ✅ `ticker.Stop()` |
	   | 是否泄漏风险 | ✅ 否                | ⚠️ 有泄漏风险           | ❌ 无泄漏风险           |
	   | 常用场景   | 单次延时等待             | 简单 demo / 无限循环     | 可控的定时任务           |

	*/
}

func qq() {
	ctx := context.Background()
	retryInterval := 3 * time.Second
	// 方式1
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(retryInterval):
			fmt.Println("doSomething()")
		}
	}
	// 方式2
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(retryInterval):
			fmt.Println("doSomething()")
		}
	}
}
