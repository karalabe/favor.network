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
const FavornetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DropPromise\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetRequestAt\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"bound\",\"type\":\"bool\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"reqIdx\",\"type\":\"uint256\"},{\"name\":\"reqId\",\"type\":\"uint256\"},{\"name\":\"promIdx\",\"type\":\"uint256\"}],\"name\":\"HonourRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MakeRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetRequest\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"bound\",\"type\":\"bool\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetPromiseAt\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"offered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetPromiseCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"GetPromise\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"favor\",\"type\":\"string\"},{\"name\":\"offered\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DropRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// FavornetBin is the compiled bytecode used for deploying new contracts.
const FavornetBin = `0x6060604052341561000c57fe5b5b60016000555b5b6114ba806100236000396000f300606060405236156100ac5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166320a30d6981146100ae5780632219ddb0146100c657806338b5442d1461019c57806365b93db7146101c05780637b836de2146101db5780639b3703171461023f578063bb1e6392146102f7578063bf6b6090146103df578063cb9db93e1461040d578063d7d75629146104c9578063f239b7bb146104f7575bfe5b34156100b657fe5b6100c460043560243561050f565b005b34156100ce57fe5b6100e5600160a060020a0360043516602435610680565b6040518086815260200185600160a060020a0316600160a060020a03168152602001806020018415151515815260200183815260200182810382528581815181526020019150805190602001908083836000831461015e575b80518252602083111561015e57601f19909201916020918201910161013e565b505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34156101a457fe5b6100c4600160a060020a03600435166024356044356107a8565b005b34156101c857fe5b6100c4600435602435604435610882565b005b34156101e357fe5b60408051602060046024803582810135601f81018590048502860185019096528585526100c4958335600160a060020a031695939460449493929092019181908401838280828437509496505093359350610d3192505050565b005b341561024757fe5b610252600435610e73565b60408051600160a060020a0386168152831515918101919091526060810182905260806020808301828152865192840192909252855160a084019187019080838382156102ba575b8051825260208311156102ba57601f19909201916020918201910161029a565b505050905090810190601f1680156102e65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156102ff57fe5b610316600160a060020a0360043516602435610f53565b6040518086815260200185600160a060020a0316600160a060020a0316815260200184600160a060020a0316600160a060020a03168152602001806020018315151515815260200182810382528481815181526020019150805190602001908083836000831461015e575b80518252602083111561015e57601f19909201916020918201910161013e565b505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34156103e757fe5b6103fb600160a060020a0360043516611081565b60408051918252519081900360200190f35b341561041557fe5b6104206004356110a0565b60408051600160a060020a038087168252851660208083019190915283151560608301526080928201838152855193830193909352845191929160a084019186019080838382156102ba575b8051825260208311156102ba57601f19909201916020918201910161029a565b505050905090810190601f1680156102e65780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34156104d157fe5b6103fb600160a060020a0360043516611181565b60408051918252519081900360200190f35b34156104ff57fe5b6100c46004356024356111a0565b005b600160a060020a0333166000908152600460205260409020548083106105355760006000fd5b600160a060020a033316600090815260046020526040902080548391908590811061055c57fe5b906000526020600020900160005b5054146105775760006000fd5b60008281526002602052604090206004015460ff16156105975760006000fd5b600160a060020a0333166000908152600460205260409020805460001983019081106105bf57fe5b906000526020600020900160005b5054600160a060020a03331660009081526004602052604090208054859081106105f357fe5b906000526020600020900160005b5055600160a060020a033316600090815260046020526040902080549061062c906000198301611340565b5060008281526002602081905260408220828155600181018054600160a060020a031990811690915591810180549092169091559061066e600383018261136a565b50600401805460ff191690555b505050565b6000600061068c6113b2565b600160a060020a038516600090815260036020526040812080548291829160019183918a9081106106b957fe5b906000526020600020900160005b5054815260208082019290925260409081016000208054600180830154600384015460048501546002808701805489516101009782161597909702600019011691909104601f81018a90048a0286018a019098528785529598509396600160a060020a039092169560ff90911693929185919083018282801561078b5780601f106107605761010080835404028352916020019161078b565b820191906000526020600020905b81548152906001019060200180831161076e57829003601f168201915b50505050509250955095509550955095505b509295509295909350565b600160a060020a0383166000908152600360205260409020548083106107ce5760006000fd5b600160a060020a03841660009081526003602052604090208054839190859081106107f557fe5b906000526020600020900160005b5054146108105760006000fd5b6000828152600160208190526040909120015433600160a060020a0390811691161461083c5760006000fd5b60008281526001602052604090206003015460ff161561085c5760006000fd5b6000828152600160208190526040909120600301805460ff191690911790555b50505050565b600160a060020a03331660009081526003602052604081205490808286106108aa5760006000fd5b600160a060020a03331660009081526003602052604090208054869190889081106108d157fe5b906000526020600020900160005b5054146108ec5760006000fd5b60008581526001602052604090206003015460ff16151561090d5760006000fd5b50506000838152600160209081526040808320600160a060020a033316845260049283905292205490820154156109925780841061094b5760006000fd5b600480830154600160a060020a0333166000908152602092909252604090912080548690811061097757fe5b906000526020600020900160005b5054146109925760006000fd5b5b60048201541515610b32576040805160a081018252868152600184810154600160a060020a0390811660208085019190915233909116838501526002808701805486516101009582161595909502600019011691909104601f81018390048302840183019095528483529293606085019391830182828015610a565780601f10610a2b57610100808354040283529160200191610a56565b820191906000526020600020905b815481529060010190602001808311610a3957829003601f168201915b50505091835250506000602091820181905287815260028083526040918290208451815584840151600182018054600160a060020a0319908116600160a060020a0393841617909155938601519282018054909416921691909117909155606083015180519192610acf926003850192909101906113c4565b50608091909101516004918201805460ff1916911515919091179055600183810154600160a060020a031660009081526020929092526040909120805490918101610b1a8382611340565b916000526020600020900160005b5086905550610c4b565b6001808301805460048086018054600090815260026020908152604080832088018054600160a060020a031916600160a060020a03978816179055925482528282208401805460ff1916905594549093168352925220805490918101610b988382611340565b916000526020600020900160005b50600480850154909155600160a060020a03331660009081526020919091526040902080549091506000198301908110610bdc57fe5b906000526020600020900160005b5054600160a060020a0333166000908152600460205260409020805486908110610c1057fe5b906000526020600020900160005b5055600160a060020a0333166000908152600460205260409020805490610c49906000198301611340565b505b600160a060020a033316600090815260036020526040902080546000198501908110610c7357fe5b906000526020600020900160005b5054600160a060020a0333166000908152600360205260409020805488908110610ca757fe5b906000526020600020900160005b5055600160a060020a0333166000908152600360205260409020805490610ce0906000198301611340565b50600085815260016020819052604082208281559081018054600160a060020a031916905590610d13600283018261136a565b5060038101805460ff1916905560006004909101555b505050505050565b60008115610d8757506000818152600260205260409020600181015433600160a060020a039081169116141580610d6c5750600481015460ff165b15610d775760006000fd5b60048101805460ff191660011790555b600160a060020a0333166000908152600360205260409020805460018101610daf8382611340565b916000526020600020900160005b5060008054918290556040805160a081018252838152600160a060020a0389811660208084019182528385018b815260608501879052608085018b90529686526001808252949095208351815590519381018054600160a060020a03191694909216939093179055925180519394509092610e3e92600285019201906113c4565b50606082015160038201805460ff19169115159190911790556080909101516004909101556000805460010190555b50505050565b6000610e7d6113b2565b600083815260016020818152604080842080840154600382015460048301546002808501805487516101009a8216159a909a02600019011691909104601f810188900488028901880190965285885288979496600160a060020a0390941695909460ff90931693919291859190830182828015610f3b5780601f10610f1057610100808354040283529160200191610f3b565b820191906000526020600020905b815481529060010190602001808311610f1e57829003601f168201915b5050505050925094509450945094505b509193509193565b600060006000610f616113b2565b600160a060020a0386166000908152600460205260408120805482916002918391908a908110610f8d57fe5b906000526020600020900160005b5054815260208082019290925260409081016000208054600180830154600280850154600486015460038701805489516101009782161597909702600019011693909304601f81018a90048a0286018a019098528785529598509396600160a060020a03928316969290941694909360ff909116929184918301828280156110645780601f1061103957610100808354040283529160200191611064565b820191906000526020600020905b81548152906001019060200180831161104757829003601f168201915b50505050509150955095509550955095505b509295509295909350565b600160a060020a0381166000908152600460205260409020545b919050565b600060006110ac6113b2565b60008481526002602081815260408084206001808201548286015460048401546003850180548751601f600019978316156101000297909701909116999099049485018890048802890188019096528388529396600160a060020a039283169691909216949360ff169284918301828280156111695780601f1061113e57610100808354040283529160200191611169565b820191906000526020600020905b81548152906001019060200180831161114c57829003601f168201915b5050505050915094509450945094505b509193509193565b600160a060020a0381166000908152600360205260409020545b919050565b600160a060020a033316600090815260036020526040812054908184106111c75760006000fd5b600160a060020a03331660009081526003602052604090208054849190869081106111ee57fe5b906000526020600020900160005b5054146112095760006000fd5b60008381526001602052604090206003015460ff16156112295760006000fd5b50600082815260016020526040902060040154801561125c576000818152600260205260409020600401805460ff191690555b600160a060020a03331660009081526003602052604090208054600019840190811061128457fe5b906000526020600020900160005b5054600160a060020a03331660009081526003602052604090208054869081106112b857fe5b906000526020600020900160005b5055600160a060020a03331660009081526003602052604090208054906112f1906000198301611340565b50600083815260016020819052604082208281559081018054600160a060020a031916905590611324600283018261136a565b5060038101805460ff1916905560006004909101555b50505050565b81548183558181151161067b5760008381526020902061067b91810190830161146d565b5b505050565b50805460018160011615610100020316600290046000825580601f1061139057506113ae565b601f0160209004906000526020600020908101906113ae919061146d565b5b50565b60408051602081019091526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061140557805160ff1916838001178555611432565b82800160010185558215611432579182015b82811115611432578251825591602001919060010190611417565b5b5061143f92915061146d565b5090565b81548183558181151161067b5760008381526020902061067b91810190830161146d565b5b505050565b61148b91905b8082111561143f5760008155600101611473565b5090565b905600a165627a7a72305820227ade235ca9c98ee0a05e6e6c14e5fbc6999394d784256fde1d02c8d577ec360029`

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

