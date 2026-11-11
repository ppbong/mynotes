package patterns

import (
	"fmt"
	"sync"

	"github.com/ppbong/go-demo/internal/demo"
)

// Registry 是一个并发安全的键值注册表。
type Registry struct {
	mu    sync.RWMutex
	items map[string]string
}

var (
	registryOnce sync.Once
	registryInst *Registry
)

// GetRegistry 使用 sync.Once 实现并发安全的懒加载单例。
func GetRegistry() *Registry {
	registryOnce.Do(func() {
		registryInst = &Registry{items: make(map[string]string)}
	})
	return registryInst
}

// Set 写入键值，写锁保证并发安全。
func (r *Registry) Set(key, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = value
}

// Get 读取键值，读锁允许并发读取。
func (r *Registry) Get(key string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[key]
	return v, ok
}

func singletonDemo() {
	reg := GetRegistry()
	reg.Set("lang", "go")

	v, ok := reg.Get("lang")
	demo.KV("读取单例", fmt.Sprintf("%s ok=%t", v, ok))
	demo.KV("两次获取同一实例", GetRegistry() == reg)
}
