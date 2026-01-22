package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudflare/tableflip"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// import "github.com/cloudflare/tableflip"
// Package tableflip implements zero downtime upgrades
// 零停机（zero downtime）重启 / 升级, 在不关闭监听端口、不中断已有连接的情况下，重启 Go 服务进程

type AppServer struct {
	ctx context.Context
	//mux    *http.ServeMux
	engine *gin.Engine
	server *http.Server
}

func NewAppServer(ctx context.Context) *AppServer {
	//mux := http.NewServeMux()
	engine := gin.Default()
	return &AppServer{
		ctx,
		engine,
		&http.Server{Handler: engine},
	}
}

func (s *AppServer) registerRouter() {
	// http://127.0.0.1:9090/version
	//mux := http.NewServeMux()
	//mux.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
	//	for i := 0; i < 15; i++ {
	//		log.Printf("Do task %d version = %s \n", i, Version)
	//		time.Sleep(1 * time.Second)
	//	}
	//})
	s.engine.GET("/version", func(c *gin.Context) {
		for i := 0; i < 15; i++ {
			log.Printf("Do task %d version = %s \n", i, Version)
			time.Sleep(1 * time.Second)
		}
	})
	log.Println("Route registered")
}

func (s *AppServer) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}

func (s *AppServer) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down")
	return s.server.Shutdown(ctx)
}

var Version = "dev" // 默认值，可被 -ldflags 覆盖

// go build -ldflags "-X main.Version=v1" -o app

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// 1. 创建 Upgrader
	upg, err := tableflip.New(tableflip.Options{})
	if err != nil {
		log.Fatalf("failed to create upgrader: %v", err)
	}
	defer upg.Stop() // 确保退出时清理

	// 监听系统的 SIGHUP 信号，以此信号触发进程重启
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			// 核心的 Upgrade 调用
			err := upg.Upgrade()
			if err != nil {
				log.Println("Upgrade failed:", err)
			}
		}
	}()

	// 2. 创建 Listener
	ln, err := upg.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 3. 创建服务
	server := NewAppServer(context.Background())
	server.registerRouter()

	// 4. 监听退出信号（优雅退出）
	go func() {
		<-upg.Exit()
		log.Println("graceful shutdown start")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
		log.Println("graceful shutdown complete")
	}()

	// 5. 启动服务
	go func() {
		if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server serve error: %v", err)
		}
	}()

	// 6. 通知 tableflip 服务已就绪
	if err := upg.Ready(); err != nil {
		log.Fatalf("failed to ready: %v", err)
	}
	time.Sleep(time.Second)
	log.Printf("server running at :%s with PID %d, version: %s", port, os.Getpid(), Version)

	// 7. 捕获系统信号，优雅退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("shutting down by signal")
}

/*验证tableflip能否完成优雅更新
1、编译app-v1：go build -ldflags "-X main.Version=v1" -o app-v1，保存好PID：84987
2、启动app-v1: ./app-v1
3、编译新版本：go build -ldflags "-X main.Version=v22" -o app-v22
4、执行接口：curl http://localhost:9090/version
5、向旧进程发送升级信号：kill -s HUP 84987
5、再次执行接口：curl http://localhost:9090/version，观察version是否变为v22

*/
