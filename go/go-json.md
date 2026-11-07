# go json

## import module

```go
import "encoding/json"
```

- 大文件用 `Decoder` / `Encoder`
- 小文件用 `Unmarshal` / `Marshal`

## example

### Decoder - Encoder 适用于流数据

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type AppConfig struct {
    AppName    string   `json:"app_name"`
    Version    string   `json:"version"`
    Port       int      `json:"port"`
    Features   []string `json:"features"`
    Debug      bool     `json:"debug"`
    LicenseKey string   `json:"-"` // 使用-则该字段不会被序列化/反序列化
}

// 准备测试文件
func createDummyJsonFile(filePath string) {
	jsonData := `{
		"app_name": "MyGoApp",
		"version": "1.0.0",
		"port": 8080,
		"features": ["logging", "metrics"],
		"debug": false
	}`

	os.WriteFile(filePath, []byte(jsonData), 0644)
}

func readJSON(filePath string) (AppConfig, error) {
	var config AppConfig

	file, err := os.Open(filePath)
	if err != nil {
		return config, fmt.Errorf("打开文件失败: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	
	// 未知字段
	decoder.DisallowUnknownFields() 

	err = decoder.Decode(&config)
	if err != nil {
		return config, fmt.Errorf("解析 JSON 失败: %w", err)
	}

	return config, nil
}

func writeJSON(filePath string, data AppConfig) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	
	// 设置缩进
	encoder.SetIndent("", "  ") 
	
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("编码 JSON 失败: %w", err)
	}

	return nil
}

func main() {
    filePath := "config.json"

    createDummyJsonFile(filePath)

    config, err := readJSON(filePath)
	if err != nil {
		log.Fatalf("读取 JSON 失败: %v", err)
	}

    config.Version = "2.0.0"
	config.Debug = true
	config.Features = append(config.Features, "new_feature")

	outputPath := "config_updated.json"
	err = writeJSON(outputPath, config)
	if err != nil {
		log.Fatalf("写入 JSON 失败: %v", err)
	}
```

### Unmarshal - Marshal 适用于内存数据

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Item struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	OrderID  string  `json:"order_id"`
	Customer string  `json:"customer"`
	Items    []Item  `json:"items,omitempty"` // 使用 omitempty 则 Items 为 nil 时，序列化时不输出该字段
}

func main() {
    jsonData := `{
		"order_id": "ORD-20261107-001",
		"customer": "Alice",
		"items": [
			{"product_id": "P01", "name": "机械键盘", "price": 399.0, "quantity": 1},
			{"product_id": "P02", "name": "无线鼠标", "price": 129.5, "quantity": 2}
		]
	}`

    var order Order

	err := json.Unmarshal([]byte(jsonData), &order)
	if err != nil {
		log.Fatalf("解析 JSON 失败: %v", err)
	}

	fmt.Printf("解析成功，订单 %s 包含 %d 件商品\n\n", order.OrderID, len(order.Items))

	fmt.Println("--- 开始处理商品数据 ---")
	
    // 错误做法：
	// for _, item := range order.Items {
	//     item.Quantity = 0  // 这里的 item 是副本，修改它不会影响原数组！
	// }

	// 正确做法：使用索引修改，或者使用指针切片
	for i := range order.Items {
		if order.Items[i].Name == "无线鼠标" {
			order.Items[i].Quantity += 1
			fmt.Printf("修改商品: %s, 新数量: %d\n", order.Items[i].Name, order.Items[i].Quantity)
		}
	}

	fmt.Println("\n--- 重新生成 JSON ---")
	
	// 使用 MarshalIndent 生成带缩进的 JSON
	prettyJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		log.Fatalf("序列化 JSON 失败: %v", err)
	}
	
	fmt.Println(string(prettyJSON))
}
```

### reference

```go
package main

import (
    "bytes"
    "encoding/json"
    "log"
)

type Config struct {
    Env string `json:"env"`
}

func main() {
    config := Config{Env: "production"}

    var buf bytes.Buffer

    encoder := json.NewEncoder(&buf)
    err := encoder.Encode(config)
    if err != nil {
        log.Fatal(err)
    }

    var parsedConfig Config

    jsonBytes := buf.Bytes()
    reader := bytes.NewReader(jsonBytes)
    decoder := json.NewDecoder(reader)
    err = decoder.Decode(&parsedConfig)
    if err != nil {
        log.Fatal(err)
    }
}
```
