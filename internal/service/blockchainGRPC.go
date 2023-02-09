package service

import (
	"blockchainEtheriumGRPC/conf"
	"blockchainEtheriumGRPC/internal/blockchain"
	"blockchainEtheriumGRPC/internal/repository"
	"blockchainEtheriumGRPC/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func StartGrpcServer() error {
	fmt.Println("gRPC Server")
	pathConfig := "conf/"
	portString := conf.ViperEnvVariable("portGrpc", pathConfig)
	port, _ := strconv.ParseInt(portString, 10, 64)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterGrpcServiceServer(grpcServer, repository.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return err
}

func ShutdownGrpcServer() {
	log.Println("shutting down...")
	repository.StopServer()
	blockchain.ClientBlockchain.Close()
}
