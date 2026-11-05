import 'dart:convert';

import 'package:web_socket_channel/web_socket_channel.dart';

import '../config/app_config.dart';

/// 单连接实时通道：收到 envelope 帧后回调；发送 ack 即完成投递闭环。
class Realtime {
  WebSocketChannel? _ch;
  bool _closed = false;

  void connect(String ticket, void Function(Map<String, dynamic> frame) onFrame) {
    final uri = Uri.parse('${AppConfig.wsUrl}?ticket=$ticket');
    final ch = WebSocketChannel.connect(uri);
    _ch = ch;
    ch.stream.listen((raw) {
      if (raw is String) {
        try {
          onFrame(jsonDecode(raw) as Map<String, dynamic>);
        } catch (_) {}
      }
    }, onError: (_) => close());
  }

  void ack(List<int> ids) {
    if (ids.isEmpty) return;
    _ch?.sink.add(jsonEncode({'type': 'ack', 'ids': ids}));
  }

  void close() {
    if (_closed) return;
    _closed = true;
    _ch?.sink.close();
  }
}
