# [OpenResty官网](https://openresty.org/cn/installation.html)




openresty安装路径：/opt/homebrew/Cellar/openresty/1.27.1.2_1
可执行文件（nginx / openresty）：/opt/homebrew/opt/openresty/bin/openresty、 /opt/homebrew/opt/openresty/bin/nginx
配置文件：/opt/homebrew/etc/openresty/nginx.conf
后台启动：sudo brew services start openresty/brew/openresty
前台启动：/opt/homebrew/opt/openresty/bin/openresty -g 'daemon off;'




好的，我给你整理一份 **适合 macOS + OpenResty 的 Lua 开发环境完整配置教程**，包含安装 Lua、LuaRocks、LuaJIT、OpenResty 及测试脚本。步骤按顺序来，保证可以顺利开发。

---

# **Mac + OpenResty Lua 开发环境配置教程**

## **1. 安装 Homebrew（如果还没有安装）**

Homebrew 是 macOS 上最方便的包管理工具。

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew --version
```

---

## **2. 安装 LuaJIT**

OpenResty 使用 LuaJIT 作为 Lua 的执行引擎，性能比标准 Lua 高。

```bash
brew install luajit
```

**验证安装：**

```bash
luajit -v
```

输出类似：

```
LuaJIT 2.1.0-beta3 -- Copyright (C) 2005-2023 Mike Pall
```

> 注意：不要用 macOS 自带的 `/usr/bin/lua`，版本可能是 5.1，且不够新。

---

## **3. 安装 LuaRocks**

LuaRocks 是 Lua 的包管理工具，方便安装各种 Lua 模块。

```bash
brew install luarocks
```

**配置 LuaRocks 使用 LuaJIT：**

```bash
luarocks config --lua-version=5.1
```

> LuaJIT 默认兼容 Lua 5.1。

**验证 LuaRocks：**

```bash
luarocks --version
```

**安装测试模块：**

```bash
luarocks install luasocket
```

**测试模块是否可用：**

```bash
luajit -e "local socket = require('socket'); print('luasocket loaded')"
```

输出：

```
luasocket loaded
```

---

## **4. 安装 OpenResty**

OpenResty 是基于 Nginx + LuaJIT 的高性能 Web 平台。

```bash
brew install openresty/brew/openresty --without-geoip
```

**验证安装：**

```bash
openresty -v
```

输出类似：

```
nginx version: openresty/1.23.4.2
```

**启动 OpenResty：**

```bash
openresty
```

> 默认配置文件路径：
>
> `/usr/local/etc/openresty/nginx.conf`

**停止 OpenResty：**

```bash
openresty -s stop
```

---

## **5. 测试 Lua 脚本**

1. 创建一个测试 Lua 脚本，例如 `test.lua`：

```lua
-- test.lua
print("Hello, OpenResty + LuaJIT!")

local socket = require("socket")
print("Current time:", socket.gettime())
```

2. 在终端运行：

```bash
luajit test.lua
```

输出类似：

```
Hello, OpenResty + LuaJIT!
Current time: 1700000000.123
```

---

## **6. 测试 OpenResty 中 Lua 脚本**

1. 修改 OpenResty 配置文件 `/usr/local/etc/openresty/nginx.conf`：

```nginx
http {
    server {
        listen 8080;

        location /lua {
            default_type 'text/plain';
            content_by_lua_block {
                ngx.say("Hello from OpenResty Lua!")
            }
        }
    }
}
```

2. 启动 OpenResty：

```bash
openresty
```

3. 访问测试：

```bash
curl http://127.0.0.1:8080/lua
```

输出：

```
Hello from OpenResty Lua!
```

> ✅ 成功说明 Lua + LuaJIT + OpenResty 配置完成。

---

## **7. 小技巧与注意事项**

* LuaRocks 安装的模块路径可能和 OpenResty LuaJIT 不同，可通过 `package.path` / `package.cpath` 调整。
* 开发中可以用 `resty` 命令快速运行 Lua 脚本：

```bash
resty test.lua
```

* 如果需要最新 Lua 版本，也可以通过 Homebrew 安装 `lua`，但 OpenResty 默认使用 LuaJIT（Lua 5.1 兼容）。

---

我可以帮你顺手写一份 **带 LuaRocks 模块路径配置、适用于 OpenResty 的 Lua 开发模板**，直接拷贝就能跑。

你希望我帮你写吗？
