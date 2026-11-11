package stdlib

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ppbong/go-demo/internal/demo"
)

// FileSystem 演示文件读写、目录遍历与错误判断。
func FileSystem() {
	demo.Section("os / bufio / filepath 文件操作")

	dir, err := os.MkdirTemp("", "go-demo-*")
	if err != nil {
		demo.KV("MkdirTemp 失败", err)
		return
	}
	defer os.RemoveAll(dir) // 函数退出时清理临时目录

	path := filepath.Join(dir, "notes.txt")

	// 权限 0600 表示仅当前用户可读写。
	if err := os.WriteFile(path, []byte("第一行\n第二行\n"), 0o600); err != nil {
		demo.KV("WriteFile 失败", err)
		return
	}

	// 小文件可以一次性读取。
	data, err := os.ReadFile(path)
	if err != nil {
		demo.KV("ReadFile 失败", err)
		return
	}
	demo.KV("ReadFile 行数", strings.Count(string(data), "\n"))

	// 大文件推荐流式读取，内存占用稳定。
	f, err := os.Open(path)
	if err != nil {
		demo.KV("Open 失败", err)
		return
	}
	defer f.Close() // 只读文件的关闭错误通常可以忽略

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		demo.KV("逐行读取", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		demo.KV("scanner 错误", err)
	}

	// 追加写入：O_APPEND 保证内容追加到文件末尾。
	f2, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		demo.KV("OpenFile 失败", err)
		return
	}
	writer := bufio.NewWriter(f2) // 缓冲写入，减少系统调用
	if _, err := writer.WriteString("第三行\n"); err != nil {
		demo.KV("写入失败", err)
	}
	if err := writer.Flush(); err != nil {
		demo.KV("Flush 失败", err)
	}
	if err := f2.Close(); err != nil {
		demo.KV("Close 失败", err)
	}

	// 遍历目录树。
	sub := filepath.Join(dir, "sub")
	_ = os.Mkdir(sub, 0o700)
	_ = os.WriteFile(filepath.Join(sub, "a.log"), []byte("log"), 0o600)

	err = filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, relErr := filepath.Rel(dir, p)
		if relErr != nil {
			return relErr
		}
		demo.KV("WalkDir", fmt.Sprintf("%s (dir=%t)", rel, d.IsDir()))
		return nil
	})
	if err != nil {
		demo.KV("WalkDir 失败", err)
	}

	// 用 errors.Is + fs.ErrNotExist 判断“文件不存在”。
	_, err = os.Stat(filepath.Join(dir, "missing.txt"))
	demo.KV("文件不存在", errors.Is(err, fs.ErrNotExist))
}
