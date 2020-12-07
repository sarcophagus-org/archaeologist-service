// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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

// ArchaeologistsABI is the input ABI used to generate the binding from.
const ArchaeologistsABI = "[]"

// ArchaeologistsBin is the compiled bytecode used for deploying new contracts.
var ArchaeologistsBin = "0x611984610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c80635d83aa67116100655780635d83aa67146101d157806377a91eac146103c3578063ca03a0c21461041d578063e6b72c961461060f57610087565b80630460a24b1461008c5780630d9616ca1461010757806310dfb3751461016c575b600080fd5b81801561009857600080fd5b506100ef600480360360608110156100af57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610674565b60405180821515815260200191505060405180910390f35b81801561011357600080fd5b5061016a6004803603606081101561012a57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610821565b005b81801561017857600080fd5b506101cf6004803603606081101561018f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061083c565b005b8180156101dd57600080fd5b506103ab60048036036101408110156101f557600080fd5b81019080803590602001909291908035906020019064010000000081111561021c57600080fd5b82018360208201111561022e57600080fd5b8035906020019184600183028401116401000000008311171561025057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102b357600080fd5b8201836020820111156102c557600080fd5b803590602001918460018302840111640100000000831117156102e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108a6565b60405180821515815260200191505060405180910390f35b61041b600480360360608110156103d957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190505050610e41565b005b81801561042957600080fd5b506105f7600480360361014081101561044157600080fd5b81019080803590602001909291908035906020019064010000000081111561046857600080fd5b82018360208201111561047a57600080fd5b8035906020019184600183028401116401000000008311171561049c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156104ff57600080fd5b82018360208201111561051157600080fd5b8035906020019184600183028401116401000000008311171561053357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f85565b60405180821515815260200191505060405180910390f35b81801561061b57600080fd5b506106726004803603606081101561063257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611426565b005b600061068284336001610e41565b60008460000160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506106d2853386611441565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561076757600080fd5b505af115801561077b573d6000803e3d6000fd5b505050506040513d602081101561079157600080fd5b8101908080519060200190929190505050508060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f856040518082815260200191505060405180910390a260019150509392505050565b61082c8383836114c5565b61083783838361083c565b505050565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905061089882826009015461152f90919063ffffffff16565b816009018190555050505050565b60006108b48b336000610e41565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__63495fe0078b6040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561091e578082015181840152602081019050610903565b50505050905090810190601f16801561094b5780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561096857600080fd5b505af415801561097c573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610a0f57600080fd5b505af1158015610a23573d6000803e3d6000fd5b505050506040513d6020811015610a3957600080fd5b810190808051906020019092919050505050610a5361172b565b6040518061016001604052806001151581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018c81526020018b81526020018a73ffffffffffffffffffffffffffffffffffffffff16815260200189815260200188815260200187815260200186815260200185815260200160008152509050808c60000160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001019080519060200190610b969291906117b3565b506060820151816002019080519060200190610bb3929190611833565b5060808201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816004015560c0820151816005015560e082015181600601556101008201518160070155610120820151816008015561014082015181600901559050508b600101339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806020015173ffffffffffffffffffffffffffffffffffffffff167f6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f8260400151836060015184608001518560a001518660c001518760e001518861010001518961012001516040518080602001806020018973ffffffffffffffffffffffffffffffffffffffff16815260200188815260200187815260200186815260200185815260200184815260200183810383528b818151815260200191508051906020019080838360005b83811015610d86578082015181840152602081019050610d6b565b50505050905090810190601f168015610db35780820380516001836020036101000a031916815260200191505b5083810382528a818151815260200191508051906020019080838360005b83811015610dec578082015181840152602081019050610dd1565b50505050905090810190601f168015610e195780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390a260019150509a9950505050505050505050565b60606040518060600160405280602981526020016118d160299139905081610e7f576040518060600160405280602981526020016118fa6029913990505b8115158460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff161515148190610f7e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f43578082015181840152602081019050610f28565b50505050905090810190601f168015610f705780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5050505050565b6000610f938b336001610e41565b60008b60000160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508a8051906020012081600101604051808280546001816001161561010002031660029004801561103f5780601f1061101d57610100808354040283529182019161103f565b820191906000526020600020905b81548152906001019060200180831161102b575b5050915050604051809103902014611142578060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a178c6040518080602001828103825283818151815260200191508051906020019080838360005b838110156110ee5780820151818401526020810190506110d3565b50505050905090810190601f16801561111b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a28a8160010190805190602001906111409291906117b3565b505b8981600201908051906020019061115a929190611833565b50888160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555087816004018190555086816005018190555085816006018190555084816007018190555060008411156112a2576111d68c33866114c5565b8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561126557600080fd5b505af1158015611279573d6000803e3d6000fd5b505050506040513d602081101561128f57600080fd5b8101908080519060200190929190505050505b8060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f826002018360030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600401548560050154866006015487600701548b60405180806020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020018581526020018481526020018381526020018281038252898181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156113ff5780601f106113d4576101008083540402835291602001916113ff565b820191906000526020600020905b8154815290600101906020018083116113e257829003601f168201915b50509850505050505050505060405180910390a260019150509a9950505050505050505050565b611431838383611441565b61143c838383611579565b505050565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506114b7826040518060600160405280602c8152602001611923602c913983600801546115e39092919063ffffffff16565b816008018190555050505050565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506115218282600801546116a390919063ffffffff16565b816008018190555050505050565b600061157183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506115e3565b905092915050565b60008360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506115d58282600901546116a390919063ffffffff16565b816009018190555050505050565b6000838311158290611690576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561165557808201518184015260208101905061163a565b50505050905090810190601f1680156116825780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015611721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b604051806101600160405280600015158152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117f457805160ff1916838001178555611822565b82800160010185558215611822579182015b82811115611821578251825591602001919060010190611806565b5b50905061182f91906118b3565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061187457805160ff19168380011785556118a2565b828001600101855582156118a2579182015b828111156118a1578251825591602001919060010190611886565b5b5090506118af91906118b3565b5090565b5b808211156118cc5760008160009055506001016118b4565b509056fe6172636861656f6c6f6769737420686173206e6f74206265656e2072656769737465726564207965746172636861656f6c6f676973742068617320616c7265616479206265656e20726567697374657265646172636861656f6c6f6769737420646f6573206e6f74206861766520656e6f756768206672656520626f6e64a26469706673582212209cf61007fc405cf9f99fa57eddee4f7af7be124477d714926482fd64fbf8efb364736f6c63430007020033"

