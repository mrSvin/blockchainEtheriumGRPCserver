package blockchain

import (
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
