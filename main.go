package main

import (
	"github.com/ingojaeckel/go-udp/client"
	"github.com/ingojaeckel/go-udp/server"
	"time"
)

func main() {
	initialMessage1 := []byte{23}
	initialMessage2 := []byte{42}

	go server.New()
	time.Sleep(1 * time.Second)

	go client.New("Alice", initialMessage1)
	time.Sleep(5 * time.Second)

	go client.New("Bob", initialMessage1)
	time.Sleep(5 * time.Second)

	go client.New("Carol", initialMessage2)
	time.Sleep(5 * time.Second)

	go client.New("Dan", initialMessage2)
	time.Sleep(5 * time.Second)
}
