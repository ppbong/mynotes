# [Golang](https://golang.google.cn/)

## 安装

```bash
# download
wget https://dl.google.com/go/go1.27.3.linux-amd64.tar.gz

# install
sudo tar xzvf go1.27.3.linux-amd64.tar.gz -C /usr/local
```

## 设置环境变量

```bash
# ~/.bashrc
export PATH=$PATH:$/usr/local/go/bin
```

## 验证

```bash
go version
```

## 运行代码

```go
// main.go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

```bash
# 直接运行（开发与调试）
go run main.go

#编译运行（生产）
go build main.go
```

## 交叉编译

```bash
# macos 编译 linux
GOOS=linux GOARCH=amd64 go build -o app-linux main.go

# linux 编译 windows
GOOS=windows GOARCH=amd64 go build -o app.exe main.go

# windows 编译 macos
$env:GOOS="darwin"
$env:GOARCH="arm64"
go build -o app-arm64 main.go
```

## 常用命令

|命令|功能|常用场景|
|-----|-----|-----|
|go build|	编译源代码| 生成可执行文件，用于部署|
|go run| 编译并运行| 快速测试代码，开发调试|
|go test| 执行测试| 运行 _test.go 文件中的测试函数|
|go mod init| 初始化模块| 创建新项目时生成 go.mod 文件|
|go mod tidy| 整理依赖| 下载缺失的依赖，移除未使用的依赖|
|go fmt| 格式化代码| 统一代码风格（Go 社区强制使用）|
|go vet| 静态分析| 检查代码中的可疑构造|
|go get| 安装依赖包| 下载并安装指定的包|
|go doc| 查看文档| 在终端查看包或函数的文档|
|go env| 查看环境变量| 查看 GOPATH、GOROOT 等配置|
|go clean| 清理构建缓存| 释放磁盘空间|

## 工程（Go Modules 管理依赖）

```bash
mkdir my-go-project && cd my-go-project
go mod init github.com/myrepo/my-go-project  # 将生成 go.mod 文件
```

## 代理设置（模块下载失败或速度慢时）

```bash
# 查看
go env
go env GOPROXY

# 启用模块（默认auto）
go env -w GO111MODULE=on

# 七牛镜像
go env -w GOPROXY=https://goproxy.cn,direct

# 阿里镜像
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```
