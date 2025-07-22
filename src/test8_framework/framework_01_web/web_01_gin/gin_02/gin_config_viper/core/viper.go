package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	localConfig "gin_config_viper/global"
)

// Viper 配置，使用 Viper 读取配置
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()        // 创建viper对象
	v.SetConfigFile(config) // 设置配置文件地址
	v.SetConfigType("yaml")

	// 读取配置
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 🔁 热更新配置（这时一个可选配置）
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置更新后回进入这个回调函数
		fmt.Println("config file changed:", e.Name)
		// 配置更新后，将 Viper 中已加载的配置数据（比如从 config.yaml 中）解析并填充到结构体 Global.GlobalConfig 中。
		if err = v.Unmarshal(&localConfig.GlobalConfig); err != nil {
			fmt.Println(err)
		}
	})
	// 将 Viper 中已加载的配置数据（比如从 config.yaml 中）解析并填充到结构体 global.GVA_CONFIG 中。
	if err = v.Unmarshal(&localConfig.GlobalConfig); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	//root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {
	// `-c` flag parse
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" { // 命令行参数不为空 将值赋值于config
		// go run main.go -c config-debug.yaml
		// go run main.go -c config-release.yaml
		// go run main.go -c config-test.yaml
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}

	switch gin.Mode() { // 根据 gin 模式文件名
	case gin.DebugMode:
		config = localConfig.ConfigDevFile
	case gin.ReleaseMode:
		config = localConfig.ConfigReleaseFile
	case gin.TestMode:
		config = localConfig.ConfigTestFile
	}
	fmt.Printf("您正在使用 gin 的 %s 模式运行, config 的路径为 %s\n", gin.Mode(), config)

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = localConfig.ConfigDevFile
		fmt.Printf("配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
