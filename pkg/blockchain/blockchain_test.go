package blockchain

import (
	"blockchainEtheriumGRPC/api"
	"log"
	"math/big"
	"strconv"
	"testing"
)

func TestCheckContract(t *testing.T) {
	got := CheckContract(smartContractHexAddress)
	want := true

	if want != got {
		t.Errorf("Expected True, but got False")
	}
}

func TestNewApi(t *testing.T) {
	client, address := ConnectContract(gateway, smartContractHexAddress)
	_, err := api.NewApi(address, client)
	if err != nil {
		t.Errorf("Error get NewApi")
	}
}

func TestNewApiCaller(t *testing.T) {
	client, address := ConnectContract(gateway, smartContractHexAddress)
	_, err := api.NewApiCaller(address, client)
	if err != nil {
		t.Errorf("Error get NewApiCaller")
	}
}

func TestNewApiTransactor(t *testing.T) {
	client, address := ConnectContract(gateway, smartContractHexAddress)
	_, err := api.NewApiTransactor(address, client)
	if err != nil {
		t.Errorf("Error get NewApiTransactor")
	}
}

func TestNewApiFilterer(t *testing.T) {
	client, address := ConnectContract(gateway, smartContractHexAddress)
	_, err := api.NewApiFilterer(address, client)
	if err != nil {

	}
}

func TestGetWallet(t *testing.T) {
	nameVallet := "one"
	client, address := ConnectContract(gateway, smartContractHexAddress)
	instance, err := api.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}

	getWallet, err := instance.GetWallet(nil, nameVallet)
	if err != nil {
		log.Fatal(err)
	}
	if nameVallet != getWallet.WalletName {
		t.Errorf("Error don't wanted Wallet with name %q", nameVallet)
	}

}

func TestWallets(t *testing.T) {
	nameVallet := "one"
	client, address := ConnectContract(gateway, smartContractHexAddress)
	instance, err := api.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}
	wallet, err := instance.Wallets(nil, nameVallet)
	if nameVallet != wallet.WalletName {
		t.Errorf("Error don't wanted Wallet with name %q", nameVallet)
	}
}

func TestSetWallet(t *testing.T) {
	auth, instance := ConnectChangeContract(gateway, smartContractHexAddress, accountPrivateKey)
	tx, err := instance.SetWallet(auth, "five", big.NewInt(50000))
	if err != nil {
		log.Fatal(err)
	}
	if len(tx.Hash().Hex()) != 66 {
		t.Errorf("Error get SetWallet")
	}
}

func TestSendMoney(t *testing.T) {

	nameWalletSender := "one"
	nameWalletRecipient := "two"
	sendMoney := big.NewInt(20)

	auth, instance := ConnectChangeContract(gateway, smartContractHexAddress, accountPrivateKey)

	getWalletOne, err := instance.GetWallet(nil, nameWalletSender)
	getWalletTwo, err := instance.GetWallet(nil, nameWalletRecipient)
	var beforeMoneyWalletOne *big.Int
	beforeMoneyWalletOne = getWalletOne.Balance
	beforeMoneyWalletTwo := getWalletTwo.Balance
	_, err = instance.SendMoney(auth, nameWalletSender, nameWalletRecipient, sendMoney)
	if err != nil {
		log.Fatal(err)
	}

	getWalletOne, err = instance.GetWallet(nil, nameWalletSender)
	getWalletTwo, err = instance.GetWallet(nil, nameWalletRecipient)
	afterMoneyWalletOne := getWalletOne.Balance
	afterMoneyWalletTwo := getWalletTwo.Balance

	expectBeforeMoneyWalletOne, err := strconv.Atoi(beforeMoneyWalletOne.String())
	expectBeforeMoneyWalletTwo, err := strconv.Atoi(beforeMoneyWalletTwo.String())
	expectAfterMoneyWalletOne, err := strconv.Atoi(afterMoneyWalletOne.String())
	expectAfterMoneyWalletTwo, err := strconv.Atoi(afterMoneyWalletTwo.String())

	if expectBeforeMoneyWalletOne-20 != expectAfterMoneyWalletOne && expectBeforeMoneyWalletTwo+20 != expectAfterMoneyWalletTwo {
		t.Errorf("Error SendMoney, beforeMoneyWalletOne: %q, beforeMoneyWalletTwo: %q, afterMoneyWalletOne: %q, afterMoneyWalletTwo: %q ", beforeMoneyWalletOne, beforeMoneyWalletTwo, afterMoneyWalletOne, afterMoneyWalletTwo)
	}

}
