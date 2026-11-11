package stdlib

import (
	"fmt"
	"regexp"

	"github.com/ppbong/go-demo/internal/demo"
)

// Regexp 演示正则的编译、匹配、分组与替换。
func Regexp() {
	demo.Section("regexp 正则表达式")

	// 编译一次、反复使用；MustCompile 在模式非法时直接 panic。
	dateRe := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	s := "日期：2024-05-01，发布：2025-01-15"

	demo.KV("FindAllString", dateRe.FindAllString(s, -1))

	// FindStringSubmatch 返回完整匹配和各个捕获组。
	m := dateRe.FindStringSubmatch("2024-05-01")
	if len(m) > 0 {
		demo.KV("完整匹配", m[0])
		demo.KV("年 / 月 / 日", fmt.Sprintf("%s / %s / %s", m[1], m[2], m[3]))
	}

	// 命名分组让代码更易读。
	named := regexp.MustCompile(`(?P<year>\d{4})/(?P<month>\d{2})`)
	if sub := named.FindStringSubmatch("2024/05"); sub != nil {
		yearIdx := named.SubexpIndex("year")
		monthIdx := named.SubexpIndex("month")
		demo.KV("命名分组", fmt.Sprintf("year=%s month=%s", sub[yearIdx], sub[monthIdx]))
	}

	demo.KV("ReplaceAllString", dateRe.ReplaceAllString(s, "YYYY-MM-DD"))
	demo.KV("MatchString", regexp.MustCompile(`^\w+@\w+\.\w+$`).MatchString("go@example.com"))
}
