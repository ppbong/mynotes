package patterns

import (
	"context"
	"fmt"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// Retry 以指数退避重试 fn，直到成功、次数用尽或 ctx 被取消。
// 适合处理网络抖动、临时性故障等可重试错误。
func Retry(ctx context.Context, attempts int, base time.Duration, fn func() error) error {
	if attempts <= 0 {
		return fmt.Errorf("attempts 必须大于 0")
	}

	var err error
	delay := base
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if i == attempts-1 {
			break
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("重试被取消: %w", ctx.Err())
		case <-time.After(delay):
			delay *= 2 // 指数退避，降低下游压力
		}
	}
	return fmt.Errorf("重试 %d 次仍失败: %w", attempts, err)
}

func retryDemo() {
	calls := 0
	err := Retry(context.Background(), 4, 5*time.Millisecond, func() error {
		calls++
		if calls < 3 {
			return fmt.Errorf("第 %d 次调用失败", calls)
		}
		return nil
	})
	demo.KV("重试结果", fmt.Sprintf("共调用 %d 次, err=%v", calls, err))
}
