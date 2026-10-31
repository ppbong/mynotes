# bcryptjs 技术文档

## 1. 库简介

bcryptjs 是一个纯 JavaScript 实现的 bcrypt 密码哈希算法库，用于安全地存储用户密码。它基于 Blowfish 加密算法，通过自动加盐和可调节的计算强度提供了强大的密码保护。

### 1.1 主要特性

- **自动加盐**：内置盐生成和管理，避免手动加盐的安全风险
- **可调节的计算强度**：通过 saltRounds 参数控制加密复杂度
- **跨平台兼容**：纯 JavaScript 实现，无需依赖任何外部库
- **防彩虹表攻击**：每次哈希生成唯一的盐值
- **自适应哈希**：随着硬件性能提升，可以增加计算强度保持安全性

## 2. 安装

### 2.1 使用 npm 安装

```bash
npm install bcryptjs
```

### 2.2 使用 yarn 安装

```bash
yarn add bcryptjs
```

### 2.3 引入方式

```javascript
// CommonJS 引入方式
const bcrypt = require('bcryptjs');

// ES6 模块引入方式
import bcrypt from 'bcryptjs';
```

## 3. 基本用法

### 3.1 密码哈希

bcryptjs 提供了同步和异步两种哈希方式，推荐使用异步方式以避免阻塞事件循环。

#### 3.1.1 异步哈希（推荐）

```javascript
const bcrypt = require('bcryptjs');

// 明文密码
const password = 'mySecurePassword123';
// 盐值轮数（推荐范围：10-12）
const saltRounds = 10;

// 异步生成密码哈希
bcrypt.hash(password, saltRounds, (err, hash) => {
  if (err) {
    console.error('哈希生成失败:', err);
    return;
  }
  console.log('生成的哈希值:', hash);
  // 示例输出: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
});
```

#### 3.1.2 同步哈希

```javascript
const bcrypt = require('bcryptjs');

// 明文密码
const password = 'mySecurePassword123';
// 盐值轮数
const saltRounds = 10;

try {
  // 同步生成密码哈希
  const hash = bcrypt.hashSync(password, saltRounds);
  console.log('生成的哈希值:', hash);
} catch (err) {
  console.error('哈希生成失败:', err);
}
```

### 3.2 密码验证

密码验证用于检查用户输入的明文密码是否与存储的哈希值匹配。

#### 3.2.1 异步验证（推荐）

```javascript
const bcrypt = require('bcryptjs');

// 存储的哈希值
const storedHash = '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
// 用户输入的密码
const userInput = 'mySecurePassword123';

// 异步验证密码
bcrypt.compare(userInput, storedHash, (err, result) => {
  if (err) {
    console.error('验证失败:', err);
    return;
  }
  if (result) {
    console.log('密码正确！');
  } else {
    console.log('密码错误！');
  }
});
```

#### 3.2.2 同步验证

```javascript
const bcrypt = require('bcryptjs');

// 存储的哈希值
const storedHash = '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
// 用户输入的密码
const userInput = 'mySecurePassword123';

try {
  // 同步验证密码
  const result = bcrypt.compareSync(userInput, storedHash);
  if (result) {
    console.log('密码正确！');
  } else {
    console.log('密码错误！');
  }
} catch (err) {
  console.error('验证失败:', err);
}
```

## 4. 高级用法

### 4.1 手动生成盐值

虽然 bcryptjs 会自动生成盐值，但在某些情况下，您可能需要手动生成盐值。

```javascript
const bcrypt = require('bcryptjs');

// 异步生成盐值
const saltRounds = 10;
bcrypt.genSalt(saltRounds, (err, salt) => {
  if (err) {
    console.error('盐值生成失败:', err);
    return;
  }
  console.log('生成的盐值:', salt);
  // 使用生成的盐值进行哈希
  bcrypt.hash('myPassword', salt, (err, hash) => {
    if (err) {
      console.error('哈希生成失败:', err);
      return;
    }
    console.log('生成的哈希值:', hash);
  });
});

// 同步生成盐值
try {
  const salt = bcrypt.genSaltSync(saltRounds);
  console.log('生成的盐值:', salt);
  const hash = bcrypt.hashSync('myPassword', salt);
  console.log('生成的哈希值:', hash);
} catch (err) {
  console.error('操作失败:', err);
}
```

### 4.2 检查哈希值是否需要重新哈希

随着硬件性能提升，您可能需要增加现有密码哈希的计算强度。

```javascript
const bcrypt = require('bcryptjs');

// 存储的哈希值
const storedHash = '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
// 新的盐值轮数
const newSaltRounds = 12;

// 检查哈希值是否需要重新哈希
const isNeedRehash = bcrypt.getRounds(storedHash) < newSaltRounds;

if (isNeedRehash) {
  console.log('需要重新哈希密码');
  // 获取用户输入的密码后重新哈希
} else {
  console.log('当前哈希强度足够');
}
```

## 5. API 参考

### 5.1 哈希相关方法

#### 5.1.1 bcrypt.hash(password, saltRounds, callback)

异步生成密码哈希。

- **参数**:
  - `password`: 明文密码（字符串）
  - `saltRounds`: 盐值轮数（数字，默认：10）
  - `callback`: 回调函数，格式为 `(err, hash) => {}`

- **返回值**: 无

