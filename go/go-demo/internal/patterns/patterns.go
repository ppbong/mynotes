// Package patterns 汇总 Go 工程中常用的设计模式与工程实践。
package patterns

import "github.com/ppbong/go-demo/internal/demo"

// Run 依次执行工程实践示例。
func Run() {
	demo.Section("函数选项模式")
	optionsDemo()

	demo.Section("并发安全单例")
	singletonDemo()

	demo.Section("HTTP 中间件（装饰器）")
	middlewareDemo()

	demo.Section("接口 + 内存仓储")
	repositoryDemo()

	demo.Section("带指数退避的重试")
	retryDemo()

	demo.Section("HTTP 服务优雅关闭")
	gracefulShutdownDemo()
}
