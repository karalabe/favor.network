// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// FavornetABI is the input ABI used to generate the binding from.
const FavornetABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetRequestAt\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"bound\",\"type\":\"bool\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MakeRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetRequest\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"bound\",\"type\":\"bool\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DropRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// FavornetBin is the compiled bytecode used for deploying new contracts.
const FavornetBin = `0x6060604052341561000c57fe5b5b60016000555b5b610973806100236000396000f300606060405236156100755763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632219ddb0811461007757806338b5442d1461014d5780637b836de2146101715780639b370317146101d5578063d7d756291461028d578063f239b7bb146102bb575bfe5b341561007f57fe5b610096600160a060020a03600435166024356102d3565b6040518086815260200185600160a060020a0316600160a060020a03168152602001806020018415151515815260200183815260200182810382528581815181526020019150805190602001908083836000831461010f575b80518252602083111561010f57601f1990920191602091820191016100ef565b505050905090810190601f16801561013b5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561015557fe5b61016f600160a060020a03600435166024356044356103fb565b005b341561017957fe5b60408051602060046024803582810135601f810185900485028601850190965285855261016f958335600160a060020a0316959394604494939290920191819084018382808284375094965050933593506104d592505050565b005b34156101dd57fe5b6101e860043561061e565b60408051600160a060020a0386168152831515918101919091526060810182905260806020808301828152865192840192909252855160a08401918701908083838215610250575b80518252602083111561025057601f199092019160209182019101610230565b505050905090810190601f16801561027c5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561029557fe5b6102a9600160a060020a03600435166106fe565b60408051918252519081900360200190f35b34156102c357fe5b61016f60043560243561071d565b005b600060006102df610841565b600160a060020a038516600090815260036020526040812080548291829160019183918a90811061030c57fe5b906000526020600020900160005b5054815260208082019290925260409081016000208054600180830154600384015460048501546002808701805489516101009782161597909702600019011691909104601f81018a90048a0286018a019098528785529598509396600160a060020a039092169560ff9091169392918591908301828280156103de5780601f106103b3576101008083540402835291602001916103de565b820191906000526020600020905b8154815290600101906020018083116103c157829003601f168201915b50505050509250955095509550955095505b509295509295909350565b600160a060020a0383166000908152600360205260409020548083106104215760006000fd5b600160a060020a038416600090815260036020526040902080548391908590811061044857fe5b906000526020600020900160005b5054146104635760006000fd5b6000828152600160208190526040909120015433600160a060020a0390811691161461048f5760006000fd5b60008281526001602052604090206003015460ff16156104af5760006000fd5b6000828152600160208190526040909120600301805460ff191690911790555b50505050565b60008115610525575060008181526002602081905260409091209081015433600160a060020a0390811691161415806105115750600381015415155b1561051c5760006000fd5b60005460038201555b600160a060020a033316600090815260036020526040902080546001810161054d8382610853565b916000526020600020900160005b5060008054918290556040805160a081018252838152600160a060020a0389811660208084019182528385018b815260608501879052608085018b9052968652600180825294909520835181559051938101805473ffffffffffffffffffffffffffffffffffffffff1916949092169390931790559251805193945090926105e9926002850192019061087d565b50606082015160038201805460ff19169115159190911790556080909101516004909101556000805460010190555b50505050565b6000610628610841565b600083815260016020818152604080842080840154600382015460048301546002808501805487516101009a8216159a909a02600019011691909104601f810188900488028901880190965285885288979496600160a060020a0390941695909460ff909316939192918591908301828280156106e65780601f106106bb576101008083540402835291602001916106e6565b820191906000526020600020905b8154815290600101906020018083116106c957829003601f168201915b5050505050925094509450945094505b509193509193565b600160a060020a0381166000908152600360205260409020545b919050565b600160a060020a0333166000908152600360205260409020548083106107435760006000fd5b600160a060020a033316600090815260036020526040902080548391908590811061076a57fe5b906000526020600020900160005b5054146107855760006000fd5b60008281526001602052604090206003015460ff16156107a55760006000fd5b600160a060020a0333166000908152600360205260409020805460001983019081106107cd57fe5b906000526020600020900160005b5054600160a060020a033316600090815260036020526040902080548590811061080157fe5b906000526020600020900160005b5055600160a060020a03331660009081526003602052604090208054906104cf906000198301610853565b505b505050565b60408051602081019091526000815290565b81548183558181151161083c5760008381526020902061083c918101908301610926565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106108be57805160ff19168380011785556108eb565b828001600101855582156108eb579182015b828111156108eb5782518255916020019190600101906108d0565b5b506108f8929150610926565b5090565b81548183558181151161083c5760008381526020902061083c918101908301610926565b5b505050565b61094491905b808211156108f8576000815560010161092c565b5090565b905600a165627a7a7230582010231fff50e224f99a9e8f01d2a24afb3f33e596a574285f21f5192f43db23c30029`

