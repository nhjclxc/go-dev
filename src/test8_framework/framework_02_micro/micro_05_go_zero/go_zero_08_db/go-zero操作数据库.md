

# 一、MySQL数据库

任务目标
- 了解 github.com/zeromicro/go-zero/core/stores/sqlx 包的使用。
- 根据 sqlx 创建一个 sql 链接。


## 1.1、代码生成

项目根目录下执行：
```shell
goctl model mysql datasource -url="root:root123@tcp(127.0.0.1:3306)/test" -table="tab_user" -dir="./model"
```

文件介绍：
- [tabusermodel.go](db_01_mysql%2Fmodel%2Ftabusermodel.go)：这个就类似于 MyBatis 的Mapper层，提供了操作的接口
- [tabusermodel_gen.go](db_01_mysql%2Fmodel%2Ftabusermodel_gen.go)：这个就类似于 MyBatis 的model数据定义
- [vars.go](db_01_mysql%2Fmodel%2Fvars.go)：



  


# 1、代码生成