// 服务端配置：通过 --dart-define=SERVER_URL=... 注入，默认连本机 8080
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
