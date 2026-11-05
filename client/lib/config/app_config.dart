// 服务端地址：换服务器时只改这一行，重新打包 APK 即可。
// 默认连 Termux 里的本地服务端（手机自己跑自己）。
// 要连 VPS，把下面的地址改成你的 VPS IP，例如 'http://1.2.3.4:8080'。
class AppConfig {
  static const String baseUrl = String.fromEnvironment(
    'SERVER_URL',
    defaultValue: 'http://127.0.0.1:8080',
  );

  static String get wsUrl {
    final u = Uri.parse(baseUrl);
    final scheme = u.scheme == 'https' ? 'wss' : 'ws';
    return '$scheme://${u.host}:${u.hasPort ? u.port : 80}/ws';
  }
}
