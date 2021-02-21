// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TypesArchaeologist is an auto generated low-level Go binding around an user-defined struct.
type TypesArchaeologist struct {
	Exists                  bool
	CurrentPublicKey        []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	FreeBond                *big.Int
	CursedBond              *big.Int
}

// TypesSarcophagus is an auto generated low-level Go binding around an user-defined struct.
type TypesSarcophagus struct {
	State                  uint8
	Archaeologist          common.Address
	ArchaeologistPublicKey []byte
	Embalmer               common.Address
	Name                   string
	ResurrectionTime       *big.Int
	ResurrectionWindow     *big.Int
	AssetId                string
	RecipientPublicKey     []byte
	StorageFee             *big.Int
	DiggingFee             *big.Int
	Bounty                 *big.Int
	CurrentCursedBond      *big.Int
	PrivateKey             [32]byte
}

// SarcophagusABI is the input ABI used to generate the binding from.
const SarcophagusABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"accuseArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"}],\"name\":\"archaeologistAccusalsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistAccusalsIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"}],\"name\":\"archaeologistCancelsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistCancelsIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"}],\"name\":\"archaeologistCleanupsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistCleanupsIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"archaeologistCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"}],\"name\":\"archaeologistSarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistSarcophagusIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"}],\"name\":\"archaeologistSuccessesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistSuccessesIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"archaeologists\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.Archaeologist\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"burySarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"cancelSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"cleanUpSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"}],\"name\":\"createSarcophagus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"}],\"name\":\"embalmerSarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"embalmerSarcophagusIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sarcoToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recipientSarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"recipientSarcophagusIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"registerArchaeologist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"rewrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcoToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"sarcophagus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumTypes.SarcophagusStates\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"archaeologistPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentCursedBond\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"privateKey\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.Sarcophagus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"sarcophagusIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"privateKey\",\"type\":\"bytes32\"}],\"name\":\"unwrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"newPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"updateArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"updateSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBond\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
func (_Sarcophagus *SarcophagusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Sarcophagus *SarcophagusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// ArchaeologistAccusalsCount is a free data retrieval call binding the contract method 0x0178a06a.
//
// Solidity: function archaeologistAccusalsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistAccusalsCount(opts *bind.CallOpts, archaeologist common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistAccusalsCount", archaeologist)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArchaeologistAccusalsCount is a free data retrieval call binding the contract method 0x0178a06a.
//
// Solidity: function archaeologistAccusalsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistAccusalsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistAccusalsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistAccusalsCount is a free data retrieval call binding the contract method 0x0178a06a.
//
// Solidity: function archaeologistAccusalsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistAccusalsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistAccusalsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistAccusalsIdentifier is a free data retrieval call binding the contract method 0x6c99bedc.
//
// Solidity: function archaeologistAccusalsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistAccusalsIdentifier(opts *bind.CallOpts, archaeologist common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistAccusalsIdentifier", archaeologist, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ArchaeologistAccusalsIdentifier is a free data retrieval call binding the contract method 0x6c99bedc.
//
// Solidity: function archaeologistAccusalsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) ArchaeologistAccusalsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistAccusalsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistAccusalsIdentifier is a free data retrieval call binding the contract method 0x6c99bedc.
//
// Solidity: function archaeologistAccusalsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistAccusalsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistAccusalsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistAddresses is a free data retrieval call binding the contract method 0x66086784.
//
// Solidity: function archaeologistAddresses(uint256 index) view returns(address)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistAddresses(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistAddresses", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// ArchaeologistCancelsCount is a free data retrieval call binding the contract method 0x5bb38160.
//
// Solidity: function archaeologistCancelsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCancelsCount(opts *bind.CallOpts, archaeologist common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistCancelsCount", archaeologist)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArchaeologistCancelsCount is a free data retrieval call binding the contract method 0x5bb38160.
//
// Solidity: function archaeologistCancelsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistCancelsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCancelsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistCancelsCount is a free data retrieval call binding the contract method 0x5bb38160.
//
// Solidity: function archaeologistCancelsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistCancelsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCancelsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistCancelsIdentifier is a free data retrieval call binding the contract method 0x99959230.
//
// Solidity: function archaeologistCancelsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCancelsIdentifier(opts *bind.CallOpts, archaeologist common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistCancelsIdentifier", archaeologist, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ArchaeologistCancelsIdentifier is a free data retrieval call binding the contract method 0x99959230.
//
// Solidity: function archaeologistCancelsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) ArchaeologistCancelsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistCancelsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistCancelsIdentifier is a free data retrieval call binding the contract method 0x99959230.
//
// Solidity: function archaeologistCancelsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistCancelsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistCancelsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistCleanupsCount is a free data retrieval call binding the contract method 0x64bd0ce9.
//
// Solidity: function archaeologistCleanupsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCleanupsCount(opts *bind.CallOpts, archaeologist common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistCleanupsCount", archaeologist)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArchaeologistCleanupsCount is a free data retrieval call binding the contract method 0x64bd0ce9.
//
// Solidity: function archaeologistCleanupsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistCleanupsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCleanupsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistCleanupsCount is a free data retrieval call binding the contract method 0x64bd0ce9.
//
// Solidity: function archaeologistCleanupsCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistCleanupsCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistCleanupsCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistCleanupsIdentifier is a free data retrieval call binding the contract method 0xfb6fffb5.
//
// Solidity: function archaeologistCleanupsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCleanupsIdentifier(opts *bind.CallOpts, archaeologist common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistCleanupsIdentifier", archaeologist, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ArchaeologistCleanupsIdentifier is a free data retrieval call binding the contract method 0xfb6fffb5.
//
// Solidity: function archaeologistCleanupsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) ArchaeologistCleanupsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistCleanupsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistCleanupsIdentifier is a free data retrieval call binding the contract method 0xfb6fffb5.
//
// Solidity: function archaeologistCleanupsIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistCleanupsIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistCleanupsIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistCount is a free data retrieval call binding the contract method 0x1a7ac550.
//
// Solidity: function archaeologistCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// ArchaeologistSarcophagusCount is a free data retrieval call binding the contract method 0xbfb98463.
//
// Solidity: function archaeologistSarcophagusCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistSarcophagusCount(opts *bind.CallOpts, archaeologist common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistSarcophagusCount", archaeologist)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArchaeologistSarcophagusCount is a free data retrieval call binding the contract method 0xbfb98463.
//
// Solidity: function archaeologistSarcophagusCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistSarcophagusCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistSarcophagusCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistSarcophagusCount is a free data retrieval call binding the contract method 0xbfb98463.
//
// Solidity: function archaeologistSarcophagusCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistSarcophagusCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistSarcophagusCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistSarcophagusIdentifier is a free data retrieval call binding the contract method 0x125a0c84.
//
// Solidity: function archaeologistSarcophagusIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistSarcophagusIdentifier(opts *bind.CallOpts, archaeologist common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistSarcophagusIdentifier", archaeologist, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ArchaeologistSarcophagusIdentifier is a free data retrieval call binding the contract method 0x125a0c84.
//
// Solidity: function archaeologistSarcophagusIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) ArchaeologistSarcophagusIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistSarcophagusIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistSarcophagusIdentifier is a free data retrieval call binding the contract method 0x125a0c84.
//
// Solidity: function archaeologistSarcophagusIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistSarcophagusIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistSarcophagusIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistSuccessesCount is a free data retrieval call binding the contract method 0xf1840ab2.
//
// Solidity: function archaeologistSuccessesCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistSuccessesCount(opts *bind.CallOpts, archaeologist common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistSuccessesCount", archaeologist)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArchaeologistSuccessesCount is a free data retrieval call binding the contract method 0xf1840ab2.
//
// Solidity: function archaeologistSuccessesCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) ArchaeologistSuccessesCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistSuccessesCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistSuccessesCount is a free data retrieval call binding the contract method 0xf1840ab2.
//
// Solidity: function archaeologistSuccessesCount(address archaeologist) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistSuccessesCount(archaeologist common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.ArchaeologistSuccessesCount(&_Sarcophagus.CallOpts, archaeologist)
}

