# 在 Termux 上跑 Paperplane（v0.3 极简版）

## 一句话方案

手机上只跑一个**纯 Go 二进制**，零外部依赖：不要 PostgreSQL、不要 Redis、不要 Docker、不要 proot。数据存在内存里，重启即清空，适合局域网内 1-2 个人测试。

## 一键启动

```bash
# 1. 装 Termux（F-Droid 版，不要 Google Play 那个过期的）
pkg update -y && pkg upgrade -y

# 2. 装 git + go
pkg install -y git golang

# 3. 拉代码
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee

# 4. 一键启动
bash scripts/termux-bootstrap.sh
```

就这四步。脚本只做三件事：
1. `pkg install golang`
2. `termux-wake-lock`（防止手机杀后台）
3. `CGO_ENABLED=0 go run ./server/cmd/paperplane-phone -addr :8080`

## 手动启动（如果脚本中途挂了）

```bash
cd paperplane-e2ee
CGO_ENABLED=0 go run ./server/cmd/paperplane-phone -addr :8080
```

## 手机上装 APK

1. 从 GitHub Releases 下 `paperplane-vX.Y.Z.apk`。
2. **改服务器地址**：用你 Termux 手机的局域网 IP，例如 `http://192.168.1.5:8080`。
3. 改地址的方法：
   - 改源码 `client/lib/config/app_config.dart` 第 8 行的 `defaultValue: 'http://127.0.0.1:8080'`，重新打包 APK。
   - 或者用 `--dart-define=SERVER_URL=http://192.168.1.5:8080` 编译。
4. 注册 → 登录 → 发消息。

> 查看 Termux 手机的局域网 IP：`ifconfig wlan0` 或 `ip addr show wlan0`。

## 对外暴露（可选，让别人也能连）

```bash
pkg install -y cloudflared
cloudflared tunnel --url http://127.0.0.1:8080
```

它会给你一个 `https://xxx.trycloudflare.com` 临时域名，改 `app_config.dart` 重新打包即可。

## 为什么不用 PostgreSQL / Redis / Docker

| 组件 | 手机上的问题 | 解决方案 |
|---|---|---|
| PostgreSQL | Bionic libc 跑不起来，proot 太重 | 手机版用纯内存存储 |
| Redis | ARM64 COW bug，需要额外参数 | 手机版不需要 |
| Docker | Termux 里跑 Docker 要 root + proot | 手机版是单二进制 |
| eBPF/XDP | 非 root 不可用 | 手机版跳过 |
| Caddy/CrowdSec | 手机是内网入口，不需要 TLS/WAF | 手机版跳过 |

完整防护栈（PostgreSQL + Redis + Caddy + CrowdSec + Kotlin sidecar + XDP）在 VPS 上跑：
```bash
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee
# 改 config/config.yaml 里的 secret
docker compose -f deploy/docker-compose.yml up -d
```

## 改服务器地址要改哪个文件

**文件：`client/lib/config/app_config.dart`**

第 8 行：
```dart
static const String baseUrl = String.fromEnvironment(
  'SERVER_URL',
  defaultValue: 'http://127.0.0.1:8080',  // ← 改这一行
);
```

把 `http://127.0.0.1:8080` 改成你的 VPS 地址或 Termux 局域网 IP，然后 push 一个新 tag，CI 会自动出新 APK。

## 常见问题

| 问题 | 解决 |
|---|---|
| `go: command not found` | `pkg install golang` |
| App 连不上 | 检查手机和 Termux 是否在同一 Wi-Fi；App 里的地址别写 127.0.0.1 |
| 编译 Go 报 CGO 错误 | `export CGO_ENABLED=0` |
| 手机杀后台 | Termux 里跑 `termux-wake-lock`；给 Termux 加电池白名单 |
| 重启后数据没了 | 正常，手机版是内存存储；生产用 VPS + PostgreSQL |
