-- ssl_handler.lua
-- HTTPS/TLS 配置处理（证书动态加载、TLS版本控制）

local _M = {}

local ssl = require "ngx.ssl"
local lrucache = require "resty.lrucache"
local settings = require "app.config.settings"
local config_helper = require "lib.config_helper"

-- 证书缓存 (每个 worker 独立)
local cert_cache, err = lrucache.new(1000)
if not cert_cache then
    ngx.log(ngx.ERR, "failed to create cert cache: ", err)
end

-- 获取共享字典用于版本控制
local function get_version_dict()
    return ngx.shared[settings.DICT_SYNC_STATE]
end

-- 获取域名配置版本号
-- 用于缓存失效机制
local function get_domain_version(domain)
    local dict = get_version_dict()
    if not dict then
        return 0
    end

    local version = dict:get("ssl_version:" .. domain) or 0
    return version
end

-- 使缓存失效（当域名配置更新时调用）
function _M.invalidate_cache(domain)
    if not domain then
        return
    end

    local dict = get_version_dict()
    if not dict then
        ngx.log(ngx.WARN, "[ssl_handler] version dict not available, cannot invalidate cache for: ", domain)
        return
    end

    -- 增加版本号，使旧缓存失效
    local new_version, err = dict:incr("ssl_version:" .. domain, 1, 0)
    if not new_version then
        ngx.log(ngx.ERR, "[ssl_handler] failed to increment version for ", domain, ": ", err)
        return
    end

    ngx.log(ngx.INFO, "[ssl_handler] invalidated SSL cache for ", domain, ", new version: ", new_version)
end

-- 生成缓存 key，包含版本号
local function get_cache_key(domain)
    local version = get_domain_version(domain)
    return domain .. ":v" .. version
end

-- TLS 版本映射表：配置格式 -> OpenSSL 格式
-- 配置格式: "1.0", "1.1", "1.2", "1.3"
-- OpenSSL 格式: "TLSv1", "TLSv1.1", "TLSv1.2", "TLSv1.3"
local TLS_VERSION_MAP = {
    ["1.0"] = "TLSv1",
    ["1.1"] = "TLSv1.1",
    ["1.2"] = "TLSv1.2",
    ["1.3"] = "TLSv1.3",
}

-- 转换 TLS 版本列表格式
local function convert_tls_versions(tls_list)
    if not tls_list or #tls_list == 0 then
        return nil
    end

    local converted = {}
    for i, version in ipairs(tls_list) do
        local mapped = TLS_VERSION_MAP[version]
        if mapped then
            converted[#converted + 1] = mapped
        else
            ngx.log(ngx.WARN, "[ssl_handler] unknown TLS version: ", version)
        end
    end

    if #converted == 0 then
        return nil
    end

    return converted
end

-- 设置 TLS 协议版本
local function set_tls_protocols(tls_list)
    if not tls_list or #tls_list == 0 then
        return true
    end

    -- 检查 set_protocols 是否可用 (OpenResty 1.19.3.1+)
    if not ssl.set_protocols then
        return true
    end

    -- 转换版本格式
    local protocols = convert_tls_versions(tls_list)
    if not protocols then
        return true
    end

    local ok, err = ssl.set_protocols(protocols)
    if not ok then
        ngx.log(ngx.ERR, "[ssl_handler] failed to set TLS protocols: ", err)
        return false
    end
    return true
end

-- 处理 SSL 证书动态加载
function _M.handle_ssl()
    local server_name = ssl.server_name()
    if not server_name then
        return  -- 无 SNI，使用默认证书
    end

    -- 1. 检查缓存（使用包含版本号的 key）
    if cert_cache then
        local cache_key = get_cache_key(server_name)
        local cached = cert_cache:get(cache_key)
        if cached then
            -- 设置 TLS 版本
            set_tls_protocols(cached.tls_list)

            local ok, err = ssl.clear_certs()
            if not ok then
                ngx.log(ngx.ERR, "failed to clear certs: ", err)
                return
            end

            ok, err = ssl.set_cert(cached.cert)
            if not ok then
                ngx.log(ngx.ERR, "failed to set cert: ", err)
                return
            end

            ok, err = ssl.set_priv_key(cached.key)
            if not ok then
                ngx.log(ngx.ERR, "failed to set key: ", err)
                return
            end
            return
        end
    end

    -- 2. 获取域名配置
    local domain_matcher = require "core.domain_matcher"
    local config = domain_matcher.get_config(server_name)
    if not config then
        return  -- 无配置，使用默认证书
    end

    local https_config = config.httpsConfig
    if not https_config or not config_helper.is_enabled(https_config.httpsOn) then
        return  -- HTTPS 未开启
    end

    local cert_info = https_config.certInfo
    if not cert_info or not cert_info.certificate or not cert_info.privateKey then
        return  -- 无证书配置
    end

    -- 3. 设置 TLS 版本
    local tls_list = https_config.tlsList
    set_tls_protocols(tls_list)

    -- 4. 解析证书
    local cert_pem = cert_info.certificate
    local key_pem = cert_info.privateKey

    local cert, err = ssl.parse_pem_cert(cert_pem)
    if not cert then
        ngx.log(ngx.ERR, "failed to parse cert for ", server_name, ": ", err)
        return
    end

    local key, err = ssl.parse_pem_priv_key(key_pem)
    if not key then
        ngx.log(ngx.ERR, "failed to parse key for ", server_name, ": ", err)
        return
    end

    -- 5. 缓存证书和 TLS 配置 (1 小时，但会通过版本号自动失效)
    if cert_cache then
        local cache_key = get_cache_key(server_name)
        cert_cache:set(cache_key, { cert = cert, key = key, tls_list = tls_list }, 3600)
    end

    -- 6. 设置证书
    local ok, err = ssl.clear_certs()
    if not ok then
        ngx.log(ngx.ERR, "failed to clear certs: ", err)
        return
    end

    ok, err = ssl.set_cert(cert)
    if not ok then
        ngx.log(ngx.ERR, "failed to set cert: ", err)
        return
    end

    ok, err = ssl.set_priv_key(key)
    if not ok then
        ngx.log(ngx.ERR, "failed to set key: ", err)
        return
    end

    ngx.log(ngx.INFO, "SSL cert loaded for: ", server_name)
end

return _M
