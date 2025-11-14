package examples

import (
	"context"
	"fmt"
	"github.com/yourorg/cdn-api-gateway/handler"
	common "github.com/yourorg/cdn-common"
	"github.com/yourorg/cdn-pay-service/pkg/pay"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	fmt.Println(pay.PayStatusPaid)
	fmt.Println(pay.PayStatusProcessing)
	fmt.Println(pay.PayStatusCanceled)
}

func Test222(t *testing.T) {
	// 创建 PayService 实例，注入实现
	var processor common.PaymentProcessor = pay.NewAlipayPayProcessor()

	// 注入到 API Handler
	payHandler := handler.NewPayHandler(processor)

	// 使用
	err := payHandler.ChargeUser("user123", 100)
	if err != nil {
		panic(err)
	}
}

func Test333(t *testing.T) {
	// 创建 PayService 实例，注入实现
	var processor common.PaymentProcessor = pay.NewWeChatPayProcessor()

	// 注入到 API Handler
	payHandler := handler.NewPayHandler(processor)

	// 使用
	err := payHandler.ChargeUser("user123", 100)
	if err != nil {
		panic(err)
	}
}

func Test123(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ch1 chan string = make(chan string)
	var ch2 chan string = make(chan string)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		n := r.Intn(10) + 1
		fmt.Println("ch1 rand ", n)
		time.Sleep(time.Duration(n) * time.Second)
		ch1 <- "ch1 sleep end!"
	}()
	go func() {
		n := r.Intn(10) + 1
		fmt.Println("ch2 rand ", n)
		time.Sleep(time.Duration(n) * time.Second)
		ch2 <- "ch2 sleep end!"
	}()
	go func() {
		n := r.Intn(10) + 1
		fmt.Println("cancel rand ", n)
		time.Sleep(time.Duration(n) * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("程序关闭！")
	case data, ok := <-ch1:
		if ok {
			fmt.Println("received ch1 data: ", data)
		}
	case data, ok := <-ch2:
		if ok {
			fmt.Println("received ch2 data: ", data)
		}
	}

}
