package main

import (
	"blockchainEtheriumGRPC/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errChan := make(chan error)
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := service.StartGrpcServer(); err != nil {
			errChan <- err
		}
	}()

	defer service.ShutdownGrpcServer()

	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case <-stopChan:
	}

}
