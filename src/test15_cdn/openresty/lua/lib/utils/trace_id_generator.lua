-- Trace ID 生成器
-- 支持多种格式的 Trace ID 生成

local random = math.random
local format = string.format
local concat = table.concat

local _M = {
    _VERSION = '0.1.0'
}

-- 初始化随机数种子
math.randomseed(ngx.now() * 1000 + ngx.worker.pid())

-- 生成随机十六进制字符串
local function random_hex(length)
    local chars = {}
    for i = 1, length do
        chars[i] = format("%x", random(0, 15))
    end
    return concat(chars)
end

-- 方式1：UUID v4 (推荐用于简单场景)
-- 格式：xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
function _M.generate_uuid()
    local template = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function(c)
        local v = (c == 'x') and random(0, 15) or random(8, 11)
        return format('%x', v)
    end)
end

-- 方式2：时间戳 + Worker PID + 随机数
-- 格式：{timestamp}-{worker_pid}-{random}
-- 优点：可读性好，包含时间和进程信息
function _M.generate_timestamp_based()
    local timestamp = ngx.now() * 1000  -- 毫秒级时间戳
    local worker_pid = ngx.worker.pid()
    local rand = random_hex(8)

    return format("%d-%d-%s",
        math.floor(timestamp),
        worker_pid,
        rand
    )
end

-- 方式3：OpenTelemetry / W3C Trace Context 格式
-- Trace ID: 32位十六进制 (128 bit)
-- Span ID:  16位十六进制 (64 bit)
function _M.generate_otel_trace_id()
    return random_hex(32)
end

function _M.generate_otel_span_id()
    return random_hex(16)
end

-- 生成完整的 W3C traceparent header
-- 格式：00-{trace_id}-{span_id}-{flags}
function _M.generate_w3c_traceparent(trace_id, span_id)
    trace_id = trace_id or _M.generate_otel_trace_id()
    span_id = span_id or _M.generate_otel_span_id()
    local flags = "01"  -- sampled

    return format("00-%s-%s-%s", trace_id, span_id, flags)
end

-- 方式4：Zipkin/Jaeger 格式
-- 64位或128位十六进制
function _M.generate_zipkin_trace_id(use_128bit)
    if use_128bit then
        return random_hex(32)
    else
        return random_hex(16)
    end
end

-- 方式5：紧凑格式（推荐用于高性能场景）
-- 格式：{timestamp_hex}{random_hex}
-- 长度：20-24 字符
function _M.generate_compact()
    local timestamp = ngx.now() * 1000
    local time_hex = format("%x", math.floor(timestamp))
    local rand_hex = random_hex(12)

    return time_hex .. rand_hex
end

-- 解析 W3C traceparent
function _M.parse_w3c_traceparent(traceparent)
    if not traceparent then
        return nil
    end

    -- 格式：version-trace_id-parent_id-flags
    local version, trace_id, parent_id, flags =
        traceparent:match("^(%x%x)%-(%x+)%-(%x+)%-(%x%x)$")

    if not version then
        return nil
    end

    return {
        version = version,
        trace_id = trace_id,
        parent_id = parent_id,
        flags = flags,
        sampled = (tonumber(flags, 16) % 2 == 1)
    }
end

-- 默认生成器（可配置）
_M.default_generator = _M.generate_compact

return _M
