package main

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("炸煎鱼: ", t.Unix())
		}
	}
}

func Test2(t2 *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("炸煎鱼: ", t.Unix())
		}
	}
}

func Test3(t3 *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for t := range ticker.C {
		fmt.Println("炸煎鱼: ", t.Unix())
	}
}
