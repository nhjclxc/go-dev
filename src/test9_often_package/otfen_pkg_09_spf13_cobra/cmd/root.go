package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"otfen_pkg_09_spf13_cobra/config"
)

var (
	cfg     = &config.Config{}
	cfgFile = ""
	//cfgFile = "./config/config.yaml"
)
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "应用根命令",
	// 如果用户没有输入任何子命令，就报错
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("必须指定子命令，比如: app http-cmd -c ./config/config.yaml")
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {

	// cobra.OnInitialize 是一个 全局钩子函数注册器，用于在 Execute() 运行之前执行一些初始化逻辑。
	cobra.OnInitialize(func() {
		fmt.Printf("执行OnInitialize \n")
	})

	// 在每一个文件的init函数里面进行子命令注册
	// 如：rootCmd.AddCommand(httpCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
