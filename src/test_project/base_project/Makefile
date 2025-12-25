# base_project Makefile
# 支持编译时注入版本信息

# 版本信息
VERSION ?= 1.0.0
BUILD_TIME := $(shell date -u '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 编译参数
LDFLAGS := -X 'base_project/pkg/version.Version=$(VERSION)' \
           -X 'base_project/pkg/version.BuildTime=$(BUILD_TIME)' \
           -X 'base_project/pkg/version.GitCommit=$(GIT_COMMIT)'

# 输出目录
BUILD_DIR := build

# 目标平台
.PHONY: all clean linux-arm64 linux-amd64 darwin-amd64 darwin-arm64 version

all: linux-arm64 linux-amd64

# 显示版本信息
version:
	@echo "Version: $(VERSION)"
	@echo "BuildTime: $(BUILD_TIME)"
	@echo "GitCommit: $(GIT_COMMIT)"

# 清理构建产物
clean:
	rm -rf $(BUILD_DIR)

# Linux ARM64（Android 设备）
linux-arm64:
	@echo "Building for linux/arm64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/base_project-$(VERSION)-linux-arm64 .
	@echo "Build complete: $(BUILD_DIR)/base_project-$(VERSION)-linux-arm64"

# Linux AMD64（服务器）
linux-amd64:
	@echo "Building for linux/amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/base_project-$(VERSION)-linux-amd64 .
	@echo "Build complete: $(BUILD_DIR)/base_project-$(VERSION)-linux-amd64"

# macOS AMD64
darwin-amd64:
	@echo "Building for darwin/amd64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/base_project-$(VERSION)-darwin-amd64 .
	@echo "Build complete: $(BUILD_DIR)/base_project-$(VERSION)-darwin-amd64"

# macOS ARM64（M1/M2）
darwin-arm64:
	@echo "Building for darwin/arm64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/base_project-$(VERSION)-darwin-arm64 .
	@echo "Build complete: $(BUILD_DIR)/base_project-$(VERSION)-darwin-arm64"

# 生成 SHA256 校验和
checksum:
	@echo "Generating checksums..."
	@cd $(BUILD_DIR) && for file in base_project-*; do \
		if [ -f "$$file" ]; then \
			sha256sum "$$file" > "$$file.sha256"; \
			echo "Generated checksum for $$file"; \
		fi \
	done

# 构建所有平台并生成校验和
release: all checksum
	@echo "Release build complete!"
	@echo "Files in $(BUILD_DIR):"
	@ls -lh $(BUILD_DIR)/

# 快速构建（用于开发测试）
dev:
	@echo "Building for current platform (dev mode)..."
	go build -ldflags "$(LDFLAGS)" -o base_project .
	@echo "Dev build complete: ./base_project"

# 测试
test:
	go test -v ./...

# 运行 go mod tidy
tidy:
	go mod tidy
