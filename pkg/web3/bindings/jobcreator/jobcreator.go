// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jobcreator

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

// JobcreatorMetaData contains all meta data concerning the Jobcreator contract.
var JobcreatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"calling_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextJobID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"resultsAdded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"runJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setRequiredDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611c8a806101205f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c8063a470295811610095578063d2a715c011610064578063d2a715c014610223578063f2fde38b14610241578063f3d3d4481461025d578063fb7cfdd714610279576100f3565b8063a4702958146101af578063c4d66de8146101b9578063c57380a2146101d5578063c75555fa146101f3576100f3565b80634c526c76116100d15780634c526c761461014d57806358e56db414610169578063715018a6146101875780638da5cb5b14610191576100f3565b806310fe9ae8146100f757806312b69a8e1461011557806326a4e8d214610131575b5f80fd5b6100ff610297565b60405161010c9190610dc2565b60405180910390f35b61012f600480360381019061012a9190610f5b565b6102bf565b005b61014b6004803603810190610146919061103d565b6102c5565b005b61016760048036038101906101629190611068565b6103df565b005b610171610434565b60405161017e91906110a2565b60405180910390f35b61018f61043d565b005b610199610450565b6040516101a69190610dc2565b60405180910390f35b6101b7610477565b005b6101d360048036038101906101ce919061103d565b61049b565b005b6101dd6105e3565b6040516101ea9190610dc2565b60405180910390f35b61020d6004803603810190610208919061119d565b61060b565b60405161021a91906110a2565b60405180910390f35b61022b610886565b60405161023891906110a2565b60405180910390f35b61025b6004803603810190610256919061103d565b61088c565b005b6102776004803603810190610272919061103d565b61090e565b005b610281610a16565b60405161028e91906110a2565b60405180910390f35b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b50505050565b6102cd610a1c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361033b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103329061127f565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103e7610a9a565b505f811161042a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610421906112e7565b60405180910390fd5b8060048190555050565b5f600454905090565b610445610a1c565b61044e5f610bc7565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61047f610a1c565b5f600160146101000a81548160ff021916908315150217905550565b5f600160169054906101000a900460ff161590508080156104cd575060018060159054906101000a900460ff1660ff16105b806104fb57506104dc30610c88565b1580156104fa575060018060159054906101000a900460ff1660ff16145b5b61053a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053190611375565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156105765760018060166101000a81548160ff0219169083151502179055505b61057f826102c5565b5f60058190555080156105df575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516105d691906113e1565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60045460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e846106556105e3565b6040518363ffffffff1660e01b81526004016106729291906113fa565b602060405180830381865afa15801561068d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106b19190611435565b10156106f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e9906114aa565b60405180910390fd5b600160055461070191906114f5565b6005819055506040518060a0016040528060055481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481525060065f60055481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190816108189190611723565b506080820151816004019080519060200190610835929190610cb1565b509050507faa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a60055433848787604051610872959493929190611955565b60405180910390a160055490509392505050565b60055481565b610894610a1c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f990611a24565b60405180910390fd5b61090b81610bc7565b50565b610916610a1c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097b90611ab2565b60405180910390fd5b600160149054906101000a900460ff166109d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ca90611b40565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b610a24610caa565b73ffffffffffffffffffffffffffffffffffffffff16610a42610450565b73ffffffffffffffffffffffffffffffffffffffff1614610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f90611ba8565b60405180910390fd5b565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610b2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2190611ab2565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b6a610caa565b73ffffffffffffffffffffffffffffffffffffffff1614610bc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb790611c36565b60405180910390fd5b6001905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b828054828255905f5260205f20908101928215610cf7579160200282015b82811115610cf6578251829081610ce69190611723565b5091602001919060010190610ccf565b5b509050610d049190610d08565b5090565b5b80821115610d27575f8181610d1e9190610d2b565b50600101610d09565b5090565b508054610d379061155f565b5f825580601f10610d485750610d65565b601f0160209004905f5260205f2090810190610d649190610d68565b5b50565b5b80821115610d7f575f815f905550600101610d69565b5090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610dac82610d83565b9050919050565b610dbc81610da2565b82525050565b5f602082019050610dd55f830184610db3565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e3a82610df4565b810181811067ffffffffffffffff82111715610e5957610e58610e04565b5b80604052505050565b5f610e6b610ddb565b9050610e778282610e31565b919050565b5f67ffffffffffffffff821115610e9657610e95610e04565b5b610e9f82610df4565b9050602081019050919050565b828183375f83830152505050565b5f610ecc610ec784610e7c565b610e62565b905082815260208101848484011115610ee857610ee7610df0565b5b610ef3848285610eac565b509392505050565b5f82601f830112610f0f57610f0e610dec565b5b8135610f1f848260208601610eba565b91505092915050565b5f819050919050565b610f3a81610f28565b8114610f44575f80fd5b50565b5f81359050610f5581610f31565b92915050565b5f805f8060808587031215610f7357610f72610de4565b5b5f85013567ffffffffffffffff811115610f9057610f8f610de8565b5b610f9c87828801610efb565b945050602085013567ffffffffffffffff811115610fbd57610fbc610de8565b5b610fc987828801610efb565b935050604085013567ffffffffffffffff811115610fea57610fe9610de8565b5b610ff687828801610efb565b925050606061100787828801610f47565b91505092959194509250565b61101c81610da2565b8114611026575f80fd5b50565b5f8135905061103781611013565b92915050565b5f6020828403121561105257611051610de4565b5b5f61105f84828501611029565b91505092915050565b5f6020828403121561107d5761107c610de4565b5b5f61108a84828501610f47565b91505092915050565b61109c81610f28565b82525050565b5f6020820190506110b55f830184611093565b92915050565b5f67ffffffffffffffff8211156110d5576110d4610e04565b5b602082029050602081019050919050565b5f80fd5b5f6110fc6110f7846110bb565b610e62565b9050808382526020820190506020840283018581111561111f5761111e6110e6565b5b835b8181101561116657803567ffffffffffffffff81111561114457611143610dec565b5b8086016111518982610efb565b85526020850194505050602081019050611121565b5050509392505050565b5f82601f83011261118457611183610dec565b5b81356111948482602086016110ea565b91505092915050565b5f805f606084860312156111b4576111b3610de4565b5b5f84013567ffffffffffffffff8111156111d1576111d0610de8565b5b6111dd86828701610efb565b935050602084013567ffffffffffffffff8111156111fe576111fd610de8565b5b61120a86828701611170565b925050604061121b86828701611029565b9150509250925092565b5f82825260208201905092915050565b7f546f6b656e2061646472657373000000000000000000000000000000000000005f82015250565b5f611269600d83611225565b915061127482611235565b602082019050919050565b5f6020820190508181035f8301526112968161125d565b9050919050565b7f4d696e206465706f7369740000000000000000000000000000000000000000005f82015250565b5f6112d1600b83611225565b91506112dc8261129d565b602082019050919050565b5f6020820190508181035f8301526112fe816112c5565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f61135f602e83611225565b915061136a82611305565b604082019050919050565b5f6020820190508181035f83015261138c81611353565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6113cb6113c66113c184611393565b6113a8565b61139c565b9050919050565b6113db816113b1565b82525050565b5f6020820190506113f45f8301846113d2565b92915050565b5f60408201905061140d5f830185610db3565b61141a6020830184610db3565b9392505050565b5f8151905061142f81610f31565b92915050565b5f6020828403121561144a57611449610de4565b5b5f61145784828501611421565b91505092915050565b7f546f6b656e20616c6c6f77616e6365206e6f7420656e6f7567680000000000005f82015250565b5f611494601a83611225565b915061149f82611460565b602082019050919050565b5f6020820190508181035f8301526114c181611488565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114ff82610f28565b915061150a83610f28565b9250828201905080821115611522576115216114c8565b5b92915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061157657607f821691505b60208210810361158957611588611532565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026115eb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826115b0565b6115f586836115b0565b95508019841693508086168417925050509392505050565b5f61162761162261161d84610f28565b6113a8565b610f28565b9050919050565b5f819050919050565b6116408361160d565b61165461164c8261162e565b8484546115bc565b825550505050565b5f90565b61166861165c565b611673818484611637565b505050565b5b818110156116965761168b5f82611660565b600181019050611679565b5050565b601f8211156116db576116ac8161158f565b6116b5846115a1565b810160208510156116c4578190505b6116d86116d0856115a1565b830182611678565b50505b505050565b5f82821c905092915050565b5f6116fb5f19846008026116e0565b1980831691505092915050565b5f61171383836116ec565b9150826002028217905092915050565b61172c82611528565b67ffffffffffffffff81111561174557611744610e04565b5b61174f825461155f565b61175a82828561169a565b5f60209050601f83116001811461178b575f8415611779578287015190505b6117838582611708565b8655506117ea565b601f1984166117998661158f565b5f5b828110156117c05784890151825560018201915060208501945060208101905061179b565b868310156117dd57848901516117d9601f8916826116ec565b8355505b6001600288020188555050505b505050505050565b5f5b8381101561180f5780820151818401526020810190506117f4565b5f8484015250505050565b5f61182482611528565b61182e8185611225565b935061183e8185602086016117f2565b61184781610df4565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f61189582611528565b61189f818561187b565b93506118af8185602086016117f2565b6118b881610df4565b840191505092915050565b5f6118ce838361188b565b905092915050565b5f602082019050919050565b5f6118ec82611852565b6118f6818561185c565b9350836020820285016119088561186c565b805f5b85811015611943578484038952815161192485826118c3565b945061192f836118d6565b925060208a0199505060018101905061190b565b50829750879550505050505092915050565b5f60a0820190506119685f830188611093565b6119756020830187610db3565b6119826040830186610db3565b8181036060830152611994818561181a565b905081810360808301526119a881846118e2565b90509695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611a0e602683611225565b9150611a19826119b4565b604082019050919050565b5f6020820190508181035f830152611a3b81611a02565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f611a9c603583611225565b9150611aa782611a42565b604082019050919050565b5f6020820190508181035f830152611ac981611a90565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611b2a603983611225565b9150611b3582611ad0565b604082019050919050565b5f6020820190508181035f830152611b5781611b1e565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611b92602083611225565b9150611b9d82611b5e565b602082019050919050565b5f6020820190508181035f830152611bbf81611b86565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f611c20603b83611225565b9150611c2b82611bc6565b604082019050919050565b5f6020820190508181035f830152611c4d81611c14565b905091905056fea2646970667358221220b6ad28bc8d70d5d3d643f60aeb87c204788e7e17b0c66e5f8887f86d692da00664736f6c63430008150033",
}

// JobcreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use JobcreatorMetaData.ABI instead.
var JobcreatorABI = JobcreatorMetaData.ABI

// JobcreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobcreatorMetaData.Bin instead.
var JobcreatorBin = JobcreatorMetaData.Bin

// DeployJobcreator deploys a new Ethereum contract, binding an instance of Jobcreator to it.
func DeployJobcreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jobcreator, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobcreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// Jobcreator is an auto generated Go binding around an Ethereum contract.
type Jobcreator struct {
	JobcreatorCaller     // Read-only binding to the contract
	JobcreatorTransactor // Write-only binding to the contract
	JobcreatorFilterer   // Log filterer for contract events
}

// JobcreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobcreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobcreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobcreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobcreatorSession struct {
	Contract     *Jobcreator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobcreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobcreatorCallerSession struct {
	Contract *JobcreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobcreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobcreatorTransactorSession struct {
	Contract     *JobcreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobcreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobcreatorRaw struct {
	Contract *Jobcreator // Generic contract binding to access the raw methods on
}

// JobcreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobcreatorCallerRaw struct {
	Contract *JobcreatorCaller // Generic read-only contract binding to access the raw methods on
}

// JobcreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobcreatorTransactorRaw struct {
	Contract *JobcreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobcreator creates a new instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreator(address common.Address, backend bind.ContractBackend) (*Jobcreator, error) {
	contract, err := bindJobcreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// NewJobcreatorCaller creates a new read-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorCaller(address common.Address, caller bind.ContractCaller) (*JobcreatorCaller, error) {
	contract, err := bindJobcreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorCaller{contract: contract}, nil
}

// NewJobcreatorTransactor creates a new write-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JobcreatorTransactor, error) {
	contract, err := bindJobcreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorTransactor{contract: contract}, nil
}

