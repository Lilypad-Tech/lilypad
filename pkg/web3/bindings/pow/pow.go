// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pow

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

// LilypadPowPOWSubmission is an auto generated low-level Go binding around an user-defined struct.
type LilypadPowPOWSubmission struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
}

// PowMetaData contains all meta data concerning the Pow contract.
var PowMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPowRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"calculate_difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"change_difficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTimeWindow\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMinerPowSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMinerPowSubmissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"internalType\":\"structLilypadPow.POWSubmission[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061251a8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c8063a738bf8f116100b6578063bb52c1371161007a578063bb52c13714610323578063bfb660de14610341578063c2c434da1461034b578063da8accf91461037b578063e0d152af14610397578063f2fde38b146103c757610140565b8063a738bf8f146102a5578063ab098945146102c3578063adf0047c146102df578063b28d87ea146102fb578063b681f2fd1461031957610140565b80636189f3ac116101085780636189f3ac14610204578063715018a61461023757806378e97925146102415780638129fc1c1461025f5780638b2db16e146102695780638da5cb5b1461028757610140565b80631633da6e146101445780631bb0d808146101625780632d68c39714610192578063331bade1146101b05780634bbe05e4146101ce575b5f80fd5b61014c6103e3565b60405161015991906114d9565b60405180910390f35b61017c6004803603810190610177919061152b565b61046e565b604051610189919061156e565b60405180910390f35b61019a6104b7565b6040516101a7919061156e565b60405180910390f35b6101b86104f4565b6040516101c5919061156e565b60405180910390f35b6101e860048036038101906101e391906115b1565b6104fa565b6040516101fb97969594939291906116a0565b60405180910390f35b61021e6004803603810190610219919061152b565b6105f7565b60405161022e9493929190611714565b60405180910390f35b61023f6106a9565b005b6102496106bc565b604051610256919061156e565b60405180910390f35b6102676106c2565b005b61027161081e565b60405161027e919061156e565b60405180910390f35b61028f610824565b60405161029c919061175e565b60405180910390f35b6102ad61084c565b6040516102ba919061156e565b60405180910390f35b6102dd60048036038101906102d891906117d8565b610852565b005b6102f960048036038101906102f49190611823565b6109ae565b005b6103036109c0565b604051610310919061156e565b60405180910390f35b6103216109c6565b005b61032b610a16565b604051610338919061156e565b60405180910390f35b610349610a22565b005b6103656004803603810190610360919061152b565b610a68565b6040516103729190611a08565b60405180910390f35b61039560048036038101906103909190611a28565b610c17565b005b6103b160048036038101906103ac9190611823565b611113565b6040516103be919061175e565b60405180910390f35b6103e160048036038101906103dc919061152b565b61114e565b005b6060606780548060200260200160405190810160405280929190818152602001828054801561046457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161041b575b5050505050905090565b5f60665f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b5f806015446104c69190611ab2565b605a6104d29190611b0f565b90506064816065546104e49190611b42565b6104ee9190611b83565b91505090565b606b5481565b6066602052815f5260405f208181548110610513575f80fd5b905f5260205f2090600702015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461055890611be0565b80601f016020809104026020016040519081016040528092919081815260200182805461058490611be0565b80156105cf5780601f106105a6576101008083540402835291602001916105cf565b820191905f5260205f20905b8154815290600101906020018083116105b257829003601f168201915b5050505050908060020154908060030154908060040154908060050154908060060154905087565b6068602052805f5260405f205f91509050805f01549080600101549080600201805461062290611be0565b80601f016020809104026020016040519081016040528092919081815260200182805461064e90611be0565b80156106995780601f1061067057610100808354040283529160200191610699565b820191905f5260205f20905b81548152906001019060200180831161067c57829003601f168201915b5050505050908060030154905084565b6106b16111d0565b6106ba5f61124e565b565b606a5481565b5f8060019054906101000a900460ff161590508080156106f1575060015f8054906101000a900460ff1660ff16105b8061071d575061070030611311565b15801561071c575060015f8054906101000a900460ff1660ff16145b5b61075c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075390611c80565b60405180910390fd5b60015f806101000a81548160ff021916908360ff16021790555080156107975760015f60016101000a81548160ff0219169083151502179055505b61079f611333565b7c036f11b000000000000000000000000000000000000000000000000000606581905550801561081b575f8060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516108129190611cec565b60405180910390a15b50565b60655481565b5f60335f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606c5481565b61085a610a22565b5f42606b54338585604051602001610876959493929190611da6565b6040516020818303038152906040528051906020012090505f6108976104b7565b9050604051806080016040528083815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015560408201518160020190816109619190611fb2565b50606082015181600301559050507f496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef82826040516109a0929190612081565b60405180910390a150505050565b6109b66111d0565b8060658190555050565b60695481565b6109ce6111d0565b43606b819055506024436109e29190611b0f565b606c819055507f10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e60405160405180910390a1565b5f606780549050905090565b606c544310610a66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5d906120f2565b60405180910390fd5b565b606060665f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610c0c578382905f5260205f2090600702016040518060e00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610b4b90611be0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7790611be0565b8015610bc25780601f10610b9957610100808354040283529160200191610bc2565b820191905f5260205f20905b815481529060010190602001808311610ba557829003601f168201915b505050505081526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152505081526020019060010190610ac6565b505050509050919050565b610c1f610a22565b5f60685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f820154815260200160018201548152602001600282018054610c8b90611be0565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb790611be0565b8015610d025780601f10610cd957610100808354040283529160200191610d02565b820191905f5260205f20905b815481529060010190602001808311610ce557829003601f168201915b5050505050815260200160038201548152505090505f8160600151606b54338686604051602001610d37959493929190611da6565b60405160208183030381529060405280519060200120905080825f015114610d94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8b90612180565b60405180910390fd5b5f8186604051602001610da89291906121be565b6040516020818303038152906040528051906020012090508260200151815f1c10610e08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dff90612259565b60405180910390fd5b60695f815480929190610e1a90612277565b91905055505f60665f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f818054905003610ecb57606733908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b806040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200189815260200186606001518152602001428152602001865f015181526020018660200151815250908060018154018082558091505060019003905f5260205f2090600702015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610fe29190611fb2565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060155505060405180608001604052805f801b81526020015f815260200160405180602001604052805f81525081526020015f81525060685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015560408201518160020190816110ac9190611fb2565b50606082015181600301559050507f172b7d59e60446d8dfc98985344fb883c871ef150b1db4d1592bcb67699037323387878a8860600151428a5f01518b602001516040516111029897969594939291906122ea565b60405180910390a150505050505050565b60678181548110611122575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111566111d0565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bb906123d0565b60405180910390fd5b6111cd8161124e565b50565b6111d861138b565b73ffffffffffffffffffffffffffffffffffffffff166111f6610824565b73ffffffffffffffffffffffffffffffffffffffff161461124c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124390612438565b60405180910390fd5b565b5f60335f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160335f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f60019054906101000a900460ff16611381576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611378906124c6565b60405180910390fd5b611389611392565b565b5f33905090565b5f60019054906101000a900460ff166113e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d7906124c6565b60405180910390fd5b6113f06113eb61138b565b61124e565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114448261141b565b9050919050565b6114548161143a565b82525050565b5f611465838361144b565b60208301905092915050565b5f602082019050919050565b5f611487826113f2565b61149181856113fc565b935061149c8361140c565b805f5b838110156114cc5781516114b3888261145a565b97506114be83611471565b92505060018101905061149f565b5085935050505092915050565b5f6020820190508181035f8301526114f1818461147d565b905092915050565b5f80fd5b5f80fd5b61150a8161143a565b8114611514575f80fd5b50565b5f8135905061152581611501565b92915050565b5f602082840312156115405761153f6114f9565b5b5f61154d84828501611517565b91505092915050565b5f819050919050565b61156881611556565b82525050565b5f6020820190506115815f83018461155f565b92915050565b61159081611556565b811461159a575f80fd5b50565b5f813590506115ab81611587565b92915050565b5f80604083850312156115c7576115c66114f9565b5b5f6115d485828601611517565b92505060206115e58582860161159d565b9150509250929050565b6115f88161143a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561163557808201518184015260208101905061161a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61165a826115fe565b6116648185611608565b9350611674818560208601611618565b61167d81611640565b840191505092915050565b5f819050919050565b61169a81611688565b82525050565b5f60e0820190506116b35f83018a6115ef565b81810360208301526116c58189611650565b90506116d4604083018861155f565b6116e1606083018761155f565b6116ee608083018661155f565b6116fb60a0830185611691565b61170860c083018461155f565b98975050505050505050565b5f6080820190506117275f830187611691565b611734602083018661155f565b81810360408301526117468185611650565b9050611755606083018461155f565b95945050505050565b5f6020820190506117715f8301846115ef565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261179857611797611777565b5b8235905067ffffffffffffffff8111156117b5576117b461177b565b5b6020830191508360018202830111156117d1576117d061177f565b5b9250929050565b5f80602083850312156117ee576117ed6114f9565b5b5f83013567ffffffffffffffff81111561180b5761180a6114fd565b5b61181785828601611783565b92509250509250929050565b5f60208284031215611838576118376114f9565b5b5f6118458482850161159d565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f611891826115fe565b61189b8185611877565b93506118ab818560208601611618565b6118b481611640565b840191505092915050565b6118c881611556565b82525050565b6118d781611688565b82525050565b5f60e083015f8301516118f25f86018261144b565b506020830151848203602086015261190a8282611887565b915050604083015161191f60408601826118bf565b50606083015161193260608601826118bf565b50608083015161194560808601826118bf565b5060a083015161195860a08601826118ce565b5060c083015161196b60c08601826118bf565b508091505092915050565b5f61198183836118dd565b905092915050565b5f602082019050919050565b5f61199f8261184e565b6119a98185611858565b9350836020820285016119bb85611868565b805f5b858110156119f657848403895281516119d78582611976565b94506119e283611989565b925060208a019950506001810190506119be565b50829750879550505050505092915050565b5f6020820190508181035f830152611a208184611995565b905092915050565b5f805f60408486031215611a3f57611a3e6114f9565b5b5f611a4c8682870161159d565b935050602084013567ffffffffffffffff811115611a6d57611a6c6114fd565b5b611a7986828701611783565b92509250509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611abc82611556565b9150611ac783611556565b925082611ad757611ad6611a85565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611b1982611556565b9150611b2483611556565b9250828201905080821115611b3c57611b3b611ae2565b5b92915050565b5f611b4c82611556565b9150611b5783611556565b9250828202611b6581611556565b91508282048414831517611b7c57611b7b611ae2565b5b5092915050565b5f611b8d82611556565b9150611b9883611556565b925082611ba857611ba7611a85565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611bf757607f821691505b602082108103611c0a57611c09611bb3565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611c6a602e83611608565b9150611c7582611c10565b604082019050919050565b5f6020820190508181035f830152611c9781611c5e565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f611cd6611cd1611ccc84611c9e565b611cb3565b611ca7565b9050919050565b611ce681611cbc565b82525050565b5f602082019050611cff5f830184611cdd565b92915050565b5f819050919050565b611d1f611d1a82611556565b611d05565b82525050565b5f8160601b9050919050565b5f611d3b82611d25565b9050919050565b5f611d4c82611d31565b9050919050565b611d64611d5f8261143a565b611d42565b82525050565b5f81905092915050565b828183375f83830152505050565b5f611d8d8385611d6a565b9350611d9a838584611d74565b82840190509392505050565b5f611db18288611d0e565b602082019150611dc18287611d0e565b602082019150611dd18286611d53565b601482019150611de2828486611d82565b91508190509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611e7a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611e3f565b611e848683611e3f565b95508019841693508086168417925050509392505050565b5f611eb6611eb1611eac84611556565b611cb3565b611556565b9050919050565b5f819050919050565b611ecf83611e9c565b611ee3611edb82611ebd565b848454611e4b565b825550505050565b5f90565b611ef7611eeb565b611f02818484611ec6565b505050565b5b81811015611f2557611f1a5f82611eef565b600181019050611f08565b5050565b601f821115611f6a57611f3b81611e1e565b611f4484611e30565b81016020851015611f53578190505b611f67611f5f85611e30565b830182611f07565b50505b505050565b5f82821c905092915050565b5f611f8a5f1984600802611f6f565b1980831691505092915050565b5f611fa28383611f7b565b9150826002028217905092915050565b611fbb826115fe565b67ffffffffffffffff811115611fd457611fd3611df1565b5b611fde8254611be0565b611fe9828285611f29565b5f60209050601f83116001811461201a575f8415612008578287015190505b6120128582611f97565b865550612079565b601f19841661202886611e1e565b5f5b8281101561204f5784890151825560018201915060208501945060208101905061202a565b8683101561206c5784890151612068601f891682611f7b565b8355505b6001600288020188555050505b505050505050565b5f6040820190506120945f830185611691565b6120a1602083018461155f565b9392505050565b7f70726f6f662077696e646f77732068617320636c6f73656400000000000000005f82015250565b5f6120dc601883611608565b91506120e7826120a8565b602082019050919050565b5f6020820190508181035f830152612109816120d0565b9050919050565b7f576f726b207375626d6974206e6f7420636f6d70617461626c652077697468205f8201527f6368616c6c656e67650000000000000000000000000000000000000000000000602082015250565b5f61216a602983611608565b915061217582612110565b604082019050919050565b5f6020820190508181035f8301526121978161215e565b9050919050565b5f819050919050565b6121b86121b382611688565b61219e565b82525050565b5f6121c982856121a7565b6020820191506121d98284611d0e565b6020820191508190509392505050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f612243602483611608565b915061224e826121e9565b604082019050919050565b5f6020820190508181035f83015261227081612237565b9050919050565b5f61228182611556565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036122b3576122b2611ae2565b5b600182019050919050565b5f6122c98385611608565b93506122d6838584611d74565b6122df83611640565b840190509392505050565b5f60e0820190506122fd5f83018b6115ef565b818103602083015261231081898b6122be565b905061231f604083018861155f565b61232c606083018761155f565b612339608083018661155f565b61234660a0830185611691565b61235360c083018461155f565b9998505050505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6123ba602683611608565b91506123c582612360565b604082019050919050565b5f6020820190508181035f8301526123e7816123ae565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f612422602083611608565b915061242d826123ee565b602082019050919050565b5f6020820190508181035f83015261244f81612416565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f7420695f8201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b5f6124b0602b83611608565b91506124bb82612456565b604082019050919050565b5f6020820190508181035f8301526124dd816124a4565b905091905056fea2646970667358221220a5b9dff04c384855a865d0ea5063cceacd0fd1c6fa0e93876968da9d327a3a9164736f6c63430008150033",
}

