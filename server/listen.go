package server

import (
	"bytes"
	"log"
	"net"
)

var (
	ip         = "0.0.0.0"
	port       = 6969
	udpNetwork = "udp"
	bufferSize = 1024
)

func Listen(ch chan []byte) error {
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: port}

	lstnr, err := net.ListenUDP(udpNetwork, &addr)
	if err != nil {
		log.Fatalf("create listener: %v", err)
	}
	defer lstnr.Close()

	buffer := make([]byte, bufferSize)

	for {
		n, _, err := lstnr.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("read udp data: %v", err)
		}
		ch <- bytes.TrimSpace(buffer[:n])
	}

}
