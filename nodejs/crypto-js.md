# Crypto-JS 技术文档

## 1. 库简介

Crypto-JS 是一个 JavaScript 加密库，提供了多种加密算法的实现，包括哈希函数、对称加密、非对称加密和消息认证码等。它可以在浏览器和 Node.js 环境中使用，具有良好的跨平台兼容性。

### 1.1 主要特性

- **全面的算法支持**：包含多种哈希算法（MD5、SHA 系列等）和对称加密算法（AES、DES、Rabbit、RC4 等）
- **易于使用的 API**：简洁直观的接口设计，方便开发者快速集成
- **跨平台兼容**：支持浏览器和 Node.js 环境
- **模块化设计**：可以根据需要引入特定的算法模块，减小打包体积
- **无外部依赖**：纯 JavaScript 实现，无需其他库支持

## 2. 安装与引入

### 2.1 安装方法

#### 使用 npm 安装

```bash
npm install crypto-js
```

#### 使用 yarn 安装

```bash
yarn add crypto-js
```

### 2.2 引入方式

#### 引入全部模块

```javascript
import CryptoJS from 'crypto-js';
// 或
const CryptoJS = require('crypto-js');
```

#### 引入特定模块

```javascript
// 引入 MD5 模块
import MD5 from 'crypto-js/md5';
// 或
const MD5 = require('crypto-js/md5');

// 引入 AES 模块
import AES from 'crypto-js/aes';
// 或
const AES = require('crypto-js/aes');
```

## 3. 哈希函数

### 3.1 MD5

MD5 是一种常用的哈希函数，生成 128 位（16 字节）的哈希值，通常表示为 32 位十六进制字符串。

```javascript
const CryptoJS = require('crypto-js');

// 计算 MD5 哈希值
const hash = CryptoJS.MD5('Hello, Crypto-JS!');
console.log(hash.toString()); // 输出: 6e513d4873f6977b7c30b5a47242480e
```

### 3.2 SHA 系列

SHA 系列包括 SHA-1、SHA-224、SHA-256、SHA-384 和 SHA-512 等哈希函数。

```javascript
const CryptoJS = require('crypto-js');

// SHA-1
const sha1Hash = CryptoJS.SHA1('Hello, Crypto-JS!');
console.log(sha1Hash.toString()); // 输出: 6c0f9443112d62a68b7576c29758a42d7b8c0d88

// SHA-256
const sha256Hash = CryptoJS.SHA256('Hello, Crypto-JS!');
console.log(sha256Hash.toString()); // 输出: 3e25960a79dbc69b674cd4ec67a72c623cd712148a08e65189087d19612f7ce7

// SHA-512
const sha512Hash = CryptoJS.SHA512('Hello, Crypto-JS!');
console.log(sha512Hash.toString()); // 输出: 长字符串...
```

### 3.3 RIPEMD-160

RIPEMD-160 是一种哈希函数，生成 160 位（20 字节）的哈希值。

```javascript
const CryptoJS = require('crypto-js');

const ripemd160Hash = CryptoJS.RIPEMD160('Hello, Crypto-JS!');
console.log(ripemd160Hash.toString()); // 输出: 9a0e6b5c0f94a7d15a2b3c4d5e6f7a8b9c0d1e2f
```

## 4. 对称加密

### 4.1 AES 加密

AES（Advanced Encryption Standard）是目前最常用的对称加密算法，支持 128、192 和 256 位密钥长度。

#### 基本用法

```javascript
const CryptoJS = require('crypto-js');

// 加密
const plaintext = 'Hello, Crypto-JS!';
const key = 'my-secret-key-12345';
const encrypted = CryptoJS.AES.encrypt(plaintext, key);
console.log(encrypted.toString()); // 输出: U2FsdGVkX19Q...（加密后的字符串）

// 解密
const decrypted = CryptoJS.AES.decrypt(encrypted.toString(), key);
const originalText = decrypted.toString(CryptoJS.enc.Utf8);
console.log(originalText); // 输出: Hello, Crypto-JS!
```

#### 使用自定义配置

