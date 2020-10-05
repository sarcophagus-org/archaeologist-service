// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sarcophagus

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

// SarcophagusABI is the input ABI used to generate the binding from.
const SarcophagusABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sarcoToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"accuseArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"archaeologistCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"}],\"name\":\"archaeologists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"burySarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"cancelSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"cleanUpSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"archaeologistPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"}],\"name\":\"createSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"registerArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"rewrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcoToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"}],\"name\":\"unwrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"updateArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"updateSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawalBond\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Sarcophagus is an auto generated Go binding around an Ethereum contract.
type Sarcophagus struct {
	SarcophagusCaller     // Read-only binding to the contract
	SarcophagusTransactor // Write-only binding to the contract
	SarcophagusFilterer   // Log filterer for contract events
}

// SarcophagusCaller is an auto generated read-only Go binding around an Ethereum contract.
type SarcophagusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SarcophagusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SarcophagusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SarcophagusSession struct {
	Contract     *Sarcophagus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SarcophagusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SarcophagusCallerSession struct {
	Contract *SarcophagusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SarcophagusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SarcophagusTransactorSession struct {
	Contract     *SarcophagusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SarcophagusRaw is an auto generated low-level Go binding around an Ethereum contract.
type SarcophagusRaw struct {
	Contract *Sarcophagus // Generic contract binding to access the raw methods on
}

// SarcophagusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SarcophagusCallerRaw struct {
	Contract *SarcophagusCaller // Generic read-only contract binding to access the raw methods on
}

// SarcophagusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SarcophagusTransactorRaw struct {
	Contract *SarcophagusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSarcophagus creates a new instance of Sarcophagus, bound to a specific deployed contract.
func NewSarcophagus(address common.Address, backend bind.ContractBackend) (*Sarcophagus, error) {
	contract, err := bindSarcophagus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sarcophagus{SarcophagusCaller: SarcophagusCaller{contract: contract}, SarcophagusTransactor: SarcophagusTransactor{contract: contract}, SarcophagusFilterer: SarcophagusFilterer{contract: contract}}, nil
}

// NewSarcophagusCaller creates a new read-only instance of Sarcophagus, bound to a specific deployed contract.
func NewSarcophagusCaller(address common.Address, caller bind.ContractCaller) (*SarcophagusCaller, error) {
	contract, err := bindSarcophagus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SarcophagusCaller{contract: contract}, nil
}

// NewSarcophagusTransactor creates a new write-only instance of Sarcophagus, bound to a specific deployed contract.
func NewSarcophagusTransactor(address common.Address, transactor bind.ContractTransactor) (*SarcophagusTransactor, error) {
	contract, err := bindSarcophagus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SarcophagusTransactor{contract: contract}, nil
}

// NewSarcophagusFilterer creates a new log filterer instance of Sarcophagus, bound to a specific deployed contract.
func NewSarcophagusFilterer(address common.Address, filterer bind.ContractFilterer) (*SarcophagusFilterer, error) {
	contract, err := bindSarcophagus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SarcophagusFilterer{contract: contract}, nil
}

// bindSarcophagus binds a generic wrapper to an already deployed contract.
func bindSarcophagus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SarcophagusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sarcophagus *SarcophagusRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sarcophagus.Contract.SarcophagusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sarcophagus *SarcophagusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sarcophagus.Contract.SarcophagusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sarcophagus *SarcophagusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sarcophagus.Contract.SarcophagusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sarcophagus *SarcophagusCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sarcophagus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sarcophagus *SarcophagusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sarcophagus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sarcophagus *SarcophagusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sarcophagus.Contract.contract.Transact(opts, method, params...)
}

// ArchaeologistAddresses is a free data retrieval call binding the contract method 0x66086784.
//
// Solidity: function archaeologistAddresses(uint256 index) view returns(address)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistAddresses(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sarcophagus.contract.Call(opts, out, "archaeologistAddresses", index)
	return *ret0, err
}

// ArchaeologistAddresses is a free data retrieval call binding the contract method 0x66086784.
//
// Solidity: function archaeologistAddresses(uint256 index) view returns(address)
func (_Sarcophagus *SarcophagusSession) ArchaeologistAddresses(index *big.Int) (common.Address, error) {
	return _Sarcophagus.Contract.ArchaeologistAddresses(&_Sarcophagus.CallOpts, index)
}

