-- API 访问控制模块
-- 只做限流检查，然后让 nginx proxy_pass 转发

local rate_limiter = require "lib.security.rate_limiter"
local logger = require "lib.utils.logger"

local _M = {}

local log = logger.new({prefix = "APIHandler"})

-- 获取客户端 IP
local function get_client_ip()
    local headers = ngx.req.get_headers()
    local ip = headers["X-Real-IP"] or headers["X-Forwarded-For"] or ngx.var.remote_addr
    if ip then
        ip = string.match(ip, "^([^,]+)")
    end
    return ip
end

-- 访问检查（在 access 阶段调用）
-- 返回 true 表示允许通过，返回 false 表示已拒绝
function _M.check_access(config)
    local client_ip = get_client_ip()

    -- 限流检查
    if config and config.rate_limit and config.rate_limit.enabled then
        local limiter = rate_limiter.new({
            rate = config.rate_limit.rate or 100,
            burst = config.rate_limit.burst or 50,
            dict_name = config.rate_limit.dict_name or "rate_limit"
        })

        if limiter then
            local allowed, info = limiter:allow(client_ip)

            -- 设置限流响应头
            ngx.header["X-RateLimit-Limit"] = tostring(info.limit or 100)
            ngx.header["X-RateLimit-Remaining"] = tostring(info.remaining or 0)

            if not allowed then
                log:warn("Rate limit exceeded for IP: ", client_ip)
                ngx.status = 429
                ngx.header["Content-Type"] = "application/json"
                ngx.header["Retry-After"] = tostring(info.retry_after or 1)
                ngx.say('{"error": "Too many requests"}')
                return ngx.exit(429)
            end
        end
    end

    -- 设置转发给后端的请求头
    ngx.req.set_header("X-Client-IP", client_ip)
    ngx.req.set_header("X-Forwarded-For", client_ip)

    log:info("Access granted: ", ngx.var.uri, " from ", client_ip)

    -- 返回 true，让 nginx 继续执行 proxy_pass
    return true
end

return _M
