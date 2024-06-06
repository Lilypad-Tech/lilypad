// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package powtoken

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

// PowtokenMetaData contains all meta data concerning the Powtoken contract.
var PowtokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPostRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHashrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611705806101065f395ff3fe608060405234801561000f575f80fd5b50600436106100b2575f3560e01c80638da5cb5b1161006f5780638da5cb5b1461016a578063ab09894514610188578063b28d87ea146101a4578063b681f2fd146101c2578063da8accf9146101cc578063f2fde38b146101e8576100b2565b80633181c380146100b65780634bbe05e4146100d45780636189f3ac14610108578063715018a61461013857806378e97925146101425780638129fc1c14610160575b5f80fd5b6100be610204565b6040516100cb9190610ac7565b60405180910390f35b6100ee60048036038101906100e99190610b6c565b61026c565b6040516100ff959493929190610c5b565b60405180910390f35b610122600480360381019061011d9190610cb3565b61035d565b60405161012f9190610cde565b60405180910390f35b610140610372565b005b61014a610385565b6040516101579190610ac7565b60405180910390f35b61016861038b565b005b6101726104c0565b60405161017f9190610cf7565b60405180910390f35b6101a2600480360381019061019d9190610d71565b6104e7565b005b6101ac610595565b6040516101b99190610ac7565b60405180910390f35b6101ca61059b565b005b6101e660048036038101906101e19190610dbc565b6105d1565b005b61020260048036038101906101fd9190610cb3565b6108c5565b005b5f80600454426102149190610e46565b90505f6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6102459190610ea6565b90505f816003546102569190610ed6565b905082816102649190610ea6565b935050505090565b6001602052815f5260405f208181548110610285575f80fd5b905f5260205f2090600502015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546102ca90610f44565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690610f44565b80156103415780601f1061031857610100808354040283529160200191610341565b820191905f5260205f20905b81548152906001019060200180831161032457829003601f168201915b5050505050908060020154908060030154908060040154905085565b6002602052805f5260405f205f915090505481565b61037a610947565b6103835f6109c5565b565b60045481565b5f8060159054906101000a900460ff161590508080156103bc575060015f60149054906101000a900460ff1660ff16105b806103ea57506103cb30610a86565b1580156103e9575060015f60149054906101000a900460ff1660ff16145b5b610429576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042090610fe4565b60405180910390fd5b60015f60146101000a81548160ff021916908360ff16021790555080156104655760015f60156101000a81548160ff0219169083151502179055505b80156104bd575f8060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516104b49190611050565b60405180910390a15b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f423384846040516020016104ff949392919061110a565b6040516020818303038152906040528051906020012090508060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507f2d3b66c3b514494f1330d833a6e605138788b0a382e7a282071485298300d770816040516105889190610cde565b60405180910390a1505050565b60035481565b6105a3610947565b7f6828ec9c4d2d50ad73ca05fd4575a57180db77534ab19a16e9fbd0e30393d20560405160405180910390a1565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8133868660405160200161066a9493929190611164565b6040516020818303038152906040528051906020012090505f818760405160200161069692919061119e565b6040516020818303038152906040528051906020012090505f6001446106bc91906111c9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6106e79190610ea6565b905080825f1c1061072d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107249061126c565b60405180910390fd5b60035f81548092919061073f9061128a565b91905055503373ffffffffffffffffffffffffffffffffffffffff167fab3368117a8c701589001ea9b7cd7edb08dbb5cd3edba3b58272d5f23dfd654988888b42896040516107929594939291906112fd565b60405180910390a2846040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018a815260200142815260200186815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161089a919061150a565b5060408201518160020155606082015181600301556080820151816004015550505050505050505050565b6108cd610947565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361093b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093290611649565b60405180910390fd5b610944816109c5565b50565b61094f610aa8565b73ffffffffffffffffffffffffffffffffffffffff1661096d6104c0565b73ffffffffffffffffffffffffffffffffffffffff16146109c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ba906116b1565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f819050919050565b610ac181610aaf565b82525050565b5f602082019050610ada5f830184610ab8565b92915050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b1182610ae8565b9050919050565b610b2181610b07565b8114610b2b575f80fd5b50565b5f81359050610b3c81610b18565b92915050565b610b4b81610aaf565b8114610b55575f80fd5b50565b5f81359050610b6681610b42565b92915050565b5f8060408385031215610b8257610b81610ae0565b5b5f610b8f85828601610b2e565b9250506020610ba085828601610b58565b9150509250929050565b610bb381610b07565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610bf0578082015181840152602081019050610bd5565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610c1582610bb9565b610c1f8185610bc3565b9350610c2f818560208601610bd3565b610c3881610bfb565b840191505092915050565b5f819050919050565b610c5581610c43565b82525050565b5f60a082019050610c6e5f830188610baa565b8181036020830152610c808187610c0b565b9050610c8f6040830186610ab8565b610c9c6060830185610ab8565b610ca96080830184610c4c565b9695505050505050565b5f60208284031215610cc857610cc7610ae0565b5b5f610cd584828501610b2e565b91505092915050565b5f602082019050610cf15f830184610c4c565b92915050565b5f602082019050610d0a5f830184610baa565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610d3157610d30610d10565b5b8235905067ffffffffffffffff811115610d4e57610d4d610d14565b5b602083019150836001820283011115610d6a57610d69610d18565b5b9250929050565b5f8060208385031215610d8757610d86610ae0565b5b5f83013567ffffffffffffffff811115610da457610da3610ae4565b5b610db085828601610d1c565b92509250509250929050565b5f805f60408486031215610dd357610dd2610ae0565b5b5f610de086828701610b58565b935050602084013567ffffffffffffffff811115610e0157610e00610ae4565b5b610e0d86828701610d1c565b92509250509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e5082610aaf565b9150610e5b83610aaf565b9250828203905081811115610e7357610e72610e19565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610eb082610aaf565b9150610ebb83610aaf565b925082610ecb57610eca610e79565b5b828204905092915050565b5f610ee082610aaf565b9150610eeb83610aaf565b9250828202610ef981610aaf565b91508282048414831517610f1057610f0f610e19565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f5b57607f821691505b602082108103610f6e57610f6d610f17565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f610fce602e83610bc3565b9150610fd982610f74565b604082019050919050565b5f6020820190508181035f830152610ffb81610fc2565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61103a61103561103084611002565b611017565b61100b565b9050919050565b61104a81611020565b82525050565b5f6020820190506110635f830184611041565b92915050565b5f819050919050565b61108361107e82610aaf565b611069565b82525050565b5f8160601b9050919050565b5f61109f82611089565b9050919050565b5f6110b082611095565b9050919050565b6110c86110c382610b07565b6110a6565b82525050565b5f81905092915050565b828183375f83830152505050565b5f6110f183856110ce565b93506110fe8385846110d8565b82840190509392505050565b5f6111158287611072565b60208201915061112582866110b7565b6014820191506111368284866110e6565b915081905095945050505050565b5f819050919050565b61115e61115982610c43565b611144565b82525050565b5f61116f828761114d565b60208201915061117f82866110b7565b6014820191506111908284866110e6565b915081905095945050505050565b5f6111a9828561114d565b6020820191506111b98284611072565b6020820191508190509392505050565b5f6111d382610aaf565b91506111de83610aaf565b92508282019050808211156111f6576111f5610e19565b5b92915050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f611256602483610bc3565b9150611261826111fc565b604082019050919050565b5f6020820190508181035f8301526112838161124a565b9050919050565b5f61129482610aaf565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036112c6576112c5610e19565b5b600182019050919050565b5f6112dc8385610bc3565b93506112e98385846110d8565b6112f283610bfb565b840190509392505050565b5f6080820190508181035f8301526113168187896112d1565b90506113256020830186610ab8565b6113326040830185610ab8565b61133f6060830184610c4c565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113d27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611397565b6113dc8683611397565b95508019841693508086168417925050509392505050565b5f61140e61140961140484610aaf565b611017565b610aaf565b9050919050565b5f819050919050565b611427836113f4565b61143b61143382611415565b8484546113a3565b825550505050565b5f90565b61144f611443565b61145a81848461141e565b505050565b5b8181101561147d576114725f82611447565b600181019050611460565b5050565b601f8211156114c25761149381611376565b61149c84611388565b810160208510156114ab578190505b6114bf6114b785611388565b83018261145f565b50505b505050565b5f82821c905092915050565b5f6114e25f19846008026114c7565b1980831691505092915050565b5f6114fa83836114d3565b9150826002028217905092915050565b61151382610bb9565b67ffffffffffffffff81111561152c5761152b611349565b5b6115368254610f44565b611541828285611481565b5f60209050601f831160018114611572575f8415611560578287015190505b61156a85826114ef565b8655506115d1565b601f19841661158086611376565b5f5b828110156115a757848901518255600182019150602085019450602081019050611582565b868310156115c457848901516115c0601f8916826114d3565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611633602683610bc3565b915061163e826115d9565b604082019050919050565b5f6020820190508181035f83015261166081611627565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61169b602083610bc3565b91506116a682611667565b602082019050919050565b5f6020820190508181035f8301526116c88161168f565b905091905056fea264697066735822122065878ce96c77a7406e43b756d55517e4b3867de07d9e23f4dfa7c000c7565c6a64736f6c63430008150033",
}

// PowtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PowtokenMetaData.ABI instead.
var PowtokenABI = PowtokenMetaData.ABI

// PowtokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PowtokenMetaData.Bin instead.
var PowtokenBin = PowtokenMetaData.Bin

// DeployPowtoken deploys a new Ethereum contract, binding an instance of Powtoken to it.
func DeployPowtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Powtoken, error) {
	parsed, err := PowtokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PowtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Powtoken{PowtokenCaller: PowtokenCaller{contract: contract}, PowtokenTransactor: PowtokenTransactor{contract: contract}, PowtokenFilterer: PowtokenFilterer{contract: contract}}, nil
}

// Powtoken is an auto generated Go binding around an Ethereum contract.
type Powtoken struct {
	PowtokenCaller     // Read-only binding to the contract
	PowtokenTransactor // Write-only binding to the contract
	PowtokenFilterer   // Log filterer for contract events
}

// PowtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PowtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowtokenSession struct {
	Contract     *Powtoken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowtokenCallerSession struct {
	Contract *PowtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PowtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowtokenTransactorSession struct {
	Contract     *PowtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PowtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowtokenRaw struct {
	Contract *Powtoken // Generic contract binding to access the raw methods on
}

// PowtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowtokenCallerRaw struct {
	Contract *PowtokenCaller // Generic read-only contract binding to access the raw methods on
}

// PowtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowtokenTransactorRaw struct {
	Contract *PowtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPowtoken creates a new instance of Powtoken, bound to a specific deployed contract.
func NewPowtoken(address common.Address, backend bind.ContractBackend) (*Powtoken, error) {
	contract, err := bindPowtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Powtoken{PowtokenCaller: PowtokenCaller{contract: contract}, PowtokenTransactor: PowtokenTransactor{contract: contract}, PowtokenFilterer: PowtokenFilterer{contract: contract}}, nil
}

// NewPowtokenCaller creates a new read-only instance of Powtoken, bound to a specific deployed contract.
func NewPowtokenCaller(address common.Address, caller bind.ContractCaller) (*PowtokenCaller, error) {
	contract, err := bindPowtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PowtokenCaller{contract: contract}, nil
}

// NewPowtokenTransactor creates a new write-only instance of Powtoken, bound to a specific deployed contract.
func NewPowtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PowtokenTransactor, error) {
	contract, err := bindPowtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PowtokenTransactor{contract: contract}, nil
}

