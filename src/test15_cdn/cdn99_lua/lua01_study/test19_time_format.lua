#!/usr/bin/env lua


-- lua时间对象和字符串时间互转
-- "2025-11-04T15:59:35+08:00"

local timeStr = "2025-11-04T15:59:35+08:00"
print("timeStr = "..timeStr)


local function parse_iso8601(str)
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

    return final_ts, t
end

local ts, tt = parse_iso8601(timeStr)
local dataTime = os.date("%Y-%m-%d %H:%M:%S", ts)
print("ts = "..ts)
print("local time = ", dataTime)
print("local time = ", dataTime.year)
print("local time = ", dataTime.month)
print(tt.year)
print(tt.month)
print(tt.day)
print(tt.hour)
print(tt.min)
print(tt.sec)
print(tt.isdst)