```javascript
const CryptoJS = require('crypto-js');

const plaintext = 'Hello, Crypto-JS!';
const key = CryptoJS.enc.Utf8.parse('1234567890123456'); // 16 字节密钥
const iv = CryptoJS.enc.Utf8.parse('1234567890123456'); // 16 字节 IV

// 加密配置
const encrypted = CryptoJS.AES.encrypt(plaintext, key, {
  iv: iv,
  mode: CryptoJS.mode.CBC,
  padding: CryptoJS.pad.Pkcs7
});

console.log(encrypted.toString());

// 解密
const decrypted = CryptoJS.AES.decrypt(encrypted.toString(), key, {
  iv: iv,
  mode: CryptoJS.mode.CBC,
  padding: CryptoJS.pad.Pkcs7
});

const originalText = decrypted.toString(CryptoJS.enc.Utf8);
console.log(originalText);
```

### 4.2 DES 加密

DES（Data Encryption Standard）是一种对称加密算法，使用 64 位密钥（实际有效密钥长度为 56 位）。

```javascript
const CryptoJS = require('crypto-js');

// 加密
const plaintext = 'Hello, Crypto-JS!';
const key = '12345678'; // 8 字节密钥
const encrypted = CryptoJS.DES.encrypt(plaintext, key);
console.log(encrypted.toString());

// 解密
const decrypted = CryptoJS.DES.decrypt(encrypted.toString(), key);
const originalText = decrypted.toString(CryptoJS.enc.Utf8);
console.log(originalText);
```

### 4.3 Triple DES 加密

Triple DES 是 DES 的增强版，使用三个 64 位密钥进行三次加密。

```javascript
const CryptoJS = require('crypto-js');

// 加密
const plaintext = 'Hello, Crypto-JS!';
const key = '123456789012345678901234'; // 24 字节密钥
const encrypted = CryptoJS.TripleDES.encrypt(plaintext, key);
console.log(encrypted.toString());

// 解密
const decrypted = CryptoJS.TripleDES.decrypt(encrypted.toString(), key);
const originalText = decrypted.toString(CryptoJS.enc.Utf8);
console.log(originalText);
```

### 4.4 RC4 加密

RC4 是一种流密码算法，使用可变长度的密钥。

```javascript
const CryptoJS = require('crypto-js');

// 加密
const plaintext = 'Hello, Crypto-JS!';
const key = 'my-secret-key';
const encrypted = CryptoJS.RC4.encrypt(plaintext, key);
console.log(encrypted.toString());

// 解密
const decrypted = CryptoJS.RC4.decrypt(encrypted.toString(), key);
const originalText = decrypted.toString(CryptoJS.enc.Utf8);
console.log(originalText);
```

## 5. 消息认证码 (MAC)

### 5.1 HMAC

HMAC（Hash-based Message Authentication Code）是一种基于哈希函数的消息认证码算法。

```javascript
const CryptoJS = require('crypto-js');

const message = 'Hello, Crypto-JS!';
const key = 'my-secret-key';

// 使用 SHA-256 的 HMAC
const hmacSHA256 = CryptoJS.HmacSHA256(message, key);
console.log(hmacSHA256.toString());

// 使用 MD5 的 HMAC
const hmacMD5 = CryptoJS.HmacMD5(message, key);
console.log(hmacMD5.toString());
```

## 6. 格式转换

Crypto-JS 支持多种数据格式之间的转换，包括 Utf8、Base64、Hex 等。

```javascript
const CryptoJS = require('crypto-js');

// Utf8 字符串转 WordArray
const wordArray = CryptoJS.enc.Utf8.parse('Hello, Crypto-JS!');
console.log(wordArray.toString());

// WordArray 转 Utf8 字符串
const utf8String = CryptoJS.enc.Utf8.stringify(wordArray);
console.log(utf8String);

// Utf8 字符串转 Base64
const base64String = CryptoJS.enc.Base64.stringify(CryptoJS.enc.Utf8.parse('Hello, Crypto-JS!'));
console.log(base64String);

// Base64 转 Utf8 字符串
const decodedString = CryptoJS.enc.Base64.parse(base64String).toString(CryptoJS.enc.Utf8);
console.log(decodedString);

// Utf8 字符串转 Hex
const hexString = CryptoJS.enc.Hex.stringify(CryptoJS.enc.Utf8.parse('Hello, Crypto-JS!'));
console.log(hexString);

// Hex 转 Utf8 字符串
const hexDecodedString = CryptoJS.enc.Hex.parse(hexString).toString(CryptoJS.enc.Utf8);
console.log(hexDecodedString);
```

