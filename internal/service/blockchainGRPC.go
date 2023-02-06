package service

import (
	"blockchainEtheriumGRPC/api/proto"
	"blockchainEtheriumGRPC/internal/blockchain"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	proto.UnimplementedGrpcServiceServer
}

func (s *Server) GetWallet(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	log.Printf("get wallet with name: %s", in.Body)

	walletName, balance := blockchain.ReadWalletSmartContract(in.Body)
	result := "Wallet name: " + walletName + ", balance: " + balance
	return &proto.Message{Body: result}, nil
}

func (s *Server) CreateWallet(ctx context.Context, wallet *proto.WalletCreate) (*proto.Message, error) {
	log.Printf("create wallet with name: %s, balance: %d", wallet.WalletName, wallet.Balance)
	return &proto.Message{Body: blockchain.CreateWalletSmartContract(wallet.WalletName, wallet.Balance)}, nil
}

func (s *Server) SendMoneyWallet(ctx context.Context, money *proto.SendMoney) (*proto.Message, error) {
	log.Printf("transfer of wallet named: %s, to wallet named: %s, send money: %d", money.WalletSender, money.WalletRecipient, money.SendMoney)
	return &proto.Message{Body: blockchain.SendMoneySnartContract(money.WalletSender, money.WalletRecipient, money.SendMoney)}, nil
}
