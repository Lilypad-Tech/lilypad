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
	DealId   *big.Int
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"indexed\":true,\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"MediationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"mediationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6119d7806101205f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c8063a470295811610064578063a47029581461010a578063c57380a214610114578063edaa017814610132578063f2fde38b1461014e578063f3d3d4481461016a5761009c565b806323a9a862146100a05780635dd049fd146100bc578063715018a6146100d85780638129fc1c146100e25780638da5cb5b146100ec575b5f80fd5b6100ba60048036038101906100b59190610c90565b610186565b005b6100d660048036038101906100d19190610c90565b610332565b005b6100e06104de565b005b6100ea6104f1565b005b6100f4610628565b6040516101019190610cfa565b60405180910390f35b61011261064f565b005b61011c610673565b6040516101299190610cfa565b60405180910390f35b61014c600480360381019061014791906110f1565b61069b565b005b61016860048036038101906101639190611138565b61082d565b005b610184600480360381019061017f9190611138565b6108af565b005b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021b906111bd565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b990611225565b60405180910390fd5b6102ca610673565b73ffffffffffffffffffffffffffffffffffffffff166323a9a862826040518263ffffffff1660e01b81526004016103029190611252565b5f604051808303815f87803b158015610319575f80fd5b505af115801561032b573d5f803e3d5ffd5b5050505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036103d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c7906111bd565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461046e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046590611225565b60405180910390fd5b610476610673565b73ffffffffffffffffffffffffffffffffffffffff16635dd049fd826040518263ffffffff1660e01b81526004016104ae9190611252565b5f604051808303815f87803b1580156104c5575f80fd5b505af11580156104d7573d5f803e3d5ffd5b5050505050565b6104e66109b7565b6104ef5f610a35565b565b5f600160169054906101000a900460ff16159050808015610523575060018060159054906101000a900460ff1660ff16105b80610551575061053230610af6565b158015610550575060018060159054906101000a900460ff1660ff16145b5b610590576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610587906112db565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156105cc5760018060166101000a81548160ff0219169083151502179055505b8015610625575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161061c9190611347565b60405180910390a15b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106576109b7565b5f600160146101000a81548160ff021916908315150217905550565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106a3610b18565b505f8160200151606001515142835f01516040516020016106c5929190611380565b604051602081830303815290604052805190602001205f1c6106e791906113d8565b90505f826020015160600151828151811061070557610704611408565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361077d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610774906111bd565b60405180910390fd5b8060025f855f015181526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826040516107dd91906116eb565b60405180910390208173ffffffffffffffffffffffffffffffffffffffff167f80c2428b46fb5f98fa27c965978150c1c83f697369a6ba488396b31e859d396260405160405180910390a3505050565b6108356109b7565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089a90611771565b60405180910390fd5b6108ac81610a35565b50565b6108b76109b7565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610925576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091c906117ff565b60405180910390fd5b600160149054906101000a900460ff16610974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096b9061188d565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6109bf610c45565b73ffffffffffffffffffffffffffffffffffffffff166109dd610628565b73ffffffffffffffffffffffffffffffffffffffff1614610a33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2a906118f5565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610ba8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9f906117ff565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610be8610c45565b73ffffffffffffffffffffffffffffffffffffffff1614610c3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3590611983565b60405180910390fd5b6001905090565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610c6f81610c5d565b8114610c79575f80fd5b50565b5f81359050610c8a81610c66565b92915050565b5f60208284031215610ca557610ca4610c55565b5b5f610cb284828501610c7c565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ce482610cbb565b9050919050565b610cf481610cda565b82525050565b5f602082019050610d0d5f830184610ceb565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d5d82610d17565b810181811067ffffffffffffffff82111715610d7c57610d7b610d27565b5b80604052505050565b5f610d8e610c4c565b9050610d9a8282610d54565b919050565b5f80fd5b610dac81610cda565b8114610db6575f80fd5b50565b5f81359050610dc781610da3565b92915050565b5f80fd5b5f67ffffffffffffffff821115610deb57610dea610d27565b5b602082029050602081019050919050565b5f80fd5b5f610e12610e0d84610dd1565b610d85565b90508083825260208201905060208402830185811115610e3557610e34610dfc565b5b835b81811015610e5e5780610e4a8882610db9565b845260208401935050602081019050610e37565b5050509392505050565b5f82601f830112610e7c57610e7b610dcd565b5b8135610e8c848260208601610e00565b91505092915050565b5f60808284031215610eaa57610ea9610d13565b5b610eb46080610d85565b90505f610ec384828501610db9565b5f830152506020610ed684828501610db9565b6020830152506040610eea84828501610db9565b604083015250606082013567ffffffffffffffff811115610f0e57610f0d610d9f565b5b610f1a84828501610e68565b60608301525092915050565b5f60408284031215610f3b57610f3a610d13565b5b610f456040610d85565b90505f610f5484828501610c7c565b5f830152506020610f6784828501610c7c565b60208301525092915050565b5f6101008284031215610f8957610f88610d13565b5b610f936080610d85565b90505f610fa284828501610f26565b5f830152506040610fb584828501610f26565b6020830152506080610fc984828501610f26565b60408301525060c0610fdd84828501610f26565b60608301525092915050565b5f60808284031215610ffe57610ffd610d13565b5b6110086080610d85565b90505f61101784828501610c7c565b5f83015250602061102a84828501610c7c565b602083015250604061103e84828501610c7c565b604083015250606061105284828501610c7c565b60608301525092915050565b5f6101c0828403121561107457611073610d13565b5b61107e6080610d85565b90505f61108d84828501610c7c565b5f83015250602082013567ffffffffffffffff8111156110b0576110af610d9f565b5b6110bc84828501610e95565b60208301525060406110d084828501610f73565b6040830152506101406110e584828501610fe9565b60608301525092915050565b5f6020828403121561110657611105610c55565b5b5f82013567ffffffffffffffff81111561112357611122610c59565b5b61112f8482850161105e565b91505092915050565b5f6020828403121561114d5761114c610c55565b5b5f61115a84828501610db9565b91505092915050565b5f82825260208201905092915050565b7f6d65646961746f722063616e6e6f7420626520307830000000000000000000005f82015250565b5f6111a7601683611163565b91506111b282611173565b602082019050919050565b5f6020820190508181035f8301526111d48161119b565b9050919050565b7f74782e6f726967696e206d75737420626520746865206d65646961746f7200005f82015250565b5f61120f601e83611163565b915061121a826111db565b602082019050919050565b5f6020820190508181035f83015261123c81611203565b9050919050565b61124c81610c5d565b82525050565b5f6020820190506112655f830184611243565b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6112c5602e83611163565b91506112d08261126b565b604082019050919050565b5f6020820190508181035f8301526112f2816112b9565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61133161132c611327846112f9565b61130e565b611302565b9050919050565b61134181611317565b82525050565b5f60208201905061135a5f830184611338565b92915050565b5f819050919050565b61137a61137582610c5d565b611360565b82525050565b5f61138b8285611369565b60208201915061139b8284611369565b6020820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6113e282610c5d565b91506113ed83610c5d565b9250826113fd576113fc6113ab565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61143e81610c5d565b82525050565b5f61144f8383611435565b60208301905092915050565b61146481610cda565b82525050565b5f611475838361145b565b60208301905092915050565b5f81519050919050565b5f81905092915050565b5f819050602082019050919050565b5f602082019050919050565b5f6114ba82611481565b6114c4818561148b565b93506114cf83611495565b805f5b838110156114ff5781516114e6888261146a565b97506114f1836114a4565b9250506001810190506114d2565b5085935050505092915050565b5f61151783836114b0565b905092915050565b5f8083015f830151611531858261146a565b9450506020830151611543858261146a565b9450506040830151611555858261146a565b9450506060830151611567858261150c565b9450508391505092915050565b5f61157f838361151f565b905092915050565b5f82015f8201516115988482611444565b93505060208201516115aa8482611444565b935050505050565b5f6115bd8383611587565b60408301905092915050565b5f82015f8201516115da84826115b2565b93505060208201516115ec84826115b2565b93505060408201516115fe84826115b2565b935050606082015161161084826115b2565b935050505050565b5f61162383836115c9565b6101008301905092915050565b5f82015f8201516116418482611444565b93505060208201516116538482611444565b93505060408201516116658482611444565b93505060608201516116778482611444565b935050505050565b5f61168a8383611630565b60808301905092915050565b5f8083015f8301516116a88582611444565b94505060208301516116ba8582611574565b94505060408301516116cc8582611618565b94505060608301516116de858261167f565b9450508391505092915050565b5f6116f68284611696565b915081905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61175b602683611163565b915061176682611701565b604082019050919050565b5f6020820190508181035f8301526117888161174f565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6117e9603583611163565b91506117f48261178f565b604082019050919050565b5f6020820190508181035f830152611816816117dd565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611877603983611163565b91506118828261181d565b604082019050919050565b5f6020820190508181035f8301526118a48161186b565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6118df602083611163565b91506118ea826118ab565b602082019050919050565b5f6020820190508181035f83015261190c816118d3565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f61196d603b83611163565b915061197882611913565b604082019050919050565b5f6020820190508181035f83015261199a81611961565b905091905056fea26469706673582212205bccb23e2472615e9b6ea1c7e6658cd9a0306514d8da67f1e52fcb043d7f346364736f6c63430008150033",
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

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Mediation *MediationTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Mediation *MediationSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Mediation *MediationTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Mediation *MediationTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Mediation *MediationSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Mediation *MediationTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRequest is a paid mutator transaction binding the contract method 0xedaa0178.
//
// Solidity: function mediationRequest((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationTransactor) MediationRequest(opts *bind.TransactOpts, deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRequest", deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0xedaa0178.
//
// Solidity: function mediationRequest((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationSession) MediationRequest(deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRequest(&_Mediation.TransactOpts, deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0xedaa0178.
//
// Solidity: function mediationRequest((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
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

// FilterMediationRequested is a free log retrieval operation binding the contract event 0x80c2428b46fb5f98fa27c965978150c1c83f697369a6ba488396b31e859d3962.
//
// Solidity: event MediationRequested(address indexed mediator, (uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
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

// WatchMediationRequested is a free log subscription operation binding the contract event 0x80c2428b46fb5f98fa27c965978150c1c83f697369a6ba488396b31e859d3962.
//
// Solidity: event MediationRequested(address indexed mediator, (uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
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

// ParseMediationRequested is a log parse operation binding the contract event 0x80c2428b46fb5f98fa27c965978150c1c83f697369a6ba488396b31e859d3962.
//
// Solidity: event MediationRequested(address indexed mediator, (uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) indexed deal)
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
