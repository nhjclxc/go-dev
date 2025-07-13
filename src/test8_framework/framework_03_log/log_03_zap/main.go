package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"strconv"

	"log_03_zap/core"
	"log_03_zap/global"
)

func main() {

	global.GlobalViper = core.Viper() // 初始化Viper

	// 初始化zap日志库
	global.GlobalZapLog = core.Zap()
	zap.ReplaceGlobals(global.GlobalZapLog) // 并将其设为全局默认的 zap 日志实例

	// 创建一个默认的路由引擎
	router := gin.Default()

	// 路由绑定
	router.GET("/logs", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")

		global.GlobalZapLog.Debug("Debug 我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v \n", zap.String("name", global.GlobalConfig.Name), zap.String("id", id), zap.String("name", name))
		global.GlobalZapLog.Info("Info 我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v \n", zap.String("name", global.GlobalConfig.Name), zap.String("id", id), zap.String("name", name))
		global.GlobalZapLog.Warn("Warn 我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v \n", zap.String("name", global.GlobalConfig.Name), zap.String("id", id), zap.String("name", name))
		global.GlobalZapLog.Error("Error 我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v \n", zap.String("name", global.GlobalConfig.Name), zap.String("id", id), zap.String("name", name))

		num1Str := context.Query("num1")
		num2Str := context.Query("num2")

		defer func() {
			if r := recover(); r != nil {
				// 如何把panic堆栈打印出来？？？
				//global.GlobalZapLog.Error("除数不能为0", )

				global.GlobalZapLog.Error("程序发生 panic",
					zap.Any("type", "除数不能为0"),              // 打印 panic 的值（可能是字符串或 error）
					zap.Any("error", r),                    // 打印 panic 的值（可能是字符串或 error）
					zap.ByteString("stack", debug.Stack()), // 打印堆栈信息
				)

				context.String(http.StatusOK, "除数不能为0")

			}
		}()
		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)

		res := div(num1, num2)

		global.GlobalZapLog.Debug("This is a debug message = " + strconv.Itoa(res)) // 通常不会在生产环境中显示
		global.GlobalZapLog.Info("This is an info message = " + strconv.Itoa(res))  // 日常操作记录
		global.GlobalZapLog.Warn("This is a warning = " + strconv.Itoa(res))        // 不影响运行，但需注意
		global.GlobalZapLog.Error("This is an error = " + strconv.Itoa(res))        // 运行出错，但系统仍可继续运行
		//global.GlobalZapLog.DPanic("This is a DPanic = " + strconv.Itoa(res))         // 开发中 panic，生产中 error
		//global.GlobalZapLog.Panic("This is a panic")           // 打印后会触发 panic
		// 不得在生产中使用以下命令，该命令会导致程序关闭
		//global.GlobalZapLog.Fatal("This is fatal")             // 打印后调用 os.Exit(1)

		context.String(http.StatusOK, "我是 news 页面，当前环境是：%s，现在请求的文章id = %v, name = %v, div(num1, num2) = %d \n", global.GlobalConfig.Name, id, name, res)
	})

	// go run main.go
	// go run main.go -c config/config-dev.yaml
	// go run main.go -c config/config-test.yaml
	// go run main.go -c config/config-prod.yaml

	// http://localhost:8080/logs?id=666&name=HelloGolang

	router.Run(":" + strconv.Itoa(global.GlobalConfig.Port))

}

func div(num1, num2 int) int {
	return num1 / num2
}
