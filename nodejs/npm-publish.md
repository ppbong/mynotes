# 项目发布到NPM仓库

## 1. 发布前准备工作

### 1.1 确保package.json文件配置正确

- 确保`name`字段设置为你要发布的包名，且在NPM仓库中是唯一的。
- 确保`version`字段设置为你要发布的版本号，遵循语义化版本号规范（如`1.0.0`）。
- 确保`description`字段设置为包的描述，用于在NPM仓库中显示。
- 确保`license`字段设置为包的开源协议（如`MIT`）。
- 确保`repository`字段设置为包的代码仓库地址（如`"git+https://github.com/your-username/your-package.git"`）。
- 确保`author`字段设置为包的作者信息（如`"Your Name <your.email@example.com>"`）。
- 确保`main`字段设置为入口文件的路径（如`index.js`）。
- 确保`scripts`字段包含`"publish": "npm publish"`脚本，用于发布到NPM仓库。
- 确保`keywords`字段包含包的关键词，用于在NPM仓库中搜索（如`["svg", "converter"]`）。
- 确保`files`字段包含所有需要发布的文件路径（如`["index.js", "lib/**/*"]`）。
- 确保`dependencies`字段包含所有运行时依赖的包名和版本号。
- 确保`devDependencies`字段包含所有开发时依赖的包名和版本号。
- 等

### 1.2 添加README.md文件

- 确保根目录下包含`README.md`文件，用于描述包的功能、安装方法、使用示例等。
- `README.md`文件应该包含以下内容：
  - 包的名称、描述、版本号、开源协议等基本信息。
  - 安装包的方法（如`npm install your-package`）。
  - 使用包的示例代码（如`const svgConverter = require('your-package');`）。
  - 贡献者名单（如`"Contributors": ["Your Name <your.email@example.com>"]`）。
  - 其他相关信息（如`"License": "MIT"`）。

### 1.3 添加.gitignore文件

- 确保根目录下包含`.gitignore`文件，用于指定Git应该忽略的文件和目录。
- `.gitignore`文件应该包含以下内容：
  - 编译生成的文件（如`dist/`、`build/`等）。
  - 临时文件（如`node_modules/`、`*.log`等）。
  - 配置文件（如`config.json`、`*.env`等）。
  - 其他不需要版本控制的文件（如`package-lock.json`、`yarn.lock`等）。

### 1.4 运行测试

- 确保在发布前运行测试，确保代码质量和功能正常。
- 可以使用`npm test`命令运行测试脚本（如`"test": "mocha"`）。
- 确保测试通过，没有错误或警告。

## 2. 发布到NPM仓库

### 2.1 登录NPM账号

- 确保你有一个NPM账号，并且已经登录。
- 可以使用`npm login`命令登录NPM账号（如`npm login --registry=https://registry.npmjs.org/`）。

### 2.2 发布包到NPM仓库

- 确保你在包的根目录下。
- 运行`npm publish`命令发布包到NPM仓库（如`npm publish --access=public --registry=https://registry.npmjs.org/`）。
- 确保发布成功，没有错误或警告。
- 可以在NPM仓库中搜索你的包，确认发布成功。

### 2.3 发布私有包到NPM仓库

- 确保你在包的根目录下。
- 运行`npm publish --access=private`命令发布私有包到NPM仓库（如`npm publish --access=private --registry=https://registry.npmjs.org/`）。
- 确保发布成功，没有错误或警告。
- 可以在NPM仓库中搜索你的私有包，确认发布成功。

## 3. 发布后维护

### 3.1 更新版本号

- 确保在发布新版本后，更新`package.json`文件中的版本号。
- 遵循语义化版本号规范（如`1.0.0`），根据代码变更的类型（如修复bug、添加新功能、重大变更等）更新版本号。

```bash
# 补丁版本更新（1.0.0 → 1.0.1）
npm version patch

# 小版本更新（1.0.0 → 1.1.0）
npm version minor

# 大版本更新（1.0.0 → 2.0.0）
npm version major
```

- 运行这些命令后，npm会自动更新package.json中的version字段并创建一个git标签，然后可以执行 npm publish 发布新版本。

### 3.2 发布新版本

- 确保在发布新版本前，已经更新了`package.json`文件中的版本号。
- 运行`npm publish`命令发布新版本到NPM仓库（如`npm publish --access=public --registry=https://registry.npmjs.org/`）。
- 确保发布成功，没有错误或警告。
- 可以在NPM仓库中搜索你的包，确认发布成功。

### 3.3 维护私有包

- 确保在发布私有包后，已经更新了`package.json`文件中的版本号。
- 运行`npm publish --access=private`命令发布私有包到NPM仓库（如`npm publish --access=private --registry=https://registry.npmjs.org/`）。
- 确保发布成功，没有错误或警告。
- 可以在NPM仓库中搜索你的私有包，确认发布成功。
