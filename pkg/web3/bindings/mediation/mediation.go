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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"MediationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getMediator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"mediationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611974806101205f395ff3fe608060405234801561000f575f80fd5b50600436106100a7575f3560e01c80638da5cb5b1161006f5780638da5cb5b14610113578063a2bffa0014610131578063a470295814610161578063c57380a21461016b578063f2fde38b14610189578063f3d3d448146101a5576100a7565b806370bea207146100ab578063715018a6146100c757806380ffdfe0146100d15780638129fc1c146100ed578063824518aa146100f7575b5f80fd5b6100c560048036038101906100c0919061120f565b6101c1565b005b6100cf610343565b005b6100eb60048036038101906100e69190611256565b610356565b005b6100f561051c565b005b610111600480360381019061010c9190611256565b610653565b005b61011b610819565b60405161012891906112ac565b60405180910390f35b61014b60048036038101906101469190611256565b610840565b60405161015891906112ac565b60405180910390f35b610169610886565b005b6101736108aa565b60405161018091906112ac565b60405180910390f35b6101a3600480360381019061019e91906112c5565b6108d2565b005b6101bf60048036038101906101ba91906112c5565b610954565b005b6101c9610a5c565b505f8160200151606001515142835f01516040516020016101eb92919061137c565b604051602081830303815290604052805190602001205f1c61020d91906113d0565b90505f826020015160600151828151811061022b5761022a611400565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029a90611487565b60405180910390fd5b806002845f01516040516102b791906114a5565b90815260200160405180910390205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086835f0151826040516103369291906114f3565b60405180910390a1505050565b61034b610b89565b6103545f610c07565b565b5f73ffffffffffffffffffffffffffffffffffffffff1660028260405161037d91906114a5565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610401576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f890611487565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660028260405161042891906114a5565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a39061156b565b60405180910390fd5b6104b46108aa565b73ffffffffffffffffffffffffffffffffffffffff166380ffdfe0826040518263ffffffff1660e01b81526004016104ec9190611589565b5f604051808303815f87803b158015610503575f80fd5b505af1158015610515573d5f803e3d5ffd5b5050505050565b5f600160169054906101000a900460ff1615905080801561054e575060018060159054906101000a900460ff1660ff16105b8061057c575061055d30610cc8565b15801561057b575060018060159054906101000a900460ff1660ff16145b5b6105bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b290611619565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156105f75760018060166101000a81548160ff0219169083151502179055505b8015610650575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516106479190611685565b60405180910390a15b50565b5f73ffffffffffffffffffffffffffffffffffffffff1660028260405161067a91906114a5565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f590611487565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660028260405161072591906114a5565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a09061156b565b60405180910390fd5b6107b16108aa565b73ffffffffffffffffffffffffffffffffffffffff1663824518aa826040518263ffffffff1660e01b81526004016107e99190611589565b5f604051808303815f87803b158015610800575f80fd5b505af1158015610812573d5f803e3d5ffd5b5050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60028260405161085191906114a5565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b61088e610b89565b5f600160146101000a81548160ff021916908315150217905550565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6108da610b89565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093f9061170e565b60405180910390fd5b61095181610c07565b50565b61095c610b89565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c19061179c565b60405180910390fd5b600160149054906101000a900460ff16610a19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a109061182a565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610aec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae39061179c565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b2c610cea565b73ffffffffffffffffffffffffffffffffffffffff1614610b82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b79906118b8565b60405180910390fd5b6001905090565b610b91610cea565b73ffffffffffffffffffffffffffffffffffffffff16610baf610819565b73ffffffffffffffffffffffffffffffffffffffff1614610c05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bfc90611920565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d4c82610d06565b810181811067ffffffffffffffff82111715610d6b57610d6a610d16565b5b80604052505050565b5f610d7d610cf1565b9050610d898282610d43565b919050565b5f80fd5b5f80fd5b5f80fd5b5f67ffffffffffffffff821115610db457610db3610d16565b5b610dbd82610d06565b9050602081019050919050565b828183375f83830152505050565b5f610dea610de584610d9a565b610d74565b905082815260208101848484011115610e0657610e05610d96565b5b610e11848285610dca565b509392505050565b5f82601f830112610e2d57610e2c610d92565b5b8135610e3d848260208601610dd8565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e6f82610e46565b9050919050565b610e7f81610e65565b8114610e89575f80fd5b50565b5f81359050610e9a81610e76565b92915050565b5f67ffffffffffffffff821115610eba57610eb9610d16565b5b602082029050602081019050919050565b5f80fd5b5f610ee1610edc84610ea0565b610d74565b90508083825260208201905060208402830185811115610f0457610f03610ecb565b5b835b81811015610f2d5780610f198882610e8c565b845260208401935050602081019050610f06565b5050509392505050565b5f82601f830112610f4b57610f4a610d92565b5b8135610f5b848260208601610ecf565b91505092915050565b5f60808284031215610f7957610f78610d02565b5b610f836080610d74565b90505f610f9284828501610e8c565b5f830152506020610fa584828501610e8c565b6020830152506040610fb984828501610e8c565b604083015250606082013567ffffffffffffffff811115610fdd57610fdc610d8e565b5b610fe984828501610f37565b60608301525092915050565b5f819050919050565b61100781610ff5565b8114611011575f80fd5b50565b5f8135905061102281610ffe565b92915050565b5f6040828403121561103d5761103c610d02565b5b6110476040610d74565b90505f61105684828501611014565b5f83015250602061106984828501611014565b60208301525092915050565b5f610100828403121561108b5761108a610d02565b5b6110956080610d74565b90505f6110a484828501611028565b5f8301525060406110b784828501611028565b60208301525060806110cb84828501611028565b60408301525060c06110df84828501611028565b60608301525092915050565b5f60808284031215611100576110ff610d02565b5b61110a6080610d74565b90505f61111984828501611014565b5f83015250602061112c84828501611014565b602083015250604061114084828501611014565b604083015250606061115484828501611014565b60608301525092915050565b5f6101c0828403121561117657611175610d02565b5b6111806080610d74565b90505f82013567ffffffffffffffff81111561119f5761119e610d8e565b5b6111ab84828501610e19565b5f83015250602082013567ffffffffffffffff8111156111ce576111cd610d8e565b5b6111da84828501610f64565b60208301525060406111ee84828501611075565b604083015250610140611203848285016110eb565b60608301525092915050565b5f6020828403121561122457611223610cfa565b5b5f82013567ffffffffffffffff81111561124157611240610cfe565b5b61124d84828501611160565b91505092915050565b5f6020828403121561126b5761126a610cfa565b5b5f82013567ffffffffffffffff81111561128857611287610cfe565b5b61129484828501610e19565b91505092915050565b6112a681610e65565b82525050565b5f6020820190506112bf5f83018461129d565b92915050565b5f602082840312156112da576112d9610cfa565b5b5f6112e784828501610e8c565b91505092915050565b5f819050919050565b61130a61130582610ff5565b6112f0565b82525050565b5f81519050919050565b5f81905092915050565b5f5b83811015611341578082015181840152602081019050611326565b5f8484015250505050565b5f61135682611310565b611360818561131a565b9350611370818560208601611324565b80840191505092915050565b5f61138782856112f9565b602082019150611397828461134c565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6113da82610ff5565b91506113e583610ff5565b9250826113f5576113f46113a3565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f6d65646961746f722063616e6e6f7420626520307830000000000000000000005f82015250565b5f61147160168361142d565b915061147c8261143d565b602082019050919050565b5f6020820190508181035f83015261149e81611465565b9050919050565b5f6114b0828461134c565b915081905092915050565b5f6114c582611310565b6114cf818561142d565b93506114df818560208601611324565b6114e881610d06565b840191505092915050565b5f6040820190508181035f83015261150b81856114bb565b905061151a602083018461129d565b9392505050565b7f74782e6f726967696e206d75737420626520746865206d65646961746f7200005f82015250565b5f611555601e8361142d565b915061156082611521565b602082019050919050565b5f6020820190508181035f83015261158281611549565b9050919050565b5f6020820190508181035f8301526115a181846114bb565b905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611603602e8361142d565b915061160e826115a9565b604082019050919050565b5f6020820190508181035f830152611630816115f7565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61166f61166a61166584611637565b61164c565b611640565b9050919050565b61167f81611655565b82525050565b5f6020820190506116985f830184611676565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6116f860268361142d565b91506117038261169e565b604082019050919050565b5f6020820190508181035f830152611725816116ec565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f61178660358361142d565b91506117918261172c565b604082019050919050565b5f6020820190508181035f8301526117b38161177a565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f61181460398361142d565b915061181f826117ba565b604082019050919050565b5f6020820190508181035f83015261184181611808565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6118a2603b8361142d565b91506118ad82611848565b604082019050919050565b5f6020820190508181035f8301526118cf81611896565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61190a60208361142d565b9150611915826118d6565b602082019050919050565b5f6020820190508181035f830152611937816118fe565b905091905056fea2646970667358221220ca4993e93f46ed61c2c9b52cbb71b5b58eb20c47bbfe4b0db81e68a611bb388a64736f6c63430008150033",
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

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationCaller) GetMediator(opts *bind.CallOpts, dealId string) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "getMediator", dealId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationSession) GetMediator(dealId string) (common.Address, error) {
	return _Mediation.Contract.GetMediator(&_Mediation.CallOpts, dealId)
}

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationCallerSession) GetMediator(dealId string) (common.Address, error) {
	return _Mediation.Contract.GetMediator(&_Mediation.CallOpts, dealId)
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
	DealId   string
	Mediator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMediationRequested is a free log retrieval operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
func (_Mediation *MediationFilterer) FilterMediationRequested(opts *bind.FilterOpts) (*MediationMediationRequestedIterator, error) {

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "MediationRequested")
	if err != nil {
		return nil, err
	}
	return &MediationMediationRequestedIterator{contract: _Mediation.contract, event: "MediationRequested", logs: logs, sub: sub}, nil
}

// WatchMediationRequested is a free log subscription operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
func (_Mediation *MediationFilterer) WatchMediationRequested(opts *bind.WatchOpts, sink chan<- *MediationMediationRequested) (event.Subscription, error) {

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "MediationRequested")
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

// ParseMediationRequested is a log parse operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
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
