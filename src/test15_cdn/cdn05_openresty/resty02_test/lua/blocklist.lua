
-- 黑名单

local _M = {}
_M.__index = _M

function _M:is_blocklist()
    local t03_blocklist = ngx.shared.t03_blocklist
    local addr = ngx.var.remote_addr
    local value, flags = t03_blocklist:get(addr)
    ngx.log(ngx.INFO, "t03_blocklist:get(addr) = ", value, ", value ~= nil => ", value ~= nil)

    return value ~= nil
end

return _M