package repository

import (
	"blockchainEtheriumGRPC/internal/blockchain"
	"blockchainEtheriumGRPC/proto"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func StopServer() {
	grpcServer.GracefulStop()
}

type WalletRepository interface {
	GetWallet(name string) (*proto.Message, error)
	CreateWallet(wallet *proto.WalletCreate) (*proto.Message, error)
	SendMoneyWallet(money *proto.SendMoney) (*proto.Message, error)
}

type WalletRepositoryImpl struct {
}

func NewWalletRepository() WalletRepository {
	return &WalletRepositoryImpl{}
}

func (w *WalletRepositoryImpl) GetWallet(name string) (*proto.Message, error) {
	walletName, balance := blockchain.ReadWalletSmartContract(name)
	result := "Wallet name: " + walletName + ", balance: " + balance
	return &proto.Message{Body: result}, nil
}

func (w *WalletRepositoryImpl) CreateWallet(wallet *proto.WalletCreate) (*proto.Message, error) {
	result := blockchain.CreateWalletSmartContract(wallet.WalletName, wallet.Balance)
	return &proto.Message{Body: result}, nil
}

func (w *WalletRepositoryImpl) SendMoneyWallet(money *proto.SendMoney) (*proto.Message, error) {
	return &proto.Message{Body: blockchain.SendMoneySnartContract(money.WalletSender, money.WalletRecipient, money.SendMoney)}, nil
}
