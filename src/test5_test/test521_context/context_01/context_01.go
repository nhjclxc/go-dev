package main

import (
	context "context"
	"fmt"
	"os"
	"time"
)

func main() {

	/*
	context 是 Go 中非常重要的标准库，主要用于：
		控制 超时（timeout）
		支持 取消操作（cancel）
		在请求链中传递 请求范围的数据

	认识context包的所有常见方法
	| 方法                                      | 说明                                 |
	| --------------------------------------- | ---------------------------------- |
	| `context.Background()`                  | 最顶层的 context，常用于 main 或测试，用于创建一个context实例          |
	| `context.TODO()`                        | 占位符，表示“以后再补 context”，用于创建一个context实例               |
	| `context.WithCancel(parent)`            | 返回可取消的子 context                    |
	| `context.WithTimeout(parent, duration)` | 带超时取消的 context                     |
	| `context.WithDeadline(parent, time)`    | 指定时间点取消的 context                   |
	| `context.WithValue(parent, key, value)` | 向 context 中存储键值对数据                 |
	| `ctx.Done()`                            | 返回一个 channel，表示 context 结束（被取消或超时） |
	| `ctx.Err()`                             | 返回 context 被取消的原因                  |
	| `ctx.Value(key)`                        | 获取 context 中的值                     |

	*/

	//test01()

	//test02()

	//test03()

	//test04()

	//test05()

	test06()

}

func test06() {
	// 6. 综合示例：结合 cancel + value + timeout

	baseCtx := context.WithValue(context.Background(), "user", "Alice")
	ctx, cancel := context.WithTimeout(baseCtx, 4*time.Second)  // 4s后超时取消
	defer cancel()

	go worker(ctx, "Worker1")
	go worker(ctx, "Worker2")

	time.Sleep(6 * time.Second) // 等待观察超时退出

}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] 退出: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] 正在处理用户: %v\n", name, ctx.Value("user"))
			time.Sleep(500 * time.Millisecond)
		}
	}
}


func test05() {
	background := context.Background()
	ctx := context.WithValue(background, "user", "Luo Xianchao")
	process(ctx)

	fmt.Println("main.background.user: ", background.Value("user"))

	go processPoint(&ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("main.background.age: ", background.Value("age"))
	fmt.Println("main.ctx.age: ", ctx.Value("age"))

}

func process(ctx context.Context) {
	user := ctx.Value("user")
	if user != nil {
		fmt.Println("当前用户是:", user)
	} else {
		fmt.Println("没有用户信息")
	}
}

func processPoint(ctx *context.Context) {
	user := (*ctx).Value("user")
	if user != nil {
		fmt.Println("当前用户是:", user)
	} else {
		fmt.Println("没有用户信息")
	}
	(*ctx) = context.WithValue((*ctx), "age", 18)
}

func test04() {
	// 4. context.WithDeadline：设定截止时间

	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("正常完成")
		case <-ctx.Done():
			fmt.Println("被 deadline 取消:", ctx.Err())
			os.Exit(-1)
		default:
			fmt.Println("工作中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// context.WithDeadline 和 context.WithTimeout 的区别是什么？？？
	//context.WithDeadline 设置明确的截止时间点，到某个时间了关闭，如2025-07-05 15:00:00关闭，精确控制任务必须在某一时间点前完成
	//context.WithTimeout 设置相对当前时间的超时时间，多久后关闭，如5s后关闭，通用超时控制（如接口请求 3s 超时）


	// 为什么context.WithDeadline和context.WithTimeout要使用defer cancel()来关闭，而context.WithCancel不适应defer可以直接cancel()
	//不管是 context.WithCancel、WithTimeout 还是 WithDeadline，都应该在使用完毕后调用 cancel()，并且推荐用 defer cancel() 来自动释放资源。
	//

	// 🧠 为什么需要 defer cancel()？
	//在 context.WithCancel / WithTimeout / WithDeadline 中，cancel() 函数的作用是：
	//通知子 goroutine 停止工作（发出取消信号）
	//释放上下文相关的资源（memory leak 防止），比如定时器、内部结构等

	// 👇 对比 3 种 context 的行为差异
	// | 类型             | 是否自动超时取消  | 是否需要 cancel()  | 推荐用 `defer cancel()` |
	//| -------------- | --------- | -------------- | -------------------- |
	//| `WithCancel`   | ❌ 不自动取消   | ✅ 需要手动取消       | ✅ 推荐                 |
	//| `WithTimeout`  | ✅ 自动超时取消  | ✅ 仍建议 cancel() | ✅ 强烈推荐               |
	//| `WithDeadline` | ✅ 到期后自动取消 | ✅ 仍建议 cancel() | ✅ 强烈推荐               |
	// 从上面的是否支持自动取消中可以看出，WithCancel是不会自动取消的，一定要去手动关闭，但是是否使用defer是由开发人员自己决定的
	// 而对于WithTimeout和WithDeadline而言，他们是可以由go自动执行cancel()的，程序员可以不用执行cancle()方法，但是还是建议显示的执行cancel()方法

	// 此外，⚠️ 超时后的自动取消 不会立即释放底层资源，只有调用 cancel() 才会真正清理。 "Code should call cancel even if the context is not needed any more, to avoid context leak."


}

func test03() {
	// 3. context.WithTimeout：自动取消
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancelFunc()

	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("完成任务")
		case <-ctx.Done():
			fmt.Println("超时了:", ctx.Err())
			os.Exit(-1)
		default:
			fmt.Println("工作中...")
			time.Sleep(500 * time.Millisecond)
		}
	}



}

func test02() {
	// 2. context.WithCancel：手动取消 goroutine
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine 退出:", ctx.Err())
				return
			default:
				fmt.Println("工作中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancelFunc() // 手动取消，当调用cancelFunc之后，该cancelFunc对应的ctx对应的ctx.Done()方法将被触发，此时去监控他的输出信道，可以取消该协程的执行
	time.Sleep(1 * time.Second)

	fmt.Println("主协程退出！！！")


}

func test01() {
	// 1. context.Background() & context.TODO()

	ctx1 := context.Background()
	ctx2 := context.TODO()

	fmt.Printf("ctx1: %v, %#v \n", ctx1, ctx1)
	fmt.Printf("ctx2: %v, %#v \n", ctx2, ctx2)
	fmt.Println()
}
