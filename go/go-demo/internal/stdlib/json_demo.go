package stdlib

import (
	"encoding/json"
	"fmt"

	"github.com/ppbong/go-demo/internal/demo"
)

// User 演示结构体标签与 JSON 序列化规则。
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"` // 零值时省略该字段
	Tags     []string `json:"tags"`
	password string   // 未导出字段不会参与序列化
}

// JSON 演示编码、解码以及常见注意事项。
func JSON() {
	demo.Section("encoding/json")

	u := User{ID: 1, Name: "Alice", Tags: []string{"go", "backend"}}
	u.password = "secret"

	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		demo.KV("Marshal 失败", err)
		return
	}
	fmt.Println(string(data))

	var decoded User
	if err := json.Unmarshal(data, &decoded); err != nil {
		demo.KV("Unmarshal 失败", err)
		return
	}
	demo.KV("反序列化", fmt.Sprintf("ID=%d Name=%s", decoded.ID, decoded.Name))

	// 结构不确定时，可以解码到 map[string]any。
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		demo.KV("map 解码失败", err)
		return
	}
	demo.KV("map 解码 id", raw["id"])
	demo.KV("map 解码 tags", raw["tags"])

	// 注意：Unmarshal 默认忽略 JSON 中不存在的结构体字段，
	// 需要严格校验时可使用 Decoder.DisallowUnknownFields。
}
