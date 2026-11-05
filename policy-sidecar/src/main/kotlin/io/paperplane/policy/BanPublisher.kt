package io.paperplane.policy

import io.lettuce.core.api.sync.RedisCommands
import io.lettuce.core.XAddArgs
import org.slf4j.LoggerFactory
import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.http.HttpResponse
import java.time.Duration

/**
 * 下发封禁的唯一出口：
 *   1) Redis ZSET（Go 内嵌层 + XDP loader 消费）
 *   2) Redis Stream（/var/lib/paperplane/bans.txt 镜像）
 *   3) CrowdSec LAPI（机器人群封禁协同）
 *
 * 任何外部调用失败都 fail-open（50ms 超时），绝不阻塞事件消费。
 */
class BanPublisher(
    private val cfg: AppConfig,
    private val sync: RedisCommands<String, String>,
) {
    private val log = LoggerFactory.getLogger("bans")
    private val http = HttpClient.newBuilder()
        .connectTimeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs))
        .build()

    fun banIP(ip: String, reason: String) {
        val nowSec = redisNowSec()
        val until = nowSec + cfg.policy.banSeconds
        try {
            sync.zadd("paperplane:ban:ip", until.toDouble(), ip)
            sync.xadd(
                XAddArgs().stream("paperplane:bans").maxLen(10_000, true)
                    .add("scope", "ip").add("value", ip).add("reason", reason)
                    .add("until", until.toString())
            )
            crowdSec("ip", ip, reason)
        } catch (e: Exception) {
            log.warn("ban publish failed, failing open", e)
        }
    }

    fun banAccount(account: String, reason: String) {
        val until = redisNowSec() + cfg.policy.banSeconds
        try {
            sync.zadd("paperplane:ban:account", until.toDouble(), account)
            sync.xadd(
                XAddArgs().stream("paperplane:bans").maxLen(10_000, true)
                    .add("scope", "account").add("value", account).add("reason", reason)
                    .add("until", until.toString())
            )
        } catch (e: Exception) {
            log.warn("account ban publish failed", e)
        }
    }

    private fun crowdSec(scope: String, value: String, reason: String) {
        val apiKey = System.getenv(cfg.crowdsec.apiKeyEnv) ?: return
        if (apiKey.isBlank()) return
        try {
            val body = """{"duration":"${cfg.crowdsec.decisionDuration}","scenario":"paperplane:$reason","type":"$scope","value":"$value","origin":"policy-sidecar"}"""
            val req = HttpRequest.newBuilder()
                .uri(URI.create(cfg.crowdsec.lapiUrl.trimEnd('/') + "/v1/decisions"))
                .timeout(Duration.ofMillis(cfg.policy.failOpenTimeoutMs))
                .header("Authorization", "ApiKey $apiKey")
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(body))
                .build()
            http.send(req, HttpResponse.BodyHandlers.discarding())
        } catch (e: Exception) {
            log.debug("crowdsec call skipped", e)
        }
    }

    private fun redisNowSec(): Long = try {
        sync.time()[0].toLong()
    } catch (_: Exception) {
        System.currentTimeMillis() / 1000
    }
}
