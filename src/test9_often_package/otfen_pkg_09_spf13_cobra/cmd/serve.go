package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"otfen_pkg_09_spf13_cobra/app"
	"otfen_pkg_09_spf13_cobra/config"
	"syscall"
)

// serveCmd 用于启动 所有 服务
// go run main.go serve-cmd -c ./config/config.yaml
// go run main.go serve-cmd --config ./config/config.yaml

var serveCmd = &cobra.Command{
	Use: "serve-cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())

		// 加载配置
		err := config.InitConfig(cfg, cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		// 启动应用程序
		a, err := app.NewApplication(ctx, cancel, cfg)

		// 启动服务
		a.StartApp()

		// 等待中断信号
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		// 优雅关闭
		a.StopApp()

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config/config.yaml", "配置文件路径")

	rootCmd.AddCommand(serveCmd)
}
