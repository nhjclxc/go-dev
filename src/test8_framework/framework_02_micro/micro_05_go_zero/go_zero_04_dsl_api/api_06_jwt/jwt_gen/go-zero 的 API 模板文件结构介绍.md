
# go-zero 的 API 模板文件结构介绍

https://go-zero.dev/docs/concepts/layout



etc[etc](etc)：存放配置文件

[internal](internal)：这个服务内部的所有代码，供内部使用，防止被其他模块误引用

[config/config.go](internal%2Fconfig%2Fconfig.go)：定义并加载配置结构体，结合 go-zero 的 `conf.MustLoad()` 函数读取如 [user-api.yaml](etc%2Fuser-api.yaml) 中的配置项，赋值到结构体中，供全局使用。

[handler/routes.go](internal%2Fhandler%2Froutes.go)：定义路由与使用中间件。即绑定前端的请求 api 路径，加入一些中间件的使用，如日志中间件、鉴权中间件等

[handler](internal%2Fhandler) 里面的 `*.Handler.go`：接收请求，定义每个 API 接口的 HTTP 路由处理函数，接收并解析请求参数，调用对应的 logic 层方法处理业务逻辑，类似于 SpringBoot 中的 Controller 层。

[logic](internal%2Flogic) 里面的 `*Logic.go`：里面就是实际的接口业务逻辑处理，类似于SpringBoot里面的Service层一样

[serviceContext.go](internal%2Fsvc%2FserviceContext.go)：定义并初始化所有依赖资源，如配置、数据库连接、RPC 客户端等，并在 logic 层通过注入方式使用，类似于一个依赖管理容器（SpringBoot的IOC）。

[types.go](internal%2Ftypes%2Ftypes.go)：定义由 .api 文件自动生成的请求与响应结构体，以及部分错误定义（如 ErrNotFound），供 handler 层使用。




以下是 go-zero API 项目的标准目录结构示意图，通常由 `goctl api new` 命令生成：([go-zero.dev](https://go-zero.dev/docs/tutorials/cli/api))

```
.
├── etc
│   └── user-api.yaml          # 配置文件
├── user.api                   # API 定义文件
├── go.mod
├── user.go                    # 程序入口
└── internal
    ├── config
    │   └── config.go          # 配置结构体定义与加载
    ├── handler
    │   ├── routes.go          # 路由与中间件定义
    │   └── *.handler.go       # 各接口的处理函数
    ├── logic
    │   └── *.logic.go         # 业务逻辑处理
    ├── svc
    │   └── serviceContext.go  # 服务上下文与依赖注入
    └── types
        └── types.go           # 请求与响应结构体定义
```

该结构有助于实现清晰的职责分离，提升代码的可维护性和可扩展性。


