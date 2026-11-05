# 架构 (ARCHITECTURE.md)

## 组件

| 组件 | 职责 |
|---|---|
| Flutter client | 账号派生 / libsignal 会话 / UI / 本地加密存储 |
| Go server | HTTP + WS 网关、零知识信封转发、内嵌限流 |
| PostgreSQL | 账号表、预密钥包、待投信封 |
| Redis | 一次性 nonce / 票据、限流计数、事件流、封禁 ZSET |
| Kotlin sidecar | 消费 Redis Streams，跨维规则，下发封禁 |
| Caddy | TLS 终止 + WS 反代 + 访问日志 |
| CrowdSec | 场景引擎 + LAPI 决策 |
| XDP loader | 把封禁 ZSET 同步到内核 LPM trie |

## 请求流

```
Client ──HTTPS──▶ Caddy ──▶ Go:8080 ──▶ Postgres / Redis
                       ▲
Client ──WSS───────────┘
                │
                ├─ Go 发布事件 ──▶ Redis Stream
                │                       │
                │                  Kotlin sidecar
                │                       │ ban:ip ZSET
                ▼                       ▼
            CrowdSec LAPI          XDP loader → LPM map
```

## 数据流

- 注册：`username+password → PHC → 存(account_id, PHC)`
- 发消息：`明文 → Signal.encrypt → 信封 → 转发 → ack 后删除`
- 在线推送：`Hub.Deliver → WebSocket frame → 客户端 ack`

## 配置流

```
config/config.yaml
        │
        ▼ scripts/render-config.sh
   ├── deploy/.env
   ├── deploy/caddy.env
   ├── deploy/crowdsec/acquis.yaml
   ├── policy-sidecar/config/application.yaml
   └── xdp/config.env
```

任何服务不允许再自己维护第二份配置。
