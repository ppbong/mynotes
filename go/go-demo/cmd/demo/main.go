// Command demo 是 Go 学习项目的统一入口。
//
// 用法：
//
//	go run ./cmd/demo            # 查看帮助
//	go run ./cmd/demo all        # 运行全部示例
//	go run ./cmd/demo basics     # 运行指定示例
package main

import (
	"fmt"
	"os"

	"github.com/ppbong/go-demo/internal/basics"
	"github.com/ppbong/go-demo/internal/collections"
	"github.com/ppbong/go-demo/internal/concurrency"
	"github.com/ppbong/go-demo/internal/errorsdemo"
	"github.com/ppbong/go-demo/internal/interfaces"
	"github.com/ppbong/go-demo/internal/patterns"
	"github.com/ppbong/go-demo/internal/stdlib"
)

// lesson 描述一个可独立运行的示例模块。
type lesson struct {
	name string
	desc string
	run  func()
}

var lessons = []lesson{
	{"basics", "变量、常量、控制流、函数、指针、结构体", basics.Run},
	{"collections", "数组、切片、map、排序、字符串", collections.Run},
	{"interfaces", "接口、类型断言、嵌入、泛型", interfaces.Run},
	{"errors", "错误包装、errors.Is/As/Join、panic/recover", errorsdemo.Run},
	{"concurrency", "goroutine、channel、select、sync、context", concurrency.Run},
	{"stdlib", "strings、json、time、文件、HTTP、regexp", stdlib.Run},
	{"patterns", "函数选项、单例、中间件、重试、优雅关闭", patterns.Run},
}

func main() {
	if len(os.Args) > 1 {
		switch arg := os.Args[1]; arg {
		case "all":
			for _, l := range lessons {
				l.run()
			}
		case "help", "-h", "--help":
			usage()
		default:
			for _, l := range lessons {
				if l.name == arg {
					l.run()
					return
				}
			}
			fmt.Fprintf(os.Stderr, "未知示例: %q\n\n", arg)
			usage()
			os.Exit(1)
		}
		return
	}
	usage()
}

// usage 打印帮助信息，列出所有示例及其说明。
func usage() {
	fmt.Println("Go 语言学习示例集")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  go run ./cmd/demo <示例名>")
	fmt.Println("  go run ./cmd/demo all")
	fmt.Println()
	fmt.Println("可用示例:")
	for _, l := range lessons {
		fmt.Printf("  %-12s %s\n", l.name, l.desc)
	}
}
