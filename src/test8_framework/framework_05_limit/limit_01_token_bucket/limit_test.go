package main

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	for i := 0; i < 30; i++ {

		limiter := getLimiterForIP("1.7.0.0.1")

		if !limiter.Allow() {
			fmt.Printf("Too many requests. Please try again later. \n")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		fmt.Printf("Successful !!! \n")

	}
}
