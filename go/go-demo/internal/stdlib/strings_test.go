package stdlib

import (
	"strings"
	"testing"
)

// BenchmarkConcatPlus 演示用 + 拼接字符串的基准测试。
func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "go"
		}
		_ = s
	}
}

// BenchmarkStringBuilder 与 + 拼接对比，Builder 通常快一个数量级。
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for j := 0; j < 100; j++ {
			sb.WriteString("go")
		}
		_ = sb.String()
	}
}
