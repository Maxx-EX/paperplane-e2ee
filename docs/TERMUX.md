# 在 Termux 上跑起来（Android 手机 / 平板）

Paperplane 服务端用 Go 写，不依赖内核模块，所以可以直接在 Termux 里跑。
客户端（APK）在本机用 `flutter build apk` 编译，或直接从 GitHub Releases 下载。

## 一、装 Termux

1. 从 F-Droid 装 Termux（不要用 Google Play 上那个过期版）。
2. 打开 Termux，先更新：

```bash
pkg update -y && pkg upgrade -y
```

## 二、在 Termux 里跑服务端

一条命令：

```bash
pkg install -y git golang postgresql redis clang make
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee
bash scripts/termux-bootstrap.sh
```

脚本会自动：
- 初始化 PostgreSQL 数据库目录
- 启动 `postgres` 与 `redis-server`
- 建 `paperplane` 用户与数据库
- 跑 `go run ./server/cmd/server -config config/config.yaml`

服务监听 `127.0.0.1:8080`。手机上用 `localhost:8080` 即可。

> 注意：Termux 里 PostgreSQL 不支持扩展，本项目未使用任何 PG 扩展，开箱即跑。

## 三、在 Termux 里编译 APK（可选，需要 proot）

Termux 不能直接跑 Flutter SDK（缺 Android SDK 与 glibc），用 proot Ubuntu：

```bash
pkg install -y proot-distro
proot-distro install ubuntu
proot-distro login ubuntu
```

进入 Ubuntu 后：

```bash
apt update && apt install -y git curl unzip openjdk-17-jdk wget

# 装 Flutter SDK
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
flutter build apk --release
# 产物：build/app/outputs/flutter-apk/app-release.apk
```

把 APK 拷到手机：

```bash
cp build/app/outputs/flutter-apk/app-release.apk /sdcard/Download/
```

然后在 Android 上点开安装即可。

## 四、对外暴露（可选）

Termux 里跑 Caddy 或用 `frp` / `cloudflared` 反代到公网。最简方式：

```bash
pkg install -y cloudflared
cloudflared tunnel --url http://127.0.0.1:8080
```

它会给你一个 `https://xxx.trycloudflare.com` 临时域名。
把这个域名写进 `config/config.yaml` 的 `server.public_url` 与 `ws_allowed_origins`，
重新 `make config` 即可。

## 五、常见问题

| 问题 | 解决 |
|---|---|
| `go: command not found` | `pkg install golang` |
| Postgres 起不来 | `pg_ctl -D $PREFIX/var/lib/postgresql start` 看日志 |
| Flutter 卡在 pub get | `flutter pub cache repair` |
| APK 装不上 | 允许「未知来源」安装；架构选 `--target-platform=android-arm64` |
| 手机杀后台 | 给 Termux 加电池白名单，用 `termux-wake-lock` |
