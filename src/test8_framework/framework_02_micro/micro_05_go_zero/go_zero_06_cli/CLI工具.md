

goctl 是 go-zero 配套的代码生成工具脚手架，其集成了 Go HTTP 服务，Go gRPC 服务，数据库 model，k8s，Dockerfile 等生成功能。

goctl 读音为 go control [ɡō kənˈtrōl]，它的主要功能是帮助开发者快速生成代码，提高开发效率。




生成 camel case 文件和目录示例 
$ goctl api new demo --style goZero


# swagger 生成

根据 api 文件生成 swagger 文档，支持生成 json 和 yaml 格式的文档。


```shell
go_zero_06_cli\cli_01_swagger>goctl api swagger -h
Generate swagger file from api

Usage:
  goctl api swagger [flags]

Flags:
      --api string        The api file                                           # *.api源文件
      --dir string        The target dir                                         # 输出目录
      --filename string   The generated swagger file name without the extension  # 生成的 swagger 文件名
  -h, --help              help for swagger                                       # 
      --yaml              Generate swagger yaml file, default to json            # 是否生成 swagger 的 yaml 
```

如：`goctl api swagger --api .\swaggerDoc.api --dir ./swaggerGen --filename swaggerDocFile --yaml`




# goctl api

goctl api 是 goctl 中的核心模块之一，其可以通过 .api 文件一键快速生成一个 api 服务，如果仅仅是启动一个 go-zero 的 api 演示项目， 你甚至都不用编码，就可以完成一个 api 服务开发及正常运行。在传统的 api 项目中，我们要创建各级目录，编写结构体， 定义路由，添加 logic 文件，这一系列操作，如果按照一条协议的业务需求计算，整个编码下来大概需要 5 ～ 6 分钟才能真正进入业务逻辑的编写， 这还不考虑编写过程中可能产生的各种错误，而随着服务的增多，随着协议的增多，这部分准备工作的时间将成正比上升， 而 goctl api 则可以完全替代你去做这一部分工作，不管你的协议要定多少个，最终来说，只需要花费 10 秒不到即可完成。

```shell
go_zero_06_cli\cli_02_api>goctl api -h
Generate api related files

Usage:
  goctl api [flags]
  goctl api [command]

Available Commands:
  dart        Generate dart files for provided api in api file
  doc         Generate doc files
  format      Format api files
  go          Generate go files for provided api in api file
  kt          Generate kotlin code for provided api file
  new         Fast create api service
  plugin      Custom file generator
  swagger     Generate swagger file from api
  ts          Generate ts files for provided api in api file
  validate    Validate api file

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for api
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --o string        Output a sample api file
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure


Use "goctl api [command] --help" for more information about a command.

```

```shell
go_zero_06_cli\cli_02_api>goctl api go -h
Generate go files for provided api in api file

Usage:
  goctl api go [flags]

Flags:
      --api string      The api file
      --branch string   The branch of the remote repo, it does work with --remote
      --dir string      The target dir
  -h, --help            help for go
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md] (default "gozero")
      --test            Generate test files

```

`goctl api go -api order.api --dir ../ --style goZero`


```shell
go_zero_06_cli\cli_02_api>goctl api new -h
Fast create api service

Usage:
  goctl api new [flags]

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for new
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md] (default "gozero")
```


# goctl config 指令

```shell
go_zero_06_cli\cli_02_api>goctl config -h
Usage:
  goctl config [command]

Available Commands:
  clean       Clean goctl config file
  init        Initialize goctl config file

Flags:
  -h, --help   help for config


Use "goctl config [command] --help" for more information about a command.
```

# goctl docker

```shell
go_zero_06_cli\cli_02_api>goctl docker -h
Generate Dockerfile

Usage:
  goctl docker [flags]

Flags:
      --base string      The base image to build the docker image, default scratch (default "scratch")
      --branch string    The branch of the remote repo, it does work with --remote
      --exe string       The executable name in the built image
      --go string        The file that contains main function
  -h, --help             help for docker
      --home string      The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --port int         The port to expose, default none
      --remote string    The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                         The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --tz string        The timezone of the container (default "Asia/Shanghai")
      --version string   The goctl builder golang image version
```

# goctl model


## goctl model

```shell
go_zero_06_cli\cli_02_api>goctl model -h
Generate model code

Usage:
  goctl model [command]

Available Commands:
  mongo       Generate mongo model
  mysql       Generate mysql model
  pg          Generate postgresql model

Flags:
  -h, --help   help for model


Use "goctl model [command] --help" for more information about a command.
```


