#!/usr/bin/env bash
# =============================================================================
# render-config.sh — 从唯一配置源 config/config.yaml 派生所有下游配置
# -----------------------------------------------------------------------------
# 用法:  ./scripts/render-config.sh [--check]
#   --check  只校验必要字段，不写任何文件（CI 用）
#
# 依赖:  yq (v4, 可选) 或 python3 + PyYAML
# 输出:  deploy/.env
#        deploy/caddy.env
#        deploy/crowdsec/acquis.yaml
#        policy-sidecar/config/application.yaml
#        xdp/config.env
# =============================================================================
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFIG="${CONFIG_FILE:-$ROOT/config/config.yaml}"
CHECK_ONLY=0
[[ "${1:-}" == "--check" ]] && CHECK_ONLY=1

die()  { printf '\033[31m[render-config] %s\033[0m\n' "$*" >&2; exit 1; }
warn() { printf '\033[33m[render-config] %s\033[0m\n' "$*" >&2; }
info() { printf '\033[36m[render-config] %s\033[0m\n' "$*"; }

[[ -f "$CONFIG" ]] || die "找不到配置文件: $CONFIG"

if command -v yq >/dev/null 2>&1 && yq --version 2>&1 | grep -q 'yq.*4'; then
  RENDERER=yq
elif command -v python3 >/dev/null 2>&1 && python3 -c 'import yaml' >/dev/null 2>&1; then
  RENDERER=python
else
  die "需要 yq v4 或 python3 + PyYAML（pip install pyyaml）"
fi
info "配置源: $CONFIG (renderer=$RENDERER)"

# ---------------------------------------------------------------------------
# cfg <dotted.path>   读取单值（数组用逗号连接）
# cfg_list <path>     读取数组，每项一行
# ---------------------------------------------------------------------------
cfg() {
  local expr="$1"
  if [[ "$RENDERER" == yq ]]; then
    yq -r "$expr" "$CONFIG"
  else
    python3 - "$CONFIG" "$expr" <<'PY'
import sys, yaml
doc = yaml.safe_load(open(sys.argv[1], encoding="utf-8"))
cur = doc
for part in sys.argv[2].lstrip(".").split("."):
    if isinstance(cur, dict) and part in cur:
        cur = cur[part]
    else:
        cur = ""
        break
if isinstance(cur, bool):
    print("true" if cur else "false")
elif isinstance(cur, list):
    print(",".join(str(x) for x in cur))
elif cur is None:
    print("")
else:
    print(cur)
PY
  fi
}

cfg_list() {
  local path="$1"
  if [[ "$RENDERER" == yq ]]; then
    yq -r "$path[]" "$CONFIG"
  else
    python3 - "$CONFIG" "$path" <<'PY'
import sys, yaml
doc = yaml.safe_load(open(sys.argv[1], encoding="utf-8"))
cur = doc
for part in sys.argv[2].lstrip(".").split("."):
    cur = (cur or {}).get(part, [])
for item in (cur or []):
    print(item)
PY
  fi
}

require() {
  local name="$1" value="$2"
  [[ -n "$value" ]] || die "必填字段缺失: $name"
}

# yaml 双引号转义
yq_escape() { printf '%s' "${1//\\/\\\\}" | sed 's/"/\\"/g'; }

# "60s" / "1m" / "3600" -> 秒
to_seconds() {
  local v="$1"
  case "$v" in
    *s) printf '%s' "${v%s}" ;;
    *m) printf '%s' $(( ${v%m} * 60 )) ;;
    *h) printf '%s' $(( ${v%h} * 3600 )) ;;
    *)  printf '%s' "$v" ;;
  esac
}

PROJECT=$(cfg .project);                       require project "$PROJECT"
ENVIRONMENT=$(cfg .environment)
SERVER_HOST=$(cfg .server.host);               require server.host "$SERVER_HOST"
SERVER_PORT=$(cfg .server.port);               require server.port "$SERVER_PORT"
PUBLIC_URL=$(cfg .server.public_url);          require server.public_url "$PUBLIC_URL"
DB_HOST=$(cfg .database.host);                 require database.host "$DB_HOST"
DB_PORT=$(cfg .database.port)
DB_NAME=$(cfg .database.name);                 require database.name "$DB_NAME"
DB_USER=$(cfg .database.user);                 require database.user "$DB_USER"
DB_PASSWORD=$(cfg .database.password);         require database.password "$DB_PASSWORD"
REDIS_HOST=$(cfg .redis.host);                 require redis.host "$REDIS_HOST"
REDIS_PORT=$(cfg .redis.port)
REDIS_PASSWORD=$(cfg .redis.password);         require redis.password "$REDIS_PASSWORD"
JWT_SECRET=$(cfg .jwt.secret)
PEPPER=$(cfg .crypto.pepper)
LAPI_PORT=$(cfg .defense.crowdsec.lapi_port)
KOTLIN_STREAM=$(cfg .defense.kotlin.redis_stream)
KOTLIN_GROUP=$(cfg .defense.kotlin.consumer_group)
KOTLIN_WINDOW=$(to_seconds "$(cfg .defense.kotlin.window)")
WS_LIMIT=$(cfg .defense.rate_limits.ws_per_ip)
PREKEY_LIMIT=$(cfg .defense.rate_limits.prekey_per_min)
ENVELOPE_LIMIT=$(cfg .defense.rate_limits.envelope_per_min)
TARGET_LIMIT=$(cfg .defense.rate_limits.envelope_distinct_targets)
ENVELOPE_WINDOW=$(to_seconds "$(cfg .defense.rate_limits.envelope_window)")
CHALLENGE_LIMIT=$(cfg .defense.rate_limits.challenge_per_min)

