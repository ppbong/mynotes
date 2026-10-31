# python 环境

- 系统级
- 用户级
- 项目级（虚拟环境）

## 虚拟环境

```bash
# 创建虚拟环境
python -m venv myvenv

# 激活虚拟环境
source myvenv/bin/activate

# 此时安装模块将保存到 myvenv 不影响系统全局
python -m pip install module_name

# 退出虚拟环境
deactivate
```

## 虚拟环境（允许访问系统环境安装的模块）

```bash
python -m venv myvenv --system-site-packages
```

## 用户级安装模块

```bash
python -m pip install -r requirements.txt --user
```

## 指定源镜像

```bash
python -m pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
```

## 生成项目依赖

```bash
python -m pip freeze > requirements.txt
```