// ArchaeologistSuccessesIdentifier is a free data retrieval call binding the contract method 0x284f34de.
//
// Solidity: function archaeologistSuccessesIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) ArchaeologistSuccessesIdentifier(opts *bind.CallOpts, archaeologist common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologistSuccessesIdentifier", archaeologist, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ArchaeologistSuccessesIdentifier is a free data retrieval call binding the contract method 0x284f34de.
//
// Solidity: function archaeologistSuccessesIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) ArchaeologistSuccessesIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistSuccessesIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// ArchaeologistSuccessesIdentifier is a free data retrieval call binding the contract method 0x284f34de.
//
// Solidity: function archaeologistSuccessesIdentifier(address archaeologist, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) ArchaeologistSuccessesIdentifier(archaeologist common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.ArchaeologistSuccessesIdentifier(&_Sarcophagus.CallOpts, archaeologist, index)
}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address account) view returns((bool,bytes,string,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Sarcophagus *SarcophagusCaller) Archaeologists(opts *bind.CallOpts, account common.Address) (TypesArchaeologist, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "archaeologists", account)

	if err != nil {
		return *new(TypesArchaeologist), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesArchaeologist)).(*TypesArchaeologist)

	return out0, err

}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address account) view returns((bool,bytes,string,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Sarcophagus *SarcophagusSession) Archaeologists(account common.Address) (TypesArchaeologist, error) {
	return _Sarcophagus.Contract.Archaeologists(&_Sarcophagus.CallOpts, account)
}

// Archaeologists is a free data retrieval call binding the contract method 0xf8fc6f7b.
//
// Solidity: function archaeologists(address account) view returns((bool,bytes,string,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Sarcophagus *SarcophagusCallerSession) Archaeologists(account common.Address) (TypesArchaeologist, error) {
	return _Sarcophagus.Contract.Archaeologists(&_Sarcophagus.CallOpts, account)
}

