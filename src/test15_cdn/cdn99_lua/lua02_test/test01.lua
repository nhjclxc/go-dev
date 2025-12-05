#!/usr/bin/env lua

local http = require("socket.http")   -- 使用 luasocket 请求接口
local ltn12 = require("ltn12")
local cjson = require("cjson.safe")

local function fetch_domains()
    local body = {}
    local res, code = http.request{
        url = "http://127.0.0.1:8090/api/v1/domain/getList",
        method = "GET",
        sink = ltn12.sink.table(body),
    }

    if code == 200 then
        local data = table.concat(body)
        local domains = cjson.decode(data)
        print("Fetched domain list:", data)
    else
        print("Failed to fetch domains, status code:", code)
    end
end

while true do
    fetch_domains()
    print("Waiting 60 seconds...")
    os.execute("sleep 60")   -- Linux / Mac
    -- Windows 下可以用: os.execute("timeout 60")
end