// PowABI is the input ABI used to generate the binding from.
// Deprecated: Use PowMetaData.ABI instead.
var PowABI = PowMetaData.ABI

// PowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PowMetaData.Bin instead.
var PowBin = PowMetaData.Bin

// DeployPow deploys a new Ethereum contract, binding an instance of Pow to it.
func DeployPow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pow, error) {
	parsed, err := PowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pow{PowCaller: PowCaller{contract: contract}, PowTransactor: PowTransactor{contract: contract}, PowFilterer: PowFilterer{contract: contract}}, nil
}

// Pow is an auto generated Go binding around an Ethereum contract.
type Pow struct {
	PowCaller     // Read-only binding to the contract
	PowTransactor // Write-only binding to the contract
	PowFilterer   // Log filterer for contract events
}

// PowCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowSession struct {
	Contract     *Pow              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowCallerSession struct {
	Contract *PowCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowTransactorSession struct {
	Contract     *PowTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowRaw struct {
	Contract *Pow // Generic contract binding to access the raw methods on
}

// PowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowCallerRaw struct {
	Contract *PowCaller // Generic read-only contract binding to access the raw methods on
}

// PowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowTransactorRaw struct {
	Contract *PowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPow creates a new instance of Pow, bound to a specific deployed contract.
func NewPow(address common.Address, backend bind.ContractBackend) (*Pow, error) {
	contract, err := bindPow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pow{PowCaller: PowCaller{contract: contract}, PowTransactor: PowTransactor{contract: contract}, PowFilterer: PowFilterer{contract: contract}}, nil
}

// NewPowCaller creates a new read-only instance of Pow, bound to a specific deployed contract.
func NewPowCaller(address common.Address, caller bind.ContractCaller) (*PowCaller, error) {
	contract, err := bindPow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PowCaller{contract: contract}, nil
}

// NewPowTransactor creates a new write-only instance of Pow, bound to a specific deployed contract.
func NewPowTransactor(address common.Address, transactor bind.ContractTransactor) (*PowTransactor, error) {
	contract, err := bindPow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PowTransactor{contract: contract}, nil
}

// NewPowFilterer creates a new log filterer instance of Pow, bound to a specific deployed contract.
func NewPowFilterer(address common.Address, filterer bind.ContractFilterer) (*PowFilterer, error) {
	contract, err := bindPow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PowFilterer{contract: contract}, nil
}

// bindPow binds a generic wrapper to an already deployed contract.
func bindPow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pow *PowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pow.Contract.PowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pow *PowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pow.Contract.PowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pow *PowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pow.Contract.PowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pow *PowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pow *PowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pow *PowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pow.Contract.contract.Transact(opts, method, params...)
}

