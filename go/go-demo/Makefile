.PHONY: help run run-all test race bench cover vet fmt build clean

BIN := bin/go-demo

help: ## 显示可用命令
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

run: ## 运行示例（不带参数时显示帮助）
	go run ./cmd/demo

run-all: ## 运行全部示例
	go run ./cmd/demo all

test: ## 运行单元测试
	go test ./...

race: ## 开启竞态检测运行测试
	go test -race ./...

bench: ## 运行基准测试
	go test -bench=. -benchmem ./...

cover: ## 生成测试覆盖率报告
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

vet: ## 静态检查
	go vet ./...

fmt: ## 格式化代码
	gofmt -w .

build: ## 编译二进制到 bin/
	go build -o $(BIN) ./cmd/demo

clean: ## 清理构建产物
	rm -rf bin coverage.out coverage.html
