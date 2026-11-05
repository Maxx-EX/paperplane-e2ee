// 账号单向派生链：
//   account_id = SHA-256(username ‖ pepper)
//   hash       = Argon2id(password, salt = account_id[:16], m=64MiB, t=3, p=1)
//   proof      = HMAC-SHA256(credential, nonce)
//
// 服务端只存 account_id + PHC，从不接触明文口令。
import 'dart:convert';
import 'dart:isolate';

import 'package:crypto/crypto.dart';
import 'package:serverpod_argon2/serverpod_argon2.dart';

class Argon2Params {
  final int memoryKib;
  final int iterations;
  final int parallelism;
  const Argon2Params({
    required this.memoryKib,
    required this.iterations,
    required this.parallelism,
  });

  factory Argon2Params.fromServer(Map<String, dynamic> json) => Argon2Params(
        memoryKib: (json['argon2_memory_kib'] as num?)?.toInt() ?? 65536,
        iterations: (json['argon2_iterations'] as num?)?.toInt() ?? 3,
        parallelism: (json['argon2_parallelism'] as num?)?.toInt() ?? 1,
      );
}

class AccountPassword {
  static String accountId(String username, String pepper) =>
      sha256.convert(utf8.encode('${username.trim().toLowerCase()}$pepper')).toString();

  /// Argon2id 故意很慢：必须在 Isolate.run 里跑，绝不阻塞 UI。
  static Future<String> hashPassword(
    String password,
    String accountId,
    Argon2Params params,
  ) async {
    final argon2 = await Argon2.load();
    final salt = utf8.encode(accountId.substring(0, 16));
    // 注意：字段名随 serverpod_argon2 版本可能微调；如编译失败，
    // 仅需修改下面这一行（memory/iterations/parallelism）。
    final parameters = Argon2Parameters(
      memory: params.memoryKib,
      iterations: params.iterations,
      parallelism: params.parallelism,
    );
    return Isolate.run(() => argon2.hash(password, salt: salt, parameters: parameters));
  }

  static String proof(String credential, String nonce) =>
      Hmac(sha256, utf8.encode(credential)).convert(utf8.encode(nonce)).toString();
}
