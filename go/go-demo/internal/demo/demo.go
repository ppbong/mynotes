// Package demo 提供示例程序统一的输出辅助函数。
package demo

import "fmt"

// Section 打印章节标题，让每个示例的输出层次清晰。
func Section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// KV 以对齐的“键: 值”形式打印结果。
func KV(key string, value any) {
	fmt.Printf("%-26s %v\n", key+":", value)
}
