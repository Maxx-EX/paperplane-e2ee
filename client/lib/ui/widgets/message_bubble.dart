import 'package:flutter/material.dart';

import '../state/chat_store.dart';
import 'theme.dart';

class MessageBubble extends StatelessWidget {
  final ChatMsg msg;
  const MessageBubble({super.key, required this.msg});

  @override
  Widget build(BuildContext context) {
    final isMe = msg.fromMe;
    return Align(
      alignment: isMe ? Alignment.centerRight : Alignment.centerLeft,
      child: Container(
        margin: const EdgeInsets.symmetric(vertical: 3, horizontal: 12),
        padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 14),
        constraints: BoxConstraints(
          maxWidth: MediaQuery.of(context).size.width * 0.75,
        ),
        decoration: BoxDecoration(
          gradient: isMe ? bubbleGradient : null,
          color: isMe ? null : Colors.white,
          borderRadius: BorderRadius.circular(bubbleRadius),
          boxShadow: const [
            BoxShadow(blurRadius: 4, color: Colors.black12, offset: Offset(0, 1)),
          ],
        ),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Flexible(
              child: Text(
                msg.text,
                style: TextStyle(
                  color: isMe ? Colors.white : const Color(0xFF1B2430),
                  fontSize: 15,
                ),
              ),
            ),
            const SizedBox(width: 6),
            if (isMe) Icon(_tick(msg.status), size: 14, color: Colors.white70),
          ],
        ),
      ),
    );
  }

  IconData _tick(MsgStatus s) => switch (s) {
        MsgStatus.sending => Icons.access_time,
        MsgStatus.sent => Icons.check,
        MsgStatus.delivered => Icons.done_all,
        MsgStatus.read => Icons.done_all,
      };
}