// DeployFavornet deploys a new Ethereum contract, binding an instance of Favornet to it.
func DeployFavornet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Favornet, error) {
	parsed, err := abi.JSON(strings.NewReader(FavornetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FavornetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Favornet{FavornetCaller: FavornetCaller{contract: contract}, FavornetTransactor: FavornetTransactor{contract: contract}}, nil
}

// Favornet is an auto generated Go binding around an Ethereum contract.
type Favornet struct {
	FavornetCaller     // Read-only binding to the contract
	FavornetTransactor // Write-only binding to the contract
}

// FavornetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FavornetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FavornetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FavornetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FavornetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FavornetSession struct {
	Contract     *Favornet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FavornetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FavornetCallerSession struct {
	Contract *FavornetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FavornetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FavornetTransactorSession struct {
	Contract     *FavornetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FavornetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FavornetRaw struct {
	Contract *Favornet // Generic contract binding to access the raw methods on
}

// FavornetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FavornetCallerRaw struct {
	Contract *FavornetCaller // Generic read-only contract binding to access the raw methods on
}

// FavornetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FavornetTransactorRaw struct {
	Contract *FavornetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFavornet creates a new instance of Favornet, bound to a specific deployed contract.
func NewFavornet(address common.Address, backend bind.ContractBackend) (*Favornet, error) {
	contract, err := bindFavornet(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Favornet{FavornetCaller: FavornetCaller{contract: contract}, FavornetTransactor: FavornetTransactor{contract: contract}}, nil
}

// NewFavornetCaller creates a new read-only instance of Favornet, bound to a specific deployed contract.
func NewFavornetCaller(address common.Address, caller bind.ContractCaller) (*FavornetCaller, error) {
	contract, err := bindFavornet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FavornetCaller{contract: contract}, nil
}

// NewFavornetTransactor creates a new write-only instance of Favornet, bound to a specific deployed contract.
func NewFavornetTransactor(address common.Address, transactor bind.ContractTransactor) (*FavornetTransactor, error) {
	contract, err := bindFavornet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FavornetTransactor{contract: contract}, nil
}

// bindFavornet binds a generic wrapper to an already deployed contract.
func bindFavornet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FavornetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Favornet *FavornetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Favornet.Contract.FavornetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Favornet *FavornetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Favornet.Contract.FavornetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Favornet *FavornetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Favornet.Contract.FavornetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Favornet *FavornetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Favornet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Favornet *FavornetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Favornet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Favornet *FavornetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Favornet.Contract.contract.Transact(opts, method, params...)
}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(id uint256) constant returns(from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetCaller) GetRequest(opts *bind.CallOpts, id *big.Int) (struct {
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	ret := new(struct {
		From   common.Address
		Favor  string
		Bound  bool
		Reward *big.Int
	})
	out := ret
	err := _Favornet.contract.Call(opts, out, "GetRequest", id)
	return *ret, err
}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(id uint256) constant returns(from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetSession) GetRequest(id *big.Int) (struct {
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	return _Favornet.Contract.GetRequest(&_Favornet.CallOpts, id)
}

// GetRequest is a free data retrieval call binding the contract method 0x9b370317.
//
// Solidity: function GetRequest(id uint256) constant returns(from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetCallerSession) GetRequest(id *big.Int) (struct {
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	return _Favornet.Contract.GetRequest(&_Favornet.CallOpts, id)
}

// GetRequestAt is a free data retrieval call binding the contract method 0x2219ddb0.
//
// Solidity: function GetRequestAt(user address, index uint256) constant returns(id uint256, from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetCaller) GetRequestAt(opts *bind.CallOpts, user common.Address, index *big.Int) (struct {
	Id     *big.Int
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	ret := new(struct {
		Id     *big.Int
		From   common.Address
		Favor  string
		Bound  bool
		Reward *big.Int
	})
	out := ret
	err := _Favornet.contract.Call(opts, out, "GetRequestAt", user, index)
	return *ret, err
}

// GetRequestAt is a free data retrieval call binding the contract method 0x2219ddb0.
//
// Solidity: function GetRequestAt(user address, index uint256) constant returns(id uint256, from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetSession) GetRequestAt(user common.Address, index *big.Int) (struct {
	Id     *big.Int
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	return _Favornet.Contract.GetRequestAt(&_Favornet.CallOpts, user, index)
}

// GetRequestAt is a free data retrieval call binding the contract method 0x2219ddb0.
//
// Solidity: function GetRequestAt(user address, index uint256) constant returns(id uint256, from address, favor string, bound bool, reward uint256)
func (_Favornet *FavornetCallerSession) GetRequestAt(user common.Address, index *big.Int) (struct {
	Id     *big.Int
	From   common.Address
	Favor  string
	Bound  bool
	Reward *big.Int
}, error) {
	return _Favornet.Contract.GetRequestAt(&_Favornet.CallOpts, user, index)
}

// GetRequestCount is a free data retrieval call binding the contract method 0xd7d75629.
//
// Solidity: function GetRequestCount(user address) constant returns(uint256)
func (_Favornet *FavornetCaller) GetRequestCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Favornet.contract.Call(opts, out, "GetRequestCount", user)
	return *ret0, err
}

// GetRequestCount is a free data retrieval call binding the contract method 0xd7d75629.
//
// Solidity: function GetRequestCount(user address) constant returns(uint256)
func (_Favornet *FavornetSession) GetRequestCount(user common.Address) (*big.Int, error) {
	return _Favornet.Contract.GetRequestCount(&_Favornet.CallOpts, user)
}

// GetRequestCount is a free data retrieval call binding the contract method 0xd7d75629.
//
// Solidity: function GetRequestCount(user address) constant returns(uint256)
func (_Favornet *FavornetCallerSession) GetRequestCount(user common.Address) (*big.Int, error) {
	return _Favornet.Contract.GetRequestCount(&_Favornet.CallOpts, user)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x38b5442d.
//
// Solidity: function AcceptRequest(from address, index uint256, id uint256) returns()
func (_Favornet *FavornetTransactor) AcceptRequest(opts *bind.TransactOpts, from common.Address, index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.contract.Transact(opts, "AcceptRequest", from, index, id)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x38b5442d.
//
// Solidity: function AcceptRequest(from address, index uint256, id uint256) returns()
func (_Favornet *FavornetSession) AcceptRequest(from common.Address, index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.AcceptRequest(&_Favornet.TransactOpts, from, index, id)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x38b5442d.
//
// Solidity: function AcceptRequest(from address, index uint256, id uint256) returns()
func (_Favornet *FavornetTransactorSession) AcceptRequest(from common.Address, index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.AcceptRequest(&_Favornet.TransactOpts, from, index, id)
}

// DropRequest is a paid mutator transaction binding the contract method 0xf239b7bb.
//
// Solidity: function DropRequest(index uint256, id uint256) returns()
func (_Favornet *FavornetTransactor) DropRequest(opts *bind.TransactOpts, index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.contract.Transact(opts, "DropRequest", index, id)
}

// DropRequest is a paid mutator transaction binding the contract method 0xf239b7bb.
//
// Solidity: function DropRequest(index uint256, id uint256) returns()
func (_Favornet *FavornetSession) DropRequest(index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.DropRequest(&_Favornet.TransactOpts, index, id)
}

// DropRequest is a paid mutator transaction binding the contract method 0xf239b7bb.
//
// Solidity: function DropRequest(index uint256, id uint256) returns()
func (_Favornet *FavornetTransactorSession) DropRequest(index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.DropRequest(&_Favornet.TransactOpts, index, id)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x7b836de2.
//
// Solidity: function MakeRequest(from address, favor string, reward uint256) returns()
func (_Favornet *FavornetTransactor) MakeRequest(opts *bind.TransactOpts, from common.Address, favor string, reward *big.Int) (*types.Transaction, error) {
	return _Favornet.contract.Transact(opts, "MakeRequest", from, favor, reward)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x7b836de2.
//
// Solidity: function MakeRequest(from address, favor string, reward uint256) returns()
func (_Favornet *FavornetSession) MakeRequest(from common.Address, favor string, reward *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.MakeRequest(&_Favornet.TransactOpts, from, favor, reward)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x7b836de2.
//
// Solidity: function MakeRequest(from address, favor string, reward uint256) returns()
func (_Favornet *FavornetTransactorSession) MakeRequest(from common.Address, favor string, reward *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.MakeRequest(&_Favornet.TransactOpts, from, favor, reward)
}
