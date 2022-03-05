hello: hello.go
	CC=clang CGO_FLAGS="-I/usr/include/bpf -I/usr/include/linux" CFLAGS="-Wno-deprecated-declarations" CGO_LDFLAGS="/usr/lib64/libbpf.so" go build -o hello

hello.bpf.o: hello.bpf.c
	clang -I /usr/include/linux/ -target bpf -o hello.bpf.o hello.bpf.c -O2 -c -g