// CalculateDifficulty is a free data retrieval call binding the contract method 0x2d68c397.
//
// Solidity: function calculate_difficulty() view returns(uint256)
func (_Pow *PowCaller) CalculateDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "calculate_difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateDifficulty is a free data retrieval call binding the contract method 0x2d68c397.
//
// Solidity: function calculate_difficulty() view returns(uint256)
func (_Pow *PowSession) CalculateDifficulty() (*big.Int, error) {
	return _Pow.Contract.CalculateDifficulty(&_Pow.CallOpts)
}

// CalculateDifficulty is a free data retrieval call binding the contract method 0x2d68c397.
//
// Solidity: function calculate_difficulty() view returns(uint256)
func (_Pow *PowCallerSession) CalculateDifficulty() (*big.Int, error) {
	return _Pow.Contract.CalculateDifficulty(&_Pow.CallOpts)
}

// CheckTimeWindow is a free data retrieval call binding the contract method 0xbfb660de.
//
// Solidity: function checkTimeWindow() view returns()
func (_Pow *PowCaller) CheckTimeWindow(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "checkTimeWindow")

	if err != nil {
		return err
	}

	return err

}

// CheckTimeWindow is a free data retrieval call binding the contract method 0xbfb660de.
//
// Solidity: function checkTimeWindow() view returns()
func (_Pow *PowSession) CheckTimeWindow() error {
	return _Pow.Contract.CheckTimeWindow(&_Pow.CallOpts)
}

