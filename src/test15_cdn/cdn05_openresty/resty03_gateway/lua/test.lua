#!/usr/bin/env lua


-- 测试nginx.conf里面是不是可以导入lua脚本

local test = {}

test.data = "lua脚本数据"

function test.get_str()
    return "来自lua脚本的字符串"
end

ngx.log(ngx.INFO, "在luqwqwa脚本12345里8899面输出的log")


-- 刷新服务端的域名列表
local function domain_refresh()
    local http = require "resty.http"
    local hc = http.new()

    local res, err = h:request_uri("http://backend/domains", { method = "GET" })
    if not res then
        ngx.log(ngx.ERR, "fetch error: ", err)
        return
    end

    local dict = ngx.shared.domains
    dict:set("list", res.body)
end

return test
