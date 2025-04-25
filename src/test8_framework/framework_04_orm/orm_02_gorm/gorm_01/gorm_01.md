
### 官方文档

[GORM官方文档](https://gorm.io/zh_CN/docs/index.html)
[GitHub源码](https://github.com/go-gorm/gorm)

### 依赖的库

gin：`go get -u github.com/gin-gonic/gin`
gorm 库：`go get -u gorm.io/gorm`
gorm封装的mysql驱动：`go get -u gorm.io/driver/mysql`
属性拷贝库,DTO to Model：`go get github.com/jinzhu/copier`
数据库转struct工具：`go get -u github.com/xxjwxc/gormt`


```mysql
CREATE TABLE `gen_table2`
(
    `table_id`      bigint NOT NULL AUTO_INCREMENT COMMENT '代码生成业务表主键id',
    `table_name2`   varchar(200) COMMENT '表名称',
    `table_comment` varchar(500) COMMENT '表描述',
    `sort`          int         DEFAULT NULL COMMENT '排序',
    `state`         tinyint     DEFAULT NULL COMMENT '状态（0=删除，1=在用）',
    `create_by`     varchar(64) DEFAULT '' COMMENT '创建者',
    `create_time`   datetime    DEFAULT NULL COMMENT '创建时间',
    `update_by`     varchar(64) DEFAULT '' COMMENT '更新者',
    `update_time`   datetime    DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`table_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb3 COMMENT ='代码生成业务表';


CREATE TABLE `gen_table_column2`
(
    `column_id`      bigint NOT NULL AUTO_INCREMENT COMMENT '代码生成业务字段表主键',
    `table_id`       bigint       DEFAULT NULL COMMENT '归属表编号',
    `column_name`    varchar(200) DEFAULT NULL COMMENT '列名称',
    `column_comment` varchar(500) DEFAULT NULL COMMENT '列描述',
    `column_type`    varchar(100) DEFAULT NULL COMMENT '列类型',
    `is_query`       char(1)      DEFAULT NULL COMMENT '是否查询字段（1是）',
    `create_by`      varchar(64)  DEFAULT '' COMMENT '创建者',
    `create_time`    datetime     DEFAULT NULL COMMENT '创建时间',
    `update_by`      varchar(64)  DEFAULT '' COMMENT '更新者',
    `update_time`    datetime     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`column_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb3 COMMENT ='代码生成业务表字段';
```