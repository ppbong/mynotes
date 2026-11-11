package patterns

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ppbong/go-demo/internal/demo"
)

// User 是仓储中的领域对象。
type User struct {
	ID   string
	Name string
}

// ErrUserNotFound 是仓储的哨兵错误。
var ErrUserNotFound = errors.New("用户不存在")

// UserRepository 是使用方需要的最小接口。
type UserRepository interface {
	Get(id string) (User, error)
	Save(u User) error
}

// MemoryUserRepo 是接口的内存实现，适合测试和本地演示。
type MemoryUserRepo struct {
	mu   sync.RWMutex
	data map[string]User
}

// NewMemoryUserRepo 返回具体类型，方便调用方按需使用。
func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{data: make(map[string]User)}
}

// Save 保存用户；ID 为空时返回错误。
func (r *MemoryUserRepo) Save(u User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u.ID == "" {
		return errors.New("用户 ID 不能为空")
	}
	r.data[u.ID] = u
	return nil
}

// Get 查询用户，未找到时返回包装后的 ErrUserNotFound。
func (r *MemoryUserRepo) Get(id string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.data[id]
	if !ok {
		return User{}, fmt.Errorf("Get(%q): %w", id, ErrUserNotFound)
	}
	return u, nil
}

func repositoryDemo() {
	var repo UserRepository = NewMemoryUserRepo() // 面向接口使用

	if err := repo.Save(User{ID: "u1", Name: "Alice"}); err != nil {
		demo.KV("保存失败", err)
		return
	}

	u, err := repo.Get("u1")
	demo.KV("仓储读取", fmt.Sprintf("%+v err=%v", u, err))

	_, err = repo.Get("missing")
	demo.KV("errors.Is(ErrUserNotFound)", errors.Is(err, ErrUserNotFound))
}
