package server

import (
	"net"
	"github.com/ingojaeckel/go-udp/common"
	"fmt"
	"time"
)

var timeBetweenServerMessages = 1000 * time.Millisecond

func New() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
	common.CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	common.CheckError(err)
	defer ServerConn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, _ := ServerConn.ReadFromUDP(buf)

		go handleNewClient(buf, n, addr, ServerConn)
	}

}

func handleNewClient(buf []byte, n int, addr *net.UDPAddr, srvConn *net.UDPConn) {
	fmt.Printf("[srv] Handling new connection. Received %d bytes from client:%v; addr:%v\n", n, buf[:n], addr)

	// Create a copy of the bytes originally received by the client.
	packet := make([]byte, n+1)
	for i:=0; i<n; i++ {
		packet[i] = buf[i]
	}
	// This will initialize the sequence number with 1.
	packet[n] = 1

	for i := 0; ; i++ {
		srvConn.WriteTo(packet, addr)
		time.Sleep(timeBetweenServerMessages)

		// Send n+1 bytes to this client.
		// The fist n bytes will be the original message the client sent.
		// The last byte will be a sequence number starting at 1.
		packet[n] = packet[n] + 1
	}
}