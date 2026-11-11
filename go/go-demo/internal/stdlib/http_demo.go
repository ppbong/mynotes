package stdlib

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// HTTPServer 演示 http 服务端与客户端的基本用法。
func HTTPServer() {
	demo.Section("net/http 客户端与服务端")

	mux := http.NewServeMux()

	// Go 1.22+ 的路由支持“方法 + 路径参数”。
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"id":   id,
			"name": "user-" + id,
		})
	})

	// httptest 启动本地测试服务器，无需外部网络即可运行示例。
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second} // 客户端一定要设置超时
	resp, err := client.Get(server.URL + "/users/7")
	if err != nil {
		demo.KV("请求失败", err)
		return
	}
	// 响应体必须关闭，否则会泄漏连接。
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		demo.KV("读取响应失败", err)
		return
	}
	demo.KV("状态码", resp.StatusCode)
	demo.KV("响应体", string(body))
}
