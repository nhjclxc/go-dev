
-- 鉴权

local _M = {}
_M.__index = _M


-- 对 token 进行校验
local function doCheck(token)
    -- 使用 string.find 判断是否包含
    if string.find(token, "aabb", 1, true) then
        ngx.log(ngx.INFO, "包含子串")
        return true
    else
        ngx.log(ngx.INFO, "包含子串")
        return false
    end

    -- 实际对校验过程 ...

end

function _M:check()
    local headers, err = ngx.req.get_headers()
    local token = headers["token"]
    if not token then
       return false
    end
    return doCheck(token)
end

return _M
