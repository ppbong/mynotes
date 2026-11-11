// Package concurrency 演示 goroutine、channel、select、sync 与 context。
package concurrency

import (
	"bytes"
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// Run 依次执行并发示例。
func Run() {
	demo.Section("goroutine 与 WaitGroup")
	goroutineWaitGroup()

	demo.Section("channel：缓冲 / 无缓冲 / 关闭")
	channels()

	demo.Section("select：多路复用与超时")
	selectTimeout()

	demo.Section("worker pool")
	workerPool()

	demo.Section("sync 与 atomic")
	syncPrimitives()

	demo.Section("context 取消与超时")
	contextDemo()

	demo.Section("流水线（pipeline）")
	pipeline()
}

func goroutineWaitGroup() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// 显式传入 n，避免依赖循环变量的捕获行为。
		go func(n int) {
			defer wg.Done()
			mu.Lock() // 共享变量必须加锁
			total += n
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	demo.KV("并发求和 1..5", total)
}

func channels() {
	// 带缓冲 channel：缓冲未满时发送不会阻塞。
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 由发送方关闭；关闭后仍可读完剩余数据

	sum := 0
	for v := range ch { // range 会持续读取直到 channel 关闭
		sum += v
	}
	demo.KV("缓冲 channel 求和", sum)

	// 无缓冲 channel：发送与接收必须同时就绪，天然形成同步点。
	done := make(chan string)
	go func() {
		done <- "握手完成"
	}()
	demo.KV("无缓冲 channel", <-done)
}

func selectTimeout() {
	ch := make(chan string)
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch <- "数据到达"
	}()

	// select 同时等待多个 channel；time.After 常用来实现超时。
	select {
	case v := <-ch:
		demo.KV("select 收到", v)
	case <-time.After(200 * time.Millisecond):
		demo.KV("select", "超时")
	}

	// 带 default 的 select 是非阻塞的。
	empty := make(chan int)
	select {
	case v := <-empty:
		demo.KV("非阻塞接收", v)
	default:
		demo.KV("非阻塞接收", "通道为空，走 default")
	}
}

func workerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ { // 启动 3 个 worker
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs { // jobs 关闭后循环自动结束
				results <- job * job
			}
		}()
	}

	for i := 1; i <= 6; i++ {
		jobs <- i
	}
	close(jobs)

	// 所有 worker 退出后再关闭结果通道，避免向已关闭通道发送数据。
	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	demo.KV("平方和 1..6", sum)
}

func syncPrimitives() {
	// Mutex 保护共享计数器。
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	demo.KV("Mutex 计数", counter)

	// 简单计数用 atomic 更轻量，无需加锁。
	var hits atomic.Int64
	hits.Add(1)
	hits.Add(2)
	demo.KV("atomic.Int64", hits.Load())

	// sync.Once 保证初始化逻辑在并发下只执行一次。
	var once sync.Once
	initCount := 0
	for i := 0; i < 3; i++ {
		once.Do(func() {
			initCount++
		})
	}
	demo.KV("sync.Once 执行次数", initCount)

	// sync.Pool 复用临时对象，降低 GC 压力；池中对象随时可能被回收。
	pool := sync.Pool{
		New: func() any { return new(bytes.Buffer) },
	}
	buf := pool.Get().(*bytes.Buffer)
	buf.WriteString("复用缓冲区")
	demo.KV("sync.Pool 取出", buf.String())
	buf.Reset()
	pool.Put(buf)
}

func contextDemo() {
	// WithTimeout：到期后 Done 关闭，ctx.Err() 返回 deadline exceeded。
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel() // 必须调用，释放关联的定时器资源

	select {
	case <-ctx.Done():
		demo.KV("context 超时", ctx.Err())
	case <-time.After(time.Second):
		demo.KV("context", "不会走到这里")
	}

	// WithCancel：把 ctx 传给 goroutine，实现协作式退出，避免泄漏。
	parent, cancelParent := context.WithCancel(context.Background())
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		<-parent.Done()
	}()
	cancelParent()
	<-finished
	demo.KV("context 取消", "goroutine 已退出")

	// WithValue 只适合传递请求域元数据（trace-id、用户身份等），不要传依赖。
	type ctxKey struct{}
	vctx := context.WithValue(context.Background(), ctxKey{}, "trace-id-123")
	demo.KV("context.Value", vctx.Value(ctxKey{}))
}

func pipeline() {
	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}

	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				out <- n * n
			}
		}()
		return out
	}

	var results []int
	for v := range square(gen(1, 2, 3, 4)) {
		results = append(results, v)
	}
	demo.KV("生成 -> 平方", results)
}
