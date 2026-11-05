package io.paperplane.policy

import org.yaml.snakeyaml.Yaml
import java.io.File

/**
 * 单一配置源：所有字段都由 scripts/render-config.sh 从 config/config.yaml 生成。
 * 本类不允许出现硬编码阈值。
 */
data class AppConfig(
    val server: Server = Server(),
    val redis: Redis = Redis(),
    val policy: Policy = Policy(),
    val crowdsec: CrowdSec = CrowdSec(),
) {
    data class Server(val healthPort: Int = 8091)
    data class Redis(val host: String = "127.0.0.1", val port: Int = 6379,
                     val password: String = "", val db: Int = 0)
    data class Policy(
        val stream: String = "paperplane:events",
        val consumerGroup: String = "paperplane-policy",
        val consumerName: String = "policy-1",
        val windowSeconds: Long = 60,
        val ipAccountRatio: Int = 8,
        val reconnectLimit: Int = 5,
        val reconnectWindowSeconds: Long = 30,
        val banSeconds: Long = 600,
        val failOpenTimeoutMs: Long = 50,
    )
    data class CrowdSec(
        val lapiUrl: String = "http://127.0.0.1:8081",
        val apiKeyEnv: String = "CROWDSEC_LAPI_KEY",
        val decisionDuration: String = "10m",
    )

    companion object {
        @Suppress("UNCHECKED_CAST")
        fun load(path: String): AppConfig {
            val raw = Yaml().load<Map<String, Any>>(File(path).inputStream()) ?: emptyMap()
            fun section(name: String) = (raw[name] as? Map<String, Any>) ?: emptyMap()
            fun Map<String, Any>.str(k: String, d: String) = this[k] as? String ?: d
            fun Map<String, Any>.long(k: String, d: Long) = (this[k] as? Number)?.toLong() ?: d
            fun Map<String, Any>.int(k: String, d: Int) = (this[k] as? Number)?.toInt() ?: d

            val r = section("redis"); val p = section("policy"); val c = section("crowdsec")
            return AppConfig(
                redis = Redis(r.str("host", "127.0.0.1"), r.int("port", 6379),
                    r.str("password", ""), r.int("database", r.int("db", 0))),
                policy = Policy(
                    p.str("stream", "paperplane:events"),
                    p.str("consumer_group", "paperplane-policy"),
                    p.str("consumer_name", "policy-1"),
                    p.long("window_seconds", 60),
                    p.int("ip_account_ratio", 8),
                    p.int("reconnect_limit", 5),
                    p.long("reconnect_window_seconds", 30),
                    p.long("ban_seconds", 600),
                    p.long("fail_open_timeout_ms", 50),
                ),
                crowdsec = CrowdSec(
                    c.str("lapi_url", "http://127.0.0.1:8081"),
                    c.str("api_key_env", "CROWDSEC_LAPI_KEY"),
                    c.str("decision_duration", "10m"),
                ),
            )
        }
    }
}