// ArchaeologistAddresses is a free data retrieval call binding the contract method 0x66086784.
//
// Solidity: function archaeologistAddresses(uint256 index) view returns(address)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistAddresses(index *big.Int) (common.Address, error) {
	return _Sarcophagus.Contract.ArchaeologistAddresses(&_Sarcophagus.CallOpts, index)
}

// ArchaeologistCount is a free data retrieval call binding the contract method 0x1a7ac550.
//
// Solidity: function archaeologistCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sarcophagus.contract.Call(opts, out, "archaeologistCount")
	return *ret0, err
}

// ArchaeologistCount is a free data retrieval call binding the contract method 0x1a7ac550.
//
// Solidity: function archaeologistCount() view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistCount() (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCount(&_Sarcophagus.CallOpts)
}

// ArchaeologistCount is a free data retrieval call binding the contract method 0x1a7ac550.
//
// Solidity: function archaeologistCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistCount() (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCount(&_Sarcophagus.CallOpts)
}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address addy) view returns(bool exists, bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusCaller) Archaeologists(opts *bind.CallOpts, addy common.Address) (struct {
	Exists                  bool
	PublicKey               []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	FreeBond                *big.Int
	CursedBond              *big.Int
}, error) {
	ret := new(struct {
		Exists                  bool
		PublicKey               []byte
		Endpoint                string
		PaymentAddress          common.Address
		FeePerByte              *big.Int
		MinimumBounty           *big.Int
		MinimumDiggingFee       *big.Int
		MaximumResurrectionTime *big.Int
		FreeBond                *big.Int
		CursedBond              *big.Int
	})
	out := ret
	err := _Sarcophagus.contract.Call(opts, out, "archaeologists", addy)
	return *ret, err
}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address addy) view returns(bool exists, bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusSession) Archaeologists(addy common.Address) (struct {
	Exists                  bool
	PublicKey               []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	FreeBond                *big.Int
	CursedBond              *big.Int
}, error) {
	return _Sarcophagus.Contract.Archaeologists(&_Sarcophagus.CallOpts, addy)
}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address addy) view returns(bool exists, bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusCallerSession) Archaeologists(addy common.Address) (struct {
	Exists                  bool
	PublicKey               []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	FreeBond                *big.Int
	CursedBond              *big.Int
}, error) {
	return _Sarcophagus.Contract.Archaeologists(&_Sarcophagus.CallOpts, addy)
}

// SarcoToken is a free data retrieval call binding the contract method 0x607e8f09.
//
// Solidity: function sarcoToken() view returns(address)
func (_Sarcophagus *SarcophagusCaller) SarcoToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sarcophagus.contract.Call(opts, out, "sarcoToken")
	return *ret0, err
}

// SarcoToken is a free data retrieval call binding the contract method 0x607e8f09.
//
// Solidity: function sarcoToken() view returns(address)
func (_Sarcophagus *SarcophagusSession) SarcoToken() (common.Address, error) {
	return _Sarcophagus.Contract.SarcoToken(&_Sarcophagus.CallOpts)
}

// SarcoToken is a free data retrieval call binding the contract method 0x607e8f09.
//
// Solidity: function sarcoToken() view returns(address)
func (_Sarcophagus *SarcophagusCallerSession) SarcoToken() (common.Address, error) {
	return _Sarcophagus.Contract.SarcoToken(&_Sarcophagus.CallOpts)
}

// SarcophagusCount is a free data retrieval call binding the contract method 0x555912b7.
//
// Solidity: function sarcophagusCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) SarcophagusCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sarcophagus.contract.Call(opts, out, "sarcophagusCount")
	return *ret0, err
}

// SarcophagusCount is a free data retrieval call binding the contract method 0x555912b7.
//
// Solidity: function sarcophagusCount() view returns(uint256)
func (_Sarcophagus *SarcophagusSession) SarcophagusCount() (*big.Int, error) {
	return _Sarcophagus.Contract.SarcophagusCount(&_Sarcophagus.CallOpts)
}

