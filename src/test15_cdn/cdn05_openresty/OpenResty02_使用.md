启动 OpenResty 只是第一步，它本质上是 **带 Lua 扩展的 Nginx**，所以使用方式和 Nginx 类似，但可以运行 Lua 脚本。下面给你一个完整的使用流程和常见操作方法：

---

## **1️⃣ 启动 OpenResty**

如果你已经安装好，可以选择：

* **后台启动（服务形式）**

```bash
sudo brew services start openresty/brew/openresty
```

* **前台启动（调试用）**

```bash
/opt/homebrew/opt/openresty/bin/openresty -g 'daemon off;'
```

* **重载配置**（修改 `nginx.conf` 后）

```bash
openresty -s reload
```

* **停止 OpenResty**

```bash
openresty -s stop
```

> 命令都可用 `openresty` 或 `nginx`，因为 OpenResty 的 Nginx 已经替换了原生 Nginx。

---

## **2️⃣ 配置 OpenResty**

配置文件默认在：

```
/opt/homebrew/etc/openresty/nginx.conf
```

你可以在 `http` 块或者 `server` 块里添加 Lua 脚本，比如：

```nginx
http {
    server {
        listen 8080;
        location /hello {
            content_by_lua_block {
                ngx.say("Hello OpenResty!")
            }
        }
    }
}
```

* `content_by_lua_block` 可以直接写 Lua 脚本
* Lua 脚本可以访问 `ngx` 全局对象，例如：请求信息、响应输出、共享内存等

---

## **3️⃣ 测试访问**

启动 OpenResty 后，访问：

```
http://127.0.0.1:8080/hello
```

应该返回：

```
Hello OpenResty!
```

---

## **4️⃣ 使用 Lua 模块**

1. LuaRocks 可以安装模块：

```bash
luarocks install luasocket --local
```

2. 在 OpenResty 中加载：

```nginx
location /time {
    content_by_lua_block {
        local socket = require("socket")
        ngx.say(socket.gettime())
    }
}
```

> 注意：使用 Lua 模块时需要确保 OpenResty 使用的是 **Lua 5.1 / LuaJIT 兼容模块**。

---

## **5️⃣ 常用命令总结**

| 命令                    | 作用             |
| --------------------- | -------------- |
| `openresty -t`        | 测试配置是否正确       |
| `openresty -s reload` | 重新加载配置         |
| `openresty -s stop`   | 停止 OpenResty   |
| `openresty -s quit`   | 优雅停止 OpenResty |
| `openresty -v`        | 查看版本           |

---

💡 **小技巧**

* 可以把 `nginx.conf` 拆成多个 `*.conf` 文件，放在 `/opt/homebrew/etc/openresty/conf.d/`，在主配置里用 `include conf.d/*.conf;`
* Lua 脚本可直接放在 `lua/` 目录，配置里用 `content_by_lua_file` 引入

---

如果你愿意，我可以帮你写一个 **最简单可运行的 OpenResty 项目目录结构示例**，包括 Lua 脚本和配置，直接启动就能访问。你希望我写吗？




在 OpenResty 下启动指定配置文件，你可以直接用 `-c` 参数（和普通 nginx 一样）：

假设你的目录结构是：

```
resty01_test/
├── conf
│   └── nginx.conf
```

---

# ✅ 方法 1：命令行指定配置文件

```bash
openresty -c /Users/lxc20250729/lxc/code/go-dev/src/test15_cdn/resty01_test/conf/nginx.conf
```

如果你在当前目录，也可以用相对路径：

```bash
openresty -c ./conf/nginx.conf
```

---

# ✅ 方法 2：启动时加 `-p` 指定 prefix

`-p` 用于指定 **prefix 路径**，即 Nginx 的 root 目录，用来解析 `logs/`、`html/` 等相对路径。

```bash
openresty -p $(pwd) -c conf/nginx.conf
```

这时：

* Nginx 会把 `logs/`、`html/` 等相对路径都基于当前目录
* 可以避免 docker 外挂路径找不到问题

---

# ✅ 方法 3：后台启动

```bash
openresty -p $(pwd) -c conf/nginx.conf -g "daemon on;"
```

* `daemon on;` 表示后台启动
* `-g` 可以覆盖 nginx.conf 全局指令

---

# 🔧 检查配置是否正确

```bash
openresty -t -c ./conf/nginx.conf
```

输出类似：

```
nginx: the configuration file ./conf/nginx.conf syntax is ok
nginx: configuration file ./conf/nginx.conf test is successful
```

---

# 🔑 总结

| 参数               | 作用                     |
| ---------------- | ---------------------- |
| `-c <file>`      | 指定 nginx.conf 文件       |
| `-p <path>`      | 指定 prefix，控制日志和相对路径    |
| `-g <directive>` | 覆盖全局配置（daemon、worker等） |

---

如果你愿意，我可以帮你写一个 **完整命令示例**，让 OpenResty 直接用你的本地 `conf/nginx.conf` 启动，并把 logs 写到 `./logs/` 文件夹里。


如果缺少nginx相关文件（如：mime.types），那么去[nginx官网](https://nginx.org/en/download.html)下载

openresty检查nginx.conf文件配置是否正确：`openresty -t -p $(pwd) -c conf/nginx.conf`
openresty启动nginx.conf：`openresty -p $(pwd) -c conf/nginx.conf`
openresty后台启动nginx.conf：`openresty -p $(pwd) -c conf/nginx.conf -g "daemon on;"`

强制停止所有 openresty / nginx：`pkill -9 nginx` 或 `pkill -9 openresty`

重启openresty：` pkill -9 nginx && openresty -p $(pwd) -c conf/nginx.conf `

pkill -9 nginx; openresty -p $(pwd) -c conf/nginx_02.conf