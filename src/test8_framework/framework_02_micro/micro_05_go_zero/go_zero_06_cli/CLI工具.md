

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




