import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../config/app_config.dart';
import '../state/chat_store.dart';

class SettingsPage extends StatefulWidget {
  const SettingsPage({super.key});

  @override
  State<SettingsPage> createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final _ctrl = TextEditingController();
  String? _saved;

  @override
  void initState() {
    super.initState();
    _ctrl.text = AppConfig.baseUrl;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('设置')),
      body: ListView(
        padding: const EdgeInsets.all(20),
        children: [
          const Text('服务器地址', style: TextStyle(fontWeight: 600)),
          const SizedBox(height: 8),
          TextField(
            controller: _ctrl,
            decoration: const InputDecoration(
              hintText: 'http://192.168.1.5:8080',
              helperText: 'Termux 里跑的服务端，或你的 VPS 地址。改完自动保存。',
            ),
            keyboardType: TextInputType.url,
          ),
          const SizedBox(height: 12),
          if (_saved != null)
            Text(_saved!, style: const TextStyle(color: Colors.green)),
          const SizedBox(height: 8),
          FilledButton(
            onPressed: () async {
              await AppConfig.setBaseUrl(_ctrl.text.trim());
              if (context.mounted) {
                final store = context.read<ChatStore>();
                store.api.base = AppConfig.baseUrl;
                setState(() => _saved = '已保存：${AppConfig.baseUrl}');
              }
            },
            child: const Text('保存'),
          ),
          const Divider(height: 40),
          const Text('关于', style: TextStyle(fontWeight: 600)),
          const ListTile(
            title: Text('Paperplane'),
            subtitle: Text('v0.2.0 · 零知识 E2EE 聊天 · AGPL-3.0'),
          ),
          const ListTile(
            title: Text('图标'),
            subtitle: Text('黑白折纸鸽子'),
          ),
        ],
      ),
    );
  }
}
