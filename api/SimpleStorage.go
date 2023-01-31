// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimpleStorageWallet is an auto generated low-level Go binding around an user-defined struct.
type SimpleStorageWallet struct {
	WalletName string
	Balance    *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWallet\",\"type\":\"string\"}],\"name\":\"getWallet\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"walletName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structSimpleStorage.Wallet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWalletSender\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nameWalletRecipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"sendMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nameWallet\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"setBalance\",\"type\":\"uint256\"}],\"name\":\"setWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"walletName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061081f806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806358ef40961461005157806372f179f814610066578063a4e2df6614610090578063b418cf32146100b0575b600080fd5b61006461005f3660046104b0565b6100c3565b005b61007961007436600461051d565b610229565b6040516100879291906105aa565b60405180910390f35b6100a361009e36600461051d565b6102d8565b60405161008791906105cc565b6100646100be3660046105fe565b6103b9565b6000836040516100d39190610643565b9081526020016040518091039020600101548111158015610124575063ffffffff816000846040516101059190610643565b9081526020016040518091039020600101546101219190610675565b11155b15610224578260008460405161013a9190610643565b908152604051908190036020019020906101549082610716565b50806000846040516101669190610643565b90815260200160405180910390206001015461018291906107d6565b6000846040516101929190610643565b908152602001604051809103902060010181905550816000836040516101b89190610643565b908152604051908190036020019020906101d29082610716565b50806000836040516101e49190610643565b9081526020016040518091039020600101546102009190610675565b6000836040516102109190610643565b908152604051908190036020019020600101555b505050565b805160208183018101805160008252928201919093012091528054819061024f9061068e565b80601f016020809104026020016040519081016040528092919081815260200182805461027b9061068e565b80156102c85780601f1061029d576101008083540402835291602001916102c8565b820191906000526020600020905b8154815290600101906020018083116102ab57829003601f168201915b5050505050908060010154905082565b6040805180820190915260608152600060208201526000826040516102fd9190610643565b90815260200160405180910390206040518060400160405290816000820180546103269061068e565b80601f01602080910402602001604051908101604052809291908181526020018280546103529061068e565b801561039f5780601f106103745761010080835404028352916020019161039f565b820191906000526020600020905b81548152906001019060200180831161038257829003601f168201915b505050505081526020016001820154815250509050919050565b816000836040516103ca9190610643565b908152604051908190036020019020906103e49082610716565b50806000836040516103f69190610643565b908152604051908190036020019020600101555050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261043457600080fd5b813567ffffffffffffffff8082111561044f5761044f61040d565b604051601f8301601f19908116603f011681019082821181831017156104775761047761040d565b8160405283815286602085880101111561049057600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156104c557600080fd5b833567ffffffffffffffff808211156104dd57600080fd5b6104e987838801610423565b945060208601359150808211156104ff57600080fd5b5061050c86828701610423565b925050604084013590509250925092565b60006020828403121561052f57600080fd5b813567ffffffffffffffff81111561054657600080fd5b61055284828501610423565b949350505050565b60005b8381101561057557818101518382015260200161055d565b50506000910152565b6000815180845261059681602086016020860161055a565b601f01601f19169290920160200192915050565b6040815260006105bd604083018561057e565b90508260208301529392505050565b6020815260008251604060208401526105e8606084018261057e565b9050602084015160408401528091505092915050565b6000806040838503121561061157600080fd5b823567ffffffffffffffff81111561062857600080fd5b61063485828601610423565b95602094909401359450505050565b6000825161065581846020870161055a565b9190910192915050565b634e487b7160e01b600052601160045260246000fd5b808201808211156106885761068861065f565b92915050565b600181811c908216806106a257607f821691505b6020821081036106c257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561022457600081815260208120601f850160051c810160208610156106ef5750805b601f850160051c820191505b8181101561070e578281556001016106fb565b505050505050565b815167ffffffffffffffff8111156107305761073061040d565b6107448161073e845461068e565b846106c8565b602080601f83116001811461077957600084156107615750858301515b600019600386901b1c1916600185901b17855561070e565b600085815260208120601f198616915b828110156107a857888601518255948401946001909101908401610789565b50858210156107c65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b818103818111156106885761068861065f56fea264697066735822122052ecce26ad6dd7251d8084ff157d50f9f86588244a0b4eb71ca6e2e6a464979064736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiCaller) GetWallet(opts *bind.CallOpts, nameWallet string) (SimpleStorageWallet, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getWallet", nameWallet)

	if err != nil {
		return *new(SimpleStorageWallet), err
	}

	out0 := *abi.ConvertType(out[0], new(SimpleStorageWallet)).(*SimpleStorageWallet)

	return out0, err

}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiSession) GetWallet(nameWallet string) (SimpleStorageWallet, error) {
	return _Api.Contract.GetWallet(&_Api.CallOpts, nameWallet)
}

// GetWallet is a free data retrieval call binding the contract method 0xa4e2df66.
//
// Solidity: function getWallet(string nameWallet) view returns((string,uint256))
func (_Api *ApiCallerSession) GetWallet(nameWallet string) (SimpleStorageWallet, error) {
	return _Api.Contract.GetWallet(&_Api.CallOpts, nameWallet)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiCaller) Wallets(opts *bind.CallOpts, arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "wallets", arg0)

	outstruct := new(struct {
		WalletName string
		Balance    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiSession) Wallets(arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	return _Api.Contract.Wallets(&_Api.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(string walletName, uint256 balance)
func (_Api *ApiCallerSession) Wallets(arg0 string) (struct {
	WalletName string
	Balance    *big.Int
}, error) {
	return _Api.Contract.Wallets(&_Api.CallOpts, arg0)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiTransactor) SendMoney(opts *bind.TransactOpts, nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendMoney", nameWalletSender, nameWalletRecipient, money)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiSession) SendMoney(nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoney(&_Api.TransactOpts, nameWalletSender, nameWalletRecipient, money)
}

// SendMoney is a paid mutator transaction binding the contract method 0x58ef4096.
//
// Solidity: function sendMoney(string nameWalletSender, string nameWalletRecipient, uint256 money) returns()
func (_Api *ApiTransactorSession) SendMoney(nameWalletSender string, nameWalletRecipient string, money *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SendMoney(&_Api.TransactOpts, nameWalletSender, nameWalletRecipient, money)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiTransactor) SetWallet(opts *bind.TransactOpts, nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setWallet", nameWallet, setBalance)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiSession) SetWallet(nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetWallet(&_Api.TransactOpts, nameWallet, setBalance)
}

// SetWallet is a paid mutator transaction binding the contract method 0xb418cf32.
//
// Solidity: function setWallet(string nameWallet, uint256 setBalance) returns()
func (_Api *ApiTransactorSession) SetWallet(nameWallet string, setBalance *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SetWallet(&_Api.TransactOpts, nameWallet, setBalance)
}
