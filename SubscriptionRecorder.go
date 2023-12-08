// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// SubscriptionRecorderABI is the input ABI used to generate the binding from.
const SubscriptionRecorderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avaxContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"enumIEvent.SubscriptionType\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"hasValidSubscription\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subscriptions\",\"outputs\":[{\"internalType\":\"enumIEvent.SubscriptionType\",\"name\":\"subscriptionType\",\"type\":\"uint8\"},{\"internalType\":\"enumIEvent.SubscriptionTier\",\"name\":\"subscriptionTier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SubscriptionRecorder is an auto generated Go binding around an Ethereum contract.
type SubscriptionRecorder struct {
	SubscriptionRecorderCaller     // Read-only binding to the contract
	SubscriptionRecorderTransactor // Write-only binding to the contract
	SubscriptionRecorderFilterer   // Log filterer for contract events
}

// SubscriptionRecorderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionRecorderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionRecorderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionRecorderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionRecorderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionRecorderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionRecorderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionRecorderSession struct {
	Contract     *SubscriptionRecorder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SubscriptionRecorderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionRecorderCallerSession struct {
	Contract *SubscriptionRecorderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SubscriptionRecorderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionRecorderTransactorSession struct {
	Contract     *SubscriptionRecorderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SubscriptionRecorderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionRecorderRaw struct {
	Contract *SubscriptionRecorder // Generic contract binding to access the raw methods on
}

// SubscriptionRecorderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionRecorderCallerRaw struct {
	Contract *SubscriptionRecorderCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionRecorderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionRecorderTransactorRaw struct {
	Contract *SubscriptionRecorderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscriptionRecorder creates a new instance of SubscriptionRecorder, bound to a specific deployed contract.
func NewSubscriptionRecorder(address common.Address, backend bind.ContractBackend) (*SubscriptionRecorder, error) {
	contract, err := bindSubscriptionRecorder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscriptionRecorder{SubscriptionRecorderCaller: SubscriptionRecorderCaller{contract: contract}, SubscriptionRecorderTransactor: SubscriptionRecorderTransactor{contract: contract}, SubscriptionRecorderFilterer: SubscriptionRecorderFilterer{contract: contract}}, nil
}

// NewSubscriptionRecorderCaller creates a new read-only instance of SubscriptionRecorder, bound to a specific deployed contract.
func NewSubscriptionRecorderCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionRecorderCaller, error) {
	contract, err := bindSubscriptionRecorder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionRecorderCaller{contract: contract}, nil
}

// NewSubscriptionRecorderTransactor creates a new write-only instance of SubscriptionRecorder, bound to a specific deployed contract.
func NewSubscriptionRecorderTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionRecorderTransactor, error) {
	contract, err := bindSubscriptionRecorder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionRecorderTransactor{contract: contract}, nil
}

// NewSubscriptionRecorderFilterer creates a new log filterer instance of SubscriptionRecorder, bound to a specific deployed contract.
func NewSubscriptionRecorderFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionRecorderFilterer, error) {
	contract, err := bindSubscriptionRecorder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionRecorderFilterer{contract: contract}, nil
}

