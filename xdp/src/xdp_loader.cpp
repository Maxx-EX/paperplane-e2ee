// SPDX-License-Identifier: AGPL-3.0
// Paperplane XDP loader: RAII ownership over libbpf objects, explicit detach on
// SIGTERM/SIGINT, and a periodic diff of /var/lib/paperplane/bans.txt into the
// LPM_TRIE map. Build: g++ -O2 -std=c++17 xdp_loader.cpp -lbpf -lelf -o xdp_loader
#include <atomic>
#include <csignal>
#include <cstdint>
#include <cstdio>
#include <cstring>
#include <fstream>
#include <iostream>
#include <net/if.h>
#include <string>
#include <thread>
#include <unordered_map>
#include <vector>

#include <bpf/bpf.h>
#include <bpf/libbpf.h>
#include <arpa/inet.h>

namespace {
std::atomic<bool> g_stop{false};
int g_bpf_obj_fd = -1;
int g_prog_fd = -1;

void on_signal(int) {
    g_stop = true;
}

// RAII guard for bpf_object: closes handle on scope exit.
class BpfObject {
  public:
    explicit BpfObject(bpf_object *o) : obj_(o) {}
    ~BpfObject() { if (obj_) bpf_object__close(obj_); }
    BpfObject(const BpfObject &) = delete;
    BpfObject &operator=(const BpfObject &) = delete;
    bpf_object *get() { return obj_; }

  private:
    bpf_object *obj_ = nullptr;
};

bool parseBanFile(const std::string &path, std::unordered_map<uint32_t, time_t> *out) {
    std::ifstream in(path);
    if (!in) return false;
    out->clear();
    std::string line;
    while (std::getline(in, line)) {
        if (line.empty() || line[0] == '#') continue;
        auto sp = line.find(' ');
        std::string ip = line.substr(0, sp);
        time_t until = 0;
        if (sp != std::string::npos) until = std::stoll(line.substr(sp + 1));
        in_addr addr{};
        if (inet_pton(AF_INET, ip.c_str(), &addr) != 1) continue;
        (*out)[addr.s_addr] = until;
    }
    return true;
}
}  // namespace

int main(int argc, char **argv) {
    if (argc < 3) {
        std::fprintf(stderr,
                     "usage: %s <iface> <bans.txt> [poll_seconds]\n"
                     "  -d   drop mode (default skb/generic fallback)\n",
                     argv[0]);
        return 1;
    }
    std::string iface = argv[1];
    std::string banFile = argv[2];
    int poll = argc > 3 ? std::atoi(argv[3]) : 5;

    std::signal(SIGINT, on_signal);
    std::signal(SIGTERM, on_signal);
    std::signal(SIGHUP, on_signal);

    int ifindex = if_nametoindex(iface.c_str());
    if (ifindex == 0) {
        std::fprintf(stderr, "interface %s not found\n", iface.c_str());
        return 1;
    }

    // Load BPF object.
    BpfObject obj(bpf_object__open("xdp_filter.bpf.o"));
    if (!obj.get()) {
        std::fprintf(stderr, "bpf_object__open failed\n");
        return 1;
    }
    if (bpf_object__load(obj.get())) {
        std::fprintf(stderr, "bpf_object__load failed (missing libbpf headers?)\n");
        return 1;
    }
    bpf_program *prog = bpf_object__find_program_by_name(obj.get(), "xdp_paperplane_filter");
    if (!prog) {
        std::fprintf(stderr, "program not found in object\n");
        return 1;
    }
    g_prog_fd = bpf_program__fd(prog);

    // Attach: try generic/skb first (safe on virtio/QEMU), then native.
    int flags = XDP_FLAGS_SKB_MODE;
    if (bpf_xdp_attach(ifindex, g_prog_fd, flags, nullptr) < 0) {
        std::fprintf(stderr, "skb attach failed, trying drv mode\n");
        flags = XDP_FLAGS_DRV_MODE;
        if (bpf_xdp_attach(ifindex, g_prog_fd, flags, nullptr) < 0) {
            std::perror("bpf_xdp_attach");
            return 1;
        }
    }
    std::printf("xdp attached on %s (flags=%d)\n", iface.c_str(), flags);

    int banMap = bpf_object__find_map_fd_by_name(obj.get(), "ban_map");

    // Periodic sync: diff file content against the BPF map.
    std::unordered_map<uint32_t, time_t> current;
    while (!g_stop) {
        std::unordered_map<uint32_t, time_t> wanted;
        if (parseBanFile(banFile, &wanted)) {
            for (const auto &[ip, until] : wanted) {
                if (current.find(ip) == current.end()) {
                    if (bpf_map_update_elem(banMap, &(struct { uint32_t p = 32; uint32_t a; }){32, ip},
                                            &(uint64_t{1}), BPF_ANY)) {
                        /* ignore transient errors */
                    }
                }
            }
            for (const auto &[ip, _] : current) {
                if (wanted.find(ip) == wanted.end()) {
                    struct { uint32_t p = 32; uint32_t a; } key{32, ip};
                    bpf_map_delete_elem(banMap, &key);
                }
            }
            current = wanted;
        }
        std::this_thread::sleep_for(std::chrono::seconds(poll));
    }

    // Explicit detach on shutdown.
    if (bpf_xdp_detach(ifindex, flags, nullptr) < 0) {
        std::perror("bpf_xdp_detach");
    } else {
        std::printf("xdp detached cleanly\n");
    }
    return 0;
}