// EmbalmerSarcophagusCount is a free data retrieval call binding the contract method 0xfcbbb735.
//
// Solidity: function embalmerSarcophagusCount(address embalmer) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) EmbalmerSarcophagusCount(opts *bind.CallOpts, embalmer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "embalmerSarcophagusCount", embalmer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmbalmerSarcophagusCount is a free data retrieval call binding the contract method 0xfcbbb735.
//
// Solidity: function embalmerSarcophagusCount(address embalmer) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) EmbalmerSarcophagusCount(embalmer common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.EmbalmerSarcophagusCount(&_Sarcophagus.CallOpts, embalmer)
}

// EmbalmerSarcophagusCount is a free data retrieval call binding the contract method 0xfcbbb735.
//
// Solidity: function embalmerSarcophagusCount(address embalmer) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) EmbalmerSarcophagusCount(embalmer common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.EmbalmerSarcophagusCount(&_Sarcophagus.CallOpts, embalmer)
}

// EmbalmerSarcophagusIdentifier is a free data retrieval call binding the contract method 0x072dcee4.
//
// Solidity: function embalmerSarcophagusIdentifier(address embalmer, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) EmbalmerSarcophagusIdentifier(opts *bind.CallOpts, embalmer common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "embalmerSarcophagusIdentifier", embalmer, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EmbalmerSarcophagusIdentifier is a free data retrieval call binding the contract method 0x072dcee4.
//
// Solidity: function embalmerSarcophagusIdentifier(address embalmer, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) EmbalmerSarcophagusIdentifier(embalmer common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.EmbalmerSarcophagusIdentifier(&_Sarcophagus.CallOpts, embalmer, index)
}

// EmbalmerSarcophagusIdentifier is a free data retrieval call binding the contract method 0x072dcee4.
//
// Solidity: function embalmerSarcophagusIdentifier(address embalmer, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) EmbalmerSarcophagusIdentifier(embalmer common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.EmbalmerSarcophagusIdentifier(&_Sarcophagus.CallOpts, embalmer, index)
}

// RecipientSarcophagusCount is a free data retrieval call binding the contract method 0x70965e59.
//
// Solidity: function recipientSarcophagusCount(address recipient) view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) RecipientSarcophagusCount(opts *bind.CallOpts, recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "recipientSarcophagusCount", recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecipientSarcophagusCount is a free data retrieval call binding the contract method 0x70965e59.
//
// Solidity: function recipientSarcophagusCount(address recipient) view returns(uint256)
func (_Sarcophagus *SarcophagusSession) RecipientSarcophagusCount(recipient common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.RecipientSarcophagusCount(&_Sarcophagus.CallOpts, recipient)
}

// RecipientSarcophagusCount is a free data retrieval call binding the contract method 0x70965e59.
//
// Solidity: function recipientSarcophagusCount(address recipient) view returns(uint256)
func (_Sarcophagus *SarcophagusCallerSession) RecipientSarcophagusCount(recipient common.Address) (*big.Int, error) {
	return _Sarcophagus.Contract.RecipientSarcophagusCount(&_Sarcophagus.CallOpts, recipient)
}

// RecipientSarcophagusIdentifier is a free data retrieval call binding the contract method 0xe58def65.
//
// Solidity: function recipientSarcophagusIdentifier(address recipient, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) RecipientSarcophagusIdentifier(opts *bind.CallOpts, recipient common.Address, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "recipientSarcophagusIdentifier", recipient, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecipientSarcophagusIdentifier is a free data retrieval call binding the contract method 0xe58def65.
//
// Solidity: function recipientSarcophagusIdentifier(address recipient, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) RecipientSarcophagusIdentifier(recipient common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.RecipientSarcophagusIdentifier(&_Sarcophagus.CallOpts, recipient, index)
}

