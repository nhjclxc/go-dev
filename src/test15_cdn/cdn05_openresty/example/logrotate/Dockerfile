# Dockerfile
FROM alpine:3.18
RUN apk add --no-cache logrotate
COPY openresty_logrotate /etc/logrotate.d/openresty_logrotate
VOLUME /var/log/openresty_logrotate
CMD ["sh", "-c", "while true; do logrotate -f /etc/logrotate.d/openresty_logrotate; sleep 60; done"]