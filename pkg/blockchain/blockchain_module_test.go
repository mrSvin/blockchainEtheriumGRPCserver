package blockchain

import (
	"blockchainEtheriumGRPC/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strconv"
	"testing"
)

func TestDeployApi(t *testing.T) {
	address, _, tx, _, _ := createSimulateContract()
	if len(address.Hex()) != 42 && len(tx.Hash().Hex()) != 66 {
		t.Errorf("Error get DeployApi")
	}
}

func TestModuleCreateWallet(t *testing.T) {
	address, _, tx, auth, instance := createSimulateContract()
	if len(address.Hex()) != 42 && len(tx.Hash().Hex()) != 66 {
		t.Errorf("Error get DeployApi")
	}
	tx, err := instance.SetWallet(auth, "five", big.NewInt(50000))
	if err != nil {
		log.Fatal(err)
	}
	if len(tx.Hash().Hex()) != 66 {
		t.Errorf("Error get SetWallet")
	}
}

func TestModuleGetWallet(t *testing.T) {
	nameWallet := "one"
	balance := big.NewInt(30000)
	_, client, _, auth, instance := createSimulateContract()
	_, err := instance.SetWallet(auth, nameWallet, balance)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()
	getWallet, err := instance.GetWallet(nil, nameWallet)
	if err != nil {
		log.Fatal(err)
	}

	if getWallet.Balance.String() != balance.String() && getWallet.WalletName != nameWallet {
		t.Errorf("Error get GetWallet")
	}
}

func TestModuleSendMoney(t *testing.T) {
	nameSenderWallet := "one"
	nameRecipientWallet := "two"
	balanceSender := big.NewInt(30000)
	balanceRecipien := big.NewInt(20000)
	moneySend := big.NewInt(300)
	address, client, tx, auth, instance := createSimulateContract()
	if len(address.Hex()) != 42 && len(tx.Hash().Hex()) != 66 {
		t.Errorf("Error get DeployApi")
	}
	tx, err := instance.SetWallet(auth, nameSenderWallet, balanceSender)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()

	tx, err = instance.SetWallet(auth, nameRecipientWallet, balanceRecipien)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()

	getWalletSender, err := instance.GetWallet(nil, nameSenderWallet)
	if err != nil {
		log.Fatal(err)
	}

	getWalletRecipient, err := instance.GetWallet(nil, nameRecipientWallet)
	if err != nil {
		log.Fatal(err)
	}

	tx, err = instance.SendMoney(auth, nameSenderWallet, nameRecipientWallet, moneySend)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()

	getWalletSender, err = instance.GetWallet(nil, nameSenderWallet)
	if err != nil {
		log.Fatal(err)
	}

	getWalletRecipient, err = instance.GetWallet(nil, nameRecipientWallet)
	if err != nil {
		log.Fatal(err)
	}

	oldGetWalletSender, err := strconv.Atoi(balanceSender.String())
	newGetWalletSender, err := strconv.Atoi(getWalletSender.Balance.String())
	oldGetWalletRecipient, err := strconv.Atoi(balanceRecipien.String())
	newGetWalletRecipient, err := strconv.Atoi(getWalletRecipient.Balance.String())

	if newGetWalletSender+300 != oldGetWalletSender && newGetWalletRecipient-300 != oldGetWalletRecipient {
		t.Errorf("Error get GetWallet SendMoney")
	}

}

func createSimulateContract() (common.Address, *backends.SimulatedBackend, *types.Transaction, *bind.TransactOpts, *api.Api) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	address, tx, _, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	instance, err := api.NewApi(address, client)

	return address, client, tx, auth, instance
}
