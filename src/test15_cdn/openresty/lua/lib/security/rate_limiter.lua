-- 限流器模块
-- 基于令牌桶算法，直接使用 ngx.shared.dict

local _M = {
    _VERSION = '0.1.0'
}

-- 创建限流器实例
function _M.new(opts)
    opts = opts or {}

    local dict_name = opts.dict_name or "rate_limit"
    local dict = ngx.shared[dict_name]

    if not dict then
        ngx.log(ngx.ERR, "shared dict not found: ", dict_name)
        return nil, "shared dict not found"
    end

    local self = {
        rate = opts.rate or 100,           -- 每秒请求数
        burst = opts.burst or 50,          -- 突发容量
        dict = dict,
        key_prefix = opts.key_prefix or "rl:"
    }

    return setmetatable(self, {__index = _M})
end

-- 令牌桶算法
function _M:allow(key)
    local full_key = self.key_prefix .. key
    local now = ngx.now()

    local last_time_key = full_key .. ":last"
    local tokens_key = full_key .. ":tokens"

    -- 获取上次时间和令牌数
    local last_time = self.dict:get(last_time_key) or now
    local tokens = tonumber(self.dict:get(tokens_key)) or self.burst

    -- 计算新增令牌
    local time_passed = now - last_time
    local new_tokens = time_passed * self.rate
    tokens = math.min(self.burst, tokens + new_tokens)

    -- 检查是否允许
    if tokens < 1 then
        return false, {
            allowed = false,
            remaining = 0,
            limit = self.rate,
            retry_after = math.ceil((1 - tokens) / self.rate)
        }
    end

    -- 消耗令牌
    tokens = tokens - 1

    -- 更新状态（TTL 60秒）
    self.dict:set(tokens_key, tokens, 60)
    self.dict:set(last_time_key, now, 60)

    return true, {
        allowed = true,
        remaining = math.floor(tokens),
        limit = self.rate
    }
end

return _M
