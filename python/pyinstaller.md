# PyInstaller 技术文档

## 1. 简介

PyInstaller 是一个用于将 Python 应用程序打包成独立可执行文件的工具，支持 Windows、macOS 和 Linux 平台。它能够将 Python 解释器、脚本以及所有依赖的库和资源文件打包在一起，生成一个无需安装 Python 即可运行的可执行文件。

### 1.1 核心功能

- 将 Python 脚本转换为独立的可执行文件（.exe、.app、.elf 等）
- 支持多平台（Windows、macOS、Linux）
- 自动识别和打包依赖库
- 支持单文件和多文件模式
- 提供丰富的配置选项
- 支持自定义图标和版本信息

## 2. 安装

### 2.1 基本安装

使用 pip 安装 PyInstaller 非常简单：

```bash
pip install pyinstaller
```

### 2.2 安装特定版本

如果需要安装特定版本的 PyInstaller，可以使用以下命令：

```bash
pip install pyinstaller==5.13.2
```

### 2.3 升级 PyInstaller

```bash
pip install --upgrade pyinstaller
```

## 3. 基本使用

### 3.1 基本打包命令

最简单的打包命令是直接指定要打包的 Python 脚本：

```bash
pyinstaller your_script.py
```

这将在当前目录下创建两个文件夹：
- `build/`：包含打包过程中生成的临时文件
- `dist/`：包含最终生成的可执行文件和相关依赖

### 3.2 单文件模式

使用 `-F` 或 `--onefile` 参数可以将所有内容打包到一个单独的可执行文件中：

```bash
pyinstaller -F your_script.py
```

### 3.3 指定输出目录

使用 `-D` 或 `--distpath` 参数可以指定输出目录：

```bash
pyinstaller -D --distpath ./output your_script.py
```

### 3.4 隐藏命令行窗口

在 Windows 平台上，使用 `-w` 或 `--windowed` 参数可以隐藏命令行窗口（适用于 GUI 应用程序）：

```bash
pyinstaller -w your_gui_script.py
```

### 3.5 指定图标

使用 `-i` 或 `--icon` 参数可以指定可执行文件的图标：

```bash
pyinstaller -i your_icon.ico your_script.py
```

## 4. 高级配置

### 4.1 排除特定模块

使用 `--exclude-module` 参数可以排除不需要的模块：

```bash
pyinstaller --exclude-module matplotlib your_script.py
```

### 4.2 包含额外文件

使用 `--add-data` 参数可以包含额外的文件或文件夹：

```bash
# Windows
pyinstaller --add-data "data.txt;data" your_script.py

# macOS/Linux
pyinstaller --add-data "data.txt:data" your_script.py
```

### 4.3 包含额外二进制文件

使用 `--add-binary` 参数可以包含额外的二进制文件：

```bash
# Windows
pyinstaller --add-binary "lib.dll;lib" your_script.py

# macOS/Linux
pyinstaller --add-binary "lib.so:lib" your_script.py
```

### 4.4 自定义可执行文件名称

使用 `-n` 或 `--name` 参数可以自定义可执行文件的名称：

```bash
pyinstaller -n MyApp your_script.py
```

### 4.5 调试模式

使用 `-d` 或 `--debug` 参数可以启用调试模式：

```bash
pyinstaller -d your_script.py
```

## 5. Spec 文件详解

当使用 PyInstaller 打包时，会生成一个 `.spec` 文件，该文件包含了打包的所有配置信息。你可以手动编辑这个文件来进行更高级的配置。

### 5.1 生成 Spec 文件

使用 `-g` 或 `--specpath` 参数可以生成 Spec 文件而不执行打包：

```bash
pyinstaller --specpath ./specs -g your_script.py
```

### 5.2 基本 Spec 文件结构

一个基本的 Spec 文件结构如下：

```python
# -*- mode: python ; coding: utf-8 -*-

block_cipher = None

a = Analysis(
    ['your_script.py'],
    pathex=[],
    binaries=[],
    datas=[],
    hiddenimports=[],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    win_no_prefer_redirects=False,
    win_private_assemblies=False,
    cipher=block_cipher,
    noarchive=False,
)
pyz = PYZ(a.pure, a.zipped_data, cipher=block_cipher)

exe = EXE(
    pyz,
    a.scripts,
    a.binaries,
    a.zipfiles,
    a.datas,
    [],
    name='your_script',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    upx_exclude=[],
    runtime_tmpdir=None,
    console=True,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)
```

### 5.3 使用 Spec 文件打包

编辑完 Spec 文件后，可以使用以下命令进行打包：

```bash
pyinstaller your_script.spec
```

### 5.4 自定义 Spec 文件

#### 添加数据文件

