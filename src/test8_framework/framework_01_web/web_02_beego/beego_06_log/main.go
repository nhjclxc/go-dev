package main

import (

	// beego 日志包
	"github.com/beego/beego/v2/core/logs"
)

func init() {
	// 主要的参数如下说明：
	//
	//filename 保存的文件名
	//maxlines 每个文件保存的最大行数，默认值 1000000
	//maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, 256 MB
	//daily 是否按照每天 logrotate，默认是 true
	//maxdays 文件最多保存多少天，默认保存 7 天
	//rotate 是否开启 logrotate，默认是 true
	//level 日志保存的时候的级别，默认是 Trace 级别
	//perm 日志文件权限

	var config string = `
		"x": 1,

	`

	logs.SetLogger(logs.AdapterConsole, config)

	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	logs.EnableFuncCallDepth(true)

	// 我们知道，计算机的执行速度是非常快的，但是如果请求里面需要不断输出日志的话，会导致请求变慢，
	// 特别是线上的时候，在不影响服务的情况下，又不想影响日志输出的话，就可以使用异步输出日志
	//为了提升性能, 可以设置异步输出:
	logs.Async()
	//异步输出允许设置缓冲 chan 的大小
	logs.Async(1e3)


	// 自定义日志格式
	// https://beegodoc.com/zh/developing/logs/#%E8%87%AA%E5%AE%9A%E4%B9%89%E6%97%A5%E5%BF%97%E6%A0%BC%E5%BC%8F





}

func main() {
	//an official log.Logger
	l := logs.GetLogger()
	l.Println("this is a message of http")
	// 2025/05/02 15:10:07.452 [M]  this is a message of http

	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")
	//2025/05/02 15:10:07.461 [M]  [ORM] this is a message of orm

	// 以下的日志级别从低到高
	//
	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	//2025/05/02 15:10:07.461 [D]  my book is bought in the year of  2016
	//2025/05/02 15:10:07.461 [I]  this yellow cat is 3 years old
	//2025/05/02 15:10:07.461 [W]  json is a type of kv like map[key:2016]
	//2025/05/02 15:10:07.461 [E]  1024 is a very good game
	//2025/05/02 15:10:07.461 [C]  oh,crash
}
