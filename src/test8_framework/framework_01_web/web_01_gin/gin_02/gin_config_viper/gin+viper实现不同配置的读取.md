在 Gin 项目中使用 Viper 实现「不同配置的读取」非常常见，适合实现多环境配置（如 dev/test/prod）、模块配置（如数据库、Redis、日志）等。

---

## ✅ 目标示例

我们希望支持以下功能：

* 支持读取多模块配置（如 `server`、`mysql`、`redis` 等）。
* 支持读取多个环境的配置（如 `config-dev.yaml` / `config-prod.yaml`）。
* 在代码中优雅地访问配置项。

---

## ✅ 一、项目结构示例

```
.
├── config
│   ├── config.yaml         # 默认配置
│   ├── config-dev.yaml     # 开发环境
│   ├── config-prod.yaml    # 生产环境
├── conf
│   └── config.go           # 读取配置的初始化逻辑
├── main.go
```

---

## ✅ 二、配置文件样例（YAML）

`config/config.yaml`（默认配置）：

```yaml
server:
  port: 8080
  mode: debug

mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: root
  database: test

redis:
  addr: 127.0.0.1:6379
  password: ""
  db: 0
```

`config/config-dev.yaml`：

```yaml
server:
  port: 8081
  mode: debug

mysql:
  database: dev_db
```

`config/config-prod.yaml`：

```yaml
server:
  port: 80
  mode: release

mysql:
  host: prod.db.host
  password: prod_pass
  database: prod_db
```

---

## ✅ 三、读取配置逻辑（conf/config.go）

```go
package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Conf *AppConfig

// 统一配置结构体
type AppConfig struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func InitConfig(env string) {
	v := viper.New()
	v.SetConfigName("config")        // 默认 config.yaml
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")      // 配置文件路径

	// 读取默认配置
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取默认配置失败: %v", err)
	}

	// 根据环境读取额外配置覆盖
	if env != "" {
		v.SetConfigName(fmt.Sprintf("config-%s", env))
		if err := v.MergeInConfig(); err != nil {
			log.Printf("读取 config-%s.yaml 失败：%v", env, err)
		}
	}

	var c AppConfig
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("配置解析失败: %v", err)
	}

	Conf = &c
	log.Printf("配置加载成功：%+v\n", Conf)
}
```

---

## ✅ 四、main.go 使用示例

```go
package main

import (
	"fmt"
	"your_project/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	// 可以通过命令行/env/env变量决定环境
	conf.InitConfig("dev") // 传 "prod" 或 "dev"

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mysql": conf.Conf.MySQL,
			"redis": conf.Conf.Redis,
			"msg":   "pong",
		})
	})

	port := conf.Conf.Server.Port
	r.Run(fmt.Sprintf(":%d", port))
}
```

---

## ✅ 五、运行示例

```bash
go run main.go
```

输出：

```
2025/07/22 09:00:00 配置加载成功：&{Server:{Port:8081 Mode:debug} MySQL:{Host:127.0.0.1 Port:3306 Username:root Password:root Database:dev_db} Redis:{Addr:127.0.0.1:6379 Password: DB:0}}
```

---

## ✅ 六、支持从命令行或环境变量设置环境名

使用 `flag` 或 `os.Getenv` 更灵活：

```go
// main.go
env := os.Getenv("GIN_ENV") // 或 flag.Parse()
conf.InitConfig(env)
```

运行：

```bash
GIN_ENV=prod go run main.go
```

---

## ✅ 总结

| 功能     | 实现方式                             |
| ------ | -------------------------------- |
| 多模块配置  | 使用结构体嵌套                          |
| 多环境切换  | 用 `viper.MergeInConfig()` 读取环境覆盖 |
| 动态环境切换 | 支持传入 env 参数，或通过环境变量配置            |
| 热加载配置  | viper 支持 `WatchConfig()`（可选实现）   |

---

如你需要支持 **配置热更新** 或 **多模块拆分成多个文件合并读取**，也可以继续告诉我，我可以补充完整。