// SarcophagusCount is a free data retrieval call binding the contract method 0x555912b7.
//
// Solidity: function sarcophagusCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) SarcophagusCount() (*big.Int, error) {
	return _Sarcophagus.Contract.SarcophagusCount(&_Sarcophagus.CallOpts)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 assetDoubleHash, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) AccuseArchaeologist(opts *bind.TransactOpts, assetDoubleHash [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "accuseArchaeologist", assetDoubleHash, singleHash, paymentAddress)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 assetDoubleHash, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusSession) AccuseArchaeologist(assetDoubleHash [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.AccuseArchaeologist(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash, paymentAddress)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 assetDoubleHash, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) AccuseArchaeologist(assetDoubleHash [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.AccuseArchaeologist(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash, paymentAddress)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) BurySarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "burySarcophagus", assetDoubleHash)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusSession) BurySarcophagus(assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.BurySarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) BurySarcophagus(assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.BurySarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CancelSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "cancelSarcophagus", assetDoubleHash)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusSession) CancelSarcophagus(assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CancelSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 assetDoubleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CancelSarcophagus(assetDoubleHash [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CancelSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 assetDoubleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CleanUpSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "cleanUpSarcophagus", assetDoubleHash, paymentAddress)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 assetDoubleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusSession) CleanUpSarcophagus(assetDoubleHash [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CleanUpSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, paymentAddress)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 assetDoubleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CleanUpSarcophagus(assetDoubleHash [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CleanUpSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, paymentAddress)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x38e2c558.
//
// Solidity: function createSarcophagus(string name, bytes archaeologistPublicKey, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CreateSarcophagus(opts *bind.TransactOpts, name string, archaeologistPublicKey []byte, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "createSarcophagus", name, archaeologistPublicKey, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x38e2c558.
//
// Solidity: function createSarcophagus(string name, bytes archaeologistPublicKey, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusSession) CreateSarcophagus(name string, archaeologistPublicKey []byte, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologistPublicKey, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x38e2c558.
//
// Solidity: function createSarcophagus(string name, bytes archaeologistPublicKey, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CreateSarcophagus(name string, archaeologistPublicKey []byte, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologistPublicKey, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) RegisterArchaeologist(opts *bind.TransactOpts, publicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "registerArchaeologist", publicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusSession) RegisterArchaeologist(publicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, publicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes publicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) RegisterArchaeologist(publicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, publicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 assetDoubleHash, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) RewrapSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "rewrapSarcophagus", assetDoubleHash, resurrectionTime, diggingFee, bounty)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 assetDoubleHash, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusSession) RewrapSarcophagus(assetDoubleHash [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RewrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, resurrectionTime, diggingFee, bounty)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 assetDoubleHash, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) RewrapSarcophagus(assetDoubleHash [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RewrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, resurrectionTime, diggingFee, bounty)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x8d58642a.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UnwrapSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte, singleHash []byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "unwrapSarcophagus", assetDoubleHash, singleHash)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x8d58642a.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash) returns(bool)
func (_Sarcophagus *SarcophagusSession) UnwrapSarcophagus(assetDoubleHash [32]byte, singleHash []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x8d58642a.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UnwrapSarcophagus(assetDoubleHash [32]byte, singleHash []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xeb15ef17.
//
// Solidity: function updateArchaeologist(string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateArchaeologist(opts *bind.TransactOpts, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateArchaeologist", endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xeb15ef17.
//
// Solidity: function updateArchaeologist(string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateArchaeologist(endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xeb15ef17.
//
// Solidity: function updateArchaeologist(string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateArchaeologist(endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xa7331797.
//
// Solidity: function updateSarcophagus(bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateSarcophagus", assetDoubleHash, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xa7331797.
//
// Solidity: function updateSarcophagus(bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateSarcophagus(assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xa7331797.
//
// Solidity: function updateSarcophagus(bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateSarcophagus(assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, assetId, v, r, s)
}

// WithdrawalBond is a paid mutator transaction binding the contract method 0xbbe026ca.
//
// Solidity: function withdrawalBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) WithdrawalBond(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "withdrawalBond", amount)
}

// WithdrawalBond is a paid mutator transaction binding the contract method 0xbbe026ca.
//
// Solidity: function withdrawalBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusSession) WithdrawalBond(amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.WithdrawalBond(&_Sarcophagus.TransactOpts, amount)
}

// WithdrawalBond is a paid mutator transaction binding the contract method 0xbbe026ca.
//
// Solidity: function withdrawalBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) WithdrawalBond(amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.WithdrawalBond(&_Sarcophagus.TransactOpts, amount)
}
