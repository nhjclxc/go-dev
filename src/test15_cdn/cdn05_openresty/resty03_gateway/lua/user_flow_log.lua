local cjson = require "cjson.safe"
local logger = require "logger"

-- get_bytes_sent 获取用户本次请求的数据大小
local function get_bytes_sent(host, uri)
    local bytes_sent = 0
    local range = ngx.var.http_range

    if range then
        -- Range: bytes=1000-1999 或 bytes=1000-
        local start_str, end_str = string.match(range, "bytes=(%d*)-(%d*)")

        local start = tonumber(start_str)
        local finish

        if end_str == "" then
            -- bytes=1000-
            finish = tonumber(ngx.var.content_length) - 1
        else
            finish = tonumber(end_str)
        end

        bytes_sent = finish - start + 1
        ngx.log(ngx.INFO, "bytes_sent2 = ", bytes_sent)

    else
        -- 非断点续传，用整个响应发送的字节数
        bytes_sent = tonumber(ngx.var.bytes_sent) or 0
    end

    -- 用: 域名 + URI 做文件+租户唯一标识
    local key = string.format("flow:%s:%s", host, uri)
    ngx.log(ngx.INFO)

    -- 累加流量
    local newVal = ngx.shared.stats:incr(key, bytes_sent, 0)
    ngx.log(ngx.INFO, "key = ", key, "bytes_sent = ", bytes_sent, "newVal = ", newVal)

    return bytes_sent
end

-- 获取本次的流量
local bytes_sent = get_bytes_sent(host, uri)

-- 构建 Lua 表
local log_data = {
    trace_id = trace_id,
    timestamp = ngx.now(),
    date = os.date("%Y-%m-%d"),
    hour = os.date("%H"),
    remote_addr = ngx.var.remote_addr,
    client_ip = ngx.var.http_x_forwarded_for or ngx.var.remote_addr,
    user_id = ngx.var.user_id,
    tenant_id = ngx.var.tenant_id,
    session_id = ngx.var.session_id,
    host = ngx.var.host,
    uri = ngx.var.uri,
    method = ngx.var.request_method,
    protocol = ngx.var.server_protocol,
    status = tonumber(ngx.var.status),
    body_bytes_sent = tonumber(ngx.var.body_bytes_sent),
    bytes_sent = bytes_sent,
    request_time = tonumber(ngx.var.request_time),
    upstream_response_time = ngx.var.upstream_response_time,
    cache_status = ngx.var.upstream_cache_status,
    range = ngx.var.http_range,
    content_range = ngx.var.content_range,
    -- percent_sent = bytes_sent / file_size,
    -- file_size = file_size,
    file_type = ngx.var.content_type,
    http_referer = ngx.var.http_referer,
    http_user_agent = ngx.var.http_user_agent,
    connection_aborted = ngx.var.connection_aborted
}


local str = cjson.encode(log_data)

ngx.log(ngx.INFO, "jsonstr = ", str)
logger:flow_log(str)

-- OpenResty(Lua) → JSON 日志 → Fluent Bit/Kafka → Go 后端 → Redis/DB → 报表/计费
-- Fluent Bit去采集输出的日志到消息中间件里面，然后由go程序去消息中间件里面消费消息
-- 将日志记录redis，然后由后端go程序定时30min/1h去redis里面拉去数据
-- // Redis Key 按日期 + 小时 + 用户 + 文件聚合
-- t := time.Unix(int64(logEntry.Timestamp), 0)
-- date := t.Format("20060102")
-- hour := t.Format("15")
-- redisKey := fmt.Sprintf("flow:%s:%s:%s:%s", date, hour, logEntry.UserID, logEntry.URI)
-- 可以使用*来获取不同要求的流量数据：flow:20251209:14:alice:/video.mp4 -> 1234567
-- 20251209这天alice用户的所有流量：flow:20251209:*:alice:*
--[[
iter := rdb.Scan(ctx, 0, "flow:20251209:*:alice:*", 1000).Iterator()
for iter.Next(ctx) {
    k := iter.Val()
    fmt.Println("Matched key:", k)
}
if err := iter.Err(); err != nil {
    log.Println(err)
}
]]
-- Redis 内部可以用 hash 或 zset 代替平铺 Key，以优化key中*的扫描，如：rdb.HGet(ctx, "flow:20251209:14:user123", "/video.mp4")



