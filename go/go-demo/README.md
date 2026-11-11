# Go 语言学习示例集

一个可运行、可测试的 Go 学习型项目，覆盖语言基础语法、标准库高频用法与常见工程实践。每个模块都可以单独运行，代码中配有简要注释，适合边看边改、快速查阅。

## 特性

- 纯标准库实现，无第三方依赖，`go run` 即可运行
- 按主题拆分模块，一个命令运行一个知识点
- 示例代码配套表驱动测试、可执行文档（Example）与基准测试
- 包含工程化的目录结构、Makefile 与构建脚本约定
- 兼顾现代 Go 写法：泛型、`slices`/`maps`、`rand/v2`、`http.ServeMux` 路由参数

## 环境要求

- Go **1.23 或更高版本**（示例使用了 `maps.Keys` 迭代器、`math/rand/v2`、`http.ServeMux` 方法与路径参数等特性）
- 可选：`make`（用于快捷命令）

检查环境：

```bash
go version
go env GOPROXY
```

国内网络可设置代理加速依赖下载（本项目无外部依赖，仅便于后续扩展）：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

## 目录结构

```text
go-demo/
├── cmd/
│   └── demo/                 # 程序入口：统一 CLI，分发到各示例模块
│       └── main.go
├── internal/                 # 内部包：只允许本模块导入，避免被外部复用
│   ├── demo/                 # 输出格式辅助（Section / KV）
│   ├── basics/               # 变量、常量、类型、控制流、函数、指针、结构体
│   ├── collections/          # 数组、切片、map、排序、字符串与 rune
│   ├── interfaces/           # 接口、类型断言、接口嵌入、泛型
│   ├── errorsdemo/           # 错误包装、errors.Is/As/Join、panic/recover
│   ├── concurrency/          # goroutine、channel、select、sync、context
│   ├── stdlib/               # strings/json/time/fs/io/http/regexp/rand
│   └── patterns/             # 函数选项、单例、中间件、仓储、重试、优雅关闭
├── go.mod                    # 模块定义与 Go 版本
├── Makefile                  # 常用开发命令
└── README.md
```

约定说明：

- `cmd/` 存放可执行程序入口，每个子目录对应一个二进制
- `internal/` 存放不对外暴露的内部实现，借助 Go 的内部包机制限制导入范围
- 每个示例包导出 `Run()`，由 `cmd/demo` 统一调度，便于单独运行与组合运行

## 快速开始

```bash
# 查看帮助与示例列表
go run ./cmd/demo

# 运行全部示例
go run ./cmd/demo all

# 运行单个主题
go run ./cmd/demo basics
go run ./cmd/demo concurrency
```

使用 Makefile：

```bash
make help      # 查看全部命令
make run-all   # 运行全部示例
make test      # 运行测试
make race      # 竞态检测
make bench     # 基准测试
make cover     # 生成覆盖率报告 coverage.html
make vet       # 静态检查
make build     # 编译到 bin/go-demo
```

## 学习路线

建议按以下顺序阅读源码，每看完一个模块就运行对应命令验证输出：

| 顺序 | 模块 | 命令 | 核心内容 |
| --- | --- | --- | --- |
| 1 | `internal/basics` | `go run ./cmd/demo basics` | 零值与短变量声明、`iota` 枚举、显式类型转换、`if/switch/for`、标签 `break`、多返回值、闭包、`defer` LIFO、`panic/recover`、指针与值语义、结构体与嵌入 |
| 2 | `internal/collections` | `go run ./cmd/demo collections` | 数组值拷贝、切片扩容与共享底层数组、`slices`/`maps` 包、set 写法、`sort.Slice` 与 `slices.SortFunc`、字符串字节与 rune |
| 3 | `internal/interfaces` | `go run ./cmd/demo interfaces` | 隐式接口实现、小接口设计、类型断言与 type switch、接口嵌入、`fmt.Stringer`、泛型约束与泛型类型 |
| 4 | `internal/errorsdemo` | `go run ./cmd/demo errors` | 哨兵错误、`%w` 包装、`errors.Is/As/Join`、`recover` 转 error、`defer` 关闭资源、typed nil 陷阱 |
| 5 | `internal/concurrency` | `go run ./cmd/demo concurrency` | `WaitGroup`、缓冲/无缓冲 channel、`select` 超时、worker pool、`Mutex`/`atomic`/`Once`/`Pool`、`context` 取消与传值、pipeline |
| 6 | `internal/stdlib` | `go run ./cmd/demo stdlib` | `strings`/`strconv`/`fmt`、`bufio`、`encoding/json`、`time`、文件与目录遍历、`io.Reader/Writer`、`net/http`、`regexp`、随机数 |
| 7 | `internal/patterns` | `go run ./cmd/demo patterns` | 函数选项模式、并发安全单例、HTTP 中间件链、接口 + 仓储、指数退避重试、服务优雅关闭 |

