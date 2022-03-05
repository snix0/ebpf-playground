# Overview

This program is a very basic eBPF program which monitors for uses of the mmap syscall and prints a simple message when such a usage is detected.

# Building and Running

In order to build this project, you will need to generate the vmlinux.h file which contains all the type defintions that your running Linux kernel uses in its own source code.

This file can be generated using the command below:

```
bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```

Simply compile the program using make and run the program with sudo.

Note: You may need to update your include paths to match your Linux distribution.