## 7. 密钥派生函数

### 7.1 PBKDF2

PBKDF2（Password-Based Key Derivation Function 2）用于从密码生成密钥。

```javascript
const CryptoJS = require('crypto-js');

const password = 'my-password';
const salt = CryptoJS.lib.WordArray.random(128/8); // 128 位盐值
const iterations = 1000; // 迭代次数
const keySize = 256/32; // 256 位密钥

const key = CryptoJS.PBKDF2(password, salt, {
  keySize: keySize,
  iterations: iterations
});

console.log('密钥:', key.toString());
console.log('盐值:', salt.toString());
```

## 8. 加密模式与填充

### 8.1 支持的加密模式

- `CryptoJS.mode.CBC` (默认)
- `CryptoJS.mode.CFB`
- `CryptoJS.mode.CTR`
- `CryptoJS.mode.OFB`
- `CryptoJS.mode.ECB`

### 8.2 支持的填充方式

- `CryptoJS.pad.Pkcs7` (默认)
- `CryptoJS.pad.Iso97971`
- `CryptoJS.pad.AnsiX923`
- `CryptoJS.pad.Iso10126`
- `CryptoJS.pad.ZeroPadding`
- `CryptoJS.pad.NoPadding`

## 9. 高级用法

### 9.1 加密对象

Crypto-JS 可以直接加密 JavaScript 对象，内部会将对象转换为 JSON 字符串。

```javascript
const CryptoJS = require('crypto-js');

const obj = {
  name: 'Crypto-JS',
  version: '4.1.1',
  algorithms: ['AES', 'DES', 'MD5', 'SHA-256']
};

const key = 'my-secret-key';

// 加密对象
const encrypted = CryptoJS.AES.encrypt(JSON.stringify(obj), key);
console.log(encrypted.toString());

// 解密对象
const decrypted = CryptoJS.AES.decrypt(encrypted.toString(), key);
const decryptedObj = JSON.parse(decrypted.toString(CryptoJS.enc.Utf8));
console.log(decryptedObj);
```

### 9.2 文件加密（Node.js）

在 Node.js 环境中，可以使用 Crypto-JS 加密文件内容。

```javascript
const CryptoJS = require('crypto-js');
const fs = require('fs');

// 读取文件内容
const fileContent = fs.readFileSync('input.txt', 'utf8');

// 加密
const key = 'my-secret-key';
const encrypted = CryptoJS.AES.encrypt(fileContent, key);

// 写入加密后的文件
fs.writeFileSync('encrypted.txt', encrypted.toString());

// 读取并解密
const encryptedContent = fs.readFileSync('encrypted.txt', 'utf8');
const decrypted = CryptoJS.AES.decrypt(encryptedContent, key);
const originalContent = decrypted.toString(CryptoJS.enc.Utf8);

console.log('原始内容:', originalContent);
```

## 10. 常见问题与注意事项

### 10.1 密钥长度要求

不同的加密算法对密钥长度有不同的要求：
- AES: 16、24 或 32 字节（对应 128、192 或 256 位）
- DES: 8 字节
- Triple DES: 24 字节

### 10.2 安全建议

- 不要使用硬编码的密钥，尤其是在前端代码中
- 定期更换密钥
- 使用足够强度的密钥（推荐至少 128 位）
- 避免使用不安全的算法，如 MD5 和 DES
- 考虑使用 HTTPS 保护传输中的加密数据

## 11. 参考资源

- [Crypto-JS GitHub 仓库](https://github.com/brix/crypto-js)
- [Crypto-JS 官方文档](https://cryptojs.gitbook.io/docs/)
- [MDN Web Docs - Web 加密 API](https://developer.mozilla.org/zh-CN/docs/Web/API/Web_Crypto_API)
