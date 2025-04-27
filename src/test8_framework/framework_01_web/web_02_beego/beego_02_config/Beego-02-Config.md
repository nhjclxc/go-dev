
### [Beego-02、配置模块](https://beegodoc.com/zh/developing/config/#%E5%88%9D%E5%A7%8B%E5%8C%96%E6%96%B9%E6%B3%95)


配置模块是基础模块之一，对不同类型的配置文件提供了一种抽象。该章节内容都可以在[配置模块例子](https://github.com/beego/beego-example/tree/master/config)

Beego 目前支持 INI、XML、JSON、YAML 格式的配置文件解析，也支持以 etcd 作为远程配置中心。默认采用了 INI 格式解析，用户可以通过简单的配置就可以获得很大的灵活性。

它们拥有的方法都是一样的，具体可以参考[Config API](https://github.com/beego/beego/blob/develop/core/config/config.go)。主要方法有：

```go
// Configer defines how to get and set value from configuration raw data.
type Configer interface {
	// support section::key type in given key when using ini type.
	Set(key, val string) error

	// support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
	String(key string) (string, error)
	// get string slice
	Strings(key string) ([]string, error)
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	// support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
	DefaultString(key string, defaultVal string) string
	// get string slice
	DefaultStrings(key string, defaultVal []string) []string
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64

	// DIY return the original value
	DIY(key string) (interface{}, error)

	GetSection(section string) (map[string]string, error)

	Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error
	Sub(key string) (Configer, error)
	OnChange(key string, fn func(value string))
	SaveConfigFile(filename string) error
}
```

这里有一些使用的注意事项：
1. 所有的Default*方法，在key不存在，或者查找的过程中，出现error，都会返回默认值；
2. DIY直接返回对应的值，而没有做任何类型的转换。当你使用这个方法的时候，你应该自己确认值的类型。只有在极少数的情况下你才应该考虑使用这个方法；
3. GetSection会返回section所对应的部分配置。section如何被解释，取决于具体的实现；
4. Unmarshaler会尝试用当且配置的值来初始化obj。需要注意的是，prefix的概念类似于section；
5. Sub类似与GetSection，都是尝试返回配置的一部分。所不同的是，GetSection将结果组织成map，而Sub将结果组织成Config实例；
6. OnChange主要用于监听配置的变化。对于大部分依赖于文件系统的实现来说，都不支持。具体而言，我们设计这个主要是为了考虑支持远程配置；
7. SaveConfigFile尝试将配置导出成为一个文件；
8. 某些实现支持分段式的key。比如说a.b.c这种，但是，并不是所有的实现都支持，也不是所有的实现都采用.作为分隔符。这是一个历史遗留问题，为了保留兼容性，我们无法在这方面保持一致。

Web 模块封装了配置模块，可以参考[Web 配置](https://beegodoc.com/zh/developing/web/config.html)

