package client

import (
	"fmt"
	"net"
)

func New(clientName string, initialMessage []byte, serverIp string, serverPort int, clientIp string) {
	ServerAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	LocalAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:0", clientIp))
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
