#!/usr/bin/env lua


-- 测试nginx.conf里面是不是可以导入lua脚本

local test = {}

test.data = "lua脚本数据"

function test.get_str()
    return "来自lua脚本的字符串"
end

ngx.log(ngx.INFO, "在lua脚本12345里8899面输出的log")

return test
