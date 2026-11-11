// Package collections 演示数组、切片、映射、排序与字符串处理。
package collections

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/ppbong/go-demo/internal/demo"
)

// Person 用作排序示例的数据结构。
type Person struct {
	Name string
	Age  int
}

// Run 依次执行集合相关示例。
func Run() {
	demo.Section("数组：长度固定、值语义")
	arrays()

	demo.Section("切片：动态数组、共享底层数组")
	sliceDemo()

	demo.Section("slices 包常用函数")
	slicesPackage()

	demo.Section("map 常用操作")
	mapsDemo()

	demo.Section("排序")
	sortDemo()

	demo.Section("字符串、字节与 rune")
	stringsRunes()
}

func arrays() {
	var a [3]int
	a[0] = 1

	// 数组是值类型，赋值会复制整个数组。
	b := a
	b[0] = 99

	demo.KV("数组 a", a)
	demo.KV("副本 b", b)
	demo.KV("len(a)", len(a))

	// 用 ... 让编译器推断数组长度。
	c := [...]string{"go", "rust", "c"}
	demo.KV("字面量数组", c)
}

func sliceDemo() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	demo.KV("append 后", fmt.Sprintf("%v len=%d cap=%d", s, len(s), cap(s)))

	// 切片只是底层数组的视图，修改子切片会影响原切片。
	sub := s[1:3]
	sub[0] = 99
	demo.KV("修改子切片后 s", s)

	// 需要独立数据时使用 Clone 或 copy。
	clone := slices.Clone(s)
	clone[0] = -1
	demo.KV("Clone 后互不影响", fmt.Sprintf("s[0]=%d clone[0]=%d", s[0], clone[0]))

	// 删除下标 2 的元素，保持剩余元素顺序。
	s = slices.Delete(s, 2, 3)
	demo.KV("Delete(s,2,3) 后", s)

	// copy 把数据复制到已存在的目标切片，返回复制的元素个数。
	dst := make([]int, len(s))
	n := copy(dst, s)
	demo.KV("copy 元素个数", n)
	demo.KV("copy 结果", dst)

	// nil 切片可以直接 append，无需提前 make。
	var nilSlice []int
	nilSlice = append(nilSlice, 1)
	demo.KV("nil 切片 append", nilSlice)
	demo.KV("nil 切片 len", len(nilSlice))

	// 三索引切片 s[low:high:max] 限制容量，防止 append 覆盖原数组。
	limited := s[0:2:2]
	demo.KV("三索引切片 cap", cap(limited))
}

func slicesPackage() {
	nums := []int{5, 2, 8, 1}
	slices.Sort(nums)
	demo.KV("slices.Sort", nums)

	demo.KV("slices.Contains(8)", slices.Contains(nums, 8))
	demo.KV("slices.Index(8)", slices.Index(nums, 8))
	demo.KV("slices.Max", slices.Max(nums))
	demo.KV("slices.Min", slices.Min(nums))

	// Insert 在下标 1 处插入元素。
	nums = slices.Insert(nums, 1, 42)
	demo.KV("slices.Insert", nums)

	slices.Reverse(nums)
	demo.KV("slices.Reverse", nums)
}

func mapsDemo() {
	scores := map[string]int{"alice": 90, "bob": 85}
	scores["carol"] = 95

	// comma-ok 是判断 key 是否存在的标准写法。
	v, ok := scores["dave"]
	demo.KV("comma-ok", fmt.Sprintf("v=%d ok=%t", v, ok))

	// delete 删除不存在的 key 是安全的，不会 panic。
	delete(scores, "bob")
	demo.KV("len(scores)", len(scores))

	// map 的遍历顺序是随机的；需要稳定输出时先对 key 排序。
	keys := slices.Sorted(maps.Keys(scores))
	for _, k := range keys {
		demo.KV("map["+k+"]", scores[k])
	}

	// 用 map 实现集合（set）：值使用空结构体，不占额外内存。
	set := map[string]struct{}{}
	set["go"] = struct{}{}
	set["rust"] = struct{}{}
	_, exists := set["go"]
	demo.KV("set 中是否存在 go", exists)

	// 使用前必须 make，向 nil map 写入会 panic。
	var bad map[string]int
	demo.KV("nil map 读取", bad["x"])
	demo.KV("nil map 长度", len(bad))
}

func sortDemo() {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 30},
		{Name: "Dave", Age: 20},
	}

	// sort.Slice 使用比较函数排序，不需要实现 sort.Interface。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	demo.KV("sort.Slice 按年龄", people)

	// sort.SliceStable 保证相等元素的原始相对顺序。
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	demo.KV("SliceStable", people)

	// slices.SortFunc 是类型安全的新写法；cmp.Compare 可直接比较有序类型。
	slices.SortFunc(people, func(a, b Person) int {
		if c := cmp.Compare(a.Age, b.Age); c != 0 {
			return c
		}
		return cmp.Compare(a.Name, b.Name)
	})
	demo.KV("多字段排序", people)
}

func stringsRunes() {
	s := "Hello, 世界"

	// len 返回字节数；统计字符数要用 utf8.RuneCountInString。
	demo.KV("len 字节数", len(s))
	demo.KV("字符数", utf8.RuneCountInString(s))
	demo.KV("[]rune 长度", len([]rune(s)))

	// range 遍历字符串时，下标是字节索引，值是 rune。
	for i, r := range s {
		if r == '世' {
			demo.KV("'世' 的字节索引", i)
			break
		}
	}

	// 频繁拼接字符串应使用 strings.Builder，避免反复分配内存。
	var sb strings.Builder
	sb.Grow(32) // 预估容量，减少扩容
	for i := 1; i <= 3; i++ {
		fmt.Fprintf(&sb, "第%d项;", i)
	}
	demo.KV("strings.Builder", sb.String())

	// 字符串比较、查找、分割都可以直接用标准库完成。
	demo.KV("strings.Cut", cutDemo("name=alice"))
}

func cutDemo(s string) string {
	// strings.Cut 从字符串中提取子字符串，返回子字符串、子字符串是否在字符串中。
	key, value, found := strings.Cut(s, "=")
	return fmt.Sprintf("key=%s value=%s found=%t", key, value, found)
}
