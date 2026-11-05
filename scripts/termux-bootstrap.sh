#!/data/data/com.termux/files/usr/bin/env bash
# Paperplane Termux 启动脚本（proot Debian 路线）
# 为什么需要 proot？
#   - PostgreSQL 在 Android Bionic libc 下原生跑不起来，必须在 proot Debian/Ubuntu 里
#   - eBPF/XDP 需要 root + 内核支持，Termux 非 root 环境直接跳过
#   - Go 服务端是纯 Go（pgx/go-redis/websocket 全纯 Go），CGO_ENABLED=0 可以在 Termux 原生或 proot 里跑
set -euo pipefail

echo "==> [1/5] 安装 Termux 基础包"
pkg update -y
pkg install -y git proot-distro

echo "==> [2/5] 安装 proot Debian（如果还没装）"
if ! proot-distro list | grep -q "debian.*installed"; then
  proot-distro install debian
fi

echo "==> [3/5] 在 Debian 里装 postgresql / redis / golang"
proot-distro login debian --shared-tmp -- bash -c "
  set -e
  export DEBIAN_FRONTEND=noninteractive
  apt-get update -y
  apt-get install -y postgresql redis golang git curl
  service postgresql start || pg_ctlcluster 16 main start || true
  service redis-server start || redis-server --daemonize yes || true
  su - postgres -c \"psql -c 'CREATE USER paperplane WITH PASSWORD '\''paperplane'\'' SUPERUSER;'\" 2>/dev/null || true
  su - postgres -c \"psql -c 'CREATE DATABASE paperplane OWNER paperplane;'\" 2>/dev/null || true
"

echo "==> [4/5] 克隆仓库到 Termux home（如果还没克隆）"
cd ~
if [ ! -d paperplane-e2ee ]; then
  git clone https://github.com/Maxx-EX/paperplane-e2ee.git
fi
cd paperplane-e2ee

echo "==> [5/5] 启动 Go 服务端（CGO 纯 Go 模式）"
echo "    注意：eBPF/XDP 在 Termux 不可用，已跳过；CrowdSec 也跳过；"
echo "    服务端只绑 127.0.0.1:8080，手机上其他设备访问不到。"
echo ""

export CGO_ENABLED=0
cd server
exec go run -ldflags "-s -w" ./cmd/server -config ../config/config.termux.yaml
