# 在 Termux 上跑 Paperplane（v0.2 修正版）

## 现实约束（先看这个）

| 组件 | Termux 原生 | 方案 |
|---|---|---|
| Go 服务端 | ✅ 纯 Go 可跑 | CGO_ENABLED=0 |
| Redis | ⚠️ 需加 `--ignore-warnings ARM64-COW-BUG` | 原生跑 |
| PostgreSQL | ❌ Bionic libc 跑不起来 | proot Ubuntu |
| eBPF/XDP | ❌ 非 root + seccomp 不可用 | 跳过 |
| Kotlin sidecar | ⚠️ 需要 JDK，可跑但没必要 | 跳过 |
| Caddy | ⚠️ 可跑但手机没必要 | 跳过 |

所以手机上跑的是**精简版**：Go 服务端 + Redis + proot PostgreSQL，只给局域网或 Cloudflare Tunnel 用。

## 一键启动

```bash
# 1. 装 Termux（从 F-Droid，不要用 Google Play 那个过期版）
pkg update -y && pkg upgrade -y

# 2. 装基础依赖
pkg install -y git proot-distro

# 3. 拉代码
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee

# 4. 一键引导（自动装 Redis、proot Ubuntu、PostgreSQL、启动 Go 服务）
bash scripts/termux-bootstrap.sh
```

脚本会做：
1. `pkg install golang redis proot-distro`
2. 启动 Redis（带 ARM64 COW 警告抑制）
3. 在 proot Ubuntu 里装 PostgreSQL 并启动
4. `CGO_ENABLED=0 go run ./server/cmd/server` 启动服务端

## 手动步骤（如果脚本中途挂了）

```bash
# Redis
redis-server --daemonize yes --ignore-warnings ARM64-COW-BUG

# PostgreSQL in proot
proot-distro install ubuntu
proot-distro login ubuntu --shared-tmp
# 进入容器后：
apt update && apt install -y postgresql
service postgresql start
su - postgres -c "createuser -s paperplane && createdb paperplane"
exit

# Go 服务端
cd paperplane-e2ee
CGO_ENABLED=0 go run ./server/cmd/server -config config/config.yaml
```

## 手机上装 APK

1. 从 GitHub Releases 下 `paperplane-vX.Y.Z.apk`。
2. 默认服务器是 `http://127.0.0.1:8080` —— 这是给模拟器用的。
3. 打开 App → 设置 → 改成你 Termux 手机的局域网地址（如 `http://192.168.1.5:8080`）。
4. 保存后重连。

> 不需要重新打包！服务器地址在 App 内可改，存在 `flutter_secure_storage`。

## 对外暴露（可选）

```bash
pkg install -y cloudflared
cloudflared tunnel --url http://127.0.0.1:8080
```

它会给你一个 `https://xxx.trycloudflare.com` 临时域名，填到 App 设置里即可。

## 为什么不在手机上跑完整防护栈

- eBPF/XDP：Android 内核有 seccomp 过滤，普通 Termux 无 CAP_BPF，加载不了。
- Kotlin sidecar：手机上就一个用户，Redis Streams 规则没意义。
- CrowdSec + Caddy：手机是内网入口，不需要 TLS 终止和 WAF。
- 完整栈在 VPS 上跑（`deploy/docker-compose.yml`），手机只当客户端。

## 常见问题

| 问题 | 解决 |
|---|---|
| `go: command not found` | `pkg install golang` |
| PostgreSQL 起不来 | 必须在 `proot-distro login ubuntu` 里跑，不要在 Termux 原生 |
| Redis 报 ARM64-COW-BUG | 加 `--ignore-warnings ARM64-COW-BUG` |
| App 连不上 | 检查 Termux 和手机是否在同一 Wi-Fi；App 设置里 IP 别写 127.0.0.1 |
| 编译 Go 报 CGO 错误 | `export CGO_ENABLED=0` |
| 手机杀后台 | Termux 里跑 `termux-wake-lock`；给 Termux 加电池白名单 |
