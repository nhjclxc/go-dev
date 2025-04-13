package main

func main() {
	/*
		time 包中有一些有趣的功能可以和通道组合使用。
		其中就包含了 time.Ticker 结构体，这个对象以指定的时间间隔重复的向通道 C 发送时间值：

		type Ticker struct {
		C <-chan Time // the channel on which the ticks are delivered.
		// contains filtered or unexported fields
		...
		}
		时间间隔的单位是 ns（纳秒，int64），在工厂函数 time.NewTicker 中以 Duration 类型的参数传入：func NewTicker(dur) *Ticker。

		在协程周期性的执行一些事情（打印状态日志，输出，计算等等）的时候非常有用。
		time.Tick() 函数声明为 Tick(d Duration) <-chan Time，当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d 是纳秒数。
		如果需要，像下边的代码一样，可以限制处理频率（函数 client.Call() 是一个 RPC 调用，这里暂不赘述
		rate_per_sec := 10
		var dur time.Duration = 1e9 // rate_per_sec
		chRate := time.Tick(dur) // a tick every 1/10th of a second
		for req := range requests {
		<- chRate // rate limit our Service.Method RPC calls
		go client.Call("Service.Method", req, ...)
		}
	*/

}
