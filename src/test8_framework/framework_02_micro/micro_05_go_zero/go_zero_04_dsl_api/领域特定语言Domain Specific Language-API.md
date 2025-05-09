
api 是 go-zero 自研的领域特性语言

goctl api go -api *.api --dir api_gen --style goZero



在 API 描述语言中，类型声明需要满足如下规则:
- 类型声明必须以 type 开头
- 不需要声明 struct关键字



# goctl 命令生成 *.api 文件模板

```
goctl api new apidemo
```

# API 规范

1. @doc 语句

@doc 语句是对单个路由的 meta 信息描述，一般为 key-value 值，可以传递给 goctl 及其插件来进行扩展生成

2. @handler 语句

@handler 语句是对单个路由的 handler 信息控制，主要用于生成 golang http.HandleFunc 的实现转换方法

2. server 语句

service 语句是对 HTTP 服务的直观描述，包含请求 handler，请求方法，请求路由，请求体，响应体，jwt 开关，中间件声明等定义。

2. @server 语句

@server 语句是对一个服务语句的 meta 信息描述，其对应特性包含但不限于：jwt 开关、中间件、路由分组、路由前缀

```api
// 空内容
@server()

// 有内容
@server (
    // jwt 声明
    // 如果 key 固定为 “jwt:”，则代表开启 jwt 鉴权声明
    // value 则为配置文件的结构体名称
    jwt: Auth

    // 路由前缀
    // 如果 key 固定为 “prefix:”
    // 则代表路由前缀声明，value 则为具体的路由前缀值，字符串中没让必须以 / 开头
    prefix: /v1

    // 路由分组
    // 如果 key 固定为 “group:”，则代表路由分组声明
    // value 则为具体分组名称，在 goctl生成代码后会根据此值进行文件夹分组
    group: Foo

    // 中间件
    // 如果 key 固定为 middleware:”，则代表中间件声明
    // value 则为具体中间件函数名称，在 goctl生成代码后会根据此值进生成对应的中间件函数
    middleware: AuthInterceptor

    // 超时控制
    // 如果 key 固定为  timeout:”，则代表超时配置
    // value 则为具体中duration，在 goctl生成代码后会根据此值进生成对应的超时配置
    timeout: 3s

    // 其他 key-value，除上述几个内置 key 外，其他 key-value
    // 也可以在作为 annotation 信息传递给 goctl 及其插件，但就
    // 目前来看，goctl 并未使用。
    foo: bar
)
```


3. HTTP 服务的请求/响应体语句

路由语句是对单此 HTTP 请求的具体描述，包括请求方法，请求路径，请求体，响应体信息

```api
// 没有请求体和响应体的写法
get /ping

// 只有请求体的写法
get /foo (foo)

// 只有响应体的写法
post /foo returns (foo)

// 有请求体和响应体的写法
post /foo (foo) returns (bar)
```

4. 注释语句
5. import 语句

import 语句是在 api 中引入其他 api 文件的语法块，其支持相对/绝对路径，

```api
// 单行 import
import "foo"
import "/path/to/file"

// import 组
import ()
import (
    "bar"
    "relative/to/file"
)
```

6. info 语句

info 语句是 api 语言的 meta 信息，其仅用于对当前 api 文件进行描述，暂不参与代码生成，其和注释还是有一些区别，注释一般是依附某个 syntax 语句存在，而 info 语句是用于描述整个 api 信息的，当然，不排除在将来会参与到代码生成里面来，

```api
// 不包含 key-value 的 info 块
info ()

// 包含 key-value 的 info 块
info (
    foo: "bar"
    bar:
)
```

7. HTTP 路由语句
8. HTTP 服务声明语句
9. syntax 语句
   
syntax 语句用于标记 api 语言的版本，不同的版本可能语法结构有所不同，随着版本的提升会做不断的优化，当前版本为 v2。

   `syntax = "v1"`

10. 结构体语句

```api
// 空结构体
type Foo {}

// 单个结构体
type Bar {
    Foo int               `json:"foo"`
    Bar bool              `json:"bar"`
    Baz []string          `json:"baz"`
    Qux map[string]string `json:"qux"`
}

type Baz {
    Bar    `json:"baz"`
    Array [3]int `json:"array"`
    // 结构体内嵌 goctl 1.6.8 版本支持
    Qux {
        Foo string `json:"foo"`
        Bar bool   `json:"bar"`
    } `json:"baz"`
}

// 空结构体组
type ()

// 结构体组
type (
    Int int
    Integer = int
    Bar {
        Foo int               `json:"foo"`
        Bar bool              `json:"bar"`
        Baz []string          `json:"baz"`
        Qux map[string]string `json:"qux"`
    }
)

```

# 完整的 API 定义示例

[API 定义完整示例](https://go-zero.dev/docs/reference)

```api
syntax = "v1"

info (
    title:   "api 文件完整示例写法"
    desc:    "演示如何编写 api 文件"
    author:  "keson.an"
    date:    "2022 年 12 月 26 日"
    version: "v1"
)

type UpdateReq {
    Arg1 string `json:"arg1"`
}

type ListItem {
    Value1 string `json:"value1"`
}

type LoginReq {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResp {
    Name string `json:"name"`
}

type FormExampleReq {
    Name string `form:"name"`
}

type PathExampleReq {
    // path 标签修饰的 id 必须与请求路由中的片段对应，如
    // id 在 service 语法块的请求路径上一定会有 :id 对应，见下文。
    ID string `path:"id"`
}

type PathExampleResp {
    Name string `json:"name"`
}

@server (
    jwt:        Auth // 对当前 Foo 语法块下的所有路由，开启 jwt 认证，不需要则请删除此行
    prefix:     /v1 // 对当前 Foo 语法块下的所有路由，新增 /v1 路由前缀，不需要则请删除此行
    group:      g1 // 对当前 Foo 语法块下的所有路由，路由归并到 g1 目录下，不需要则请删除此行
    timeout:    3s // 对当前 Foo 语法块下的所有路由进行超时配置，不需要则请删除此行
    middleware: AuthInterceptor // 对当前 Foo 语法块下的所有路由添加中间件，不需要则请删除此行
    maxBytes:   1048576 // 对当前 Foo 语法块下的所有路由添加请求体大小控制，单位为 byte,goctl 版本 >= 1.5.0 才支持
)
service Foo {
    // 定义没有请求体和响应体的接口，如 ping
    @handler ping
    get /ping

    // 定义只有请求体的接口，如更新信息
    @handler update
    post /update (UpdateReq)

    // 定义只有响应体的结构，如获取全部信息列表
    @handler list
    get /list returns ([]ListItem)

    // 定义有结构体和响应体的接口，如登录
    @handler login
    post /login (LoginReq) returns (LoginResp)

    // 定义表单请求
    @handler formExample
    post /form/example (FormExampleReq)

    // 定义 path 参数
    @handler pathExample
    get /path/example/:id (PathExampleReq) returns (PathExampleResp)
}

```


