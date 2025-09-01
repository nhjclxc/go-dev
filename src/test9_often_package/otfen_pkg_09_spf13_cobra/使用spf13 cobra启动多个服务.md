很好 👍
`github.com/spf13/cobra` 就是 Go 里面做 **多子命令 CLI 程序**的标准库，常见的 kubectl、helm、etcdctl 都是用它写的。

我给你写一个完整的示例：既可以单独启动 `serviceA` 或 `serviceB`，也可以用 `all` 一次性并行启动两个服务。

---

## 示例：cobra 多子命令服务启动器

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func runService(name string, port int) {
	fmt.Printf("启动 %s 服务，监听端口: %d\n", name, port)
	// 这里写具体服务逻辑，比如 http.ListenAndServe(...)
	select {}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "多服务启动器",
	}

	// serviceA
	var portA int
	var serviceACmd = &cobra.Command{
		Use:   "serviceA",
		Short: "启动服务 A",
		Run: func(cmd *cobra.Command, args []string) {
			runService("serviceA", portA)
		},
	}
	serviceACmd.Flags().IntVar(&portA, "port", 8080, "服务 A 端口")

	// serviceB
	var portB int
	var serviceBCmd = &cobra.Command{
		Use:   "serviceB",
		Short: "启动服务 B",
		Run: func(cmd *cobra.Command, args []string) {
			runService("serviceB", portB)
		},
	}
	serviceBCmd.Flags().IntVar(&portB, "port", 9090, "服务 B 端口")

	// all: 同时启动 A 和 B
	var allCmd = &cobra.Command{
		Use:   "all",
		Short: "同时启动 serviceA 和 serviceB",
		Run: func(cmd *cobra.Command, args []string) {
			go runService("serviceA", portA)
			go runService("serviceB", portB)

			// 捕获退出信号，保持阻塞
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
			<-ch
		},
	}
	allCmd.Flags().IntVar(&portA, "portA", 8080, "服务 A 端口")
	allCmd.Flags().IntVar(&portB, "portB", 9090, "服务 B 端口")

	// 注册子命令
	rootCmd.AddCommand(serviceACmd, serviceBCmd, allCmd)
	_ = rootCmd.Execute()
}
```

---

### 使用方式

1. 启动单个服务：

```bash
go run main.go serviceA --port=8888
go run main.go serviceB --port=9999
```

2. 同时启动两个服务：

```bash
go run main.go all --portA=8888 --portB=9999
```

---

要不要我帮你扩展一下，让 `all` 启动的服务能 **优雅退出**（比如 Ctrl+C 时能调用 `http.Server.Shutdown` 关闭端口）？