```python
datas=[('data.txt', '.'), ('images/', 'images')],
```

#### 添加隐藏导入

```python
hiddenimports=['numpy', 'pandas'],
```

#### 排除模块

```python
excludes=['matplotlib', 'tkinter'],
```

#### 自定义图标和版本信息

```python
executable = EXE(
    pyz,
    a.scripts,
    a.binaries,
    a.zipfiles,
    a.datas,
    [],
    name='MyApp',
    icon='app_icon.ico',  # 自定义图标
    version='version_info.txt',  # 版本信息文件
    debug=False,
    console=False,  # 隐藏命令行窗口
)
```

## 6. 常见问题与解决方案

### 6.1 缺少依赖库

**问题**：打包后的程序运行时提示缺少某个模块。

**解决方案**：
1. 使用 `--hidden-import` 参数手动添加缺少的模块：
   ```bash
   pyinstaller --hidden-import=missing_module your_script.py
   ```
2. 在 Spec 文件的 `hiddenimports` 列表中添加：
   ```python
   hiddenimports=['missing_module'],
   ```

### 6.2 文件路径问题

**问题**：打包后的程序无法找到数据文件。

**解决方案**：
1. 使用 `--add-data` 参数将数据文件打包到程序中
2. 在代码中使用 `sys._MEIPASS` 来获取打包后的数据文件路径：

```python
import sys
import os

def resource_path(relative_path):
    """获取资源文件的绝对路径"""
    try:
        # PyInstaller 创建的临时目录
        base_path = sys._MEIPASS
    except Exception:
        # 正常运行时的路径
        base_path = os.path.abspath(".")
    return os.path.join(base_path, relative_path)

# 使用示例
data_file = resource_path("data.txt")
```

### 6.3 程序体积过大

**问题**：打包后的程序体积过大。

**解决方案**：
1. 使用 `-F` 参数打包成单文件时体积会更大，可以考虑使用多文件模式
2. 使用 `--exclude-module` 参数排除不需要的模块
3. 检查是否打包了不必要的依赖库
4. 考虑使用 UPX 压缩（需要先安装 UPX）：
   ```bash
   pyinstaller --upx your_script.py
   ```

### 6.4 防病毒软件误报

**问题**：打包后的程序被防病毒软件误报为病毒。

**解决方案**：
1. 升级 PyInstaller 到最新版本
2. 不使用 UPX 压缩
3. 向防病毒软件厂商提交误报报告
4. 考虑使用数字签名

## 7. 最佳实践

### 7.1 使用虚拟环境

在虚拟环境中打包可以减少不必要的依赖：

```bash
# 创建虚拟环境
python -m venv venv

# 激活虚拟环境
# Windows
venv\Scripts\activate
# macOS/Linux
source venv/bin/activate

# 安装依赖
pip install -r requirements.txt

# 安装 PyInstaller
pip install pyinstaller

# 打包
pyinstaller your_script.py
```

### 7.2 编写 requirements.txt

始终保持 `requirements.txt` 文件的更新，确保依赖版本一致：

```bash
pip freeze > requirements.txt
```

### 7.3 测试打包后的程序

在不同的环境中测试打包后的程序，确保其正常运行。

### 7.4 使用版本控制系统

将 Spec 文件添加到版本控制系统中，方便后续修改和维护。

## 8. 命令行参数参考

### 8.1 常用参数

| 参数 | 简写 | 描述 |
|------|------|------|
| `--onefile` | `-F` | 打包为单文件 |
| `--onedir` | `-D` | 打包为多文件（默认） |
| `--windowed` | `-w` | 隐藏命令行窗口（GUI 应用） |
| `--icon` | `-i` | 指定图标文件 |
| `--name` | `-n` | 自定义可执行文件名称 |
| `--specpath` | `-p` | 指定 Spec 文件生成目录 |
| `--add-data` | | 添加数据文件 |
| `--add-binary` | | 添加二进制文件 |
| `--hidden-import` | | 添加隐藏导入的模块 |
| `--exclude-module` | | 排除特定模块 |
| `--debug` | `-d` | 启用调试模式 |
| `--clean` | | 清理临时文件 |

### 8.2 所有参数

要查看所有可用参数，可以使用以下命令：

```bash
pyinstaller --help
```

## 9. 总结

PyInstaller 是一个功能强大的 Python 打包工具，能够帮助开发者将 Python 应用程序转换为独立的可执行文件。通过本文的介绍，你应该已经了解了 PyInstaller 的基本使用方法、高级配置选项以及常见问题的解决方案。

在实际使用中，建议根据项目的具体需求选择合适的打包方式和配置参数，并通过测试确保打包后的程序能够在目标平台上正常运行。