// NewJobcreatorFilterer creates a new log filterer instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JobcreatorFilterer, error) {
	contract, err := bindJobcreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobcreatorFilterer{contract: contract}, nil
}

// bindJobcreator binds a generic wrapper to an already deployed contract.
func bindJobcreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.JobcreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) GetRequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getRequiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) NextJobID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "nextJobID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// ResultsAdded is a paid mutator transaction binding the contract method 0x12b69a8e.
//
// Solidity: function resultsAdded(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Jobcreator *JobcreatorTransactor) ResultsAdded(opts *bind.TransactOpts, dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "resultsAdded", dealId, resultsId, dataId, instructionCount)
}

// ResultsAdded is a paid mutator transaction binding the contract method 0x12b69a8e.
//
// Solidity: function resultsAdded(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Jobcreator *JobcreatorSession) ResultsAdded(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.ResultsAdded(&_Jobcreator.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// ResultsAdded is a paid mutator transaction binding the contract method 0x12b69a8e.
//
// Solidity: function resultsAdded(string dealId, string resultsId, string dataId, uint256 instructionCount) returns()
func (_Jobcreator *JobcreatorTransactorSession) ResultsAdded(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.ResultsAdded(&_Jobcreator.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactor) RunJob(opts *bind.TransactOpts, module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "runJob", module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactor) SetRequiredDeposit(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setRequiredDeposit", cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// JobcreatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Jobcreator contract.
type JobcreatorInitializedIterator struct {
	Event *JobcreatorInitialized // Event containing the contract specifics and raw log

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
func (it *JobcreatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorInitialized)
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
		it.Event = new(JobcreatorInitialized)
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
func (it *JobcreatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorInitialized represents a Initialized event raised by the Jobcreator contract.
type JobcreatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobcreatorInitializedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobcreatorInitializedIterator{contract: _Jobcreator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobcreatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorInitialized)
				if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseInitialized(log types.Log) (*JobcreatorInitialized, error) {
	event := new(JobcreatorInitialized)
	if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the Jobcreator contract.
type JobcreatorJobAddedIterator struct {
	Event *JobcreatorJobAdded // Event containing the contract specifics and raw log

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
func (it *JobcreatorJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorJobAdded)
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
		it.Event = new(JobcreatorJobAdded)
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
func (it *JobcreatorJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorJobAdded represents a JobAdded event raised by the Jobcreator contract.
type JobcreatorJobAdded struct {
	Id              *big.Int
	CallingContract common.Address
	Payee           common.Address
	Module          string
	Inputs          []string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) FilterJobAdded(opts *bind.FilterOpts) (*JobcreatorJobAddedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &JobcreatorJobAddedIterator{contract: _Jobcreator.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *JobcreatorJobAdded) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorJobAdded)
				if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
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

// ParseJobAdded is a log parse operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) ParseJobAdded(log types.Log) (*JobcreatorJobAdded, error) {
	event := new(JobcreatorJobAdded)
	if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Jobcreator contract.
type JobcreatorOwnershipTransferredIterator struct {
	Event *JobcreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JobcreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorOwnershipTransferred)
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
		it.Event = new(JobcreatorOwnershipTransferred)
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
func (it *JobcreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorOwnershipTransferred represents a OwnershipTransferred event raised by the Jobcreator contract.
type JobcreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobcreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobcreatorOwnershipTransferredIterator{contract: _Jobcreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobcreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorOwnershipTransferred)
				if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseOwnershipTransferred(log types.Log) (*JobcreatorOwnershipTransferred, error) {
	event := new(JobcreatorOwnershipTransferred)
	if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
