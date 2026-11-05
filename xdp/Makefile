# Paperplane XDP filter build
# Requirements: clang (bpf target), libbpf-dev, bpftool, libelf-dev, root/CAP_BPF
CLANG   ?= clang
BPFTOOL ?= bpftool
CC      ?= g++
CFLAGS  ?= -g -O2 -Wall
BPF_CFLAGS ?= -target bpf -D__TARGET_ARCH_x86 -O2 -g -Wall -Werror

all: xdp_filter.bpf.o xdp_loader

xdp_filter.bpf.o: src/xdp_filter.c
	$(CLANG) $(BPF_CFLAGS) -c $< -o $@

xdp_loader: src/xdp_loader.cpp xdp_filter.bpf.o
	$(CC) $(CFLAGS) -std=c++17 $< -lbpf -lelf -o $@

vmlinux:
	$(BPFTOOL) btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h

check:
	bpftool feature probe | grep -q "eXtended DataPath" || echo "XDP not available on this host"

clean:
	rm -f xdp_filter.bpf.o xdp_loader vmlinux.h

.PHONY: all clean check