## goctl model mysql -h

```shell
go_zero_06_cli\cli_02_api>goctl model mysql -h
Generate mysql model

Usage:
  goctl model mysql [command]

Available Commands:
  datasource  Generate model from datasource
  ddl         Generate mysql model from ddl

Flags:
  -h, --help                     help for mysql
  -i, --ignore-columns strings   Ignore columns while creating or updating rows (default [create_at,created_at,create_time,update_at,updated_at,update_time])
  -p, --prefix string            The cache prefix, effective when --cache is true (default "cache")
      --strict                   Generate model in strict mode


Use "goctl model mysql [command] --help" for more information about a command.
```


### goctl model mysql ddl -h

```shell
go_zero_06_cli\cli_02_api>goctl model mysql ddl -h
Generate mysql model from ddl

Usage:
  goctl model mysql ddl [flags]

Flags:
      --branch string     The branch of the remote repo, it does work with --remote
  -c, --cache             Generate code with cache [optional]
      --database string
  -d, --dir string        The target dir
  -h, --help              help for ddl
      --home string       The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --idea              For idea plugin [optional]
      --remote string     The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                          The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
  -s, --src string        The path or path globbing patterns of the ddl
      --style string      The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]


Global Flags:
  -i, --ignore-columns strings   Ignore columns while creating or updating rows (default [create_at,created_at,create_time,update_at,updated_at,update_time])
  -p, --prefix string            The cache prefix, effective when --cache is true (default "cache")
      --strict                   Generate model in strict mode
```

```goctl model mysql ddl -src "path/to/your.sql" -dir "path/to/output" -caching```

```goctl model mysql ddl -src ./user.sql -dir ./model -caching```

| 参数         | 说明                                     |
| ---------- | -------------------------------------- |
| `-src`     | 指定包含建表语句（DDL）的 `.sql` 文件路径             |
| `-dir`     | 指定生成的 model 文件输出目录                     |
| `-caching` | 是否启用缓存逻辑（可选，启用后会生成 Redis 缓存相关代码）       |
| `-style`   | 命名风格，如 `go_zero`（默认）、`camel`、`snake` 等 |
| `-home`    | 指定 goctl 的模板目录（可选）                     |





### goctl model mysql datasource -h 

```shell
go_zero_06_cli\cli_02_api>goctl model mysql datasource -h
Generate model from datasource

Usage:
  goctl model mysql datasource [flags]

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -c, --cache           Generate code with cache [optional]
  -d, --dir string      The target dir
  -h, --help            help for datasource
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --idea            For idea plugin [optional]
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format, see [https://github.com/zeromicro/go-zero/blob/master/tools/goctl/config/readme.md]
  -t, --table strings   The table or table globbing patterns in the database
      --url string      The data source of database,like "root:password@tcp(127.0.0.1:3306)/database


Global Flags:
  -i, --ignore-columns strings   Ignore columns while creating or updating rows (default [create_at,created_at,create_time,update_at,updated_at,update_time])
  -p, --prefix string            The cache prefix, effective when --cache is true (default "cache")
      --strict                   Generate model in strict mode
```


```
goctl model mysql datasource -url="user:password@tcp(host:port)/dbname" -table="tableName1,tableName2,table*" -dir="./model"
```

```
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/testdb" -table="user" -dir="./model"
```

| 参数名        | 说明                                               |
| ---------- | ------------------------------------------------ |
| `-url`     | 数据源连接串，格式为 `user:password@tcp(host:port)/dbname` |
| `-table`   | 要生成模型的表名，可以是多个表名用英文逗号隔开，也可以使用通配符 `*`             |
| `-dir`     | 生成的代码放置的目录                                       |
| `-caching` | 是否开启缓存功能（需要配置 Redis）                             |
| `-style`   | 命名风格，如 `go_zero`、`camel` 等（默认是 `go_zero` 风格）     |




# goctl rpc

```shell
go_zero_06_cli\cli_02_api>goctl rpc -h
Generate rpc code

Usage:
  goctl rpc [flags]
  goctl rpc [command]

Available Commands:
  new         Generate rpc demo service
  protoc      Generate grpc code
  template    Generate proto template

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -h, --help            help for rpc
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --o string        Output a sample proto file
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure


Use "goctl rpc [command] --help" for more information about a command.
```



