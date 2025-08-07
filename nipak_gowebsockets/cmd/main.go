package main

import (
	"log"
	"nipak_gowebsockets/internal/wsserver"
)

const (
	addr = "localhost:9090"
)

func main() {
	wsServer := wsserver.New(addr)
	if err := wsServer.Start(); err != nil {
		log.Fatalf("Error with ws server: %v", err)
	}
}
