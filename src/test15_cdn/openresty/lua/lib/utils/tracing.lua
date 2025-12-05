-- 链路追踪模块
-- 只使用 X-Trace-Id

local trace_id_generator = require "lib.utils.trace_id_generator"

local _M = {}

local HEADER_NAME = "X-Trace-Id"

-- 生成 Trace ID
local function generate_trace_id()
    return trace_id_generator.generate_compact()
end

-- 初始化追踪
function _M.init()
    -- 1. 尝试从请求头获取
    local trace_id = ngx.req.get_headers()[HEADER_NAME]

    -- 2. 如果没有，生成新的
    if not trace_id or trace_id == "" then
        trace_id = generate_trace_id()
    end

    -- 3. 存储到上下文
    ngx.ctx.trace_id = trace_id

    -- 4. 设置 nginx 变量（用于日志）
    ngx.var.trace_id = trace_id

    -- 5. 设置传递给后端的请求头
    ngx.req.set_header(HEADER_NAME, trace_id)

    return trace_id
end

-- 设置响应头
function _M.set_response_header()
    local trace_id = ngx.ctx.trace_id
    if trace_id then
        ngx.header[HEADER_NAME] = trace_id
    end
end

-- 获取当前 Trace ID
function _M.get_trace_id()
    return ngx.ctx.trace_id
end

return _M
