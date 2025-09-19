# 构建阶段
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gin_kubectl .

# 运行阶段
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/gin_kubectl .
COPY ./config.yaml /app/config.yaml

# 设置时区
RUN apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 给应用添加所有权限
RUN chmod 775 /app/gin_kubectl

EXPOSE 18080 18090 19090

CMD ["/app/gin_kubectl", "-c", "/app/config.yaml"]
