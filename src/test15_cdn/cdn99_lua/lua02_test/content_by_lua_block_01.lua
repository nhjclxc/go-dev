local cache = ngx.shared.domain_cache
local json_str = cache:get("domain_list")

if not json_str then
    ngx.say("domain list not loaded yet")
    return
end

local cjson = require("cjson.safe")
local domains = cjson.decode(json_str)

ngx.say("Current domains: ", json_str)
