package patterns

import (
	"fmt"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// ServerConfig 演示函数选项模式：默认值安全，可选参数清晰。
type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	MaxOpenConns int
}

// Option 是函数类型的配置项，Option 本身可被测试和组合。
type Option func(*ServerConfig)

// WithAddr 设置监听地址。
func WithAddr(addr string) Option {
	return func(c *ServerConfig) { c.Addr = addr }
}

// WithReadTimeout 设置读超时。
func WithReadTimeout(d time.Duration) Option {
	return func(c *ServerConfig) { c.ReadTimeout = d }
}

// WithMaxOpenConns 设置最大连接数，非法值会被忽略。
func WithMaxOpenConns(n int) Option {
	return func(c *ServerConfig) {
		if n > 0 {
			c.MaxOpenConns = n
		}
	}
}

// NewServerConfig 先写入默认值，再依次应用选项。
func NewServerConfig(opts ...Option) *ServerConfig {
	cfg := &ServerConfig{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		MaxOpenConns: 10,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func optionsDemo() {
	cfg := NewServerConfig(WithAddr(":9090"), WithMaxOpenConns(50))
	demo.KV("配置", fmt.Sprintf("addr=%s timeout=%s conns=%d",
		cfg.Addr, cfg.ReadTimeout, cfg.MaxOpenConns))
}
