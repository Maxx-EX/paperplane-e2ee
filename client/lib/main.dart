import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'state/chat_store.dart';
import 'ui/auth_page.dart';
import 'ui/chat_page.dart';
import 'ui/theme.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const PaperplaneApp());
}

class PaperplaneApp extends StatelessWidget {
  const PaperplaneApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => ChatStore()..bootstrap(),
      child: MaterialApp(
        title: 'Paperplane',
        debugShowCheckedModeBanner: false,
        theme: buildLightTheme(),
        darkTheme: buildDarkTheme(),
        themeMode: ThemeMode.system,
        home: const AuthGate(),
      ),
    );
  }
}

class AuthGate extends StatelessWidget {
  const AuthGate({super.key});

  @override
  Widget build(BuildContext context) {
    return Consumer<ChatStore>(
      builder: (_, store, __) =>
          store.loggedIn ? const ChatPage() : const AuthPage(),
    );
  }
}
