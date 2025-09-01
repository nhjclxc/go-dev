package main

import "otfen_pkg_09_spf13_cobra/cmd"

func main() {
	cmd.Execute()
}

//
//func runService(name string, port int) {
//	fmt.Printf("启动 %s 服务，监听端口: %d\n", name, port)
//	// 这里写具体服务逻辑，比如 http.ListenAndServe(...)
//	select {}
//}
//
//func main() {
//	var rootCmd = &cobra.Command{
//		Use:   "app",
//		Short: "多服务启动器",
//	}
//
//	// serviceA
//	var portA int
//	var serviceACmd = &cobra.Command{
//		Use:   "serviceA",
//		Short: "启动服务 A",
//		Run: func(cmd *cobra.Command, args []string) {
//			runService("serviceA", portA)
//		},
//	}
//	serviceACmd.Flags().IntVar(&portA, "port", 8080, "服务 A 端口")
//
//	// serviceB
//	var portB int
//	var serviceBCmd = &cobra.Command{
//		Use:   "serviceB",
//		Short: "启动服务 B",
//		Run: func(cmd *cobra.Command, args []string) {
//			runService("serviceB", portB)
//		},
//	}
//	serviceBCmd.Flags().IntVar(&portB, "port", 9090, "服务 B 端口")
//
//	// all: 同时启动 A 和 B
//	var allCmd = &cobra.Command{
//		Use:   "all",
//		Short: "同时启动 serviceA 和 serviceB",
//		Run: func(cmd *cobra.Command, args []string) {
//			go runService("serviceA", portA)
//			go runService("serviceB", portB)
//
//			// 捕获退出信号，保持阻塞
//			ch := make(chan os.Signal, 1)
//			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
//			<-ch
//		},
//	}
//	allCmd.Flags().IntVar(&portA, "portA", 8080, "服务 A 端口")
//	allCmd.Flags().IntVar(&portB, "portB", 9090, "服务 B 端口")
//
//	// 注册子命令
//	rootCmd.AddCommand(serviceACmd, serviceBCmd, allCmd)
//	_ = rootCmd.Execute()
//}
