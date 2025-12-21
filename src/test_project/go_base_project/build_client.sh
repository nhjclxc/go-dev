#!/bin/bash

# 获取版本信息
VERSION=$(git describe --tags --always 2>/dev/null || echo "v0.0.0")
BUILD_TIME=$(date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GO_VERSION=$(go version | awk '{print $3}')

# 构建 ldflags
LDFLAGS="-w -s \
  -X 'go_base_project/version.Version=${VERSION}' \
  -X 'go_base_project/version.BuildTime=${BUILD_TIME}' \
  -X 'go_base_project/version.GitCommit=${GIT_COMMIT}' \
  -X 'go_base_project/version.GoVersion=${GO_VERSION}'"

echo "Building go_base_project Client..."
echo "Version: ${VERSION}"
echo "Build Time: ${BUILD_TIME}"
echo "Git Commit: ${GIT_COMMIT}"
echo "Go Version: ${GO_VERSION}"
echo ""

# 编译
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
  -tags client \
  -ldflags "${LDFLAGS}" \
  -o bin/go_base_project \
  main.go

if [ $? -eq 0 ]; then
  echo "Build successful! Binary: bin/go_base_project"
  echo "File size: $(ls -lh bin/go_base_project | awk '{print $5}')"
else
  echo "Build failed!"
  exit 1
fi
