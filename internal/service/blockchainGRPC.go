package service

import (
	"blockchainEtheriumGRPC/conf"
	"blockchainEtheriumGRPC/internal/blockchain"
	"blockchainEtheriumGRPC/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var grpcServer *grpc.Server

type WalletRepository struct{}

func (r *WalletRepository) ReadWallet(walletName string) (string, string) {
	return blockchain.ReadWalletSmartContract(walletName)
}

func (r *WalletRepository) CreateWallet(walletName string, balance int64) string {
	return blockchain.CreateWalletSmartContract(walletName, balance)
}

func (r *WalletRepository) SendMoney(walletSender string, walletRecipient string, sendMoney int64) string {
	return blockchain.SendMoneySnartContract(walletSender, walletRecipient, sendMoney)
}

type Server struct {
	proto.UnimplementedGrpcServiceServer
	walletRepository *WalletRepository
}

func NewServer() *Server {
	return &Server{walletRepository: &WalletRepository{}}
}

func (s *Server) GetWallet(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	log.Printf("get wallet with name: %s", in.Body)

	walletName, balance := blockchain.ReadWalletSmartContract(in.Body)
	result := "Wallet name: " + walletName + ", balance: " + balance
	return &proto.Message{Body: result}, nil
}

func (s *Server) CreateWallet(ctx context.Context, wallet *proto.WalletCreate) (*proto.Message, error) {
	log.Printf("create wallet with name: %s, balance: %d", wallet.WalletName, wallet.Balance)
	result := blockchain.CreateWalletSmartContract(wallet.WalletName, wallet.Balance)
	return &proto.Message{Body: result}, nil
}

func (s *Server) SendMoneyWallet(ctx context.Context, money *proto.SendMoney) (*proto.Message, error) {
	log.Printf("transfer of wallet named: %s, to wallet named: %s, send money: %d", money.WalletSender, money.WalletRecipient, money.SendMoney)
	return &proto.Message{Body: blockchain.SendMoneySnartContract(money.WalletSender, money.WalletRecipient, money.SendMoney)}, nil
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
	proto.RegisterGrpcServiceServer(grpcServer, NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	return err
}

func ShutdownGrpcServer() {
	log.Println("shutting down...")
	grpcServer.GracefulStop()
	blockchain.ClientBlockchain.Close()
}
