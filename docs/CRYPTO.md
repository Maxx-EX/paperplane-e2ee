# 密码学模型 (CRYPTO.md)

## 1. 账号单向派生

```
account_id = SHA-256( username.lower() ‖ pepper )           # hex64
argon2id   = Argon2id( password, salt = account_id[0:16],
                       m = 64 MiB, t = 3, p = 1 )           # PHC 字符串
```

- `pepper` 通过 `GET /api/v1/params` 公开下发，只做域名分隔，不是秘密。
- 服务端只存 `account_id` + `argon2id_hash`，永不接触明文密码。
- 登录：客户端拿一次性 `nonce`，证明 `HMAC-SHA256(credential, nonce)`。
  抓包者拿到的 JWT / nonce 都不能重放。

## 2. E2EE（X3DH + Double Ratchet）

- 首次登录生成：
  - IdentityKey (Curve25519 密钥对)
  - SignedPreKey (签名公钥)
  - 100 个 OneTimePreKey
  - KyberPreKey（预留字段，当前 libsignal 纯 Dart 版尚未启用 PQXDH）
- 公钥上传，私钥仅存 `flutter_secure_storage`。
- 会话建立：`SessionBuilder.processPreKeyBundle(remoteBundle)`。
- 收发：`SessionCipher.encrypt / decrypt`；密文只以 `base64` 信封过服务端。

## 3. 信封协议

```
POST /api/v1/envelopes
{
  "to": "<account_id hex64>",
  "cipher_type": 3 | 4,          // 3=Whisper, 4=PreKeyWhisper
  "body": "<base64(SignalMessage)>",
  "client_msg_id": "<uuid>",
  "sealed_sender": true|false,
  "sender_token": "<盲令牌>"
}
```

- 服务端不解析 `body`，只按 `to` 转发。
- 投递成功即 `sender_id = NULL`（清除路由元数据）。
- 收件人 `ack` 后整行删除；离线密文 24h TTL 兜底清理。

## 4. Sealed Sender

- `POST /api/v1/sealed-tokens` 一次领 5 个盲令牌（24h 有效）。
- 信封携带 `sender_token`；服务端验票后**不记录**发送方账号。
- 收件人只看到密文，不知道是谁发的。

## 5. 零知识承诺

| 服务端允许 | 服务端禁止 |
|---|---|
| 转发密文信封 | 解密/记录密文 |
| 存 account_id + PHC | 存可逆用户名 |
| 在线推送 + ack 删除 | 内存 Map 当生产 store |
| 单点限流与事件统计 | 记录 JWT / Origin / 口令 |

## 6. 不要做的事

- 不要自研 E2EE；只允许集成 libsignal。
- 不要把 Argon2 跑在 UI 线程；必须 `Isolate.run`。
- 不要在客户端存 `pepper` 以外的密钥；不要持久化会话状态。
- 不要在 WS / HTTP 处理里 `return true` 跳过 Origin 校验。
