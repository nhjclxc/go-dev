-- 应用入口文件
local logger = require "lib.utils.logger"

local _M = {
    _VERSION = '1.0.0'
}

local log = logger.new({prefix = "App"})
local config

-- 加载配置
function _M.load_config()
    if not config then
        config = require "app.config.default"
        log:info("Configuration loaded")
    end
    return config
end

-- 路由表
local routes = {
    {
        pattern = "^/api/",
        handler = "app.handlers.api_handler"
    },
    {
        pattern = "^/health$",
        handler = function(cfg)
            ngx.header["Content-Type"] = "text/plain"
            ngx.say("OK")
        end
    }
}

-- 路由匹配
local function match_route(uri)
    for _, route in ipairs(routes) do
        if ngx.re.match(uri, route.pattern, "jo") then
            return route.handler
        end
    end
    return nil
end

-- 处理请求
function _M.handle_request()
    -- 确保配置已加载
    if not config then
        config = _M.load_config()
    end

    local uri = ngx.var.uri
    log:debug("Handling request: ", uri)

    -- 匹配路由
    local handler = match_route(uri)

    if not handler then
        -- 默认处理：显示欢迎页
        ngx.header["Content-Type"] = "text/html"
        ngx.say([[
<!DOCTYPE html>
<html>
<head>
    <title>CDN Gateway V2</title>
    <style>
        body { font-family: Arial; margin: 50px; }
        h1 { color: #333; }
        .info { background: #f0f0f0; padding: 15px; border-radius: 5px; }
        code { background: #e0e0e0; padding: 2px 6px; border-radius: 3px; }
    </style>
</head>
<body>
    <h1>OpenResty CDN Gateway V2</h1>
    <div class="info">
        <p><strong>Version:</strong> 1.0.0</p>
        <p><strong>Architecture:</strong> Modular Lua Application</p>
    </div>

    <h2>Available Endpoints:</h2>
    <ul>
        <li><code>GET /health</code> - Health check</li>
        <li><code>GET /api/*</code> - API proxy with rate limiting</li>
        <li><code>GET /static/*</code> - Static resources with CDN caching</li>
    </ul>
</body>
</html>
        ]])
        return
    end

    -- 执行处理器
    if type(handler) == "string" then
        -- 动态加载模块
        local ok, module = pcall(require, handler)
        if not ok then
            log:error("Failed to load handler: ", handler, " - ", module)
            ngx.status = ngx.HTTP_INTERNAL_SERVER_ERROR
            ngx.say('{"error": "Handler not found"}')
            return
        end
        handler = module.handle
    end

    -- 调用处理器
    local ok, err = pcall(handler, config)
    if not ok then
        log:error("Handler execution failed: ", err)
        ngx.status = ngx.HTTP_INTERNAL_SERVER_ERROR
        ngx.header["Content-Type"] = "application/json"
        ngx.say('{"error": "Internal server error"}')
    end
end

return _M