// GetPromise is a free data retrieval call binding the contract method 0xcb9db93e.
//
// Solidity: function GetPromise(id uint256) constant returns(owner address, from address, favor string, offered bool)
func (_Favornet *FavornetCaller) GetPromise(opts *bind.CallOpts, id *big.Int) (struct {
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	ret := new(struct {
		Owner   common.Address
		From    common.Address
		Favor   string
		Offered bool
	})
	out := ret
	err := _Favornet.contract.Call(opts, out, "GetPromise", id)
	return *ret, err
}

// GetPromise is a free data retrieval call binding the contract method 0xcb9db93e.
//
// Solidity: function GetPromise(id uint256) constant returns(owner address, from address, favor string, offered bool)
func (_Favornet *FavornetSession) GetPromise(id *big.Int) (struct {
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	return _Favornet.Contract.GetPromise(&_Favornet.CallOpts, id)
}

// GetPromise is a free data retrieval call binding the contract method 0xcb9db93e.
//
// Solidity: function GetPromise(id uint256) constant returns(owner address, from address, favor string, offered bool)
func (_Favornet *FavornetCallerSession) GetPromise(id *big.Int) (struct {
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	return _Favornet.Contract.GetPromise(&_Favornet.CallOpts, id)
}

// GetPromiseAt is a free data retrieval call binding the contract method 0xbb1e6392.
//
// Solidity: function GetPromiseAt(user address, index uint256) constant returns(id uint256, owner address, from address, favor string, offered bool)
func (_Favornet *FavornetCaller) GetPromiseAt(opts *bind.CallOpts, user common.Address, index *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	ret := new(struct {
		Id      *big.Int
		Owner   common.Address
		From    common.Address
		Favor   string
		Offered bool
	})
	out := ret
	err := _Favornet.contract.Call(opts, out, "GetPromiseAt", user, index)
	return *ret, err
}

// GetPromiseAt is a free data retrieval call binding the contract method 0xbb1e6392.
//
// Solidity: function GetPromiseAt(user address, index uint256) constant returns(id uint256, owner address, from address, favor string, offered bool)
func (_Favornet *FavornetSession) GetPromiseAt(user common.Address, index *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	return _Favornet.Contract.GetPromiseAt(&_Favornet.CallOpts, user, index)
}

// GetPromiseAt is a free data retrieval call binding the contract method 0xbb1e6392.
//
// Solidity: function GetPromiseAt(user address, index uint256) constant returns(id uint256, owner address, from address, favor string, offered bool)
func (_Favornet *FavornetCallerSession) GetPromiseAt(user common.Address, index *big.Int) (struct {
	Id      *big.Int
	Owner   common.Address
	From    common.Address
	Favor   string
	Offered bool
}, error) {
	return _Favornet.Contract.GetPromiseAt(&_Favornet.CallOpts, user, index)
}

// GetPromiseCount is a free data retrieval call binding the contract method 0xbf6b6090.
//
// Solidity: function GetPromiseCount(user address) constant returns(uint256)
func (_Favornet *FavornetCaller) GetPromiseCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Favornet.contract.Call(opts, out, "GetPromiseCount", user)
	return *ret0, err
}

// GetPromiseCount is a free data retrieval call binding the contract method 0xbf6b6090.
//
// Solidity: function GetPromiseCount(user address) constant returns(uint256)
func (_Favornet *FavornetSession) GetPromiseCount(user common.Address) (*big.Int, error) {
	return _Favornet.Contract.GetPromiseCount(&_Favornet.CallOpts, user)
}

// GetPromiseCount is a free data retrieval call binding the contract method 0xbf6b6090.
//
// Solidity: function GetPromiseCount(user address) constant returns(uint256)
func (_Favornet *FavornetCallerSession) GetPromiseCount(user common.Address) (*big.Int, error) {
	return _Favornet.Contract.GetPromiseCount(&_Favornet.CallOpts, user)
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

// DropPromise is a paid mutator transaction binding the contract method 0x20a30d69.
//
// Solidity: function DropPromise(index uint256, id uint256) returns()
func (_Favornet *FavornetTransactor) DropPromise(opts *bind.TransactOpts, index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.contract.Transact(opts, "DropPromise", index, id)
}

// DropPromise is a paid mutator transaction binding the contract method 0x20a30d69.
//
// Solidity: function DropPromise(index uint256, id uint256) returns()
func (_Favornet *FavornetSession) DropPromise(index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.DropPromise(&_Favornet.TransactOpts, index, id)
}

// DropPromise is a paid mutator transaction binding the contract method 0x20a30d69.
//
// Solidity: function DropPromise(index uint256, id uint256) returns()
func (_Favornet *FavornetTransactorSession) DropPromise(index *big.Int, id *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.DropPromise(&_Favornet.TransactOpts, index, id)
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

// HonourRequest is a paid mutator transaction binding the contract method 0x65b93db7.
//
// Solidity: function HonourRequest(reqIdx uint256, reqId uint256, promIdx uint256) returns()
func (_Favornet *FavornetTransactor) HonourRequest(opts *bind.TransactOpts, reqIdx *big.Int, reqId *big.Int, promIdx *big.Int) (*types.Transaction, error) {
	return _Favornet.contract.Transact(opts, "HonourRequest", reqIdx, reqId, promIdx)
}

// HonourRequest is a paid mutator transaction binding the contract method 0x65b93db7.
//
// Solidity: function HonourRequest(reqIdx uint256, reqId uint256, promIdx uint256) returns()
func (_Favornet *FavornetSession) HonourRequest(reqIdx *big.Int, reqId *big.Int, promIdx *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.HonourRequest(&_Favornet.TransactOpts, reqIdx, reqId, promIdx)
}

// HonourRequest is a paid mutator transaction binding the contract method 0x65b93db7.
//
// Solidity: function HonourRequest(reqIdx uint256, reqId uint256, promIdx uint256) returns()
func (_Favornet *FavornetTransactorSession) HonourRequest(reqIdx *big.Int, reqId *big.Int, promIdx *big.Int) (*types.Transaction, error) {
	return _Favornet.Contract.HonourRequest(&_Favornet.TransactOpts, reqIdx, reqId, promIdx)
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
