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
	ReportHashrates   *big.Int
}

// PowMetaData contains all meta data concerning the Pow contract.
var PowMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPowRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"report_hashrates\",\"type\":\"uint256\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"calculate_difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"change_difficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTimeWindow\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMinerPowSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getMinerPowSubmissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"report_hashrates\",\"type\":\"uint256\"}],\"internalType\":\"structLilypadPow.POWSubmission[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"report_hashrates\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"report_hashrates\",\"type\":\"uint256\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506125868061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c80638da5cb5b116100b6578063b681f2fd1161007a578063b681f2fd14610336578063bb52c13714610340578063bfb660de1461035e578063c2c434da14610368578063e0d152af14610398578063f2fde38b146103c857610140565b80638da5cb5b146102a4578063a738bf8f146102c2578063ab098945146102e0578063adf0047c146102fc578063b28d87ea1461031857610140565b80634bbe05e4116101085780634bbe05e4146101ea5780636189f3ac14610221578063715018a61461025457806378e979251461025e5780638129fc1c1461027c5780638b2db16e1461028657610140565b80631633da6e146101445780631bb0d808146101625780632c23c10c146101925780632d68c397146101ae578063331bade1146101cc575b5f80fd5b61014c6103e4565b60405161015991906114ff565b60405180910390f35b61017c60048036038101906101779190611551565b61046f565b6040516101899190611594565b60405180910390f35b6101ac60048036038101906101a79190611638565b6104b8565b005b6101b66109c8565b6040516101c39190611594565b60405180910390f35b6101d4610a05565b6040516101e19190611594565b60405180910390f35b61020460048036038101906101ff91906116a9565b610a0b565b604051610218989796959493929190611798565b60405180910390f35b61023b60048036038101906102369190611551565b610b0e565b60405161024b949392919061181b565b60405180910390f35b61025c610bc0565b005b610266610bd3565b6040516102739190611594565b60405180910390f35b610284610bd9565b005b61028e610d35565b60405161029b9190611594565b60405180910390f35b6102ac610d3b565b6040516102b99190611865565b60405180910390f35b6102ca610d63565b6040516102d79190611594565b60405180910390f35b6102fa60048036038101906102f5919061187e565b610d69565b005b610316600480360381019061031191906118c9565b610ec5565b005b610320610ed7565b60405161032d9190611594565b60405180910390f35b61033e610edd565b005b610348610f2d565b6040516103559190611594565b60405180910390f35b610366610f39565b005b610382600480360381019061037d9190611551565b610f7f565b60405161038f9190611ac2565b60405180910390f35b6103b260048036038101906103ad91906118c9565b611139565b6040516103bf9190611865565b60405180910390f35b6103e260048036038101906103dd9190611551565b611174565b005b6060606780548060200260200160405190810160405280929190818152602001828054801561046557602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161041c575b5050505050905090565b5f60665f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805490509050919050565b6104c0610f39565b5f60685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f82015481526020016001820154815260200160028201805461052c90611b0f565b80601f016020809104026020016040519081016040528092919081815260200182805461055890611b0f565b80156105a35780601f1061057a576101008083540402835291602001916105a3565b820191905f5260205f20905b81548152906001019060200180831161058657829003601f168201915b5050505050815260200160038201548152505090505f8160600151606b543387876040516020016105d8959493929190611be0565b60405160208183030381529060405280519060200120905080825f015114610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062c90611c9b565b60405180910390fd5b5f8187604051602001610649929190611cd9565b6040516020818303038152906040528051906020012090508260200151815f1c106106a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a090611d74565b60405180910390fd5b60695f8154809291906106bb90611dbf565b91905055505f60665f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f81805490500361076c57606733908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b806040518061010001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018a815260200186606001518152602001428152602001865f015181526020018660200151815260200187815250908060018154018082558091505060019003905f5260205f2090600802015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161088a9190611fd0565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155505060405180608001604052805f801b81526020015f815260200160405180602001604052805f81525081526020015f81525060685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015560208201518160010155604082015181600201908161095e9190611fd0565b50606082015181600301559050507ff6feb9ff838f9c6cb818374d0e0c8191e3653e218de80e7f84cce1560b758c653388888b8860600151428a5f01518b602001518d6040516109b6999897969594939291906120cb565b60405180910390a15050505050505050565b5f806015446109d7919061217d565b605a6109e391906121ad565b90506064816065546109f591906121e0565b6109ff9190612221565b91505090565b606b5481565b6066602052815f5260405f208181548110610a24575f80fd5b905f5260205f2090600802015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610a6990611b0f565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9590611b0f565b8015610ae05780601f10610ab757610100808354040283529160200191610ae0565b820191905f5260205f20905b815481529060010190602001808311610ac357829003601f168201915b5050505050908060020154908060030154908060040154908060050154908060060154908060070154905088565b6068602052805f5260405f205f91509050805f015490806001015490806002018054610b3990611b0f565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6590611b0f565b8015610bb05780601f10610b8757610100808354040283529160200191610bb0565b820191905f5260205f20905b815481529060010190602001808311610b9357829003601f168201915b5050505050908060030154905084565b610bc86111f6565b610bd15f611274565b565b606a5481565b5f8060019054906101000a900460ff16159050808015610c08575060015f8054906101000a900460ff1660ff16105b80610c345750610c1730611337565b158015610c33575060015f8054906101000a900460ff1660ff16145b5b610c73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6a906122c1565b60405180910390fd5b60015f806101000a81548160ff021916908360ff1602179055508015610cae5760015f60016101000a81548160ff0219169083151502179055505b610cb6611359565b7c149a6a30000000000000000000000000000000000000000000000000006065819055508015610d32575f8060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610d299190612324565b60405180910390a15b50565b60655481565b5f60335f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606c5481565b610d71610f39565b5f42606b54338585604051602001610d8d959493929190611be0565b6040516020818303038152906040528051906020012090505f610dae6109c8565b9050604051806080016040528083815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060685f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0155602082015181600101556040820151816002019081610e789190611fd0565b50606082015181600301559050507f496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef8282604051610eb792919061233d565b60405180910390a150505050565b610ecd6111f6565b8060658190555050565b60695481565b610ee56111f6565b43606b81905550601e43610ef991906121ad565b606c819055507f10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e60405160405180910390a1565b5f606780549050905090565b606c544310610f7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f74906123ae565b60405180910390fd5b565b606060665f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b8282101561112e578382905f5260205f209060080201604051806101000160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461106390611b0f565b80601f016020809104026020016040519081016040528092919081815260200182805461108f90611b0f565b80156110da5780601f106110b1576101008083540402835291602001916110da565b820191905f5260205f20905b8154815290600101906020018083116110bd57829003601f168201915b50505050508152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505081526020019060010190610fdd565b505050509050919050565b60678181548110611148575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61117c6111f6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e19061243c565b60405180910390fd5b6111f381611274565b50565b6111fe6113b1565b73ffffffffffffffffffffffffffffffffffffffff1661121c610d3b565b73ffffffffffffffffffffffffffffffffffffffff1614611272576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611269906124a4565b60405180910390fd5b565b5f60335f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160335f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f60019054906101000a900460ff166113a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139e90612532565b60405180910390fd5b6113af6113b8565b565b5f33905090565b5f60019054906101000a900460ff16611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fd90612532565b60405180910390fd5b6114166114116113b1565b611274565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61146a82611441565b9050919050565b61147a81611460565b82525050565b5f61148b8383611471565b60208301905092915050565b5f602082019050919050565b5f6114ad82611418565b6114b78185611422565b93506114c283611432565b805f5b838110156114f25781516114d98882611480565b97506114e483611497565b9250506001810190506114c5565b5085935050505092915050565b5f6020820190508181035f83015261151781846114a3565b905092915050565b5f80fd5b5f80fd5b61153081611460565b811461153a575f80fd5b50565b5f8135905061154b81611527565b92915050565b5f602082840312156115665761156561151f565b5b5f6115738482850161153d565b91505092915050565b5f819050919050565b61158e8161157c565b82525050565b5f6020820190506115a75f830184611585565b92915050565b6115b68161157c565b81146115c0575f80fd5b50565b5f813590506115d1816115ad565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126115f8576115f76115d7565b5b8235905067ffffffffffffffff811115611615576116146115db565b5b602083019150836001820283011115611631576116306115df565b5b9250929050565b5f805f80606085870312156116505761164f61151f565b5b5f61165d878288016115c3565b945050602085013567ffffffffffffffff81111561167e5761167d611523565b5b61168a878288016115e3565b9350935050604061169d878288016115c3565b91505092959194509250565b5f80604083850312156116bf576116be61151f565b5b5f6116cc8582860161153d565b92505060206116dd858286016115c3565b9150509250929050565b6116f081611460565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561172d578082015181840152602081019050611712565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611752826116f6565b61175c8185611700565b935061176c818560208601611710565b61177581611738565b840191505092915050565b5f819050919050565b61179281611780565b82525050565b5f610100820190506117ac5f83018b6116e7565b81810360208301526117be818a611748565b90506117cd6040830189611585565b6117da6060830188611585565b6117e76080830187611585565b6117f460a0830186611789565b61180160c0830185611585565b61180e60e0830184611585565b9998505050505050505050565b5f60808201905061182e5f830187611789565b61183b6020830186611585565b818103604083015261184d8185611748565b905061185c6060830184611585565b95945050505050565b5f6020820190506118785f8301846116e7565b92915050565b5f80602083850312156118945761189361151f565b5b5f83013567ffffffffffffffff8111156118b1576118b0611523565b5b6118bd858286016115e3565b92509250509250929050565b5f602082840312156118de576118dd61151f565b5b5f6118eb848285016115c3565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f611937826116f6565b611941818561191d565b9350611951818560208601611710565b61195a81611738565b840191505092915050565b61196e8161157c565b82525050565b61197d81611780565b82525050565b5f61010083015f8301516119995f860182611471565b50602083015184820360208601526119b1828261192d565b91505060408301516119c66040860182611965565b5060608301516119d96060860182611965565b5060808301516119ec6080860182611965565b5060a08301516119ff60a0860182611974565b5060c0830151611a1260c0860182611965565b5060e0830151611a2560e0860182611965565b508091505092915050565b5f611a3b8383611983565b905092915050565b5f602082019050919050565b5f611a59826118f4565b611a6381856118fe565b935083602082028501611a758561190e565b805f5b85811015611ab05784840389528151611a918582611a30565b9450611a9c83611a43565b925060208a01995050600181019050611a78565b50829750879550505050505092915050565b5f6020820190508181035f830152611ada8184611a4f565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611b2657607f821691505b602082108103611b3957611b38611ae2565b5b50919050565b5f819050919050565b611b59611b548261157c565b611b3f565b82525050565b5f8160601b9050919050565b5f611b7582611b5f565b9050919050565b5f611b8682611b6b565b9050919050565b611b9e611b9982611460565b611b7c565b82525050565b5f81905092915050565b828183375f83830152505050565b5f611bc78385611ba4565b9350611bd4838584611bae565b82840190509392505050565b5f611beb8288611b48565b602082019150611bfb8287611b48565b602082019150611c0b8286611b8d565b601482019150611c1c828486611bbc565b91508190509695505050505050565b7f576f726b207375626d6974206e6f7420636f6d70617461626c652077697468205f8201527f6368616c6c656e67650000000000000000000000000000000000000000000000602082015250565b5f611c85602983611700565b9150611c9082611c2b565b604082019050919050565b5f6020820190508181035f830152611cb281611c79565b9050919050565b5f819050919050565b611cd3611cce82611780565b611cb9565b82525050565b5f611ce48285611cc2565b602082019150611cf48284611b48565b6020820191508190509392505050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f611d5e602483611700565b9150611d6982611d04565b604082019050919050565b5f6020820190508181035f830152611d8b81611d52565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611dc98261157c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611dfb57611dfa611d92565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611e8f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611e54565b611e998683611e54565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611ed4611ecf611eca8461157c565b611eb1565b61157c565b9050919050565b5f819050919050565b611eed83611eba565b611f01611ef982611edb565b848454611e60565b825550505050565b5f90565b611f15611f09565b611f20818484611ee4565b505050565b5b81811015611f4357611f385f82611f0d565b600181019050611f26565b5050565b601f821115611f8857611f5981611e33565b611f6284611e45565b81016020851015611f71578190505b611f85611f7d85611e45565b830182611f25565b50505b505050565b5f82821c905092915050565b5f611fa85f1984600802611f8d565b1980831691505092915050565b5f611fc08383611f99565b9150826002028217905092915050565b611fd9826116f6565b67ffffffffffffffff811115611ff257611ff1611e06565b5b611ffc8254611b0f565b612007828285611f47565b5f60209050601f831160018114612038575f8415612026578287015190505b6120308582611fb5565b865550612097565b601f19841661204686611e33565b5f5b8281101561206d57848901518255600182019150602085019450602081019050612048565b8683101561208a5784890151612086601f891682611f99565b8355505b6001600288020188555050505b505050505050565b5f6120aa8385611700565b93506120b7838584611bae565b6120c083611738565b840190509392505050565b5f610100820190506120df5f83018c6116e7565b81810360208301526120f2818a8c61209f565b90506121016040830189611585565b61210e6060830188611585565b61211b6080830187611585565b61212860a0830186611789565b61213560c0830185611585565b61214260e0830184611585565b9a9950505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6121878261157c565b91506121928361157c565b9250826121a2576121a1612150565b5b828206905092915050565b5f6121b78261157c565b91506121c28361157c565b92508282019050808211156121da576121d9611d92565b5b92915050565b5f6121ea8261157c565b91506121f58361157c565b92508282026122038161157c565b9150828204841483151761221a57612219611d92565b5b5092915050565b5f61222b8261157c565b91506122368361157c565b92508261224657612245612150565b5b828204905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6122ab602e83611700565b91506122b682612251565b604082019050919050565b5f6020820190508181035f8301526122d88161229f565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f61230e612309612304846122df565b611eb1565b6122e8565b9050919050565b61231e816122f4565b82525050565b5f6020820190506123375f830184612315565b92915050565b5f6040820190506123505f830185611789565b61235d6020830184611585565b9392505050565b7f70726f6f662077696e646f77732068617320636c6f73656400000000000000005f82015250565b5f612398601883611700565b91506123a382612364565b602082019050919050565b5f6020820190508181035f8301526123c58161238c565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612426602683611700565b9150612431826123cc565b604082019050919050565b5f6020820190508181035f8301526124538161241a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61248e602083611700565b91506124998261245a565b602082019050919050565b5f6020820190508181035f8301526124bb81612482565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f7420695f8201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b5f61251c602b83611700565b9150612527826124c2565b604082019050919050565b5f6020820190508181035f83015261254981612510565b905091905056fea26469706673582212207939612bc798117c27b9fc32c8b5186d731659dd091802330d4f10c786e00f8964736f6c63430008150033",
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
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256,uint256)[])
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
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256,uint256)[])
func (_Pow *PowSession) GetMinerPowSubmissions(addr common.Address) ([]LilypadPowPOWSubmission, error) {
	return _Pow.Contract.GetMinerPowSubmissions(&_Pow.CallOpts, addr)
}

