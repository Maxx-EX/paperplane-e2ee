// E2EE 层：v0.3 编译版。
// v0.4 接真正的 X25519 + AES-GCM（package:cryptography API 正在对齐）。
import 'dart:convert';
import 'dart:typed_data';

class SignalService {
  final Map<String, dynamic> _bundle = {};
  final Set<String> _established = {};

  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_bundle.isNotEmpty) return;
    _bundle['identity_key'] = base64Encode(List.generate(32, (i) => i + 1));
    _bundle['device_id'] = deviceId;
  }

  Future<Map<String, dynamic>> buildBundle() async {
    await ensureIdentity();
    return Map<String, dynamic>.from(_bundle);
  }

  Future<void> establishSession(
      String remoteId, Map<String, dynamic> bundle) async {
    _established.add(remoteId);
  }

  Future<({int type, Uint8List body})> encrypt(
      String remoteId, String plaintext) async {
    return (type: 3, body: utf8.encode(base64Encode(utf8.encode(plaintext))));
  }

  Future<String> decrypt(String remoteId, int type, Uint8List body) async {
    return utf8.decode(base64Decode(utf8.decode(body)));
  }
}
