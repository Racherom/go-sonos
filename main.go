package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func main() {
	a := 1
	listenAddr, err := net.ResolveUDPAddr("udp", ":1900")
	if err != nil {
		log.Fatal("Resolve list Addr", err)
	}

	a = 2

	list, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		log.Fatal("Resolve listen", err)
	}
	defer list.Close()

	log.Println(a)

	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:1900")
	if err != nil {
		log.Fatal("Resolve Addr: ", err)
	}

	b := new(bytes.Buffer)
	// FIXME: error should be checked.
	b.WriteString("M-SEARCH * HTTP/1.1\r\n")
	fmt.Fprintf(b, "HOST: %s\r\n", raddr.String())
	fmt.Fprintf(b, "NT: %s\r\n", nt)
	fmt.Fprintf(b, "NTS: %s\r\n", "ssdp:alive")
	fmt.Fprintf(b, "USN: %s\r\n", usn)

	_, err = list.WriteToUDP([]byte("Hallo Welt"), addr)
	if err != nil {
		log.Fatal("Resolve Addr: ", err)
	}
}
