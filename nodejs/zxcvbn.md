# zxcvbn 技术文档

## 1. 简介

zxcvbn 是由 Dropbox 开发的开源密码强度评估库，它使用先进的模式匹配和机器学习算法来评估密码的安全性。与传统的密码强度检测（如要求特定字符组合）不同，zxcvbn 通过分析密码的实际破解难度来提供更准确的评估。

### 主要特性

- **智能评估**：基于密码的实际破解难度而非简单的字符组合要求
- **多语言支持**：支持多种语言的密码模式识别
- **详细反馈**：提供密码强度评分和改进建议
- **前端/后端兼容**：可在浏览器和 Node.js 环境中使用
- **无依赖**：独立库，不依赖其他第三方库

## 2. 安装

### Node.js 环境

```bash
npm install zxcvbn
```

### 浏览器环境

```html
<!-- 引入 zxcvbn 库 -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/zxcvbn/4.4.2/zxcvbn.js"></script>
```

## 3. 基本用法

### Node.js 环境

```javascript
// 导入 zxcvbn
const zxcvbn = require('zxcvbn');

// 评估密码强度
const result = zxcvbn('my_password123');

// 输出评估结果
console.log(result);
```

### 浏览器环境

```javascript
// 评估密码强度
const result = zxcvbn('my_password123');

// 输出评估结果
console.log(result);
```

## 4. API 参考

### zxcvbn(password, user_inputs?)

评估密码强度的主要函数。

#### 参数

- `password` (string): 要评估的密码
- `user_inputs` (array, optional): 用户相关信息数组（如用户名、邮箱等），用于提高评估准确性

#### 返回值

返回一个包含密码强度评估结果的对象，包含以下属性：

- `score` (number): 密码强度评分（0-4，0 最弱，4 最强）
- `guesses` (number): 估计的破解尝试次数
- `guesses_log10` (number): 破解尝试次数的对数（base 10）
- `crack_times_seconds` (object): 不同场景下的破解时间（秒）
  - `online_throttling_100_per_hour`: 在线攻击（100次/小时）
  - `online_no_throttling_10_per_second`: 在线攻击（无限制）
  - `offline_slow_hashing_1e4_per_second`: 离线攻击（慢哈希）
  - `offline_fast_hashing_1e10_per_second`: 离线攻击（快哈希）
- `crack_times_display` (object): 不同场景下的破解时间（人类可读格式）
- `feedback` (object): 密码改进建议
  - `warning`: 密码存在的问题
  - `suggestions`: 改进建议数组
- `sequence` (array): 密码中识别的模式序列
- `calc_time` (number): 评估耗时（毫秒）

## 5. 使用示例

### 基本示例

```javascript
const zxcvbn = require('zxcvbn');

// 评估简单密码
const result1 = zxcvbn('password123');
console.log('密码 "password123" 的评分:', result1.score); // 0
console.log('破解时间:', result1.crack_times_display.offline_fast_hashing_1e10_per_second); // "不到一秒"

// 评估强密码
const result2 = zxcvbn('C0mpl3xP@ssw0rd!');
console.log('密码 "C0mpl3xP@ssw0rd!" 的评分:', result2.score); // 4
console.log('破解时间:', result2.crack_times_display.offline_fast_hashing_1e10_per_second); // " centuries"
```

### 使用用户输入提高准确性

```javascript
const zxcvbn = require('zxcvbn');

const username = 'john.doe';
const email = 'john.doe@example.com';
const password = 'john.doe123';

// 提供用户相关信息以提高评估准确性
const result = zxcvbn(password, [username, email]);

console.log('评分:', result.score); // 0
console.log('警告:', result.feedback.warning); // "用户信息"
console.log('建议:', result.feedback.suggestions); // ["避免使用个人信息", "添加更多单词或字符"]
```

### 在 Express 应用中使用

```javascript
const express = require('express');
const zxcvbn = require('zxcvbn');
const app = express();

app.use(express.json());

app.post('/api/check-password', (req, res) => {
  const { password, username, email } = req.body;
  const result = zxcvbn(password, [username, email]);
  res.json({
    score: result.score,
    feedback: result.feedback,
    crackTime: result.crack_times_display.offline_fast_hashing_1e10_per_second
  });
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
});
```

### 在 Vue.js 中使用

