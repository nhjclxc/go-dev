package main

import "fmt"

func main02() {
	/*
		// 封装更佳通用的生成器的方法
			这些原则可以概括为：通过巧妙地使用空接口、闭包和高阶函数，我们能实现一个通用的惰性生产器的工厂函数 BuildLazyEvaluator（这个应该放在一个工具包中实现）。
			工厂函数需要一个函数和一个初始状态作为输入参数，返回一个无参、返回值是生成序列的函数。传入的函数需要计算出下一个返回值以及下一个状态参数。
			在工厂函数中，创建一个通道和无限循环的 go 协程。返回值被放到了该通道中，返回函数稍后被调用时从该通道中取得该返回值。每当取得一个值时，下一个值即被计算。

			在下面的例子中，定义了一个 evenFunc 函数，其是一个惰性生成函数：在 main() 函数中，我们创建了前 10 个偶数，每个都是通过调用 even() 函数取得下一个值的。
			为此，我们需要在 BuildLazyIntEvaluator 函数中具体化我们的生成函数，然后我们能够基于此做出定义。
	*/

	eg := NewEvenGenerate()
	fmt.Println(eg.generaterEven())
	fmt.Println(eg.generaterEven())
	fmt.Println(eg.generaterEven())
	fmt.Println(eg.generaterEven())
	fmt.Println(eg.generaterEven())
	fmt.Println(eg.generaterEven())

	fmt.Println("-----------------------------")

	evenFunc := func(eg *Generate, initState int) {

		currState := initState
		for {
			eg.resume <- currState
			currState += 2
		}
	}

	evenGen := NewGenerate(evenFunc, 0)

	fmt.Println(evenGen.generater())
	fmt.Println(evenGen.generater())
	fmt.Println(evenGen.generater())
	fmt.Println(evenGen.generater())
	fmt.Println(evenGen.generater())
	fmt.Println(evenGen.generater())

	fmt.Println("-----------------------------")

	fiveFunc := func(eg *Generate, initState int) {

		currState := initState
		for {
			eg.resume <- currState
			currState += 5
		}
	}

	fiveGen := NewGenerate(fiveFunc, 3)

	fmt.Println(fiveGen.generater())
	fmt.Println(fiveGen.generater())
	fmt.Println(fiveGen.generater())
	fmt.Println(fiveGen.generater())
	fmt.Println(fiveGen.generater())
	fmt.Println(fiveGen.generater())

}

type EvenGenerate struct {
	resume chan int
}

func NewEvenGenerate() EvenGenerate {

	eg := EvenGenerate{
		resume: make(chan int),
	}

	counter := 0

	go func() {
		for {
			eg.resume <- counter
			counter++
		}
	}()

	return eg
}

func (this *EvenGenerate) generaterEven() int {
	return <-this.resume
}

// 定义通用的生成器

type Generate struct {
	resume chan int
}

func NewGenerate(optFunc func(*Generate, int), initState int) *Generate {

	eg := &Generate{
		resume: make(chan int),
	}

	go optFunc(eg, initState)

	return eg
}

func (this *Generate) generater() int {
	return <-this.resume
}
