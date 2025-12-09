
-- 限流

local _M = {}

function _M:allow()
    local t03_dict = ngx.shared.t03_dict
    local key = ngx.var.remote_addr
    local newval, err = t03_dict:incr(key, 1, 0, 20)
    ngx.log(ngx.INFO, "限流 newval = ", newval)
    newval = newval or 99999

    return newval <= 5
end

return _M