// DeployArchaeologists deploys a new Ethereum contract, binding an instance of Archaeologists to it.
func DeployArchaeologists(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Archaeologists, error) {
	parsed, err := abi.JSON(strings.NewReader(ArchaeologistsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	utilsAddr, _, _, _ := DeployUtils(auth, backend)
	ArchaeologistsBin = strings.Replace(ArchaeologistsBin, "__$b4ae538d42a0e0bb8330684d08e51a40b4$__", utilsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArchaeologistsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Archaeologists{ArchaeologistsCaller: ArchaeologistsCaller{contract: contract}, ArchaeologistsTransactor: ArchaeologistsTransactor{contract: contract}, ArchaeologistsFilterer: ArchaeologistsFilterer{contract: contract}}, nil
}

// Archaeologists is an auto generated Go binding around an Ethereum contract.
type Archaeologists struct {
	ArchaeologistsCaller     // Read-only binding to the contract
	ArchaeologistsTransactor // Write-only binding to the contract
	ArchaeologistsFilterer   // Log filterer for contract events
}

// ArchaeologistsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArchaeologistsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArchaeologistsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArchaeologistsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArchaeologistsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArchaeologistsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArchaeologistsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArchaeologistsSession struct {
	Contract     *Archaeologists   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArchaeologistsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArchaeologistsCallerSession struct {
	Contract *ArchaeologistsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ArchaeologistsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArchaeologistsTransactorSession struct {
	Contract     *ArchaeologistsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ArchaeologistsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArchaeologistsRaw struct {
	Contract *Archaeologists // Generic contract binding to access the raw methods on
}

// ArchaeologistsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArchaeologistsCallerRaw struct {
	Contract *ArchaeologistsCaller // Generic read-only contract binding to access the raw methods on
}

// ArchaeologistsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArchaeologistsTransactorRaw struct {
	Contract *ArchaeologistsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArchaeologists creates a new instance of Archaeologists, bound to a specific deployed contract.
func NewArchaeologists(address common.Address, backend bind.ContractBackend) (*Archaeologists, error) {
	contract, err := bindArchaeologists(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Archaeologists{ArchaeologistsCaller: ArchaeologistsCaller{contract: contract}, ArchaeologistsTransactor: ArchaeologistsTransactor{contract: contract}, ArchaeologistsFilterer: ArchaeologistsFilterer{contract: contract}}, nil
}

// NewArchaeologistsCaller creates a new read-only instance of Archaeologists, bound to a specific deployed contract.
func NewArchaeologistsCaller(address common.Address, caller bind.ContractCaller) (*ArchaeologistsCaller, error) {
	contract, err := bindArchaeologists(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArchaeologistsCaller{contract: contract}, nil
}

// NewArchaeologistsTransactor creates a new write-only instance of Archaeologists, bound to a specific deployed contract.
func NewArchaeologistsTransactor(address common.Address, transactor bind.ContractTransactor) (*ArchaeologistsTransactor, error) {
	contract, err := bindArchaeologists(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArchaeologistsTransactor{contract: contract}, nil
}

// NewArchaeologistsFilterer creates a new log filterer instance of Archaeologists, bound to a specific deployed contract.
func NewArchaeologistsFilterer(address common.Address, filterer bind.ContractFilterer) (*ArchaeologistsFilterer, error) {
	contract, err := bindArchaeologists(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArchaeologistsFilterer{contract: contract}, nil
}

// bindArchaeologists binds a generic wrapper to an already deployed contract.
func bindArchaeologists(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArchaeologistsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Archaeologists *ArchaeologistsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Archaeologists.Contract.ArchaeologistsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Archaeologists *ArchaeologistsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Archaeologists.Contract.ArchaeologistsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Archaeologists *ArchaeologistsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Archaeologists.Contract.ArchaeologistsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Archaeologists *ArchaeologistsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Archaeologists.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Archaeologists *ArchaeologistsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Archaeologists.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Archaeologists *ArchaeologistsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Archaeologists.Contract.contract.Transact(opts, method, params...)
}

// DatasABI is the input ABI used to generate the binding from.
const DatasABI = "[]"

// DatasBin is the compiled bytecode used for deploying new contracts.
var DatasBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f25fef85ae5e4a533820a1d80a97140bf64281da952e7dd7383318bc482f0c6a64736f6c63430007020033"

// DeployDatas deploys a new Ethereum contract, binding an instance of Datas to it.
func DeployDatas(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Datas, error) {
	parsed, err := abi.JSON(strings.NewReader(DatasABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatasBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datas{DatasCaller: DatasCaller{contract: contract}, DatasTransactor: DatasTransactor{contract: contract}, DatasFilterer: DatasFilterer{contract: contract}}, nil
}

// Datas is an auto generated Go binding around an Ethereum contract.
type Datas struct {
	DatasCaller     // Read-only binding to the contract
	DatasTransactor // Write-only binding to the contract
	DatasFilterer   // Log filterer for contract events
}

// DatasCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatasSession struct {
	Contract     *Datas            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatasCallerSession struct {
	Contract *DatasCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DatasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatasTransactorSession struct {
	Contract     *DatasTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatasRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatasRaw struct {
	Contract *Datas // Generic contract binding to access the raw methods on
}

// DatasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatasCallerRaw struct {
	Contract *DatasCaller // Generic read-only contract binding to access the raw methods on
}

// DatasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatasTransactorRaw struct {
	Contract *DatasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatas creates a new instance of Datas, bound to a specific deployed contract.
func NewDatas(address common.Address, backend bind.ContractBackend) (*Datas, error) {
	contract, err := bindDatas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datas{DatasCaller: DatasCaller{contract: contract}, DatasTransactor: DatasTransactor{contract: contract}, DatasFilterer: DatasFilterer{contract: contract}}, nil
}

// NewDatasCaller creates a new read-only instance of Datas, bound to a specific deployed contract.
func NewDatasCaller(address common.Address, caller bind.ContractCaller) (*DatasCaller, error) {
	contract, err := bindDatas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatasCaller{contract: contract}, nil
}

// NewDatasTransactor creates a new write-only instance of Datas, bound to a specific deployed contract.
func NewDatasTransactor(address common.Address, transactor bind.ContractTransactor) (*DatasTransactor, error) {
	contract, err := bindDatas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatasTransactor{contract: contract}, nil
}

// NewDatasFilterer creates a new log filterer instance of Datas, bound to a specific deployed contract.
func NewDatasFilterer(address common.Address, filterer bind.ContractFilterer) (*DatasFilterer, error) {
	contract, err := bindDatas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatasFilterer{contract: contract}, nil
}

// bindDatas binds a generic wrapper to an already deployed contract.
func bindDatas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatasABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datas *DatasRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datas.Contract.DatasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datas *DatasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datas.Contract.DatasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datas *DatasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datas.Contract.DatasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datas *DatasCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Datas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datas *DatasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datas *DatasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datas.Contract.contract.Transact(opts, method, params...)
}

// EventsABI is the input ABI used to generate the binding from.
const EventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accuserBondReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"embalmerBondReward\",\"type\":\"uint256\"}],\"name\":\"AccuseArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"BurySarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"CancelSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cleaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cleanerBondReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"embalmerBondReward\",\"type\":\"uint256\"}],\"name\":\"CleanUpSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"archaeologistPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"name\":\"CreateSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sarcophagusContract\",\"type\":\"address\"}],\"name\":\"Creation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"}],\"name\":\"RegisterArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"name\":\"RewrapSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"privatekey\",\"type\":\"bytes32\"}],\"name\":\"UnwrapSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedBond\",\"type\":\"uint256\"}],\"name\":\"UpdateArchaeologist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"}],\"name\":\"UpdateArchaeologistPublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"}],\"name\":\"UpdateSarcophagus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnBond\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFreeBond\",\"type\":\"event\"}]"

// EventsBin is the compiled bytecode used for deploying new contracts.
var EventsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aa2efafde06495a6e9391f93ab3aeaacc4b136c51e1fe5a0de417f8aa4a4100164736f6c63430007020033"

// DeployEvents deploys a new Ethereum contract, binding an instance of Events to it.
func DeployEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Events, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EventsAccuseArchaeologistIterator is returned from FilterAccuseArchaeologist and is used to iterate over the raw logs and unpacked data for AccuseArchaeologist events raised by the Events contract.
type EventsAccuseArchaeologistIterator struct {
	Event *EventsAccuseArchaeologist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsAccuseArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsAccuseArchaeologist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsAccuseArchaeologist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsAccuseArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsAccuseArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsAccuseArchaeologist represents a AccuseArchaeologist event raised by the Events contract.
type EventsAccuseArchaeologist struct {
	AssetDoubleHash    [32]byte
	Accuser            common.Address
	AccuserBondReward  *big.Int
	EmbalmerBondReward *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAccuseArchaeologist is a free log retrieval operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) FilterAccuseArchaeologist(opts *bind.FilterOpts, assetDoubleHash [][32]byte, accuser []common.Address) (*EventsAccuseArchaeologistIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "AccuseArchaeologist", assetDoubleHashRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return &EventsAccuseArchaeologistIterator{contract: _Events.contract, event: "AccuseArchaeologist", logs: logs, sub: sub}, nil
}

// WatchAccuseArchaeologist is a free log subscription operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) WatchAccuseArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsAccuseArchaeologist, assetDoubleHash [][32]byte, accuser []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "AccuseArchaeologist", assetDoubleHashRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsAccuseArchaeologist)
				if err := _Events.contract.UnpackLog(event, "AccuseArchaeologist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccuseArchaeologist is a log parse operation binding the contract event 0x6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd4.
//
// Solidity: event AccuseArchaeologist(bytes32 indexed assetDoubleHash, address indexed accuser, uint256 accuserBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) ParseAccuseArchaeologist(log types.Log) (*EventsAccuseArchaeologist, error) {
	event := new(EventsAccuseArchaeologist)
	if err := _Events.contract.UnpackLog(event, "AccuseArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsBurySarcophagusIterator is returned from FilterBurySarcophagus and is used to iterate over the raw logs and unpacked data for BurySarcophagus events raised by the Events contract.
type EventsBurySarcophagusIterator struct {
	Event *EventsBurySarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsBurySarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsBurySarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsBurySarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsBurySarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsBurySarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsBurySarcophagus represents a BurySarcophagus event raised by the Events contract.
type EventsBurySarcophagus struct {
	AssetDoubleHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurySarcophagus is a free log retrieval operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) FilterBurySarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsBurySarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "BurySarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsBurySarcophagusIterator{contract: _Events.contract, event: "BurySarcophagus", logs: logs, sub: sub}, nil
}

// WatchBurySarcophagus is a free log subscription operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) WatchBurySarcophagus(opts *bind.WatchOpts, sink chan<- *EventsBurySarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "BurySarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsBurySarcophagus)
				if err := _Events.contract.UnpackLog(event, "BurySarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurySarcophagus is a log parse operation binding the contract event 0x3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a.
//
// Solidity: event BurySarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) ParseBurySarcophagus(log types.Log) (*EventsBurySarcophagus, error) {
	event := new(EventsBurySarcophagus)
	if err := _Events.contract.UnpackLog(event, "BurySarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCancelSarcophagusIterator is returned from FilterCancelSarcophagus and is used to iterate over the raw logs and unpacked data for CancelSarcophagus events raised by the Events contract.
type EventsCancelSarcophagusIterator struct {
	Event *EventsCancelSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsCancelSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCancelSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsCancelSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsCancelSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCancelSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCancelSarcophagus represents a CancelSarcophagus event raised by the Events contract.
type EventsCancelSarcophagus struct {
	AssetDoubleHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelSarcophagus is a free log retrieval operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) FilterCancelSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsCancelSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CancelSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsCancelSarcophagusIterator{contract: _Events.contract, event: "CancelSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCancelSarcophagus is a free log subscription operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) WatchCancelSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCancelSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CancelSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCancelSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CancelSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelSarcophagus is a log parse operation binding the contract event 0x829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e7750.
//
// Solidity: event CancelSarcophagus(bytes32 indexed assetDoubleHash)
func (_Events *EventsFilterer) ParseCancelSarcophagus(log types.Log) (*EventsCancelSarcophagus, error) {
	event := new(EventsCancelSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CancelSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCleanUpSarcophagusIterator is returned from FilterCleanUpSarcophagus and is used to iterate over the raw logs and unpacked data for CleanUpSarcophagus events raised by the Events contract.
type EventsCleanUpSarcophagusIterator struct {
	Event *EventsCleanUpSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsCleanUpSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCleanUpSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsCleanUpSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsCleanUpSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCleanUpSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCleanUpSarcophagus represents a CleanUpSarcophagus event raised by the Events contract.
type EventsCleanUpSarcophagus struct {
	AssetDoubleHash    [32]byte
	Cleaner            common.Address
	CleanerBondReward  *big.Int
	EmbalmerBondReward *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCleanUpSarcophagus is a free log retrieval operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) FilterCleanUpSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte, cleaner []common.Address) (*EventsCleanUpSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var cleanerRule []interface{}
	for _, cleanerItem := range cleaner {
		cleanerRule = append(cleanerRule, cleanerItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CleanUpSarcophagus", assetDoubleHashRule, cleanerRule)
	if err != nil {
		return nil, err
	}
	return &EventsCleanUpSarcophagusIterator{contract: _Events.contract, event: "CleanUpSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCleanUpSarcophagus is a free log subscription operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) WatchCleanUpSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCleanUpSarcophagus, assetDoubleHash [][32]byte, cleaner []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var cleanerRule []interface{}
	for _, cleanerItem := range cleaner {
		cleanerRule = append(cleanerRule, cleanerItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CleanUpSarcophagus", assetDoubleHashRule, cleanerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCleanUpSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CleanUpSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCleanUpSarcophagus is a log parse operation binding the contract event 0x106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb9.
//
// Solidity: event CleanUpSarcophagus(bytes32 indexed assetDoubleHash, address indexed cleaner, uint256 cleanerBondReward, uint256 embalmerBondReward)
func (_Events *EventsFilterer) ParseCleanUpSarcophagus(log types.Log) (*EventsCleanUpSarcophagus, error) {
	event := new(EventsCleanUpSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CleanUpSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCreateSarcophagusIterator is returned from FilterCreateSarcophagus and is used to iterate over the raw logs and unpacked data for CreateSarcophagus events raised by the Events contract.
type EventsCreateSarcophagusIterator struct {
	Event *EventsCreateSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsCreateSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCreateSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsCreateSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsCreateSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCreateSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCreateSarcophagus represents a CreateSarcophagus event raised by the Events contract.
type EventsCreateSarcophagus struct {
	AssetDoubleHash        [32]byte
	Archaeologist          common.Address
	ArchaeologistPublicKey []byte
	Embalmer               common.Address
	Name                   string
	ResurrectionTime       *big.Int
	ResurrectionWindow     *big.Int
	StorageFee             *big.Int
	DiggingFee             *big.Int
	Bounty                 *big.Int
	RecipientPublicKey     []byte
	CursedBond             *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCreateSarcophagus is a free log retrieval operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) FilterCreateSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte, archaeologist []common.Address) (*EventsCreateSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "CreateSarcophagus", assetDoubleHashRule, archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsCreateSarcophagusIterator{contract: _Events.contract, event: "CreateSarcophagus", logs: logs, sub: sub}, nil
}

// WatchCreateSarcophagus is a free log subscription operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) WatchCreateSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsCreateSarcophagus, assetDoubleHash [][32]byte, archaeologist []common.Address) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}
	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "CreateSarcophagus", assetDoubleHashRule, archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCreateSarcophagus)
				if err := _Events.contract.UnpackLog(event, "CreateSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateSarcophagus is a log parse operation binding the contract event 0x61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f850.
//
// Solidity: event CreateSarcophagus(bytes32 indexed assetDoubleHash, address indexed archaeologist, bytes archaeologistPublicKey, address embalmer, string name, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes recipientPublicKey, uint256 cursedBond)
func (_Events *EventsFilterer) ParseCreateSarcophagus(log types.Log) (*EventsCreateSarcophagus, error) {
	event := new(EventsCreateSarcophagus)
	if err := _Events.contract.UnpackLog(event, "CreateSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsCreationIterator is returned from FilterCreation and is used to iterate over the raw logs and unpacked data for Creation events raised by the Events contract.
type EventsCreationIterator struct {
	Event *EventsCreation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsCreation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsCreation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsCreation represents a Creation event raised by the Events contract.
type EventsCreation struct {
	SarcophagusContract common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCreation is a free log retrieval operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) FilterCreation(opts *bind.FilterOpts) (*EventsCreationIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "Creation")
	if err != nil {
		return nil, err
	}
	return &EventsCreationIterator{contract: _Events.contract, event: "Creation", logs: logs, sub: sub}, nil
}

// WatchCreation is a free log subscription operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) WatchCreation(opts *bind.WatchOpts, sink chan<- *EventsCreation) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "Creation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsCreation)
				if err := _Events.contract.UnpackLog(event, "Creation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreation is a log parse operation binding the contract event 0x2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c.
//
// Solidity: event Creation(address sarcophagusContract)
func (_Events *EventsFilterer) ParseCreation(log types.Log) (*EventsCreation, error) {
	event := new(EventsCreation)
	if err := _Events.contract.UnpackLog(event, "Creation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsRegisterArchaeologistIterator is returned from FilterRegisterArchaeologist and is used to iterate over the raw logs and unpacked data for RegisterArchaeologist events raised by the Events contract.
type EventsRegisterArchaeologistIterator struct {
	Event *EventsRegisterArchaeologist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsRegisterArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsRegisterArchaeologist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsRegisterArchaeologist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsRegisterArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsRegisterArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsRegisterArchaeologist represents a RegisterArchaeologist event raised by the Events contract.
type EventsRegisterArchaeologist struct {
	Archaeologist           common.Address
	CurrentPublicKey        []byte
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	Bond                    *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRegisterArchaeologist is a free log retrieval operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) FilterRegisterArchaeologist(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsRegisterArchaeologistIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "RegisterArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsRegisterArchaeologistIterator{contract: _Events.contract, event: "RegisterArchaeologist", logs: logs, sub: sub}, nil
}

// WatchRegisterArchaeologist is a free log subscription operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) WatchRegisterArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsRegisterArchaeologist, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "RegisterArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsRegisterArchaeologist)
				if err := _Events.contract.UnpackLog(event, "RegisterArchaeologist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterArchaeologist is a log parse operation binding the contract event 0x6f06bc5924965fb88f5a66a11a425a6f09d3abc6988f40dab496d5e9fede6f2f.
//
// Solidity: event RegisterArchaeologist(address indexed archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 bond)
func (_Events *EventsFilterer) ParseRegisterArchaeologist(log types.Log) (*EventsRegisterArchaeologist, error) {
	event := new(EventsRegisterArchaeologist)
	if err := _Events.contract.UnpackLog(event, "RegisterArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsRewrapSarcophagusIterator is returned from FilterRewrapSarcophagus and is used to iterate over the raw logs and unpacked data for RewrapSarcophagus events raised by the Events contract.
type EventsRewrapSarcophagusIterator struct {
	Event *EventsRewrapSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsRewrapSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsRewrapSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsRewrapSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsRewrapSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsRewrapSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsRewrapSarcophagus represents a RewrapSarcophagus event raised by the Events contract.
type EventsRewrapSarcophagus struct {
	AssetId            string
	AssetDoubleHash    [32]byte
	ResurrectionTime   *big.Int
	ResurrectionWindow *big.Int
	DiggingFee         *big.Int
	Bounty             *big.Int
	CursedBond         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewrapSarcophagus is a free log retrieval operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) FilterRewrapSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsRewrapSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "RewrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsRewrapSarcophagusIterator{contract: _Events.contract, event: "RewrapSarcophagus", logs: logs, sub: sub}, nil
}

// WatchRewrapSarcophagus is a free log subscription operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) WatchRewrapSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsRewrapSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "RewrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsRewrapSarcophagus)
				if err := _Events.contract.UnpackLog(event, "RewrapSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewrapSarcophagus is a log parse operation binding the contract event 0x17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46.
//
// Solidity: event RewrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, uint256 resurrectionTime, uint256 resurrectionWindow, uint256 diggingFee, uint256 bounty, uint256 cursedBond)
func (_Events *EventsFilterer) ParseRewrapSarcophagus(log types.Log) (*EventsRewrapSarcophagus, error) {
	event := new(EventsRewrapSarcophagus)
	if err := _Events.contract.UnpackLog(event, "RewrapSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUnwrapSarcophagusIterator is returned from FilterUnwrapSarcophagus and is used to iterate over the raw logs and unpacked data for UnwrapSarcophagus events raised by the Events contract.
type EventsUnwrapSarcophagusIterator struct {
	Event *EventsUnwrapSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsUnwrapSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUnwrapSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsUnwrapSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsUnwrapSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUnwrapSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUnwrapSarcophagus represents a UnwrapSarcophagus event raised by the Events contract.
type EventsUnwrapSarcophagus struct {
	AssetId         string
	AssetDoubleHash [32]byte
	SingleHash      []byte
	Privatekey      [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnwrapSarcophagus is a free log retrieval operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) FilterUnwrapSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsUnwrapSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UnwrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsUnwrapSarcophagusIterator{contract: _Events.contract, event: "UnwrapSarcophagus", logs: logs, sub: sub}, nil
}

// WatchUnwrapSarcophagus is a free log subscription operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) WatchUnwrapSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsUnwrapSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UnwrapSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUnwrapSarcophagus)
				if err := _Events.contract.UnpackLog(event, "UnwrapSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnwrapSarcophagus is a log parse operation binding the contract event 0x8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c788.
//
// Solidity: event UnwrapSarcophagus(string assetId, bytes32 indexed assetDoubleHash, bytes singleHash, bytes32 privatekey)
func (_Events *EventsFilterer) ParseUnwrapSarcophagus(log types.Log) (*EventsUnwrapSarcophagus, error) {
	event := new(EventsUnwrapSarcophagus)
	if err := _Events.contract.UnpackLog(event, "UnwrapSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateArchaeologistIterator is returned from FilterUpdateArchaeologist and is used to iterate over the raw logs and unpacked data for UpdateArchaeologist events raised by the Events contract.
type EventsUpdateArchaeologistIterator struct {
	Event *EventsUpdateArchaeologist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsUpdateArchaeologistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateArchaeologist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsUpdateArchaeologist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsUpdateArchaeologistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateArchaeologistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateArchaeologist represents a UpdateArchaeologist event raised by the Events contract.
type EventsUpdateArchaeologist struct {
	Archaeologist           common.Address
	Endpoint                string
	PaymentAddress          common.Address
	FeePerByte              *big.Int
	MinimumBounty           *big.Int
	MinimumDiggingFee       *big.Int
	MaximumResurrectionTime *big.Int
	AddedBond               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdateArchaeologist is a free log retrieval operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) FilterUpdateArchaeologist(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsUpdateArchaeologistIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateArchaeologistIterator{contract: _Events.contract, event: "UpdateArchaeologist", logs: logs, sub: sub}, nil
}

// WatchUpdateArchaeologist is a free log subscription operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) WatchUpdateArchaeologist(opts *bind.WatchOpts, sink chan<- *EventsUpdateArchaeologist, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateArchaeologist", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateArchaeologist)
				if err := _Events.contract.UnpackLog(event, "UpdateArchaeologist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateArchaeologist is a log parse operation binding the contract event 0x8efcecf1d91087696dce579361122ec7045fa8634810682fc50cfb1b1a2e302f.
//
// Solidity: event UpdateArchaeologist(address indexed archaeologist, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 addedBond)
func (_Events *EventsFilterer) ParseUpdateArchaeologist(log types.Log) (*EventsUpdateArchaeologist, error) {
	event := new(EventsUpdateArchaeologist)
	if err := _Events.contract.UnpackLog(event, "UpdateArchaeologist", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateArchaeologistPublicKeyIterator is returned from FilterUpdateArchaeologistPublicKey and is used to iterate over the raw logs and unpacked data for UpdateArchaeologistPublicKey events raised by the Events contract.
type EventsUpdateArchaeologistPublicKeyIterator struct {
	Event *EventsUpdateArchaeologistPublicKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsUpdateArchaeologistPublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateArchaeologistPublicKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsUpdateArchaeologistPublicKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsUpdateArchaeologistPublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateArchaeologistPublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateArchaeologistPublicKey represents a UpdateArchaeologistPublicKey event raised by the Events contract.
type EventsUpdateArchaeologistPublicKey struct {
	Archaeologist    common.Address
	CurrentPublicKey []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateArchaeologistPublicKey is a free log retrieval operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) FilterUpdateArchaeologistPublicKey(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsUpdateArchaeologistPublicKeyIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateArchaeologistPublicKey", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateArchaeologistPublicKeyIterator{contract: _Events.contract, event: "UpdateArchaeologistPublicKey", logs: logs, sub: sub}, nil
}

// WatchUpdateArchaeologistPublicKey is a free log subscription operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) WatchUpdateArchaeologistPublicKey(opts *bind.WatchOpts, sink chan<- *EventsUpdateArchaeologistPublicKey, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateArchaeologistPublicKey", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateArchaeologistPublicKey)
				if err := _Events.contract.UnpackLog(event, "UpdateArchaeologistPublicKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateArchaeologistPublicKey is a log parse operation binding the contract event 0x6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17.
//
// Solidity: event UpdateArchaeologistPublicKey(address indexed archaeologist, bytes currentPublicKey)
func (_Events *EventsFilterer) ParseUpdateArchaeologistPublicKey(log types.Log) (*EventsUpdateArchaeologistPublicKey, error) {
	event := new(EventsUpdateArchaeologistPublicKey)
	if err := _Events.contract.UnpackLog(event, "UpdateArchaeologistPublicKey", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsUpdateSarcophagusIterator is returned from FilterUpdateSarcophagus and is used to iterate over the raw logs and unpacked data for UpdateSarcophagus events raised by the Events contract.
type EventsUpdateSarcophagusIterator struct {
	Event *EventsUpdateSarcophagus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsUpdateSarcophagusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsUpdateSarcophagus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsUpdateSarcophagus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsUpdateSarcophagusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsUpdateSarcophagusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsUpdateSarcophagus represents a UpdateSarcophagus event raised by the Events contract.
type EventsUpdateSarcophagus struct {
	AssetDoubleHash [32]byte
	AssetId         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateSarcophagus is a free log retrieval operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) FilterUpdateSarcophagus(opts *bind.FilterOpts, assetDoubleHash [][32]byte) (*EventsUpdateSarcophagusIterator, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "UpdateSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return &EventsUpdateSarcophagusIterator{contract: _Events.contract, event: "UpdateSarcophagus", logs: logs, sub: sub}, nil
}

// WatchUpdateSarcophagus is a free log subscription operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) WatchUpdateSarcophagus(opts *bind.WatchOpts, sink chan<- *EventsUpdateSarcophagus, assetDoubleHash [][32]byte) (event.Subscription, error) {

	var assetDoubleHashRule []interface{}
	for _, assetDoubleHashItem := range assetDoubleHash {
		assetDoubleHashRule = append(assetDoubleHashRule, assetDoubleHashItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "UpdateSarcophagus", assetDoubleHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsUpdateSarcophagus)
				if err := _Events.contract.UnpackLog(event, "UpdateSarcophagus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateSarcophagus is a log parse operation binding the contract event 0x2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212.
//
// Solidity: event UpdateSarcophagus(bytes32 indexed assetDoubleHash, string assetId)
func (_Events *EventsFilterer) ParseUpdateSarcophagus(log types.Log) (*EventsUpdateSarcophagus, error) {
	event := new(EventsUpdateSarcophagus)
	if err := _Events.contract.UnpackLog(event, "UpdateSarcophagus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EventsWithdrawalFreeBondIterator is returned from FilterWithdrawalFreeBond and is used to iterate over the raw logs and unpacked data for WithdrawalFreeBond events raised by the Events contract.
type EventsWithdrawalFreeBondIterator struct {
	Event *EventsWithdrawalFreeBond // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsWithdrawalFreeBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsWithdrawalFreeBond)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsWithdrawalFreeBond)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsWithdrawalFreeBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsWithdrawalFreeBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsWithdrawalFreeBond represents a WithdrawalFreeBond event raised by the Events contract.
type EventsWithdrawalFreeBond struct {
	Archaeologist common.Address
	WithdrawnBond *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFreeBond is a free log retrieval operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) FilterWithdrawalFreeBond(opts *bind.FilterOpts, archaeologist []common.Address) (*EventsWithdrawalFreeBondIterator, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.FilterLogs(opts, "WithdrawalFreeBond", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return &EventsWithdrawalFreeBondIterator{contract: _Events.contract, event: "WithdrawalFreeBond", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFreeBond is a free log subscription operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) WatchWithdrawalFreeBond(opts *bind.WatchOpts, sink chan<- *EventsWithdrawalFreeBond, archaeologist []common.Address) (event.Subscription, error) {

	var archaeologistRule []interface{}
	for _, archaeologistItem := range archaeologist {
		archaeologistRule = append(archaeologistRule, archaeologistItem)
	}

	logs, sub, err := _Events.contract.WatchLogs(opts, "WithdrawalFreeBond", archaeologistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsWithdrawalFreeBond)
				if err := _Events.contract.UnpackLog(event, "WithdrawalFreeBond", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalFreeBond is a log parse operation binding the contract event 0x8ae00c9807f956157140b8b0d7839708e14a60ded9714c6d2760d6ad492ac78f.
//
// Solidity: event WithdrawalFreeBond(address indexed archaeologist, uint256 withdrawnBond)
func (_Events *EventsFilterer) ParseWithdrawalFreeBond(log types.Log) (*EventsWithdrawalFreeBond, error) {
	event := new(EventsWithdrawalFreeBond)
	if err := _Events.contract.UnpackLog(event, "WithdrawalFreeBond", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PrivateKeysABI is the input ABI used to generate the binding from.
const PrivateKeysABI = "[{\"inputs\":[],\"name\":\"gx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"privKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"keyVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PrivateKeysBin is the compiled bytecode used for deploying new contracts.
var PrivateKeysBin = "0x61035c610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c806306c91ce31461005057806310789d861461006e5780632d58c9a214610149575b600080fd5b610058610167565b6040518082815260200191505060405180910390f35b6101316004803603604081101561008457600080fd5b8101908080359060200190929190803590602001906401000000008111156100ab57600080fd5b8201836020820111156100bd57600080fd5b803590602001918460018302840111640100000000831117156100df57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061018b565b60405180821515815260200191505060405180910390f35b6101516101e1565b6040518082815260200191505060405180910390f35b7f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f8179881565b60006101d97f79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f817987f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b88585610205565b905092915050565b7f483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b881565b6000807ffffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03641419050600060016000806002898161023c57fe5b06141561024a57601b61024d565b601c5b8960001b858061025957fe5b8b8a60001c0960001b60405160008152602001604052604051808560001b81526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156102b9573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff858051906020012060001c1690508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614935050505094935050505056fea264697066735822122021b284587e808732059c386ca11cd1114733a942bff0af22182531fb0ee78d1a64736f6c63430007020033"

// DeployPrivateKeys deploys a new Ethereum contract, binding an instance of PrivateKeys to it.
func DeployPrivateKeys(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrivateKeys, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateKeysABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateKeysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivateKeys{PrivateKeysCaller: PrivateKeysCaller{contract: contract}, PrivateKeysTransactor: PrivateKeysTransactor{contract: contract}, PrivateKeysFilterer: PrivateKeysFilterer{contract: contract}}, nil
}

// PrivateKeys is an auto generated Go binding around an Ethereum contract.
type PrivateKeys struct {
	PrivateKeysCaller     // Read-only binding to the contract
	PrivateKeysTransactor // Write-only binding to the contract
	PrivateKeysFilterer   // Log filterer for contract events
}

// PrivateKeysCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateKeysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateKeysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateKeysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateKeysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateKeysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateKeysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateKeysSession struct {
	Contract     *PrivateKeys      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivateKeysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateKeysCallerSession struct {
	Contract *PrivateKeysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PrivateKeysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateKeysTransactorSession struct {
	Contract     *PrivateKeysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrivateKeysRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateKeysRaw struct {
	Contract *PrivateKeys // Generic contract binding to access the raw methods on
}

// PrivateKeysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateKeysCallerRaw struct {
	Contract *PrivateKeysCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateKeysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateKeysTransactorRaw struct {
	Contract *PrivateKeysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateKeys creates a new instance of PrivateKeys, bound to a specific deployed contract.
func NewPrivateKeys(address common.Address, backend bind.ContractBackend) (*PrivateKeys, error) {
	contract, err := bindPrivateKeys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateKeys{PrivateKeysCaller: PrivateKeysCaller{contract: contract}, PrivateKeysTransactor: PrivateKeysTransactor{contract: contract}, PrivateKeysFilterer: PrivateKeysFilterer{contract: contract}}, nil
}

// NewPrivateKeysCaller creates a new read-only instance of PrivateKeys, bound to a specific deployed contract.
func NewPrivateKeysCaller(address common.Address, caller bind.ContractCaller) (*PrivateKeysCaller, error) {
	contract, err := bindPrivateKeys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateKeysCaller{contract: contract}, nil
}

// NewPrivateKeysTransactor creates a new write-only instance of PrivateKeys, bound to a specific deployed contract.
func NewPrivateKeysTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateKeysTransactor, error) {
	contract, err := bindPrivateKeys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateKeysTransactor{contract: contract}, nil
}

// NewPrivateKeysFilterer creates a new log filterer instance of PrivateKeys, bound to a specific deployed contract.
func NewPrivateKeysFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateKeysFilterer, error) {
	contract, err := bindPrivateKeys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateKeysFilterer{contract: contract}, nil
}

// bindPrivateKeys binds a generic wrapper to an already deployed contract.
func bindPrivateKeys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateKeysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateKeys *PrivateKeysRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateKeys.Contract.PrivateKeysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateKeys *PrivateKeysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateKeys.Contract.PrivateKeysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateKeys *PrivateKeysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateKeys.Contract.PrivateKeysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateKeys *PrivateKeysCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivateKeys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateKeys *PrivateKeysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateKeys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateKeys *PrivateKeysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateKeys.Contract.contract.Transact(opts, method, params...)
}

// Gx is a free data retrieval call binding the contract method 0x06c91ce3.
//
// Solidity: function gx() view returns(uint256)
func (_PrivateKeys *PrivateKeysCaller) Gx(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateKeys.contract.Call(opts, out, "gx")
	return *ret0, err
}

// Gx is a free data retrieval call binding the contract method 0x06c91ce3.
//
// Solidity: function gx() view returns(uint256)
func (_PrivateKeys *PrivateKeysSession) Gx() (*big.Int, error) {
	return _PrivateKeys.Contract.Gx(&_PrivateKeys.CallOpts)
}

// Gx is a free data retrieval call binding the contract method 0x06c91ce3.
//
// Solidity: function gx() view returns(uint256)
func (_PrivateKeys *PrivateKeysCallerSession) Gx() (*big.Int, error) {
	return _PrivateKeys.Contract.Gx(&_PrivateKeys.CallOpts)
}

// Gy is a free data retrieval call binding the contract method 0x2d58c9a2.
//
// Solidity: function gy() view returns(uint256)
func (_PrivateKeys *PrivateKeysCaller) Gy(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivateKeys.contract.Call(opts, out, "gy")
	return *ret0, err
}

// Gy is a free data retrieval call binding the contract method 0x2d58c9a2.
//
// Solidity: function gy() view returns(uint256)
func (_PrivateKeys *PrivateKeysSession) Gy() (*big.Int, error) {
	return _PrivateKeys.Contract.Gy(&_PrivateKeys.CallOpts)
}

// Gy is a free data retrieval call binding the contract method 0x2d58c9a2.
//
// Solidity: function gy() view returns(uint256)
func (_PrivateKeys *PrivateKeysCallerSession) Gy() (*big.Int, error) {
	return _PrivateKeys.Contract.Gy(&_PrivateKeys.CallOpts)
}

// KeyVerification is a free data retrieval call binding the contract method 0x10789d86.
//
// Solidity: function keyVerification(bytes32 privKey, bytes pubKey) pure returns(bool)
func (_PrivateKeys *PrivateKeysCaller) KeyVerification(opts *bind.CallOpts, privKey [32]byte, pubKey []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivateKeys.contract.Call(opts, out, "keyVerification", privKey, pubKey)
	return *ret0, err
}

// KeyVerification is a free data retrieval call binding the contract method 0x10789d86.
//
// Solidity: function keyVerification(bytes32 privKey, bytes pubKey) pure returns(bool)
func (_PrivateKeys *PrivateKeysSession) KeyVerification(privKey [32]byte, pubKey []byte) (bool, error) {
	return _PrivateKeys.Contract.KeyVerification(&_PrivateKeys.CallOpts, privKey, pubKey)
}

// KeyVerification is a free data retrieval call binding the contract method 0x10789d86.
//
// Solidity: function keyVerification(bytes32 privKey, bytes pubKey) pure returns(bool)
func (_PrivateKeys *PrivateKeysCallerSession) KeyVerification(privKey [32]byte, pubKey []byte) (bool, error) {
	return _PrivateKeys.Contract.KeyVerification(&_PrivateKeys.CallOpts, privKey, pubKey)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c125d2353422a2540360ddf1900a18d4fbcf856799ab19582159927bd24ea4ab64736f6c63430007020033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SarcophagusABI is the input ABI used to generate the binding from.
const SarcophagusABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sarcoToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"accuseArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"archaeologistAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"archaeologistCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addy\",\"type\":\"address\"}],\"name\":\"archaeologists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cursedBond\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"burySarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"}],\"name\":\"cancelSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"cleanUpSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"recipientPublicKey\",\"type\":\"bytes\"}],\"name\":\"createSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"registerArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"rewrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcoToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"doubleHash\",\"type\":\"bytes32\"}],\"name\":\"sarcophagus\",\"outputs\":[{\"internalType\":\"enumTypes.SarcophagusStates\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"archaeologist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"storageFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sarcophagusCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"sarcophagusDoubleHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"privateKey\",\"type\":\"bytes32\"}],\"name\":\"unwrapSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"currentPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"paymentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeBond\",\"type\":\"uint256\"}],\"name\":\"updateArchaeologist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"assetDoubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"updateSarcophagus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBond\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SarcophagusBin is the compiled bytecode used for deploying new contracts.
var SarcophagusBin = "0x608060405234801561001057600080fd5b506040516128593803806128598339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2ae68230dd9c1c06842490953a33785a09292650f8c76b1a20b2fbbb214a896c81604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150612778806100e16000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638917aa28116100a2578063bf1a9d7711610071578063bf1a9d7714610a7e578063bfb62a7814610ae0578063c3daab9614610c73578063f456f36714610cb7578063f8fc6f7b14610db257610116565b80638917aa28146106e6578063972770ac14610728578063a2b1fb3c1461087f578063b8c77c38146108c357610116565b8063607e8f09116100e9578063607e8f091461028057806366086784146102b45780637d99cda11461030c5780637eb76a041461037057806386d05d581461052b57610116565b80630511f2ec1461011b5780631a7ac5501461015f57806349a5624f1461017d578063555912b714610262575b600080fd5b6101476004803603602081101561013157600080fd5b8101908080359060200190929190505050610f48565b60405180821515815260200191505060405180910390f35b610167611022565b6040518082815260200191505060405180910390f35b61024a6004803603606081101561019357600080fd5b8101908080359060200190929190803590602001906401000000008111156101ba57600080fd5b8201836020820111156101cc57600080fd5b803590602001918460018302840111640100000000831117156101ee57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611031565b60405180821515815260200191505060405180910390f35b61026a611182565b6040518082815260200191505060405180910390f35b610288611192565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102e0600480360360208110156102ca57600080fd5b81019080803590602001909291905050506111b6565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103586004803603604081101561032257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111f9565b60405180821515815260200191505060405180910390f35b610513600480360361010081101561038757600080fd5b81019080803590602001906401000000008111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111640100000000831117156103d857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561048d57600080fd5b82018360208201111561049f57600080fd5b803590602001918460018302840111640100000000831117156104c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506112f2565b60405180821515815260200191505060405180910390f35b6106ce600480360361010081101561054257600080fd5b810190808035906020019064010000000081111561055f57600080fd5b82018360208201111561057157600080fd5b8035906020019184600183028401116401000000008311171561059357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156105f657600080fd5b82018360208201111561060857600080fd5b8035906020019184600183028401116401000000008311171561062a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291905050506114eb565b60405180821515815260200191505060405180910390f35b610712600480360360208110156106fc57600080fd5b81019080803590602001909291905050506116e4565b6040518082815260200191505060405180910390f35b6107546004803603602081101561073e57600080fd5b8101908080359060200190929190505050611708565b6040518088600281111561076457fe5b81526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156107d75780820151818401526020810190506107bc565b50505050905090810190601f1680156108045780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561083d578082015181840152602081019050610822565b50505050905090810190601f16801561086a5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b6108ab6004803603602081101561089557600080fd5b8101908080359060200190929190505050611b24565b60405180821515815260200191505060405180910390f35b610a6660048036036101008110156108da57600080fd5b81019080803590602001906401000000008111156108f757600080fd5b82018360208201111561090957600080fd5b8035906020019184600183028401116401000000008311171561092b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561098e57600080fd5b8201836020820111156109a057600080fd5b803590602001918460018302840111640100000000831117156109c257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611bfe565b60405180821515815260200191505060405180910390f35b610ac860048036036080811015610a9457600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611df7565b60405180821515815260200191505060405180910390f35b610c5b600480360360c0811015610af657600080fd5b8101908080359060200190640100000000811115610b1357600080fd5b820183602082011115610b2557600080fd5b80359060200191846001830284011164010000000083111715610b4757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919080359060200190640100000000811115610bb457600080fd5b820183602082011115610bc657600080fd5b80359060200191846001830284011164010000000083111715610be857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff1690602001909291908035906020019092919080359060200190929190505050611eec565b60405180821515815260200191505060405180910390f35b610c9f60048036036020811015610c8957600080fd5b81019080803590602001909291905050506120c0565b60405180821515815260200191505060405180910390f35b610d9a60048036036060811015610ccd57600080fd5b810190808035906020019092919080359060200190640100000000811115610cf457600080fd5b820183602082011115610d0657600080fd5b80359060200191846001830284011164010000000083111715610d2857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061219a565b60405180821515815260200191505060405180910390f35b610df460048036036020811015610dc857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612301565b604051808c151581526020018b73ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018a73ffffffffffffffffffffffffffffffffffffffff16815260200189815260200188815260200187815260200186815260200185815260200184815260200183810383528c818151815260200191508051906020019080838360005b83811015610e9c578082015181840152602081019050610e81565b50505050905090810190601f168015610ec95780820380516001836020036101000a031916815260200191505b5083810382528b818151815260200191508051906020019080838360005b83811015610f02578082015181840152602081019050610ee7565b50505050905090810190601f168015610f2f5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405180910390f35b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63efd4eb5b60018460008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b8152600401808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001935050505060206040518083038186803b158015610fe057600080fd5b505af4158015610ff4573d6000803e3d6000fd5b505050506040513d602081101561100a57600080fd5b81019080805190602001909291905050509050919050565b60006001800180549050905090565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63f0ed8cf9600186868660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040180868152602001858152602001806020018481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b838110156110f05780820151818401526020810190506110d5565b50505050905090810190601f16801561111d5780820380516001836020036101000a031916815260200191505b50965050505050505060206040518083038186803b15801561113e57600080fd5b505af4158015611152573d6000803e3d6000fd5b505050506040513d602081101561116857600080fd5b810190808051906020019092919050505090509392505050565b6000600160080180549050905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001800182815481106111c757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63eba4cd776001858560008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518563ffffffff1660e01b8152600401808581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060206040518083038186803b1580156112af57600080fd5b505af41580156112c3573d6000803e3d6000fd5b505050506040513d60208110156112d957600080fd5b8101908080519060200190929190505050905092915050565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__6355c22e8960018b8b8b8b8b8b8b8b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518b63ffffffff1660e01b8152600401808b8152602001806020018a73ffffffffffffffffffffffffffffffffffffffff168152602001898152602001888152602001878152602001868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528c818151815260200191508051906020019080838360005b838110156113e85780820151818401526020810190506113cd565b50505050905090810190601f1680156114155780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561144e578082015181840152602081019050611433565b50505050905090810190601f16801561147b5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060206040518083038186803b1580156114a257600080fd5b505af41580156114b6573d6000803e3d6000fd5b505050506040513d60208110156114cc57600080fd5b8101908080519060200190929190505050905098975050505050505050565b600073__$aa38728129cff9f68d145c7eaa3d916b3c$__635d83aa6760018b8b8b8b8b8b8b8b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518b63ffffffff1660e01b8152600401808b815260200180602001806020018a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528c818151815260200191508051906020019080838360005b838110156115e15780820151818401526020810190506115c6565b50505050905090810190601f16801561160e5780820380516001836020036101000a031916815260200191505b5083810382528b818151815260200191508051906020019080838360005b8381101561164757808201518184015260208101905061162c565b50505050905090810190601f1680156116745780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060206040518083038186803b15801561169b57600080fd5b505af41580156116af573d6000803e3d6000fd5b505050506040513d60208110156116c557600080fd5b8101908080519060200190929190505050905098975050505050505050565b6000600160080182815481106116f657fe5b90600052602060002001549050919050565b600080600080606080600061171b612611565b600160070160008a8152602001908152602001600020604051806101c00160405290816000820160009054906101000a900460ff16600281111561175b57fe5b600281111561176657fe5b81526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118595780601f1061182e57610100808354040283529160200191611859565b820191906000526020600020905b81548152906001019060200180831161183c57829003601f168201915b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156119515780601f1061192657610100808354040283529160200191611951565b820191906000526020600020905b81548152906001019060200180831161193457829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611a075780601f106119dc57610100808354040283529160200191611a07565b820191906000526020600020905b8154815290600101906020018083116119ea57829003601f168201915b50505050508152602001600782018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611aa95780601f10611a7e57610100808354040283529160200191611aa9565b820191906000526020600020905b815481529060010190602001808311611a8c57829003601f168201915b505050505081526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c820154815250509050806000015181602001518260a001518360c0015184608001518560e00151866101200151975097509750975097509750975050919395979092949650565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63a6ea2a3d60018460008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b8152600401808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001935050505060206040518083038186803b158015611bbc57600080fd5b505af4158015611bd0573d6000803e3d6000fd5b505050506040513d6020811015611be657600080fd5b81019080805190602001909291905050509050919050565b600073__$aa38728129cff9f68d145c7eaa3d916b3c$__63ca03a0c260018a8c8b8b8b8b8b8b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518b63ffffffff1660e01b8152600401808b815260200180602001806020018a73ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018881526020018781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528c818151815260200191508051906020019080838360005b83811015611cf4578082015181840152602081019050611cd9565b50505050905090810190601f168015611d215780820380516001836020036101000a031916815260200191505b5083810382528b818151815260200191508051906020019080838360005b83811015611d5a578082015181840152602081019050611d3f565b50505050905090810190601f168015611d875780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060206040518083038186803b158015611dae57600080fd5b505af4158015611dc2573d6000803e3d6000fd5b505050506040513d6020811015611dd857600080fd5b8101908080519060200190929190505050905098975050505050505050565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63544730ba60018787878760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518763ffffffff1660e01b8152600401808781526020018681526020018581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001965050505050505060206040518083038186803b158015611ea757600080fd5b505af4158015611ebb573d6000803e3d6000fd5b505050506040513d6020811015611ed157600080fd5b81019080805190602001909291905050509050949350505050565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__63fc0742cc600189898989898960008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518963ffffffff1660e01b81526004018089815260200180602001888152602001806020018760ff1681526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183810383528a818151815260200191508051906020019080838360005b83811015611fc1578082015181840152602081019050611fa6565b50505050905090810190601f168015611fee5780820380516001836020036101000a031916815260200191505b50838103825288818151815260200191508051906020019080838360005b8381101561202757808201518184015260208101905061200c565b50505050905090810190601f1680156120545780820380516001836020036101000a031916815260200191505b509a505050505050505050505060206040518083038186803b15801561207957600080fd5b505af415801561208d573d6000803e3d6000fd5b505050506040513d60208110156120a357600080fd5b810190808051906020019092919050505090509695505050505050565b600073__$aa38728129cff9f68d145c7eaa3d916b3c$__630460a24b60018460008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518463ffffffff1660e01b8152600401808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff168152602001935050505060206040518083038186803b15801561215857600080fd5b505af415801561216c573d6000803e3d6000fd5b505050506040513d602081101561218257600080fd5b81019080805190602001909291905050509050919050565b600073__$e91813ee0e1b88ba13a7b292e30aaf308e$__6340ebc6a9600186868660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040180868152602001858152602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b8381101561226f578082015181840152602081019050612254565b50505050905090810190601f16801561229c5780820380516001836020036101000a031916815260200191505b50965050505050505060206040518083038186803b1580156122bd57600080fd5b505af41580156122d1573d6000803e3d6000fd5b505050506040513d60208110156122e757600080fd5b810190808051906020019092919050505090509392505050565b600080606080600080600080600080600061231a6126ba565b600160000160008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806101600160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156124725780601f1061244757610100808354040283529160200191612472565b820191906000526020600020905b81548152906001019060200180831161245557829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156125145780601f106124e957610100808354040283529160200191612514565b820191906000526020600020905b8154815290600101906020018083116124f757829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815260200160078201548152602001600882015481526020016009820154815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e001518861010001518961012001518a61014001519b509b509b509b509b509b509b509b509b509b509b505091939597999b90929496989a50565b604051806101c001604052806000600281111561262a57fe5b8152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016000815260200160008152602001606081526020016060815260200160008152602001600081526020016000815260200160008152602001600080191681525090565b604051806101600160405280600015158152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600081526020016000815260200160008152509056fea2646970667358221220f962c67aa74a108ca5915c581dfe478dd70a40cc3f7ec8781f594e2bd08c494b64736f6c63430007020033"

// DeploySarcophagus deploys a new Ethereum contract, binding an instance of Sarcophagus to it.
func DeploySarcophagus(auth *bind.TransactOpts, backend bind.ContractBackend, _sarcoToken common.Address) (common.Address, *types.Transaction, *Sarcophagus, error) {
	parsed, err := abi.JSON(strings.NewReader(SarcophagusABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	archaeologistsAddr, _, _, _ := DeployArchaeologists(auth, backend)
	SarcophagusBin = strings.Replace(SarcophagusBin, "__$aa38728129cff9f68d145c7eaa3d916b3c$__", archaeologistsAddr.String()[2:], -1)

	sarcophagusesAddr, _, _, _ := DeploySarcophaguses(auth, backend)
	SarcophagusBin = strings.Replace(SarcophagusBin, "__$e91813ee0e1b88ba13a7b292e30aaf308e$__", sarcophagusesAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SarcophagusBin), backend, _sarcoToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sarcophagus{SarcophagusCaller: SarcophagusCaller{contract: contract}, SarcophagusTransactor: SarcophagusTransactor{contract: contract}, SarcophagusFilterer: SarcophagusFilterer{contract: contract}}, nil
}

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
// Solidity: function archaeologists(address addy) view returns(bool exists, address archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusCaller) Archaeologists(opts *bind.CallOpts, addy common.Address) (struct {
	Exists                  bool
	Archaeologist           common.Address
	CurrentPublicKey        []byte
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
		Archaeologist           common.Address
		CurrentPublicKey        []byte
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
// Solidity: function archaeologists(address addy) view returns(bool exists, address archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusSession) Archaeologists(addy common.Address) (struct {
	Exists                  bool
	Archaeologist           common.Address
	CurrentPublicKey        []byte
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
// Solidity: function archaeologists(address addy) view returns(bool exists, address archaeologist, bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond, uint256 cursedBond)
func (_Sarcophagus *SarcophagusCallerSession) Archaeologists(addy common.Address) (struct {
	Exists                  bool
	Archaeologist           common.Address
	CurrentPublicKey        []byte
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

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 doubleHash) view returns(uint8 state, address archaeologist, uint256 resurrectionTime, uint256 resurrectionWindow, string name, string assetId, uint256 storageFee)
func (_Sarcophagus *SarcophagusCaller) Sarcophagus(opts *bind.CallOpts, doubleHash [32]byte) (struct {
	State              uint8
	Archaeologist      common.Address
	ResurrectionTime   *big.Int
	ResurrectionWindow *big.Int
	Name               string
	AssetId            string
	StorageFee         *big.Int
}, error) {
	ret := new(struct {
		State              uint8
		Archaeologist      common.Address
		ResurrectionTime   *big.Int
		ResurrectionWindow *big.Int
		Name               string
		AssetId            string
		StorageFee         *big.Int
	})
	out := ret
	err := _Sarcophagus.contract.Call(opts, out, "sarcophagus", doubleHash)
	return *ret, err
}

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 doubleHash) view returns(uint8 state, address archaeologist, uint256 resurrectionTime, uint256 resurrectionWindow, string name, string assetId, uint256 storageFee)
func (_Sarcophagus *SarcophagusSession) Sarcophagus(doubleHash [32]byte) (struct {
	State              uint8
	Archaeologist      common.Address
	ResurrectionTime   *big.Int
	ResurrectionWindow *big.Int
	Name               string
	AssetId            string
	StorageFee         *big.Int
}, error) {
	return _Sarcophagus.Contract.Sarcophagus(&_Sarcophagus.CallOpts, doubleHash)
}

// Sarcophagus is a free data retrieval call binding the contract method 0x972770ac.
//
// Solidity: function sarcophagus(bytes32 doubleHash) view returns(uint8 state, address archaeologist, uint256 resurrectionTime, uint256 resurrectionWindow, string name, string assetId, uint256 storageFee)
func (_Sarcophagus *SarcophagusCallerSession) Sarcophagus(doubleHash [32]byte) (struct {
	State              uint8
	Archaeologist      common.Address
	ResurrectionTime   *big.Int
	ResurrectionWindow *big.Int
	Name               string
	AssetId            string
	StorageFee         *big.Int
}, error) {
	return _Sarcophagus.Contract.Sarcophagus(&_Sarcophagus.CallOpts, doubleHash)
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

// SarcophagusDoubleHash is a free data retrieval call binding the contract method 0x8917aa28.
//
// Solidity: function sarcophagusDoubleHash(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCaller) SarcophagusDoubleHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Sarcophagus.contract.Call(opts, out, "sarcophagusDoubleHash", index)
	return *ret0, err
}

// SarcophagusDoubleHash is a free data retrieval call binding the contract method 0x8917aa28.
//
// Solidity: function sarcophagusDoubleHash(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusSession) SarcophagusDoubleHash(index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.SarcophagusDoubleHash(&_Sarcophagus.CallOpts, index)
}

// SarcophagusDoubleHash is a free data retrieval call binding the contract method 0x8917aa28.
//
// Solidity: function sarcophagusDoubleHash(uint256 index) view returns(bytes32)
func (_Sarcophagus *SarcophagusCallerSession) SarcophagusDoubleHash(index *big.Int) ([32]byte, error) {
	return _Sarcophagus.Contract.SarcophagusDoubleHash(&_Sarcophagus.CallOpts, index)
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

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) CreateSarcophagus(opts *bind.TransactOpts, name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "createSarcophagus", name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusSession) CreateSarcophagus(name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// CreateSarcophagus is a paid mutator transaction binding the contract method 0x7eb76a04.
//
// Solidity: function createSarcophagus(string name, address archaeologist, uint256 resurrectionTime, uint256 storageFee, uint256 diggingFee, uint256 bounty, bytes32 assetDoubleHash, bytes recipientPublicKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) CreateSarcophagus(name string, archaeologist common.Address, resurrectionTime *big.Int, storageFee *big.Int, diggingFee *big.Int, bounty *big.Int, assetDoubleHash [32]byte, recipientPublicKey []byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.CreateSarcophagus(&_Sarcophagus.TransactOpts, name, archaeologist, resurrectionTime, storageFee, diggingFee, bounty, assetDoubleHash, recipientPublicKey)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) RegisterArchaeologist(opts *bind.TransactOpts, currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "registerArchaeologist", currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusSession) RegisterArchaeologist(currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// RegisterArchaeologist is a paid mutator transaction binding the contract method 0x86d05d58.
//
// Solidity: function registerArchaeologist(bytes currentPublicKey, string endpoint, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) RegisterArchaeologist(currentPublicKey []byte, endpoint string, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.RegisterArchaeologist(&_Sarcophagus.TransactOpts, currentPublicKey, endpoint, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
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

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x49a5624f.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UnwrapSarcophagus(opts *bind.TransactOpts, assetDoubleHash [32]byte, singleHash []byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "unwrapSarcophagus", assetDoubleHash, singleHash, privateKey)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x49a5624f.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusSession) UnwrapSarcophagus(assetDoubleHash [32]byte, singleHash []byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash, privateKey)
}

// UnwrapSarcophagus is a paid mutator transaction binding the contract method 0x49a5624f.
//
// Solidity: function unwrapSarcophagus(bytes32 assetDoubleHash, bytes singleHash, bytes32 privateKey) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UnwrapSarcophagus(assetDoubleHash [32]byte, singleHash []byte, privateKey [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UnwrapSarcophagus(&_Sarcophagus.TransactOpts, assetDoubleHash, singleHash, privateKey)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes currentPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateArchaeologist(opts *bind.TransactOpts, endpoint string, currentPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateArchaeologist", endpoint, currentPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes currentPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateArchaeologist(endpoint string, currentPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, currentPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateArchaeologist is a paid mutator transaction binding the contract method 0xb8c77c38.
//
// Solidity: function updateArchaeologist(string endpoint, bytes currentPublicKey, address paymentAddress, uint256 feePerByte, uint256 minimumBounty, uint256 minimumDiggingFee, uint256 maximumResurrectionTime, uint256 freeBond) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateArchaeologist(endpoint string, currentPublicKey []byte, paymentAddress common.Address, feePerByte *big.Int, minimumBounty *big.Int, minimumDiggingFee *big.Int, maximumResurrectionTime *big.Int, freeBond *big.Int) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateArchaeologist(&_Sarcophagus.TransactOpts, endpoint, currentPublicKey, paymentAddress, feePerByte, minimumBounty, minimumDiggingFee, maximumResurrectionTime, freeBond)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactor) UpdateSarcophagus(opts *bind.TransactOpts, newPublicKey []byte, assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.contract.Transact(opts, "updateSarcophagus", newPublicKey, assetDoubleHash, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusSession) UpdateSarcophagus(newPublicKey []byte, assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, newPublicKey, assetDoubleHash, assetId, v, r, s)
}

// UpdateSarcophagus is a paid mutator transaction binding the contract method 0xbfb62a78.
//
// Solidity: function updateSarcophagus(bytes newPublicKey, bytes32 assetDoubleHash, string assetId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Sarcophagus *SarcophagusTransactorSession) UpdateSarcophagus(newPublicKey []byte, assetDoubleHash [32]byte, assetId string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Sarcophagus.Contract.UpdateSarcophagus(&_Sarcophagus.TransactOpts, newPublicKey, assetDoubleHash, assetId, v, r, s)
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

// SarcophagusesABI is the input ABI used to generate the binding from.
const SarcophagusesABI = "[]"

// SarcophagusesBin is the compiled bytecode used for deploying new contracts.
var SarcophagusesBin = "0x61442d610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c8063eba4cd7711610065578063eba4cd77146104cf578063efd4eb5b1461056a578063f0ed8cf9146105e5578063fc0742cc1461070157610092565b806340ebc6a914610097578063544730ba146101c957806355c22e8914610262578063a6ea2a3d14610454575b600080fd5b8180156100a357600080fd5b506101b1600480360360a08110156100ba57600080fd5b810190808035906020019092919080359060200190929190803590602001906401000000008111156100eb57600080fd5b8201836020820111156100fd57600080fd5b8035906020019184600183028401116401000000008311171561011f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108cc565b60405180821515815260200191505060405180910390f35b8180156101d557600080fd5b5061024a600480360360c08110156101ec57600080fd5b810190808035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b6c565b60405180821515815260200191505060405180910390f35b81801561026e57600080fd5b5061043c600480360361014081101561028657600080fd5b8101908080359060200190929190803590602001906401000000008111156102ad57600080fd5b8201836020820111156102bf57600080fd5b803590602001918460018302840111640100000000831117156102e157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561039657600080fd5b8201836020820111156103a857600080fd5b803590602001918460018302840111640100000000831117156103ca57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061133b565b60405180821515815260200191505060405180910390f35b81801561046057600080fd5b506104b76004803603606081101561047757600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d57565b60405180821515815260200191505060405180910390f35b8180156104db57600080fd5b50610552600480360360808110156104f257600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612526565b60405180821515815260200191505060405180910390f35b81801561057657600080fd5b506105cd6004803603606081101561058d57600080fd5b810190808035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506126f8565b60405180821515815260200191505060405180910390f35b8180156105f157600080fd5b506106e9600480360360a081101561060857600080fd5b8101908080359060200190929190803590602001909291908035906020019064010000000081111561063957600080fd5b82018360208201111561064b57600080fd5b8035906020019184600183028401116401000000008311171561066d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612abf565b60405180821515815260200191505060405180910390f35b81801561070d57600080fd5b506108b4600480360361010081101561072557600080fd5b81019080803590602001909291908035906020019064010000000081111561074c57600080fd5b82018360208201111561075e57600080fd5b8035906020019184600183028401116401000000008311171561078057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190803590602001906401000000008111156107ed57600080fd5b8201836020820111156107ff57600080fd5b8035906020019184600183028401116401000000008311171561082157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506131d7565b60405180821515815260200191505060405180910390f35b60008086600701600087815260200190815260200160002090506109028160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__631807715f82600401546040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561095557600080fd5b505af4158015610969573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63a25bb36587876040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156109de5780820151818401526020810190506109c3565b50505050905090810190601f168015610a0b5780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b158015610a2957600080fd5b505af4158015610a3d573d6000803e3d6000fd5b50505050600080610a75898460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16888689613bce565b915091508860040160008460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208890806001815401808255809150506001900390600052602060002001600090919091909150553373ffffffffffffffffffffffffffffffffffffffff16887f6cee286923fc4de6af78ba8b06ff9f4236e5aad4282a5e13a3cea8446c803fd48484604051808381526020018281526020019250505060405180910390a36001935050505095945050505050565b6000808760070160008881526020019081526020016000209050610ba28160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__6377a5b2d98260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060006040518083038186803b158015610c2b57600080fd5b505af4158015610c3f573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__631807715f82600401546040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610c9657600080fd5b505af4158015610caa573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__631807715f876040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b158015610cfd57600080fd5b505af4158015610d11573d6000803e3d6000fd5b5050505060008860000160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63c029b7b48888888560070154866006015487600501546040518763ffffffff1660e01b815260040180878152602001868152602001858152602001848152602001838152602001828152602001965050505050505060006040518083038186803b158015610e0157600080fd5b505af4158015610e15573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330896040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505050506040513d6020811015610ed257600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600901546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610f7d57600080fd5b505af1158015610f91573d6000803e3d6000fd5b505050506040513d6020811015610fa757600080fd5b8101908080519060200190929190505050506000610fce8688613e7e90919063ffffffff16565b905082600b01548111156110ae576000610ff584600b015483613f0690919063ffffffff16565b905073__$aa38728129cff9f68d145c7eaa3d916b3c$__63e6b72c968c8660000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b15801561109057600080fd5b505af41580156110a4573d6000803e3d6000fd5b5050505050611189565b82600b01548110156111885760006110d38285600b0154613f0690919063ffffffff16565b905073__$aa38728129cff9f68d145c7eaa3d916b3c$__630d9616ca8c8660000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b15801561116e57600080fd5b505af4158015611182573d6000803e3d6000fd5b50505050505b5b600073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63c54820118a6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156111da57600080fd5b505af41580156111ee573d6000803e3d6000fd5b505050506040513d602081101561120457600080fd5b810190808051906020019092919050505090508884600401819055508784600901819055508684600a01819055508184600b0181905550808460050181905550897f17f5acaa1fbabbe2ec2373f9adb3a1b2f191bac493461e9f64b1db1e0a3f5e46856006018b848c8c8860405180806020018781526020018681526020018581526020018481526020018381526020018281038252888181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156113165780601f106112eb57610100808354040283529160200191611316565b820191906000526020600020905b8154815290600101906020018083116112f957829003601f168201915b505097505050505050505060405180910390a260019450505050509695505050505050565b600073__$aa38728129cff9f68d145c7eaa3d916b3c$__6377a91eac8c8b60016040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018215158152602001935050505060006040518083038186803b1580156113b557600080fd5b505af41580156113c9573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63495fe007846040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561143757808201518184015260208101905061141c565b50505050905090810190601f1680156114645780820380516001836020036101000a031916815260200191505b509250505060006040518083038186803b15801561148157600080fd5b505af4158015611495573d6000803e3d6000fd5b505050506114c98b600701600086815260200190815260200160002060000160009054906101000a900460ff166000613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__631807715f896040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561151857600080fd5b505af415801561152c573d6000803e3d6000fd5b5050505060008b60000160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63c029b7b48a89898560070154866006015487600501546040518763ffffffff1660e01b815260040180878152602001868152602001858152602001848152602001838152602001828152602001965050505050505060006040518083038186803b1580156115f857600080fd5b505af415801561160c573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff166323b872dd33306116538c6116458c8e613e7e90919063ffffffff16565b613e7e90919063ffffffff16565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156116c357600080fd5b505af11580156116d7573d6000803e3d6000fd5b505050506040513d60208110156116ed57600080fd5b81019080805190602001909291905050505060006117148789613e7e90919063ffffffff16565b905073__$aa38728129cff9f68d145c7eaa3d916b3c$__63e6b72c968e8d846040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b15801561178b57600080fd5b505af415801561179f573d6000803e3d6000fd5b505050506117ab61414b565b604051806101c00160405280600160028111156117c457fe5b81526020018d73ffffffffffffffffffffffffffffffffffffffff168152602001846001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561187d5780601f106118525761010080835404028352916020019161187d565b820191906000526020600020905b81548152906001019060200180831161186057829003601f168201915b505050505081526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018e81526020018c815260200173__$b4ae538d42a0e0bb8330684d08e51a40b4$__63c54820118e6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156118fe57600080fd5b505af4158015611912573d6000803e3d6000fd5b505050506040513d602081101561192857600080fd5b810190808051906020019092919050505081526020016040518060200160405280600081525081526020018781526020018b81526020018a81526020018981526020018381526020016000801b8152509050808e600701600089815260200190815260200160002060008201518160000160006101000a81548160ff021916908360028111156119b457fe5b021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816001019080519060200190611a1c9291906141f4565b5060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816003019080519060200190611a80929190614274565b5060a0820151816004015560c0820151816005015560e0820151816006019080519060200190611ab1929190614274565b50610100820151816007019080519060200190611acf9291906141f4565b506101208201518160080155610140820151816009015561016082015181600a015561018082015181600b01556101a082015181600c01559050508d600801879080600181540180825580915050600190039060005260206000200160009091909190915055806020015173ffffffffffffffffffffffffffffffffffffffff16877f61377d9cca1a0627203c24e7200ac291d185768dd96a7278cbc7a2db6d81f8508360400151846060015185608001518660a001518760c001518861012001518961014001518a61016001518b61010001518c610180015160405180806020018b73ffffffffffffffffffffffffffffffffffffffff168152602001806020018a81526020018981526020018881526020018781526020018681526020018060200185815260200184810384528e818151815260200191508051906020019080838360005b83811015611c31578082015181840152602081019050611c16565b50505050905090810190601f168015611c5e5780820380516001836020036101000a031916815260200191505b5084810383528c818151815260200191508051906020019080838360005b83811015611c97578082015181840152602081019050611c7c565b50505050905090810190601f168015611cc45780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b83811015611cfd578082015181840152602081019050611ce2565b50505050905090810190601f168015611d2a5780820380516001836020036101000a031916815260200191505b509d505050505050505050505050505060405180910390a3600193505050509a9950505050505050505050565b6000808460070160008581526020019081526020016000209050611d8d8160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__636d74499f826006016040518263ffffffff1660e01b81526004018080602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611e435780601f10611e1857610100808354040283529160200191611e43565b820191906000526020600020905b815481529060010190602001808311611e2657829003601f168201915b50509250505060006040518083038186803b158015611e6157600080fd5b505af4158015611e75573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__6377a5b2d98260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060006040518083038186803b158015611f0257600080fd5b505af4158015611f16573d6000803e3d6000fd5b50505050611f226142f4565b8560000160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806101600160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561209d5780601f106120725761010080835404028352916020019161209d565b820191906000526020600020905b81548152906001019060200180831161208057829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561213f5780601f106121145761010080835404028352916020019161213f565b820191906000526020600020905b81548152906001019060200180831161212257829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152505090508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612238856008015486600a0154613e7e90919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561228b57600080fd5b505af115801561229f573d6000803e3d6000fd5b505050506040513d60208110156122b557600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb826080015184600901546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561234057600080fd5b505af1158015612354573d6000803e3d6000fd5b505050506040513d602081101561236a57600080fd5b81019080805190602001909291905050505073__$aa38728129cff9f68d145c7eaa3d916b3c$__630d9616ca878460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600b01546040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b15801561241957600080fd5b505af415801561242d573d6000803e3d6000fd5b5050505061243a82613f50565b60028260000160006101000a81548160ff0219169083600281111561245b57fe5b02179055508560030160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020859080600181540180825580915050600190039060005260206000200160009091909190915055847f829e85bd7eda128d8f00c1b0b36d28a9357d6c88d69a552de0d3be7b766e775060405160405180910390a26001925050509392505050565b600080856007016000868152602001908152602001600020905061255c8160000160009054906101000a900460ff166001613a93565b4261257882600501548360040154613e7e90919063ffffffff16565b106125ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603381526020018061439a6033913960400191505060405180910390fd5b600080612602888460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16888689613bce565b915091508760050160008460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208790806001815401808255809150506001900390600052602060002001600090919091909150553373ffffffffffffffffffffffffffffffffffffffff16877f106e723534f3d39516c80d5e49eb5ac7788139c661267137c40382659d409fb98484604051808381526020018281526020019250505060405180910390a360019350505050949350505050565b600080846007016000858152602001908152602001600020905061272e8160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__6377a5b2d98260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060006040518083038186803b1580156127b757600080fd5b505af41580156127cb573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__631807715f82600401546040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561282257600080fd5b505af4158015612836573d6000803e3d6000fd5b5050505060008560000160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905073__$aa38728129cff9f68d145c7eaa3d916b3c$__630d9616ca878460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600b01546040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b15801561294057600080fd5b505af4158015612954573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600901546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156129f157600080fd5b505af1158015612a05573d6000803e3d6000fd5b505050506040513d6020811015612a1b57600080fd5b810190808051906020019092919050505050612a3682613f50565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826004018190555060028260000160006101000a81548160ff02191690836002811115612a8057fe5b0217905550847f3a0205aa93b1a96da7d6405b2967f91b441383623de2d1ee3b27bfd1642b167a60405160405180910390a26001925050509392505050565b6000808660070160008781526020019081526020016000209050612af58160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__632be5910a826004015483600501546040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b158015612b5457600080fd5b505af4158015612b68573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__63a25bb36587876040518363ffffffff1660e01b81526004018083815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612bdd578082015181840152602081019050612bc2565b50505050905090810190601f168015612c0a5780820380516001836020036101000a031916815260200191505b50935050505060006040518083038186803b158015612c2857600080fd5b505af4158015612c3c573d6000803e3d6000fd5b5050505073__$cc89d10cddaafd15e21a5cd49626695b96$__6310789d8685836001016040518363ffffffff1660e01b81526004018083815260200180602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015612cfd5780601f10612cd257610100808354040283529160200191612cfd565b820191906000526020600020905b815481529060010190602001808311612ce057829003601f168201915b5050935050505060206040518083038186803b158015612d1c57600080fd5b505af4158015612d30573d6000803e3d6000fd5b505050506040513d6020811015612d4657600080fd5b8101908080519060200190929190505050612dc9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f21707269766174654b657900000000000000000000000000000000000000000081525060200191505060405180910390fd5b8381600c018190555060008760000160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612e9785600a01548660090154613e7e90919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015612eea57600080fd5b505af1158015612efe573d6000803e3d6000fd5b505050506040513d6020811015612f1457600080fd5b81019080805190602001909291905050505073__$aa38728129cff9f68d145c7eaa3d916b3c$__630d9616ca898460000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600b01546040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b158015612fc357600080fd5b505af4158015612fd7573d6000803e3d6000fd5b50505050612fe482613f50565b60028260000160006101000a81548160ff0219169083600281111561300557fe5b02179055508760020160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020879080600181540180825580915050600190039060005260206000200160009091909190915055867f8d31f85522fabfbd8bdc64e740a84ac72cd810e0bb788f4454c17f5fc392c7888360060188886040518080602001806020018481526020018381038352868181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156131515780601f1061312657610100808354040283529160200191613151565b820191906000526020600020905b81548152906001019060200180831161313457829003601f168201915b5050838103825285818151815260200191508051906020019080838360005b8381101561318b578082015181840152602081019050613170565b50505050905090810190601f1680156131b85780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a260019250505095945050505050565b600080896007016000898152602001908152602001600020905061320d8160000160009054906101000a900460ff166001613a93565b73__$b4ae538d42a0e0bb8330684d08e51a40b4$__6377a5b2d98260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060006040518083038186803b15801561329657600080fd5b505af41580156132aa573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__638899ecf782600601896040518363ffffffff1660e01b81526004018080602001806020018381038352858181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156133695780601f1061333e57610100808354040283529160200191613369565b820191906000526020600020905b81548152906001019060200180831161334c57829003601f168201915b5050838103825284818151815260200191508051906020019080838360005b838110156133a3578082015181840152602081019050613388565b50505050905090810190601f1680156133d05780820380516001836020036101000a031916815260200191505b5094505050505060006040518083038186803b1580156133ef57600080fd5b505af4158015613403573d6000803e3d6000fd5b5050505073__$b4ae538d42a0e0bb8330684d08e51a40b4$__638ad5f4468a896040516020018083805190602001908083835b602083106134595780518252602082019150602081019050602083039250613436565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083106134aa5780518252602082019150602081019050602083039250613487565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528888888660000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518663ffffffff1660e01b815260040180806020018660ff1681526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828103825287818151815260200191508051906020019080838360005b8381101561358757808201518184015260208101905061356c565b50505050905090810190601f1680156135b45780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b1580156135d557600080fd5b505af41580156135e9573d6000803e3d6000fd5b50505050896006018160010160405180828054600181600116156101000203166002900480156136505780601f1061362e576101008083540402835291820191613650565b820191906000526020600020905b81548152906001019060200180831161363c575b5050915050908152602001604051809103902060009054906101000a900460ff16156136e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f7075626c6963206b657920616c7265616479207573656400000000000000000081525060200191505060405180910390fd5b60018a6006018260010160405180828054600181600116156101000203166002900480156137495780601f10613727576101008083540402835291820191613749565b820191906000526020600020905b815481529060010190602001808311613735575b5050915050908152602001604051809103902060006101000a81548160ff0219169083151502179055508681600601908051906020019061378b929190614274565b5060008a60000160008360000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508981600101908051906020019061380d9291906141f4565b508373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600801546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156138a757600080fd5b505af11580156138bb573d6000803e3d6000fd5b505050506040513d60208110156138d157600080fd5b81019080805190602001909291905050505060008260080181905550887f2a08f73a04c4a1376d4405159d2bdeb189df1ae580eb3ec566844dc3330d1212896040518080602001828103825283818151815260200191508051906020019080838360005b83811015613950578082015181840152602081019050613935565b50505050905090810190601f16801561397d5780820380516001836020036101000a031916815260200191505b509250505060405180910390a28060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f6f8affc6f8d70bdb860a7009d79090e7b415131ce3c1030a6a7aaa7391586a17826001016040518080602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015613a735780601f10613a4857610100808354040283529160200191613a73565b820191906000526020600020905b815481529060010190602001808311613a5657829003601f168201915b50509250505060405180910390a260019250505098975050505050505050565b60606040518060400160405280601a81526020017f736172636f70686167757320616c726561647920657869737473000000000000815250905060016002811115613ada57fe5b826002811115613ae657fe5b1415613b08576040518060600160405280602b81526020016143cd602b913990505b816002811115613b1457fe5b836002811115613b2057fe5b148190613bc8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613b8d578082015181840152602081019050613b72565b50505050905090810190601f168015613bba5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505050565b6000806000613beb600286600b0154613f7b90919063ffffffff16565b90506000613c068287600b0154613f0690919063ffffffff16565b90508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8760020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16613c7685613c688b600901548c600a0154613e7e90919063ffffffff16565b613e7e90919063ffffffff16565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613cc957600080fd5b505af1158015613cdd573d6000803e3d6000fd5b505050506040513d6020811015613cf357600080fd5b8101908080519060200190929190505050508473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb88836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015613d7657600080fd5b505af1158015613d8a573d6000803e3d6000fd5b505050506040513d6020811015613da057600080fd5b81019080805190602001909291905050505073__$aa38728129cff9f68d145c7eaa3d916b3c$__6310dfb3758a8a89600b01546040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060006040518083038186803b158015613e2b57600080fd5b505af4158015613e3f573d6000803e3d6000fd5b50505050613e4c86613f50565b60028660000160006101000a81548160ff02191690836002811115613e6d57fe5b021790555050509550959350505050565b600080828401905083811015613efc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000613f4883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613fc5565b905092915050565b6000816008018190555060008160090181905550600081600a0181905550600081600b018190555050565b6000613fbd83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250614085565b905092915050565b6000838311158290614072576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561403757808201518184015260208101905061401c565b50505050905090810190601f1680156140645780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008083118290614131576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156140f65780820151818401526020810190506140db565b50505050905090810190601f1680156141235780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161413d57fe5b049050809150509392505050565b604051806101c001604052806000600281111561416457fe5b8152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016000815260200160008152602001606081526020016060815260200160008152602001600081526020016000815260200160008152602001600080191681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061423557805160ff1916838001178555614263565b82800160010185558215614263579182015b82811115614262578251825591602001919060010190614247565b5b509050614270919061437c565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106142b557805160ff19168380011785556142e3565b828001600101855582156142e3579182015b828111156142e25782518255916020019190600101906142c7565b5b5090506142f0919061437c565b5090565b604051806101600160405280600015158152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b5b8082111561439557600081600090555060010161437d565b509056fe736172636f70686167757320726573757272656374696f6e20706572696f64206d75737420626520696e207468652070617374736172636f70686167757320646f6573206e6f74206578697374206f72206973206e6f7420616374697665a2646970667358221220312dc8830d21a2664447ea339ca5301b0d054ca42904486a82fe7c6943c340e964736f6c63430007020033"

// DeploySarcophaguses deploys a new Ethereum contract, binding an instance of Sarcophaguses to it.
func DeploySarcophaguses(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sarcophaguses, error) {
	parsed, err := abi.JSON(strings.NewReader(SarcophagusesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	archaeologistsAddr, _, _, _ := DeployArchaeologists(auth, backend)
	SarcophagusesBin = strings.Replace(SarcophagusesBin, "__$aa38728129cff9f68d145c7eaa3d916b3c$__", archaeologistsAddr.String()[2:], -1)

	utilsAddr, _, _, _ := DeployUtils(auth, backend)
	SarcophagusesBin = strings.Replace(SarcophagusesBin, "__$b4ae538d42a0e0bb8330684d08e51a40b4$__", utilsAddr.String()[2:], -1)

	privateKeysAddr, _, _, _ := DeployPrivateKeys(auth, backend)
	SarcophagusesBin = strings.Replace(SarcophagusesBin, "__$cc89d10cddaafd15e21a5cd49626695b96$__", privateKeysAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SarcophagusesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sarcophaguses{SarcophagusesCaller: SarcophagusesCaller{contract: contract}, SarcophagusesTransactor: SarcophagusesTransactor{contract: contract}, SarcophagusesFilterer: SarcophagusesFilterer{contract: contract}}, nil
}

// Sarcophaguses is an auto generated Go binding around an Ethereum contract.
type Sarcophaguses struct {
	SarcophagusesCaller     // Read-only binding to the contract
	SarcophagusesTransactor // Write-only binding to the contract
	SarcophagusesFilterer   // Log filterer for contract events
}

// SarcophagusesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SarcophagusesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SarcophagusesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SarcophagusesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SarcophagusesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SarcophagusesSession struct {
	Contract     *Sarcophaguses    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SarcophagusesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SarcophagusesCallerSession struct {
	Contract *SarcophagusesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SarcophagusesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SarcophagusesTransactorSession struct {
	Contract     *SarcophagusesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SarcophagusesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SarcophagusesRaw struct {
	Contract *Sarcophaguses // Generic contract binding to access the raw methods on
}

// SarcophagusesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SarcophagusesCallerRaw struct {
	Contract *SarcophagusesCaller // Generic read-only contract binding to access the raw methods on
}

// SarcophagusesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SarcophagusesTransactorRaw struct {
	Contract *SarcophagusesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSarcophaguses creates a new instance of Sarcophaguses, bound to a specific deployed contract.
func NewSarcophaguses(address common.Address, backend bind.ContractBackend) (*Sarcophaguses, error) {
	contract, err := bindSarcophaguses(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sarcophaguses{SarcophagusesCaller: SarcophagusesCaller{contract: contract}, SarcophagusesTransactor: SarcophagusesTransactor{contract: contract}, SarcophagusesFilterer: SarcophagusesFilterer{contract: contract}}, nil
}

// NewSarcophagusesCaller creates a new read-only instance of Sarcophaguses, bound to a specific deployed contract.
func NewSarcophagusesCaller(address common.Address, caller bind.ContractCaller) (*SarcophagusesCaller, error) {
	contract, err := bindSarcophaguses(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SarcophagusesCaller{contract: contract}, nil
}

// NewSarcophagusesTransactor creates a new write-only instance of Sarcophaguses, bound to a specific deployed contract.
func NewSarcophagusesTransactor(address common.Address, transactor bind.ContractTransactor) (*SarcophagusesTransactor, error) {
	contract, err := bindSarcophaguses(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SarcophagusesTransactor{contract: contract}, nil
}

// NewSarcophagusesFilterer creates a new log filterer instance of Sarcophaguses, bound to a specific deployed contract.
func NewSarcophagusesFilterer(address common.Address, filterer bind.ContractFilterer) (*SarcophagusesFilterer, error) {
	contract, err := bindSarcophaguses(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SarcophagusesFilterer{contract: contract}, nil
}

// bindSarcophaguses binds a generic wrapper to an already deployed contract.
func bindSarcophaguses(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SarcophagusesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sarcophaguses *SarcophagusesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sarcophaguses.Contract.SarcophagusesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sarcophaguses *SarcophagusesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sarcophaguses.Contract.SarcophagusesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sarcophaguses *SarcophagusesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sarcophaguses.Contract.SarcophagusesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sarcophaguses *SarcophagusesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sarcophaguses.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sarcophaguses *SarcophagusesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sarcophaguses.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sarcophaguses *SarcophagusesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sarcophaguses.Contract.contract.Transact(opts, method, params...)
}

// TypesABI is the input ABI used to generate the binding from.
const TypesABI = "[]"

// TypesBin is the compiled bytecode used for deploying new contracts.
var TypesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200c2728e78cb8738153b39bf9709c3ea30f7743cc91a81fa5afa5041c4ed629a964736f6c63430007020033"

// DeployTypes deploys a new Ethereum contract, binding an instance of Types to it.
func DeployTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Types, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Types{TypesCaller: TypesCaller{contract: contract}, TypesTransactor: TypesTransactor{contract: contract}, TypesFilterer: TypesFilterer{contract: contract}}, nil
}

// Types is an auto generated Go binding around an Ethereum contract.
type Types struct {
	TypesCaller     // Read-only binding to the contract
	TypesTransactor // Write-only binding to the contract
	TypesFilterer   // Log filterer for contract events
}

// TypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypesSession struct {
	Contract     *Types            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypesCallerSession struct {
	Contract *TypesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypesTransactorSession struct {
	Contract     *TypesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypesRaw struct {
	Contract *Types // Generic contract binding to access the raw methods on
}

// TypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypesCallerRaw struct {
	Contract *TypesCaller // Generic read-only contract binding to access the raw methods on
}

// TypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypesTransactorRaw struct {
	Contract *TypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypes creates a new instance of Types, bound to a specific deployed contract.
func NewTypes(address common.Address, backend bind.ContractBackend) (*Types, error) {
	contract, err := bindTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Types{TypesCaller: TypesCaller{contract: contract}, TypesTransactor: TypesTransactor{contract: contract}, TypesFilterer: TypesFilterer{contract: contract}}, nil
}

// NewTypesCaller creates a new read-only instance of Types, bound to a specific deployed contract.
func NewTypesCaller(address common.Address, caller bind.ContractCaller) (*TypesCaller, error) {
	contract, err := bindTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypesCaller{contract: contract}, nil
}

// NewTypesTransactor creates a new write-only instance of Types, bound to a specific deployed contract.
func NewTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*TypesTransactor, error) {
	contract, err := bindTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypesTransactor{contract: contract}, nil
}

// NewTypesFilterer creates a new log filterer instance of Types, bound to a specific deployed contract.
func NewTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*TypesFilterer, error) {
	contract, err := bindTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypesFilterer{contract: contract}, nil
}

// bindTypes binds a generic wrapper to an already deployed contract.
func bindTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.TypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Types.Contract.TypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Types.Contract.TypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Types.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Types.Contract.contract.Transact(opts, method, params...)
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"existingAssetId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newAssetId\",\"type\":\"string\"}],\"name\":\"assetIdsCheck\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"}],\"name\":\"confirmAssetIdNotSet\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"}],\"name\":\"getGracePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"doubleHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"singleHash\",\"type\":\"bytes\"}],\"name\":\"hashCheck\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"publicKeyLength\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"}],\"name\":\"resurrectionInFuture\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"embalmer\",\"type\":\"address\"}],\"name\":\"sarcophagusUpdater\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"archAddress\",\"type\":\"address\"}],\"name\":\"signatureCheck\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resurrectionWindow\",\"type\":\"uint256\"}],\"name\":\"unwrapTime\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"resurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumResurrectionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumDiggingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBounty\",\"type\":\"uint256\"}],\"name\":\"withinArchaeologistLimits\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x611092610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c80638899ecf7116100705780638899ecf7146102cd5780638ad5f4461461041f578063a25bb3651461051b578063c029b7b4146105e0578063c548201114610640576100a8565b80631807715f146100ad5780632be5910a146100db578063495fe007146101135780636d74499f146101ce57806377a5b2d914610289575b600080fd5b6100d9600480360360208110156100c357600080fd5b8101908080359060200190929190505050610682565b005b610111600480360360408110156100f157600080fd5b8101908080359060200190929190803590602001909291905050506106dd565b005b6101cc6004803603602081101561012957600080fd5b810190808035906020019064010000000081111561014657600080fd5b82018360208201111561015857600080fd5b8035906020019184600183028401116401000000008311171561017a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107a5565b005b610287600480360360208110156101e457600080fd5b810190808035906020019064010000000081111561020157600080fd5b82018360208201111561021357600080fd5b8035906020019184600183028401116401000000008311171561023557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061081f565b005b6102cb6004803603602081101561029f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610899565b005b61041d600480360360408110156102e357600080fd5b810190808035906020019064010000000081111561030057600080fd5b82018360208201111561031257600080fd5b8035906020019184600183028401116401000000008311171561033457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561039757600080fd5b8201836020820111156103a957600080fd5b803590602001918460018302840111640100000000831117156103cb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610920565b005b610519600480360360a081101561043557600080fd5b810190808035906020019064010000000081111561045257600080fd5b82018360208201111561046457600080fd5b8035906020019184600183028401116401000000008311171561048657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803560ff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109a4565b005b6105de6004803603604081101561053157600080fd5b81019080803590602001909291908035906020019064010000000081111561055857600080fd5b82018360208201111561056a57600080fd5b8035906020019184600183028401116401000000008311171561058c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610a9f565b005b61063e600480360360c08110156105f657600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190505050610b1f565b005b61066c6004803603602081101561065657600080fd5b8101908080359060200190929190505050610c7e565b6040518082815260200191505060405180910390f35b4281116106da576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610f716027913960400191505060405180910390fd5b50565b42821115610736576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610fbb6027913960400191505060405180910390fd5b4261074a8284610cce90919063ffffffff16565b10156107a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610f986023913960400191505060405180910390fd5b5050565b604081511461081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f7075626c6963206b6579206d757374206265203634206279746573000000000081525060200191505060405180910390fd5b50565b6000815114610896576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f617373657449642068617320616c7265616479206265656e207365740000000081525060200191505060405180910390fd5b50565b3373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180611009602b913960400191505060405180910390fd5b50565b6109298261081f565b60008151116109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f61737365744964206d757374206e6f7420686176652030206c656e677468000081525060200191505060405180910390fd5b5050565b60006001868051906020012086868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610a07573d6000803e3d6000fd5b5050506020604051035190508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a97576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806110346029913960400191505060405180910390fd5b505050505050565b80805190602001208214610b1b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f68617368657320646f206e6f74206d617463680000000000000000000000000081525060200191505060405180910390fd5b5050565b610b328342610cce90919063ffffffff16565b861115610b8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610fe26027913960400191505060405180910390fd5b81851015610c00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f64696767696e672066656520697320746f6f206c6f770000000000000000000081525060200191505060405180910390fd5b80841015610c76576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f626f756e747920697320746f6f206c6f7700000000000000000000000000000081525060200191505060405180910390fd5b505050505050565b60008061070890506000610cae6064610ca04287610d5690919063ffffffff16565b610da090919063ffffffff16565b90508161ffff16811015610cc4578161ffff1690505b8092505050919050565b600080828401905083811015610d4c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000610d9883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610dea565b905092915050565b6000610de283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610eaa565b905092915050565b6000838311158290610e97576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e5c578082015181840152602081019050610e41565b50505050905090810190601f168015610e895780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008083118290610f56576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f1b578082015181840152602081019050610f00565b50505050905090810190601f168015610f485780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581610f6257fe5b04905080915050939250505056fe726573757272656374696f6e2074696d65206d75737420626520696e207468652066757475726574686520726573757272656374696f6e2077696e646f7720686173206578706972656469742773206e6f742074696d6520746f20756e777261702074686520736172636f706861677573726573757272656374696f6e2074696d6520746f6f2066617220696e2074686520667574757265736172636f7068616775732063616e206f6e6c79206265207570646174656420627920656d62616c6d65727369676e617475726520646964206e6f7420636f6d652066726f6d206172636861656f6c6f67697374a26469706673582212204f4b012d1f7065f7f0072fb4660d91b3aa42b33590c327df636f2d5a1e7b8c0564736f6c63430007020033"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// AssetIdsCheck is a free data retrieval call binding the contract method 0x8899ecf7.
//
// Solidity: function assetIdsCheck(string existingAssetId, string newAssetId) pure returns()
func (_Utils *UtilsCaller) AssetIdsCheck(opts *bind.CallOpts, existingAssetId string, newAssetId string) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "assetIdsCheck", existingAssetId, newAssetId)
	return err
}

// AssetIdsCheck is a free data retrieval call binding the contract method 0x8899ecf7.
//
// Solidity: function assetIdsCheck(string existingAssetId, string newAssetId) pure returns()
func (_Utils *UtilsSession) AssetIdsCheck(existingAssetId string, newAssetId string) error {
	return _Utils.Contract.AssetIdsCheck(&_Utils.CallOpts, existingAssetId, newAssetId)
}

// AssetIdsCheck is a free data retrieval call binding the contract method 0x8899ecf7.
//
// Solidity: function assetIdsCheck(string existingAssetId, string newAssetId) pure returns()
func (_Utils *UtilsCallerSession) AssetIdsCheck(existingAssetId string, newAssetId string) error {
	return _Utils.Contract.AssetIdsCheck(&_Utils.CallOpts, existingAssetId, newAssetId)
}

// ConfirmAssetIdNotSet is a free data retrieval call binding the contract method 0x6d74499f.
//
// Solidity: function confirmAssetIdNotSet(string assetId) pure returns()
func (_Utils *UtilsCaller) ConfirmAssetIdNotSet(opts *bind.CallOpts, assetId string) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "confirmAssetIdNotSet", assetId)
	return err
}

// ConfirmAssetIdNotSet is a free data retrieval call binding the contract method 0x6d74499f.
//
// Solidity: function confirmAssetIdNotSet(string assetId) pure returns()
func (_Utils *UtilsSession) ConfirmAssetIdNotSet(assetId string) error {
	return _Utils.Contract.ConfirmAssetIdNotSet(&_Utils.CallOpts, assetId)
}

// ConfirmAssetIdNotSet is a free data retrieval call binding the contract method 0x6d74499f.
//
// Solidity: function confirmAssetIdNotSet(string assetId) pure returns()
func (_Utils *UtilsCallerSession) ConfirmAssetIdNotSet(assetId string) error {
	return _Utils.Contract.ConfirmAssetIdNotSet(&_Utils.CallOpts, assetId)
}

// GetGracePeriod is a free data retrieval call binding the contract method 0xc5482011.
//
// Solidity: function getGracePeriod(uint256 resurrectionTime) view returns(uint256)
func (_Utils *UtilsCaller) GetGracePeriod(opts *bind.CallOpts, resurrectionTime *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Utils.contract.Call(opts, out, "getGracePeriod", resurrectionTime)
	return *ret0, err
}

// GetGracePeriod is a free data retrieval call binding the contract method 0xc5482011.
//
// Solidity: function getGracePeriod(uint256 resurrectionTime) view returns(uint256)
func (_Utils *UtilsSession) GetGracePeriod(resurrectionTime *big.Int) (*big.Int, error) {
	return _Utils.Contract.GetGracePeriod(&_Utils.CallOpts, resurrectionTime)
}

// GetGracePeriod is a free data retrieval call binding the contract method 0xc5482011.
//
// Solidity: function getGracePeriod(uint256 resurrectionTime) view returns(uint256)
func (_Utils *UtilsCallerSession) GetGracePeriod(resurrectionTime *big.Int) (*big.Int, error) {
	return _Utils.Contract.GetGracePeriod(&_Utils.CallOpts, resurrectionTime)
}

// HashCheck is a free data retrieval call binding the contract method 0xa25bb365.
//
// Solidity: function hashCheck(bytes32 doubleHash, bytes singleHash) pure returns()
func (_Utils *UtilsCaller) HashCheck(opts *bind.CallOpts, doubleHash [32]byte, singleHash []byte) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "hashCheck", doubleHash, singleHash)
	return err
}

// HashCheck is a free data retrieval call binding the contract method 0xa25bb365.
//
// Solidity: function hashCheck(bytes32 doubleHash, bytes singleHash) pure returns()
func (_Utils *UtilsSession) HashCheck(doubleHash [32]byte, singleHash []byte) error {
	return _Utils.Contract.HashCheck(&_Utils.CallOpts, doubleHash, singleHash)
}

// HashCheck is a free data retrieval call binding the contract method 0xa25bb365.
//
// Solidity: function hashCheck(bytes32 doubleHash, bytes singleHash) pure returns()
func (_Utils *UtilsCallerSession) HashCheck(doubleHash [32]byte, singleHash []byte) error {
	return _Utils.Contract.HashCheck(&_Utils.CallOpts, doubleHash, singleHash)
}

// PublicKeyLength is a free data retrieval call binding the contract method 0x495fe007.
//
// Solidity: function publicKeyLength(bytes publicKey) pure returns()
func (_Utils *UtilsCaller) PublicKeyLength(opts *bind.CallOpts, publicKey []byte) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "publicKeyLength", publicKey)
	return err
}

// PublicKeyLength is a free data retrieval call binding the contract method 0x495fe007.
//
// Solidity: function publicKeyLength(bytes publicKey) pure returns()
func (_Utils *UtilsSession) PublicKeyLength(publicKey []byte) error {
	return _Utils.Contract.PublicKeyLength(&_Utils.CallOpts, publicKey)
}

// PublicKeyLength is a free data retrieval call binding the contract method 0x495fe007.
//
// Solidity: function publicKeyLength(bytes publicKey) pure returns()
func (_Utils *UtilsCallerSession) PublicKeyLength(publicKey []byte) error {
	return _Utils.Contract.PublicKeyLength(&_Utils.CallOpts, publicKey)
}

// ResurrectionInFuture is a free data retrieval call binding the contract method 0x1807715f.
//
// Solidity: function resurrectionInFuture(uint256 resurrectionTime) view returns()
func (_Utils *UtilsCaller) ResurrectionInFuture(opts *bind.CallOpts, resurrectionTime *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "resurrectionInFuture", resurrectionTime)
	return err
}

// ResurrectionInFuture is a free data retrieval call binding the contract method 0x1807715f.
//
// Solidity: function resurrectionInFuture(uint256 resurrectionTime) view returns()
func (_Utils *UtilsSession) ResurrectionInFuture(resurrectionTime *big.Int) error {
	return _Utils.Contract.ResurrectionInFuture(&_Utils.CallOpts, resurrectionTime)
}

// ResurrectionInFuture is a free data retrieval call binding the contract method 0x1807715f.
//
// Solidity: function resurrectionInFuture(uint256 resurrectionTime) view returns()
func (_Utils *UtilsCallerSession) ResurrectionInFuture(resurrectionTime *big.Int) error {
	return _Utils.Contract.ResurrectionInFuture(&_Utils.CallOpts, resurrectionTime)
}

// SarcophagusUpdater is a free data retrieval call binding the contract method 0x77a5b2d9.
//
// Solidity: function sarcophagusUpdater(address embalmer) view returns()
func (_Utils *UtilsCaller) SarcophagusUpdater(opts *bind.CallOpts, embalmer common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "sarcophagusUpdater", embalmer)
	return err
}

// SarcophagusUpdater is a free data retrieval call binding the contract method 0x77a5b2d9.
//
// Solidity: function sarcophagusUpdater(address embalmer) view returns()
func (_Utils *UtilsSession) SarcophagusUpdater(embalmer common.Address) error {
	return _Utils.Contract.SarcophagusUpdater(&_Utils.CallOpts, embalmer)
}

// SarcophagusUpdater is a free data retrieval call binding the contract method 0x77a5b2d9.
//
// Solidity: function sarcophagusUpdater(address embalmer) view returns()
func (_Utils *UtilsCallerSession) SarcophagusUpdater(embalmer common.Address) error {
	return _Utils.Contract.SarcophagusUpdater(&_Utils.CallOpts, embalmer)
}

// SignatureCheck is a free data retrieval call binding the contract method 0x8ad5f446.
//
// Solidity: function signatureCheck(bytes data, uint8 v, bytes32 r, bytes32 s, address archAddress) pure returns()
func (_Utils *UtilsCaller) SignatureCheck(opts *bind.CallOpts, data []byte, v uint8, r [32]byte, s [32]byte, archAddress common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "signatureCheck", data, v, r, s, archAddress)
	return err
}

// SignatureCheck is a free data retrieval call binding the contract method 0x8ad5f446.
//
// Solidity: function signatureCheck(bytes data, uint8 v, bytes32 r, bytes32 s, address archAddress) pure returns()
func (_Utils *UtilsSession) SignatureCheck(data []byte, v uint8, r [32]byte, s [32]byte, archAddress common.Address) error {
	return _Utils.Contract.SignatureCheck(&_Utils.CallOpts, data, v, r, s, archAddress)
}

// SignatureCheck is a free data retrieval call binding the contract method 0x8ad5f446.
//
// Solidity: function signatureCheck(bytes data, uint8 v, bytes32 r, bytes32 s, address archAddress) pure returns()
func (_Utils *UtilsCallerSession) SignatureCheck(data []byte, v uint8, r [32]byte, s [32]byte, archAddress common.Address) error {
	return _Utils.Contract.SignatureCheck(&_Utils.CallOpts, data, v, r, s, archAddress)
}

// UnwrapTime is a free data retrieval call binding the contract method 0x2be5910a.
//
// Solidity: function unwrapTime(uint256 resurrectionTime, uint256 resurrectionWindow) view returns()
func (_Utils *UtilsCaller) UnwrapTime(opts *bind.CallOpts, resurrectionTime *big.Int, resurrectionWindow *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "unwrapTime", resurrectionTime, resurrectionWindow)
	return err
}

// UnwrapTime is a free data retrieval call binding the contract method 0x2be5910a.
//
// Solidity: function unwrapTime(uint256 resurrectionTime, uint256 resurrectionWindow) view returns()
func (_Utils *UtilsSession) UnwrapTime(resurrectionTime *big.Int, resurrectionWindow *big.Int) error {
	return _Utils.Contract.UnwrapTime(&_Utils.CallOpts, resurrectionTime, resurrectionWindow)
}

// UnwrapTime is a free data retrieval call binding the contract method 0x2be5910a.
//
// Solidity: function unwrapTime(uint256 resurrectionTime, uint256 resurrectionWindow) view returns()
func (_Utils *UtilsCallerSession) UnwrapTime(resurrectionTime *big.Int, resurrectionWindow *big.Int) error {
	return _Utils.Contract.UnwrapTime(&_Utils.CallOpts, resurrectionTime, resurrectionWindow)
}

// WithinArchaeologistLimits is a free data retrieval call binding the contract method 0xc029b7b4.
//
// Solidity: function withinArchaeologistLimits(uint256 resurrectionTime, uint256 diggingFee, uint256 bounty, uint256 maximumResurrectionTime, uint256 minimumDiggingFee, uint256 minimumBounty) view returns()
func (_Utils *UtilsCaller) WithinArchaeologistLimits(opts *bind.CallOpts, resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int, maximumResurrectionTime *big.Int, minimumDiggingFee *big.Int, minimumBounty *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _Utils.contract.Call(opts, out, "withinArchaeologistLimits", resurrectionTime, diggingFee, bounty, maximumResurrectionTime, minimumDiggingFee, minimumBounty)
	return err
}

// WithinArchaeologistLimits is a free data retrieval call binding the contract method 0xc029b7b4.
//
// Solidity: function withinArchaeologistLimits(uint256 resurrectionTime, uint256 diggingFee, uint256 bounty, uint256 maximumResurrectionTime, uint256 minimumDiggingFee, uint256 minimumBounty) view returns()
func (_Utils *UtilsSession) WithinArchaeologistLimits(resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int, maximumResurrectionTime *big.Int, minimumDiggingFee *big.Int, minimumBounty *big.Int) error {
	return _Utils.Contract.WithinArchaeologistLimits(&_Utils.CallOpts, resurrectionTime, diggingFee, bounty, maximumResurrectionTime, minimumDiggingFee, minimumBounty)
}

// WithinArchaeologistLimits is a free data retrieval call binding the contract method 0xc029b7b4.
//
// Solidity: function withinArchaeologistLimits(uint256 resurrectionTime, uint256 diggingFee, uint256 bounty, uint256 maximumResurrectionTime, uint256 minimumDiggingFee, uint256 minimumBounty) view returns()
func (_Utils *UtilsCallerSession) WithinArchaeologistLimits(resurrectionTime *big.Int, diggingFee *big.Int, bounty *big.Int, maximumResurrectionTime *big.Int, minimumDiggingFee *big.Int, minimumBounty *big.Int) error {
	return _Utils.Contract.WithinArchaeologistLimits(&_Utils.CallOpts, resurrectionTime, diggingFee, bounty, maximumResurrectionTime, minimumDiggingFee, minimumBounty)
}
