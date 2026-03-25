# multi-stage Dockerfile for a Go app
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git curl

# 最后加一条保持容器运行的命令
CMD ["tail", "-f", "/dev/null"]