package main

import (
	_ "web_06_input/routers"

	beego "github.com/beego/beego/v2/server/web"
)

// 学习 Beego 的输入处理
// https://beegodoc.com/zh/developing/web/input/
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//总体来说，处理输入主要依赖于 Controller 提供的方法。而具体输入可以来源于：
	//	路径参数：这一部分主要是指参数路由
	//	查询参数
	//	请求体：要想从请求体里面读取数据，大多数时候将BConfig.CopyRequestBody 设置为true就足够了。而如果你是创建了多个 web.Server，那么必须每一个Server实例里面的配置都将CopyRequestBody设置为true了

	//而获取参数的方法可以分成两大类：
	//第一类是以 Get 为前缀的方法：这一大类的方法，主要获得某个特定参数的值
	// 针对这一类方法，Beego 主要从两个地方读取：查询参数和表单，如果两个地方都有相同名字的参数，那么 Beego 会返回表单里面的数据。
	//		GetString(key string, def ...string) string
	//		GetStrings(key string, def ...[]string) []string
	//		GetInt(key string, def ...int) (int, error)
	//		GetInt8(key string, def ...int8) (int8, error)
	//		GetUint8(key string, def ...uint8) (uint8, error)
	//		GetInt16(key string, def ...int16) (int16, error)
	//		GetUint16(key string, def ...uint16) (uint16, error)
	//		GetInt32(key string, def ...int32) (int32, error)
	//		GetUint32(key string, def ...uint32) (uint32, error)
	//		GetInt64(key string, def ...int64) (int64, error)
	//		GetUint64(key string, def ...uint64) (uint64, error)
	//		GetBool(key string, def ...bool) (bool, error)
	//		GetFloat(key string, def ...float64) (float64, error)

	//第二类是以 Bind 为前缀的方法：这一大类的方法，试图将输入转化为结构体
	//		Bind(obj interface{}) error: 默认是依据输入的 Content-Type字段，来判断该如何反序列化；
	//		BindYAML(obj interface{}) error: 处理YAML输入
	//		BindForm(obj interface{}) error: 处理表单输入
	//		BindJSON(obj interface{}) error: 处理JSON输入
	//		BindProtobuf(obj proto.Message) error: 处理proto输入
	//		BindXML(obj interface{}) error: 处理XML输入

	beego.BConfig.CopyRequestBody = true

	//
	//

	// 文件上传设置
	// https://beegodoc.com/zh/developing/web/file/
	// 设置缓存内存大小大小为 64 MB
	beego.BConfig.MaxMemory = 1 << 22

	// 设置文件上传大小为 64 MB
	// MaxUploadSize来限制最大上传文件大小——如果你一次长传多个文件，那么它限制的就是这些所有文件合并在一起的大小。
	beego.BConfig.MaxUploadSize = 1 << 22

	beego.Run()
}
