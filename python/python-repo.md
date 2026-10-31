# Python 镜像源设置

## 配置文件

- **Windows:** pip.ini
- **Linux/macOS:** pip.conf

## 全局设置

- 自动生成配置文件

```bash
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
pip config set global.trusted-host pypi.tuna.tsinghua.edu.cn
```

- 检查配置

```bash
pip config list
```

## 当前虚拟环境设置

- 创建虚拟环境（my-venv）

```bash
python -m venv my-venv
```

- 创建配置文件(my-venv/pip.conf)

```ini
[global]
index-url = https://pypi.tuna.tsinghua.edu.cn/simple
[install]
trusted-host = pypi.tuna.tsinghua.edu.cn
```

- 激活环境

```bash
source my-venv/bin/activate
```

## 当前终端环境设置

- Linux/macOS

```bash
export PIP_INDEX_URL=https://pypi.tuna.tsinghua.edu.cn/simple
```

- Windows

```bash
# CMD
set PIP_INDEX_URL=https://pypi.tuna.tsinghua.edu.cn/simple

# PowerShell
$env:PIP_INDEX_URL="https://pypi.tuna.tsinghua.edu.cn/simple"
```

## 临时设置

```bash
pip install requests -i https://pypi.tuna.tsinghua.edu.cn/simple
```

## 国内常用优质镜像源

- **清华大学** https://pypi.tuna.tsinghua.edu.cn/simple
- **阿里云** https://mirrors.aliyun.com/pypi/simple
- **腾讯云** https://mirrors.cloud.tencent.com/pypi/simple
- **华为云** https://repo.huaweicloud.com/repository/pypi/simple
