# 自托管指南 (SELF_HOSTING.md)

## 0. 前置

- 一台 Linux 服务器（推荐 2C2G），域名已解析到本机。
- Docker 24+ / docker compose 插件。
- 开放 80 / 443（Caddy 自动签 Let's Encrypt）。

## 1. 拉代码

```bash
git clone https://github.com/<your-org>/paperplane-e2ee.git
cd paperplane-e2ee
```

## 2. 改唯一配置

编辑 `config/config.yaml`：

- `server.public_url`: `https://chat.example.com`
- `server.ws_allowed_origins`: 你的域名
- `jwt.secret`: `openssl rand -hex 32`
- `crypto.pepper`: `openssl rand -hex 16`
- `database.password` / `redis.password`: 随机
- `defense.xdp.iface`: 公网网卡名（`ip -o link`）

```bash
make config-check        # 校验必填
make config              # 生成 deploy/.env / caddy.env / acquis.yaml / sidecar 配置
```

## 3. 启动

```bash
cd deploy
CROWDSEC_LAPI_KEY=$(openssl rand -hex 24) docker compose up -d
```

服务端监听 `127.0.0.1:8080`；Caddy 占 80/443；CrowdSec LAPI 在 `127.0.0.1:8081`。

## 4. 验证

```bash
curl https://chat.example.com/healthz
# {"status":"ok","uptime_s":3,"time":"..."}
```

客户端里：把 `SERVER_URL` 指向 `https://chat.example.com` 即可。

## 5. XDP 内核过滤（可选）

```bash
sudo apt install -y clang libbpf-dev linux-tools-common
cd xdp && make
sudo ./xdp_loader eth0 /var/lib/paperplane/bans.txt
```

封禁 IP 会由 Go 服务端定期写入 `bans.txt`，loader 每 5s 同步到 LPM trie。

## 6. 备份与升级

- 只需要备份 Postgres：`pg_dump paperplane > backup.sql`。
- Redis 不存任何持久数据，重启即清。
- 升级：`git pull && make config && cd deploy && docker compose up -d --build`。

## 常见问题

- **WS 连不上**：先看 Caddy 日志 `/var/log/caddy/access.log`，确认 `Upgrade` 头被转发；不要把 Origin 白名单写错。
- **注册 409**：用户名已存在；要重新注册先换用户名（当前不提供账号恢复/删除接口）。
- **客户端连不上**：浏览器跨域需把 `https://chat.example.com` 加进 `ws_allowed_origins`，Caddy 必须 HTTPS。
