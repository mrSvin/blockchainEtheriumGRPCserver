package blockchain

import (
	"blockchainEtheriumGRPC/binding"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strconv"
	"testing"
)

func TestSetWallet(t *testing.T) {
	auth, instance := ConnectChangeContract(smartContractHexAddress, accountPrivateKey)
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

	auth, instance := ConnectChangeContract(smartContractHexAddress, accountPrivateKey)

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

func TestDeploySmartContract(t *testing.T) {
	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(3000000)        // in units
	auth.GasPrice = big.NewInt(1000000000) // in wei

	address, tx, _, err := binding.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("smartContractHexAddress: ", address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println("tx: ", tx.Hash().Hex())                    // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

}
