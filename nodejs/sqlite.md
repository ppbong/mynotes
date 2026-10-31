# SQLite 技术文档

## 1. 数据库简介

SQLite 是一个轻量级的嵌入式关系型数据库管理系统，它不需要独立的服务器进程，而是直接嵌入到应用程序中。SQLite 以单一文件形式存储整个数据库，具有体积小、速度快、功能完善等特点。

### 1.1 主要特性

- **零配置**：无需安装和配置服务器
- **单文件存储**：整个数据库存储在单个文件中
- **跨平台兼容**：支持 Windows、Linux、macOS、Android、iOS 等
- **ACID 兼容**：支持原子性、一致性、隔离性和持久性
- **支持标准 SQL**：支持大部分 SQL92 标准
- **事务支持**：完整的事务管理功能
- **轻量级**：库文件体积小，占用资源少

### 1.2 适用场景

- **移动应用**：Android、iOS 等移动平台的本地数据存储
- **桌面应用**：小型桌面应用的本地数据库
- **嵌入式系统**：智能设备、IoT 设备的数据存储
- **开发测试**：开发阶段的临时数据库
- **小型网站**：访问量较小的网站后端

## 2. 安装与设置

### 2.1 直接下载 SQLite

可以从 SQLite 官方网站下载预编译的二进制文件：

