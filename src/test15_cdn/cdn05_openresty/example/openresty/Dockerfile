# -------------------------
# Build stage
# -------------------------
FROM openresty/openresty:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装构建依赖
RUN apk add --no-cache wget unzip git build-base luajit-dev cmake pcre-dev zlib-dev openssl-dev brotli-dev

WORKDIR /tmp

# 创建目标目录
RUN mkdir -p /usr/local/openresty/lualib/resty

# -------------------------
# 1. lua-resty-http
# -------------------------
RUN wget https://github.com/ledgetech/lua-resty-http/archive/refs/tags/v0.17.1.zip && \
    unzip v0.17.1.zip && \
    cp -r lua-resty-http-0.17.1/lib/resty/* /usr/local/openresty/lualib/resty/ && \
    rm -rf v0.17.1.zip lua-resty-http-0.17.1

# -------------------------
# 2. lua-resty-ipmatcher
# -------------------------
RUN wget https://github.com/api7/lua-resty-ipmatcher/archive/refs/tags/v0.6.1.zip && \
    unzip v0.6.1.zip && \
    cp lua-resty-ipmatcher-0.6.1/resty/ipmatcher.lua /usr/local/openresty/lualib/resty/ && \
    rm -rf v0.6.1.zip lua-resty-ipmatcher-0.6.1

# -------------------------
# 3. lua-resty-jit-uuid
# -------------------------
RUN wget https://github.com/thibaultcha/lua-resty-jit-uuid/archive/refs/tags/0.0.7.zip && \
    unzip 0.0.7.zip && \
    cp lua-resty-jit-uuid-0.0.7/lib/resty/jit-uuid.lua /usr/local/openresty/lualib/resty/ && \
    rm -rf 0.0.7.zip lua-resty-jit-uuid-0.0.7

# -------------------------
# 4. lua-resty-balancer (resty.chash)
# -------------------------
RUN wget https://github.com/openresty/lua-resty-balancer/archive/refs/tags/v0.04.zip && \
    unzip v0.04.zip && \
    cp -r lua-resty-balancer-0.04/lib/resty/* /usr/local/openresty/lualib/resty/ && \
    rm -rf v0.04.zip lua-resty-balancer-0.04

# -------------------------
# 5. luafilesystem
# -------------------------
RUN wget https://github.com/lunarmodules/luafilesystem/archive/refs/tags/v1_9_0.zip && \
    unzip v1_9_0.zip && \
    cd luafilesystem-1_9_0/src && \
    gcc -O2 -fPIC -shared \
        -I/usr/local/openresty/luajit/include/luajit-2.1 \
        -o lfs.so lfs.c && \
    mv lfs.so /usr/local/openresty/lualib/ && \
    rm -rf v1_9_0.zip luafilesystem-1_9_0

# -------------------------
# 6. ngx_brotli (dynamic modules)
# -------------------------
# FULL_NGINX_VERSION=1.27.1.2，nginx.org下没有1.27.x.y，只有1.27.x
RUN set -eux; \
    FULL_NGINX_VERSION="$(openresty -V 2>&1 | awk -F/ '/nginx version/ {print $2}')"; \
    NGINX_VERSION="${FULL_NGINX_VERSION%.*}"; \
    mkdir -p /usr/local/openresty/nginx/modules && \
    git clone --recursive --depth 1 --shallow-submodules https://github.com/google/ngx_brotli.git /tmp/ngx_brotli && \
    wget -q -O /tmp/nginx.tar.gz "https://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz" && \
    tar -C /tmp -xzf /tmp/nginx.tar.gz && \
    cd "/tmp/nginx-${NGINX_VERSION}" && \
    ./configure --with-compat --add-dynamic-module=/tmp/ngx_brotli && \
    make modules && \
    cp objs/ngx_http_brotli_filter_module.so /usr/local/openresty/nginx/modules/ && \
    cp objs/ngx_http_brotli_static_module.so /usr/local/openresty/nginx/modules/ && \
    rm -rf /tmp/nginx.tar.gz "/tmp/nginx-${NGINX_VERSION}" /tmp/ngx_brotli

# -------------------------
# Runtime stage
# -------------------------
FROM openresty/openresty:alpine

# 拷贝编译好的 Lua 模块
COPY --from=builder /usr/local/openresty/lualib /usr/local/openresty/lualib
# 拷贝 brotli 动态模块
COPY --from=builder /usr/local/openresty/nginx/modules /usr/local/openresty/nginx/modules

# 保证日志与证书目录存在（挂载点必须提前创建）
RUN mkdir -p /usr/local/openresty/nginx/logs \
    /usr/local/openresty/nginx/conf/certs