## 从零构建本项目

如果你要复刻或扩展这个项目，可以按下面的步骤操作。

### 1. 初始化模块

```bash
mkdir go-demo && cd go-demo
go mod init github.com/ppbong/go-demo
```

`go.mod` 中的模块路径是包导入前缀，实际项目应替换为自己的仓库地址：

```go
module github.com/ppbong/go-demo

go 1.23
```

### 2. 创建目录骨架

```bash
mkdir -p cmd/demo internal/basics internal/collections internal/interfaces \
         internal/errorsdemo internal/concurrency internal/stdlib internal/patterns
```

### 3. 编写示例包

每个示例包遵循同一约定：

```go
// Package xxx 说明该包演示的主题。
package xxx

// Run 是该示例的统一入口。
func Run() {
    // 演示代码
}
```

### 4. 注册到 CLI

在 `cmd/demo/main.go` 的 `lessons` 切片中增加一条记录即可：

```go
var lessons = []lesson{
    {"basics", "变量、常量、控制流、函数、指针、结构体", basics.Run},
    // 新增示例：
    {"mytopic", "示例说明", mytopic.Run},
}
```

### 5. 编译与交叉编译

```bash
# 本机编译
go build -o bin/go-demo ./cmd/demo

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o bin/go-demo-linux-amd64 ./cmd/demo
GOOS=windows GOARCH=amd64 go build -o bin/go-demo-windows-amd64 ./cmd/demo
GOOS=darwin GOARCH=amd64 go build -o bin/go-demo-macos-amd64 ./cmd/demo
GOOS=darwin GOARCH=arm64 go build -o bin/go-demo-macos-arm64 ./cmd/demo
```

`go run` 适合日常调试，产物在临时目录；需要分发时使用 `go build`。

## 测试与质量保障

```bash
go test ./...                 # 单元测试
go test -race ./...           # 竞态检测（并发代码必跑）
go test -cover ./...          # 覆盖率
go test -bench=. -benchmem ./...   # 基准测试
go vet ./...                  # 官方静态检查
gofmt -l .                    # 列出未格式化的文件
```

项目内的测试示例：

- `internal/basics/basics_test.go`：表驱动测试，覆盖正常与异常分支
- `internal/interfaces/interfaces_test.go`：泛型函数测试 + `ExampleSum` 可执行文档
- `internal/patterns/patterns_test.go`：重试与仓储的行为测试
- `internal/stdlib/strings_test.go`：`+` 拼接与 `strings.Builder` 的基准对比

基准测试实测结果参考（不同机器会有差异）：

```text
BenchmarkConcatPlus-4       176814    7955 ns/op   10736 B/op   99 allocs/op
BenchmarkStringBuilder-4   2500447     505 ns/op     504 B/op    6 allocs/op
```

## 最佳实践清单

代码中已尽量体现以下惯例，可作为日常开发的检查项：

1. **错误处理**：用 `fmt.Errorf("...: %w", err)` 保留错误链；用 `errors.Is` 判断哨兵错误、`errors.As` 提取自定义错误；不要用字符串比较判断错误。
2. **资源释放**：`defer f.Close()` 紧跟资源创建；需要上报关闭错误时使用命名返回值在 `defer` 中合并错误。
3. **并发安全**：共享状态加锁或用 `atomic`；goroutine 必须有明确的退出路径；`context.Context` 作为第一个参数并向下传递。
4. **接口设计**：接口定义在使用方（调用方）而不是实现方；接口尽量小；遵循 “accept interfaces, return structs”。
5. **零值可用**：设计类型时尽量让零值即可使用，例如 `sync.Mutex`、`bytes.Buffer`、自定义 `Stack`。
6. **显式优于隐式**：不做隐式类型转换；配置通过函数选项显式传入；时间解析显式指定时区。
7. **测试习惯**：表驱动测试 + `t.Run` 子测试；并发代码跑 `-race`；性能敏感路径用 `-bench` 验证。
8. **格式化与检查**：提交前执行 `gofmt`、`go vet`、`go test ./...`。
9. **避免 `init` 副作用**：初始化逻辑写成显式函数，便于测试与控制执行顺序。
10. **HTTP 客户端**：始终设置超时并关闭 `resp.Body`，避免连接泄漏。

## 常见问题

**运行时报 `go: command not found`**

Go 未安装或未加入 `PATH`。安装后确认：

```bash
export PATH=$PATH:/usr/local/go/bin
go version
```

**`go run ./cmd/demo` 输出“未知示例”**

查看帮助中的示例名，注意大小写，例如 `concurrency` 而不是 `Concurrency`。

**示例中的 HTTP 服务会占用端口吗？**

不会。`stdlib` 与 `patterns` 使用 `httptest` 或 `127.0.0.1:0` 由系统分配临时端口，运行结束后自动关闭。

**想要更严格的静态检查？**

可额外安装 `staticcheck`：

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```
