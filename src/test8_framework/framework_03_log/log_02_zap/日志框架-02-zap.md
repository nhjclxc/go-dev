`zap` 是 Go 语言中非常流行的一款高性能结构化日志库，由 Uber 开发。它专注于 **高性能** 和 **结构化日志记录**，支持 JSON 格式输出、不同日志级别、日志分组、日志切割（需配合 lumberjack）等特性。

---

## ✅ zap 框架特点

| 特性      | 说明                                 |
| ------- | ---------------------------------- |
| 高性能     | 比 logrus 快 4\~10 倍，尤其在高并发场景        |
| 支持结构化日志 | 以 key-value 形式记录日志，利于机器读取与分析       |
| 两种风格    | `SugaredLogger`（易用型），`Logger`（性能型） |
| 灵活配置    | 支持生产/开发环境预设配置，支持文件输出               |
| 与其他库兼容  | 可以与 `lumberjack`、`gorm` 等结合        |

---

## 🛠️ zap 使用方法详解

### 1. 安装依赖

```bash
go get go.uber.org/zap
```

---

### 2. 基本用法

#### （1）使用 SugaredLogger（推荐开发环境）

```go
package main

import (
    "go.uber.org/zap"
)

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
```

#### （2）使用性能更高的 Logger（推荐生产环境）

```go
package main

import (
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("User login",
        zap.String("username", "Tom"),
        zap.Int("age", 30),
    )
}
```

---

### 3. 自定义配置

可以通过 zap 的 `zap.Config` 创建更灵活的日志器：

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
    config := zap.Config{
        Encoding:         "json", // "console" or "json"
        Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
        OutputPaths:      []string{"stdout", "./log/app.log"},
        ErrorOutputPaths: []string{"stderr"},
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "time",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            StacktraceKey:  "stacktrace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 彩色等级
            EncodeTime:     zapcore.ISO8601TimeEncoder,       // 时间格式
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
    }

    logger, _ := config.Build()
    return logger
}
```

---

### 4. 配合 lumberjack 实现日志切割

`zap` 本身不支持日志轮转，但可以通过 `lumberjack` 实现：

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

func InitLogger() *zap.Logger {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   "./log/app.log",
        MaxSize:    100, // MB
        MaxBackups: 3,
        MaxAge:     28,   // days
        Compress:   true, // gzip
    }

    writeSyncer := zapcore.AddSync(lumberJackLogger)
    encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

    return zap.New(core, zap.AddCaller())
}
```

---

### 5. 日志级别设置

| 日志级别   | 说明               |
| ------ | ---------------- |
| Debug  | 调试用日志            |
| Info   | 常规日志             |
| Warn   | 警告信息             |
| Error  | 错误信息             |
| DPanic | panic前的开发环境触发    |
| Panic  | 日志后直接 panic      |
| Fatal  | 日志后程序退出（os.Exit） |

---

### 6. 输出示例（JSON）

```json
{
  "level":"info",
  "time":"2025-05-03T15:04:05.000+0800",
  "caller":"main.go:14",
  "msg":"User login",
  "username":"Tom",
  "age":30
}
```

---

## 📦 项目中统一封装 zap 工具类（建议）

你可以封装一个通用日志初始化模块，然后在项目中全局调用：

```go
// logger/logger.go
package logger

var Log *zap.SugaredLogger

func Init() {
    l, _ := zap.NewProduction()
    Log = l.Sugar()
}
```

然后在 main 或其他地方使用：

```go
logger.Log.Infow("Startup",
    "env", "prod",
    "version", "1.0.0",
)
```

---

## ✅ 总结

| 类型     | 方法名                      | 适合环境 |
| ------ | ------------------------ | ---- |
| 开发调试   | `zap.NewDevelopment()`   | 开发   |
| 生产部署   | `zap.NewProduction()`    | 生产   |
| 自定义设置  | `zap.Config` + `Build()` | 高级用户 |
| 支持日志切割 | 搭配 `lumberjack` 使用       | 日志归档 |

---

是否需要我帮你生成一个适合你项目结构的完整 zap 封装模板？
