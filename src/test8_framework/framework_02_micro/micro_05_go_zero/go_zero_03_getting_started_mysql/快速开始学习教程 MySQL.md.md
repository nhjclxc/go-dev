
[mysql 代码生成](https://go-zero.dev/docs/tasks/cli/mysql)



# 一、准备 sql 文件

准备一个sql文件，用于goctl根据sql的ddl定义来生成 go-zero 的代码。

详细的sql文件看 [tab_user.sql](model%2Fuser%2Ftab_user.sql)

将 tab_user.sql 文件 移动到 [user](model%2Fuser) 目录


# 二、执行 goctl 命令生成 model 代码

进入 [user](model%2Fuser) 文件夹，执行 `goctl model mysql ddl --src tab_user.sql --dir . --style=goZero`

将生成三个文件，tabUserModel.go 、tabUserModel_gen.go 、vars.go

```
├── tab_user.sql
├── tabUserModel.go
├── tabUserModel_gen.go
└── vars.go
```

