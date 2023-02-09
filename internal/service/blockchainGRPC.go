package service

import (
	"blockchainEtheriumGRPC/conf"
	"blockchainEtheriumGRPC/internal/blockchain"
	"blockchainEtheriumGRPC/internal/repository"
	"blockchainEtheriumGRPC/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
	proto.UnimplementedGrpcServiceServer
	walletRepository repository.WalletRepository
}

func NewServer(walletRepository repository.WalletRepository) *Server {
	return &Server{walletRepository: walletRepository}
}

func (s *Server) GetWallet(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	log.Printf("get wallet with name: %s", in.Body)
	return s.walletRepository.GetWallet(in.Body)
}

func (s *Server) CreateWallet(ctx context.Context, wallet *proto.WalletCreate) (*proto.Message, error) {
	log.Printf("create wallet with name: %s, balance: %d", wallet.WalletName, wallet.Balance)
	return s.walletRepository.CreateWallet(wallet)
}

func (s *Server) SendMoneyWallet(ctx context.Context, money *proto.SendMoney) (*proto.Message, error) {
	log.Printf("transfer of wallet named: %s, to wallet named: %s, send money: %d", money.WalletSender, money.WalletRecipient, money.SendMoney)
	return s.walletRepository.SendMoneyWallet(money)
}

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
	proto.RegisterGrpcServiceServer(grpcServer, NewServer(repository.NewWalletRepository()))
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
