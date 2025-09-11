package main

import (
	"fmt"
	"testing"
	"time"
)

// time.NewTimer和time.NewTicker的区别
// timer是延迟器：延迟多少秒执行，仅执行一次，也就是睡眠n秒后执行
// ticker是定时器：每个多少秒执行，定时器不关闭则一直间隔n秒执行
func Test1(t *testing.T) {

	ti := time.NewTimer(3 * time.Second)

	fmt.Println("111 ", time.Now())
	<-ti.C
	fmt.Println("222 ", time.Now())
}

func Test2(t *testing.T) {

	ti := time.NewTicker(3 * time.Second)
	defer ti.Stop()

	for {
		select {
		case t := <-ti.C:
			fmt.Println("发送心跳:", t)
		}
	}
}
