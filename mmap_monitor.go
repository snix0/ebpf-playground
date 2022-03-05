package main

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"

	bpf "github.com/aquasecurity/libbpfgo"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	fmt.Println("mmap monitor starting...")
	b, err := bpf.NewModuleFromFile("mmap_monitor.bpf.o")
	defer b.Close()

	must(b.BPFLoadObject())

	p, err := b.GetProgram("kprobe__sys_mmap")
	must(err)

	_, err = p.AttachKprobe("__x64_sys_mmap")
	must(err)

	eventsChannel := make(chan []byte)
	rb, err := b.InitRingBuf("events", eventsChannel)
	must(err)

	rb.Start()
	for {
		eventBytes := <-eventsChannel
		pid := int(binary.LittleEndian.Uint32(eventBytes[0:4]))
		comm := string(bytes.TrimRight(eventBytes[4:], "\x00"))

		fmt.Printf("%d %v\n", pid, comm)
	}

	<-sig
	fmt.Println("Cleaning up")
}

func must(err error) {
	if err != nil {

		panic(err)
	}
}
