package main

// https://denganliang.github.io/the-way-to-go_ZH_CN/18.8.html
func main() {
	/*
	   	18.8 协程 (goroutine) 与通道 (channel)
	      出于性能考虑的建议：

	      实践经验表明，为了使并行运算获得高于串行运算的效率，在协程内部完成的工作量，必须远远高于协程的创建和相互来回通信的开销。

	      出于性能考虑建议使用带缓存的通道：

	      使用带缓存的通道可以很轻易成倍提高它的吞吐量，某些场景其性能可以提高至 10 倍甚至更多。通过调整通道的容量，甚至可以尝试着更进一步的优化其性能。

	      限制一个通道的数据数量并将它们封装成一个数组：

	      如果使用通道传递大量单独的数据，那么通道将变成性能瓶颈。然而，将数据块打包封装成数组，在接收端解压数据时，性能可以提高至 10 倍。

	      现在创建一个带缓存的通道：ch := make(chan type,buf)

	      （1）如何使用 for 或者 for-range 遍历一个通道：

	      for v := range ch {
	          // do something with v
	      }
	      （2）如何检测一个通道 ch 是否关闭：

	      //read channel until it closes or error-condition
	      for {
	          if input, open := <-ch; !open {
	              break
	          }
	          fmt.Printf("%s", input)
	      }
	      或者使用（1）自动检测。

	      （3）如何通过一个通道让主程序等待直到协程完成（信号量模式）：

	      ch := make(chan int) // Allocate a channel.
	      // Start something in a goroutine; when it completes, signal on the channel.
	      go func() {
	          // doSomething
	          ch <- 1 // Send a signal; value does not matter.
	      }()
	      doSomethingElseForAWhile()
	      <-ch // Wait for goroutine to finish; discard sent value.
	      如果希望程序一直阻塞，在匿名函数中省略 ch <- 1 即可。

	      （4）通道的工厂模板：以下函数是一个通道工厂，启动一个匿名函数作为协程以生产通道：

	      func pump() chan int {
	          ch := make(chan int)
	          go func() {
	              for i := 0; ; i++ {
	                  ch <- i
	              }
	          }()
	          return ch
	      }
	      （5）通道迭代器模板：（注：这里原书没有写东西，但是应该是参考章节 14.2.10）

	      （6）如何限制并发处理请求的数量：参考章节 14.11

	      （7）如何在多核CPU上实现并行计算：参考章节 14.13

	      （8）如何终止一个协程：runtime.Goexit()

	      （9）简单的超时模板：

	      timeout := make(chan bool, 1)
	      go func() {
	          time.Sleep(1e9) // one second
	          timeout <- true
	      }()
	      select {
	          case <-ch:
	          // a read from ch has occurred
	          case <-timeout:
	          // the read from ch has timed out
	      }
	      （10）如何使用输入通道和输出通道代替锁：

	      func Worker(in, out chan *Task) {
	          for {
	              t := <-in
	              process(t)
	              out <- t
	          }
	      }
	      （11）如何在同步调用运行时间过长时将之丢弃：参考章节 14.5 第二个变体

	      （12）如何在通道中使用计时器和定时器：参考章节 14.5

	      （13）典型的服务器后端模型：参考章节 14.4
	*/
}
