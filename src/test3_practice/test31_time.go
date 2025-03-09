package main

import (
	"fmt"
	"time"
)

/*
*

	golang包time用法详解
*/
func main() {

	//test31_1()

	//test31_2()

	//test31_3()

	//test31_4()

	//test31_5()

	test31_6()

}

func test31_6() {
	/*
		6. 获取一年中的某个信息
		方法	作用
		t.Year()	获取年份
		t.Month()	获取月份
		t.Day()		获取天
		t.Weekday()	获取星期
		t.YearDay()	获取一年中的第几天
	*/

	t := time.Now()
	fmt.Println(t.Year())
	fmt.Println(t.Month(), int(t.Month()))
	fmt.Println(t.Day())
	fmt.Println(t.Weekday(), int(t.Weekday()))
	fmt.Println(t.YearDay())

	println(getDaysInMonth(2025, 2))

}

/*
*
获取某年某月的天数
*/
func getDaysInMonth(year, month int) int {
	// 创建该月的最后一天的时间
	location := time.Local // 可以根据需要修改时区
	lastDay := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, location)
	// 返回该月的天数
	return lastDay.Day()
}

func test31_5() {
	/*
		5. 定时器 & 睡眠
		方法	作用
		time.Sleep(d)	休眠 d 时间
		time.After(d)	d 时间后返回数据
		time.NewTimer(d)	创建定时器
		time.Tick(d)	每隔 d 时间触发
	*/

	fmt.Println(time.Now().Unix())
	time.Sleep(time.Second)
	fmt.Println(time.Now().Unix())
	// 由上可知，time.Now().Unix()表示的是秒的时间戳

	//// 定时器
	//timer := time.NewTimer(3 * time.Second)
	//<-timer.C
	//fmt.Println("3 秒后触发:", time.Now())
	//
	//// 周期性触发
	//ticker := time.NewTicker(1 * time.Second)
	//for i := 0; i < 3; i++ {
	//	fmt.Println("Tick:", <-ticker.C)
	//}
	//ticker.Stop()

}

func test31_4() {
	/*
		4. 时间戳
		方法	作用
		t.Unix()	返回秒级时间戳
		t.UnixNano()	返回纳秒级时间戳
	*/

	// 时间转时间戳
	t := time.Now()
	nuix := t.Unix()
	fmt.Println(nuix)
	fmt.Println(t.UnixNano())

	// 时间戳转时间
	// 在 Go 语言的 time.Unix(sec int64, nsec int64) 方法中，第二个参数 nsec 代表纳秒（nanoseconds）偏移量，用于增加额外的时间精度。
	fmt.Println(time.Unix(nuix, 0))
	fmt.Println(time.Unix(nuix, 1))
	fmt.Println(time.Unix(nuix, 2))
	fmt.Println(time.Unix(nuix, 2*1000))
	fmt.Println(time.Unix(nuix, 2*1000*2))
	//秒 (second) = 1,000 毫秒 (millisecond)
	//1 毫秒 (millisecond) = 1,000,000 纳秒 (nanoseconds)
	//1 秒 (second) = 1,000,000,000 纳秒 (nanoseconds)
	fmt.Println(time.Unix(nuix, 1000000))
	fmt.Println(time.Unix(nuix, 1000000*1000))
	fmt.Println(time.Unix(nuix, 1000000*1000*60))

}

func test31_3() {

	/*
		3. 时间计算
		方法	作用
		t.Add(d)	增加时间
		t.Sub(t2)	计算时间间隔
		t.Before(t2)	是否在某时间之前
		t.After(t2)	是否在某时间之后
		t.Equal(t2)	时间是否相等
	*/
	// 以下就是Duration，表示毫秒，秒，分组，一小时和一天
	fmt.Println("", time.Millisecond)
	fmt.Println("", time.Second)
	fmt.Println("", time.Minute)
	fmt.Println("", time.Hour)
	fmt.Println("", time.Hour*24)

	var time1 = time.Date(2025, 3, 9, 10, 16, 11, 22, time.Local)
	var time2 = time.Date(2026, 3, 9, 10, 16, 11, 22, time.Local)
	fmt.Println("time1 = ", time1)
	fmt.Println("time2 = ", time2)
	fmt.Println(time1.Add(time.Minute * 2))
	fmt.Println(time1.Add(time.Hour * 2))
	time11 := time1.Add(time.Hour*2 + time.Minute*2)
	fmt.Println(time11)
	time11Subtime1 := time11.Sub(time1)
	fmt.Println(time11Subtime1)           // 时间间隔
	fmt.Println(time11Subtime1.Hours())   // 总的间隔小时
	fmt.Println(time11Subtime1.Minutes()) // 总的间隔分钟
	fmt.Println(time11Subtime1.Seconds()) // 总的间隔秒

	time111 := time11.Add(-time.Hour*2 - time.Minute*2)
	fmt.Println(time1.After(time11))  //time1是否在time11之后，否返回false
	fmt.Println(time1.Before(time11)) // time1是否在time11之前，是返回true
	fmt.Println(time1.Equal(time111)) // time1和time111的年月日时分秒是否全部相等，相等返回true
	fmt.Println(time11.Equal(time111))

}

func test31_2() {

	/*
				2. 时间格式化 & 解析
				方法	作用
				t.Format(layout)	格式化时间
				time.Parse(layout, str)	解析时间字符串
				time.ParseInLocation(layout, str, loc)	解析时区时间

		Go 设计上强制规定了这个时间格式,	layout 不是占位符，而是一个 时间实例，必须按照 2006-01-02 15:04:05 的结构来编写，否则 time.Parse() 解析会失败。
			Go 语言 必须 使用 "2006-01-02 15:04:05" 作为 layout，
	*/
	now := time.Now()
	// 时间格式化
	fmt.Println("YYYY-MM-DD:", now.Format("2006-01-02"))
	fmt.Println("YYYY/MM/DD HH:MM:SS:", now.Format("2006/01/02 15:04:05"))
	fmt.Println("YYYY/MM/DD HH:MM:SS:", now.Format("06/01/02 15:04:05"))
	fmt.Println("YYYY/MM/DD HH:MM:SS:", now.Format("2003/01/02 15:04:05"))
	fmt.Println("YYYY/MM/DD HH:MM:SS:", now.Format("2009/01/02 15:04:05"))

	layout := "2006年01月02号，15点04分05"
	timeStr := "2025年03月09号，10点33分06"
	t, err := time.Parse(layout, timeStr)
	fmt.Println(t)
	fmt.Println(err)
	layout2 := "2006年01月02号，15点"
	timeStr2 := "2025年03月09号，10点"
	t2, err2 := time.Parse(layout2, timeStr2)
	fmt.Println(t2)
	fmt.Println(err2)

}

func test31_1() {
	/*
		1. 时间获取
		方法	作用
		time.Now()	获取当前时间
		time.Date(year, month, day, hour, min, sec, nsec, loc)	创建指定时间
		time.Parse(layout, str)	解析时间字符串
		time.Unix(sec, nsec)	从 Unix 时间戳创建时间
	*/
	now := time.Now()
	fmt.Println(now)

	var nowTime = time.Date(2025, 3, 9, 10, 16, 11, 22, time.Local)
	fmt.Println(nowTime)

	// 从时间戳创建时间
	timestamp := time.Unix(1709875200, 0)
	fmt.Println("Unix 时间:", timestamp)

}
