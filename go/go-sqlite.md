# go-sqlite

## install module

```bash
# 初始化 Go 模块
go mod init sqlite-demo

# 安装纯 Go 实现的 SQLite 驱动
go get modernc.org/sqlite
```

## Example

```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	// 导入纯 Go 实现的 SQLite 驱动，驱动名为 "sqlite"
	_ "modernc.org/sqlite" 
)

func main() {
	// 1. 打开数据库连接（如果 example.db 不存在，会自动创建）
	db, err := sql.Open("sqlite", "example.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

    // SQLite 默认不支持高并发写入，通常建议将最大打开连接数限制为 1，以避免文件锁冲突
    db.SetMaxOpenConns(1)

	// 💡 最佳实践：开启 WAL 模式，大幅提升并发读取性能
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatal("开启 WAL 模式失败:", err)
	}

	// 2. 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER,
		email TEXT UNIQUE
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	fmt.Println("✅ 表创建成功或已存在！")

	// 3. 插入单条数据
	insertSQL := `INSERT INTO users (name, age, email) VALUES (?, ?, ?)`
	result, err := db.Exec(insertSQL, "张三", 28, "zhangsan@example.com")
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("✅ 插入单条数据成功，ID: %d\n", id)

	// 4. 批量插入（使用事务提高性能）
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("开启事务失败:", err)
	}
	
	// 预编译 SQL 语句
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal("预编译失败:", err)
	}
	defer stmt.Close()

	users := []struct {
		name  string
		age   int
		email string
	}{
		{"李四", 25, "lisi@example.com"},
		{"王五", 30, "wangwu@example.com"},
	}

	for _, u := range users {
		_, err = stmt.Exec(u.name, u.age, u.email)
		if err != nil {
			tx.Rollback() // 发生错误回滚事务
			log.Fatal("批量插入失败:", err)
		}
	}
	tx.Commit() // 提交事务
	fmt.Println("✅ 批量插入数据成功！")

	// 5. 查询多条数据
	fmt.Println("\n--- 📖 查询所有用户 ---")
	rows, err := db.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	defer rows.Close() // 务必关闭 rows

	for rows.Next() {
		var id, age int
		var name, email string
		// Scan 将列数据映射到变量中（必须传指针）
		err = rows.Scan(&id, &name, &age, &email)
		if err != nil {
			log.Fatal("数据扫描失败:", err)
		}
		fmt.Printf("ID: %d | 姓名: %-5s | 年龄: %d | 邮箱: %s\n", id, name, age, email)
	}
	// 检查迭代过程中是否有错误
	if err = rows.Err(); err != nil {
		log.Fatal("遍历结果集出错:", err)
	}

	// 6. 查询单条数据
	fmt.Println("\n--- 📖 查询单条用户 ---")
	var userName string
	var userAge int
	// QueryRow 用于确保只返回一行数据
	err = db.QueryRow("SELECT name, age FROM users WHERE id = ?", 1).Scan(&userName, &userAge)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("未找到该用户")
		} else {
			log.Fatal("查询单条失败:", err)
		}
	} else {
		fmt.Printf("ID 为 1 的用户: %s, 年龄: %d\n", userName, userAge)
	}

	// 7. 更新数据
	fmt.Println("\n--- ✏️ 更新数据 ---")
	updateSQL := `UPDATE users SET age = ? WHERE name = ?`
	res, err := db.Exec(updateSQL, 29, "张三")
	if err != nil {
		log.Fatal("更新失败:", err)
	}
	affected, _ := res.RowsAffected()
	fmt.Printf("更新了 %d 条记录\n", affected)

	// 8. 删除数据
	fmt.Println("\n--- 🗑️ 删除数据 ---")
	deleteSQL := `DELETE FROM users WHERE name = ?`
	res, err = db.Exec(deleteSQL, "王五")
	if err != nil {
		log.Fatal("删除失败:", err)
	}
	affected, _ = res.RowsAffected()
	fmt.Printf("删除了 %d 条记录\n", affected)
}
```
