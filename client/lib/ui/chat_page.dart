import 'package:flutter/material.dart';
import 'package:flutter_chat_ui/flutter_chat_ui.dart';
import 'package:provider/provider.dart';

import '../state/chat_store.dart';
import 'widgets/composer.dart';
import 'widgets/message_bubble.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Paperplane'),
        actions: [
          IconButton(
            onPressed: () => context.read<ChatStore>().logout(),
            icon: const Icon(Icons.logout),
          ),
        ],
      ),
      body: Consumer<ChatStore>(
        builder: (_, store, __) {
          if (store.contactId == null) {
            return Center(
              child: OutlinedButton.icon(
                onPressed: () => _openTargetDialog(context, store),
                icon: const Icon(Icons.person_add),
                label: const Text('输入对方账号 ID 开始聊天'),
              ),
            );
          }
          // ChatList：Flyer Chat 提供动画分页与底部输入栏容器。
          return ChatList(
            items: store.messages,
            itemBuilder: (item, index) =>
                MessageBubble(msg: item as ChatMsg),
            bottomWidget: Composer(onSend: store.sendText),
            isLastPage: true,
            bubbleRtlAlignment: BubbleRtlAlignment.right,
            typingIndicatorOptions: const TypingIndicatorOptions(),
          );
        },
      ),
    );
  }

  Future<void> _openTargetDialog(BuildContext context, ChatStore store) async {
    final ctrl = TextEditingController();
    final id = await showDialog<String>(
      context: context,
      builder: (_) => AlertDialog(
        title: const Text('目标账号 ID'),
        content: TextField(
          controller: ctrl,
          decoration: const InputDecoration(hintText: '64 位 hex account id'),
        ),
        actions: [
          TextButton(
              onPressed: () => Navigator.pop(context), child: const Text('取消')),
          FilledButton(
              onPressed: () => Navigator.pop(context, ctrl.text),
              child: const Text('开始')),
        ],
      ),
    );
    if (id == null || id.isEmpty) return;
    try {
      await store.startChatWith(id);
    } catch (e) {
      if (context.mounted) {
        ScaffoldMessenger.of(context)
            .showSnackBar(SnackBar(content: Text('会话建立失败: $e')));
      }
    }
  }
}
