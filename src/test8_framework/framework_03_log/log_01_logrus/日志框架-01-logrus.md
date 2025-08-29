`logrus` 是 Go 语言中另一款非常流行的日志框架，以**结构化日志**和**丰富的格式化支持**而闻名。相比 `zap` 的性能导向，`logrus` 更注重**易用性、灵活性和兼容性**，非常适合中小型项目或快速原型开发。

---

## 🚀 logrus 简介

* GitHub: [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* 作者：`sirupsen`
* 风格：OOP 风格 API + 支持字段（Fields）+ 多输出

---

## ✅ logrus 特点

| 特性       | 描述                                   |
| -------- | ------------------------------------ |
| 类似标准库接口  | `log.Info()`, `log.Warn()` 等方式       |
| 支持结构化日志  | `WithFields(map[string]interface{})` |
| 多格式支持    | 支持 JSON、文本、钩子等                       |
| 支持日志钩子   | 可将日志发送到 syslog、slack、ELK 等           |
| 支持日志级别控制 | Debug、Info、Warn、Error、Fatal、Panic    |
| 可定制输出目标  | 输出到文件、网络或其他接口                        |

---

## 📦 安装 logrus

```bash
go get github.com/sirupsen/logrus
```

---

## ✍️ 基本用法

```go
package main

import (
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetFormatter(&log.TextFormatter{
        FullTimestamp: true,
    })
    log.SetLevel(log.DebugLevel)

    log.Info("This is an info message")
    log.WithFields(log.Fields{
        "username": "tom",
        "age":      30,
    }).Warn("This is a warning with fields")
}
```

---

## 🔧 设置输出格式

### 1. 文本格式（默认）

```go
log.SetFormatter(&log.TextFormatter{
    FullTimestamp: true,
})
```

### 2. JSON 格式

```go
log.SetFormatter(&log.JSONFormatter{})
```

---

## 🪵 设置日志级别

```go
log.SetLevel(log.InfoLevel) // 只输出 info 及以上级别日志
```

支持的级别：

```go
log.Trace() // 最详细
log.Debug()
log.Info()
log.Warn()
log.Error()
log.Fatal() // 输出后调用 os.Exit(1)
log.Panic() // 输出后 panic
```

---

## 📂 输出到文件

```go
import (
    "os"
)

func init() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        log.SetOutput(file)
    } else {
        log.Info("Failed to log to file, using default stderr")
    }
}
```

---

## 📌 带字段日志（结构化日志）

```go
log.WithFields(log.Fields{
    "order_id": 12345,
    "status":   "shipped",
}).Info("Order processed")
```

---

## 🔗 Hook：将日志转发至第三方

logrus 支持 hook 插件机制：

```go
type MyHook struct{}

func (h *MyHook) Levels() []log.Level {
    return log.AllLevels
}

func (h *MyHook) Fire(entry *log.Entry) error {
    fmt.Println("hook triggered:", entry.Message)
    return nil
}

func main() {
    log.AddHook(&MyHook{})
    log.Info("Test hook")
}
```

还可以使用已有的插件，比如：

* `logrus_syslog`
* `logrus_elasticsearch`
* `logrus_sentry`

---

## 📦 封装成工具类（建议）

你可以将 `logrus` 封装为项目通用日志工具：

```go
// logger/logger.go
package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Log *logrus.Logger

func Init() {
    Log = logrus.New()
    Log.SetFormatter(&logrus.JSONFormatter{})
    Log.SetOutput(os.Stdout)
    Log.SetLevel(logrus.InfoLevel)
}
```

使用：

```go
logger.Log.WithFields(logrus.Fields{
    "anonymous_user": "jack",
    "op":   "login",
}).Info("User operation")
```

---

## ✅ 总结对比（logrus vs zap）

| 特性      | `logrus`     | `zap`             |
| ------- | ------------ | ----------------- |
| 性能      | 中（反射型）       | 高（零内存分配）          |
| 结构化日志支持 | ✅ WithFields | ✅ zap.Field       |
| 日志切割支持  | ❌（需外部实现）     | ❌（可搭配 lumberjack） |
| 输出格式支持  | ✅ JSON/Text  | ✅ JSON/自定义        |
| 开发体验    | 简洁，入门快       | 更复杂，性能优先          |
| 推荐用途    | 中小型项目/快速开发   | 大型项目/高性能服务        |

---

需要我帮你搭建一个完整的 logrus + 文件输出 + hook 的生产级日志模块吗？
