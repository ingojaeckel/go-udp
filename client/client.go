package client

import (
	"fmt"
	"net"
)

func New(clientName string, initialMessage []byte) {
	ServerAddr, _ := net.ResolveUDPAddr("udp","127.0.0.1:10001")
	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	Conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)

	defer Conn.Close()

	Conn.Write(initialMessage)

	for {
		readBuf := make([]byte, 10)
		numberOfBytesRead, _ := Conn.Read(readBuf)
		fmt.Printf("[%s] Received %d bytes: %v\n", clientName, numberOfBytesRead, readBuf[:numberOfBytesRead])
	}
}
