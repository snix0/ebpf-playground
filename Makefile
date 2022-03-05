TARGETS := mmap_monitor
TARGETS += mmap_monitor.bpf.o

all: $(TARGETS)

mmap_monitor: mmap_monitor.go
	CC=clang CGO_FLAGS="-I/usr/include/bpf -I/usr/include/linux" CGO_LDFLAGS="/usr/lib64/libbpf.so" go build -o mmap_monitor

mmap_monitor.bpf.o: mmap_monitor.bpf.c
	clang -I /usr/include/linux/ -target bpf -o mmap_monitor.bpf.o mmap_monitor.bpf.c -O2 -c -g

clean:
	rm -f $(TARGETS)
