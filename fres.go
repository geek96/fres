// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AirdropContractABI is the input ABI used to generate the binding from.
const AirdropContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dests\",\"type\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"send\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AirdropContractBin is the compiled bytecode used for deploying new contracts.
const AirdropContractBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a031991821617909155600180549091167329218e791cb7def0cae8fed26d06697d808a62341790556102c08061005d6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063a645ff5f1461008e575b600080fd5b34801561005c57600080fd5b5061006561012e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506040805160206004803580820135838102808601850190965280855261011c95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a99890198929750908201955093508392508501908490808284375094975061014a9650505050505050565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000805481903373ffffffffffffffffffffffffffffffffffffffff90811691161461017557600080fd5b5060005b835181101561028d57600154845173ffffffffffffffffffffffffffffffffffffffff9091169063a9059cbb908690849081106101b257fe5b9060200190602002015185848151811015156101ca57fe5b906020019060200201516040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561025957600080fd5b505af115801561026d573d6000803e3d6000fd5b505050506040513d602081101561028357600080fd5b5050600101610179565b93925050505600a165627a7a72305820048e9f9aa0cbb788399ec179e19913cc6bc657d287f3c547135bd36daf1c7bf20029`

// DeployAirdropContract deploys a new Ethereum contract, binding an instance of AirdropContract to it.
func DeployAirdropContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AirdropContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AirdropContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AirdropContract{AirdropContractCaller: AirdropContractCaller{contract: contract}, AirdropContractTransactor: AirdropContractTransactor{contract: contract}, AirdropContractFilterer: AirdropContractFilterer{contract: contract}}, nil
}

// AirdropContract is an auto generated Go binding around an Ethereum contract.
type AirdropContract struct {
	AirdropContractCaller     // Read-only binding to the contract
	AirdropContractTransactor // Write-only binding to the contract
	AirdropContractFilterer   // Log filterer for contract events
}

// AirdropContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropContractSession struct {
	Contract     *AirdropContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropContractCallerSession struct {
	Contract *AirdropContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AirdropContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropContractTransactorSession struct {
	Contract     *AirdropContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AirdropContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropContractRaw struct {
	Contract *AirdropContract // Generic contract binding to access the raw methods on
}

// AirdropContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropContractCallerRaw struct {
	Contract *AirdropContractCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropContractTransactorRaw struct {
	Contract *AirdropContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdropContract creates a new instance of AirdropContract, bound to a specific deployed contract.
func NewAirdropContract(address common.Address, backend bind.ContractBackend) (*AirdropContract, error) {
	contract, err := bindAirdropContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AirdropContract{AirdropContractCaller: AirdropContractCaller{contract: contract}, AirdropContractTransactor: AirdropContractTransactor{contract: contract}, AirdropContractFilterer: AirdropContractFilterer{contract: contract}}, nil
}

// NewAirdropContractCaller creates a new read-only instance of AirdropContract, bound to a specific deployed contract.
func NewAirdropContractCaller(address common.Address, caller bind.ContractCaller) (*AirdropContractCaller, error) {
	contract, err := bindAirdropContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropContractCaller{contract: contract}, nil
}

// NewAirdropContractTransactor creates a new write-only instance of AirdropContract, bound to a specific deployed contract.
func NewAirdropContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropContractTransactor, error) {
	contract, err := bindAirdropContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropContractTransactor{contract: contract}, nil
}

// NewAirdropContractFilterer creates a new log filterer instance of AirdropContract, bound to a specific deployed contract.
func NewAirdropContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropContractFilterer, error) {
	contract, err := bindAirdropContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropContractFilterer{contract: contract}, nil
}

// bindAirdropContract binds a generic wrapper to an already deployed contract.
func bindAirdropContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirdropContract *AirdropContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AirdropContract.Contract.AirdropContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirdropContract *AirdropContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirdropContract.Contract.AirdropContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirdropContract *AirdropContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirdropContract.Contract.AirdropContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AirdropContract *AirdropContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AirdropContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AirdropContract *AirdropContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AirdropContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AirdropContract *AirdropContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AirdropContract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AirdropContract *AirdropContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AirdropContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AirdropContract *AirdropContractSession) Owner() (common.Address, error) {
	return _AirdropContract.Contract.Owner(&_AirdropContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AirdropContract *AirdropContractCallerSession) Owner() (common.Address, error) {
	return _AirdropContract.Contract.Owner(&_AirdropContract.CallOpts)
}

// Send is a paid mutator transaction binding the contract method 0xa645ff5f.
//
// Solidity: function send(dests address[], values uint256[]) returns(uint256)
func (_AirdropContract *AirdropContractTransactor) Send(opts *bind.TransactOpts, dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirdropContract.contract.Transact(opts, "send", dests, values)
}

// Send is a paid mutator transaction binding the contract method 0xa645ff5f.
//
// Solidity: function send(dests address[], values uint256[]) returns(uint256)
func (_AirdropContract *AirdropContractSession) Send(dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirdropContract.Contract.Send(&_AirdropContract.TransactOpts, dests, values)
}

// Send is a paid mutator transaction binding the contract method 0xa645ff5f.
//
// Solidity: function send(dests address[], values uint256[]) returns(uint256)
func (_AirdropContract *AirdropContractTransactorSession) Send(dests []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _AirdropContract.Contract.Send(&_AirdropContract.TransactOpts, dests, values)
}

// FrescoTokenABI is the input ABI used to generate the binding from.
const FrescoTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FrescoTokenBin is the compiled bytecode used for deploying new contracts.
const FrescoTokenBin = `0x`

// DeployFrescoToken deploys a new Ethereum contract, binding an instance of FrescoToken to it.
func DeployFrescoToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FrescoToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FrescoTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FrescoTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FrescoToken{FrescoTokenCaller: FrescoTokenCaller{contract: contract}, FrescoTokenTransactor: FrescoTokenTransactor{contract: contract}, FrescoTokenFilterer: FrescoTokenFilterer{contract: contract}}, nil
}

// FrescoToken is an auto generated Go binding around an Ethereum contract.
type FrescoToken struct {
	FrescoTokenCaller     // Read-only binding to the contract
	FrescoTokenTransactor // Write-only binding to the contract
	FrescoTokenFilterer   // Log filterer for contract events
}

// FrescoTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FrescoTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrescoTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FrescoTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrescoTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FrescoTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FrescoTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FrescoTokenSession struct {
	Contract     *FrescoToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FrescoTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FrescoTokenCallerSession struct {
	Contract *FrescoTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FrescoTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FrescoTokenTransactorSession struct {
	Contract     *FrescoTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FrescoTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FrescoTokenRaw struct {
	Contract *FrescoToken // Generic contract binding to access the raw methods on
}

// FrescoTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FrescoTokenCallerRaw struct {
	Contract *FrescoTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FrescoTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FrescoTokenTransactorRaw struct {
	Contract *FrescoTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFrescoToken creates a new instance of FrescoToken, bound to a specific deployed contract.
func NewFrescoToken(address common.Address, backend bind.ContractBackend) (*FrescoToken, error) {
	contract, err := bindFrescoToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FrescoToken{FrescoTokenCaller: FrescoTokenCaller{contract: contract}, FrescoTokenTransactor: FrescoTokenTransactor{contract: contract}, FrescoTokenFilterer: FrescoTokenFilterer{contract: contract}}, nil
}

// NewFrescoTokenCaller creates a new read-only instance of FrescoToken, bound to a specific deployed contract.
func NewFrescoTokenCaller(address common.Address, caller bind.ContractCaller) (*FrescoTokenCaller, error) {
	contract, err := bindFrescoToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FrescoTokenCaller{contract: contract}, nil
}

// NewFrescoTokenTransactor creates a new write-only instance of FrescoToken, bound to a specific deployed contract.
func NewFrescoTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FrescoTokenTransactor, error) {
	contract, err := bindFrescoToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FrescoTokenTransactor{contract: contract}, nil
}

// NewFrescoTokenFilterer creates a new log filterer instance of FrescoToken, bound to a specific deployed contract.
func NewFrescoTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FrescoTokenFilterer, error) {
	contract, err := bindFrescoToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FrescoTokenFilterer{contract: contract}, nil
}

// bindFrescoToken binds a generic wrapper to an already deployed contract.
func bindFrescoToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FrescoTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FrescoToken *FrescoTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FrescoToken.Contract.FrescoTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FrescoToken *FrescoTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FrescoToken.Contract.FrescoTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FrescoToken *FrescoTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FrescoToken.Contract.FrescoTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FrescoToken *FrescoTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FrescoToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FrescoToken *FrescoTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FrescoToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FrescoToken *FrescoTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FrescoToken.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_FrescoToken *FrescoTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FrescoToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_FrescoToken *FrescoTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FrescoToken.Contract.Transfer(&_FrescoToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_FrescoToken *FrescoTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FrescoToken.Contract.Transfer(&_FrescoToken.TransactOpts, to, value)
}
