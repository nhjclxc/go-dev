


本教程使用 Protocol Buffers 语言的 proto3 版本，为 Go 程序员提供了使用 Protocol Buffers 的基本介绍。通过创建一个简单的示例应用程序，它向您展示了如何

- 在 .proto 文件中定义消息格式。
- 使用协议缓冲区编译器。
- 使用 Go protocol buffer API 来写入和读取消息。



# 服务分组概述

go-zero 采用 gRPC 进行服务间的通信，我们通过 proto 文件来定义服务的接口，但是在实际的开发中，我们可能会有多个服务，如果不对服务进行文件分组，那么 goctl 生成的代码将会是一个大的文件夹，这样会导致代码的可维护性变差，因此服务分组可以提高代码的可读性和可维护性。



# 服务分组

在 go-zero 中，我们通过在 proto 文件中以 service 为维度来进行文件分组，我们可以在 proto 文件中定义多个 service，每个 service 都会生成一个独立的文件夹，这样就可以将不同的服务进行分组，从而提高代码的可读性和可维护性。

除了 proto 文件中定义了 service 外，分组与否还需要在 goctl 中控制，生成带分组或者不带分组的代码取决于开发者，我们通过示例来演示一下。




## 不带分组


```protobuf
syntax = "proto3";

package user;

option go_package = "github.com/example/anonymous_user";

message LoginReq{}
message LoginResp{}
message UserInfoReq{}
message UserInfoResp{}
message UserInfoUpdateReq{}
message UserInfoUpdateResp{}
message UserListReq{}
message UserListResp{}

message UserRoleListReq{}
message UserRoleListResp{}
message UserRoleUpdateReq{}
message UserRoleUpdateResp{}
message UserRoleInfoReq{}
message UserRoleInfoResp{}
message UserRoleAddReq{}
message UserRoleAddResp{}
message UserRoleDeleteReq{}
message UserRoleDeleteResp{}


message UserClassListReq{}
message UserClassListResp{}
message UserClassUpdateReq{}
message UserClassUpdateResp{}
message UserClassInfoReq{}
message UserClassInfoResp{}
message UserClassAddReq{}
message UserClassAddResp{}
message UserClassDeleteReq{}
message UserClassDeleteResp{}

service UserService{
  rpc Login (LoginReq) returns (LoginResp);
  rpc UserInfo (UserInfoReq) returns (UserInfoResp);
  rpc UserInfoUpdate (UserInfoUpdateReq) returns (UserInfoUpdateResp);
  rpc UserList (UserListReq) returns (UserListResp);

  rpc UserRoleList (UserRoleListReq) returns (UserRoleListResp);
  rpc UserRoleUpdate (UserRoleUpdateReq) returns (UserRoleUpdateResp);
  rpc UserRoleInfo (UserRoleInfoReq) returns (UserRoleInfoResp);
  rpc UserRoleAdd (UserRoleAddReq) returns (UserRoleAddResp);
  rpc UserRoleDelete (UserRoleDeleteReq) returns (UserRoleDeleteResp);

  rpc UserClassList (UserClassListReq) returns (UserClassListResp);
  rpc UserClassUpdate (UserClassUpdateReq) returns (UserClassUpdateResp);
  rpc UserClassInfo (UserClassInfoReq) returns (UserClassInfoResp);
  rpc UserClassAdd (UserClassAddReq) returns (UserClassAddResp);
  rpc UserClassDelete (UserClassDeleteReq) returns (UserClassDeleteResp);
}
```



`goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.`


所有的rpc代码都将在 [userservice.go](group%2Fuserservice%2Fuserservice.go) 里面

```shell
├── etc
│   └── anonymous_user.yaml
├── github.com
│   └── example
│       └── anonymous_user
│           ├── anonymous_user.pb.go
│           └── user_grpc.pb.go
├── go.mod
├── internal
│   ├── config
│   │   └── config.go
│   ├── logic
│   │   ├── loginlogic.go
│   │   ├── userclassaddlogic.go
│   │   ├── userclassdeletelogic.go
│   │   ├── userclassinfologic.go
│   │   ├── userclasslistlogic.go
│   │   ├── userclassupdatelogic.go
│   │   ├── userinfologic.go
│   │   ├── userinfoupdatelogic.go
│   │   ├── userlistlogic.go
│   │   ├── userroleaddlogic.go
│   │   ├── userroledeletelogic.go
│   │   ├── userroleinfologic.go
│   │   ├── userrolelistlogic.go
│   │   └── userroleupdatelogic.go
│   ├── server
│   │   └── userserviceserver.go
│   └── svc
│       └── servicecontext.go
├── anonymous_user.go
├── anonymous_user.proto
└── userservice
    └── userservice.go
```


## 带分组

首先，我们需要在 proto 文件中定义多个 service，如下：

```protobuf
syntax = "proto3";

package user;

option go_package = "github.com/example/anonymous_user";

message LoginReq{}
message LoginResp{}
message UserInfoReq{}
message UserInfoResp{}
message UserInfoUpdateReq{}
message UserInfoUpdateResp{}
message UserListReq{}
message UserListResp{}
service UserService{
  rpc Login (LoginReq) returns (LoginResp);
  rpc UserInfo (UserInfoReq) returns (UserInfoResp);
  rpc UserInfoUpdate (UserInfoUpdateReq) returns (UserInfoUpdateResp);
  rpc UserList (UserListReq) returns (UserListResp);
}

message UserRoleListReq{}
message UserRoleListResp{}
message UserRoleUpdateReq{}
message UserRoleUpdateResp{}
message UserRoleInfoReq{}
message UserRoleInfoResp{}
message UserRoleAddReq{}
message UserRoleAddResp{}
message UserRoleDeleteReq{}
message UserRoleDeleteResp{}
service UserRoleService{
  rpc UserRoleList (UserRoleListReq) returns (UserRoleListResp);
  rpc UserRoleUpdate (UserRoleUpdateReq) returns (UserRoleUpdateResp);
  rpc UserRoleInfo (UserRoleInfoReq) returns (UserRoleInfoResp);
  rpc UserRoleAdd (UserRoleAddReq) returns (UserRoleAddResp);
  rpc UserRoleDelete (UserRoleDeleteReq) returns (UserRoleDeleteResp);
}

message UserClassListReq{}
message UserClassListResp{}
message UserClassUpdateReq{}
message UserClassUpdateResp{}
message UserClassInfoReq{}
message UserClassInfoResp{}
message UserClassAddReq{}
message UserClassAddResp{}
message UserClassDeleteReq{}
message UserClassDeleteResp{}
service UserClassService{
  rpc UserClassList (UserClassListReq) returns (UserClassListResp);
  rpc UserClassUpdate (UserClassUpdateReq) returns (UserClassUpdateResp);
  rpc UserClassInfo (UserClassInfoReq) returns (UserClassInfoResp);
  rpc UserClassAdd (UserClassAddReq) returns (UserClassAddResp);
  rpc UserClassDelete (UserClassDeleteReq) returns (UserClassDeleteResp);
}
```

接着，通过 -m 指定 goctl 生成分组的代码
```goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m```

在 [client](group%2Fclient) 下生成分组的 rpc 代码



# 常见问题

1. 在使用 goctl 生成 gRPC 代码时，编写的所有 rpc 方法的请求体和响应体必须在主 proto 中声明 message，即不支持从外包外 message，
2. proto import 只支持 message 引入，不支持 service 引入。


