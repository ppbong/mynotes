// Package interfaces 演示接口、类型断言、接口嵌入与泛型。
package interfaces

import (
	"bytes"
	"fmt"
	"io"
	"math"

	"github.com/ppbong/go-demo/internal/demo"
)

// Shape 是由调用方（使用方）定义的小接口：
// 需要哪些行为就声明哪些方法，接口越小越容易被复用。
type Shape interface {
	Name() string
	Area() float64
}

// Rect 隐式实现 Shape，无需显式声明 implements。
type Rect struct{ W, H float64 }

// Name 返回图形名称。
func (r Rect) Name() string { return "矩形" }

// Area 返回矩形面积。
func (r Rect) Area() float64 { return r.W * r.H }

// Circle 同样实现 Shape。
type Circle struct{ R float64 }

// Name 返回图形名称。
func (c Circle) Name() string { return "圆" }

// Area 返回圆面积。
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

// Money 演示 fmt.Stringer：实现 String 后，fmt 打印时会自动调用。
type Money struct {
	Amount   int64 // 以“分”为单位存储，避免浮点误差
	Currency string
}

// String 实现 fmt.Stringer 接口。
func (m Money) String() string {
	return fmt.Sprintf("%s %d.%02d", m.Currency, m.Amount/100, m.Amount%100)
}

// Number 是泛型约束：允许底层类型为 int/int64/float64 的类型。
type Number interface {
	~int | ~int64 | ~float64
}

// Sum 对任意数值切片求和。
func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Unique 利用 comparable 约束对元素去重，并保持首次出现的顺序。
func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	result := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}

// Stack 是泛型类型；零值可以直接使用（nil 切片可 append）。
type Stack[T any] struct {
	items []T
}

// Push 压栈。
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop 出栈；第二个返回值表示栈是否非空。
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

// Run 依次执行接口与泛型示例。
func Run() {
	demo.Section("接口实现多态")
	polymorphism()

	demo.Section("类型断言与 type switch")
	assertions()

	demo.Section("接口嵌入与 Stringer")
	embedding()

	demo.Section("泛型函数与泛型类型")
	generics()
}

func polymorphism() {
	shapes := []Shape{Rect{W: 2, H: 3}, Circle{R: 1}}
	total := 0.0
	for _, s := range shapes {
		demo.KV(s.Name(), fmt.Sprintf("%.2f", s.Area()))
		total += s.Area()
	}
	demo.KV("总面积", fmt.Sprintf("%.2f", total))
}

func assertions() {
	var v any = "hello"

	// 带 ok 的类型断言不会 panic，是安全写法。
	s, ok := v.(string)
	demo.KV("断言为 string", fmt.Sprintf("%q ok=%t", s, ok))

	if n, ok := v.(int); ok {
		demo.KV("断言为 int", n)
	} else {
		demo.KV("断言为 int", "失败，实际类型不是 int")
	}

	// type switch 按动态类型分支处理。
	switch x := v.(type) {
	case string:
		demo.KV("type switch", "string: "+x)
	case int:
		demo.KV("type switch", fmt.Sprintf("int: %d", x))
	default:
		demo.KV("type switch", fmt.Sprintf("未知类型 %T", x))
	}

	// any 只是 interface{} 的别名（Go 1.18+）。
	var anything any = 3.14
	demo.KV("any 的动态类型", fmt.Sprintf("%T", anything))
}

func embedding() {
	// 接口嵌入：用小接口组合出更大的接口，避免“胖接口”。
	var buf bytes.Buffer // bytes.Buffer 同时实现 io.Reader 与 io.Writer
	var rw io.ReadWriter = &buf

	if _, err := rw.Write([]byte("hello, io")); err != nil {
		demo.KV("写入失败", err)
		return
	}
	data, err := io.ReadAll(rw)
	if err != nil {
		demo.KV("读取失败", err)
		return
	}
	demo.KV("io.ReadWriter 读回", string(data))

	price := Money{Amount: 123456, Currency: "CNY"}
	demo.KV("Stringer 输出", price)
}

func generics() {
	demo.KV("Sum[int]", Sum([]int{1, 2, 3}))
	demo.KV("Sum[float64]", Sum([]float64{1.5, 2.5}))
	demo.KV("Unique", Unique([]string{"a", "b", "a", "c"}))

	var st Stack[string]
	st.Push("go")
	st.Push("generics")
	top, ok := st.Pop()
	demo.KV("Stack.Pop", fmt.Sprintf("%s ok=%t", top, ok))
	demo.KV("Stack.Len", len(st.items))
}
