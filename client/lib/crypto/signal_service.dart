// Signal 协议层（v0.2 编译占位）：
// 完整 X3DH + Double Ratchet 由 libsignal_protocol_dart 提供，
// 此处先用可编译的最小实现跑通 APK 打包；
// 真正的 libsignal 接线在 v0.3 替换本文件即可（接口已对齐）。
import 'dart:convert';
import 'dart:typed_data';

class SignalService {
  final Map<String, dynamic> _bundle = {};
  final Set<String> _established = {};

  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_bundle.isNotEmpty) return;
    _bundle['identity_key'] = base64Encode(List.generate(33, (i) => i + 1));
    _bundle['signed_prekey'] = {
      'id': 1,
      'public': base64Encode(List.generate(32, (i) => i + 2)),
      'signature': base64Encode(List.generate(64, (i) => i + 3)),
    };
    _bundle['one_time_prekeys'] = [
      for (var i = 0; i < 10; i++)
        {'id': i, 'public': base64Encode(List.generate(32, (j) => i + j + 4))},
    ];
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
