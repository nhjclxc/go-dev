package main

import "fmt"

func main() {

	// 使用生成器实现斐波那契数列
	// fib(n) = fib(n-1) + fib(n-2)

	fibGen := NewFibGenerate()

	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())
	fmt.Println(fibGen.generateFib())

}

type FibGenerate struct {
	resume chan int
}

func NewFibGenerate() *FibGenerate {

	fg := &FibGenerate{
		resume: make(chan int),
	}

	n1, n2 := 0, 1

	counter := 0

	// fib(n) = fib(n-1) + fib(n-2)

	go func() {
		for {
			if counter == 0 {
				fg.resume <- 0
			} else if counter == 1 {
				fg.resume <- 1
			} else {
				n1, n2 = n2, n1+n2
				fg.resume <- n2
			}
			counter++
		}
	}()

	return fg
}

func (this *FibGenerate) generateFib() int {
	return <-this.resume
}

// ✅ 可选拓展方向（根据你的需求）
//GenerateFibN(n int)：一次性获取前 N 个 Fibonacci 数
//
//Reset()：重置生成器
//
//Close()：优雅关闭生成器
//
//context.Context 控制取消
//
//如果你想要这些扩展功能，我可以继续帮你写～你有兴趣加一个 停止生成 的功能吗？比如生成 20 个后就不继续？
