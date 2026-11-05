import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../config/app_config.dart';
import '../state/chat_store.dart';

class SettingsPage extends StatelessWidget {
  const SettingsPage({super.key});

  @override
  Widget build(BuildContext context) {
    final store = context.watch<ChatStore>();
    return Scaffold(
      appBar: AppBar(title: const Text('设置')),
      body: ListView(
        children: [
          const ListTile(
            title: Text('服务器地址'),
            subtitle: Text(AppConfig.baseUrl),
            leading: Icon(Icons.dns),
          ),
          ListTile(
            title: Text('账号 ID'),
            subtitle: Text(store.accountId ?? '未登录'),
            leading: Icon(Icons.fingerprint),
          ),
          const Divider(),
          ListTile(
            title: const Text('关于'),
            subtitle: const Text('Paperplane v0.2.0 · AGPL-3.0'),
            leading: const Icon(Icons.info_outline),
            onTap: () => showAboutDialog(
              context: context,
              applicationName: 'Paperplane',
              applicationVersion: '0.2.0',
              applicationLegalese: '零知识 E2EE 加密聊天',
            ),
          ),
          const SizedBox(height: 24),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: FilledButton.icon(
              onPressed: () async {
                await context.read<ChatStore>().logout();
                if (context.mounted) Navigator.pop(context);
              },
              icon: const Icon(Icons.logout),
              label: const Text('退出登录'),
              style: FilledButton.styleFrom(
                minimumSize: const Size.fromHeight(48),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
