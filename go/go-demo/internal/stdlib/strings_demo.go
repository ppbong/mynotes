package stdlib

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/ppbong/go-demo/internal/demo"
)

// Strings 演示 strings / strconv / fmt / bufio 的常用函数。
func Strings() {
	demo.Section("strings 常用函数")
	s := "  Go 标准库学习  "

	demo.KV("TrimSpace", strings.TrimSpace(s))
	demo.KV("Contains", strings.Contains(s, "标准库"))
	demo.KV("HasPrefix", strings.HasPrefix(strings.TrimSpace(s), "Go"))
	demo.KV("ReplaceAll", strings.ReplaceAll("a-b-c", "-", "/"))
	demo.KV("Split", strings.Split("a,b,c", ","))
	demo.KV("Fields", strings.Fields("  foo   bar  baz "))
	demo.KV("Join", strings.Join([]string{"go", "rust", "c"}, ", "))
	demo.KV("Repeat", strings.Repeat("go", 3))
	demo.KV("ToUpper", strings.ToUpper("gopher"))

	// Cut 比 SplitN 语义更清晰：一次切分并返回是否找到。
	key, value, found := strings.Cut("name=alice", "=")
	demo.KV("Cut", fmt.Sprintf("key=%s value=%s found=%t", key, value, found))

	demo.Section("strconv 字符串与数字互转")
	n, err := strconv.Atoi("42")
	if err != nil {
		demo.KV("Atoi 失败", err)
		return
	}
	demo.KV("Atoi", n)
	demo.KV("Itoa", strconv.Itoa(99))

	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		demo.KV("ParseFloat 失败", err)
		return
	}
	demo.KV("ParseFloat", f)
	demo.KV("FormatFloat", strconv.FormatFloat(f, 'f', 3, 64))

	b, err := strconv.ParseBool("true")
	if err != nil {
		demo.KV("ParseBool 失败", err)
		return
	}
	demo.KV("ParseBool", b)

	demo.Section("fmt 常用动词")
	type user struct {
		Name string
		Age  int
	}
	u := user{"Alice", 20}
	demo.KV("%v", fmt.Sprintf("%v", u))
	demo.KV("%+v", fmt.Sprintf("%+v", u))
	demo.KV("%#v", fmt.Sprintf("%#v", u))
	demo.KV("%T", fmt.Sprintf("%T", u))
	demo.KV("%q", fmt.Sprintf("%q", "带引号"))
	demo.KV("%.2f", fmt.Sprintf("%.2f", 3.14159))
	demo.KV("%x", fmt.Sprintf("%x", 255))

	demo.Section("bufio 按行读取")
	scanner := bufio.NewScanner(strings.NewReader("第一行\n第二行"))
	line := 0
	for scanner.Scan() {
		line++
		demo.KV(fmt.Sprintf("第 %d 行", line), scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		demo.KV("scanner 错误", err)
	}
}
