// Package basics 演示 Go 的基础语法与常用写法。
package basics

import (
	"errors"
	"fmt"

	"github.com/ppbong/go-demo/internal/demo"
)

// Run 依次执行基础语法示例。
func Run() {
	demo.Section("变量与零值")
	variables()

	demo.Section("常量与 iota")
	constants()

	demo.Section("基本类型与显式转换")
	types()

	demo.Section("控制流")
	controlFlow()

	demo.Section("函数特性")
	functions()

	demo.Section("指针与值语义")
	pointers()

	demo.Section("结构体、方法与嵌入")
	structs()
}

func variables() {
	// 声明但未赋值时，变量会被初始化为该类型的“零值”。
	var i int
	var s string
	var ok bool

	// 短变量声明只能用在函数内部，类型由右侧表达式推导。
	name := "gopher"
	pi := 3.14159

	demo.KV("int 零值", i)
	demo.KV("string 零值", fmt.Sprintf("%q", s))
	demo.KV("bool 零值", ok)
	demo.KV("短变量声明", name)
	demo.KV("类型推导", fmt.Sprintf("%T = %v", pi, pi))

	// Go 支持多重赋值，交换变量不需要临时变量。
	a, b := 1, 2
	a, b = b, a
	demo.KV("交换后 a,b", fmt.Sprintf("%d,%d", a, b))
}

// Weekday 用自定义类型 + iota 实现枚举，避免直接使用魔法数字。
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// String 让 Weekday 实现 fmt.Stringer，打印时会自动调用。
func (d Weekday) String() string {
	return [...]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}[d]
}

const (
	// 同一个 const 块中 iota 从 0 开始递增。
	KB = 1 << (10 * (iota + 1)) // 1 << 10
	MB                          // 1 << 20
	GB                          // 1 << 30
)

func constants() {
	// 无类型常量拥有更高的精度，使用时才确定具体类型。
	const Pi = 3.141592653589793
	demo.KV("无类型常量 Pi", Pi)
	demo.KV("枚举 Friday", Friday)
	demo.KV("KB/MB/GB", fmt.Sprintf("%d/%d/%d", KB, MB, GB))

	const (
		StatusActive   = "active"
		StatusInactive = "inactive"
	)
	demo.KV("字符串常量", StatusActive)
	demo.KV("常量参与计算", fmt.Sprintf("%.2f", Pi*2))
}

func types() {
	// Go 不会做隐式类型转换，不同类型之间必须显式转换。
	var i32 int32 = 100
	i64 := int64(i32)
	f := float64(i32)

	demo.KV("int32 -> int64", i64)
	demo.KV("int32 -> float64", f)

	// rune 是 int32 的别名，表示一个 Unicode 码点；
	// byte 是 uint8 的别名，表示一个字节。
	var r rune = '中'
	var b byte = 'A'
	demo.KV("rune '中'", fmt.Sprintf("%c = %d", r, r))
	demo.KV("byte 'A'", fmt.Sprintf("%c = %d", b, b))

	// 整数除法的结果仍是整数，需要小数时先转 float64。
	demo.KV("7/2", 7/2)
	demo.KV("float64(7)/2", float64(7)/2)
}

func controlFlow() {
	// if 支持初始化语句，变量作用域仅限该 if/else 链。
	if score := 82; score >= 90 {
		demo.KV("成绩等级", "A")
	} else if score >= 80 {
		demo.KV("成绩等级", "B")
	} else {
		demo.KV("成绩等级", "C")
	}

	// 无表达式的 switch 等价于更清晰的 if-else 链。
	switch hour := 9; {
	case hour < 6:
		demo.KV("问候", "凌晨好")
	case hour < 12:
		demo.KV("问候", "上午好")
	default:
		demo.KV("问候", "你好")
	}

	// 默认 case 执行完即结束；只有 fallthrough 才会进入下一个分支。
	switch 2 {
	case 1:
		demo.KV("switch", "one")
	case 2:
		demo.KV("switch", "two")
		fallthrough
	case 3:
		demo.KV("fallthrough", "three")
	}

	// for 的三种常见写法：三段式、while 风格、无限循环。
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	demo.KV("for 三段式 1..5", sum)

	n := 0
	for n < 3 {
		n++
	}
	demo.KV("for while 风格", n)

	count := 0
	for {
		count++
		if count == 2 {
			break
		}
	}
	demo.KV("for 无限循环 + break", count)

	// range 同时返回下标和值；遍历字符串时值类型是 rune。
	for i, v := range []string{"a", "b"} {
		demo.KV(fmt.Sprintf("range items[%d]", i), v)
	}

	// 标签（label）可以一次性跳出多层嵌套循环。
	found := false
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j == 2 {
				found = true
				break outer
			}
		}
	}
	demo.KV("标签 break", found)
}

