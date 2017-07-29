package client

import (
	"fmt"
	"net"
	// "time"
	"github.com/ingojaeckel/go-udp/common"
)

func New(clientName string, initialMessage []byte) {
	ServerAddr, err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
	common.CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	common.CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	common.CheckError(err)

	defer Conn.Close()

	Conn.Write(initialMessage)

	for {
		readBuf := make([]byte, 10)
		numberOfBytesRead, _ := Conn.Read(readBuf)
		fmt.Printf("[%s] Received %d bytes: %v\n", clientName, numberOfBytesRead, readBuf[:numberOfBytesRead])

		// time.Sleep(100 * time.Millisecond)
	}
}