// RecipientSarcophagusIdentifier is a free data retrieval call binding the contract method 0xe58def65.
//
// Solidity: function recipientSarcophagusIdentifier(address recipient, uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) RecipientSarcophagusIdentifier(recipient common.Address, index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.RecipientSarcophagusIdentifier(&_Sarcophagus.CallOpts, recipient, index)
}

// SarcoToken is a free data retrieval call binding the contract method 0x607e8f09.
//
// Solidity: function sarcoToken() view returns(address)
func (_Sarcophagus *SarcophagusCaller) SarcoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "sarcoToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 identifier) view returns((uint8,address,bytes,address,string,uint256,uint256,string,bytes,uint256,uint256,uint256,uint256,bytes32))
func (_Sarcophagus *SarcophagusCaller) Sarcophagus(opts *bind.CallOpts, identifier [32]byte) (TypesSarcophagus, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "sarcophagus", identifier)

	if err != nil {
		return *new(TypesSarcophagus), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesSarcophagus)).(*TypesSarcophagus)

	return out0, err

}

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 identifier) view returns((uint8,address,bytes,address,string,uint256,uint256,string,bytes,uint256,uint256,uint256,uint256,bytes32))
func (_Sarcophagus *SarcophagusSession) Sarcophagus(identifier [32]byte) (TypesSarcophagus, error) {
	return _Sarcophagus.Contract.Sarcophagus(&_Sarcophagus.CallOpts, identifier)
}

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 identifier) view returns((uint8,address,bytes,address,string,uint256,uint256,string,bytes,uint256,uint256,uint256,uint256,bytes32))
func (_Sarcophagus *SarcophagusCallerSession) Sarcophagus(identifier [32]byte) (TypesSarcophagus, error) {
	return _Sarcophagus.Contract.Sarcophagus(&_Sarcophagus.CallOpts, identifier)
}

// SarcophagusCount is a free data retrieval call binding the contract method 0x555912b7.
//
// Solidity: function sarcophagusCount() view returns(uint256)
func (_Sarcophagus *SarcophagusCaller) SarcophagusCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "sarcophagusCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// SarcophagusIdentifier is a free data retrieval call binding the contract method 0x1a8800a1.
//
// Solidity: function sarcophagusIdentifier(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) SarcophagusIdentifier(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Sarcophagus.contract.Call(opts, &out, "sarcophagusIdentifier", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SarcophagusIdentifier is a free data retrieval call binding the contract method 0x1a8800a1.
//
// Solidity: function sarcophagusIdentifier(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) SarcophagusIdentifier(index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.SarcophagusIdentifier(&_Sarcophagus.CallOpts, index)
}

// SarcophagusIdentifier is a free data retrieval call binding the contract method 0x1a8800a1.
//
// Solidity: function sarcophagusIdentifier(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) SarcophagusIdentifier(index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.SarcophagusIdentifier(&_Sarcophagus.CallOpts, index)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 identifier, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) AccuseArchaeologist(opts *bind.TransactOpts, identifier [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "accuseArchaeologist", identifier, singleHash, paymentAddress)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 identifier, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusSession) AccuseArchaeologist(identifier [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.AccuseArchaeologist(&_Sarcophagus.TransactOpts, identifier, singleHash, paymentAddress)
}

// AccuseArchaeologist is a paid mutator transaction binding the contract method 0xf456f367.
//
// Solidity: function accuseArchaeologist(bytes32 identifier, bytes singleHash, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) AccuseArchaeologist(identifier [32]byte, singleHash []byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.AccuseArchaeologist(&_Sarcophagus.TransactOpts, identifier, singleHash, paymentAddress)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) BurySarcophagus(opts *bind.TransactOpts, identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "burySarcophagus", identifier)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusSession) BurySarcophagus(identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.BurySarcophagus(&_Sarcophagus.TransactOpts, identifier)
}

// BurySarcophagus is a paid mutator transaction binding the contract method 0x0511f2ec.
//
// Solidity: function burySarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) BurySarcophagus(identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.BurySarcophagus(&_Sarcophagus.TransactOpts, identifier)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CancelSarcophagus(opts *bind.TransactOpts, identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "cancelSarcophagus", identifier)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusSession) CancelSarcophagus(identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CancelSarcophagus(&_Sarcophagus.TransactOpts, identifier)
}

// CancelSarcophagus is a paid mutator transaction binding the contract method 0xa2b1fb3c.
//
// Solidity: function cancelSarcophagus(bytes32 identifier) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CancelSarcophagus(identifier [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CancelSarcophagus(&_Sarcophagus.TransactOpts, identifier)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 identifier, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CleanUpSarcophagus(opts *bind.TransactOpts, identifier [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "cleanUpSarcophagus", identifier, paymentAddress)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 identifier, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusSession) CleanUpSarcophagus(identifier [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CleanUpSarcophagus(&_Sarcophagus.TransactOpts, identifier, paymentAddress)
}

