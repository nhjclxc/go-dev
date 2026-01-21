FROM openresty/openresty:alpine

# Install lua-resty-http, lua-resty-ipmatcher, lua-resty-jit-uuid manually
RUN apk add --no-cache wget unzip && \
    mkdir -p /usr/local/openresty/lualib/resty && \
    cd /tmp && \
    # Install lua-resty-http
    wget https://github.com/ledgetech/lua-resty-http/archive/refs/tags/v0.17.1.zip && \
    unzip v0.17.1.zip && \
    cp -r lua-resty-http-0.17.1/lib/resty/* /usr/local/openresty/lualib/resty/ && \
    rm -rf v0.17.1.zip lua-resty-http-0.17.1 && \
    rm -rf /tmp/* && \
    apk del wget unzip


# Copy configuration and Lua code
COPY conf/nginx.conf /usr/local/openresty/nginx/conf/nginx.conf
COPY lua /usr/local/openresty/nginx/lua

# Create logs directory
RUN mkdir -p /usr/local/openresty/nginx/logs

# Expose port 80
EXPOSE 80

# Start OpenResty
CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]
