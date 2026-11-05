# 在 Termux 上跑 Paperplane（proot Debian 路线）

## 现实预期

Paperplane 的完整防护栈（PostgreSQL + Redis + eBPF/XDP + Kotlin sidecar + CrowdSec + Caddy）是给 Linux 服务器/Docker 设计的。在非 root 的 Android Termux 上：

| 组件 | Termux 原生 | proot Debian | 备注 |
|---|---|---|---|
| Go 服务端（纯 Go） | ✅ 可跑 | ✅ 可跑 | CGO_ENABLED=0，无 CGO 依赖 |
| Redis | ✅ 基本可用 | ✅ 可用 | ARM64 偶发 cow-bug，加 `--ignore-warnings ARM64-COW-BUG` |
| PostgreSQL | ❌ 跑不起来 | ✅ 可跑 | Bionic libc 不兼容，必须 proot |
| eBPF/XDP | ❌ 非 root 不行 | ❌ 非 root 不行 | 需要 CAP_BPF + 内核支持 |
| CrowdSec / Caddy | 跳过 | 可装但无意义 | 单机测试不需要 |

**结论**：在 Termux 上只跑「Go 服务端 + Postgres + Redis」，跳过 XDP/CrowdSec/Caddy。

## 一键脚本（推荐）

```bash
pkg update && pkg install -y git
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee
bash scripts/termux-bootstrap.sh
```

脚本会：
1. 装 `proot-distro`
2. 装 Debian
3. 在 Debian 里装 postgresql / redis / golang 并启动
4. 建 `paperplane` 用户和数据库
5. 在 Termux 里用 `CGO_ENABLED=0` 跑 Go 服务端

服务监听 `127.0.0.1:8080`，手机上的 Flutter APK 连这个地址即可。

## 手动步骤（如果你想知道脚本干了什么）

```bash
# 1. Termux 基础
pkg update -y
pkg install -y git proot-distro

# 2. 装 proot Debian
proot-distro install debian

# 3. 进 Debian，装数据库和 redis
proot-distro login debian --shared-tmp
# 以下在 Debian 里：
apt update && apt install -y postgresql redis-server golang git
service postgresql start
service redis-server start
su - postgres
psql -c "CREATE USER paperplane WITH PASSWORD 'paperplane' SUPERUSER;"
psql -c "CREATE DATABASE paperplane OWNER paperplane;"
exit
exit

# 4. 回 Termux，跑 Go 服务端
cd ~/paperplane-e2ee
cd server
CGO_ENABLED=0 go run -ldflags "-s -w" ./cmd/server -config ../config/config.termux.yaml
```

## Flutter APK 连本机服务端

APK 默认写死 `http://127.0.0.1:8080`，手机上跑服务端时这个地址正好是本机，不用改。

如果你想在另一台设备上跑服务端，需要改 `client/lib/config/app_config.dart` 里的 `baseUrl`，然后重新 `flutter build apk`。

## 在 Termux 里编译 APK（可选，需要 proot Ubuntu）

Termux 原生不能跑 Flutter SDK（缺 Android SDK 和 glibc）。用 proot Ubuntu：

```bash
pkg install -y proot-distro
proot-distro install ubuntu
proot-distro login ubuntu
# 以下在 Ubuntu 里：
apt update && apt install -y git curl unzip openjdk-17-jdk wget
git clone --depth 1 -b stable https://github.com/flutter/flutter.git /opt/flutter
export PATH="/opt/flutter/bin:$PATH"
flutter --version

# 装 Android 命令行工具
mkdir -p /opt/android-sdk/cmdline-tools
wget https://dl.google.com/android/repository/commandlinetools-linux-11076708_latest.zip
unzip commandlinetools-linux-11076708_latest.zip -d /opt/android-sdk/cmdline-tools
mv /opt/android-sdk/cmdline-tools/cmdline-tools /opt/android-sdk/cmdline-tools/latest
export ANDROID_SDK_ROOT=/opt/android-sdk
export PATH="$ANDROID_SDK_ROOT/cmdline-tools/latest/bin:$ANDROID_SDK_ROOT/platform-tools:$PATH"
yes | sdkmanager --licenses
sdkmanager "platforms;android-34" "build-tools;34.0.0"

# 编译
cd paperplane-e2ee/client
flutter pub get
flutter build apk --release --target-platform=android-arm64
cp build/app/outputs/flutter-apk/app-release.apk /sdcard/Download/
```

## 常见问题

| 问题 | 解决 |
|---|---|
| `proot-distro install` 下载慢 | 用手机热点或换镜像 |
| PostgreSQL 起不来 | 在 proot Debian 里 `service postgresql start`，不要在 Termux 原生跑 |
| Redis ARM64 报错 | 加 `--ignore-warnings ARM64-COW-BUG` |
| `go run` 编译失败 | 确认 `CGO_ENABLED=0`，pgx/go-redis 都是纯 Go |
| 手机杀后台 | 给 Termux 加电池白名单，跑 `termux-wake-lock` |
| 想让其他设备连 | 服务端绑 `0.0.0.0` 并跑 `termux-tunnel` 或 `cloudflared` 反代 |

## 结论

如果你只是想在手机上体验聊天，最省力的方式是：
- 服务端跑在任意 Linux VPS（`docker compose up -d`）
- 手机只装 APK，连 VPS 地址

Termux 路线是给「我就想在手机上跑完整服务端」的人准备的，能跑但折腾。
