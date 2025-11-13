#!/bin/bash
# =====================================================
# 文件名：init_cdn_framework.sh
# 设置运行权限（在本机上）：chmod +x init_cdn_framework.sh
# 功能：创建cdn所需要的所有模块
# 用法： ./init_cdn_framework.sh
# =====================================================

set -e

# 项目根目录（可根据需要修改）
ROOT_DIR="$(pwd)/cdn00_framework"

# CDN 微服务列表
MODULES=(
  "cdn-common"
  "cdn-gslb-service"
  "cdn-slb-service"
  "cdn-origin-service"
  "cdn-ecs-service"
  "cdn-log-service"
  "cdn-monitor-service"
  "cdn-pay-service"
  "cdn-config-service"
  "cdn-cache-manager-service"
  "cdn-auth-service"
  "cdn-traffic-control-service"
  "cdn-backup-service"
  "cdn-alert-service"
  "cdn-api-gateway"
  "cdn-metrics-service"
  "cdn-preload-service"
)

echo "创建 CDN 项目目录：$ROOT_DIR"
mkdir -p "$ROOT_DIR"
cd "$ROOT_DIR"

# 1. 创建模块目录并初始化 go.mod
for module in "${MODULES[@]}"; do
  echo "创建模块目录：$module"
  mkdir -p "$module"
  cd "$module"

  # 初始化 go.mod
  MODULE_NAME="github.com/yourorg/$module"
  if [ ! -f "go.mod" ]; then
    echo "初始化 go.mod: $MODULE_NAME"
    go mod init "$MODULE_NAME"
  fi

  cd ..
done

# 2. 初始化 go.work 并添加所有模块
echo "初始化 go.work 并添加模块"
if [ ! -f "go.work" ]; then
  go work init
fi

for module in "${MODULES[@]}"; do
  go work use "./$module"
done

echo "✅ CDN 项目结构和 go.work 初始化完成！"