#### 5.1.2 bcrypt.hashSync(password, saltRounds)

同步生成密码哈希。

- **参数**:
  - `password`: 明文密码（字符串）
  - `saltRounds`: 盐值轮数（数字，默认：10）

- **返回值**: 生成的哈希值（字符串）

### 5.2 验证相关方法

#### 5.2.1 bcrypt.compare(password, hash, callback)

异步验证密码是否与哈希值匹配。

- **参数**:
  - `password`: 明文密码（字符串）
  - `hash`: 存储的哈希值（字符串）
  - `callback`: 回调函数，格式为 `(err, result) => {}`

- **返回值**: 无

#### 5.2.2 bcrypt.compareSync(password, hash)

同步验证密码是否与哈希值匹配。

- **参数**:
  - `password`: 明文密码（字符串）
  - `hash`: 存储的哈希值（字符串）

- **返回值**: 匹配结果（布尔值）

### 5.3 盐值相关方法

#### 5.3.1 bcrypt.genSalt(saltRounds, callback)

异步生成盐值。

- **参数**:
  - `saltRounds`: 盐值轮数（数字，默认：10）
  - `callback`: 回调函数，格式为 `(err, salt) => {}`

- **返回值**: 无

#### 5.3.2 bcrypt.genSaltSync(saltRounds)

同步生成盐值。

- **参数**:
  - `saltRounds`: 盐值轮数（数字，默认：10）

- **返回值**: 生成的盐值（字符串）

### 5.4 工具方法

#### 5.4.1 bcrypt.getRounds(hash)

获取哈希值使用的盐值轮数。

- **参数**:
  - `hash`: 存储的哈希值（字符串）

- **返回值**: 盐值轮数（数字）

#### 5.4.2 bcrypt.getSalt(hash)

从哈希值中提取盐值。

- **参数**:
  - `hash`: 存储的哈希值（字符串）

- **返回值**: 盐值（字符串）

## 6. 最佳实践

### 6.1 选择合适的 saltRounds

- **默认值**: 10
- **推荐范围**: 10-12
- **考虑因素**:
  - 安全性：值越大，破解难度越高
  - 性能：值越大，哈希生成时间越长
  - 硬件：根据服务器性能调整

```javascript
// 平衡安全性和性能的推荐值
const saltRounds = 10;
```

### 6.2 使用异步方法

在 Node.js 应用中，推荐使用异步方法以避免阻塞事件循环：

```javascript
// 推荐
password.hash(password, saltRounds, (err, hash) => {
  // 处理结果
});

// 不推荐（会阻塞事件循环）
const hash = bcrypt.hashSync(password, saltRounds);
```

### 6.3 存储完整的哈希值

bcrypt 生成的哈希值包含了算法标识、盐值轮数、盐值和哈希结果，因此只需存储完整的哈希值即可：

```javascript
// 存储完整的哈希值
const user = {
  username: 'user123',
  passwordHash: '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'
};
```

### 6.4 定期更新哈希强度

随着硬件性能的提升，建议定期增加 saltRounds 值以保持相同的安全级别：

```javascript
// 检查并更新哈希强度
function checkAndUpdateHash(user, plainPassword, newSaltRounds) {
  if (bcrypt.getRounds(user.passwordHash) < newSaltRounds) {
    return bcrypt.hashSync(plainPassword, newSaltRounds);
  }
  return user.passwordHash;
}
```

## 7. 常见问题

### 7.1 哈希值长度

bcrypt 生成的哈希值长度固定为 60 个字符：

```javascript
const hash = '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy';
console.log(hash.length); // 输出: 60
```

### 7.2 哈希值格式

bcrypt 哈希值格式为：`$算法$盐值轮数$盐值和哈希结果`

- `$2a$`: 算法标识（bcrypt）
- `10$`: 盐值轮数
- `N9qo8uLOickgx2ZMRZoMye`: 22 字符的盐值
- `IjZAgcfl7p92ldGxad68LJZdL17lhWy`: 31 字符的哈希结果

### 7.3 性能考量

- saltRounds = 10: 大约需要 100ms 生成哈希
- saltRounds = 12: 大约需要 400ms 生成哈希
- saltRounds = 14: 大约需要 1.6 秒生成哈希

建议根据应用的用户量和服务器性能选择合适的盐值轮数。

## 8. 与 bcrypt 的区别

| 特性 | bcrypt | bcryptjs |
|------|--------|----------|
| 实现语言 | C++ | JavaScript |
| 依赖 | 需编译依赖 | 无依赖 |
| 性能 | 更快 | 较慢 |
| 跨平台 | 需要编译不同版本 | 纯 JS，天然跨平台 |
| 安全性 | 相同 | 相同 |

## 9. 安全注意事项

- **不要限制密码长度**: bcrypt 可以处理任意长度的密码
- **不要预处理密码**: 不要对密码进行哈希或编码后再传给 bcrypt
- **使用 HTTPS**: 确保密码在传输过程中受到保护
- **定期备份**: 定期备份用户数据，包括密码哈希
- **避免明文存储**: 永远不要以明文形式存储用户密码

## 10. 参考资源

- [bcryptjs GitHub 仓库](https://github.com/dcodeIO/bcrypt.js)
- [bcrypt 算法 Wikipedia](https://zh.wikipedia.org/wiki/Bcrypt)
- [OWASP 密码存储指南](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
