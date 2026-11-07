# go error

## example

### errors.New() 不需要变量

```go
import (
    "errors"
    "fmt"
)

// 推荐全局定义
var ErrInvalidInput = errors.New("invalid input")

// 避免在函数体内使用（每次调用将分配新内存）
err := errors.New("invalid input")
err2 := errors.New("invalid input")
fmt.Println(err == err2) // false

// 判断是否是特定错误
if errors.Is(err, ErrInvalidInput) {
    fmt.Println("输入无效")
}
```

### fmt.Errorf() 需要拼接变量

```go
import "fmt"

err := fmt.Errorf("invalid input: %s", inputStr)
```

### fmt.Errorf() 需包装底层错误（%w）

```go
import (
    "fmt"
    "os"
)

file, err := os.Open("test.txt")
if err != nil {
    return fmt.Errorf("open config file failed: %w", err)
}
```

### 自定义错误（实现接口方法 error.Error() ）

```go
import "fmt"

type MyError struct {
    Code    int    // 错误码
    Message string // 错误信息
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func doSomeWork() error {
	return &MyError{
		Code:    4004,
		Message: "balance insufficient",
	}
}
```
