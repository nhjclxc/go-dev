package app

import (
	"context"
	"fmt"
	"otfen_pkg_09_spf13_cobra/config"
	"sync"
	"time"
)

type Application struct {
	ctx    context.Context
	cancel context.CancelFunc
	cfg    *config.Config
	wg     sync.WaitGroup
}

func NewApplication(ctx context.Context, cancel context.CancelFunc, cfg *config.Config) (*Application, error) {
	return &Application{
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}, nil
}

func (a *Application) StartApp() {
	// 初始化应用
	//a.initCompose()

	// 启动应用程序
	fmt.Printf("应用程序启动中. \n")
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("应用程序启动中.. \n")
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("应用程序启动中... \n")
	time.Sleep(500 * time.Millisecond)

	numWorkers := 3
	a.wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		// worker就是每一个具体的服务
		go worker(a.ctx, &a.wg, i)
	}

	fmt.Printf("应用程序启动成功！！！\n")

	fmt.Printf("配置文件：%#v \n", a.cfg)

}

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d 启动\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d 接收到退出信号，着手停止工作\n", id)
			time.Sleep(1 * time.Second) // 模拟停止执行
			fmt.Printf("Worker %d 已停止工作！\n", id)
			return
		default:
			fmt.Printf("Worker %d 正在工作...\n", id)
			time.Sleep(1 * time.Second) // 模拟任务执行
		}
	}
}

func (a *Application) StopApp() {
	fmt.Printf("优雅关闭服务")

	a.cancel()
	a.wg.Wait()
	fmt.Println("所有 worker 已优雅退出，程序结束")
}
