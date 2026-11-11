package patterns

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// Middleware 包装 http.Handler，用于插入日志、鉴权、恢复等横切逻辑。
type Middleware func(http.Handler) http.Handler

// Chain 把多个中间件串联起来；第一个中间件最先执行。
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

// Logging 记录每个请求的方法、路径和耗时。
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("  [middleware] %s %s 耗时 %s\n", r.Method, r.URL.Path, time.Since(start))
	})
}

// Recovery 捕获处理过程中的 panic，返回 500，避免整个服务崩溃。
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("  [middleware] recovered: %v", rec)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func middlewareDemo() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("boom")
	})

	handler := Chain(mux, Logging, Recovery)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/ping")
	if err != nil {
		demo.KV("请求失败", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	demo.KV("/ping", string(body))

	resp2, err := http.Get(srv.URL + "/panic")
	if err != nil {
		demo.KV("请求失败", err)
		return
	}
	_ = resp2.Body.Close()
	demo.KV("/panic 状态码", resp2.StatusCode)
}
