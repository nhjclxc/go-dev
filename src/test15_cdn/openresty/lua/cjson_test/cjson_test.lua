#!/usr/bin/env lua

local cjson = require "cjson.safe"

local t = {a = 1, b = 2}
local ok, err = dict:set("my_table", cjson.encode(t))

if not ok then
    ngx.log(ngx.ERR, "failed to set: ", err)
end



local data = dict:get("my_table")
if data then
    local table_val = cjson.decode(data)
    ngx.say(table_val.a, table_val.b)
end
