# 在 Termux 上跑 Paperplane（Android 手机 / 平板）

服务端用 Go 写，不依赖内核模块，直接在 Termux 里跑。客户端 APK 从 Releases 下载。

## 一、装 Termux（必须 F-Droid 版）

1. 手机浏览器打开 https://f-droid.org/packages/com.termux/
2. 下载安装（不要用 Google Play 上那个过期版）。

## 二、一键启动（复制粘贴即可）

打开 Termux，粘贴：

```bash
pkg update -y && pkg install -y git curl
git clone https://github.com/Maxx-EX/paperplane-e2ee.git
cd paperplane-e2ee
bash scripts/termux-bootstrap.sh
```

脚本会自动：
- 装 golang / postgresql / redis / clang
- 初始化 PostgreSQL（用户 `paperplane`，库 `paperplane`）
- 启动 Redis（6379）
- 唤醒 Termux 防止后台被杀
- 跑起 Go 服务端，监听 `http://127.0.0.1:8080`

## 三、验证

另开一个 Termux 会话（左滑出抽屉 → New Session）：

```bash
curl http://127.0.0.1:8080/healthz
# {"status":"ok",...}
```

## 四、连 APK

1. 打开 Paperplane APK（已装到手机）。
2. 默认服务器地址是 `http://127.0.0.1:8080`，和 Termux 跑在同一台手机上，直接注册/登录即可。
3. 如果 Termux 和 APK 在不同设备，改服务器地址：进 App 里「设置」→ 服务器地址 → 填 Termux 手机的局域网 IP（比如 `http://192.168.1.5:8080`），重新打包 APK。

## 五、在 Termux 里自己编 APK（可选）

Termux 不能直接跑 Flutter，用 proot Ubuntu：

```bash
pkg install -y proot-distro
proot-distro install ubuntu
proot-distro login ubuntu
# 进入 Ubuntu 后：
apt update && apt install -y git curl unzip openjdk-17-jdk wget
git clone --depth 1 -b stable https://github.com/flutter/flutter.git /opt/flutter
export PATH="/opt/flutter/bin:$PATH"
flutter --version
# 装 Android SDK 命令行工具后:
cd paperplane-e2ee/client
flutter create --platforms=android .
flutter pub get
flutter build apk --release --target-platform=android-arm64
cp build/app/outputs/flutter-apk/app-release.apk /sdcard/Download/
```

## 六、常见问题

| 问题 | 解决 |
|---|---|
| `go: command not found` | `pkg install golang` |
| Postgres 起不来 | `cat $PREFIX/var/log/postgresql.log` |
| Redis 连不上 | `redis-server --daemonize yes --port 6379` |
| App 连不上 | 确认 Termux 里服务端在跑；同手机用 127.0.0.1，跨设备用局域网 IP |
| 手机杀后台 | 给 Termux 加电池白名单；脚本已跑 `termux-wake-lock` |
| 编译 APK 卡住 | `flutter config --no-analytics` 再试 |
