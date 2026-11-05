import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;

import '../config/app_config.dart';

class ApiException implements Exception {
  final int status;
  final String code;
  final String message;
  ApiException(this.status, this.code, this.message);
  @override
  String toString() => '[$code] $message';
}

class ApiClient {
  String base;
  String? token;
  final http.Client _client = http.Client();

  ApiClient({this.base = 'http://127.0.0.1:8080', this.token});

  Map<String, String> get _headers => {
        'Content-Type': 'application/json',
        if (token != null) 'Authorization': 'Bearer $token',
      };

  Future<Map<String, dynamic>> _handle(http.Response resp) async {
    final body = resp.body.isEmpty
        ? <String, dynamic>{}
        : jsonDecode(resp.body) as Map<String, dynamic>;
    if (resp.statusCode >= 200 && resp.statusCode < 300) return body;
    final err = body['error'];
    throw ApiException(
      resp.statusCode,
      err is Map ? err['code']?.toString() ?? 'error' : 'error',
      err is Map ? err['message']?.toString() ?? '' : resp.body,
    );
  }

  Future<Map<String, dynamic>> getParams() async {
    final r = await _client.get(Uri.parse('$base/api/v1/params'));
    return _handle(r);
  }

  Future<Map<String, dynamic>> challenge(String username, String intent) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/challenge'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username, 'intent': intent}),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> register(
      String username, String hash, String nonce, String proof) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'username': username,
        'argon2id_hash': hash,
        'nonce': nonce,
        'proof': proof,
      }),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> login(
      String username, String nonce, String proof) async {
    final r = await _client.post(
      Uri.parse('$base/api/v1/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'username': username, 'nonce': nonce, 'proof': proof}),
    );
    return _handle(r);
  }

  Future<Map<String, dynamic>> uploadKeys(Map<String, dynamic> bundle) async {
    final r = await _client.post(Uri.parse('$base/api/v1/keys'),
        headers: _headers, body: jsonEncode(bundle));
    return _handle(r);
  }

  Future<Map<String, dynamic>> fetchBundle(String accountId) async {
    final r = await _client.get(Uri.parse('$base/api/v1/keys/$accountId'),
        headers: _headers);
    return _handle(r);
  }

  Future<Map<String, dynamic>> sendEnvelope({
    required String to,
    required int cipherType,
    required String body,
    required String clientMsgId,
  }) async {
    final r = await _client.post(Uri.parse('$base/api/v1/envelopes'),
        headers: _headers,
        body: jsonEncode({
          'to': to,
          'cipher_type': cipherType,
          'body': body,
          'client_msg_id': clientMsgId,
        }));
    return _handle(r);
  }

  Future<List<dynamic>> pullEnvelopes() async {
    final r = await _client.get(Uri.parse('$base/api/v1/envelopes'),
        headers: _headers);
    final m = await _handle(r);
    return (m['envelopes'] as List? ?? const []);
  }

  Future<void> ack(List<int> ids) async {
    if (ids.isEmpty) return;
    final r = await _client.post(Uri.parse('$base/api/v1/envelopes/ack'),
        headers: _headers, body: jsonEncode({'ids': ids}));
    await _handle(r);
  }

  Future<String> wsTicket() async {
    final r = await _client.post(Uri.parse('$base/api/v1/ws-ticket'),
        headers: _headers, body: jsonEncode({}));
    final m = await _handle(r);
    return m['ticket'] as String;
  }
}
