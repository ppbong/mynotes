# ubuntu apt mirror

## ubuntu 24.04 before ( sources.list )

- **/etc/apt/sources.list**
  - `http://archive.ubuntu.com` replace with `https://mirrors.aliyun.com`

```bash
# backup
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak 

# replace
sudo sed -i 's/http:\/\/archive.ubuntu.com/https:\/\/mirrors.aliyun.com/g' /etc/apt/sources.list
```

## ubuntu 24.04 or latest ( ubuntu.sources )

- **/etc/apt/sources.list.d/ubuntu.sources**
  - `http://cn.archive.ubuntu.com` replace with `https://mirrors.aliyun.com`
  - `http://security.ubuntu.com` replace with `https://mirrors.aliyun.com`

```bash
# backup
sudo cp /etc/apt/sources.list.d/ubuntu.sources /etc/apt/sources.list.d/ubuntu.sources.bak

# replace
sudo sed -i 's/http:\/\/cn.archive.ubuntu.com/https:\/\/mirrors.aliyun.com/g' /etc/apt/sources.list.d/ubuntu.sources
sudo sed -i 's/http:\/\/security.ubuntu.com/https:\/\/mirrors.aliyun.com/g' /etc/apt/sources.list.d/ubuntu.sources
```

## other mirrors

- `https://mirrors.tuna.tsinghua.edu.cn/ubuntu`
  - `cn.archive.ubuntu.com` default redirect to tsinghua mirrors
- `https://mirrors.aliyun.com/ubuntu`