```vue
<template>
  <div class="password-strength-checker">
    <input 
      v-model="password" 
      type="password" 
      placeholder="输入密码" 
      @input="checkPasswordStrength"
    />
    <div class="strength-meter">
      <div 
        class="strength-bar" 
        :style="{ width: strengthBarWidth + '%', backgroundColor: strengthBarColor }"
      ></div>
    </div>
    <div class="feedback">
      <div v-if="feedback.warning" class="warning">{{ feedback.warning }}</div>
      <ul v-if="feedback.suggestions.length">
        <li v-for="(suggestion, index) in feedback.suggestions" :key="index">
          {{ suggestion }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import zxcvbn from 'zxcvbn';

export default {
  data() {
    return {
      password: '',
      score: 0,
      feedback: {
        warning: '',
        suggestions: []
      }
    };
  },
  computed: {
    strengthBarWidth() {
      return this.score * 25;
    },
    strengthBarColor() {
      const colors = ['#ff4444', '#ffbb33', '#ffff00', '#99cc00', '#00cc88'];
      return colors[this.score];
    }
  },
  methods: {
    checkPasswordStrength() {
      if (!this.password) {
        this.score = 0;
        this.feedback = { warning: '', suggestions: [] };
        return;
      }
      
      const result = zxcvbn(this.password);
      this.score = result.score;
      this.feedback = result.feedback;
    }
  }
};
</script>

<style scoped>
.password-strength-checker {
  max-width: 400px;
  margin: 0 auto;
}

input {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  font-size: 16px;
}

.strength-meter {
  width: 100%;
  height: 10px;
  background-color: #eee;
  border-radius: 5px;
  overflow: hidden;
  margin-bottom: 10px;
}

.strength-bar {
  height: 100%;
  transition: width 0.3s ease, background-color 0.3s ease;
}

.feedback {
  font-size: 14px;
}

.warning {
  color: #ff4444;
  margin-bottom: 5px;
}

ul {
  list-style-type: disc;
  padding-left: 20px;
  color: #666;
}
</style>
```

## 6. 评分标准

zxcvbn 使用 0-4 的评分系统来评估密码强度：

| 评分 | 描述 | 破解难度 |
|------|------|----------|
| 0 | 非常弱（very weak） | 不到一秒即可破解 |
| 1 | 弱（weak） | 几分钟内可破解 |
| 2 | 中等（medium） | 几小时内可破解 |
| 3 | 强（strong） | 几天或几周内可破解 |
| 4 | 非常强（very strong） | 几年或几个世纪内可破解 |

## 7. 高级用法

### 自定义字典

```javascript
const zxcvbn = require('zxcvbn');

// 自定义字典（如公司内部术语）
const customWords = ['companyname', 'productname', 'internalsystem'];

// 评估密码，包含自定义字典
const result = zxcvbn('companyname123', customWords);

console.log(result);
```

### 处理多个密码

```javascript
const zxcvbn = require('zxcvbn');

const passwords = [
  'password123',
  'C0mpl3xP@ssw0rd!',
  '1234567890'
];

passwords.forEach((password, index) => {
  const result = zxcvbn(password);
  console.log(`密码 ${index + 1} ("${password}") 的评分: ${result.score}`);
});
```

## 8. 最佳实践

1. **实时反馈**：在用户输入密码时提供实时强度反馈
2. **结合用户信息**：使用用户的用户名、邮箱等信息提高评估准确性
3. **合理的强度要求**：根据应用的安全需求设置合理的密码强度要求（如评分 ≥ 3）
4. **清晰的反馈**：使用视觉指示器（如进度条）和文字建议帮助用户创建强密码
5. **不要过分限制**：避免过于严格的密码规则，这可能导致用户创建更差的密码

## 9. 常见问题

### Q: zxcvbn 如何评估密码强度？

A: zxcvbn 通过分析密码中的模式（如常见单词、键盘模式、重复字符等）来估计破解难度，而不是简单地检查字符组合。

### Q: zxcvbn 支持哪些语言？

A: zxcvbn 支持多种语言的密码模式识别，包括英语、西班牙语、德语、法语等。

### Q: zxcvbn 的性能如何？

A: zxcvbn 设计为高性能库，评估密码通常只需要几毫秒时间，适合实时使用。

### Q: zxcvbn 可以在移动端使用吗？

A: 是的，zxcvbn 可以在移动浏览器和 React Native 等移动开发框架中使用。

## 10. 参考资源

- [zxcvbn GitHub 仓库](https://github.com/dropbox/zxcvbn)
- [zxcvbn 官方文档](https://github.com/dropbox/zxcvbn#usage)
- [npm zxcvbn 包](https://www.npmjs.com/package/zxcvbn)
