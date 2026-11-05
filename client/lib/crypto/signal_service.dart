// Signal 协议层：X3DH + Double Ratchet，完全基于 libsignal_protocol_dart 0.8.2。
// 本项目不自研任何 E2EE 算法。
// 参考：https://pub.dev/packages/libsignal_protocol_dart/example
import 'dart:convert';
import 'dart:typed_data';

import 'package:libsignal_protocol_dart/libsignal_protocol_dart.dart';

class SignalService {
  InMemorySignalProtocolStore? _store;
  IdentityKeyPair? _identityKeyPair;
  int _registrationId = 0;
  bool _installed = false;

  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_store != null) return;
    if (!_installed) {
      await install();
      _installed = true;
    }
    _identityKeyPair = generateIdentityKeyPair();
    _registrationId = generateRegistrationId(false);
    _store = InMemorySignalProtocolStore(_identityKeyPair!, _registrationId);
    final spk = generateSignedPreKey(_identityKeyPair!, 0);
    await _store!.storeSignedPreKey(spk.id, spk);
    final preKeys = generatePreKeys(0, 100);
    for (final p in preKeys) {
      await _store!.storePreKey(p.id, p);
    }
  }

  /// 生成预密钥公钥包，只传公钥；私钥不出本机。
  Future<Map<String, dynamic>> buildBundle() async {
    await ensureIdentity();
    final s = _store!;
    final spk = await s.loadSignedPreKey(0);
    final ots = <Map<String, dynamic>>[];
    for (var id = 0; id < 100; id++) {
      try {
        final k = await s.loadPreKey(id);
        ots.add({
          'id': id,
          'public': base64Encode(k.getKeyPair().publicKey.serialize()),
        });
      } catch (_) {}
    }
    return {
      'identity_key': base64Encode(_identityKeyPair!.getPublicKey().serialize()),
      'signed_prekey': {
        'id': spk.id,
        'public': base64Encode(spk.getKeyPair().publicKey.serialize()),
        'signature': base64Encode(spk.signature),
      },
      'one_time_prekeys': ots,
    };
  }

  SignalProtocolAddress _addr(String remoteId) =>
      SignalProtocolAddress(remoteId, deviceId);

  /// 用服务端返回的预密钥包建立 X3DH 会话。
  Future<void> establishSession(
      String remoteId, Map<String, dynamic> bundle) async {
    await ensureIdentity();
    final builder = SessionBuilder(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    final signed = bundle['signed_prekey'] as Map<String, dynamic>;
    final ot = bundle['one_time_prekey'] as Map<String, dynamic>?;
    final preKeyBundle = PreKeyBundle(
      _registrationId,
      deviceId,
      ot != null ? (ot['id'] as num).toInt() : 0,
      ot != null
          ? _publicKeyFromBytes(base64Decode(ot['public'] as String))
          : null,
      (signed['id'] as num).toInt(),
      _publicKeyFromBytes(base64Decode(signed['public'] as String)),
      base64Decode(signed['signature'] as String),
      IdentityKey(base64Decode(bundle['identity_key'] as String)),
    );
    await builder.processPreKeyBundle(preKeyBundle);
  }

  ECPublicKey _publicKeyFromBytes(Uint8List bytes) {
    // libsignal 内部用 ECPublicKey 接口；从字节构造用 Curve25519PublicKey。
    return Curve25519PublicKey.fromBytes(bytes);
  }

  Future<({int type, Uint8List body})> encrypt(
      String remoteId, String plaintext) async {
    await ensureIdentity();
    final cipher = SessionCipher.fromStore(_store!, _addr(remoteId));
    final ct = await cipher.encrypt(utf8.encode(plaintext));
    return (type: ct.getType(), body: ct.serialize());
  }

  Future<String> decrypt(String remoteId, int type, Uint8List body) async {
    await ensureIdentity();
    final cipher = SessionCipher.fromStore(_store!, _addr(remoteId));
    late Uint8List plain;
    if (type == CiphertextMessage.prekeyType) {
      final msg = PreKeySignalMessage.fromSerialized(body);
      await cipher.decryptWithCallback(msg, (p) => plain = p);
    } else {
      final msg = SignalMessage.fromSerialized(body);
      await cipher.decryptWithCallback(msg, (p) => plain = p);
    }
    return utf8.decode(plain);
  }
}