- [SQLite 官方下载页面](https://www.sqlite.org/download.html)

### 2.2 在 Node.js 中使用 SQLite

在 Node.js 环境中，常用的 SQLite 驱动有两个：`sqlite3` 和 `better-sqlite3`。

#### 2.2.1 安装 sqlite3

```bash
npm install sqlite3
```

#### 2.2.2 安装 better-sqlite3

```bash
npm install better-sqlite3
```

## 3. sqlite3 库使用指南

`sqlite3` 是 Node.js 中最常用的 SQLite 驱动之一，提供了异步 API。

### 3.1 创建数据库连接

```javascript
const sqlite3 = require('sqlite3').verbose();

// 连接到数据库（不存在则创建）
const db = new sqlite3.Database('./database.db', (err) => {
  if (err) {
    console.error('连接数据库失败:', err.message);
  } else {
    console.log('成功连接到 SQLite 数据库');
  }
});

// 关闭数据库连接
db.close((err) => {
  if (err) {
    console.error('关闭数据库失败:', err.message);
  } else {
    console.log('数据库连接已关闭');
  }
});
```

### 3.2 创建表

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 创建用户表
const createTableSQL = `
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`;

db.run(createTableSQL, (err) => {
  if (err) {
    console.error('创建表失败:', err.message);
  } else {
    console.log('用户表创建成功');
  }
});

db.close();
```

### 3.3 插入数据

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 插入单个用户
const insertUserSQL = `
INSERT INTO users (name, email, password)
VALUES (?, ?, ?);
`;

db.run(insertUserSQL, ['张三', 'zhangsan@example.com', 'password123'], function(err) {
  if (err) {
    console.error('插入数据失败:', err.message);
  } else {
    console.log(`用户创建成功，ID: ${this.lastID}`);
  }
});

db.close();
```

### 3.4 查询数据

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 查询所有用户
db.all('SELECT * FROM users', [], (err, rows) => {
  if (err) {
    console.error('查询数据失败:', err.message);
  } else {
    console.log('所有用户:');
    rows.forEach((row) => {
      console.log(`${row.id} - ${row.name} (${row.email})`);
    });
  }
});

// 查询单个用户
db.get('SELECT * FROM users WHERE id = ?', [1], (err, row) => {
  if (err) {
    console.error('查询数据失败:', err.message);
  } else if (row) {
    console.log(`用户详情: ${row.id} - ${row.name} (${row.email})`);
  } else {
    console.log('未找到用户');
  }
});

db.close();
```

### 3.5 更新数据

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 更新用户信息
const updateUserSQL = `
UPDATE users
SET name = ?,
    email = ?
WHERE id = ?;
`;

db.run(updateUserSQL, ['李四', 'lisi@example.com', 1], function(err) {
  if (err) {
    console.error('更新数据失败:', err.message);
  } else {
    console.log(`更新了 ${this.changes} 条记录`);
  }
});

db.close();
```

### 3.6 删除数据

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 删除用户
const deleteUserSQL = 'DELETE FROM users WHERE id = ?';

db.run(deleteUserSQL, [1], function(err) {
  if (err) {
    console.error('删除数据失败:', err.message);
  } else {
    console.log(`删除了 ${this.changes} 条记录`);
  }
});

db.close();
```

### 3.7 事务处理

```javascript
const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('./database.db');

// 开始事务
db.serialize(() => {
  db.run('BEGIN TRANSACTION');
  
  // 执行多个操作
  const insertSQL = 'INSERT INTO users (name, email, password) VALUES (?, ?, ?)';
  
  try {
    db.run(insertSQL, ['王五', 'wangwu@example.com', 'password123']);
    db.run(insertSQL, ['赵六', 'zhaoliu@example.com', 'password456']);
    db.run(insertSQL, ['孙七', 'sunqi@example.com', 'password789']);
    
    // 提交事务
    db.run('COMMIT');
    console.log('事务执行成功');
  } catch (err) {
    // 回滚事务
    db.run('ROLLBACK');
    console.error('事务执行失败，已回滚:', err.message);
  }
});

db.close();
```

## 4. better-sqlite3 库使用指南

`better-sqlite3` 是一个高性能的 SQLite 驱动，提供了同步 API，比 `sqlite3` 更快、更易用。

### 4.1 创建数据库连接

```javascript
const Database = require('better-sqlite3');

// 连接到数据库（不存在则创建）
try {
  const db = new Database('./database.db', {
    verbose: console.log // 启用详细日志
  });
  console.log('成功连接到 SQLite 数据库');
  
  // 关闭数据库连接
  db.close();
} catch (err) {
  console.error('连接数据库失败:', err.message);
}
```

### 4.2 创建表

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 创建用户表
const createTableSQL = `
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`;

try {
  db.exec(createTableSQL);
  console.log('用户表创建成功');
} catch (err) {
  console.error('创建表失败:', err.message);
}

db.close();
```

### 4.3 插入数据

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 准备插入语句
const insertUser = db.prepare('INSERT INTO users (name, email, password) VALUES (?, ?, ?)');

try {
  // 插入单个用户
  const result = insertUser.run('张三', 'zhangsan@example.com', 'password123');
  console.log(`用户创建成功，ID: ${result.lastInsertRowid}`);
  
  // 批量插入用户
  const insertMany = db.transaction((users) => {
    for (const user of users) {
      insertUser.run(user.name, user.email, user.password);
    }
  });
  
  const users = [
    { name: '李四', email: 'lisi@example.com', password: 'password456' },
    { name: '王五', email: 'wangwu@example.com', password: 'password789' }
  ];
  
  insertMany(users);
  console.log('批量插入用户成功');
  
} catch (err) {
  console.error('插入数据失败:', err.message);
}

db.close();
```

### 4.4 查询数据

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 查询所有用户
const selectAllUsers = db.prepare('SELECT * FROM users');
const users = selectAllUsers.all();
console.log('所有用户:');
users.forEach(user => {
  console.log(`${user.id} - ${user.name} (${user.email})`);
});

// 查询单个用户
const selectUserById = db.prepare('SELECT * FROM users WHERE id = ?');
const user = selectUserById.get(1);
if (user) {
  console.log(`用户详情: ${user.id} - ${user.name} (${user.email})`);
} else {
  console.log('未找到用户');
}

// 查询特定条件的用户
const selectUsersByName = db.prepare('SELECT * FROM users WHERE name LIKE ?');
const usersWithName = selectUsersByName.all('%张%');
console.log('名字包含"张"的用户:');
usersWithName.forEach(user => {
  console.log(`${user.id} - ${user.name} (${user.email})`);
});

db.close();
```

### 4.5 更新数据

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 更新用户信息
const updateUser = db.prepare('UPDATE users SET name = ?, email = ? WHERE id = ?');

try {
  const result = updateUser.run('李四', 'lisi-new@example.com', 1);
  console.log(`更新了 ${result.changes} 条记录`);
} catch (err) {
  console.error('更新数据失败:', err.message);
}

db.close();
```

### 4.6 删除数据

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 删除用户
const deleteUser = db.prepare('DELETE FROM users WHERE id = ?');

try {
  const result = deleteUser.run(1);
  console.log(`删除了 ${result.changes} 条记录`);
} catch (err) {
  console.error('删除数据失败:', err.message);
}

db.close();
```

### 4.7 事务处理

```javascript
const Database = require('better-sqlite3');
const db = new Database('./database.db');

// 创建事务函数
const insertUsersTransaction = db.transaction((users) => {
  const insertUser = db.prepare('INSERT INTO users (name, email, password) VALUES (?, ?, ?)');
  
  for (const user of users) {
    insertUser.run(user.name, user.email, user.password);
  }
});

try {
  const users = [
    { name: '赵六', email: 'zhaoliu@example.com', password: 'password123' },
    { name: '孙七', email: 'sunqi@example.com', password: 'password456' },
    { name: '周八', email: 'zhouba@example.com', password: 'password789' }
  ];
  
  insertUsersTransaction(users);
  console.log('事务执行成功');
} catch (err) {
  console.error('事务执行失败:', err.message);
}

db.close();
```

## 5. SQLite 命令行工具

### 5.1 基本命令

```bash
# 启动 SQLite 命令行
sqlite3 database.db

# 查看所有表
.tables

# 查看表结构
.schema users

# 执行 SQL 查询
SELECT * FROM users;

# 导出数据库
.output dump.sql
.dump
.output stdout

# 导入数据库
.read dump.sql

# 退出 SQLite 命令行
.quit
```

## 6. 最佳实践

### 6.1 数据库设计

- **合理设计表结构**：根据业务需求设计合理的表结构
- **使用合适的数据类型**：选择合适的数据类型以节省空间
- **创建索引**：为常用查询字段创建索引以提高查询性能
- **设置约束**：使用主键、唯一约束、外键等保证数据完整性

### 6.2 性能优化

- **使用预编译语句**：减少 SQL 解析时间
- **批量操作**：使用事务批量处理数据
- **避免全表扫描**：使用索引和合适的查询条件
- **优化查询**：避免复杂查询和嵌套查询

### 6.3 安全注意事项

- **防止 SQL 注入**：使用参数化查询
- **加密敏感数据**：对密码等敏感数据进行加密存储
- **备份数据库**：定期备份数据库文件
- **限制文件权限**：设置合适的数据库文件权限

## 7. 常见问题与解决方案

### 7.1 数据库文件损坏

```bash
# 检查数据库完整性
sqlite3 database.db "PRAGMA integrity_check;"

# 修复数据库
sqlite3 database.db ".recover" | sqlite3 new_database.db
```

### 7.2 并发访问问题

SQLite 支持多线程读操作，但只支持单线程写操作。对于高并发场景，可以考虑：

- 使用 WAL（Write-Ahead Logging）模式
- 减少写操作频率
- 考虑使用其他数据库系统

```javascript
// 启用 WAL 模式
const db = new Database('./database.db');
db.pragma('journal_mode = WAL');
```

## 8. 参考资源

- [SQLite 官方网站](https://www.sqlite.org/)
- [SQLite 官方文档](https://www.sqlite.org/docs.html)
- [Node.js sqlite3 文档](https://github.com/TryGhost/node-sqlite3)
- [better-sqlite3 文档](https://github.com/WiseLibs/better-sqlite3)
