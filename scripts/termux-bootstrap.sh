#!/data/data/com.termux/files/usr/bin/env bash
# Paperplane 在 Termux 里直接运行（服务端）。
# 客户端仍需在 Android 上以 APK 形式安装；这里跑的是 Go 服务端 + PG/Redis。
set -euo pipefail

echo "==> 安装依赖（golang / postgresql / redis / git）"
pkg update -y
pkg install -y golang git postgresql redis clang make

export PREFIX=${PREFIX:-/data/data/com.termux/files/usr}

echo "==> 初始化 PostgreSQL"
if [ ! -d "$PREFIX/var/lib/postgresql" ]; then
  initdb "$PREFIX/var/lib/postgresql"
fi
pg_ctl -D "$PREFIX/var/lib/postgresql" -l "$PREFIX/var/log/postgresql.log" start || true
sleep 1
createuser -s paperplane 2>/dev/null || true
createdb paperplane 2>/dev/null || true

echo "==> 启动 Redis"
redis-server --daemonize yes --port 6379

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "==> 启动 Paperplane 服务端"
exec go run ./server/cmd/server -config config/config.yaml
