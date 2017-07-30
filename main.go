package main

import (
	"github.com/ingojaeckel/go-udp/client"
	"github.com/ingojaeckel/go-udp/server"
	"time"
	"os"
	"strconv"
	"fmt"
)

func main() {
	clientIp := "127.0.0.1"
	serverIp := "127.0.0.1"
	serverPort := 10001

	if len(os.Args) == 4 {
		serverIp = os.Args[1]
		p, _ := strconv.ParseInt(os.Args[2], 10, 32)
		serverPort = int(p)
		clientIp = os.Args[3]
	}
	fmt.Printf("Using server %s:%d, client: %s\n", serverIp, serverPort, clientIp)

	initialMessage1 := []byte{23}
	initialMessage2 := []byte{42}

	go server.New()
	time.Sleep(1 * time.Second)

	go client.New("Alice", initialMessage1, serverIp, serverPort, clientIp)
	time.Sleep(5 * time.Second)

	go client.New("Bob", initialMessage1, serverIp, serverPort, clientIp)
	time.Sleep(5 * time.Second)

	go client.New("Carol", initialMessage2, serverIp, serverPort, clientIp)
	time.Sleep(5 * time.Second)

	go client.New("Dan", initialMessage2, serverIp, serverPort, clientIp)
	time.Sleep(5 * time.Second)
}
