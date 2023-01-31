package blockchain

import (
	"blockchainEtheriumGRPC/api"
	"blockchainEtheriumGRPC/conf"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"
)

var pathConfig = "conf/"
var gateway = conf.ViperEnvVariable("gateway", pathConfig)
var accountPrivateKey = conf.ViperEnvVariable("accountPrivateKey", pathConfig)
var accountHexAddress = conf.ViperEnvVariable("accountHexAddress", pathConfig)
var smartContractHexAddress = conf.ViperEnvVariable("smartContractHexAddress", pathConfig)

func CheckContract(contractHexAddress string) bool {

	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress(contractHexAddress)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	//fmt.Printf("is contract: %v\n", isContract) // is contract: true
	return isContract
}

func getBlock() {

	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144

}

func getTransaction() {
	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), big.NewInt(1)); err == nil {
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
	}

	blockHash := common.HexToHash("0xabe9f44bd044f0a74273a254f7b5f99c7202c78c1739dde244f819218e38aa39")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}

	txHash := common.HexToHash("0x687cc806cf938403b156ad51e897409571285f520c2eb2d3fa260ad21b056de1")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending) // false
}

func transferETH() {
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

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x57Bd80cEf17E2091824261497DF820d45ce93dDD")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func createTransaction() string {

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

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4B240f7b195462BBb73093CEFd002C3031406861")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	b := new(bytes.Buffer)
	ts.EncodeIndex(0, b)
	rawTxBytes := b.Bytes()
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println(rawTxHex) // f86...772
	return rawTxHex
}

func sendTransaction() {

	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
	}

	rawTx := createTransaction()

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: ", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f

}

func loadSmartContract() {
	client, err := ethclient.Dial(gateway)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(smartContractHexAddress)
	instance, err := api.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}

func DeploySmartContract() {
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

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("smartContractHexAddress: ", address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println("tx: ", tx.Hash().Hex())                    // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}

func ReadWalletSmartContract(nameWallet string) (string, string) {
	client, address := ConnectContract(gateway, smartContractHexAddress)
	instance, err := api.NewApi(address, client)
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
	auth, instance := ConnectChangeContract(gateway, smartContractHexAddress, accountPrivateKey)
	tx, err := instance.SetWallet(auth, nameWallet, big.NewInt(balance))
	if err != nil {
		log.Fatal(err)
	}

	return tx.Hash().Hex()
}

func SendMoneySnartContract(nameWalletSender string, nameWalletRecipient string, sendMoney int64) string {
	auth, instance := ConnectChangeContract(gateway, smartContractHexAddress, accountPrivateKey)
	tx, err := instance.SendMoney(auth, nameWalletSender, nameWalletRecipient, big.NewInt(sendMoney))
	if err != nil {
		log.Fatal(err)
	}

	return tx.Hash().Hex()

}

func ConnectContract(gatewayInput string, smartContractHexAddressInput string) (*ethclient.Client, common.Address) {
	client, err := ethclient.Dial(gatewayInput)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(smartContractHexAddressInput)

	return client, address
}

func ConnectChangeContract(gateway string, smartContractHexAddress string, accountPrivateKey string) (*bind.TransactOpts, *api.Api) {
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

	address := common.HexToAddress(smartContractHexAddress)
	instance, err := api.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return auth, instance
}
