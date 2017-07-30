package main

import (
	"github.com/ingojaeckel/go-udp/client"
	"github.com/ingojaeckel/go-udp/server"
	"time"
)

func main() {
	clientIp := "127.0.0.1"
	serverIp := "127.0.0.1"
	serverPort := 10001

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