// bindSubscriptionRecorder binds a generic wrapper to an already deployed contract.
func bindSubscriptionRecorder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionRecorderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionRecorder *SubscriptionRecorderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionRecorder.Contract.SubscriptionRecorderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionRecorder *SubscriptionRecorderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.SubscriptionRecorderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionRecorder *SubscriptionRecorderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.SubscriptionRecorderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionRecorder *SubscriptionRecorderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubscriptionRecorder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionRecorder *SubscriptionRecorderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionRecorder *SubscriptionRecorderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.contract.Transact(opts, method, params...)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_SubscriptionRecorder *SubscriptionRecorderCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubscriptionRecorder.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_SubscriptionRecorder *SubscriptionRecorderSession) GetRouter() (common.Address, error) {
	return _SubscriptionRecorder.Contract.GetRouter(&_SubscriptionRecorder.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0xb0f479a1.
//
// Solidity: function getRouter() view returns(address)
func (_SubscriptionRecorder *SubscriptionRecorderCallerSession) GetRouter() (common.Address, error) {
	return _SubscriptionRecorder.Contract.GetRouter(&_SubscriptionRecorder.CallOpts)
}

// HasValidSubscription is a free data retrieval call binding the contract method 0x57d203b1.
//
// Solidity: function hasValidSubscription(address _address, uint8 _type) view returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderCaller) HasValidSubscription(opts *bind.CallOpts, _address common.Address, _type uint8) (bool, error) {
	var out []interface{}
	err := _SubscriptionRecorder.contract.Call(opts, &out, "hasValidSubscription", _address, _type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasValidSubscription is a free data retrieval call binding the contract method 0x57d203b1.
//
// Solidity: function hasValidSubscription(address _address, uint8 _type) view returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderSession) HasValidSubscription(_address common.Address, _type uint8) (bool, error) {
	return _SubscriptionRecorder.Contract.HasValidSubscription(&_SubscriptionRecorder.CallOpts, _address, _type)
}

// HasValidSubscription is a free data retrieval call binding the contract method 0x57d203b1.
//
// Solidity: function hasValidSubscription(address _address, uint8 _type) view returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderCallerSession) HasValidSubscription(_address common.Address, _type uint8) (bool, error) {
	return _SubscriptionRecorder.Contract.HasValidSubscription(&_SubscriptionRecorder.CallOpts, _address, _type)
}

// Subscriptions is a free data retrieval call binding the contract method 0x623da3d1.
//
// Solidity: function subscriptions(address , uint256 ) view returns(uint8 subscriptionType, uint8 subscriptionTier, uint256 endDate, address subscriber)
func (_SubscriptionRecorder *SubscriptionRecorderCaller) Subscriptions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	SubscriptionType uint8
	SubscriptionTier uint8
	EndDate          *big.Int
	Subscriber       common.Address
}, error) {
	var out []interface{}
	err := _SubscriptionRecorder.contract.Call(opts, &out, "subscriptions", arg0, arg1)

	outstruct := new(struct {
		SubscriptionType uint8
		SubscriptionTier uint8
		EndDate          *big.Int
		Subscriber       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SubscriptionType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SubscriptionTier = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.EndDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Subscriber = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Subscriptions is a free data retrieval call binding the contract method 0x623da3d1.
//
// Solidity: function subscriptions(address , uint256 ) view returns(uint8 subscriptionType, uint8 subscriptionTier, uint256 endDate, address subscriber)
func (_SubscriptionRecorder *SubscriptionRecorderSession) Subscriptions(arg0 common.Address, arg1 *big.Int) (struct {
	SubscriptionType uint8
	SubscriptionTier uint8
	EndDate          *big.Int
	Subscriber       common.Address
}, error) {
	return _SubscriptionRecorder.Contract.Subscriptions(&_SubscriptionRecorder.CallOpts, arg0, arg1)
}

// Subscriptions is a free data retrieval call binding the contract method 0x623da3d1.
//
// Solidity: function subscriptions(address , uint256 ) view returns(uint8 subscriptionType, uint8 subscriptionTier, uint256 endDate, address subscriber)
func (_SubscriptionRecorder *SubscriptionRecorderCallerSession) Subscriptions(arg0 common.Address, arg1 *big.Int) (struct {
	SubscriptionType uint8
	SubscriptionTier uint8
	EndDate          *big.Int
	Subscriber       common.Address
}, error) {
	return _SubscriptionRecorder.Contract.Subscriptions(&_SubscriptionRecorder.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SubscriptionRecorder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubscriptionRecorder.Contract.SupportsInterface(&_SubscriptionRecorder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_SubscriptionRecorder *SubscriptionRecorderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubscriptionRecorder.Contract.SupportsInterface(&_SubscriptionRecorder.CallOpts, interfaceId)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_SubscriptionRecorder *SubscriptionRecorderTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SubscriptionRecorder.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_SubscriptionRecorder *SubscriptionRecorderSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.CcipReceive(&_SubscriptionRecorder.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_SubscriptionRecorder *SubscriptionRecorderTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _SubscriptionRecorder.Contract.CcipReceive(&_SubscriptionRecorder.TransactOpts, message)
}
