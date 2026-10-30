# Golang 语法

## 变量与常量

```go
package main

import "fmt"

func main() {
    // 方式一：标准声明（显式类型）
    var age int = 25

    // 方式二：声明后赋值（类型在声明时指定）
    var score float64
    score = 98.5

    // 方式三：短变量声明（最常用，仅限函数内部）
    name := "jacK"

    // 方式四：多变量同时声明
    x, y := 10, 20

    // 常量声明（编译时确定，不可修改）
    const Pi = 3.14159
    const AppName = "TEST-APP"

    fmt.Println(name, age, score, x, y, Pi, AppName)
}
```

## 流程控制

```go
package main

import "fmt"

func main() {
    // if 语句：条件表达式不需要括号，但大括号必须有
    score := 85
    if score >= 90 {
        fmt.Println("评级: A")
    } else if score >= 60 {
        fmt.Println("评级: B")
    } else {
        fmt.Println("评级: C")
    }

    // for 循环一：标准三要素形式（类似 C 的 for）
    fmt.Print("计数: ")
    for i := 0; i < 5; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // for 循环二：仅条件（类似 while）
    sum := 0
    n := 1
    for n <= 10 {
        sum += n
        n++
    }
    fmt.Println("1 到 10 之和:", sum)

    // for 循环三：range 遍历（最常用）
    fruits := []string{"apple", "banana", "cherry"}
    for index, fruit := range fruits {
        fmt.Printf("索引 %d: %s\n", index, fruit)
    }

    // switch 语句：每个 case 自动 break，无需手动添加
    day := "周一"
    switch day {
    case "周一":
        fmt.Println("新的一周开始了")
    case "周五":
        fmt.Println("周末快到了")
    default:
        fmt.Println("普通的一天")
    }
}
```

## 函数

```go
package main

import (
    "errors"
    "fmt"
)

// add 函数：两个参数同为 int 类型，返回一个 int
func add(a, b int) int {
    return a + b
}

// divide 函数：返回两个值——结果和错误
// Go 惯用 (result, error) 模式进行错误处理
func divide(a, b float64) (float64, error) {
    if b == 0 {
        // 返回零值和自定义错误
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil // nil 表示没有错误
}

// swap 函数：多返回值用于返回多个计算结果
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    // 调用普通函数
    sum := add(10, 20)
    fmt.Println("10 + 20 =", sum)

    // 调用多返回值函数并检查错误
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Println("10 / 2 =", result)
    }

    // 测试除零情况
    _, err = divide(10, 0) // _ 用于忽略不需要的返回值
    if err != nil {
        fmt.Println("除零错误:", err) // 预期输出此错误
    }

    // 多返回值交换
    first, second := swap("World", "Hello")
    fmt.Println("交换后:", first, second)
}
```

## 结构体与方法

```go
package main

import "fmt"

// User 定义用户结构体（类似于其他语言中的类）
type User struct {
    Name  string // 首字母大写 = 公开字段（可被外部包访问）
    Email string
    age   int    // 首字母小写 = 私有字段（仅包内可访问）
}

// Greet 是为 User 类型定义的方法
// (u User) 是值接收者，方法内对 u 的修改不会影响原始值
func (u User) Greet() string {
    return fmt.Sprintf("你好，我是 %s，邮箱是 %s", u.Name, u.Email)
}

// Birthday 使用指针接收者，方法内可以修改原始值
func (u *User) Birthday() {
    u.age++ // 通过指针直接修改原始结构体的 age 字段
}

func main() {
    // 创建结构体实例
    user := User{
        Name:  "jack",
        Email: "jack@example.com",
        age:   25,
    }

    // 调用方法
    fmt.Println(user.Greet())
    fmt.Printf("生日前年龄: %d\n", user.age)

    user.Birthday() // 使用指针接收者，age 被加 1
    fmt.Printf("生日后年龄: %d\n", user.age)
}
```

## 接口

```go
package main

import (
    "fmt"
    "math"
)

// Shape 接口：定义几何形状的行为规范
type Shape interface {
    Area() float64      // 计算面积
    Perimeter() float64 // 计算周长
}

// Circle 圆形结构体
type Circle struct {
    Radius float64
}

// Circle 实现 Shape 接口（隐式实现，无需显式声明）
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Rectangle 矩形结构体
type Rectangle struct {
    Width, Height float64
}

// Rectangle 也隐式实现了 Shape 接口
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// PrintShape 接受任何实现了 Shape 接口的类型
func PrintShape(s Shape) {
    fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    c := Circle{Radius: 5}          // 半径为 5 的圆
    r := Rectangle{Width: 4, Height: 6} // 宽 4 高 6 的矩形

    fmt.Println("圆形:")
    PrintShape(c)

    fmt.Println("矩形:")
    PrintShape(r)
}
```

## 常量组与枚举

