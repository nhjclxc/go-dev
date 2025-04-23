package main

import "errors"

func main() {
	/*
		   	1.1.1. Go怎么写测试用例
		      开发程序其中很重要的一点是测试，我们如何保证代码的质量，如何保证每个函数是可运行，运行结果是正确的，又如何保证写出来的代码性能是好的，
		   	我们知道单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，
		   	让线上的程序能够在高并发的情况下还能保持稳定。本小节将带着这一连串的问题来讲解Go语言中如何来实现单元测试和性能测试。

		      Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，
		   	你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例，那么接下来让我们一一来看一下怎么写。

			另外建议安装gotests插件自动生成测试代码:go get -u -v github.com/cweill/gotests/
	*/

	/*
			由于go test命令只能在一个相应的目录下执行所有文件，所以我们接下来新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下。

			接下来我们在该目录下面创建两个文件：gotest.go和gotest_test.go

			gotest.go:这个文件里面我们是创建了一个包，里面有一个函数实现了除法运算:


		gotest_test.go:这是我们的单元测试文件，但是记住下面的这些原则：

		文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码

		你必须import testing这个包

		所有的测试用例函数必须是Test开头

		测试用例会按照源代码中写的顺序依次执行

		测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态

		测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。

		函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。


	*/

	/*
		压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

		压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

		    func BenchmarkXXX(b *testing.B) { ... }
		go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

		在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行 文件名也必须以_test.go结尾

		下面我们新建一个压力测试文件webbench_test.go，代码如下所示：


	*/

	/*
		ns/op 是什么意思？
		这个是 Go benchmark 的标准输出格式之一，含义是：

		每次操作（op）平均耗时 0.3084 纳秒（ns = 纳秒 = 10^-9 秒）。

		这个值越小，表示 每一次操作执行得越快，性能越好。
	*/
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}
