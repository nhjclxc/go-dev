package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//test02_01()

	//test02_02()

	//test02_03()

	//test02_04()

	//test02_05()

	test02_06()

}

func test02_06() {
	// 练习题 6：组合传值 + 超时（进阶）
	//写一个 apiCall(ctx) 函数：
	//读取 traceID（用 WithValue 设置）
	//模拟远程调用 5 秒
	//设置调用超时为 2 秒，如果超时输出 "timeout"，否则输出 "api done: traceID=xxxx"
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "traceID", "qwertyuhygfd15616156156")
	ctx, cancelFunc := context.WithTimeout(ctx, 2*time.Second)
	//ctx, cancelFunc := context.WithTimeout(ctx, 6*time.Second)
	defer cancelFunc()

	apiCall(ctx)

}

func apiCall(ctx context.Context) {
	// 从 context 中获取 traceID
	traceID, ok := ctx.Value("traceID").(string)
	if !ok {
		traceID = "unknown"
	}

	fmt.Println("apiCall 开始，traceID:", traceID)

	select {
	case <-time.After(5 * time.Second): // 模拟远程调用
		fmt.Printf("✅ api done: traceID=%s\n", traceID)
	case <-ctx.Done():
		fmt.Printf("❌ timeout: traceID=%s, err=%v\n", traceID, ctx.Err())
	}


	//fmt.Println("进入 apiCall")
	//select {
	//case <-ctx.Done():
	//	fmt.Println("远程调用超时 timeout ！！！")
	//	return
	//case <-time.After(5 * time.Second):
	//	fmt.Print("api done: traceID=%s！！！", ctx.Value("traceID"))
	//}
}

func test02_05() {
	// 练习题 5：goroutine 泄漏防止（结合 cancel）
	//写一个并发函数 watchDog(ctx)，每 1 秒输出一次 "watching..."，要求主程序退出前手动取消 ctx，防止 goroutine 泄漏。
	baseCtx := context.Background()
	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	// 启动 watchDog goroutine
	go watchDog(ctx)

	// 模拟主程序运行 4 秒后取消
	time.Sleep(4 * time.Second)
	fmt.Println("main 准备取消 watchDog")
	cancel() // 主动取消 context，通知 goroutine 退出

	// 等待 goroutine 打印完退出信息
	time.Sleep(1 * time.Second)
	fmt.Println("main 退出")

}

func watchDog(ctx context.Context) {

	// 避免泄漏的两种方式：
	//context 控制生命周期（推荐）
	//明确的退出信号（channel、flag）
	for {
		select {
		case <-ctx.Done():
			fmt.Println("watchDog 接收到取消信号:", ctx.Err())
			fmt.Println("watchDog 退出")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("watching...")
		}
	}


}

func test02_04() {
	// 练习题 4：定时截止（WithDeadline）
	//编写一个函数 queryDB(ctx)，模拟数据库查询 3 秒。
	//但设置 context 的 deadline 是当前时间加 2 秒。
	//如果到了 deadline，还没查完，就输出 "deadline exceeded"。

	deadlineTime := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadlineTime)
	defer cancel()

	queryDB(ctx)



}

func queryDB(ctx context.Context) {
	fmt.Println("开始查询数据库...")

	select {
	case <-time.After(3 * time.Second): // 模拟查询耗时 // 可打断的sleep
		fmt.Println("查询成功完成！")
	case <-ctx.Done(): // deadline 到了或被取消
		fmt.Println("查询失败，原因：", ctx.Err())
	}

	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("deadline exceeded")
	//		fmt.Println("退出 queryDB")
	//		return
	//	case <-time.After(500 * time.Millisecond):
	//		fmt.Println("查询数据库ing...")
	//	}
	//}


	//fmt.Println("进入 queryDB")
	//for {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("deadline exceeded")
	//		return
	//	default:
	//		fmt.Println("查询数据库ing...")
	//		time.Sleep(3 * time.Second) // 这是阻塞调用，一旦进入 sleep，期间就无法检查 ctx 是否已被 cancel。如果 deadline 是 2 秒，而 sleep 是 3 秒，它仍会 sleep 完之后才响应超时。
	//	}
	//}
	//fmt.Println("退出 queryDB")
}

func test02_03() {
	//练习题 3：传值（WithValue）
	//编写一个函数 handleRequest(ctx)，它会从 ctx 中读取 userID 和 authToken，然后输出。
	//在 main() 中构造 context，并传入这两个值，传给 handleRequest。

	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "userId", "666")
	ctx = context.WithValue(ctx, "authToken", "authTokenauthTokenauthTokenauthTokenauthToken")

	go handleRequest(ctx)

	time.Sleep(1 * time.Second)


}

func handleRequest(ctx context.Context) {

	fmt.Println("ctx.userId: ", ctx.Value("userId"))
	fmt.Println("ctx.authToken: ", ctx.Value("authToken"))

}


func test02_02() {
	//练习题 2：超时控制（WithTimeout）
	//编写一个函数 doWork(ctx)，模拟一个耗时任务（比如 sleep 5 秒）。
	//要求设置超时时间为 3 秒，一旦超时就中断任务，输出 "timeout!"，否则输出 "work done!"。

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	doWork(ctx)

}

func doWork(ctx context.Context) {
	fmt.Println("进入 doWork")
	time.Sleep(5 * time.Second)
	fmt.Println("退出 doWork")
}

func test02_01() {
	// 练习题 1：基本使用（WithCancel）
	//编写一个函数 startWorker(ctx)，它会每秒打印 "working..."，当 ctx 被取消后，打印 "stopped" 并退出。
	//在 main() 中运行该 worker，然后在 3 秒后取消 ctx。

	baseCtx := context.Background()
	ctx, cancel := context.WithCancel(baseCtx)

	startWorker(ctx)

	time.Sleep(3 * time.Second)
	cancel()

}

func startWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopped\n")
			return
		default:
			fmt.Printf("working... \n")
			time.Sleep(1*time.Second)
		}
	}
}

