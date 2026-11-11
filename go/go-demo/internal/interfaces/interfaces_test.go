package interfaces

import (
	"fmt"
	"slices"
	"testing"
)

// TestUnique 验证去重结果与顺序。
func TestUnique(t *testing.T) {
	got := Unique([]string{"a", "b", "a", "c", "b"})
	want := []string{"a", "b", "c"}

	if !slices.Equal(got, want) {
		t.Fatalf("Unique() = %v, want %v", got, want)
	}
}

// TestSum 验证泛型函数对不同数值类型的支持。
func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([]int) = %d, want 6", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4.0 {
		t.Errorf("Sum([]float64) = %v, want 4.0", got)
	}
}

// ExampleSum 是可执行文档：go test 会校验输出是否与注释一致。
func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3}))
	// Output: 6
}
