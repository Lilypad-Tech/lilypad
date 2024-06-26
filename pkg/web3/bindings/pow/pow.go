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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPowRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"calculate_difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"change_difficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTimeWindow\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMinerPosSubmissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"internalType\":\"structLilypadPow.POWSubmission[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040527c149a6a3000000000000000000000000000000000000000000000000000600155348015610030575f80fd5b5061004d61004261005260201b60201c565b61005960201b60201c565b61011a565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6122f180620001285f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c80638da5cb5b116100ab578063b681f2fd1161006f578063b681f2fd14610303578063bfb660de1461030d578063da8accf914610317578063e0d152af14610333578063f2fde38b146103635761012a565b80638da5cb5b14610271578063a738bf8f1461028f578063ab098945146102ad578063adf0047c146102c9578063b28d87ea146102e55761012a565b80636189f3ac116100f25780636189f3ac146101ee578063715018a61461022157806378e979251461022b5780638129fc1c146102495780638b2db16e146102535761012a565b80631633da6e1461012e5780632d68c3971461014c578063331bade11461016a5780634bbe05e41461018857806354bfd759146101be575b5f80fd5b61013661037f565b604051610143919061133e565b60405180910390f35b61015461040a565b6040516101619190611376565b60405180910390f35b610172610447565b60405161017f9190611376565b60405180910390f35b6101a2600480360381019061019d91906113eb565b61044d565b6040516101b597969594939291906114da565b60405180910390f35b6101d860048036038101906101d3919061154e565b61054a565b6040516101e59190611733565b60405180910390f35b6102086004803603810190610203919061154e565b6106f9565b6040516102189493929190611753565b60405180910390f35b6102296107ab565b005b6102336107be565b6040516102409190611376565b60405180910390f35b6102516107c4565b005b61025b6108f9565b6040516102689190611376565b60405180910390f35b6102796108ff565b604051610286919061179d565b60405180910390f35b610297610926565b6040516102a49190611376565b60405180910390f35b6102c760048036038101906102c29190611817565b61092c565b005b6102e360048036038101906102de9190611862565b610a88565b005b6102ed610a9a565b6040516102fa9190611376565b60405180910390f35b61030b610aa0565b005b610315610af0565b005b610331600480360381019061032c919061188d565b610b36565b005b61034d60048036038101906103489190611862565b611032565b60405161035a919061179d565b60405180910390f35b61037d6004803603810190610378919061154e565b61106d565b005b6060600380548060200260200160405190810160405280929190818152602001828054801561040057602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103b7575b5050505050905090565b5f806015446104199190611917565b605a6104259190611974565b905060648160015461043791906119a7565b61044191906119e8565b91505090565b60075481565b6002602052815f5260405f208181548110610466575f80fd5b905f5260205f2090600702015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546104ab90611a45565b80601f01602080910402602001604051908101604052809291908181526020018280546104d790611a45565b80156105225780601f106104f957610100808354040283529160200191610522565b820191905f5260205f20905b81548152906001019060200180831161050557829003601f168201915b5050505050908060020154908060030154908060040154908060050154908060060154905087565b606060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156106ee578382905f5260205f2090600702016040518060e00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461062d90611a45565b80601f016020809104026020016040519081016040528092919081815260200182805461065990611a45565b80156106a45780601f1061067b576101008083540402835291602001916106a4565b820191905f5260205f20905b81548152906001019060200180831161068757829003601f168201915b5050505050815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481525050815260200190600101906105a8565b505050509050919050565b6004602052805f5260405f205f91509050805f01549080600101549080600201805461072490611a45565b80601f016020809104026020016040519081016040528092919081815260200182805461075090611a45565b801561079b5780601f106107725761010080835404028352916020019161079b565b820191905f5260205f20905b81548152906001019060200180831161077e57829003601f168201915b5050505050908060030154905084565b6107b36110ef565b6107bc5f61116d565b565b60065481565b5f8060159054906101000a900460ff161590508080156107f5575060015f60149054906101000a900460ff1660ff16105b8061082357506108043061122e565b158015610822575060015f60149054906101000a900460ff1660ff16145b5b610862576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085990611ae5565b60405180910390fd5b60015f60146101000a81548160ff021916908360ff160217905550801561089e5760015f60156101000a81548160ff0219169083151502179055505b80156108f6575f8060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516108ed9190611b51565b60405180910390a15b50565b60015481565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60085481565b610934610af0565b5f42600754338585604051602001610950959493929190611c0b565b6040516020818303038152906040528051906020012090505f61097161040a565b9050604051806080016040528083815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0155602082015181600101556040820151816002019081610a3b9190611e17565b50606082015181600301559050507f496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef8282604051610a7a929190611ee6565b60405180910390a150505050565b610a906110ef565b8060018190555050565b60055481565b610aa86110ef565b43600781905550601e43610abc9190611974565b6008819055507f10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e60405160405180910390a1565b6008544310610b34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2b90611f57565b60405180910390fd5b565b610b3e610af0565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f820154815260200160018201548152602001600282018054610baa90611a45565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd690611a45565b8015610c215780601f10610bf857610100808354040283529160200191610c21565b820191905f5260205f20905b815481529060010190602001808311610c0457829003601f168201915b5050505050815260200160038201548152505090505f8160600151600754338686604051602001610c56959493929190611c0b565b60405160208183030381529060405280519060200120905080825f015114610cb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610caa90611fe5565b60405180910390fd5b5f8186604051602001610cc7929190612023565b6040516020818303038152906040528051906020012090508260200151815f1c10610d27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1e906120be565b60405180910390fd5b60055f815480929190610d39906120dc565b91905055505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f818054905003610dea57600333908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b806040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200189815260200186606001518152602001428152602001865f015181526020018660200151815250908060018154018082558091505060019003905f5260205f2090600702015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610f019190611e17565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060155505060405180608001604052805f801b81526020015f815260200160405180602001604052805f81525081526020015f81525060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0155602082015181600101556040820151816002019081610fcb9190611e17565b50606082015181600301559050507f172b7d59e60446d8dfc98985344fb883c871ef150b1db4d1592bcb67699037323387878a8860600151428a5f01518b6020015160405161102198979695949392919061214f565b60405180910390a150505050505050565b60038181548110611041575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6110756110ef565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036110e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110da90612235565b60405180910390fd5b6110ec8161116d565b50565b6110f7611250565b73ffffffffffffffffffffffffffffffffffffffff166111156108ff565b73ffffffffffffffffffffffffffffffffffffffff161461116b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111629061229d565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112a982611280565b9050919050565b6112b98161129f565b82525050565b5f6112ca83836112b0565b60208301905092915050565b5f602082019050919050565b5f6112ec82611257565b6112f68185611261565b935061130183611271565b805f5b8381101561133157815161131888826112bf565b9750611323836112d6565b925050600181019050611304565b5085935050505092915050565b5f6020820190508181035f83015261135681846112e2565b905092915050565b5f819050919050565b6113708161135e565b82525050565b5f6020820190506113895f830184611367565b92915050565b5f80fd5b5f80fd5b6113a08161129f565b81146113aa575f80fd5b50565b5f813590506113bb81611397565b92915050565b6113ca8161135e565b81146113d4575f80fd5b50565b5f813590506113e5816113c1565b92915050565b5f80604083850312156114015761140061138f565b5b5f61140e858286016113ad565b925050602061141f858286016113d7565b9150509250929050565b6114328161129f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561146f578082015181840152602081019050611454565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61149482611438565b61149e8185611442565b93506114ae818560208601611452565b6114b78161147a565b840191505092915050565b5f819050919050565b6114d4816114c2565b82525050565b5f60e0820190506114ed5f83018a611429565b81810360208301526114ff818961148a565b905061150e6040830188611367565b61151b6060830187611367565b6115286080830186611367565b61153560a08301856114cb565b61154260c0830184611367565b98975050505050505050565b5f602082840312156115635761156261138f565b5b5f611570848285016113ad565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f6115bc82611438565b6115c681856115a2565b93506115d6818560208601611452565b6115df8161147a565b840191505092915050565b6115f38161135e565b82525050565b611602816114c2565b82525050565b5f60e083015f83015161161d5f8601826112b0565b506020830151848203602086015261163582826115b2565b915050604083015161164a60408601826115ea565b50606083015161165d60608601826115ea565b50608083015161167060808601826115ea565b5060a083015161168360a08601826115f9565b5060c083015161169660c08601826115ea565b508091505092915050565b5f6116ac8383611608565b905092915050565b5f602082019050919050565b5f6116ca82611579565b6116d48185611583565b9350836020820285016116e685611593565b805f5b85811015611721578484038952815161170285826116a1565b945061170d836116b4565b925060208a019950506001810190506116e9565b50829750879550505050505092915050565b5f6020820190508181035f83015261174b81846116c0565b905092915050565b5f6080820190506117665f8301876114cb565b6117736020830186611367565b8181036040830152611785818561148a565b90506117946060830184611367565b95945050505050565b5f6020820190506117b05f830184611429565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126117d7576117d66117b6565b5b8235905067ffffffffffffffff8111156117f4576117f36117ba565b5b6020830191508360018202830111156118105761180f6117be565b5b9250929050565b5f806020838503121561182d5761182c61138f565b5b5f83013567ffffffffffffffff81111561184a57611849611393565b5b611856858286016117c2565b92509250509250929050565b5f602082840312156118775761187661138f565b5b5f611884848285016113d7565b91505092915050565b5f805f604084860312156118a4576118a361138f565b5b5f6118b1868287016113d7565b935050602084013567ffffffffffffffff8111156118d2576118d1611393565b5b6118de868287016117c2565b92509250509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6119218261135e565b915061192c8361135e565b92508261193c5761193b6118ea565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61197e8261135e565b91506119898361135e565b92508282019050808211156119a1576119a0611947565b5b92915050565b5f6119b18261135e565b91506119bc8361135e565b92508282026119ca8161135e565b915082820484148315176119e1576119e0611947565b5b5092915050565b5f6119f28261135e565b91506119fd8361135e565b925082611a0d57611a0c6118ea565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a5c57607f821691505b602082108103611a6f57611a6e611a18565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611acf602e83611442565b9150611ada82611a75565b604082019050919050565b5f6020820190508181035f830152611afc81611ac3565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f611b3b611b36611b3184611b03565b611b18565b611b0c565b9050919050565b611b4b81611b21565b82525050565b5f602082019050611b645f830184611b42565b92915050565b5f819050919050565b611b84611b7f8261135e565b611b6a565b82525050565b5f8160601b9050919050565b5f611ba082611b8a565b9050919050565b5f611bb182611b96565b9050919050565b611bc9611bc48261129f565b611ba7565b82525050565b5f81905092915050565b828183375f83830152505050565b5f611bf28385611bcf565b9350611bff838584611bd9565b82840190509392505050565b5f611c168288611b73565b602082019150611c268287611b73565b602082019150611c368286611bb8565b601482019150611c47828486611be7565b91508190509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611cdf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ca4565b611ce98683611ca4565b95508019841693508086168417925050509392505050565b5f611d1b611d16611d118461135e565b611b18565b61135e565b9050919050565b5f819050919050565b611d3483611d01565b611d48611d4082611d22565b848454611cb0565b825550505050565b5f90565b611d5c611d50565b611d67818484611d2b565b505050565b5b81811015611d8a57611d7f5f82611d54565b600181019050611d6d565b5050565b601f821115611dcf57611da081611c83565b611da984611c95565b81016020851015611db8578190505b611dcc611dc485611c95565b830182611d6c565b50505b505050565b5f82821c905092915050565b5f611def5f1984600802611dd4565b1980831691505092915050565b5f611e078383611de0565b9150826002028217905092915050565b611e2082611438565b67ffffffffffffffff811115611e3957611e38611c56565b5b611e438254611a45565b611e4e828285611d8e565b5f60209050601f831160018114611e7f575f8415611e6d578287015190505b611e778582611dfc565b865550611ede565b601f198416611e8d86611c83565b5f5b82811015611eb457848901518255600182019150602085019450602081019050611e8f565b86831015611ed15784890151611ecd601f891682611de0565b8355505b6001600288020188555050505b505050505050565b5f604082019050611ef95f8301856114cb565b611f066020830184611367565b9392505050565b7f70726f6f662077696e646f77732068617320636c6f73656400000000000000005f82015250565b5f611f41601883611442565b9150611f4c82611f0d565b602082019050919050565b5f6020820190508181035f830152611f6e81611f35565b9050919050565b7f576f726b207375626d6974206e6f7420636f6d70617461626c652077697468205f8201527f6368616c6c656e67650000000000000000000000000000000000000000000000602082015250565b5f611fcf602983611442565b9150611fda82611f75565b604082019050919050565b5f6020820190508181035f830152611ffc81611fc3565b9050919050565b5f819050919050565b61201d612018826114c2565b612003565b82525050565b5f61202e828561200c565b60208201915061203e8284611b73565b6020820191508190509392505050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f6120a8602483611442565b91506120b38261204e565b604082019050919050565b5f6020820190508181035f8301526120d58161209c565b9050919050565b5f6120e68261135e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361211857612117611947565b5b600182019050919050565b5f61212e8385611442565b935061213b838584611bd9565b6121448361147a565b840190509392505050565b5f60e0820190506121625f83018b611429565b818103602083015261217581898b612123565b90506121846040830188611367565b6121916060830187611367565b61219e6080830186611367565b6121ab60a08301856114cb565b6121b860c0830184611367565b9998505050505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61221f602683611442565b915061222a826121c5565b604082019050919050565b5f6020820190508181035f83015261224c81612213565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f612287602083611442565b915061229282612253565b602082019050919050565b5f6020820190508181035f8301526122b48161227b565b905091905056fea2646970667358221220032f4fb0ab1219d91d75aff0d56aabaa226dd1f7f6a0b457ab66acdf99e4773a64736f6c63430008150033",
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

// GetMinerPosSubmissions is a free data retrieval call binding the contract method 0x54bfd759.
//
// Solidity: function getMinerPosSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowCaller) GetMinerPosSubmissions(opts *bind.CallOpts, addr common.Address) ([]LilypadPowPOWSubmission, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "getMinerPosSubmissions", addr)

	if err != nil {
		return *new([]LilypadPowPOWSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new([]LilypadPowPOWSubmission)).(*[]LilypadPowPOWSubmission)

	return out0, err

}

// GetMinerPosSubmissions is a free data retrieval call binding the contract method 0x54bfd759.
//
// Solidity: function getMinerPosSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowSession) GetMinerPosSubmissions(addr common.Address) ([]LilypadPowPOWSubmission, error) {
	return _Pow.Contract.GetMinerPosSubmissions(&_Pow.CallOpts, addr)
}

// GetMinerPosSubmissions is a free data retrieval call binding the contract method 0x54bfd759.
//
// Solidity: function getMinerPosSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256)[])
func (_Pow *PowCallerSession) GetMinerPosSubmissions(addr common.Address) ([]LilypadPowPOWSubmission, error) {
	return _Pow.Contract.GetMinerPosSubmissions(&_Pow.CallOpts, addr)
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
