#!/data/data/com.termux/files/usr/bin/env bash
# Paperplane 手机/Termux 一键启动脚本（v0.3 极简版）
#
# 零外部依赖：不要 PostgreSQL、不要 Redis、不要 Docker、不要 proot。
# 纯内存 relay，重启即清空，给局域网 1-2 人测试用。
# 生产部署用 deploy/docker-compose.yml。
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "==> [1/3] 安装 Go"
pkg install -y golang

echo "==> [2/3] 唤醒锁（防止手机杀后台）"
termux-wake-lock 2>/dev/null || true

echo "==> [3/3] 启动 paperplane-phone（监听 :8080）"
cd "$ROOT/server"
export CGO_ENABLED=0
exec go run ./cmd/paperplane-phone -addr :8080