// CleanUpSarcophagus is a paid mutator transaction binding the contract method 0x7d99cda1.
//
// Solidity: function cleanUpSarcophagus(bytes32 identifier, address paymentAddress) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CleanUpSarcophagus(identifier [32]byte, paymentAddress common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CleanUpSarcophagus(&_Sarcophagus.TransactOpts, identifier, paymentAddress)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 identifier, bytes recipientPublicKey) returns(uint256)
func (_Sarcophagus *SarcophagusTransactor) CreateSarcophagus(opts *bind.TransactOpts, name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, identifier [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "createSarcophagus", name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, identifier, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 identifier, bytes recipientPublicKey) returns(uint256)
func (_Sarcophagus *SarcophagusSession) CreateSarcophagus(name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, identifier [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, identifier, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 identifier, bytes recipientPublicKey) returns(uint256)
func (_Sarcophagus *SarcophagusTransactorSession) CreateSarcophagus(name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, identifier [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, identifier, recipientPublicKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _sarcoToken) returns()
func (_Sarcophagus *SarcophagusTransactor) Initialize(opts *bind.TransactOpts, _sarcoToken common.Address) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "initialize", _sarcoToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _sarcoToken) returns()
func (_Sarcophagus *SarcophagusSession) Initialize(_sarcoToken common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.Initialize(&_Sarcophagus.TransactOpts, _sarcoToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _sarcoToken) returns()
func (_Sarcophagus *SarcophagusTransactorSession) Initialize(_sarcoToken common.Address) (*types.Transaction, error) {
	return _Sarcophagus.Contract.Initialize(&_Sarcophagus.TransactOpts, _sarcoToken)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(uint256)
func (_Sarcophagus *SarcophagusTransactor) RegisterArchaeologist(opts *bind.TransactOpts, currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "registerArchaeologist", currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(uint256)
func (_Sarcophagus *SarcophagusSession) RegisterArchaeologist(currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(uint256)
func (_Sarcophagus *SarcophagusTransactorSession) RegisterArchaeologist(currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 identifier, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) RewrapSarcophagus(opts *bind.TransactOpts, identifier [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "rewrapSarcophagus", identifier, resurrectionTime, diggingFee, bounty)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 identifier, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusSession) RewrapSarcophagus(identifier [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RewrapSarcophagus(&_Sarcophagus.TransactOpts, identifier, resurrectionTime, diggingFee, bounty)
}

// RewrapSarcophagus is a paid mutator transaction binding the contract method 0xbf1a9d77.
//
// Solidity: function rewrapSarcophagus(bytes32 identifier, uint256 resurrectionTime, uint256 diggingFee, uint256 bounty) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) RewrapSarcophagus(identifier [32]byte, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RewrapSarcophagus(&_Sarcophagus.TransactOpts, identifier, resurrectionTime, diggingFee, bounty)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0xf4a5574b.
//
// Solidity: function unwrapSarcophagus(bytes32 identifier, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UnwrapSarcophagus(opts *bind.TransactOpts, identifier [32]byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "unwrapSarcophagus", identifier, privateKey)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0xf4a5574b.
//
// Solidity: function unwrapSarcophagus(bytes32 identifier, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusSession) UnwrapSarcophagus(identifier [32]byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, identifier, privateKey)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0xf4a5574b.
//
// Solidity: function unwrapSarcophagus(bytes32 identifier, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UnwrapSarcophagus(identifier [32]byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, identifier, privateKey)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes newPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateArchaeologist(opts *bind.TransactOpts, endpoint string, newPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateArchaeologist", endpoint, newPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes newPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateArchaeologist(endpoint string, newPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, newPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes newPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateArchaeologist(endpoint string, newPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, newPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 identifier, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateSarcophagus(opts *bind.TransactOpts, newPublicKey []byte, identifier [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateSarcophagus", newPublicKey, identifier, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 identifier, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateSarcophagus(newPublicKey []byte, identifier [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, newPublicKey, identifier, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 identifier, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateSarcophagus(newPublicKey []byte, identifier [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, newPublicKey, identifier, assetId, v, r, s)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) WithdrawBond(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "withdrawBond", amount)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusSession) WithdrawBond(amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.WithdrawBond(&_Sarcophagus.TransactOpts, amount)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) WithdrawBond(amount *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.WithdrawBond(&_Sarcophagus.TransactOpts, amount)
}
