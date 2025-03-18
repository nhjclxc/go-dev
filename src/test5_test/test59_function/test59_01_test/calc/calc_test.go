package calc

import "testing"

// 以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
//func main() {
/*

	在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
		类型			格式					作用
		测试函数	函数名前缀为Test		测试程序的一些逻辑行为是否正确
		基准函数	函数名前缀为Benchmark	测试函数的性能
		示例函数	函数名前缀为Example	为文档提供示例文档
*/
/*
	Golang单元测试对文件名和方法名，参数都有很严格的要求。
		1、文件名必须以xx_test.go命名
	    2、方法必须是Test[^a-z]开头
	    3、方法参数必须 t *testing.T
	    4、使用go test执行单元测试
*/
/*
		import "testing"
		每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：

		func TestName(t *testing.T){
		    // ...
		}

	其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：
		func (c *T) Error(args ...interface{})
		func (c *T) Errorf(format string, args ...interface{})
		func (c *T) Fail()
		func (c *T) FailNow()
		func (c *T) Failed() bool
		func (c *T) Fatal(args ...interface{})
		func (c *T) Fatalf(format string, args ...interface{})
		func (c *T) Log(args ...interface{})
		func (c *T) Logf(format string, args ...interface{})
		func (c *T) Name() string
		func (t *T) Parallel()
		func (t *T) Run(name string, f func(t *T)) bool
		func (c *T) Skip(args ...interface{})
		func (c *T) SkipNow()
		func (c *T) Skipf(format string, args ...interface{})
		func (c *T) Skipped() bool
*/
//}

// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
func TestAdd(t *testing.T) {
	num1 := 333
	num2 := 222
	got := Add(num1, num2)
	want := 333 + 222
	//want := 1111
	if got != want {
		t.Errorf("Add函数异常，got = %d, want =%d", got, want)
	}

}
func TestSub(t *testing.T) {
	num1 := 333
	num2 := 222
	got := Sub(num1, num2)
	want := 888
	//want := 1111
	if got != want {
		t.Errorf("Sub函数异常，got = %d, want =%d", got, want)
	}
}
func TestLog(t *testing.T) {

}
