

# PostgreSQL

PostgreSQL（简称 Postgres）是一个功能强大的 开源对象关系型数据库管理系统（ORDBMS），具有企业级的功能、丰富的数据类型支持和扩展能力。它广泛应用于Web开发、数据仓库、地理信息系统（GIS）、金融、物联网等领域。

PostgreSQL 是一个免费、开源、可扩展、支持标准 SQL 的数据库系统，被认为是最先进的开源关系型数据库之一。

[PostgreSQL的官方地址](https://www.postgresql.org)

[PostgreSQL的国内社区](http://www.postgres.cn/v2/home)

[12.2中文开发文档](http://www.postgres.cn/docs/12/)



## 特点
🔍 PostgreSQL 的核心特点
- 💾 开源免费	使用 PostgreSQL 不需要授权费用，完全开源（PostgreSQL License，类 BSD）。
- 🔄 支持事务	原子性、隔离性、一致性和持久性（ACID）支持强大。
- 🧠 支持复杂查询	支持子查询、联结、窗口函数、CTE 等高级 SQL 特性。
- 🧩 可扩展性强	可以自定义数据类型、函数、操作符，甚至插件（如 TimescaleDB）。
- 🌍 多种数据类型	原生支持 JSON、XML、UUID、数组、地理数据（PostGIS）等。
- 🔐 安全性好	提供 SSL、行级安全、审计扩展等机制。
- 🚀 并发性能强	使用多版本并发控制（MVCC）来处理高并发读写。
- 📦 支持分布式	通过 FDW（外部数据包装器）连接其他数据源，或结合 Citus 实现分布式存储。


✅ 适用场景
- 企业级系统（如 ERP、CRM）
- 地理信息系统（GIS，结合 PostGIS 插件）
- 数据仓库（结合 Citus 或 TimescaleDB）
- 高并发 Web 服务（配合 Golang / Node.js / Python）


PostgreSQL的版本选择一般有两种：
- 如果为了稳定的运行，推荐使用12.x版本。
- 如果想体验新特性，推荐使用14.x版本。



## 安装
[PostgreSQL的安装、配置与使用指南](https://blog.csdn.net/qq_36433289/article/details/135058755)


### 安装

https://www.postgresql.org/

https://www.postgresql.org/download/
 
以下演示从源码安装 postgresql 

Linux平台编译安装的快捷参考(Centos平台/Pg12.2为例)：[官方安装步骤](http://www.postgres.cn/v2/download)

```shell
# 下载源码
wget https://ftp.postgresql.org/pub/source/v12.2/postgresql-12.2.tar.bz2
# 上传到指定文件夹解压，这里选择 /usr/env/pgsql/
tar xjvf postgresql-12.2.tar.bz2 
# 进入解压后的目录
cd potgresql-12.2
# 生成 Makefile（准备阶段）， 拟安装至/usr/env/postgresql 是编译输出位置。--without-readline是因为readline出错了，这里把他忽略，或尝试修复sudo yum install readline-devel
./configure --prefix=/usr/env/postgresql --without-readline
# 执行编译（构建阶段），这需要一些时间
make world
make install-world
# 服务器增加 postgres 用户，#增加新用户，系统提示要给定新用户密码
adduser postgres 
# 创建数据挂载目录 #创建数据库目录
mkdir /usr/env/postgresql/data 
# 给 postgres 用户赋予操作数据挂载目录的权限
chown -R postgres:postgres /usr/env/postgresql/data
# 切换用户 root --->>> postgres   #使用postgres帐号操作
su - postgres
# 初始化数据库
/usr/env/postgresql/bin/initdb -D /usr/env/postgresql/data 
# 启动数据库
/usr/env/postgresql/bin/pg_ctl -D /usr/env/postgresql/data -l logfile start 
# 创建一个数据库，假定数据库名为 pgsqldb
/usr/env/postgresql/bin/createdb pgsqldb 
# 进入数据库内部
/usr/env/postgresql/bin/psql pgsqldb 
# Enable automatic start【开启开机启动】
sudo systemctl enable postgresql-12
sudo systemctl start postgresql-12
```
至此，安装pgsql完毕！！！


### 修改pgsql密码

https://blog.csdn.net/qq_19283249/article/details/139048277
```
ALTER ROLE postgres WITH ENCRYPTED PASSWORD 'pgsqldb123';
```

SELECT usename, passwd FROM pg_shadow WHERE usename = 'postgres';
md5e1bdb10ca5452507a2573c4eb84d14e2

### 配置IP和远程访问

https://blog.csdn.net/qq_19283249/article/details/139048277

找到 `/usr/env/postgresql/data/postgresql.conf` 文件增加监听配置。大概在文件的第59行的位置，*表示接受所有ip可以链接postgres数据库

``` listen_addresses='*' ```


允许所有 IP 访问，找到 `/usr/env/postgresql//data/pg_hba.conf` 配置如下：

``` host    all             all             0.0.0.0/0                md5 ```

最新的配置
```shell

# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
host    all             all             0.0.0.0/0                md5
# IPv6 local connections:
host    all             all             ::1/128                 trust
# Allow replication connections from localhost, by a anonymous_user with the
# replication privilege.
local   replication     all                                     trust
host    replication     all             127.0.0.1/32            trust
host    replication     all             ::1/128                 trust
```


### 启动 postgres

./postgres -D /usr/env/postgresql/data

出现如下信息：
```shell
2025-05-02 21:03:45.597 CST [1421695] LOG:  starting PostgreSQL 12.2 on x86_64-pc-linux-gnu, compiled by gcc (GCC) 8.5.0 20210514 (Red Hat 8.5.0-4), 64-bit
2025-05-02 21:03:45.597 CST [1421695] LOG:  listening on IPv4 address "0.0.0.0", port 5432
2025-05-02 21:03:45.597 CST [1421695] LOG:  listening on IPv6 address "::", port 5432
2025-05-02 21:03:45.601 CST [1421695] LOG:  listening on Unix socket "/tmp/.s.PGSQL.5432"
2025-05-02 21:03:45.620 CST [1421696] LOG:  database system was interrupted; last known up at 2025-05-02 20:33:40 CST
2025-05-02 21:03:45.749 CST [1421696] LOG:  database system was not properly shut down; automatic recovery in progress
2025-05-02 21:03:45.752 CST [1421696] LOG:  redo starts at 0/15D7070
2025-05-02 21:03:45.752 CST [1421696] LOG:  invalid record length at 0/15D7158: wanted 24, got 0
2025-05-02 21:03:45.752 CST [1421696] LOG:  redo done at 0/15D7120
2025-05-02 21:03:45.763 CST [1421695] LOG:  database system is ready to accept connections
```

### 链接 postgres

IP

端口：5432

用户名：postgres，密码：pgsqldb123，数据库：pgsqldb





## 学习资源
1. [一天学完 PostgreSQL](https://www.bilibili.com/video/BV1av411r7yB)
2. [如何搭建一套学习、开发PostgreSQL的环境](https://www.cnblogs.com/88223100/p/How-to-build-an-environment-to-learn-and-develop-PostgreSQL.html)
3. 

