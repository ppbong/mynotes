package stdlib

import (
	"bytes"
	"io"
	"strings"

	"github.com/ppbong/go-demo/internal/demo"
)

// IO 演示 io.Reader / io.Writer 这一组核心抽象。
func IO() {
	demo.Section("io.Reader / io.Writer")

	reader := strings.NewReader("hello, io")
	var sb strings.Builder
	if _, err := io.Copy(&sb, reader); err != nil {
		demo.KV("Copy 失败", err)
		return
	}
	demo.KV("io.Copy", sb.String())

	// io.ReadAll 会一次性读完全部内容，只适合小数据。
	data, err := io.ReadAll(strings.NewReader("全部读取"))
	if err != nil {
		demo.KV("ReadAll 失败", err)
		return
	}
	demo.KV("io.ReadAll", string(data))

	// LimitReader 限制最多读取的字节数，防止读取超大响应。
	limited, err := io.ReadAll(io.LimitReader(strings.NewReader("0123456789"), 4))
	if err != nil {
		demo.KV("LimitReader 失败", err)
		return
	}
	demo.KV("io.LimitReader(4)", string(limited))

	// TeeReader 在读取的同时把数据写入另一个 Writer，适合计算哈希。
	var buf bytes.Buffer
	tee := io.TeeReader(strings.NewReader("tee"), &buf)
	if _, err := io.ReadAll(tee); err != nil {
		demo.KV("TeeReader 失败", err)
		return
	}
	demo.KV("io.TeeReader 旁路写入", buf.String())
}
