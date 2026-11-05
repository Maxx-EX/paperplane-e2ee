// SPDX-License-Identifier: AGPL-3.0
// Paperplane XDP packet filter: drop banned source IPs before the socket layer.
// Build: clang -target bpf -O2 -g -c xdp_filter.c -o xdp_filter.bpf.o
// Requires CONFIG_BPF_LSM, XDP socket support; run on a real NIC (not veth).
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <bpf/bpf_helpers.h>

struct lpm_key {
    __u32 prefixlen; /* network byte order length for LPM trie */
    __u32 addr;
};

/* LPM_TRIE: longest-prefix match ban table, populated by the loader. */
struct {
    __uint(type, BPF_MAP_TYPE_LPM_TRIE);
    __uint(max_entries, 1024);
    __uint(key_size, sizeof(struct lpm_key));
    __uint(value_size, sizeof(__u64));
    __uint(map_flags, BPF_F_NO_PREALLOC);
} ban_map SEC(".maps");

/* Per-CPU drop counter for observability. */
struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __uint(max_entries, 1);
    __uint(key_size, sizeof(__u32));
    __uint(value_size, sizeof(__u64));
} stats_map SEC(".maps");

static __always_inline int is_banned(__u32 addr) {
    struct lpm_key key = {
        .prefixlen = 32,
        .addr = addr, /* already in network byte order */
    };
    return bpf_map_lookup_elem(&ban_map, &key) != NULL;
}

SEC("xdp")
int xdp_paperplane_filter(struct xdp_md *ctx) {
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;

    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end)
        return XDP_PASS;

    /* UDP/TCP/ICMP v4 only; everything else bypasses. */
    if (eth->h_proto != __constant_htons(ETH_P_IP))
        return XDP_PASS;

    struct iphdr *ip = (void *)(eth + 1);
    if ((void *)(ip + 1) > data_end)
        return XDP_PASS;

    if (is_banned(ip->saddr)) {
        __u32 idx = 0;
        __u64 *cnt = bpf_map_lookup_elem(&stats_map, &idx);
        if (cnt)
            __sync_fetch_and_add(cnt, 1);
        return XDP_DROP;
    }
    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