// CheckTimeWindow is a free data retrieval call binding the contract method 0xbfb660de.
//
// Solidity: function checkTimeWindow() view returns()
func (_Pow *PowCallerSession) CheckTimeWindow() error {
	return _Pow.Contract.CheckTimeWindow(&_Pow.CallOpts)
}

// GetMinerCount is a free data retrieval call binding the contract method 0xbb52c137.
//
// Solidity: function getMinerCount() view returns(uint256)
func (_Pow *PowCaller) GetMinerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "getMinerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinerCount is a free data retrieval call binding the contract method 0xbb52c137.
//
// Solidity: function getMinerCount() view returns(uint256)
func (_Pow *PowSession) GetMinerCount() (*big.Int, error) {
	return _Pow.Contract.GetMinerCount(&_Pow.CallOpts)
}

// GetMinerCount is a free data retrieval call binding the contract method 0xbb52c137.
//
// Solidity: function getMinerCount() view returns(uint256)
func (_Pow *PowCallerSession) GetMinerCount() (*big.Int, error) {
	return _Pow.Contract.GetMinerCount(&_Pow.CallOpts)
}

// GetMinerPowSubmissionCount is a free data retrieval call binding the contract method 0x1bb0d808.
//
// Solidity: function getMinerPowSubmissionCount(address addr) view returns(uint256)
func (_Pow *PowCaller) GetMinerPowSubmissionCount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "getMinerPowSubmissionCount", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinerPowSubmissionCount is a free data retrieval call binding the contract method 0x1bb0d808.
//
// Solidity: function getMinerPowSubmissionCount(address addr) view returns(uint256)
func (_Pow *PowSession) GetMinerPowSubmissionCount(addr common.Address) (*big.Int, error) {
	return _Pow.Contract.GetMinerPowSubmissionCount(&_Pow.CallOpts, addr)
}

