package client

import (
	"fmt"
	"net"
)

func New(clientName string, initialMessage []byte) {
	ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	Conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)

	defer Conn.Close()

	Conn.Write(initialMessage)

	lastSequenceNumber := 0
	packetsReceived := 0
	packetsLost := 0

	for {
		readBuf := make([]byte, 10)
		numberOfBytesRead, _ := Conn.Read(readBuf)

		if numberOfBytesRead > 0 {
			newSequenceNumber := int(readBuf[numberOfBytesRead - 1])

			numberOfPacketsSkipped := newSequenceNumber - lastSequenceNumber - 1
			packetsLost = packetsLost + numberOfPacketsSkipped
			lastSequenceNumber = newSequenceNumber
			packetsReceived = packetsReceived + 1
		}
		fmt.Printf("[%5s] rcvd:%2v [pkts.rcvd:%2d,lost:%2d]\n", clientName, readBuf[:numberOfBytesRead], packetsReceived, packetsLost)
	}
}
