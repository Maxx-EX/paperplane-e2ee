#!/data/data/com.termux/files/usr/bin/env bash
# Paperplane Termux 引导脚本（v0.2 修正版）
#
# 现实约束：
#   - PostgreSQL 原生跑不起来（Bionic libc），必须放 proot Ubuntu 里
#   - eBPF/XDP 非 root 不可用，本脚本跳过
#   - Go 在 Termux 里 CGO 关闭，本项目服务端是纯 Go（net/http + gorilla/websocket + pgx 纯 Go），可直接跑
#
# 用法：
#   pkg install -y proot-distro git
#   proot-distro install ubuntu
#   bash scripts/termux-bootstrap.sh
set -euo pipefail

export PREFIX=${PREFIX:-/data/data/com.termux/files/usr}
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "==> [1/4] 安装 Termux 原生依赖（Redis + Go）"
pkg install -y git golang redis proot-distro

echo "==> [2/4] 启动 Redis（ARM64 用 --ignore-warnings 避免 COW bug）"
redis-server --daemonize yes --port 6379 --ignore-warnings ARM64-COW-BUG || \
  redis-server --daemonize yes --port 6379

echo "==> [3/4] 在 proot Ubuntu 里起 PostgreSQL"
if ! proot-distro list | grep -q "ubuntu.*installed"; then
  proot-distro install ubuntu
fi

# 在 proot Ubuntu 里一次性初始化 + 启动 postgres
proot-distro login ubuntu --shared-tmp -- bash -c '
  set -e
  apt-get update -qq
  apt-get install -y -qq postgresql
  service postgresql start
  su - postgres -c "createuser -s paperplane" || true
  su - postgres -c "createdb paperplane" || true
  echo "PostgreSQL 在 proot Ubuntu 里已起，监听 5432"
'

echo "==> [4/4] 启动 Paperplane Go 服务端（纯 Go，CGO_ENABLED=0）"
cd "$ROOT"
export CGO_ENABLED=0
export CONFIG_PATH="${CONFIG_PATH:-$ROOT/config/config.yaml}"
exec go run ./server/cmd/server -config "$CONFIG_PATH"
