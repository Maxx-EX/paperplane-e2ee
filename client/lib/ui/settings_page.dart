import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../state/chat_store.dart';

class SettingsPage extends StatefulWidget {
  const SettingsPage({super.key});

  @override
  State<SettingsPage> createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final _ctrl = TextEditingController();
  bool _saved = false;

  @override
  void initState() {
    super.initState();
    final store = Provider.of<ChatStore>(context, listen: false);
    _ctrl.text = store.api.base;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('设置')),
      body: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const Text('服务器地址', style: TextStyle(fontWeight: FontWeight.w600)),
            const SizedBox(height: 8),
            TextField(
              controller: _ctrl,
              decoration: const InputDecoration(
                hintText: 'http://127.0.0.1:8080',
                helperText: '默认 127.0.0.1:8080；换地址后需重启 App 生效',
              ),
            ),
            const SizedBox(height: 20),
            FilledButton(
              onPressed: () async {
                final store = Provider.of<ChatStore>(context, listen: false);
                await store.setServerUrl(_ctrl.text.trim());
                setState(() => _saved = true);
              },
              child: const Text('保存'),
            ),
            if (_saved)
              const Padding(
                padding: EdgeInsets.only(top: 12),
                child: Text('已保存，重启 App 生效',
                    style: TextStyle(color: Colors.green)),
              ),
            const SizedBox(height: 30),
            const Divider(),
            const SizedBox(height: 10),
            ListTile(
              leading: const Icon(Icons.badge),
              title: const Text('我的账号 ID'),
              subtitle: Text(
                Provider.of<ChatStore>(context).accountId ?? '未登录',
                style: const TextStyle(fontFamily: 'monospace', fontSize: 12),
              ),
              onTap: () {
                final id = Provider.of<ChatStore>(context, listen: false).accountId ?? '';
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('已复制: $id')),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}