// GetMinerPowSubmissionCount is a free data retrieval call binding the contract method 0x1bb0d808.
//
// Solidity: function getMinerPowSubmissionCount(address addr) view returns(uint256)
func (_Pow *PowCallerSession) GetMinerPowSubmissionCount(addr common.Address) (*big.Int, error) {
	return _Pow.Contract.GetMinerPowSubmissionCount(&_Pow.CallOpts, addr)
}

// GetMinerPowSubmissions is a free data retrieval call binding the contract method 0xc2c434da.
//
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowCaller) GetMinerPowSubmissions(opts *bind.CallOpts, addr common.Address) ([]LilypadPowPOWSubmission, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "getMinerPowSubmissions", addr)

	if err != nil {
		return *new([]LilypadPowPOWSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new([]LilypadPowPOWSubmission)).(*[]LilypadPowPOWSubmission)

	return out0, err

}

// GetMinerPowSubmissions is a free data retrieval call binding the contract method 0xc2c434da.
//
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowSession) GetMinerPowSubmissions(addr common.Address) ([]LilypadPowPOWSubmission, error) {
	return _Pow.Contract.GetMinerPowSubmissions(&_Pow.CallOpts, addr)
}

// GetMinerPowSubmissions is a free data retrieval call binding the contract method 0xc2c434da.
//
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowCallerSession) GetMinerPowSubmissions(addr common.Address) ([]LilypadPowPOWSubmission, error) {
	return _Pow.Contract.GetMinerPowSubmissions(&_Pow.CallOpts, addr)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_Pow *PowCaller) GetMiners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_Pow *PowSession) GetMiners() ([]common.Address, error) {
	return _Pow.Contract.GetMiners(&_Pow.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_Pow *PowCallerSession) GetMiners() ([]common.Address, error) {
	return _Pow.Contract.GetMiners(&_Pow.CallOpts)
}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32 challenge, uint256 difficulty, string nodeId, uint256 timestamp)
func (_Pow *PowCaller) LastChallenges(opts *bind.CallOpts, arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "lastChallenges", arg0)

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
func (_Pow *PowSession) LastChallenges(arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
	return _Pow.Contract.LastChallenges(&_Pow.CallOpts, arg0)
}

