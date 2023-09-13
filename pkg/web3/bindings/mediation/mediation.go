// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mediation

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

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId   string
	Members  SharedStructsDealMembers
	Timeouts SharedStructsDealTimeouts
	Pricing  SharedStructsDealPricing
}

// SharedStructsDealMembers is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealMembers struct {
	Solver           common.Address
	JobCreator       common.Address
	ResourceProvider common.Address
	Mediators        []common.Address
}

// SharedStructsDealPricing is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPricing struct {
	InstructionPrice          *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsDealTimeout is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeout struct {
	Timeout    *big.Int
	Collateral *big.Int
}

// SharedStructsDealTimeouts is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeouts struct {
	Agree          SharedStructsDealTimeout
	SubmitResults  SharedStructsDealTimeout
	JudgeResults   SharedStructsDealTimeout
	MediateResults SharedStructsDealTimeout
}

// MediationMetaData contains all meta data concerning the Mediation contract.
var MediationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"indexed\":true,\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"MediationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"mediationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611c03806101205f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c80638da5cb5b116100645780638da5cb5b14610108578063a470295814610126578063c57380a214610130578063f2fde38b1461014e578063f3d3d4481461016a5761009c565b806370bea207146100a0578063715018a6146100bc57806380ffdfe0146100c65780638129fc1c146100e2578063824518aa146100ec575b5f80fd5b6100ba60048036038101906100b591906111ab565b610186565b005b6100c4610325565b005b6100e060048036038101906100db91906111f2565b610338565b005b6100ea6104fe565b005b610106600480360381019061010191906111f2565b610635565b005b6101106107fb565b60405161011d9190611248565b60405180910390f35b61012e610822565b005b610138610846565b6040516101459190611248565b60405180910390f35b61016860048036038101906101639190611261565b61086e565b005b610184600480360381019061017f9190611261565b6108f0565b005b61018e6109f8565b505f8160200151606001515142835f01516040516020016101b0929190611318565b604051602081830303815290604052805190602001205f1c6101d2919061136c565b90505f82602001516060015182815181106101f0576101ef61139c565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f90611423565b60405180910390fd5b806002845f015160405161027c9190611441565b90815260200160405180910390205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826040516102d59190611762565b60405180910390208173ffffffffffffffffffffffffffffffffffffffff167f6ad2e830ba43f7d241eeddf5b4fa3adf34716f63d4a44db61b48bd8618909ff560405160405180910390a3505050565b61032d610b25565b6103365f610ba3565b565b5f73ffffffffffffffffffffffffffffffffffffffff1660028260405161035f9190611441565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036103e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103da90611423565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660028260405161040a9190611441565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461048e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610485906117c2565b60405180910390fd5b610496610846565b73ffffffffffffffffffffffffffffffffffffffff166380ffdfe0826040518263ffffffff1660e01b81526004016104ce9190611818565b5f604051808303815f87803b1580156104e5575f80fd5b505af11580156104f7573d5f803e3d5ffd5b5050505050565b5f600160169054906101000a900460ff16159050808015610530575060018060159054906101000a900460ff1660ff16105b8061055e575061053f30610c64565b15801561055d575060018060159054906101000a900460ff1660ff16145b5b61059d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610594906118a8565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156105d95760018060166101000a81548160ff0219169083151502179055505b8015610632575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516106299190611914565b60405180910390a15b50565b5f73ffffffffffffffffffffffffffffffffffffffff1660028260405161065c9190611441565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d790611423565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff166002826040516107079190611441565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461078b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610782906117c2565b60405180910390fd5b610793610846565b73ffffffffffffffffffffffffffffffffffffffff1663824518aa826040518263ffffffff1660e01b81526004016107cb9190611818565b5f604051808303815f87803b1580156107e2575f80fd5b505af11580156107f4573d5f803e3d5ffd5b5050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61082a610b25565b5f600160146101000a81548160ff021916908315150217905550565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610876610b25565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108db9061199d565b60405180910390fd5b6108ed81610ba3565b50565b6108f8610b25565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610966576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095d90611a2b565b60405180910390fd5b600160149054906101000a900460ff166109b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ac90611ab9565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610a88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7f90611a2b565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610ac8610c86565b73ffffffffffffffffffffffffffffffffffffffff1614610b1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1590611b47565b60405180910390fd5b6001905090565b610b2d610c86565b73ffffffffffffffffffffffffffffffffffffffff16610b4b6107fb565b73ffffffffffffffffffffffffffffffffffffffff1614610ba1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9890611baf565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ce882610ca2565b810181811067ffffffffffffffff82111715610d0757610d06610cb2565b5b80604052505050565b5f610d19610c8d565b9050610d258282610cdf565b919050565b5f80fd5b5f80fd5b5f80fd5b5f67ffffffffffffffff821115610d5057610d4f610cb2565b5b610d5982610ca2565b9050602081019050919050565b828183375f83830152505050565b5f610d86610d8184610d36565b610d10565b905082815260208101848484011115610da257610da1610d32565b5b610dad848285610d66565b509392505050565b5f82601f830112610dc957610dc8610d2e565b5b8135610dd9848260208601610d74565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e0b82610de2565b9050919050565b610e1b81610e01565b8114610e25575f80fd5b50565b5f81359050610e3681610e12565b92915050565b5f67ffffffffffffffff821115610e5657610e55610cb2565b5b602082029050602081019050919050565b5f80fd5b5f610e7d610e7884610e3c565b610d10565b90508083825260208201905060208402830185811115610ea057610e9f610e67565b5b835b81811015610ec95780610eb58882610e28565b845260208401935050602081019050610ea2565b5050509392505050565b5f82601f830112610ee757610ee6610d2e565b5b8135610ef7848260208601610e6b565b91505092915050565b5f60808284031215610f1557610f14610c9e565b5b610f1f6080610d10565b90505f610f2e84828501610e28565b5f830152506020610f4184828501610e28565b6020830152506040610f5584828501610e28565b604083015250606082013567ffffffffffffffff811115610f7957610f78610d2a565b5b610f8584828501610ed3565b60608301525092915050565b5f819050919050565b610fa381610f91565b8114610fad575f80fd5b50565b5f81359050610fbe81610f9a565b92915050565b5f60408284031215610fd957610fd8610c9e565b5b610fe36040610d10565b90505f610ff284828501610fb0565b5f83015250602061100584828501610fb0565b60208301525092915050565b5f610100828403121561102757611026610c9e565b5b6110316080610d10565b90505f61104084828501610fc4565b5f83015250604061105384828501610fc4565b602083015250608061106784828501610fc4565b60408301525060c061107b84828501610fc4565b60608301525092915050565b5f6080828403121561109c5761109b610c9e565b5b6110a66080610d10565b90505f6110b584828501610fb0565b5f8301525060206110c884828501610fb0565b60208301525060406110dc84828501610fb0565b60408301525060606110f084828501610fb0565b60608301525092915050565b5f6101c0828403121561111257611111610c9e565b5b61111c6080610d10565b90505f82013567ffffffffffffffff81111561113b5761113a610d2a565b5b61114784828501610db5565b5f83015250602082013567ffffffffffffffff81111561116a57611169610d2a565b5b61117684828501610f00565b602083015250604061118a84828501611011565b60408301525061014061119f84828501611087565b60608301525092915050565b5f602082840312156111c0576111bf610c96565b5b5f82013567ffffffffffffffff8111156111dd576111dc610c9a565b5b6111e9848285016110fc565b91505092915050565b5f6020828403121561120757611206610c96565b5b5f82013567ffffffffffffffff81111561122457611223610c9a565b5b61123084828501610db5565b91505092915050565b61124281610e01565b82525050565b5f60208201905061125b5f830184611239565b92915050565b5f6020828403121561127657611275610c96565b5b5f61128384828501610e28565b91505092915050565b5f819050919050565b6112a66112a182610f91565b61128c565b82525050565b5f81519050919050565b5f81905092915050565b5f5b838110156112dd5780820151818401526020810190506112c2565b5f8484015250505050565b5f6112f2826112ac565b6112fc81856112b6565b935061130c8185602086016112c0565b80840191505092915050565b5f6113238285611295565b60208201915061133382846112e8565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61137682610f91565b915061138183610f91565b9250826113915761139061133f565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f6d65646961746f722063616e6e6f7420626520307830000000000000000000005f82015250565b5f61140d6016836113c9565b9150611418826113d9565b602082019050919050565b5f6020820190508181035f83015261143a81611401565b9050919050565b5f61144c82846112e8565b915081905092915050565b5f81905092915050565b5f61146b826112ac565b6114758185611457565b93506114858185602086016112c0565b61148e81610ca2565b840191505092915050565b5f6114a48383611461565b905092915050565b6114b581610e01565b82525050565b5f6114c683836114ac565b60208301905092915050565b5f81519050919050565b5f81905092915050565b5f819050602082019050919050565b5f602082019050919050565b5f61150b826114d2565b61151581856114dc565b9350611520836114e6565b805f5b8381101561155057815161153788826114bb565b9750611542836114f5565b925050600181019050611523565b5085935050505092915050565b5f6115688383611501565b905092915050565b5f8083015f83015161158285826114bb565b945050602083015161159485826114bb565b94505060408301516115a685826114bb565b94505060608301516115b8858261155d565b9450508391505092915050565b5f6115d08383611570565b905092915050565b6115e181610f91565b82525050565b5f6115f283836115d8565b60208301905092915050565b5f82015f82015161160f84826115e7565b935050602082015161162184826115e7565b935050505050565b5f61163483836115fe565b60408301905092915050565b5f82015f8201516116518482611629565b93505060208201516116638482611629565b93505060408201516116758482611629565b93505060608201516116878482611629565b935050505050565b5f61169a8383611640565b6101008301905092915050565b5f82015f8201516116b884826115e7565b93505060208201516116ca84826115e7565b93505060408201516116dc84826115e7565b93505060608201516116ee84826115e7565b935050505050565b5f61170183836116a7565b60808301905092915050565b5f8083015f83015161171f8582611499565b945050602083015161173185826115c5565b9450506040830151611743858261168f565b945050606083015161175585826116f6565b9450508391505092915050565b5f61176d828461170d565b915081905092915050565b7f74782e6f726967696e206d75737420626520746865206d65646961746f7200005f82015250565b5f6117ac601e836113c9565b91506117b782611778565b602082019050919050565b5f6020820190508181035f8301526117d9816117a0565b9050919050565b5f6117ea826112ac565b6117f481856113c9565b93506118048185602086016112c0565b61180d81610ca2565b840191505092915050565b5f6020820190508181035f83015261183081846117e0565b905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611892602e836113c9565b915061189d82611838565b604082019050919050565b5f6020820190508181035f8301526118bf81611886565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6118fe6118f96118f4846118c6565b6118db565b6118cf565b9050919050565b61190e816118e4565b82525050565b5f6020820190506119275f830184611905565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6119876026836113c9565b91506119928261192d565b604082019050919050565b5f6020820190508181035f8301526119b48161197b565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f611a156035836113c9565b9150611a20826119bb565b604082019050919050565b5f6020820190508181035f830152611a4281611a09565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611aa36039836113c9565b9150611aae82611a49565b604082019050919050565b5f6020820190508181035f830152611ad081611a97565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f611b31603b836113c9565b9150611b3c82611ad7565b604082019050919050565b5f6020820190508181035f830152611b5e81611b25565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611b996020836113c9565b9150611ba482611b65565b602082019050919050565b5f6020820190508181035f830152611bc681611b8d565b905091905056fea2646970667358221220112956ad61a492ed4f1296897e6c98d6c3dc846dde44ed5ae3745c51359280c764736f6c63430008150033",
}

