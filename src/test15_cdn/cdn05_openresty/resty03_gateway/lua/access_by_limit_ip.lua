local function close_redis(red)  
    if not red then  
        return
    end  
    --释放连接(连接池实现)  
    local pool_max_idle_time = 10000 --毫秒  
    local pool_size = 100 --连接池大小  
    local ok, err = red:set_keepalive(pool_max_idle_time, pool_size)  
    if not ok then  
        ngx.say("set keepalive error : ", err)  
    end  
end

local function errlog(...)
    ngx.log(ngx.ERR, "redis: ", ...)
end

local function duglog(...)
    ngx.log(ngx.DEBUG, "redis: ", ...)
end

local function getIp()
	local myIP = ngx.req.get_headers()["X-Real-IP"]
	if myIP == nil then
		myIP = ngx.req.get_headers()["x_forwarded_for"]
	end
	if myIP == nil then
		myIP = ngx.var.remote_addr
	end
	return myIP;
end

local key = "limit:ip:blacklist"
local ip = getIp();
local shared_ip_blacklist = ngx.shared.shared_ip_blacklist

--获得本地缓存的最新刷新时间
local last_update_time = shared_ip_blacklist:get("last_update_time");

if last_update_time ~= nil then 
	local dif_time = ngx.now() - last_update_time 
	if dif_time < 60 then --缓存1分钟,没有过期
		if shared_ip_blacklist:get(ip) then
			return ngx.exit(ngx.HTTP_FORBIDDEN) --直接返回403
		end
		return
	end
end

local redis = require "resty.redis"  --引入redis模块
local red = redis:new()  --创建一个对象，注意是用冒号调用的

--设置超时（毫秒）  
red:set_timeout(1000) 
--建立连接  
local ip = "127.0.0.1"  
local port = 6379
local ok, err = red:connect(ip, port)
if not ok then  
	close_redis(red)
	errlog("limit ip cannot connect redis");
else
	local ip_blacklist, err = red:smembers(key);
	
	if err then
		errlog("limit ip smembers");
	else
		--刷新本地缓存，重新设置
		shared_ip_blacklist:flush_all();
		
		--同步redis黑名单 到 本地缓存
		for i,bip in ipairs(ip_blacklist) do
			--本地缓存redis中的黑名单
			shared_ip_blacklist:set(bip,true);
		end
		--设置本地缓存的最新更新时间
		shared_ip_blacklist:set("last_update_time",ngx.now());
	end
end  

if shared_ip_blacklist:get(ip) then
	return ngx.exit(ngx.HTTP_FORBIDDEN) --直接返回403
end



-- Lua + Redis 实现动态封禁 IP
-- https://feixiang.blog.csdn.net/article/details/136658128


