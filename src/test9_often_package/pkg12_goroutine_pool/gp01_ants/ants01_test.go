package gp01_ants

import (
	"fmt"
	"github.com/panjf2000/ants"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	pool, _ := ants.NewPool(10) // 最大10个协程
	defer pool.Release()

	for i := 0; i < 100; i++ {
		_ = pool.Submit(func() {
			fmt.Println("run task", i)
			time.Sleep(100 * time.Millisecond)
		})
	}
}
