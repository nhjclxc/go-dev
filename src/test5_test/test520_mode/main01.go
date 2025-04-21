package main

// 17.1 逗号 ok 模式
func main01() {

	//在学习本书第二部分和第三部分时，我们经常在一个表达式返回 2 个参数时使用这种模式：something, ok，
	//第一个参数是一个值或者 nil，第二个参数是 true/false 或者一个错误 error。在一个需要赋值的 if 条件语句中，
	//使用这种模式去检测第二个参数值会让代码显得优雅简.

	/*
		（1）在函数返回时检测错误（参考第 5.2 小节）:

		value, err := pack1.Func1(param1)

		if err != nil {
		    fmt.Printf("Error %s in pack1.Func1 with parameter %v", err.Error(), param1)
		    return err
		}

		// 函数Func1没有错误:
		Process(value)

		e.g.: os.Open(file) strconv.Atoi(str)
	*/

	/*
		（2）检测映射中是否存在一个键值（参考第 8.2 小节）：key1 在映射 map1 中是否有值？

		if value, isPresent = map1[key1]; isPresent {
		        Process(value)
		}
		// key1 不存在
		…
	*/

	/*
		（3）检测一个接口类型变量 varI 是否包含了类型 T：类型断言（参考第 11.3 小节）：

		if value, ok := varI.(T); ok {
		    Process(value)
		}
		// 接口类型 varI 没有包含类型 T
	*/

	/*（4）检测一个通道 ch 是否关闭（参考第 14.3 小节）：

	    for input := range ch {
	        Process(input)
	    }
	或者:

	    for {
	        if input, open := <-ch; !open {
	            break // 通道是关闭的
	        }
	        Process(input)
	    }
	*/

}
