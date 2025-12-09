local uri    = ngx.var.uri
local domain = ngx.var.host

local bytes_sent = 0
local range = ngx.var.http_range

ngx.say("range = ", range)

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
    ngx.say("bytes_sent2 = ", bytes_sent)

else
    -- 非断点续传，用整个响应发送的字节数
    bytes_sent = tonumber(ngx.var.bytes_sent) or 0
end
ngx.say("bytes_sent3 = ", bytes_sent)

-- 用: 域名 + URI 做文件+租户唯一标识
local key = string.format("flow:%s:%s", domain, uri)
ngx.say("key = ", key)

-- 累加流量
local newVal = ngx.shared.stats:incr(key, bytes_sent, 0)
ngx.say("newVal = ", newVal)

ngx.log(ngx.INFO,
    "flow stat => domain=", domain,
    " uri=", uri,
    " bytes=", bytes_sent
)