// Divide 返回 a/b：演示“结果 + error”这一 Go 最常见的返回值组合。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// minMax 演示命名返回值与变参；裸 return 只建议用于短函数。
func minMax(nums ...int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

// counter 返回一个闭包，闭包捕获并共享 count 变量。
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// safeDivide 用 recover 把 panic 转换成 error，避免进程崩溃。
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic 已被恢复: %v", r)
		}
	}()
	return a / b, nil // b == 0 时触发 runtime panic
}

func deferOrder() {
	for i := 1; i <= 3; i++ {
		// defer 的参数在注册时求值，执行顺序为后进先出（LIFO）。
		defer fmt.Printf("  defer %d 执行\n", i)
	}
	fmt.Println("  函数体结束，开始执行 defer")
}

func functions() {
	q, err := Divide(10, 3)
	demo.KV("Divide(10,3)", fmt.Sprintf("q=%d err=%v", q, err))

	min, max := minMax(3, 1, 4, 1, 5)
	demo.KV("minMax", fmt.Sprintf("min=%d max=%d", min, max))

	next := counter()
	demo.KV("闭包第 1 次调用", next())
	demo.KV("闭包第 2 次调用", next())

	deferOrder()

	res, e := safeDivide(1, 0)
	demo.KV("safeDivide", fmt.Sprintf("result=%d err=%v", res, e))
}

func incByValue(n int)    { n++ }
func incByPointer(n *int) { *n++ }

func pointers() {
	v := 42
	p := &v // & 取地址
	*p = 43 // * 解引用
	demo.KV("通过指针修改", v)

	x := 1
	incByValue(x)
	demo.KV("值传递后 x", x)
	incByPointer(&x)
	demo.KV("指针传递后 x", x)

	// new(T) 分配一个 T 类型的零值并返回指针。
	pp := new(int)
	demo.KV("new(int) 的零值", *pp)

	// 结构体是值类型：赋值会整体拷贝；指针则共享同一份数据。
	type point struct{ X, Y int }
	pt1 := point{1, 2}
	pt2 := pt1
	pt2.X = 100
	demo.KV("值拷贝 pt1.X", pt1.X)

	pt3 := &pt1
	pt3.X = 200
	demo.KV("指针共享 pt1.X", pt1.X)
}

// Animal 作为被嵌入的类型，其字段和方法会被提升到外层结构体。
type Animal struct {
	Name string
}

// Speak 是值接收者方法。
func (a Animal) Speak() string { return a.Name + " 发出声音" }

// Dog 通过匿名嵌入 Animal 复用字段和方法（组合优于继承）。
type Dog struct {
	Animal
	Breed string
}

// Counter 演示值接收者与指针接收者的区别。
type Counter struct {
	n int
}

// Value 只读，使用值接收者即可。
func (c Counter) Value() int { return c.n }

// Inc 需要修改状态，必须使用指针接收者。
func (c *Counter) Inc() { c.n++ }

func structs() {
	d := Dog{Animal: Animal{Name: "旺财"}, Breed: "柴犬"}
	demo.KV("提升字段 d.Name", d.Name)
	demo.KV("提升方法 d.Speak()", d.Speak())

	c := Counter{}
	for i := 0; i < 3; i++ {
		c.Inc()
	}
	demo.KV("指针接收者修改", c.Value())

	// 拷贝出来的结构体拥有独立状态。
	c2 := c
	c2.Inc()
	demo.KV("副本自增", fmt.Sprintf("c=%d c2=%d", c.Value(), c2.Value()))

	// 匿名结构体适合临时数据结构，例如测试数据。
	u := struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 18}
	demo.KV("匿名结构体", fmt.Sprintf("%+v", u))
}
