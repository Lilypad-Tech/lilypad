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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficult\",\"type\":\"uint256\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPowRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHashrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611a2e806101065f395ff3fe608060405234801561000f575f80fd5b50600436106100b2575f3560e01c80638da5cb5b1161006f5780638da5cb5b1461016e578063ab0989451461018c578063b28d87ea146101a8578063b681f2fd146101c6578063da8accf9146101d0578063f2fde38b146101ec576100b2565b80633181c380146100b65780634bbe05e4146100d45780636189f3ac14610109578063715018a61461013c57806378e97925146101465780638129fc1c14610164575b5f80fd5b6100be610208565b6040516100cb9190610d28565b60405180910390f35b6100ee60048036038101906100e99190610dcd565b610270565b60405161010096959493929190610ebc565b60405180910390f35b610123600480360381019061011e9190610f22565b610367565b6040516101339493929190610f4d565b60405180910390f35b610144610419565b005b61014e61042c565b60405161015b9190610d28565b60405180910390f35b61016c610432565b005b610176610567565b6040516101839190610f97565b60405180910390f35b6101a660048036038101906101a19190611011565b61058e565b005b6101b061070e565b6040516101bd9190610d28565b60405180910390f35b6101ce610714565b005b6101ea60048036038101906101e5919061105c565b61074a565b005b61020660048036038101906102019190610f22565b610b26565b005b5f806004544261021891906110e6565b90505f6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6102499190611146565b90505f8160035461025a9190611176565b905082816102689190611146565b935050505090565b6001602052815f5260405f208181548110610289575f80fd5b905f5260205f2090600602015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546102ce906111e4565b80601f01602080910402602001604051908101604052809291908181526020018280546102fa906111e4565b80156103455780601f1061031c57610100808354040283529160200191610345565b820191905f5260205f20905b81548152906001019060200180831161032857829003601f168201915b5050505050908060020154908060030154908060040154908060050154905086565b6002602052805f5260405f205f91509050805f015490806001015490806002018054610392906111e4565b80601f01602080910402602001604051908101604052809291908181526020018280546103be906111e4565b80156104095780601f106103e057610100808354040283529160200191610409565b820191905f5260205f20905b8154815290600101906020018083116103ec57829003601f168201915b5050505050908060030154905084565b610421610ba8565b61042a5f610c26565b565b60045481565b5f8060159054906101000a900460ff16159050808015610463575060015f60149054906101000a900460ff1660ff16105b80610491575061047230610ce7565b158015610490575060015f60149054906101000a900460ff1660ff16145b5b6104d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c790611284565b60405180910390fd5b60015f60146101000a81548160ff021916908360ff160217905550801561050c5760015f60156101000a81548160ff0219169083151502179055505b8015610564575f8060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161055b91906112f0565b60405180910390a15b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f423384846040516020016105a694939291906113aa565b6040516020818303038152906040528051906020012090505f6001446105cc91906113e4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6105f79190611146565b9050604051806080016040528083815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015560408201518160020190816106c191906115d8565b50606082015181600301559050507f496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef82826040516107009291906116a7565b60405180910390a150505050565b60035481565b61071c610ba8565b7f10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e60405160405180910390a1565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f8201548152602001600182015481526020016002820180546107b6906111e4565b80601f01602080910402602001604051908101604052809291908181526020018280546107e2906111e4565b801561082d5780601f106108045761010080835404028352916020019161082d565b820191905f5260205f20905b81548152906001019060200180831161081057829003601f168201915b5050505050815260200160038201548152505090505f816060015133858560405160200161085e94939291906113aa565b60405160208183030381529060405280519060200120905080825f0151106108bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b29061173e565b60405180910390fd5b5f81866040516020016108cf92919061177c565b6040516020818303038152906040528051906020012090508260200151815f1c1061092f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092690611817565b60405180910390fd5b60035f81548092919061094190611835565b91905055505f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001898152602001428152602001865f015181526020018660200151815250908060018154018082558091505060019003905f5260205f2090600602015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610a9391906115d8565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015550503373ffffffffffffffffffffffffffffffffffffffff167f5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c86087878a42895f01518a60200151604051610b15969594939291906118a8565b60405180910390a250505050505050565b610b2e610ba8565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9390611972565b60405180910390fd5b610ba581610c26565b50565b610bb0610d09565b73ffffffffffffffffffffffffffffffffffffffff16610bce610567565b73ffffffffffffffffffffffffffffffffffffffff1614610c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1b906119da565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f819050919050565b610d2281610d10565b82525050565b5f602082019050610d3b5f830184610d19565b92915050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d7282610d49565b9050919050565b610d8281610d68565b8114610d8c575f80fd5b50565b5f81359050610d9d81610d79565b92915050565b610dac81610d10565b8114610db6575f80fd5b50565b5f81359050610dc781610da3565b92915050565b5f8060408385031215610de357610de2610d41565b5b5f610df085828601610d8f565b9250506020610e0185828601610db9565b9150509250929050565b610e1481610d68565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e51578082015181840152602081019050610e36565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610e7682610e1a565b610e808185610e24565b9350610e90818560208601610e34565b610e9981610e5c565b840191505092915050565b5f819050919050565b610eb681610ea4565b82525050565b5f60c082019050610ecf5f830189610e0b565b8181036020830152610ee18188610e6c565b9050610ef06040830187610d19565b610efd6060830186610d19565b610f0a6080830185610ead565b610f1760a0830184610d19565b979650505050505050565b5f60208284031215610f3757610f36610d41565b5b5f610f4484828501610d8f565b91505092915050565b5f608082019050610f605f830187610ead565b610f6d6020830186610d19565b8181036040830152610f7f8185610e6c565b9050610f8e6060830184610d19565b95945050505050565b5f602082019050610faa5f830184610e0b565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112610fd157610fd0610fb0565b5b8235905067ffffffffffffffff811115610fee57610fed610fb4565b5b60208301915083600182028301111561100a57611009610fb8565b5b9250929050565b5f806020838503121561102757611026610d41565b5b5f83013567ffffffffffffffff81111561104457611043610d45565b5b61105085828601610fbc565b92509250509250929050565b5f805f6040848603121561107357611072610d41565b5b5f61108086828701610db9565b935050602084013567ffffffffffffffff8111156110a1576110a0610d45565b5b6110ad86828701610fbc565b92509250509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110f082610d10565b91506110fb83610d10565b9250828203905081811115611113576111126110b9565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61115082610d10565b915061115b83610d10565b92508261116b5761116a611119565b5b828204905092915050565b5f61118082610d10565b915061118b83610d10565b925082820261119981610d10565b915082820484148315176111b0576111af6110b9565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111fb57607f821691505b60208210810361120e5761120d6111b7565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f61126e602e83610e24565b915061127982611214565b604082019050919050565b5f6020820190508181035f83015261129b81611262565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6112da6112d56112d0846112a2565b6112b7565b6112ab565b9050919050565b6112ea816112c0565b82525050565b5f6020820190506113035f8301846112e1565b92915050565b5f819050919050565b61132361131e82610d10565b611309565b82525050565b5f8160601b9050919050565b5f61133f82611329565b9050919050565b5f61135082611335565b9050919050565b61136861136382610d68565b611346565b82525050565b5f81905092915050565b828183375f83830152505050565b5f611391838561136e565b935061139e838584611378565b82840190509392505050565b5f6113b58287611312565b6020820191506113c58286611357565b6014820191506113d6828486611386565b915081905095945050505050565b5f6113ee82610d10565b91506113f983610d10565b9250828201905080821115611411576114106110b9565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026114a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611465565b6114aa8683611465565b95508019841693508086168417925050509392505050565b5f6114dc6114d76114d284610d10565b6112b7565b610d10565b9050919050565b5f819050919050565b6114f5836114c2565b611509611501826114e3565b848454611471565b825550505050565b5f90565b61151d611511565b6115288184846114ec565b505050565b5b8181101561154b576115405f82611515565b60018101905061152e565b5050565b601f8211156115905761156181611444565b61156a84611456565b81016020851015611579578190505b61158d61158585611456565b83018261152d565b50505b505050565b5f82821c905092915050565b5f6115b05f1984600802611595565b1980831691505092915050565b5f6115c883836115a1565b9150826002028217905092915050565b6115e182610e1a565b67ffffffffffffffff8111156115fa576115f9611417565b5b61160482546111e4565b61160f82828561154f565b5f60209050601f831160018114611640575f841561162e578287015190505b61163885826115bd565b86555061169f565b601f19841661164e86611444565b5f5b8281101561167557848901518255600182019150602085019450602081019050611650565b86831015611692578489015161168e601f8916826115a1565b8355505b6001600288020188555050505b505050505050565b5f6040820190506116ba5f830185610ead565b6116c76020830184610d19565b9392505050565b7f576f726b207375626d6974206e6f7420636f6d70617461626c652077697468205f8201527f6368616c6c656e67650000000000000000000000000000000000000000000000602082015250565b5f611728602983610e24565b9150611733826116ce565b604082019050919050565b5f6020820190508181035f8301526117558161171c565b9050919050565b5f819050919050565b61177661177182610ea4565b61175c565b82525050565b5f6117878285611765565b6020820191506117978284611312565b6020820191508190509392505050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f611801602483610e24565b915061180c826117a7565b604082019050919050565b5f6020820190508181035f83015261182e816117f5565b9050919050565b5f61183f82610d10565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611871576118706110b9565b5b600182019050919050565b5f6118878385610e24565b9350611894838584611378565b61189d83610e5c565b840190509392505050565b5f60a0820190508181035f8301526118c181888a61187c565b90506118d06020830187610d19565b6118dd6040830186610d19565b6118ea6060830185610ead565b6118f76080830184610d19565b979650505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61195c602683610e24565b915061196782611902565b604082019050919050565b5f6020820190508181035f83015261198981611950565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6119c4602083610e24565b91506119cf82611990565b602082019050919050565b5f6020820190508181035f8301526119f1816119b8565b905091905056fea2646970667358221220348bc079102b8d6f82a5bc2a8404f3fb7820a3c902733fb3c6db6a86749c747864736f6c63430008150033",
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
// Solidity: function lastChallenges(address ) view returns(bytes32 challenge, uint256 difficulty, string nodeId, uint256 timestamp)
func (_Powtoken *PowtokenCaller) LastChallenges(opts *bind.CallOpts, arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "lastChallenges", arg0)

	outstruct := new(struct {
		Challenge  [32]byte
		Difficulty *big.Int
		NodeId     string
		Timestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NodeId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32 challenge, uint256 difficulty, string nodeId, uint256 timestamp)
func (_Powtoken *PowtokenSession) LastChallenges(arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
	return _Powtoken.Contract.LastChallenges(&_Powtoken.CallOpts, arg0)
}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32 challenge, uint256 difficulty, string nodeId, uint256 timestamp)
func (_Powtoken *PowtokenCallerSession) LastChallenges(arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
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
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Powtoken *PowtokenCaller) PowSubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
	Difficulty    *big.Int
}, error) {
	var out []interface{}
	err := _Powtoken.contract.Call(opts, &out, "powSubmissions", arg0, arg1)

	outstruct := new(struct {
		WalletAddress common.Address
		NodeId        string
		Nonce         *big.Int
		Timestamp     *big.Int
		Challenge     [32]byte
		Difficulty    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Challenge = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Powtoken *PowtokenSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
	Difficulty    *big.Int
}, error) {
	return _Powtoken.Contract.PowSubmissions(&_Powtoken.CallOpts, arg0, arg1)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Powtoken *PowtokenCallerSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress common.Address
	NodeId        string
	Nonce         *big.Int
	Timestamp     *big.Int
	Challenge     [32]byte
	Difficulty    *big.Int
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
	Difficult *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenerateChallenge is a free log retrieval operation binding the contract event 0x496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef.
//
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficult)
func (_Powtoken *PowtokenFilterer) FilterGenerateChallenge(opts *bind.FilterOpts) (*PowtokenGenerateChallengeIterator, error) {

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "GenerateChallenge")
	if err != nil {
		return nil, err
	}
	return &PowtokenGenerateChallengeIterator{contract: _Powtoken.contract, event: "GenerateChallenge", logs: logs, sub: sub}, nil
}

