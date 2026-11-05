#!/data/data/com.termux/files/usr/bin/env bash
# ============================================================
# Paperplane Termux 一键启动脚本
# 用法:
#   curl -fsSL https://raw.githubusercontent.com/Maxx-EX/paperplane-e2ee/main/scripts/termux-bootstrap.sh | bash
# 或者克隆仓库后在仓库根目录:
#   bash scripts/termux-bootstrap.sh
# ============================================================
set -euo pipefail

echo "==> [1/6] 更新 Termux 源"
pkg update -y && pkg upgrade -y

echo "==> [2/6] 安装依赖 (golang/postgresql/redis/git/clang)"
pkg install -y golang git postgresql redis clang make termux-tools

export PREFIX=${PREFIX:-/data/data/com.termux/files/usr}
PGDIR=$PREFIX/var/lib/postgresql
PGLOG=$PREFIX/var/log/postgresql.log

echo "==> [3/6] 初始化 PostgreSQL"
if [ ! -d "$PGDIR" ]; then
  initdb "$PGDIR"
fi
pg_ctl -D "$PGDIR" -l "$PGLOG" -o "-p 5432" start || true
sleep 2
createuser -s paperplane 2>/dev/null || true
createdb paperplane 2>/dev/null || true

echo "==> [4/6] 启动 Redis"
redis-server --daemonize yes --port 6379

echo "==> [5/6] 唤醒 Termux (防止后台被杀)"
termux-wake-lock || true

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "==> [6/6] 启动 Paperplane 服务端 (http://127.0.0.1:8080)"
echo "------------------------------------------------------------"
echo "  健康检查: curl http://127.0.0.1:8080/healthz"
echo "  关闭: 按 Ctrl+C，然后 pg_ctl -D $PGDIR stop"
echo "------------------------------------------------------------"
exec go run ./server/cmd/server -config config/config.yaml
