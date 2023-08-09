// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// GatewayVerifierMetaData contains all meta data concerning the GatewayVerifier contract.
var GatewayVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"verifyToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"verifyToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GatewayVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayVerifierMetaData.ABI instead.
var GatewayVerifierABI = GatewayVerifierMetaData.ABI

// GatewayVerifier is an auto generated Go binding around an Ethereum contract.
type GatewayVerifier struct {
	GatewayVerifierCaller     // Read-only binding to the contract
	GatewayVerifierTransactor // Write-only binding to the contract
	GatewayVerifierFilterer   // Log filterer for contract events
}

// GatewayVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayVerifierSession struct {
	Contract     *GatewayVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayVerifierCallerSession struct {
	Contract *GatewayVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GatewayVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayVerifierTransactorSession struct {
	Contract     *GatewayVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GatewayVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayVerifierRaw struct {
	Contract *GatewayVerifier // Generic contract binding to access the raw methods on
}

// GatewayVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayVerifierCallerRaw struct {
	Contract *GatewayVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayVerifierTransactorRaw struct {
	Contract *GatewayVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayVerifier creates a new instance of GatewayVerifier, bound to a specific deployed contract.
func NewGatewayVerifier(address common.Address, backend bind.ContractBackend) (*GatewayVerifier, error) {
	contract, err := bindGatewayVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayVerifier{GatewayVerifierCaller: GatewayVerifierCaller{contract: contract}, GatewayVerifierTransactor: GatewayVerifierTransactor{contract: contract}, GatewayVerifierFilterer: GatewayVerifierFilterer{contract: contract}}, nil
}

// NewGatewayVerifierCaller creates a new read-only instance of GatewayVerifier, bound to a specific deployed contract.
func NewGatewayVerifierCaller(address common.Address, caller bind.ContractCaller) (*GatewayVerifierCaller, error) {
	contract, err := bindGatewayVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayVerifierCaller{contract: contract}, nil
}

// NewGatewayVerifierTransactor creates a new write-only instance of GatewayVerifier, bound to a specific deployed contract.
func NewGatewayVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayVerifierTransactor, error) {
	contract, err := bindGatewayVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayVerifierTransactor{contract: contract}, nil
}

// NewGatewayVerifierFilterer creates a new log filterer instance of GatewayVerifier, bound to a specific deployed contract.
func NewGatewayVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayVerifierFilterer, error) {
	contract, err := bindGatewayVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayVerifierFilterer{contract: contract}, nil
}

// bindGatewayVerifier binds a generic wrapper to an already deployed contract.
func bindGatewayVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayVerifier *GatewayVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayVerifier.Contract.GatewayVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayVerifier *GatewayVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayVerifier.Contract.GatewayVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayVerifier *GatewayVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayVerifier.Contract.GatewayVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayVerifier *GatewayVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayVerifier *GatewayVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayVerifier *GatewayVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyToken is a free data retrieval call binding the contract method 0xff17e232.
//
// Solidity: function verifyToken(address owner, uint256 network) view returns(bool)
func (_GatewayVerifier *GatewayVerifierCaller) VerifyToken(opts *bind.CallOpts, owner common.Address, network *big.Int) (bool, error) {
	var out []interface{}
	err := _GatewayVerifier.contract.Call(opts, &out, "verifyToken", owner, network)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyToken is a free data retrieval call binding the contract method 0xff17e232.
//
// Solidity: function verifyToken(address owner, uint256 network) view returns(bool)
func (_GatewayVerifier *GatewayVerifierSession) VerifyToken(owner common.Address, network *big.Int) (bool, error) {
	return _GatewayVerifier.Contract.VerifyToken(&_GatewayVerifier.CallOpts, owner, network)
}

// VerifyToken is a free data retrieval call binding the contract method 0xff17e232.
//
// Solidity: function verifyToken(address owner, uint256 network) view returns(bool)
func (_GatewayVerifier *GatewayVerifierCallerSession) VerifyToken(owner common.Address, network *big.Int) (bool, error) {
	return _GatewayVerifier.Contract.VerifyToken(&_GatewayVerifier.CallOpts, owner, network)
}

// VerifyToken0 is a free data retrieval call binding the contract method 0xff85a975.
//
// Solidity: function verifyToken(uint256 tokenId) view returns(bool)
func (_GatewayVerifier *GatewayVerifierCaller) VerifyToken0(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _GatewayVerifier.contract.Call(opts, &out, "verifyToken0", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyToken0 is a free data retrieval call binding the contract method 0xff85a975.
//
// Solidity: function verifyToken(uint256 tokenId) view returns(bool)
func (_GatewayVerifier *GatewayVerifierSession) VerifyToken0(tokenId *big.Int) (bool, error) {
	return _GatewayVerifier.Contract.VerifyToken0(&_GatewayVerifier.CallOpts, tokenId)
}

// VerifyToken0 is a free data retrieval call binding the contract method 0xff85a975.
//
// Solidity: function verifyToken(uint256 tokenId) view returns(bool)
func (_GatewayVerifier *GatewayVerifierCallerSession) VerifyToken0(tokenId *big.Int) (bool, error) {
	return _GatewayVerifier.Contract.VerifyToken0(&_GatewayVerifier.CallOpts, tokenId)
}
