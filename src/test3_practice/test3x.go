package main

import (
	"fmt"
	"time"
)

/*
*

https://studygolang.com/articles/30412
https://studygolang.com/user/haicoder_ibai
*/
func main() {

	//test3x_1()

	//test3x_2()

	//test3x_3()

	//test3x_4()

	//test3x_5()

	//test3x_6()

	//test3x_7()

}

/*
题目十一
描述：用 Golang 实现，求解兔子总数问题。

题目：古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
*/
func test3x_7() {
	// 经典的动态规划问题
	// 可以使用递归做吗？？？
	// 是不是斐波那契数列问题？？？

	println(f(1))
	println(f(2))
	println(f(3))
	println(f(4))
	println(f(5))
	println(f(6))
	println(f(7))
	println(f(8))
}

func f(month int) int {
	if month == 1 || month == 2 {
		return 2
	}
	var previous, current int = 2, 2
	for i := 3; i <= month; i++ {
		previous, current = current, current+previous
	}

	return current
}

/*
题目八
描述：用 Golang 实现，打印乘法口诀。

题目：输出 9*9 乘法口诀表。
*/
func test3x_6() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d, ", j, i, i*j)
		}
		println()
	}

}

/*
题目五
描述：用 Golang 实现，数字从小到大排序。

题目：输入三个 整数 x，y，z，请把这三个数由小到大输出。
*/
func test3x_5() {
	var num1 = 113
	var num2 = 11
	var num3 = 18

	var arr []int = []int{num1, num2, num3}

	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	//fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	//fmt.Println(arr)

}

/*
*
描述：用 Golang 实现，随便输入日期，判断时一年的第几天。

题目：输入某年某月某日，判断这一天是这一年的第几天？
*/
func test3x_4() {

	var inputTime = "2025/03/04"
	layout := "2006/01/02"

	parseTime, _ := time.Parse(layout, inputTime)
	fmt.Println(parseTime.YearDay()) // 直接使用系统定义的方法

	var year = parseTime.Year()
	var sum int = 0
	var monthDay = int(parseTime.Month()) // 第几个月
	// 计算前面所有月份的天数
	for i := 1; i < monthDay; i++ {
		sum += getDaysInMonth1(year, i)
	}
	// 计算最后一个月的天数
	sum += parseTime.Day()
	fmt.Println(sum)

}

/*
*
获取某年某月的天数
*/
func getDaysInMonth1(year, month int) int {
	// 创建该月的最后一天的时间
	location := time.Local // 可以根据需要修改时区
	lastDay := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, location)
	// 返回该月的天数
	return lastDay.Day()
}

/*
*
描述：用 Golang 实现，计算完全平方数。

题目：一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
*/
func test3x_3() {
	// 完全平方数（Perfect Square）是指能够写成某个整数的平方的数，即：n = m^2，则n为完全平方数1,4,9,16,25,36

	//for i := 0; ; i++ {
	//	for j := 0; ; j++ {
	//		num := (i + j) * (i - j)
	//		if num == 168 {
	//			fmt.Printf("i = %d, j = %d", i, j)
	//			return
	//		}
	//		if num > 168 {
	//			break
	//		}
	//	}
	//	if i > 168 {
	//		break
	//	}
	//	println("i = ", i)
	//}

	//for num := 0; ; num++ {
	//	num1 := num + 100
	//
	//	for i := 0; ; i++ {
	//		if i*i == num1 {
	//			fmt.Printf("num = %d, num*num = %d, i = %d, i*i = %d", num, num1, i, i*i)
	//
	//			num2 := num1 + 168
	//
	//			for j := 0; ; j++ {
	//				if j*j == num2 {
	//					fmt.Printf("num = %d, num*num = %d, j = %d, j*j = %d", num, num2, j, j*j)
	//					return
	//				}
	//			}
	//		}
	//	}

	//num2 := num1 + 168
	//
	//for i := 0; ; i++ {
	//	if i*i == num1 {
	//
	//	}
	//}

	//}

}

/*
*
描述：用 Golang 实现，企业发放的奖金根据利润提成的计算。

题目：

	 企业发放的奖金根据利润提成。
		利润(I)低于或等于 10 万元时，奖金可提成 10%；
		高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%;
		20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；
		40 万到 60 万之间时高于 40 万元的部分，可提成 3%；
		60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%;
		高于 100 万元时，超过 100 万元的部分按 1% 提成。

	 从键盘输入当月利润 I，求应发放奖金总数？
*/
func test3x_2() {
	var i float32 = 1200000.0
	var ticheng float32

	switch {
	case i <= 100000.0:
		ticheng = i * 0.1

	case 100000.0 < i && i <= 200000.0:
		ticheng = 100000.0*0.1 + (i-100000.0)*0.075

	case 200000.0 < i && i <= 400000.0:
		ticheng = (i - 200000.0) * 0.05

	case 400000.0 < i && i <= 600000.0:
		ticheng = (i - 400000.0) * 0.03

	case 600000.0 < i && i <= 1000000.0:
		ticheng = (i - 600000.0) * 0.015
	case i > 1000000.0:
		ticheng = (i - 1000000.0) * 0.01

	}

	fmt.Printf("i = %f, ticheng = %f", i, ticheng)

}

/*
*
题目一
描述：用 Golang 实现，将四个数进行排列组合。

题目：有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
*/
func test3x_1() {

	var sum int = 0

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 1; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				sum++
				fmt.Printf("sum = %d, value = %d%d%d \n", sum, i, j, k)
			}
		}
	}
	println("sum = ", sum)

}
