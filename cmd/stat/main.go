package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type StatService struct {
	ctx context.Context
	wg  *sync.WaitGroup
	mu  sync.Mutex
	cnt int
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	ss := StatService{
		ctx: ctx,
		wg:  wg,
	}

	ss.wg.Add(1)
	go ss.Cycle()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	cancel()

	ss.wg.Wait()

}

func (s *StatService) Cycle() {

	defer s.wg.Done()

	for {
		select {
		case <-time.After(500 * time.Millisecond):
			s.mu.Lock()
			s.cnt++
			s.mu.Unlock()
		case <-s.ctx.Done():
			log.Printf("Context canceled\n")
			fmt.Printf("Count: %d", s.cnt)
			return
		}
	}

}
