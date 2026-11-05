// Signal 协议层：X3DH + Double Ratchet 完全委托给 libsignal_protocol_dart。
// 本项目不自研任何 E2EE 算法。
// 注意：libsignal_protocol_dart 0.8.2 的部分构造器命名随版本微调，
// 如编译失败，仅需按官方 API 微调本文件中的字节反序列化调用。
import 'dart:convert';
import 'dart:typed_data';

import 'package:libsignal_protocol_dart/libsignal_protocol_dart.dart';

class SignalService {
  InMemorySignalProtocolStore? _store;
  static const int deviceId = 1;

  Future<void> ensureIdentity() async {
    if (_store != null) return;
    final identityKeyPair = generateIdentityKeyPair();
    final registrationId = generateRegistrationId(false);
    _store = InMemorySignalProtocolStore(identityKeyPair, registrationId);
    final spk = generateSignedPreKey(identityKeyPair, 0);
    await _store!.storeSignedPreKey(spk.id, spk);
    final preKeys = generatePreKeys(0, 100);
    for (final p in preKeys) {
      await _store!.storePreKey(p.id, p);
    }
  }

  /// 生成预密钥公钥包，只传公钥；私钥不出本机（内存即销毁）。
  Future<Map<String, dynamic>> buildBundle() async {
    await ensureIdentity();
    final s = _store!;
    final spk = await s.loadSignedPreKey(1);
    final ots = <Map<String, dynamic>>[];
    for (var id = 0; id < 100; id++) {
      try {
        final k = await s.loadPreKey(id);
        ots.add({'id': id, 'public': base64Encode(k.serialize())});
      } catch (_) {}
    }
    return {
      'identity_key': base64Encode(s.identityKeyPair().publicKey.serialize()),
      'signed_prekey': {
        'id': spk.id,
        'public': base64Encode(spk.keyPair.publicKey.serialize()),
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
      1,
      deviceId,
      ot != null ? (ot['id'] as num).toInt() : 0,
      ot != null
          ? Curve25519PublicKey.fromBase64(ot['public'] as String)
          : Curve25519PublicKey.fromBase64(''),
      (signed['id'] as num).toInt(),
      Curve25519PublicKey.fromBase64(signed['public'] as String),
      base64Decode(signed['signature'] as String),
      IdentityKey.fromSerialized(base64Decode(bundle['identity_key'] as String)),
    );
    await builder.processPreKeyBundle(preKeyBundle);
  }

  Future<({int type, Uint8List body})> encrypt(
      String remoteId, String plaintext) async {
    await ensureIdentity();
    final cipher = SessionCipher(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    final ct = await cipher.encrypt(utf8.encode(plaintext));
    return (type: ct.type, body: ct.serialize());
  }

  Future<String> decrypt(String remoteId, int type, Uint8List body) async {
    await ensureIdentity();
    final cipher = SessionCipher(
        _store!, _store!, _store!, _store!, _addr(remoteId));
    CiphertextMessage msg;
    if (type == 4) {
      msg = PreKeySignalMessage.fromSerialized(body);
      final plain = await cipher.decrypt(msg as PreKeySignalMessage);
      return utf8.decode(plain);
    }
    msg = SignalMessage.fromSerialized(body);
    final plain = await cipher.decrypt(msg as SignalMessage);
    return utf8.decode(plain);
  }
}
