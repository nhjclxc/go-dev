
-- 速率限制/防刷

local _M = {}
_M.__index = _M

function _M:allow()
    local t03_throttle = ngx.shared.t03_throttle
    local addr = ngx.var.remote_addr
    local newval, err, forcible = t03_throttle:incr(addr, 1, 0, 10)

    ngx.log(ngx.INFO, "t03_throttle newval = ", newval)
    newval = newval or 1

    ngx.log(ngx.INFO, "t03_throttle newval2 = ", newval)
    return newval <= 3
end

return _M