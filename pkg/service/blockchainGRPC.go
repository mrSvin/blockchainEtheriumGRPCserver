package service

import (
	"blockchainEtheriumGRPC/pkg/blockchain"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedGrpcServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetWallet(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("get wallet with name: %s", in.Body)

	walletName, balance := blockchain.ReadWalletSmartContract(in.Body)
	result := "Wallet name: " + walletName + ", balance: " + balance
	return &Message{Body: result}, nil
}

func (s *Server) CreateWallet(ctx context.Context, wallet *WalletCreate) (*Message, error) {
	log.Printf("create wallet with name: %s, balance: %d", wallet.WalletName, wallet.Balance)
	return &Message{Body: blockchain.CreateWalletSmartContract(wallet.WalletName, wallet.Balance)}, nil
}

func (s *Server) SendMoneyWallet(ctx context.Context, money *SendMoney) (*Message, error) {
	log.Printf("transfer of wallet named: %s, to wallet named: %s, send money: %d", money.WalletSender, money.WalletRecipient, money.SendMoney)
	return &Message{Body: blockchain.SendMoneySnartContract(money.WalletSender, money.WalletRecipient, money.SendMoney)}, nil
}
