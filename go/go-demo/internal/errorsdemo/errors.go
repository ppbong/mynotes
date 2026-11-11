// Package errorsdemo 演示 Go 的错误处理惯用法与最佳实践。
package errorsdemo

import (
	"errors"
	"fmt"
	"os"

	"github.com/ppbong/go-demo/internal/demo"
)

// ErrNotFound 是哨兵错误（sentinel error），调用方可以用 errors.Is 判断。
var ErrNotFound = errors.New("资源不存在")

// ValidationError 是带上下文的自定义错误类型，调用方可用 errors.As 提取。
type ValidationError struct {
	Field string
	Value any
}

// Error 实现 error 接口。
func (e *ValidationError) Error() string {
	return fmt.Sprintf("校验失败: 字段 %s 的值 %v 不合法", e.Field, e.Value)
}

// FindUser 用 %w 包装底层错误，形成可追溯的错误链。
func FindUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("FindUser(id=%d): %w", id, ErrNotFound)
	}
	return fmt.Sprintf("user-%d", id), nil
}

// Validate 返回自定义错误类型。
func Validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Value: age}
	}
	return nil
}

// safeDivide 用 recover 把不可恢复的 panic 转换为可处理的 error。
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic 已被恢复: %v", r)
		}
	}()
	return a / b, nil
}

// writeAndClose 演示：defer 中关闭资源，并把关闭错误返回给调用方。
func writeAndClose() (err error) {
	f, err := os.CreateTemp("", "go-demo-*.txt")
	if err != nil {
		return err
	}
	name := f.Name()
	defer os.Remove(name) // 先注册，最后执行（在 Close 之后）

	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	_, err = f.WriteString("hello, defer")
	return err
}

// Run 依次执行错误处理示例。
func Run() {
	demo.Section("哨兵错误与错误包装")
	_, err := FindUser(0)
	demo.KV("错误信息", err)
	demo.KV("errors.Is(ErrNotFound)", errors.Is(err, ErrNotFound))
	demo.KV("errors.Unwrap 一层", errors.Unwrap(err))

	demo.Section("自定义错误与 errors.As")
	err = Validate(-1)
	demo.KV("错误信息", err)
	var ve *ValidationError
	if errors.As(err, &ve) {
		demo.KV("As 提取字段", fmt.Sprintf("Field=%s Value=%v", ve.Field, ve.Value))
	}

	demo.Section("errors.Join 合并多个错误")
	joined := errors.Join(errors.New("数据库连接失败"), os.ErrPermission)
	demo.KV("Join", joined)
	demo.KV("Is(os.ErrPermission)", errors.Is(joined, os.ErrPermission))

	demo.Section("panic / recover")
	res, e := safeDivide(10, 0)
	demo.KV("safeDivide(10,0)", fmt.Sprintf("result=%d err=%v", res, e))

	demo.Section("defer 处理资源关闭")
	demo.KV("写临时文件", writeAndClose())

	demo.Section("typed nil 陷阱")
	var p *ValidationError
	var wrapped error = p
	// 接口值包含“类型 + 数据”两部分：类型非空，接口就不等于 nil。
	demo.KV("wrapped != nil", wrapped != nil)
	demo.KV("wrapped == nil", wrapped == nil)
}
