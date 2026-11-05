import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:uuid/uuid.dart';

import '../config/app_config.dart';
import '../crypto/password.dart';
import '../crypto/signal_service.dart';
import '../net/api_client.dart';
import '../net/realtime.dart';

enum MsgStatus { sending, sent, delivered, read }

class ChatMsg {
  final String id;
  final bool fromMe;
  final String text;
  final MsgStatus status;
  final DateTime ts;
  const ChatMsg({
    required this.id,
    required this.fromMe,
    required this.text,
    required this.status,
    required this.ts,
  });
  ChatMsg copy({MsgStatus? status}) => ChatMsg(
        id: id, fromMe: fromMe, text: text,
        status: status ?? this.status, ts: ts,
      );
}

class ChatStore extends ChangeNotifier {
  final api = ApiClient();
  final signal = SignalService();
  final realtime = Realtime();
  final storage = const FlutterSecureStorage();

  String? accountId;
  String? contactId;
  final List<ChatMsg> messages = [];

  bool get loggedIn => accountId != null;

  Future<void> bootstrap() async {
    accountId = await storage.read(key: 'account_id');
    api.token = await storage.read(key: 'token');
    if (accountId != null && api.token != null) {
      try {
        await _openSocket();
      } catch (e) {
        if (kDebugMode) print('resume socket failed: $e');
      }
    }
    notifyListeners();
  }

  Future<void> authenticate(
      String username, String password, bool register) async {
    final params = await api.getParams();
    final pepper = params['pepper'] as String;
    final derived = AccountPassword.accountId(username, pepper);
    final nonceResp = await api.challenge(username, register ? 'register' : 'login');
    final nonce = nonceResp['nonce'] as String;
    final argon = Argon2Params.fromServer(params);
    // Argon2id 在 Isolate 里跑，不卡 UI。
    final hash = await AccountPassword.hashPassword(password, derived, argon);
    final proof = AccountPassword.proof(hash, nonce);

    final Map<String, dynamic> resp;
    if (register) {
      resp = await api.register(username, hash, nonce, proof);
      await signal.ensureIdentity();
      await api.uploadKeys(await signal.buildBundle());
    } else {
      resp = await api.login(username, nonce, proof);
    }
    accountId = resp['account_id'] as String;
    api.token = resp['token'] as String;
    await storage.write(key: 'account_id', value: accountId!);
    await storage.write(key: 'token', value: api.token!);
    notifyListeners();
    await _openSocket();
  }

  Future<void> _openSocket() async {
    final ticket = await api.wsTicket();
    realtime.connect(ticket, _onFrame);
  }

  Future<void> startChatWith(String targetId) async {
    contactId = targetId.trim().toLowerCase();
    final bundle = await api.fetchBundle(contactId!);
    await signal.establishSession(contactId!, bundle);
    final pend = await api.pullEnvelopes();
    final ackIds = <int>[];
    for (final e in pend) {
      final m = e as Map<String, dynamic>;
      await _consume(m);
      ackIds.add(m['id'] as int);
    }
    await api.ack(ackIds);
    notifyListeners();
  }

  Future<void> _consume(Map<String, dynamic> frame) async {
    final body = base64Decode(frame['body'] as String);
    final from = frame['from'] as String? ?? '';
    final plain = await signal.decrypt(from, frame['cipher_type'] as int, body);
    messages.add(ChatMsg(
      id: frame['client_msg_id'] as String? ?? const Uuid().v4(),
      fromMe: false,
      text: plain,
      status: MsgStatus.delivered,
      ts: DateTime.now(),
    ));
  }

  void _onFrame(Map<String, dynamic> frame) async {
    if (frame['type'] != 'envelope') return;
    await _consume(frame);
    realtime.ack([frame['id'] as int]);
    notifyListeners();
  }

  Future<void> sendText(String text) async {
    final target = contactId;
    if (target == null || text.trim().isEmpty) return;
    final id = const Uuid().v4();
    messages.add(ChatMsg(
        id: id, fromMe: true, text: text, status: MsgStatus.sending, ts: DateTime.now()));
    notifyListeners();
    try {
      final ct = await signal.encrypt(target, text);
      final r = await api.sendEnvelope(
        to: target,
        cipherType: ct.type,
        body: base64Encode(ct.body),
        clientMsgId: id,
      );
      final i = messages.indexWhere((m) => m.id == id);
      if (i >= 0) messages[i] = messages[i].copy(status: MsgStatus.sent);
      if (r['delivered'] == true) {
        messages[i] = messages[i].copy(status: MsgStatus.delivered);
      }
    } catch (e) {
      final i = messages.indexWhere((m) => m.id == id);
      if (i >= 0) messages[i] = messages[i].copy(status: MsgStatus.read);
    }
    notifyListeners();
  }

  Future<void> logout() async {
    realtime.close();
    await storage.delete(key: 'token');
    await storage.delete(key: 'account_id');
    accountId = null;
    messages.clear();
    notifyListeners();
  }

  Future<void> setServerUrl(String url) async {
    await storage.write(key: 'server_url', value: url);
    // 重启 App 后由 ApiClient 读取；这里只持久化。
  }
}
