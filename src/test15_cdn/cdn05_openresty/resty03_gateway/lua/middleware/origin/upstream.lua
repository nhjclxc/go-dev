-- 回源负载均衡模块（Phase 5 功能）
-- 实现源站的加权轮询和主备切换

local _M = {}


math.randomseed(os.time())

local OriginServer = {}
OriginServer.__index = OriginServer

-- weightedRandom 在所有有权重的服务器中进行加权随机
local function weightedRandom(weightServers)
    if #weightServers == 0 then
        return nil, "no servers len(weightServers) == 0"
    end

    -- 根据 weight 从大到小排序
    table.sort(weightServers, function(a, b)
        return a.Weight > b.Weight
    end)

    local totalWeight = 0
    for _, server in ipairs(weightServers) do
        if server.Weight > 0 then
            totalWeight = totalWeight + server.Weight
        end
    end

    if totalWeight == 0 then
        return nil, "no servers totalWeight == 0"
    end

    -- 随机数rnum区间范围 -> [0, totalWeight)
    local rnum = math.random(totalWeight) - 1

    local current = 0
    for _, server in ipairs(weightServers) do
        if server.Weight > 0 then
            current = current + server.Weight
            if rnum < current then
                return server, nil
            end
        end
    end

    -- 兜底，lua索引从1开始
    return weightServers[#weightServers], nil
end

-- smoothWeightedRoundRobin 平滑加权轮询
-- wrr_state = {}
local function smoothWeightedRoundRobin(weightServers)
    if not weightServers or #weightServers == 0 then
        return nil, "no servers"
    end

    local totalWeight = 0
    local wrr_state = ngx.shared['wrr_state']

    -- 计算 totalWeight，并初始化 currentWeight
    for _, s in ipairs(weightServers) do
        if s.Weight and s.Weight > 0 then
            totalWeight = totalWeight + s.Weight
            local wrr_key = s.Address .. ":" .. s.Port
            if wrr_state:get(wrr_key) == nil then
                wrr_state:set(wrr_key, 0)
            end
        end
    end

    if totalWeight == 0 then
        return nil, "total weight is zero"
    end

    local selected = nil

    -- 核心算法
    for _, s in ipairs(weightServers) do
        if s.Weight and s.Weight > 0 then
            local wrr_key = s.Address .. ":" .. s.Port
            local wrr_selected_key = selected.Address .. ":" .. selected.Port
            wrr_state:set(wrr_key, wrr_state:get(wrr_key) + s.Weight)
            if not selected or wrr_state:get(wrr_key) > wrr_state:get(wrr_selected_key) then
                selected = s
            end
        end
    end

    if not selected then
        return nil, "no server selected"
    end

    -- 平滑
    wrr_state:incr(selected.Address .. ":" .. selected.Port, -totalWeight)

    -- 某个服务器下线时：
    --ngx.shared.wrr_state:delete(s.Address .. ":" .. s.Port)

    return selected, nil
end

-- selectServer：优先主节点（backup=false）
local function select(originServers)
    if #originServers == 0 then
        return nil, "no servers len(originServers) == 0"
    end

    local primaryServers = {}
    local backupServers = {}

    for _, server in ipairs(originServers) do
        if server.Backup then
            table.insert(backupServers, server)
        else
            table.insert(primaryServers, server)
        end
    end

    -- 优先使用主服务器
    if #primaryServers > 0 then
        return smoothWeightedRoundRobin(primaryServers)
    end

    print("primaryServers is null.")

    -- 再使用备用服务器
    if #backupServers > 0 then
        return smoothWeightedRoundRobin(backupServers)
    end

    return nil, "no server found"
end

-- 选择源站服务器
-- @param servers 服务器列表
--   [
--     {
--       address = "192.168.10.1",
--       port = 80,
--       weight = 100,
--       backup = false
--     }
--   ]
-- @return server 选中的服务器，nil 表示无可用服务器
function _M.select_server(servers)
    -- TODO: 实现加权轮询负载均衡算法
    -- 1. 筛选主服务器（backup=false）
    -- 2. 根据 weight 进行加权轮询选择
    -- 3. 如果所有主服务器不可用，切换到备用服务器（backup=true）
    return select(servers)
end

-- 标记服务器为不可用
-- @param server 服务器信息
function _M.mark_server_down(server)
    -- TODO: 实现服务器健康检查和标记逻辑
end

-- 标记服务器为可用
-- @param server 服务器信息
function _M.mark_server_up(server)
    -- TODO: 实现服务器恢复逻辑
end

return _M
