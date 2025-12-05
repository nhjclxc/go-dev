-- 调接口获取域名数据示例

local http = require "resty.http"
local cjson = require "cjson.safe"
local date = require "resty.date"
local time = require "resty.core.time"

-- my_dict 要在nginx.conf里面定义，否则就是nil
local dict = ngx.shared.my_dict

ngx.log(ngx.INFO, "enter fetch_domain")

local function fetch()
    local httpc = http.new()
    -- 注意接口域名（ip）
    local res, err = httpc:request_uri("http://192.168.200.222:8899/api/v1/domain/getList", {
        method = "GET",
        ssl_verify = false,
    })
    if not res then
        ngx.log(ngx.ERR, "fetch domains failed: ", err)
        return
    end
    ngx.log(ngx.INFO, res.body)
    local bodyTab = cjson.decode(res.body)
    ngx.log(ngx.INFO, "bodyTab['code'] = ", bodyTab["code"])
    ngx.log(ngx.INFO, "bodyTab['msg'] = ", bodyTab["msg"])
    ngx.log(ngx.INFO, "bodyTab['data'] = ", cjson.encode(bodyTab["data"]))
    ngx.log(ngx.INFO, "bodyTab['data'][1] = ", cjson.encode(bodyTab["data"][1]))
    ngx.log(ngx.INFO, "bodyTab['data'][1]['domain_name'] = ", cjson.encode(bodyTab["data"][1]["domain_name"]))
    ngx.log(ngx.INFO, "bodyTab['data'][1]['created_at'] = ", cjson.encode(bodyTab["data"][1]["created_at"]))

    local created_at_str = cjson.encode(bodyTab["data"][1]["created_at"])

    ngx.log(ngx.INFO, str_to_time("2025-11-04 15:59:35"))

    ngx.log(ngx.INFO, "created_at_str = ", created_at_str)
    ngx.log(ngx.INFO, parse_iso8601(created_at_str))


    dict:set("domain_list", res.body)

    ngx.log(ngx.INFO, "code = ", dict:get("domain_list"))


end

function str_to_time(str)
    local y, m, d, H, M, S =
        str:match("(%d+)%-(%d+)%-(%d+)%s+(%d+):(%d+):(%d+)")
    return os.time({
        year = y,
        month = m,
        day = d,
        hour = H,
        min = M,
        sec = S
    })
end

function parse_iso8601(str)
    local year, month, day, hour, min, sec, tz_sign, tz_hour, tz_min =
        str:match("^(%d+)%-(%d+)%-(%d+)T(%d+):(%d+):(%d+)([%+%-])(%d%d):(%d%d)$")

    if not year then
        return nil, "invalid ISO8601: " .. tostring(str)
    end

    -- 当前时间（按本地时区 +08:00）
    local t = {
        year = tonumber(year),
        month = tonumber(month),
        day = tonumber(day),
        hour = tonumber(hour),
        min = tonumber(min),
        sec = tonumber(sec),
        isdst = false,
    }

    -- 转换为时间戳（本地时区）
    local local_ts = os.time(t)

    -- 处理时区偏移
    local offset = tonumber(tz_hour) * 3600 + tonumber(tz_min) * 60
    if tz_sign == "-" then
        offset = -offset
    end

    --- 当前系统的本地时区（通常是 +8）
    local local_offset = os.difftime(os.time(), os.time(os.date("!*t")))

    -- 目标 UTC 时间戳 = 本地时间戳 - 本地时区 + 字符串的时区
    local final_ts = local_ts - local_offset + offset

    return final_ts
end



return fetch
