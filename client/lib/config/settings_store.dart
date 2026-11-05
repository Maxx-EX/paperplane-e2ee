import 'package:flutter_secure_storage/flutter_secure_storage.dart';

/// 持久化客户端设置（服务器地址等）。
/// 改地址不需要重新打包，在 App 内设置页改即可。
class SettingsStore {
  static const _keyServer = 'server_url';
  static const _default = 'http://127.0.0.1:8080';

  final _storage = const FlutterSecureStorage();

  Future<String> getServerUrl() async {
    final v = await _storage.read(key: _keyServer);
    return v ?? _default;
  }

  Future<void> setServerUrl(String url) async {
    await _storage.write(key: _keyServer, value: url.trim());
  }

  String get defaultServerUrl => _default;
}
