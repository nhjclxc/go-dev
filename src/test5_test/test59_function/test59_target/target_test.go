package test59_target

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	count := 0
LOOP:
	for {
		count++
		fmt.Println("for ", count)
		time.Sleep(1 * time.Second)

		if count > 5 {
			break LOOP // 跳出到 LOOP 标签处
		}
	}

	fmt.Println("LOOP")

}

func Test02(t *testing.T) {
	count := 0
	for {
		count++
		fmt.Println("for ", count)
		time.Sleep(1 * time.Second)

		if count > 5 {
			goto END
		}
	}

END:
	fmt.Println("退出for")
	fmt.Println("退出for 2")
}