// WatchGenerateChallenge is a free log subscription operation binding the contract event 0x496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef.
//
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficult)
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

// ParseGenerateChallenge is a log parse operation binding the contract event 0x496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef.
//
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficult)
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

// PowtokenNewPowRoundIterator is returned from FilterNewPowRound and is used to iterate over the raw logs and unpacked data for NewPowRound events raised by the Powtoken contract.
type PowtokenNewPowRoundIterator struct {
	Event *PowtokenNewPowRound // Event containing the contract specifics and raw log

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
func (it *PowtokenNewPowRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowtokenNewPowRound)
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
		it.Event = new(PowtokenNewPowRound)
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
func (it *PowtokenNewPowRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowtokenNewPowRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowtokenNewPowRound represents a NewPowRound event raised by the Powtoken contract.
type PowtokenNewPowRound struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewPowRound is a free log retrieval operation binding the contract event 0x10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e.
//
// Solidity: event NewPowRound()
func (_Powtoken *PowtokenFilterer) FilterNewPowRound(opts *bind.FilterOpts) (*PowtokenNewPowRoundIterator, error) {

	logs, sub, err := _Powtoken.contract.FilterLogs(opts, "NewPowRound")
	if err != nil {
		return nil, err
	}
	return &PowtokenNewPowRoundIterator{contract: _Powtoken.contract, event: "NewPowRound", logs: logs, sub: sub}, nil
}

// WatchNewPowRound is a free log subscription operation binding the contract event 0x10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e.
//
// Solidity: event NewPowRound()
func (_Powtoken *PowtokenFilterer) WatchNewPowRound(opts *bind.WatchOpts, sink chan<- *PowtokenNewPowRound) (event.Subscription, error) {

	logs, sub, err := _Powtoken.contract.WatchLogs(opts, "NewPowRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowtokenNewPowRound)
				if err := _Powtoken.contract.UnpackLog(event, "NewPowRound", log); err != nil {
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

// ParseNewPowRound is a log parse operation binding the contract event 0x10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e.
//
// Solidity: event NewPowRound()
func (_Powtoken *PowtokenFilterer) ParseNewPowRound(log types.Log) (*PowtokenNewPowRound, error) {
	event := new(PowtokenNewPowRound)
	if err := _Powtoken.contract.UnpackLog(event, "NewPowRound", log); err != nil {
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
	Difficulty    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidPOWSubmitted is a free log retrieval operation binding the contract event 0x5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c860.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
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

// WatchValidPOWSubmitted is a free log subscription operation binding the contract event 0x5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c860.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
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

// ParseValidPOWSubmitted is a log parse operation binding the contract event 0x5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c860.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Powtoken *PowtokenFilterer) ParseValidPOWSubmitted(log types.Log) (*PowtokenValidPOWSubmitted, error) {
	event := new(PowtokenValidPOWSubmitted)
	if err := _Powtoken.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
