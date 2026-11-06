// E2EE 层：X25519 密钥协商 + AES-256-GCM 加密。
// 基于 package:cryptography（Dart 官方加密库，非自研算法）。
// 注意：这是简化版 E2EE（无 Double Ratchet），适合 1:1 聊天；
// 完整 X3DH + Double Ratchet 在后续版本接 libsignal。
import 'dart:convert';
import 'dart:typed_data';

import 'package:cryptography/cryptography.dart';

class SignalService {
  X25519KeyPair? _identity;
  final Map<String, SimplePublicKey> _remoteKeys = {};

  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_identity != null) return;
    _identity = await X25519().newKeyPair();
  }

  Future<Map<String, dynamic>> buildBundle() async {
    await ensureIdentity();
    final pk = await _identity!.extractPublicKey();
    return {
      'identity_key': base64Encode(pk.bytes),
      'device_id': deviceId,
    };
  }

  Future<void> establishSession(
      String remoteId, Map<String, dynamic> bundle) async {
    final pkB64 = bundle['identity_key'] as String;
    _remoteKeys[remoteId] =
        SimplePublicKey(base64Decode(pkB64), type: KeyPairType.x25519);
  }

  Future<({int type, Uint8List body})> encrypt(
      String remoteId, String plaintext) async {
    await ensureIdentity();
    final remotePk = _remoteKeys[remoteId];
    if (remotePk == null) throw StateError('no session with $remoteId');
    final shared = await X25519().sharedSecret(
      keyPair: _identity!,
      remotePublicKey: remotePk,
    );
    final gcm = AesGcm.with256bits();
    final box = await gcm.encrypt(
      utf8.encode(plaintext),
      secretKey: SecretKey(shared.bytes),
      nonce: List.generate(12, (i) => i * 7 + 1),
    );
    final body = Uint8List.fromList([
      ...box.nonce,
      ...box.mac.bytes,
      ...box.cipherText,
    ]);
    return (type: 3, body: body);
  }

  Future<String> decrypt(String remoteId, int type, Uint8List body) async {
    await ensureIdentity();
    final remotePk = _remoteKeys[remoteId];
    if (remotePk == null) throw StateError('no session with $remoteId');
    final nonce = body.sublist(0, 12);
    final mac = body.sublist(12, 28);
    final cipherText = body.sublist(28);
    final shared = await X25519().sharedSecret(
      keyPair: _identity!,
      remotePublicKey: remotePk,
    );
    final gcm = AesGcm.with256bits();
    final plain = await gcm.decrypt(
      SecretBox(cipherText, nonce: nonce, mac: Mac(mac)),
      secretKey: SecretKey(shared.bytes),
    );
    return utf8.decode(plain);
  }
}
