# ubuntu fcitx5 rime 中文输入法

## 基本概念

- `fcitx5`: 输入法框架
- `fcitx5-rime`: 输入法引擎
- `rime-ice`: 开源输入法实现（雾凇拼音） 

## 安装方法

### 卸载 ibus 框架（避免冲突）

```bash
sudo apt purge -y ibus
sudo apt autoremove -y
```

### 安装 fcitx5 框架及 rime 引擎

```bash
sudo apt update
sudo apt install -y fcitx5 fcitx5-rime fcitx5-config-qt 
```

### 设置 ~/.bashrc 环境变量（让系统相关组建使用 fcitx 输入法框架）

```bash
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
```

### 设置键盘输入法系统为 `Fcitx 5`

- Settings
  - System
    - Region & Language
      - Manage Installed Languages
        - Keyboard input method system : Fcitx 5

```bash
im-config -n fcitx5
```

### 安装雾凇拼音输入法和词库

```bash
export RIME_BASE=~/.local/share/fcitx5
mkdir -p $RIME_BASE 
git clone --depth 1 https://github.com/iDvel/rime-ice.git $RIME_BASE/rime
```

### 配置并激活雾凇拼音输入法

- Reload fcitx config & Activate input method

```bash
# 执行后应能显示键盘图标
fcitx5-remote -ro
```

- `Input Method Settings` add `Rime` and then chocie `雾凇拼音` in Rime Option

```bash
# 执行后应能显示配置界面，可点击键盘图标打开
fcitx5-config-qt 
```

## 其他说明

- 输入法切换 `Ctrl + Space` 或 `Shift`
- 输入法启动 `fcitx5 -dr`
- 安装过程遇到问题可尝试重启系统或重新登录或可解决
 