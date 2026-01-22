-- 实现向调度中心 scheduler 每隔5秒动态拉取 openresty 上游的所有后端服务
-- 实现源站的加权轮询和主备切换

local cjson = require "cjson.safe"
local balancer = require "ngx.balancer"
local upstream_dict = ngx.shared.upstreams -- 导入共享字典
local upstream_dict_key = "upstream"

local ups = upstream_dict:get(upstream_dict_key)
if not ups then
    ngx.log(ngx.ERR, "no upstream addr ")
    ngx.exit(503)
end

ngx.log(ngx.INFO, "ups: ", ups)
local data_table = cjson.decode(ups)
if not data_table or #data_table == 0 then
    ngx.log(ngx.ERR, "no upstream addr 0")
    ngx.exit(503)
end


-- 负载均衡

local total_weight = 0
local selected  -- 被选中的后端
local selected_key

for _, node in ipairs(data_table) do
    local w = tonumber(node.weight) or 1
    total_weight = total_weight + w

    local key = "cw:upstream:" .. node.host .. ":" .. node.port
    local cw = upstream_dict:get(key) or 0

    cw = cw + w
    upstream_dict:set(key, cw)

    if not selected or cw > selected.cw then
        selected = {
            host = node.host,
            ip = node.ip,
            port = node.port,
            cw = cw,
            weight = w,
        }
        selected_key = key
    end
end

-- 减去 total_weight
upstream_dict:incr(selected_key, -total_weight)

ngx.log(ngx.INFO, "selected upstream: ", "host:", selected.host, "ip:", selected.ip, "port:", selected.port)

-- 设置被选中的后端
local ok, err = balancer.set_current_peer(selected.ip, selected.port)
if not ok then
    ngx.log(ngx.ERR, "set peer failed: ", err)
    return ngx.exit(502)
end


