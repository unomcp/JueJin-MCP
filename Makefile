.PHONY: help build build-all build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 clean

help: ## 显示帮助信息
	@echo "JueJin-MCP 构建命令："
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""

build: ## 构建当前平台
	@bash scripts/build.sh

build-all: ## 构建所有平台
	@bash scripts/build-all.sh

build-linux-amd64: ## 构建 Linux AMD64 版本
	@bash scripts/build-linux-amd64.sh

build-linux-arm64: ## 构建 Linux ARM64 版本
	@bash scripts/build-linux-arm64.sh

build-darwin-amd64: ## 构建 macOS Intel 版本
	@bash scripts/build-darwin-amd64.sh

build-darwin-arm64: ## 构建 macOS Apple Silicon 版本
	@bash scripts/build-darwin-arm64.sh

build-windows-amd64: ## 构建 Windows AMD64 版本
	@bash scripts/build-windows-amd64.sh

clean: ## 清理构建产物
	@echo "Cleaning build artifacts..."
	@rm -rf dist
	@echo "✅ Clean completed!"

test: ## 运行测试
	@go test -v ./...

deps: ## 下载依赖
	@go mod download
	@go mod tidy

run: ## 运行应用
	@go run .