// LastChallenges is a free data retrieval call binding the contract method 0x6189f3ac.
//
// Solidity: function lastChallenges(address ) view returns(bytes32 challenge, uint256 difficulty, string nodeId, uint256 timestamp)
func (_Pow *PowCallerSession) LastChallenges(arg0 common.Address) (struct {
	Challenge  [32]byte
	Difficulty *big.Int
	NodeId     string
	Timestamp  *big.Int
}, error) {
	return _Pow.Contract.LastChallenges(&_Pow.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xe0d152af.
//
// Solidity: function miners(uint256 ) view returns(address)
func (_Pow *PowCaller) Miners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "miners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Miners is a free data retrieval call binding the contract method 0xe0d152af.
//
// Solidity: function miners(uint256 ) view returns(address)
func (_Pow *PowSession) Miners(arg0 *big.Int) (common.Address, error) {
	return _Pow.Contract.Miners(&_Pow.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0xe0d152af.
//
// Solidity: function miners(uint256 ) view returns(address)
func (_Pow *PowCallerSession) Miners(arg0 *big.Int) (common.Address, error) {
	return _Pow.Contract.Miners(&_Pow.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pow *PowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pow *PowSession) Owner() (common.Address, error) {
	return _Pow.Contract.Owner(&_Pow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pow *PowCallerSession) Owner() (common.Address, error) {
	return _Pow.Contract.Owner(&_Pow.CallOpts)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowCaller) PowSubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
}, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "powSubmissions", arg0, arg1)

	outstruct := new(struct {
		WalletAddress     common.Address
		NodeId            string
		Nonce             *big.Int
		StartTimestamp    *big.Int
		CompleteTimestamp *big.Int
		Challenge         [32]byte
		Difficulty        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CompleteTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Challenge = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
}, error) {
	return _Pow.Contract.PowSubmissions(&_Pow.CallOpts, arg0, arg1)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowCallerSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
}, error) {
	return _Pow.Contract.PowSubmissions(&_Pow.CallOpts, arg0, arg1)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Pow *PowCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Pow *PowSession) StartTime() (*big.Int, error) {
	return _Pow.Contract.StartTime(&_Pow.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Pow *PowCallerSession) StartTime() (*big.Int, error) {
	return _Pow.Contract.StartTime(&_Pow.CallOpts)
}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_Pow *PowCaller) TargetDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "targetDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_Pow *PowSession) TargetDifficulty() (*big.Int, error) {
	return _Pow.Contract.TargetDifficulty(&_Pow.CallOpts)
}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_Pow *PowCallerSession) TargetDifficulty() (*big.Int, error) {
	return _Pow.Contract.TargetDifficulty(&_Pow.CallOpts)
}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Pow *PowCaller) ValidProofs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "validProofs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Pow *PowSession) ValidProofs() (*big.Int, error) {
	return _Pow.Contract.ValidProofs(&_Pow.CallOpts)
}

// ValidProofs is a free data retrieval call binding the contract method 0xb28d87ea.
//
// Solidity: function validProofs() view returns(uint256)
func (_Pow *PowCallerSession) ValidProofs() (*big.Int, error) {
	return _Pow.Contract.ValidProofs(&_Pow.CallOpts)
}

// WindowEnd is a free data retrieval call binding the contract method 0xa738bf8f.
//
// Solidity: function window_end() view returns(uint256)
func (_Pow *PowCaller) WindowEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "window_end")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WindowEnd is a free data retrieval call binding the contract method 0xa738bf8f.
//
// Solidity: function window_end() view returns(uint256)
func (_Pow *PowSession) WindowEnd() (*big.Int, error) {
	return _Pow.Contract.WindowEnd(&_Pow.CallOpts)
}

// WindowEnd is a free data retrieval call binding the contract method 0xa738bf8f.
//
// Solidity: function window_end() view returns(uint256)
func (_Pow *PowCallerSession) WindowEnd() (*big.Int, error) {
	return _Pow.Contract.WindowEnd(&_Pow.CallOpts)
}

// WindowStart is a free data retrieval call binding the contract method 0x331bade1.
//
// Solidity: function window_start() view returns(uint256)
func (_Pow *PowCaller) WindowStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "window_start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WindowStart is a free data retrieval call binding the contract method 0x331bade1.
//
// Solidity: function window_start() view returns(uint256)
func (_Pow *PowSession) WindowStart() (*big.Int, error) {
	return _Pow.Contract.WindowStart(&_Pow.CallOpts)
}

