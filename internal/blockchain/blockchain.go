package blockchain

import (
	"blockchainEtheriumGRPC/binding"
	"blockchainEtheriumGRPC/conf"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var pathConfig = "conf/"
var gateway = conf.ViperEnvVariable("gateway", pathConfig)
var accountPrivateKey = conf.ViperEnvVariable("accountPrivateKey", pathConfig)
var accountHexAddress = conf.ViperEnvVariable("accountHexAddress", pathConfig)
var smartContractHexAddress = conf.ViperEnvVariable("smartContractHexAddress", pathConfig)

var ClientBlockchain *ethclient.Client

func ReadWalletSmartContract(nameWallet string) (string, string) {
	client, address := ConnectContract(smartContractHexAddress)
	instance, err := binding.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}

	getWallet, err := instance.GetWallet(nil, nameWallet)
	if err != nil {
		log.Fatal(err)
	}

	return getWallet.WalletName, getWallet.Balance.String()
}

func CreateWalletSmartContract(nameWallet string, balance int64) string {

	auth, instance := ConnectChangeContract(smartContractHexAddress, accountPrivateKey)
	tx, err := instance.SetWallet(auth, nameWallet, big.NewInt(balance))
	if err != nil {
		log.Fatal(err)
	}

	return tx.Hash().Hex()
}

func SendMoneySnartContract(nameWalletSender string, nameWalletRecipient string, sendMoney int64) string {
	auth, instance := ConnectChangeContract(smartContractHexAddress, accountPrivateKey)
	tx, err := instance.SendMoney(auth, nameWalletSender, nameWalletRecipient, big.NewInt(sendMoney))
	if err != nil {
		log.Fatal(err)
	}

	return tx.Hash().Hex()

}

func ConnectContract(smartContractHexAddressInput string) (*ethclient.Client, common.Address) {
	if ClientBlockchain == nil {
		ClientBlockchain = getClient()
	}
	address := common.HexToAddress(smartContractHexAddressInput)

	return ClientBlockchain, address
}

func ConnectChangeContract(smartContractHexAddress string, accountPrivateKey string) (*bind.TransactOpts, *binding.Api) {

	if ClientBlockchain == nil {
		ClientBlockchain = getClient()
	}

	privateKey, err := crypto.HexToECDSA(accountPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ClientBlockchain.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		ClientBlockchain = nil
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(3000000)        // in units
	auth.GasPrice = big.NewInt(1000000000) // in wei

	address := common.HexToAddress(smartContractHexAddress)
	instance, err := binding.NewApi(address, ClientBlockchain)
	if err != nil {
		log.Fatal(err)
		ClientBlockchain = nil
	}
	return auth, instance
}

func getClient() *ethclient.Client {
	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)

	}
	return client
}
