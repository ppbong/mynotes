package patterns

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// gracefulShutdownDemo 演示 HTTP 服务的优雅关闭流程。
// 生产环境通常使用 signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
// 等待终止信号，再调用 Shutdown。
func gracefulShutdownDemo() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	// 端口传 0，由系统分配一个空闲端口，避免示例之间冲突。
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		demo.KV("监听失败", err)
		return
	}
	demo.KV("服务地址", ln.Addr().String())

	srv := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second, // 防止慢速请求攻击
	}

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- srv.Serve(ln)
	}()

	// Shutdown 会停止接收新请求，并等待正在处理的请求完成。
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		demo.KV("Shutdown 失败", err)
		return
	}
	demo.KV("Shutdown", "已优雅关闭")
	demo.KV("Serve 返回", <-serveErr) // 正常情况为 http.ErrServerClosed
}
