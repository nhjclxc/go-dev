-- 默认配置
return {
    -- 应用配置
    app = {
        name = "CDN Gateway",
        version = "1.0.0",
        env = os.getenv("APP_ENV") or "production"
    },

    -- 限流配置
    rate_limit = {
        enabled = true,
        dict_name = "rate_limit",
        rate = 100,         -- 100 req/s
        burst = 50
    },

    -- 后端配置
    backend = {
        host = os.getenv("BACKEND_HOST") or "your-backend-host",
        port = tonumber(os.getenv("BACKEND_PORT")) or 80,
        timeout = 5000,     -- 5秒
        keepalive = 32
    },

    -- 链路追踪配置
    tracing = {
        enabled = true,
        generator_type = "compact",  -- compact, uuid, otel
        return_in_response = true,
        header_name = "X-Trace-ID"
    },

    -- 日志配置
    logging = {
        level = "warn",     -- debug, info, warn, error
        json_format = true
    }
}
