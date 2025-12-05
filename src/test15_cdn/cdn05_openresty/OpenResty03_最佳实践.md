# lua + openresty


[学习网站1](https://github.com/openresty/lua-nginx-module?tab=readme-ov-file#nginx-api-for-lua)

[学习网站2](https://github.com/bungle/awesome-resty)

[学习网站3](https://docs.nginx.com/#nginx-api-for-lua)

一些学习链接：
[https://juejin.cn/post/7345387782112133160](https://juejin.cn/post/7345387782112133160)、
[https://www.cnblogs.com/Chary/p/18952730](https://www.cnblogs.com/Chary/p/18952730)、


[OpenResty执行流程与阶段详解](https://feixiang.blog.csdn.net/article/details/136590258)

[resty+lua实现一个网关框架](https://feixiang.blog.csdn.net/article/details/136658438)






## 出现【module 'resty.http' not found:】

在[lua-resty-http的github官网下载](https://github.com/ledgetech/lua-resty-http)库文件，将该仓库`lua-resty-http/lib/resty/`下的三个文件`http.lua`、`http_connect.lua`、`http_headers.lua`放入当前项目`lua/resty`下面，如下所示：
```shell
➜  resty git:(main) ✗ pwd
/Users/lxc20250729/lxc/code/go-dev/src/test15_cdn/openresty01/lua/resty
➜  resty git:(main) ✗ tree
.
├── http.lua
├── http_connect.lua
└── http_headers.lua

0 directories, 3 files

```

## 出现【failed to load module `resty.openssl.*`, mTLS isn't supported without lua-resty-openssl:】 或 【module 'resty.openssl.x509.chain' not found:】


在[lua-resty-openssl的github官网下载](https://github.com/fffonion/lua-resty-openssl)库文件，将该仓库`lua-resty-openssl/lib/resty/`下`openssl.lua`的文件和`openssl`文件夹放入当前项目`lua/resty`下面，如下所示：
```shell
➜  resty git:(main) ✗ pwd
/Users/lxc20250729/lxc/code/go-dev/src/test15_cdn/openresty01/lua/resty
➜  resty git:(main) ✗ tree
.
├── http.lua
├── http_connect.lua
├── http_headers.lua
├── openssl
│   ├── asn1.lua
│   ├── auxiliary
│   ├── ...
│   └── x509
│       ├── altname.lua
│       ├── ...
│       └── store.lua
└── openssl.lua

8 directories, 82 files
```


## nginx内部变量
```shell
$arg_name 请求中的name参数
$args 请求中的参数
$binary_remote_addr 远程地址的二进制表示
$body_bytes_sent  已发送的消息体字节数
$content_length HTTP请求信息里的"Content-Length"
$content_type 请求信息里的"Content-Type"
$document_root  针对当前请求的根路径设置值
$document_uri 与$uri相同; 比如 /test2/test.php
$host 请求信息中的"Host"，如果请求中没有Host行，则等于设置的服务器名
$hostname 机器名使用 gethostname系统调用的值
$http_cookie  cookie 信息
$http_referer 引用地址
$http_user_agent  客户端代理信息
$http_via 最后一个访问服务器的Ip地址。
$http_x_forwarded_for 相当于网络访问路径
$is_args  如果请求行带有参数，返回“?”，否则返回空字符串
$limit_rate 对连接速率的限制
$nginx_version  当前运行的nginx版本号
$pid  worker进程的PID
$query_string 与$args相同
$realpath_root  按root指令或alias指令算出的当前请求的绝对路径。其中的符号链接都会解析成真是文件路径
$remote_addr  客户端IP地址
$remote_port  客户端端口号
$remote_user  客户端用户名，认证用
$request  用户请求
$request_body 这个变量（0.7.58+）包含请求的主要信息。在使用proxy_pass或fastcgi_pass指令的location中比较有意义
$request_body_file  客户端请求主体信息的临时文件名
$request_completion 如果请求成功，设为"OK"；如果请求未完成或者不是一系列请求中最后一部分则设为空
$request_filename 当前请求的文件路径名，比如/opt/nginx/www/test.php
$request_method 请求的方法，比如"GET"、"POST"等
$request_uri  请求的URI，带参数; 比如http://localhost:88/test1/
$scheme 所用的协议，比如http或者是https
$server_addr  服务器地址，如果没有用listen指明服务器地址，使用这个变量将发起一次系统调用以取得地址(造成资源浪费)
$server_name  请求到达的服务器名
$server_port  请求到达的服务器端口号
$server_protocol  请求的协议版本，"HTTP/1.0"或"HTTP/1.1"
$uri  请求的URI，可能和最初的值有不同，比如经过重定向之类的

```

## openresty的执行流程
```shell
2、Nginx请求处理的11个阶段
Nginx处理请求的过程一共划分为11个阶段，按照执行顺序依次是post-read、server-rewrite、
find-config、rewrite、post-rewrite、 preaccess、access、post-access、try-files、content、log。

所以，整个请求的过程，是按照不同的阶段执行的，在某个阶段执行完该阶段的指令之后，再进行下一个阶段的指令执行。

1、post-read
读取请求内容阶段，nginx读取并解析完请求头之后就立即开始运行；
例如模块 ngx_realip 就在 post-read 阶段注册了处理程序，
它的功能是迫使 Nginx 认为当前请求的来源地址是指定的某一个请求头的值。

2、server-rewrite
server请求地址重写阶段；当ngx_rewrite模块的set配置指令直接书写在server配置块中时，
基本上都是运行在server-rewrite 阶段

3、find-config
配置查找阶段，这个阶段并不支持Nginx模块注册处理程序，
而是由Nginx核心来完成当前请求与location配置块之间的配对工作。

4、rewrite
location请求地址重写阶段，当ngx_rewrite指令用于location中，就是再这个阶段运行的；
另外ngx_set_misc(设置md5、encode_base64等)模块的指令，
还有ngx_lua模块的set_by_lua指令和rewrite_by_lua指令也在此阶段。

5、post-rewrite
请求地址重写提交阶段，当nginx完成rewrite阶段所要求的内部跳转动作，如果rewrite阶段有这个要求的话；

6、preaccess
访问权限检查准备阶段，ngx_limit_req和ngx_limit_zone在这个阶段运行，
ngx_limit_req可以控制请求的访问频率，ngx_limit_zone可以控制访问的并发度；

7、access
访问权限检查阶段，标准模块ngx_access、第三方模块ngx_auth_request以及第三方模块ngx_lua的access_by_lua
指令就运行在这个阶段。配置指令多是执行访问控制相关的任务，如检查用户的访问权限，检查用户的来源IP是否合法；

8、post-access
访问权限检查提交阶段；主要用于配合access阶段实现标准ngx_http_core模块提供的配置指令satisfy的功能。
satisfy all(与关系),satisfy any(或关系)

9、try-files
配置项try_files处理阶段；专门用于实现标准配置指令try_files的功能,
如果前 N-1 个参数所对应的文件系统对象都不存在，
try-files 阶段就会立即发起“内部跳转”到最后一个参数(即第 N 个参数)所指定的URI.

10、content
内容产生阶段，是所有请求处理阶段中最为重要的阶段，
因为这个阶段的指令通常是用来生成HTTP响应内容并输出 HTTP 响应的使命；

11、log
日志模块处理阶段；记录日志

NGX_HTTP_POST_READ_PHASE:
#读取请求内容阶段
NGX_HTTP_SERVER_REWRITE_PHASE:
#Server请求地址重写阶段
NGX_HTTP_FIND_CONFIG_PHASE:
#配置查找阶段:
NGX_HTTP_REWRITE_PHASE:
#Location请求地址重写阶段，常用
NGX_HTTP_POST_REWRITE_PHASE:
#请求地址重写提交阶段
NGX_HTTP_PREACCESS_PHASE:
#访问权限检查准备阶段
NGX_HTTP_ACCESS_PHASE:
#访问权限检查阶段，常用
NGX_HTTP_POST_ACCESS_PHASE:
#访问权限检查提交阶段
NGX_HTTP_TRY_FILES_PHASE:
#配置项try_files处理阶段
NGX_HTTP_CONTENT_PHASE:
#内容产生阶段 最常用
NGX_HTTP_LOG_PHASE:
#日志模块处理阶段 常用

到此，应该明白Nginx 的conf中指令的书写顺序和执行顺序是两码事。

有些阶段是支持 Nginx 模块注册处理程序，有些阶段并不可以。

最常用的是 rewrite阶段，access阶段 以及 content阶段；
不支持 Nginx 模块注册处理程序的阶段 find-config, post-rewrite, post-access,
主要是 Nginx 核心完成自己的一些逻辑。
原文链接：https://blog.csdn.net/A_art_xiang/article/details/136590258


```

## 