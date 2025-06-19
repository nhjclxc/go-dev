

https://go-zero.dev/docs/tutorials/mysql/connection


go-zero 提供了一个强大的 sqlx 工具，用于操作数据库。 所有 SQL 相关操作的包在 github.com/zeromicro/go-zero/core/stores/sqlx


# 一、使用go-zero的代码生成

```shell
goctl model mysql datasource -url="root:root123@tcp(127.0.0.1:3306)/test" -table="tab_user" -dir="./internal/model"
```


