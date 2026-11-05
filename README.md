# Paperplane

[![CI](https://github.com/Maxx-EX/paperplane-e2ee/actions/workflows/ci.yml/badge.svg)](https://github.com/Maxx-EX/paperplane-e2ee/actions/workflows/ci.yml)
[![License: AGPL-3.0](https://img.shields.io/badge/License-AGPL--3.0-blue.svg)](LICENSE)

**纸飞机风格的零知识 E2EE 加密聊天系统** — 服务端不解密、不记录密文、不存储可逆账号字段，私钥仅客户端安全存储。

- 客户端：Flutter 3.x（`flutter_chat_ui`）+ `libsignal_protocol_dart`（X3DH + Double Ratchet）
- 服务端：Go 1.25（net/http + gorilla/websocket）+ PostgreSQL + Redis
- 防护：Kotlin sidecar（协程 + Redis Streams）、C/C++ eBPF/XDP（libbpf LPM 黑名单）、CrowdSec IPS、Caddy 终止 TLS
- 许可证：**AGPL-3.0**，开源仓库

```
┌──────────┐  注册/登录  ┌──────────┐   转发密文信封   ┌────────────┐
│  Flutter │ ──────────▶ │   Go     │ ◀─────────────▶ │  PostgreSQL│
│  客户端  │             │  Relay   │                  │   envelopes│
│ (libsignal)│ ◀────────── │ (零知识) │  不存私钥/明文   └────────────┘
└──────────┘  X3DH/DR    └────┬─────┘
                              │ 事件流
                              ▼
                        ┌──────────┐   ┌──────────┐
                        │ Kotlin  │──▶│ CrowdSec │
                        │ sidecar │   └──────────┘
                        └────┬─────┘
                             │ ban:ip ZSET
                             ▼
                        ┌──────────┐
                        │ XDP/TC  │ 内核态丢包
                        │  eBPF    │
                        └──────────┘
```

## 核心设计

- **账号单向链路**：
  `account_id = SHA256(username ‖ pepper)`；
  `argon2id = Argon2id(password, salt=account_id[:16], m=64MiB, t=3, p=1)`
  服务端只存 `account_id + PHC`，登录用 `HMAC-SHA256(credential, nonce)` 挑战响应，密码零接触。
- **E2EE**：首次登录生成 IdentityKey / SignedPreKey / OneTimePreKey / KyberPreKey（占位），公钥上传，私钥入 `flutter_secure_storage`；`SessionBuilder.processPreKeyBundle` 完成 X3DH，`SessionCipher` 收发。
- **零知识**：服务端只转发密文与信封，投递后清除路由元数据，支持 Sealed Sender（盲令牌）。
- **单一配置源**：一切服务器配置收敛到 `config/config.yaml`，`make config` 派生 `.env` / Caddy / CrowdSec / Kotlin / XDP。
- **UI**：Material 3 深/浅双主题，蓝灰主色；渐变气泡 18px 圆角、单/双勾已读、文本/图片/文件输入栏。

## 快速开始

```bash
# 1. 配置
cp config/config.yaml config/config.yaml   # 改密钥（至少 32 字节）
make config-check                         # 校验字段

# 2. 本地闭环
docker compose -f docker-compose.dev.yml up -d
cd client && flutter pub get && flutter run -d chrome
# 在 Chrome 里注册两个账号即可互发密文
```

## 生产部署

```bash
./scripts/render-config.sh                 # 生成 deploy/.env 等
cd deploy && docker compose up -d
# DNS 指向本机后 Caddy 自动签 Let's Encrypt 证书
```

详见 [`docs/SELF_HOSTING.md`](docs/SELF_HOSTING.md)。

## 在 Termux 上跑（Android 手机 / 平板）

```bash
pkg install -y git golang postgresql redis
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee
bash scripts/termux-bootstrap.sh
```

完整教程（含在 proot Ubuntu 里编译 release APK）见 [`docs/TERMUX.md`](docs/TERMUX.md)。

打 tag `v*.*.*` 后，CI 会自动构建 release APK 并发布到 GitHub Releases。

## 目录结构

```
config/config.yaml         # 唯一配置源
server/                    # Go 服务端（已通过 go build）
client/                    # Flutter 客户端
policy-sidecar/            # Kotlin 策略 sidecar
xdp/                      # C/C++ eBPF/XDP 内核过滤
deploy/                   # Caddy / CrowdSec / docker-compose
docs/                     # GitHub Pages 官网 + 文档
scripts/                  # 配置渲染 / 清单生成
```

## 防御层

| 层 | 技术 | 位置 |
|---|---|---|
| 协议层 | WS Origin 精确白名单 + IP 连接限流 | Go `CheckOrigin` |
| 速率层 | Redis Lua 令牌桶 + 目标滑动窗口 | Go `defense/limits.go` |
| 策略层 | Redis Streams 事件流 + 跨维规则 | Kotlin sidecar |
| 机器层 | CrowdSec LAPI 决策 | CrowdSec container |
| 内核层 | LPM_TRIE 黑名单 + SIGTERM detach | C/C++ XDP |

## 已知坑

见 [`docs/KNOWN_ERRORS.md`](docs/KNOWN_ERRORS.md)（A–J 协议层 + K-1–K-9 内核/策略层）。

## 许可证

AGPL-3.0 © Paperplane Contributors