// WindowStart is a free data retrieval call binding the contract method 0x331bade1.
//
// Solidity: function window_start() view returns(uint256)
func (_Pow *PowCallerSession) WindowStart() (*big.Int, error) {
	return _Pow.Contract.WindowStart(&_Pow.CallOpts)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0xadf0047c.
//
// Solidity: function change_difficulty(uint256 difficulty) returns()
func (_Pow *PowTransactor) ChangeDifficulty(opts *bind.TransactOpts, difficulty *big.Int) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "change_difficulty", difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0xadf0047c.
//
// Solidity: function change_difficulty(uint256 difficulty) returns()
func (_Pow *PowSession) ChangeDifficulty(difficulty *big.Int) (*types.Transaction, error) {
	return _Pow.Contract.ChangeDifficulty(&_Pow.TransactOpts, difficulty)
}

// ChangeDifficulty is a paid mutator transaction binding the contract method 0xadf0047c.
//
// Solidity: function change_difficulty(uint256 difficulty) returns()
func (_Pow *PowTransactorSession) ChangeDifficulty(difficulty *big.Int) (*types.Transaction, error) {
	return _Pow.Contract.ChangeDifficulty(&_Pow.TransactOpts, difficulty)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Pow *PowTransactor) GenerateChallenge(opts *bind.TransactOpts, nodeId string) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "generateChallenge", nodeId)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Pow *PowSession) GenerateChallenge(nodeId string) (*types.Transaction, error) {
	return _Pow.Contract.GenerateChallenge(&_Pow.TransactOpts, nodeId)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xab098945.
//
// Solidity: function generateChallenge(string nodeId) returns()
func (_Pow *PowTransactorSession) GenerateChallenge(nodeId string) (*types.Transaction, error) {
	return _Pow.Contract.GenerateChallenge(&_Pow.TransactOpts, nodeId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pow *PowTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pow *PowSession) Initialize() (*types.Transaction, error) {
	return _Pow.Contract.Initialize(&_Pow.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pow *PowTransactorSession) Initialize() (*types.Transaction, error) {
	return _Pow.Contract.Initialize(&_Pow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pow *PowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pow *PowSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pow.Contract.RenounceOwnership(&_Pow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pow *PowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pow.Contract.RenounceOwnership(&_Pow.TransactOpts)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Pow *PowTransactor) SubmitWork(opts *bind.TransactOpts, nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "submitWork", nonce, nodeId)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Pow *PowSession) SubmitWork(nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Pow.Contract.SubmitWork(&_Pow.TransactOpts, nonce, nodeId)
}

// SubmitWork is a paid mutator transaction binding the contract method 0xda8accf9.
//
// Solidity: function submitWork(uint256 nonce, string nodeId) returns()
func (_Pow *PowTransactorSession) SubmitWork(nonce *big.Int, nodeId string) (*types.Transaction, error) {
	return _Pow.Contract.SubmitWork(&_Pow.TransactOpts, nonce, nodeId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pow *PowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pow *PowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pow.Contract.TransferOwnership(&_Pow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pow *PowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pow.Contract.TransferOwnership(&_Pow.TransactOpts, newOwner)
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Pow *PowTransactor) TriggerNewPowRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "triggerNewPowRound")
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Pow *PowSession) TriggerNewPowRound() (*types.Transaction, error) {
	return _Pow.Contract.TriggerNewPowRound(&_Pow.TransactOpts)
}

// TriggerNewPowRound is a paid mutator transaction binding the contract method 0xb681f2fd.
//
// Solidity: function triggerNewPowRound() returns()
func (_Pow *PowTransactorSession) TriggerNewPowRound() (*types.Transaction, error) {
	return _Pow.Contract.TriggerNewPowRound(&_Pow.TransactOpts)
}

// PowGenerateChallengeIterator is returned from FilterGenerateChallenge and is used to iterate over the raw logs and unpacked data for GenerateChallenge events raised by the Pow contract.
type PowGenerateChallengeIterator struct {
	Event *PowGenerateChallenge // Event containing the contract specifics and raw log

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
func (it *PowGenerateChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowGenerateChallenge)
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
		it.Event = new(PowGenerateChallenge)
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
func (it *PowGenerateChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowGenerateChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowGenerateChallenge represents a GenerateChallenge event raised by the Pow contract.
type PowGenerateChallenge struct {
	Challenge  [32]byte
	Difficulty *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGenerateChallenge is a free log retrieval operation binding the contract event 0x496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef.
//
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) FilterGenerateChallenge(opts *bind.FilterOpts) (*PowGenerateChallengeIterator, error) {

	logs, sub, err := _Pow.contract.FilterLogs(opts, "GenerateChallenge")
	if err != nil {
		return nil, err
	}
	return &PowGenerateChallengeIterator{contract: _Pow.contract, event: "GenerateChallenge", logs: logs, sub: sub}, nil
}

// WatchGenerateChallenge is a free log subscription operation binding the contract event 0x496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef.
//
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) WatchGenerateChallenge(opts *bind.WatchOpts, sink chan<- *PowGenerateChallenge) (event.Subscription, error) {

	logs, sub, err := _Pow.contract.WatchLogs(opts, "GenerateChallenge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowGenerateChallenge)
				if err := _Pow.contract.UnpackLog(event, "GenerateChallenge", log); err != nil {
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
// Solidity: event GenerateChallenge(bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) ParseGenerateChallenge(log types.Log) (*PowGenerateChallenge, error) {
	event := new(PowGenerateChallenge)
	if err := _Pow.contract.UnpackLog(event, "GenerateChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pow contract.
type PowInitializedIterator struct {
	Event *PowInitialized // Event containing the contract specifics and raw log

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
func (it *PowInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowInitialized)
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
		it.Event = new(PowInitialized)
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
func (it *PowInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowInitialized represents a Initialized event raised by the Pow contract.
type PowInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pow *PowFilterer) FilterInitialized(opts *bind.FilterOpts) (*PowInitializedIterator, error) {

	logs, sub, err := _Pow.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PowInitializedIterator{contract: _Pow.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pow *PowFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PowInitialized) (event.Subscription, error) {

	logs, sub, err := _Pow.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowInitialized)
				if err := _Pow.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pow *PowFilterer) ParseInitialized(log types.Log) (*PowInitialized, error) {
	event := new(PowInitialized)
	if err := _Pow.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowNewPowRoundIterator is returned from FilterNewPowRound and is used to iterate over the raw logs and unpacked data for NewPowRound events raised by the Pow contract.
type PowNewPowRoundIterator struct {
	Event *PowNewPowRound // Event containing the contract specifics and raw log

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
func (it *PowNewPowRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowNewPowRound)
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
		it.Event = new(PowNewPowRound)
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
func (it *PowNewPowRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowNewPowRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowNewPowRound represents a NewPowRound event raised by the Pow contract.
type PowNewPowRound struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewPowRound is a free log retrieval operation binding the contract event 0x10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e.
//
// Solidity: event NewPowRound()
func (_Pow *PowFilterer) FilterNewPowRound(opts *bind.FilterOpts) (*PowNewPowRoundIterator, error) {

	logs, sub, err := _Pow.contract.FilterLogs(opts, "NewPowRound")
	if err != nil {
		return nil, err
	}
	return &PowNewPowRoundIterator{contract: _Pow.contract, event: "NewPowRound", logs: logs, sub: sub}, nil
}

// WatchNewPowRound is a free log subscription operation binding the contract event 0x10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e.
//
// Solidity: event NewPowRound()
func (_Pow *PowFilterer) WatchNewPowRound(opts *bind.WatchOpts, sink chan<- *PowNewPowRound) (event.Subscription, error) {

	logs, sub, err := _Pow.contract.WatchLogs(opts, "NewPowRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowNewPowRound)
				if err := _Pow.contract.UnpackLog(event, "NewPowRound", log); err != nil {
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
func (_Pow *PowFilterer) ParseNewPowRound(log types.Log) (*PowNewPowRound, error) {
	event := new(PowNewPowRound)
	if err := _Pow.contract.UnpackLog(event, "NewPowRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pow contract.
type PowOwnershipTransferredIterator struct {
	Event *PowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowOwnershipTransferred)
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
		it.Event = new(PowOwnershipTransferred)
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
func (it *PowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowOwnershipTransferred represents a OwnershipTransferred event raised by the Pow contract.
type PowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pow *PowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PowOwnershipTransferredIterator{contract: _Pow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pow *PowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowOwnershipTransferred)
				if err := _Pow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pow *PowFilterer) ParseOwnershipTransferred(log types.Log) (*PowOwnershipTransferred, error) {
	event := new(PowOwnershipTransferred)
	if err := _Pow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowValidPOWSubmittedIterator is returned from FilterValidPOWSubmitted and is used to iterate over the raw logs and unpacked data for ValidPOWSubmitted events raised by the Pow contract.
type PowValidPOWSubmittedIterator struct {
	Event *PowValidPOWSubmitted // Event containing the contract specifics and raw log

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
func (it *PowValidPOWSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PowValidPOWSubmitted)
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
		it.Event = new(PowValidPOWSubmitted)
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
func (it *PowValidPOWSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PowValidPOWSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PowValidPOWSubmitted represents a ValidPOWSubmitted event raised by the Pow contract.
type PowValidPOWSubmitted struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidPOWSubmitted is a free log retrieval operation binding the contract event 0x172b7d59e60446d8dfc98985344fb883c871ef150b1db4d1592bcb6769903732.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) FilterValidPOWSubmitted(opts *bind.FilterOpts) (*PowValidPOWSubmittedIterator, error) {

	logs, sub, err := _Pow.contract.FilterLogs(opts, "ValidPOWSubmitted")
	if err != nil {
		return nil, err
	}
	return &PowValidPOWSubmittedIterator{contract: _Pow.contract, event: "ValidPOWSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidPOWSubmitted is a free log subscription operation binding the contract event 0x172b7d59e60446d8dfc98985344fb883c871ef150b1db4d1592bcb6769903732.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) WatchValidPOWSubmitted(opts *bind.WatchOpts, sink chan<- *PowValidPOWSubmitted) (event.Subscription, error) {

	logs, sub, err := _Pow.contract.WatchLogs(opts, "ValidPOWSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PowValidPOWSubmitted)
				if err := _Pow.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
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

// ParseValidPOWSubmitted is a log parse operation binding the contract event 0x172b7d59e60446d8dfc98985344fb883c871ef150b1db4d1592bcb6769903732.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) ParseValidPOWSubmitted(log types.Log) (*PowValidPOWSubmitted, error) {
	event := new(PowValidPOWSubmitted)
	if err := _Pow.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
