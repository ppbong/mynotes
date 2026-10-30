# [Rust](https://www.rust-lang.org/)

## 使用官方脚本安装（首选）

```bash
# Linux / macOS
# 自动安装 rustup 及 cargo/clippy/rust-docs/rust-std/rustc/rustfmt 组件
# 自动安装并设定默认 toochain (Ubuntu26.04 系统为 stable-x86_64-unknown-linux-gnu)
# 默认安装地址为 `~/.cargo`
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## 使用阿里云镜像源安装（可选）

官方脚本默认资源下载地址 `https://static.rust-lang.org` 速度较慢，建议使用阿里云镜像源安装 Rust。

- 配置环境变量（~/.bashrc）

```bash
export RUSTUP_UPDATE_ROOT=https://mirrors.aliyun.com/rustup/rustup
export RUSTUP_DIST_SERVER=https://mirrors.aliyun.com/rustup
```

- 安装

```bash
# 官网安装脚本
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# 或者使用阿里云镜像源安装脚本
curl --proto '=https' --tlsv1.2 -sSf https://mirrors.aliyun.com/repo/rust/rustup-init.sh | sh
```

## 验证

```bash
rustup --version  // 工具链管理器
rustc --version   // 编译器
cargo --version   // 包管理器
```

## 升级

```bash
rustup update
```

## 卸载

```bash
rustup self uninstall
```
