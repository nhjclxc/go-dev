# 1. 使用官方 Go 镜像作为构建阶段
# FROM golang:1.22 AS builder
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.22.3 AS builder

# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct

# 2. 设置工作目录
WORKDIR /app

# 3. 将 go.mod 和 go.sum 拷贝进去，并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 4. 拷贝源代码
COPY . .

# 5. 编译可执行文件（关闭 CGO 更适用于部署）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# 6. 使用更小的镜像进行部署
# FROM alpine:latest
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/alpine:latest

WORKDIR /root/

# 拷贝编译好的文件到最终镜像  # 配置文件的文件夹开关
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config

# 暴露端口（根据项目端口修改）EXPOSE 本身不会真正把端口暴露到宿主机，它只是用来声明容器内部会监听哪些端口。
# 这边将所有可能要给外部提供接入的端口都暴露出来，然后实际使用的时候在用-p命令映射，
# EXPOSE 支持写多行，一行可以暴露多个端口，可以将第一行认为是dev环境要暴露的端口，第二行认为是test要暴露的端口，第三行认为是prod要暴露的端口
#（只有一台服务器的情况下这样操作，如果有多台服务器就不用了，注：一台服务器ip相同端口不同，通过端口区分不同的环境；多台服务器ip不同端口相同，通过ip区分不同的环境）
EXPOSE 18080
EXPOSE 18090
EXPOSE 19090


# 设定启动命令
CMD ["./server"]
