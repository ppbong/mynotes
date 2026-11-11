package stdlib

import (
	"time"

	"github.com/ppbong/go-demo/internal/demo"
)

// Time 演示时间格式化、解析、运算与定时器。
func Time() {
	demo.Section("time 常用操作")

	now := time.Now()
	// Go 使用固定布局 "2006-01-02 15:04:05"，而不是 yyyy-MM-dd 占位符。
	demo.KV("当前时间", now.Format("2006-01-02 15:04:05"))
	demo.KV("RFC3339", now.Format(time.RFC3339))

	// 解析时间时显式指定时区，避免依赖服务器本地时区。
	loc := time.FixedZone("CST", 8*3600)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-05-01 09:30:00", loc)
	if err != nil {
		demo.KV("Parse 失败", err)
		return
	}
	demo.KV("解析结果", t)
	demo.KV("Unix 时间戳", t.Unix())

	// 时间运算：Add / Sub 返回新的时间与 Duration，原值不变。
	later := t.Add(90 * time.Minute)
	demo.KV("加 90 分钟", later)
	demo.KV("时间差", later.Sub(t))

	d := 2*time.Hour + 30*time.Minute
	demo.KV("Duration", d)

	// Ticker 适合周期任务，务必 Stop 释放资源。
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 2; i++ {
		<-ticker.C
	}
	demo.KV("Ticker", "已经过 2 个周期")

	// Timer 只触发一次。
	timer := time.NewTimer(10 * time.Millisecond)
	defer timer.Stop()
	<-timer.C
	demo.KV("Timer", "已触发一次")
}