// NewPowtokenFilterer creates a new log filterer instance of Powtoken, bound to a specific deployed contract.
func NewPowtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PowtokenFilterer, error) {
	contract, err := bindPowtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PowtokenFilterer{contract: contract}, nil
}

// bindPowtoken binds a generic wrapper to an already deployed contract.
func bindPowtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PowtokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powtoken *PowtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powtoken.Contract.PowtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powtoken *PowtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powtoken.Contract.PowtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powtoken *PowtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powtoken.Contract.PowtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Powtoken *PowtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Powtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Powtoken *PowtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Powtoken *PowtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Powtoken.Contract.contract.Transact(opts, method, params...)
}

// GetHashrate is a free data retrieval call binding the contract method 0x3181c380.
//
// Solidity: function getHashrate() view returns(uint256)
func (_Powtoken *PowtokenCaller) GetHashrate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "getHashrate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHashrate is a free data retrieval call binding the contract method 0x3181c380.
//
// Solidity: function getHashrate() view returns(uint256)
func (_Powtoken *PowtokenSession) GetHashrate() (*big.Int, error) {
	return _Powtoken.Contract.GetHashrate(&_Powtoken.CallOpts)
}

// GetHashrate is a free data retrieval call binding the contract method 0x3181c380.
//
// Solidity: function getHashrate() view returns(uint256)
func (_Powtoken *PowtokenCallerSession) GetHashrate() (*big.Int, error) {
	return _Powtoken.Contract.GetHashrate(&_Powtoken.CallOpts)
}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32)
func (_Powtoken *PowtokenCaller) LastChallenges(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "lastChallenges", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32)
func (_Powtoken *PowtokenSession) LastChallenges(arg0 common.Address) ([32]byte, error) {
	return _Powtoken.Contract.LastChallenges(&_Powtoken.CallOpts, arg0)
}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32)
func (_Powtoken *PowtokenCallerSession) LastChallenges(arg0 common.Address) ([32]byte, error) {
	return _Powtoken.Contract.LastChallenges(&_Powtoken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Powtoken *PowtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Powtoken *PowtokenSession) Owner() (common.Address, error) {
	return _Powtoken.Contract.Owner(&_Powtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Powtoken *PowtokenCallerSession) Owner() (common.Address, error) {
	return _Powtoken.Contract.Owner(&_Powtoken.CallOpts)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenCaller) PowSubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
}, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "powSubmissions", arg0, arg1)

	outstruct := new(struct {
		WalletAddress common.Address
		NodeId        string
		Nonce         *big.Int
		Timestamp     *big.Int
		Challenge     [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Challenge = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
}, error) {
	return _Powtoken.Contract.PowSubmissions(&_Powtoken.CallOpts, arg0, arg1)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenCallerSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
}, error) {
	return _Powtoken.Contract.PowSubmissions(&_Powtoken.CallOpts, arg0, arg1)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Powtoken *PowtokenCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Powtoken *PowtokenSession) StartTime() (*big.Int, error) {
	return _Powtoken.Contract.StartTime(&_Powtoken.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Powtoken *PowtokenCallerSession) StartTime() (*big.Int, error) {
	return _Powtoken.Contract.StartTime(&_Powtoken.CallOpts)
}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Powtoken *PowtokenCaller) ValidProofs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "validProofs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Powtoken *PowtokenSession) ValidProofs() (*big.Int, error) {
	return _Powtoken.Contract.ValidProofs(&_Powtoken.CallOpts)
}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Powtoken *PowtokenCallerSession) ValidProofs() (*big.Int, error) {
	return _Powtoken.Contract.ValidProofs(&_Powtoken.CallOpts)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Powtoken *PowtokenTransactor) GenerateChallenge(opts *bind.TransactOpts, nodeId string) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "generateChallenge", nodeId)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Powtoken *PowtokenSession) GenerateChallenge(nodeId string) (*types.Transaction, error) {
	return _Powtoken.Contract.GenerateChallenge(&_Powtoken.TransactOpts, nodeId)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Powtoken *PowtokenTransactorSession) GenerateChallenge(nodeId string) (*types.Transaction, error) {
	return _Powtoken.Contract.GenerateChallenge(&_Powtoken.TransactOpts, nodeId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Powtoken *PowtokenTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Powtoken *PowtokenSession) Initialize() (*types.Transaction, error) {
	return _Powtoken.Contract.Initialize(&_Powtoken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Powtoken *PowtokenTransactorSession) Initialize() (*types.Transaction, error) {
	return _Powtoken.Contract.Initialize(&_Powtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Powtoken *PowtokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Powtoken *PowtokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Powtoken.Contract.RenounceOwnership(&_Powtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Powtoken *PowtokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Powtoken.Contract.RenounceOwnership(&_Powtoken.TransactOpts)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Powtoken *PowtokenTransactor) SubmitWork(opts *bind.TransactOpts, nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "submitWork", nonce, nodeId)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Powtoken *PowtokenSession) SubmitWork(nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Powtoken.Contract.SubmitWork(&_Powtoken.TransactOpts, nonce, nodeId)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Powtoken *PowtokenTransactorSession) SubmitWork(nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Powtoken.Contract.SubmitWork(&_Powtoken.TransactOpts, nonce, nodeId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Powtoken *PowtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Powtoken *PowtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Powtoken.Contract.TransferOwnership(&_Powtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Powtoken *PowtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Powtoken.Contract.TransferOwnership(&_Powtoken.TransactOpts, newOwner)
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Powtoken *PowtokenTransactor) TriggerNewPowRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Powtoken.contract.Transact(opts, "triggerNewPowRound")
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Powtoken *PowtokenSession) TriggerNewPowRound() (*types.Transaction, error) {
	return _Powtoken.Contract.TriggerNewPowRound(&_Powtoken.TransactOpts)
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Powtoken *PowtokenTransactorSession) TriggerNewPowRound() (*types.Transaction, error) {
	return _Powtoken.Contract.TriggerNewPowRound(&_Powtoken.TransactOpts)
}

// PowtokenGenerateChallengeIterator is returned from FilterGenerateChallenge and is used to iterate over the raw logs and unpacked data for GenerateChallenge events raised by the Powtoken contract.
type PowtokenGenerateChallengeIterator struct {
	Event *PowtokenGenerateChallenge // Event containing the contract specifics and raw log

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
func (it *PowtokenGenerateChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenGenerateChallenge)
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
		it.Event = new(PowtokenGenerateChallenge)
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
func (it *PowtokenGenerateChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenGenerateChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenGenerateChallenge represents a GenerateChallenge event raised by the Powtoken contract.
type PowtokenGenerateChallenge struct {
	Challenge [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenerateChallenge is a free log retrieval operation binding the contract event 0x2d3b66c3b514494f1330d833a6e605138788b0a382e7a282071485298300d770.
//
// Solidity: event GenerateChallenge(bytes32 challenge)
func (_Powtoken *PowtokenFilterer) FilterGenerateChallenge(opts *bind.FilterOpts) (*PowtokenGenerateChallengeIterator, error) {

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "GenerateChallenge")
	if err != nil {
		return nil, err
	}
	return &PowtokenGenerateChallengeIterator{contract: _Powtoken.contract, event: "GenerateChallenge", logs: logs, sub: sub}, nil
}

// WatchGenerateChallenge is a free log subscription operation binding the contract event 0x2d3b66c3b514494f1330d833a6e605138788b0a382e7a282071485298300d770.
//
// Solidity: event GenerateChallenge(bytes32 challenge)
func (_Powtoken *PowtokenFilterer) WatchGenerateChallenge(opts *bind.WatchOpts, sink chan<- *PowtokenGenerateChallenge) (event.Subscription, error) {

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "GenerateChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenGenerateChallenge)
				if err := _Powtoken.contract.UnpackLog(event, "GenerateChallenge", log); err != nil {
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

// ParseGenerateChallenge is a log parse operation binding the contract event 0x2d3b66c3b514494f1330d833a6e605138788b0a382e7a282071485298300d770.
//
// Solidity: event GenerateChallenge(bytes32 challenge)
func (_Powtoken *PowtokenFilterer) ParseGenerateChallenge(log types.Log) (*PowtokenGenerateChallenge, error) {
	event := new(PowtokenGenerateChallenge)
	if err := _Powtoken.contract.UnpackLog(event, "GenerateChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowtokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Powtoken contract.
type PowtokenInitializedIterator struct {
	Event *PowtokenInitialized // Event containing the contract specifics and raw log

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
func (it *PowtokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenInitialized)
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
		it.Event = new(PowtokenInitialized)
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
func (it *PowtokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenInitialized represents a Initialized event raised by the Powtoken contract.
type PowtokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Powtoken *PowtokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*PowtokenInitializedIterator, error) {

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PowtokenInitializedIterator{contract: _Powtoken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Powtoken *PowtokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PowtokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenInitialized)
				if err := _Powtoken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Powtoken *PowtokenFilterer) ParseInitialized(log types.Log) (*PowtokenInitialized, error) {
	event := new(PowtokenInitialized)
	if err := _Powtoken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowtokenNewPostRoundIterator is returned from FilterNewPostRound and is used to iterate over the raw logs and unpacked data for NewPostRound events raised by the Powtoken contract.
type PowtokenNewPostRoundIterator struct {
	Event *PowtokenNewPostRound // Event containing the contract specifics and raw log

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
func (it *PowtokenNewPostRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenNewPostRound)
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
		it.Event = new(PowtokenNewPostRound)
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
func (it *PowtokenNewPostRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenNewPostRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenNewPostRound represents a NewPostRound event raised by the Powtoken contract.
type PowtokenNewPostRound struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewPostRound is a free log retrieval operation binding the contract event 0x6828ec9c4d2d50ad73ca05fd4575a57180db77534ab19a16e9fbd0e30393d205.
//
// Solidity: event NewPostRound()
func (_Powtoken *PowtokenFilterer) FilterNewPostRound(opts *bind.FilterOpts) (*PowtokenNewPostRoundIterator, error) {

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "NewPostRound")
	if err != nil {
		return nil, err
	}
	return &PowtokenNewPostRoundIterator{contract: _Powtoken.contract, event: "NewPostRound", logs: logs, sub: sub}, nil
}

// WatchNewPostRound is a free log subscription operation binding the contract event 0x6828ec9c4d2d50ad73ca05fd4575a57180db77534ab19a16e9fbd0e30393d205.
//
// Solidity: event NewPostRound()
func (_Powtoken *PowtokenFilterer) WatchNewPostRound(opts *bind.WatchOpts, sink chan<- *PowtokenNewPostRound) (event.Subscription, error) {

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "NewPostRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenNewPostRound)
				if err := _Powtoken.contract.UnpackLog(event, "NewPostRound", log); err != nil {
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

// ParseNewPostRound is a log parse operation binding the contract event 0x6828ec9c4d2d50ad73ca05fd4575a57180db77534ab19a16e9fbd0e30393d205.
//
// Solidity: event NewPostRound()
func (_Powtoken *PowtokenFilterer) ParseNewPostRound(log types.Log) (*PowtokenNewPostRound, error) {
	event := new(PowtokenNewPostRound)
	if err := _Powtoken.contract.UnpackLog(event, "NewPostRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowtokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Powtoken contract.
type PowtokenOwnershipTransferredIterator struct {
	Event *PowtokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PowtokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenOwnershipTransferred)
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
		it.Event = new(PowtokenOwnershipTransferred)
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
func (it *PowtokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenOwnershipTransferred represents a OwnershipTransferred event raised by the Powtoken contract.
type PowtokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Powtoken *PowtokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PowtokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PowtokenOwnershipTransferredIterator{contract: _Powtoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Powtoken *PowtokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PowtokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenOwnershipTransferred)
				if err := _Powtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Powtoken *PowtokenFilterer) ParseOwnershipTransferred(log types.Log) (*PowtokenOwnershipTransferred, error) {
	event := new(PowtokenOwnershipTransferred)
	if err := _Powtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowtokenValidPOWSubmittedIterator is returned from FilterValidPOWSubmitted and is used to iterate over the raw logs and unpacked data for ValidPOWSubmitted events raised by the Powtoken contract.
type PowtokenValidPOWSubmittedIterator struct {
	Event *PowtokenValidPOWSubmitted // Event containing the contract specifics and raw log

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
func (it *PowtokenValidPOWSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenValidPOWSubmitted)
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
		it.Event = new(PowtokenValidPOWSubmitted)
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
func (it *PowtokenValidPOWSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenValidPOWSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenValidPOWSubmitted represents a ValidPOWSubmitted event raised by the Powtoken contract.
type PowtokenValidPOWSubmitted struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidPOWSubmitted is a free log retrieval operation binding the contract event 0xab3368117a8c701589001ea9b7cd7edb08dbb5cd3edba3b58272d5f23dfd6549.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenFilterer) FilterValidPOWSubmitted(opts *bind.FilterOpts, walletAddress []common.Address) (*PowtokenValidPOWSubmittedIterator, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "ValidPOWSubmitted", walletAddressRule)
	if err != nil {
		return nil, err
	}
	return &PowtokenValidPOWSubmittedIterator{contract: _Powtoken.contract, event: "ValidPOWSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidPOWSubmitted is a free log subscription operation binding the contract event 0xab3368117a8c701589001ea9b7cd7edb08dbb5cd3edba3b58272d5f23dfd6549.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenFilterer) WatchValidPOWSubmitted(opts *bind.WatchOpts, sink chan<- *PowtokenValidPOWSubmitted, walletAddress []common.Address) (event.Subscription, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "ValidPOWSubmitted", walletAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenValidPOWSubmitted)
				if err := _Powtoken.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
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

// ParseValidPOWSubmitted is a log parse operation binding the contract event 0xab3368117a8c701589001ea9b7cd7edb08dbb5cd3edba3b58272d5f23dfd6549.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge)
func (_Powtoken *PowtokenFilterer) ParseValidPOWSubmitted(log types.Log) (*PowtokenValidPOWSubmitted, error) {
	event := new(PowtokenValidPOWSubmitted)
	if err := _Powtoken.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