[[ ${#JWT_SECRET} -ge 32 ]] || die "jwt.secret 至少 32 字节（当前 ${#JWT_SECRET}）"
[[ ${#PEPPER} -ge 16 ]]    || die "crypto.pepper 至少 16 字节（当前 ${#PEPPER}）"
[[ "$PUBLIC_URL" == https://* && "$ENVIRONMENT" == production ]] || warn "生产环境 public_url 必须是 https://"
[[ "$SERVER_HOST" == 127.0.0.1 || "$SERVER_HOST" == localhost ]] || warn "server.host 建议绑回环地址 127.0.0.1，公网流量交给 Caddy"

WS_ORIGINS=$(cfg_list .server.ws_allowed_origins | paste -sd, -)
[[ -n "$WS_ORIGINS" ]] || die "server.ws_allowed_origins 不能为空"

if grep -q 'change-me' <<<"$JWT_SECRET$PEPPER$DB_PASSWORD$REDIS_PASSWORD"; then
  if [[ "$ENVIRONMENT" == production ]]; then
    die "检测到占位密钥（change-me）。请在 config/config.yaml 里换成真实随机值后再部署。"
  fi
  warn "检测到占位密钥，仅 development 环境允许。"
fi

if [[ $CHECK_ONLY -eq 1 ]]; then
  info "校验通过：配置字段完整、长度合规。"
  exit 0
fi

mkdir -p "$ROOT/deploy" "$ROOT/deploy/crowdsec" "$ROOT/policy-sidecar/config" "$ROOT/xdp"

# ------------------------------- deploy/.env -------------------------------
cat > "$ROOT/deploy/.env" <<EOF
# 由 scripts/render-config.sh 从 config/config.yaml 生成，请勿手改。
COMPOSE_PROJECT_NAME=${PROJECT}
PUBLIC_URL=${PUBLIC_URL}
SERVER_HOST=${SERVER_HOST}
SERVER_PORT=${SERVER_PORT}
UPSTREAM=${SERVER_HOST}:${SERVER_PORT}
POSTGRES_DB=${DB_NAME}
POSTGRES_USER=${DB_USER}
POSTGRES_PASSWORD=${DB_PASSWORD}
POSTGRES_HOST=${DB_HOST}
POSTGRES_PORT=${DB_PORT}
REDIS_HOST=${REDIS_HOST}
REDIS_PORT=${REDIS_PORT}
REDIS_PASSWORD=${REDIS_PASSWORD}
WS_ALLOWED_ORIGINS=${WS_ORIGINS}
CONFIG_PATH=/etc/paperplane/config.yaml
CROWDSEC_LAPI_PORT=${LAPI_PORT}
KOTLIN_REDIS_STREAM=${KOTLIN_STREAM}
KOTLIN_CONSUMER_GROUP=${KOTLIN_GROUP}
KOTLIN_WINDOW_SECONDS=${KOTLIN_WINDOW}
WS_PER_IP=${WS_LIMIT}
PREKEY_PER_MIN=${PREKEY_LIMIT}
ENVELOPE_PER_MIN=${ENVELOPE_LIMIT}
ENVELOPE_DISTINCT_TARGETS=${TARGET_LIMIT}
ENVELOPE_WINDOW_SECONDS=${ENVELOPE_WINDOW}
CHALLENGE_PER_MIN=${CHALLENGE_LIMIT}
EOF

# ------------------------------ deploy/caddy.env ----------------------------
cat > "$ROOT/deploy/caddy.env" <<EOF
PUBLIC_URL=${PUBLIC_URL}
UPSTREAM=${SERVER_HOST}:${SERVER_PORT}
EOF

# -------------------------- crowdsec/acquis.yaml ---------------------------
cat > "$ROOT/deploy/crowdsec/acquis.yaml" <<EOF
# 由 render-config.sh 生成：Caddy 访问日志 -> CrowdSec 场景引擎
filenames:
  - /var/log/caddy/access.log
labels:
  type: caddy
EOF

# ------------------- policy-sidecar/config/application.yaml ----------------
cat > "$ROOT/policy-sidecar/config/application.yaml" <<EOF
# 由 render-config.sh 从 config/config.yaml 生成，请勿手改。
server:
  health_port: 8091
redis:
  host: "$(yq_escape "$REDIS_HOST")"
  port: ${REDIS_PORT}
  password: "$(yq_escape "$REDIS_PASSWORD")"
  database: $(cfg .redis.db)
policy:
  stream: "$(yq_escape "$KOTLIN_STREAM")"
  consumer_group: "$(yq_escape "$KOTLIN_GROUP")"
  consumer_name: "policy-1"
  window_seconds: ${KOTLIN_WINDOW}
  ip_account_ratio: 8          # 同一 IP 关联账号数超阈值 -> 封禁 IP
  reconnect_limit: 5           # 窗口内重连次数超阈值 -> 降速/封禁
  reconnect_window_seconds: 30
  ban_seconds: 600
  fail_open_timeout_ms: 50
crowdsec:
  lapi_url: "http://127.0.0.1:${LAPI_PORT}"
  api_key_env: "CROWDSEC_LAPI_KEY"
  decision_duration: "10m"
EOF

# ------------------------------ xdp/config.env ------------------------------
cat > "$ROOT/xdp/config.env" <<EOF
XDP_ENABLED=$(cfg .defense.xdp.enabled)
XDP_IFACE=$(cfg .defense.xdp.iface)
BAN_FILE=/var/lib/paperplane/bans.txt
BAN_POLL_SECONDS=5
EOF

info "已生成："
info "  deploy/.env"
info "  deploy/caddy.env"
info "  deploy/crowdsec/acquis.yaml"
info "  policy-sidecar/config/application.yaml"
info "  xdp/config.env"
