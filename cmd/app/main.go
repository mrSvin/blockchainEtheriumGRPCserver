package main

import (
	"blockchainEtheriumGRPC/api/proto"
	"blockchainEtheriumGRPC/conf"
	"blockchainEtheriumGRPC/internal/blockchain"
	"blockchainEtheriumGRPC/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var grpcServer *grpc.Server

func main() {
	errChan := make(chan error)
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := startGrpcServer(); err != nil {
			errChan <- err
		}
	}()

	defer shutdownGrpcServer()

	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case <-stopChan:
	}

}

func startGrpcServer() error {
	fmt.Println("gRPC Server")
	pathConfig := "conf/"
	portString := conf.ViperEnvVariable("portGrpc", pathConfig)
	port, _ := strconv.ParseInt(portString, 10, 64)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := service.Server{}
	grpcServer = grpc.NewServer()
	proto.RegisterGrpcServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return err
}

func shutdownGrpcServer() {
	log.Println("shutting down...")
	grpcServer.GracefulStop()
	blockchain.ClientBlockchain.Close()
}
