# OpenResty CDN 网关

基于 OpenResty + Lua 的 CDN 边缘节点网关。

## 功能

- **链路追踪**：自动生成/传递 Trace ID
- **限流保护**：基于令牌桶算法的请求限流
- **访问控制**：IP 黑白名单、Token 验证
- **请求日志**：JSON 格式日志，包含 Trace ID

## 目录结构

```
.
├── docker-compose.yml      # Docker 编排
├── conf/
│   └── nginx.conf          # OpenResty 配置
└── lua/                    # Lua 模块
    ├── app/
    │   ├── init.lua        # 应用入口
    │   ├── config/         # 配置
    │   └── handlers/       # 处理器
    └── lib/                # 可复用库
        ├── cache/          # 缓存模块
        ├── security/       # 安全模块
        └── utils/          # 工具模块
```

## 快速开始

### 1. 配置后端地址

编辑 `conf/nginx.conf`，修改 upstream：

```nginx
upstream backend {
    server your-backend-host:80;  # 改为你的后端地址
    keepalive 32;
}
```

### 2. 启动服务

```bash
docker-compose up -d
```

### 3. 测试

```bash
# 健康检查
curl http://localhost/health

# API 请求（查看 Trace ID）
curl -i http://localhost/api/test

# 查看日志
docker-compose logs -f openresty
```

## 核心模块

### 链路追踪 (`lib/utils/tracing.lua`)

```lua
-- 在 nginx.conf 中使用
rewrite_by_lua_block {
    local tracing = require "lib.utils.tracing"
    tracing.quick_init()  -- 自动生成/接收 Trace ID
}
```

**功能**：
- 自动生成 Trace ID（如果客户端没有）
- 接收客户端传来的 Trace ID（X-Trace-ID, traceparent）
- 传递给后端服务
- 在响应头返回

### 限流器 (`lib/security/rate_limiter.lua`)

```lua
local rate_limiter = require "lib.security.rate_limiter"
local limiter = rate_limiter.new({
    rate = 100,   -- 100 req/s
    burst = 50    -- 突发 50
})

local allowed, info = limiter:allow(client_ip)
```

### 缓存 (`lib/cache/memory.lua`)

```lua
local cache = require "lib.cache.memory"
local c = cache.new("cdn_cache", 300)  -- 缓存 5 分钟

c:set("key", "value")
local value = c:get("key")
```

## 配置说明

### 环境变量

```bash
# docker-compose.yml 或 .env
BACKEND_HOST=your-backend-host
BACKEND_PORT=80
```

### Nginx 配置

| 配置项 | 说明 |
|-------|------|
| `lua_shared_dict cdn_cache 100m` | CDN 缓存，100MB |
| `lua_shared_dict rate_limit 10m` | 限流字典，10MB |

### 日志格式

JSON 格式，包含 Trace ID：

```json
{
  "trace_id": "abc123",
  "time": "2025-12-03T10:30:15+00:00",
  "request": "GET /api/test HTTP/1.1",
  "status": 200,
  "body_bytes_sent": 1024,
  "request_time": 0.005
}
```

## 扩展

### 添加新的 Lua 模块

1. 在 `lua/lib/` 下创建模块
2. 在 nginx.conf 中使用

```lua
-- lua/lib/my_module.lua
local _M = {}

function _M.do_something()
    -- ...
end

return _M
```

```nginx
# nginx.conf
access_by_lua_block {
    local my_module = require "lib.my_module"
    my_module.do_something()
}
```

### 添加 HTTPS

```nginx
server {
    listen 443 ssl http2;
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    # ...
}
```

## 执行阶段

OpenResty Lua 执行顺序：

```
init_by_lua         → 启动时执行一次
    ↓
rewrite_by_lua      → URL 重写，添加 Trace ID
    ↓
access_by_lua       → 访问控制，限流
    ↓
proxy_pass          → 转发后端
    ↓
header_filter_by_lua → 修改响应头
    ↓
log_by_lua          → 记录日志
```

## 注意事项

1. **模块缓存**：Lua 模块只加载一次，修改后需重启
2. **日志级别**：生产环境建议 `error_log ... warn`
3. **共享内存**：根据流量调整 `lua_shared_dict` 大小

## 常用命令

```bash
# 启动
docker-compose up -d

# 停止
docker-compose down

# 重启（修改配置后）
docker-compose restart

# 查看日志
docker-compose logs -f openresty

# 进入容器
docker-compose exec openresty sh

# 测试配置
docker-compose exec openresty nginx -t
```
