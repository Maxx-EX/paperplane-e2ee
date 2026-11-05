# 已知错误与防御清单 (KNOWN_ERRORS.md)

## A – J：协议 / 服务端层

- **A. 会话状态必须事务提交后才返回**：bundle 取 OTP 必须在事务里 `SELECT ... FOR UPDATE SKIP LOCKED` 后再 `DELETE`，commit 完才给客户端 200。
- **B. 按 ProtocolAddress 加 `synchronized`**：对 `SignalProtocolStore` 实例级串行化，避免两个设备并发写同一会话状态。
- **C. 跳过密钥依赖 libsignal 内置**：不要自己跳过 ratchet；密钥交换失败让协议自己报 `DuplicateMessageException`。
- **D. MAC 失败提供重置会话**：`decrypt` 抛 `InvalidMessageException` 时删本地会话、重取 bundle 重发。
- **E. 跨平台序列化 schema 一致**：密文信封字段名 `cipher_type` / `body` / `client_msg_id` 两端对齐；客户端 base64 标准编码。
- **F. WS Origin 精确白名单**：空 Origin 放行（原生客户端），非空必须精确匹配 `scheme://host:port`，不要用 `strings.HasPrefix`。
- **G. Caddy 后 WS 不通**：不要用 `r.Host` 比 Origin；Caddy 会改 Host 但保留 `Origin`，只信 `Origin` 头。
- **H. 服务端绑回环**：只监听 `127.0.0.1:8080`，公网入口必须经 Caddy。
- **I. Argon2id 入 `Isolate.run`**：UI 线程不能跑 m=64MiB；慢哈希必须后台。
- **J. 自托管 Runner 构建 403**：CI 用 `GOTOOLCHAIN=local`，别在公开 Runner 上拉私有工具链。

## K-1 – K-9：内核 / 策略层

- **K-1. XDP virtio 回退 generic**：virtio 网卡不支持 native XDP，`ethtool -i` 看 driver；先尝试 `XDP_FLAGS_SKB_MODE`，失败再 `DRV_MODE`。
- **K-2. BPF map flock 串行化**：loader 用单线程读写 `ban_map`；多线程同时 `update_elem` 会破坏 LPM trie 内部结构。
- **K-3. XDP detach 显式卸载**：`SIGTERM/SIGINT/SIGHUP` 必须 `bpf_xdp_detach(ifindex, flags, NULL)`，不要依赖进程退出自动清理。
- **K-4. C++ ASan/UBSan**：CI 里 `clang++ -fsanitize=address,undefined` 跑 loader，禁止裸 `new[]/delete[]` 错配。
- **K-5. Kotlin 协程结构化并发**：顶层用 `coroutineScope` + `SupervisorJob`，单条消息异常不能杀消费循环。
- **K-6. Redis 故障 fail-open 50ms 超时**：限流/封禁查询一律 `context.WithTimeout(50ms)`，Redis 挂时服务照常（降级为无防护）。
- **K-7. 时钟统一 Redis TIME**：sidecar 用 `redis.time()` 做时间戳，不要本机时钟漂移导致规则误判。
- **K-8. bouncer 缓存+长连接**：CrowdSec bouncer 用 `cache_duration` 缓存决策，避免每个连接打 LAPI。
- **K-9. never-ban 排除长连接**：WebSocket 长连接在 CrowdSec 里走白名单（`source` 排除），不要让 bouncer 因 idle 误封正常客户端。
