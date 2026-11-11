package patterns

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestRetry_SuccessAfterFailure 验证失败后重试成功的情况。
func TestRetry_SuccessAfterFailure(t *testing.T) {
	calls := 0
	err := Retry(context.Background(), 3, time.Millisecond, func() error {
		calls++
		if calls < 2 {
			return errors.New("临时错误")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Retry() error = %v, want nil", err)
	}
	if calls != 2 {
		t.Errorf("calls = %d, want 2", calls)
	}
}

// TestRetry_Exhausted 验证重试次数用尽后返回错误。
func TestRetry_Exhausted(t *testing.T) {
	calls := 0
	err := Retry(context.Background(), 3, time.Millisecond, func() error {
		calls++
		return errors.New("持续失败")
	})

	if err == nil {
		t.Fatal("Retry() error = nil, want non-nil")
	}
	if calls != 3 {
		t.Errorf("calls = %d, want 3", calls)
	}
}

// TestRetry_Canceled 验证 context 取消后能及时退出。
func TestRetry_Canceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := Retry(ctx, 5, 10*time.Millisecond, func() error {
		return errors.New("失败")
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Retry() error = %v, want context.Canceled", err)
	}
}

// TestMemoryUserRepo 验证仓储的保存、查询与未找到错误。
func TestMemoryUserRepo(t *testing.T) {
	repo := NewMemoryUserRepo()

	user := User{ID: "u1", Name: "Alice"}
	if err := repo.Save(user); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got, err := repo.Get("u1")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got != user {
		t.Errorf("Get() = %+v, want %+v", got, user)
	}

	if _, err := repo.Get("missing"); !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Get(missing) error = %v, want ErrUserNotFound", err)
	}
}