// GetMinerPowSubmissions is a free data retrieval call binding the contract method 0xc2c434da.
//
// Solidity: function getMinerPowSubmissions(address addr) view returns((address,string,uint256,uint256,uint256,bytes32,uint256,uint256)[])
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
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
func (_Pow *PowCaller) PowSubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
	ReportHashrates   *big.Int
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
		ReportHashrates   *big.Int
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
	outstruct.ReportHashrates = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
func (_Pow *PowSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
	ReportHashrates   *big.Int
}, error) {
	return _Pow.Contract.PowSubmissions(&_Pow.CallOpts, arg0, arg1)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
func (_Pow *PowCallerSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress     common.Address
	NodeId            string
	Nonce             *big.Int
	StartTimestamp    *big.Int
	CompleteTimestamp *big.Int
	Challenge         [32]byte
	Difficulty        *big.Int
	ReportHashrates   *big.Int
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

// SubmitWork is a paid mutator transaction binding the contract method 0x2c23c10c.
//
// Solidity: function submitWork(uint256 nonce, string nodeId, uint256 report_hashrates) returns()
func (_Pow *PowTransactor) SubmitWork(opts *bind.TransactOpts, nonce *big.Int, nodeId string, report_hashrates *big.Int) (*types.Transaction, error) {
	return _Pow.contract.Transact(opts, "submitWork", nonce, nodeId, report_hashrates)
}

// SubmitWork is a paid mutator transaction binding the contract method 0x2c23c10c.
//
// Solidity: function submitWork(uint256 nonce, string nodeId, uint256 report_hashrates) returns()
func (_Pow *PowSession) SubmitWork(nonce *big.Int, nodeId string, report_hashrates *big.Int) (*types.Transaction, error) {
	return _Pow.Contract.SubmitWork(&_Pow.TransactOpts, nonce, nodeId, report_hashrates)
}

// SubmitWork is a paid mutator transaction binding the contract method 0x2c23c10c.
//
// Solidity: function submitWork(uint256 nonce, string nodeId, uint256 report_hashrates) returns()
func (_Pow *PowTransactorSession) SubmitWork(nonce *big.Int, nodeId string, report_hashrates *big.Int) (*types.Transaction, error) {
	return _Pow.Contract.SubmitWork(&_Pow.TransactOpts, nonce, nodeId, report_hashrates)
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
	ReportHashrates   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidPOWSubmitted is a free log retrieval operation binding the contract event 0xf6feb9ff838f9c6cb818374d0e0c8191e3653e218de80e7f84cce1560b758c65.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
func (_Pow *PowFilterer) FilterValidPOWSubmitted(opts *bind.FilterOpts) (*PowValidPOWSubmittedIterator, error) {

	logs, sub, err := _Pow.contract.FilterLogs(opts, "ValidPOWSubmitted")
	if err != nil {
		return nil, err
	}
	return &PowValidPOWSubmittedIterator{contract: _Pow.contract, event: "ValidPOWSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidPOWSubmitted is a free log subscription operation binding the contract event 0xf6feb9ff838f9c6cb818374d0e0c8191e3653e218de80e7f84cce1560b758c65.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
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

// ParseValidPOWSubmitted is a log parse operation binding the contract event 0xf6feb9ff838f9c6cb818374d0e0c8191e3653e218de80e7f84cce1560b758c65.
//
// Solidity: event ValidPOWSubmitted(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestamp, uint256 complete_timestamp, bytes32 challenge, uint256 difficulty, uint256 report_hashrates)
func (_Pow *PowFilterer) ParseValidPOWSubmitted(log types.Log) (*PowValidPOWSubmitted, error) {
	event := new(PowValidPOWSubmitted)
	if err := _Pow.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
