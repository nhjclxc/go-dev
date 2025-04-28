


# [beego-web-02-router学习]

[web模块-路由-Web 注册路由最佳实践](https://beegodoc.com/zh/developing/web/router/best_practice.html)

## 快速开始
启动步骤
1. 创建项目：`bee api web_01_router[]()`
2. 进入 `cd .\web_01_router\[]()` 目录
3. 更新包：`go mod tidy`
5. 运行项目执行 `bee run`

## 目录结构

```markdown
web_05_router
├─conf
├─controllers
│  ├─admin
│  │  ├─goods
│  │  │  └─AdminGoodsController.go
│  │  └─user
│  │     └─AdminUserController.go
│  └─api
│      ├─goods
│      │  └─ApiGoodsController.go
│      └─user
│         └─ApiUserController.go
├─models
│  ├─Goods.go
│  └─User.go
├─routers
│  ├─admin
│  │  ├─goods
│  │  │   └─APIGoodsRouter.go
│  │  └─user
│  │      └─APIUserRouter.go
│  └─api
│      ├─goods
│      │  └─APIGoodsRouter.go
│      └─user
│         └─APIUserRouter.go
├─tests
│   └─default_test.go
go.mod
main.go
```