```go
package main

import "fmt"

// 定义性别枚举（Go 语言中没有枚举 enum 关键字）
const (
    Unknown = 0
    Female  = 1
    Male    = 2
)

func main() {
    fmt.Printf("未知=%d, 女性=%d, 男性=%d\n", Unknown, Female, Male)
}
```

## 特殊常量计算器 iota

```go
package main

import "fmt"

// iota 从 0 开始，每行加 1
const (
    a = iota   // 0
    b          // 1（省略值，默认使用上一行的 iota）
    c          // 2
)

func main() {
    fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}
```

```go
package main

import "fmt"

func main() {
    const (
        a = iota   // 0
        b          // 1
        c          // 2
        d = "ha"   // 独立值，iota += 1
        e          // "ha"，iota += 1
        f = 100    // 独立值，iota += 1
        g          // 100，iota += 1
        h = iota   // 7，恢复使用 iota 值
        i          // 8
    )
    fmt.Printf("a=%d, b=%d, c=%d, d=%s, e=%s, f=%d, g=%d, h=%d, i=%d\n", a, b, c, d, e, f, g, h, i)
}
```

```go
package main

import "fmt"

const (
    i = 1 << iota  // 1 << 0 = 1
    j = 3 << iota  // 3 << 1 = 6
    k              // 3 << 2 = 12
    l              // 3 << 3 = 24
)

func main() {
    fmt.Printf("i=%d, j=%d, k=%d, l=%d\n", i, j, k, l)
}
```

## 数组（var arrayName [size]dataType）

```go
var numbers = [5]int{1, 2, 3, 4, 5} // 确定长度
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 不确定长度
```

## 切片（Slice）

```go
// nil 空切片
var slice []type

// 通过内置函数 make 创建指定 len 初始大小和 cap 最大容量的切片
var slice []type = make([]type, len)
var slice []type = make([]type, len, cap)

// 切片初始化
slice := []int{1,2,3}

// 作为数组引用的切片
slice := arr[startIndex:endIndex]
slice := arr[startIndex:]
slice := arr[:endIndex]
slice := arr[:]

// 切片大小和容量
len(slice)
cap(slice)

// 切片扩容（相当于新建切片）
slice = append(slice, element)
slice = append(slice, element1, element2)

// 切片拷贝
srcSlice := []int {1,2,3}
destSlice := make([]int, len(srcSlice), (cap(srcSlice))*2)  // 容量翻倍
copy(destSlice, srcSlice)
```

## 集合（Map）

```go
m := make(map[keyType]ValueType, initialCapacity)

m := make(map[string]int)
m := make(map[string]int, 10)
// 使用字面量创建 Map
m := map[string]int{
    "apple": 1,
    "banana": 2,
    "orange": 3,
}
// 遍历
for k, v := range m {
    fmt.Printf("key=%s, value=%d\n", k, v)
}
// 删除元素
delete(m, "banana")
```

## 遍历（Range）

```go
for key, value := range oldMap {
    newMap[key] = value
}

for key, _ := range oldMap  // 忽略 value
for _, value := range oldMap // 忽略 key
```

## 并发（Goroutines、Channel、Scheduler）

- Goroutines: 并发线程，使用 go 关键字启动
- Channel: Goroutine 之间的通信机制， 使用 chan 关键字创建， 通过 `<-` 发送和接收数据
- Scheduler: 基于 GMP 模型，将 Goroutine 分配到系统线程中执行，通过 M 和 P 的配合高效管理并发
  - G: Goroutine
  - M: Machine 系统线程
  - P: Processor 逻辑处理器

```go
package main

import (
    "fmt"
    "time"
)

// worker 模拟一个工作协程，从通道接收任务并处理
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("worker %d 开始处理任务 %d\n", id, j)
        time.Sleep(time.Second) // 模拟耗时操作
        results <- j * 2        // 将处理结果发送到 results 通道
        fmt.Printf("worker %d 完成任务 %d\n", id, j)
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)    // 创建带缓冲的任务通道
    results := make(chan int, numJobs) // 创建带缓冲的结果通道

    // 启动 3 个 worker goroutine
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 发送 5 个任务到 jobs 通道
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // 关闭通道，通知 worker 不会再有新任务

    // 收集所有处理结果
    for a := 1; a <= numJobs; a++ {
        fmt.Printf("结果: %d\n", <-results)
    }
    close(results) // 关闭通道
}
```

## 使用 WaitGroup 控制 Goroutine

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Goroutine 完成时调用 Done()
    fmt.Printf("Worker %d started\n", id)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // 增加计数器
        go worker(i, &wg)
    }

    wg.Wait() // 等待所有 Goroutine 完成
    fmt.Println("All workers done")
}
```

## Goroutine 高级特性

- 创建有缓冲的通道（Buffered Channel）

```go
ch := make(chan int, 2)
```

- 控制线程生命周期（Context）

```go
context.WithCancel
context.WithTimeout
```

- 解决资源共享与竞争（Mutex 和 RWMutex）

```go
var mu sync.Mutex

mu.Lock()
// critical section
mu.Unlock()
```
