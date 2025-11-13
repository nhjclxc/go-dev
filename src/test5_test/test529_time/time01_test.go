package main

import (
	"fmt"
	"testing"
	"time"
)

var cst *time.Location

func init() {
	var err error
	cst, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		cst = time.FixedZone("CST", 8*60*60)
	}
}

func Test111(t *testing.T) {

	//func (t Time) Truncate(d Duration) Time
	// 功能：
	//将时间 t 向下舍入到 指定的时间间隔 d。
	//舍入到最接近的、整倍数的 d 时间点。
	//返回一个新的 time.Time 对象（原来的 t 不变）

	tt := time.Date(2025, 11, 12, 10, 58, 0, 0, cst)
	fmt.Println("origin time", tt)
	fmt.Println(tt.Truncate(5 * time.Minute))
	fmt.Println(tt.Truncate(10 * time.Minute))
	fmt.Println(tt.Truncate(30 * time.Minute))
	fmt.Println(tt.Truncate(60 * time.Minute))

	// 如果是0-5的舍入0、5-10的舍入5、10-15的舍入10，可以使用.Truncate(5 * time.Minute)
	// 如果是2:30-7:30的舍入5、7:30-12:30的舍入10、12:30-17:30的舍入15，又该如何呢？
}
