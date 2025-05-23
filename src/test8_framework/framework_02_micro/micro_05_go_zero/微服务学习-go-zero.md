
# go-zero

[源代码](https://github.com/zeromicro/go-zero.git)

[官网](https://go-zero.dev/)

[官网文档](https://go-zero.dev/docs/concepts/overview)

[zero-doc](https://github.com/zeromicro/zero-doc.git)

[goctl-swagger](https://github.com/zeromicro/goctl-swagger.git)


## go-zero 环境搭建

可参考[go-zero官方安装指南](https://go-zero.dev/docs/tasks)

### 2.1、go-zero源码
安装：```go get -u github.com/zeromicro/go-zero```


### 2.2、goctl

[goctl源码地址](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)

安装 goctl 工具：```go install github.com/zeromicro/go-zero/tools/goctl@latest```,在https://github.com/zeromicro/go-zero/blob/master/readme-cn.md文件里面给出的地址


之后在$GOPATH/bin目录下将生成一个goctl.exe文件，进入cmd检查安装是否成功：`goctl --version`，输出：`goctl version 1.8.3 windows/amd64`

### 2.3、protoc (protobuf编译器)

protoc.exe 工具：[https://github.com/protocolbuffers/protobuf/releases/tag/v31.0-rc2](https://objects.githubusercontent.com/github-production-release-asset-2e65be/23357588/04ba4af3-1503-4779-90a3-f4bad4044593?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=releaseassetproduction%2F20250506%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250506T080411Z&X-Amz-Expires=300&X-Amz-Signature=d6dbc3ff21c57647af20c831915d0c96b33294d13dea5f6e99832fbde511bbb3&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3Dprotoc-31.0-rc-2-win64.zip&response-content-type=application%2Foctet-stream)

将下载好的protoc.exe放在$GOPATH/bin目录下，进入cmd检查安装是否成功：`protoc --version`，输出：`libprotoc 31.0-rc2`

此外，关于proto语法的相关知识详细，可以参考 go-zero里面的 [proto语法](https://go-zero.dev/docs/tasks/dsl/proto) 一节。

### 2.4、protoc-gen-go

protoc-gen-go.exe 工具：使用命令安装 ```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest``` 或 （go install google.golang.org/protobuf/protoc-gen-go@latest）

之后在$GOPATH/bin目录下将生成一个protoc-gen-go.exe文件


### 2.5、protoc-gen-go-grpc

protoc-gen-go-grpc.exe 工具：使用命令安装 ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```

之后在$GOPATH/bin目录下将生成一个protoc-gen-go-grpc.exe文件


### 2.6、其他

如果有要使用grpc-gateway，也请安装如下两个插件 , 没有使用就忽略下面2个插件

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest


## 开源学习仓库

### go-zero-looklook


[go-zero-looklook](https://github.com/Mikaelemmmm/go-zero-looklook.git)是go-zero框架的最佳实践

[go-zero-looklook对应的最佳实际教程文档](https://github.com/Mikaelemmmm/go-zero-looklook/tree/main/doc/chinese)

[go-zero-looklook 视频教程](https://www.bilibili.com/video/BV1P3411p79J)

[go-zero-looklook 社区](https://www.dongaigc.com/p/Mikaelemmmm/go-zero-looklook)



## 其他社区 

[go-zero学习 第一章 基础](https://blog.csdn.net/Mr_XiMu/article/details/131294294)

[go-zero入门，看这一篇就够了](https://juejin.cn/post/7225565801791799354)


[03. go-zero简介及如何学go-zero](https://www.cnblogs.com/haima/p/16057786.html)

[Golang 慢慢学-go-zero](https://haimait.top/docs/golang/go-zero)

[]()



## 学习视频 

[【码神之路】go-zero框架教程，十年大厂程序员讲解，通俗易懂](https://www.bilibili.com/video/BV1Fg4y1W7Na/)


[一、go-zero简介及如何学go-zero](https://www.bilibili.com/video/BV1LS4y1U72n/)
[三十三、go-zero-looklook开发教程](https://www.bilibili.com/list/389552232?sid=2122723&oid=214031571&bvid=BV1Ea411J7nj)


[【go-zero教程】01-快速入门，2024新版教程，十年大厂程序员讲解，通俗易懂](https://www.bilibili.com/video/BV1vRxzefExM/)

[go-zero零基础入门教程|go微服务开发必学教程](https://www.bilibili.com/video/BV1kM411X7Cp/)

[【项目实战】基于go-zero（微服务）实现物联网平台](https://www.bilibili.com/video/BV13G4y1R71m/)



[【项目实战】基于Go-zero、Xorm的网盘系统](https://www.bilibili.com/video/BV1cr4y1s7H4/)，开源地址：https://gitee.com/getcharzp/cloud-disk

[#101 晓黑板 go-zero 微服务框架的架构设计](https://www.bilibili.com/video/BV1rD4y127PD/)  https://talkgo.org/t/topic/729

[基于go-zero的Go微服务实战干货教程，第一课 项目介绍和核心技术介绍](https://www.bilibili.com/video/BV1op4y177iS/)

[]()

[]()



## 开源项目

[基于 go-zero 的开源项目](https://github.com/zeromicro/go-zero/issues/653) 从这里面找，有很多


[zero-examples](https://github.com/zeromicro/zero-examples)

[awesome-zero](https://github.com/zeromicro/awesome-zero)