// MediationABI is the input ABI used to generate the binding from.
// Deprecated: Use MediationMetaData.ABI instead.
var MediationABI = MediationMetaData.ABI

// MediationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MediationMetaData.Bin instead.
var MediationBin = MediationMetaData.Bin

// DeployMediation deploys a new Ethereum contract, binding an instance of Mediation to it.
func DeployMediation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mediation, error) {
	parsed, err := MediationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MediationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mediation{MediationCaller: MediationCaller{contract: contract}, MediationTransactor: MediationTransactor{contract: contract}, MediationFilterer: MediationFilterer{contract: contract}}, nil
}

// Mediation is an auto generated Go binding around an Ethereum contract.
type Mediation struct {
	MediationCaller     // Read-only binding to the contract
	MediationTransactor // Write-only binding to the contract
	MediationFilterer   // Log filterer for contract events
}

// MediationCaller is an auto generated read-only Go binding around an Ethereum contract.
type MediationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MediationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MediationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MediationSession struct {
	Contract     *Mediation        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MediationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MediationCallerSession struct {
	Contract *MediationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MediationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MediationTransactorSession struct {
	Contract     *MediationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MediationRaw is an auto generated low-level Go binding around an Ethereum contract.
type MediationRaw struct {
	Contract *Mediation // Generic contract binding to access the raw methods on
}

// MediationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MediationCallerRaw struct {
	Contract *MediationCaller // Generic read-only contract binding to access the raw methods on
}

// MediationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MediationTransactorRaw struct {
	Contract *MediationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMediation creates a new instance of Mediation, bound to a specific deployed contract.
func NewMediation(address common.Address, backend bind.ContractBackend) (*Mediation, error) {
	contract, err := bindMediation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mediation{MediationCaller: MediationCaller{contract: contract}, MediationTransactor: MediationTransactor{contract: contract}, MediationFilterer: MediationFilterer{contract: contract}}, nil
}

// NewMediationCaller creates a new read-only instance of Mediation, bound to a specific deployed contract.
func NewMediationCaller(address common.Address, caller bind.ContractCaller) (*MediationCaller, error) {
	contract, err := bindMediation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MediationCaller{contract: contract}, nil
}

// NewMediationTransactor creates a new write-only instance of Mediation, bound to a specific deployed contract.
func NewMediationTransactor(address common.Address, transactor bind.ContractTransactor) (*MediationTransactor, error) {
	contract, err := bindMediation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MediationTransactor{contract: contract}, nil
}

// NewMediationFilterer creates a new log filterer instance of Mediation, bound to a specific deployed contract.
func NewMediationFilterer(address common.Address, filterer bind.ContractFilterer) (*MediationFilterer, error) {
	contract, err := bindMediation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MediationFilterer{contract: contract}, nil
}

// bindMediation binds a generic wrapper to an already deployed contract.
func bindMediation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MediationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediation *MediationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mediation.Contract.MediationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediation *MediationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.Contract.MediationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediation *MediationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediation.Contract.MediationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediation *MediationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mediation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediation *MediationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediation *MediationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediation.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationSession) GetControllerAddress() (common.Address, error) {
	return _Mediation.Contract.GetControllerAddress(&_Mediation.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationCallerSession) GetControllerAddress() (common.Address, error) {
	return _Mediation.Contract.GetControllerAddress(&_Mediation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationSession) Owner() (common.Address, error) {
	return _Mediation.Contract.Owner(&_Mediation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationCallerSession) Owner() (common.Address, error) {
	return _Mediation.Contract.Owner(&_Mediation.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Mediation.Contract.DisableChangeControllerAddress(&_Mediation.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Mediation.Contract.DisableChangeControllerAddress(&_Mediation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationSession) Initialize() (*types.Transaction, error) {
	return _Mediation.Contract.Initialize(&_Mediation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationTransactorSession) Initialize() (*types.Transaction, error) {
	return _Mediation.Contract.Initialize(&_Mediation.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationTransactorSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationTransactorSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationTransactor) MediationRequest(opts *bind.TransactOpts, deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRequest", deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationSession) MediationRequest(deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRequest(&_Mediation.TransactOpts, deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationTransactorSession) MediationRequest(deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRequest(&_Mediation.TransactOpts, deal)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediation.Contract.RenounceOwnership(&_Mediation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediation.Contract.RenounceOwnership(&_Mediation.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.SetControllerAddress(&_Mediation.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.SetControllerAddress(&_Mediation.TransactOpts, _controllerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.TransferOwnership(&_Mediation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.TransferOwnership(&_Mediation.TransactOpts, newOwner)
}

// MediationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mediation contract.
type MediationInitializedIterator struct {
	Event *MediationInitialized // Event containing the contract specifics and raw log

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
func (it *MediationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationInitialized)
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
		it.Event = new(MediationInitialized)
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
func (it *MediationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationInitialized represents a Initialized event raised by the Mediation contract.
type MediationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mediation *MediationFilterer) FilterInitialized(opts *bind.FilterOpts) (*MediationInitializedIterator, error) {

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MediationInitializedIterator{contract: _Mediation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mediation *MediationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MediationInitialized) (event.Subscription, error) {

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationInitialized)
				if err := _Mediation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Mediation *MediationFilterer) ParseInitialized(log types.Log) (*MediationInitialized, error) {
	event := new(MediationInitialized)
	if err := _Mediation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MediationMediationRequestedIterator is returned from FilterMediationRequested and is used to iterate over the raw logs and unpacked data for MediationRequested events raised by the Mediation contract.
type MediationMediationRequestedIterator struct {
	Event *MediationMediationRequested // Event containing the contract specifics and raw log

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
func (it *MediationMediationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationMediationRequested)
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
		it.Event = new(MediationMediationRequested)
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
func (it *MediationMediationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationMediationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationMediationRequested represents a MediationRequested event raised by the Mediation contract.
type MediationMediationRequested struct {
	Mediator common.Address
	Deal     SharedStructsDeal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMediationRequested is a free log retrieval operation binding the contract event 0x6ad2e830ba43f7d241eeddf5b4fa3adf34716f63d4a44db61b48bd8618909ff5.
//
// Solidity: event MediationRequested(address indexed mediator, (string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
func (_Mediation *MediationFilterer) FilterMediationRequested(opts *bind.FilterOpts, mediator []common.Address, deal []SharedStructsDeal) (*MediationMediationRequestedIterator, error) {

	var mediatorRule []interface{}
	for _, mediatorItem := range mediator {
		mediatorRule = append(mediatorRule, mediatorItem)
	}
	var dealRule []interface{}
	for _, dealItem := range deal {
		dealRule = append(dealRule, dealItem)
	}

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "MediationRequested", mediatorRule, dealRule)
	if err != nil {
		return nil, err
	}
	return &MediationMediationRequestedIterator{contract: _Mediation.contract, event: "MediationRequested", logs: logs, sub: sub}, nil
}

// WatchMediationRequested is a free log subscription operation binding the contract event 0x6ad2e830ba43f7d241eeddf5b4fa3adf34716f63d4a44db61b48bd8618909ff5.
//
// Solidity: event MediationRequested(address indexed mediator, (string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
func (_Mediation *MediationFilterer) WatchMediationRequested(opts *bind.WatchOpts, sink chan<- *MediationMediationRequested, mediator []common.Address, deal []SharedStructsDeal) (event.Subscription, error) {

	var mediatorRule []interface{}
	for _, mediatorItem := range mediator {
		mediatorRule = append(mediatorRule, mediatorItem)
	}
	var dealRule []interface{}
	for _, dealItem := range deal {
		dealRule = append(dealRule, dealItem)
	}

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "MediationRequested", mediatorRule, dealRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationMediationRequested)
				if err := _Mediation.contract.UnpackLog(event, "MediationRequested", log); err != nil {
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

// ParseMediationRequested is a log parse operation binding the contract event 0x6ad2e830ba43f7d241eeddf5b4fa3adf34716f63d4a44db61b48bd8618909ff5.
//
// Solidity: event MediationRequested(address indexed mediator, (string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
func (_Mediation *MediationFilterer) ParseMediationRequested(log types.Log) (*MediationMediationRequested, error) {
	event := new(MediationMediationRequested)
	if err := _Mediation.contract.UnpackLog(event, "MediationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MediationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mediation contract.
type MediationOwnershipTransferredIterator struct {
	Event *MediationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MediationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationOwnershipTransferred)
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
		it.Event = new(MediationOwnershipTransferred)
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
func (it *MediationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationOwnershipTransferred represents a OwnershipTransferred event raised by the Mediation contract.
type MediationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mediation *MediationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MediationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MediationOwnershipTransferredIterator{contract: _Mediation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mediation *MediationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MediationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationOwnershipTransferred)
				if err := _Mediation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mediation *MediationFilterer) ParseOwnershipTransferred(log types.Log) (*MediationOwnershipTransferred, error) {
	event := new(MediationOwnershipTransferred)
	if err := _Mediation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
