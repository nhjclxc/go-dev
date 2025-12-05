FROM openresty/openresty:alpine

# 安装 luarocks
RUN apk add --no-cache gcc musl-dev make wget curl unzip luarocks

# 安装 lua-resty-http
RUN luarocks install lua-resty-http
RUN luarocks install lua-resty-openssl

