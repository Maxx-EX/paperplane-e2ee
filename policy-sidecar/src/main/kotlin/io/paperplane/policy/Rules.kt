package io.paperplane.policy

import io.lettuce.core.api.sync.RedisCommands
import org.slf4j.LoggerFactory

/**
 * 跨维度规则引擎（纯协程结构化并发，每条消息独立异常隔离）。
 * 时钟统一来自 Redis TIME，避免本地时钟漂移（K-8）。
 */
class Rules(
    private val cfg: AppConfig,
    private val sync: RedisCommands<String, String>,
    private val bans: BanPublisher,
) {
    private val log = LoggerFactory.getLogger("rules")

    fun handle(fields: Map<String, String>) {
        val kind = fields["kind"] ?: return
        val ip = fields["ip"].orEmpty()
        val account = fields["account"].orEmpty()
        if (ip.isEmpty() && account.isEmpty()) return
        val nowMs = redisNowMs()
        when (kind) {
            "connect", "envelope" -> {
                if (ip.isNotEmpty() && account.isNotEmpty()) trackIpAccount(ip, account, nowMs)
                if (kind == "connect") trackReconnects(account)
            }
            "register", "auth_fail" -> {
                if (ip.isNotEmpty()) trackIpAccount(ip, account.ifEmpty { "anon" }, nowMs)
            }
        }
    }

    /** 同一 IP 短时聚合异常数量的账号：代理/批量注册特征。 */
    private fun trackIpAccount(ip: String, account: String, nowMs: Long) {
        val key = "paperplane:ipaccts:$ip"
        sync.zadd(key, nowMs.toDouble(), account)
        sync.zremrangebyscore(key, 0.0, (nowMs - cfg.policy.windowSeconds * 1000).toDouble())
        val uniq = sync.zcard(key)
        sync.expire(key, cfg.policy.windowSeconds + 10)
        if (uniq > cfg.policy.ipAccountRatio) {
            log.warn("ip_account_ratio exceeded ip={} uniq={}", ip, uniq)
            bans.banIP(ip, "ip_account_ratio")
        }
    }

    /** 同一账号高频重连：重连攻击/网络层抖动放大。 */
    private fun trackReconnects(account: String) {
        val key = "paperplane:acct_reconnect:$account"
        val n = sync.incr(key)
        sync.expire(key, cfg.policy.reconnectWindowSeconds)
        if (n > cfg.policy.reconnectLimit) {
            log.warn("reconnect storm account={} count={}", account, n)
            bans.banAccount(account, "reconnect_storm")
        }
    }

    /** Redis TIME -> epoch millis（K-8 统一时钟源）。 */
    private fun redisNowMs(): Long {
        return try {
            val t = sync.time()
            t[0].toLong() * 1000 + t[1].toLong() / 1000
        } catch (_: Exception) {
            System.currentTimeMillis()
        }
    }
}
