local http = require("resty.http")
local cjson = require("cjson.safe")

-- 防止多个 worker 同时执行
local lock = require("resty.lock")
local domain_lock = lock:new("domain_cache", { timeout = 0, exptime = 5 })

local function fetch_domain_list()
    -- 尝试上锁
    local elapsed, err = domain_lock:lock("domain_update_lock")
    if not elapsed then
        -- 其他 worker 正在更新，当前 worker 不执行
        return
    end

    local httpc = http.new()
    local res, err = httpc:request_uri("http://127.0.0.1:8090/api/v1/domain/getList", {
        method = "GET",
        ssl_verify = false,
        timeout = 5000,
    })

    if not res then
        print("Failed to fetch domain list: ", err)
        domain_lock:unlock()
        return
    end

    -- 解析 JSON
    local data = cjson.decode(res.body)
    if not data or not data.data then
        -- ngx.log(ngx.ERR, "Invalid domain list response")
        print("Invalid domain list response")
        domain_lock:unlock()
        return
    end

    -- 存入共享内存
    -- cache:set("domai_list", res.body) -- 直接保存原始 JSON 字符串
    print("domain_list", res.body)

    print("Domain list updated: ")

    domain_lock:unlock()
end


fetch_domain_list()
