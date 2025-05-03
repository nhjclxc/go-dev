package main

import (
	"go.uber.org/zap"
)


// go get go.uber.org/zap
// （1）使用 SugaredLogger（推荐开发环境）
func main() {
	logger, _ := zap.NewDevelopment() // 适合开发，输出带颜色
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("User login",
		"username", "Tom",
		"age", 30,
	)

	sugar.Infof("Hello %s", "World")
}
