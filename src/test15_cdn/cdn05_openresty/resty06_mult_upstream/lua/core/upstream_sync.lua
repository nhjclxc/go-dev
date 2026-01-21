-- 实现向调度中心 scheduler 每隔5秒动态拉取 openresty 上游的所有后端服务
-- 实现源站的加权轮询和主备切换

local http = require "resty.http"
local cjson = require "cjson.safe"
local upstream_dict = ngx.shared.upstreams -- 导入共享字典
local upstream_dict_key = "upstream"

-- http://127.0.0.1:9090/api/upstreams/{nodename}
local INTERVAL = 5
local base_url = os.getenv("BASE_URL") or "http://scheduler:9090"
local node_name = os.getenv("NODE_NAME") or "node1"

-- 默认 upstream（只在第一次没有数据时设置）
local DEFAULT = {
    { host = "upstream", port = 8080, weight = 50 }
}

-- 调用调度中心接口拉取上游, 这个方法用于给定时任务调度
-- @param: node 当前节点名称
local function fetch_upstreams(node)
    local cli = http.new()
    cli:set_timeout(2000)

    -- http://127.0.0.1:9090/api/upstreams/node1
    local url = base_url  .. "/api/upstreams/" .. node
    local res, err = cli:request_uri(url, {
        method = "GET",
    })

    if not res then
        return false, "scheduler error"
    end

    -- res.body 是字符串
    ngx.log(ngx.INFO, "fetch_upstreams url: ", url, " res.body: ", res.body, res.status, not res.body)

    if res.status ~= 200 or not res.body or res.body == "" then
        return false, "scheduler response body is null"
    end

    -- 解析接口数据, json -> string
    local data = cjson.decode(res.body)
    if not data or type(data) ~= "table" or #data == 0 then
        return false, "scheduler response data is null"
    end

    -- 有数据那么覆盖贡献字典的数据
    upstream_dict:set(upstream_dict_key, res.body)
    ngx.log(ngx.INFO, "fetch_upstreams success, len(data) = ", #data)

    return true
end


-- 这个方法用于不断创建定时器
local function fetch_upstreams_sync(premature)
    if premature then
        return
    end

    local flag, fetch_err = fetch_upstreams(node_name)
    if not flag then
        ngx.log(ngx.ERR, "failed to fetch_upstreams : ", fetch_err, ", use default configration")

        upstream_dict:set(upstream_dict_key, cjson.encode(DEFAULT))
        ngx.log(ngx.INFO, "set default upstream: 127.0.0.1:8080")
    end

    -- 创建下一个定时器
    local ok, err = ngx.timer.at(INTERVAL, fetch_upstreams_sync)
    if not ok then
        ngx.log(ngx.ERR, "failed to create timer: ", err)
    end
end

-- 只让 worker_id = 0 的进程去执行调度
if ngx.worker.id() == 0 then
    -- 一启动的时候把默认的地址给上
    upstream_dict:set(upstream_dict_key, cjson.encode(DEFAULT))

    -- 立即执行一个定时任务
    local ok, err = ngx.timer.at(0, fetch_upstreams_sync)
    if not ok then
        ngx.log(ngx.ERR, "failed to create timer: ", err)
    end
end
