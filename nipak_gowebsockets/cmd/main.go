package main

import (
	"nipak_gowebsockets/internal/wsserver"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	addr = "localhost:9090"
)

func main() {
	wsServer := wsserver.New(addr)
	log.Info("Started ws server")
	if err := wsServer.Start(); err != nil {
		log.Fatalf("Error with ws server: %v", err)
	}
	log.Error(wsServer.Stop())
	time.Sleep(5 * time.Second)
}
