-- 日志工具模块
local cjson = require "cjson.safe"

local _M = {
    _VERSION = '0.1.0'
}

-- 日志级别
_M.LEVELS = {
    DEBUG = ngx.DEBUG,
    INFO = ngx.INFO,
    WARN = ngx.WARN,
    ERROR = ngx.ERR,
    CRIT = ngx.CRIT
}

-- 创建日志实例
function _M.new(opts)
    opts = opts or {}

    local self = {
        prefix = opts.prefix or "",
        level = opts.level or ngx.INFO,
        json_format = opts.json_format or false
    }

    return setmetatable(self, {__index = _M})
end

-- 格式化日志消息
local function format_message(prefix, ...)
    local args = {...}
    local parts = {}

    if prefix and prefix ~= "" then
        table.insert(parts, "[" .. prefix .. "]")
    end

    for _, arg in ipairs(args) do
        if type(arg) == "table" then
            table.insert(parts, cjson.encode(arg))
        else
            table.insert(parts, tostring(arg))
        end
    end

    return table.concat(parts, " ")
end

-- 基础日志方法
local function log(self, level, ...)
    if level < self.level then
        return
    end

    local message = format_message(self.prefix, ...)
    ngx.log(level, message)
end

-- 便捷方法
function _M:debug(...)
    log(self, ngx.DEBUG, ...)
end

function _M:info(...)
    log(self, ngx.INFO, ...)
end

function _M:warn(...)
    log(self, ngx.WARN, ...)
end

function _M:error(...)
    log(self, ngx.ERR, ...)
end

function _M:crit(...)
    log(self, ngx.CRIT, ...)
end

-- 结构化日志
function _M:log_json(level, data)
    if level < self.level then
        return
    end

    data.timestamp = ngx.now()
    data.prefix = self.prefix

    local json_str = cjson.encode(data)
    if json_str then
        ngx.log(level, json_str)
    end
end

-- append_to_file 将日志记录到指定文件
local function append_to_file(premature, filename, str)
    if premature then return end

    local file, err = io.open(filename, "a")  -- 追加记录
    if not file then
        ngx.log(ngx.ERR, "failed to open log file: ", err)
        return
    end
    file:write(str, "\n")
    file:close()
end

function _M:flow_log(data)
    -- 异步记录日志
    ngx.timer.at(0, append_to_file, "/usr/local/openresty/nginx/logs/user_flow_log.log", data)
    ngx.log(ngx.INFO, "日志记录完成")
end


-- 默认实例
_M.default = _M.new()

ngx.log(ngx.INFO, "日志模块加载完成")

return _M
