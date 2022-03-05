#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

struct process_info {
    int pid;
    char comm[100];
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24);
} events SEC(".maps");

char LICENSE[] SEC("license") = "GPL";
SEC("kprobe/sys_mmap")

int kprobe__sys_mmap(struct pt_regs *ctx) {


    __u64 id = bpf_get_current_pid_tgid();
    __u32 tgid = id >> 32;
    __u64 ringbuffer_flags = 0;

    struct process_info *process;
    process = bpf_ringbuf_reserve(&events, sizeof(struct process_info), ringbuffer_flags);

    if (!process) {
        return 0;
    }

    process->pid = tgid;

    bpf_get_current_comm(&process->comm, 100);
    bpf_ringbuf_submit(process, ringbuffer_flags);

    return 0;
}
