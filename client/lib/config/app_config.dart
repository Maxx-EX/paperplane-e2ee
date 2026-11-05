import 'settings_store.dart';

/// 运行时服务端配置。
/// 默认 127.0.0.1:8080；在 App 设置页改地址后会存到 secure storage，重启保留。
class AppConfig {
  static final SettingsStore settings = SettingsStore();

  static String _baseUrl = 'http://127.0.0.1:8080';
  static String get baseUrl => _baseUrl;

  static Future<void> load() async {
    _baseUrl = await settings.getServerUrl();
  }

  static Future<void> setBaseUrl(String url) async {
    _baseUrl = url;
    await settings.setServerUrl(url);
  }

  static String get wsUrl {
    final u = Uri.parse(_baseUrl);
    final scheme = u.scheme == 'https' ? 'wss' : 'ws';
    return '$scheme://${u.host}:${u.hasPort ? u.port : 80}/ws';
  }
}
