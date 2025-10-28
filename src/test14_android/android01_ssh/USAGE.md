# Android ADB SSH Tunnel Demo

这是一个简单的 demo，通过 SSH 反向隧道将 Android 设备的 ADB 协议暴露到公网。

## 项目结构

- `server/` - SSH 服务器端（运行在公网机器）
- `client/` - SSH 客户端（运行在 Android 设备上，作为二进制文件）

## 工作原理

1. **Server** 在公网监听 SSH 连接（2222 端口）
2. **Client** 运行在 Android 设备上，连接到公网 Server，建立 SSH 反向隧道
3. Client 请求 Server 在公网监听 6666 端口
4. 当有人连接 Server 的 6666 端口时，流量通过 SSH 隧道转发到 Android 设备的本地 ADB（127.0.0.1:5555）

## 使用方法

### 1. 启动服务器（公网机器）

```bash
cd server
go run main.go
```

服务器将监听 2222 端口，等待 Android 客户端连接。

### 2. 编译客户端（为 Android 编译）

由于客户端要运行在 Android 设备上，需要交叉编译：

```bash
cd client

# 编译为 ARM64 架构（大多数现代 Android 设备）
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o androidjlink-client main.go

# 或编译为 ARM 架构（较老的 Android 设备）
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o androidjlink-client main.go
```

### 3. 部署到 Android 设备

通过 ADB 将编译好的二进制文件推送到 Android 设备：

```bash
adb push androidjlink-client /data/local/tmp/
adb shell chmod +x /data/local/tmp/androidjlink-client
```

### 4. 在 Android 设备上运行客户端

```bash
# 先修改 client/main.go 中的 serverAddr 为你的公网服务器 IP
# 然后重新编译后推送到设备

adb shell
cd /data/local/tmp
./androidjlink-client
```

客户端会：
1. 连接到公网 SSH 服务器
2. 建立反向隧道，让服务器监听 6666 端口
3. 将连接转发到本地 Android 的 ADB（127.0.0.1:5555）

### 5. 远程连接 ADB

在任何机器上，通过公网服务器 IP 连接：

```bash
adb connect YOUR_SERVER_IP:6666
adb devices
adb shell
```

## 配置说明

### 服务器端配置

在 `server/main.go` 中修改：
- SSH 监听端口：默认 `0.0.0.0:2222`

### 客户端配置

在 `client/main.go` 中修改：
- SSH 服务器地址：`serverAddr` 变量，默认 `YOUR_SERVER_IP:2222`（**必须修改**）
- 反向隧道监听地址：默认 `0.0.0.0:6666`
- 本地 ADB 地址：默认 `127.0.0.1:5555`

## 前置条件

1. Android 设备需要开启 ADB over TCP：
```bash
adb tcpip 5555
```

2. 公网服务器需要开放端口：
   - 2222：SSH 连接端口
   - 6666：ADB 访问端口

## 注意事项

这只是一个可行性验证的 demo，**未考虑安全性和性能**：

- 使用了 NoClientAuth，没有客户端认证
- 使用了硬编码的 SSH host key
- 忽略了 SSH host key 验证
- 没有错误重连机制
- 没有日志管理
- 端口直接暴露在公网，无访问控制

在生产环境使用时需要添加：
- 正确的 SSH 密钥认证
- TLS/SSL 加密
- 连接重试和心跳机制
- 访问控制和日志审计
- 端口白名单限制
