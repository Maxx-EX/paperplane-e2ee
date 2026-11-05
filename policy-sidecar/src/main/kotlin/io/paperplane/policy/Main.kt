package io.paperplane.policy

import io.lettuce.core.*
import kotlinx.coroutines.*
import org.slf4j.LoggerFactory
import java.time.Duration

/**
 * Kotlin sidecar 入口。
 * 消费 Redis Stream 事件流，运行跨维度封禁规则，结构化并发。
 */
suspend fun main(): Unit = coroutineScope {
    val log = LoggerFactory.getLogger("main")
    val cfgPath = System.getenv("CONFIG_FILE") ?: "config/application.yaml"
    val cfg = AppConfig.load(cfgPath)

    val redisUri = RedisURI.Builder
        .redis(cfg.redis.host, cfg.redis.port)
        .apply { if (cfg.redis.password.isNotEmpty()) withPassword(cfg.redis.password) }
        .withDatabase(cfg.redis.db)
        .withTimeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs + 100))
        .build()

    val client = RedisClient.create(redisUri)
    val conn = client.connect()
    val sync = conn.sync()

    // mkstream(true): 流不存在时自动创建；BUSYGROUP 异常吞掉。
    try {
        sync.xgroupCreate(
            XReadArgs.StreamOffset.latest(cfg.policy.stream),
            cfg.policy.consumerGroup,
            XGroupCreateArgs().mkstream(true),
        )
    } catch (e: RedisBusyException) {
        log.info("consumer group already exists")
    }

    val bans = BanPublisher(cfg, sync)
    val rules = Rules(cfg, sync, bans)

    log.info(
        "policy sidecar ready stream={} group={} consumer={}",
        cfg.policy.stream, cfg.policy.consumerGroup, cfg.policy.consumerName
    )

    while (isActive) {
        try {
            val msgs = sync.xreadgroup(
                Consumer.from(cfg.policy.consumerGroup, cfg.policy.consumerName),
                XReadArgs().count(50).block(Duration.ofSeconds(2)),
                XReadArgs.StreamOffset.lastConsumed(cfg.policy.stream),
            ) ?: emptyList()

            for (msg in msgs) {
                try {
                    // Lettuce 返回的 Map<key,value> 全部是 String。
                    rules.handle(msg.body)
                } catch (e: Exception) {
                    log.warn("rule failed on message {}", msg.id, e)
                }
                try {
                    conn.async().xack(cfg.policy.stream, cfg.policy.consumerGroup, msg.id)
                } catch (e: Exception) {
                    log.warn("xack failed {}", msg.id, e)
                }
            }
        } catch (e: Exception) {
            log.warn("stream read error, backoff", e)
            delay(2000)
        }
    }
}
