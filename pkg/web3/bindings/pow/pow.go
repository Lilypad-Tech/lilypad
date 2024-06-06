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

// PowMetaData contains all meta data concerning the Pow contract.
var PowMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"GenerateChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewPowRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"ValidPOWSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TARGET_DIFFICULTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculate_difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTimeWindow\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"generateChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastChallenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minerSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powSubmissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complete_timestap\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nodeId\",\"type\":\"string\"}],\"name\":\"submitWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerNewPowRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"window_start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040527c5269a8a800000000000000000000000000000000000000000000000000608090815250348015610033575f80fd5b5061005061004561005560201b60201c565b61005c60201b60201c565b61011d565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b608051611ffe61013c5f395f818161041a0152610fd00152611ffe5ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c80638da5cb5b116100ab578063bfb660de1161006f578063bfb660de146102c8578063da8accf9146102d2578063e0d152af146102ee578063f2fde38b1461031e578063f6f851681461033a5761011f565b80638da5cb5b14610248578063a738bf8f14610266578063ab09894514610284578063b28d87ea146102a0578063b681f2fd146102be5761011f565b80634bbe05e4116100f25780634bbe05e4146101ad5780636189f3ac146101e3578063715018a61461021657806378e97925146102205780638129fc1c1461023e5761011f565b8063148c684a146101235780631633da6e146101535780632d68c39714610171578063331bade11461018f575b5f80fd5b61013d600480360381019061013891906111bc565b610358565b60405161014a91906111ff565b60405180910390f35b61015b61036d565b60405161016891906112cf565b60405180910390f35b6101796103f8565b60405161018691906111ff565b60405180910390f35b610197610453565b6040516101a491906111ff565b60405180910390f35b6101c760048036038101906101c29190611319565b610459565b6040516101da9796959493929190611408565b60405180910390f35b6101fd60048036038101906101f891906111bc565b610556565b60405161020d949392919061147c565b60405180910390f35b61021e610608565b005b61022861061b565b60405161023591906111ff565b60405180910390f35b610246610621565b005b610250610756565b60405161025d91906114c6565b60405180910390f35b61026e61077d565b60405161027b91906111ff565b60405180910390f35b61029e60048036038101906102999190611540565b610783565b005b6102a86108df565b6040516102b591906111ff565b60405180910390f35b6102c66108e5565b005b6102d0610935565b005b6102ec60048036038101906102e7919061158b565b61097b565b005b610308600480360381019061030391906115e8565b610f11565b60405161031591906114c6565b60405180910390f35b610338600480360381019061033391906111bc565b610f4c565b005b610342610fce565b60405161034f91906111ff565b60405180910390f35b6002602052805f5260405f205f915090505481565b606060038054806020026020016040519081016040528092919081815260200182805480156103ee57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103a5575b5050505050905090565b5f806015446104079190611640565b605a610413919061169d565b90506064817f000000000000000000000000000000000000000000000000000000000000000061044391906116d0565b61044d9190611711565b91505090565b60075481565b6001602052815f5260405f208181548110610472575f80fd5b905f5260205f2090600702015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546104b79061176e565b80601f01602080910402602001604051908101604052809291908181526020018280546104e39061176e565b801561052e5780601f106105055761010080835404028352916020019161052e565b820191905f5260205f20905b81548152906001019060200180831161051157829003601f168201915b5050505050908060020154908060030154908060040154908060050154908060060154905087565b6004602052805f5260405f205f91509050805f0154908060010154908060020180546105819061176e565b80601f01602080910402602001604051908101604052809291908181526020018280546105ad9061176e565b80156105f85780601f106105cf576101008083540402835291602001916105f8565b820191905f5260205f20905b8154815290600101906020018083116105db57829003601f168201915b5050505050908060030154905084565b610610610ff2565b6106195f611070565b565b60065481565b5f8060159054906101000a900460ff16159050808015610652575060015f60149054906101000a900460ff1660ff16105b80610680575061066130611131565b15801561067f575060015f60149054906101000a900460ff1660ff16145b5b6106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b69061180e565b60405180910390fd5b60015f60146101000a81548160ff021916908360ff16021790555080156106fb5760015f60156101000a81548160ff0219169083151502179055505b8015610753575f8060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161074a919061187a565b60405180910390a15b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60085481565b61078b610935565b5f426007543385856040516020016107a7959493929190611934565b6040516020818303038152906040528051906020012090505f6107c86103f8565b9050604051806080016040528083815260200182815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281525060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f01556020820151816001015560408201518160020190816108929190611b40565b50606082015181600301559050507f496186a9d930bac0744acdcd1d0c054b18283eecbe99d30bda98102d3d06b8ef82826040516108d1929190611c0f565b60405180910390a150505050565b60055481565b6108ed610ff2565b43600781905550601e43610901919061169d565b6008819055507f10cc99616aca050d810fd487c95f968e516d0fa25318530ed50753153d060a1e60405160405180910390a1565b6008544310610979576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097090611c80565b60405180910390fd5b565b610983610935565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f8201548152602001600182015481526020016002820180546109ef9061176e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1b9061176e565b8015610a665780601f10610a3d57610100808354040283529160200191610a66565b820191905f5260205f20905b815481529060010190602001808311610a4957829003601f168201915b5050505050815260200160038201548152505090505f8160600151600754338686604051602001610a9b959493929190611934565b60405160208183030381529060405280519060200120905080825f015114610af8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aef90611d0e565b60405180910390fd5b5f8186604051602001610b0c929190611d4c565b6040516020818303038152906040528051906020012090508260200151815f1c10610b6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6390611de7565b60405180910390fd5b60055f815480929190610b7e90611e05565b91905055505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610c2857600333908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610c7590611e05565b91905055505f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806040518060e001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200188888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200189815260200186606001518152602001428152602001865f015181526020018660200151815250908060018154018082558091505060019003905f5260205f2090600702015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610dd19190611b40565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060155505060405180608001604052805f801b81526020015f815260200160405180602001604052805f81525081526020015f81525060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0155602082015181600101556040820151816002019081610e9b9190611b40565b50606082015181600301559050503373ffffffffffffffffffffffffffffffffffffffff167f5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c86087878a42895f01518a60200151604051610f0096959493929190611e78565b60405180910390a250505050505050565b60038181548110610f20575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610f54610ff2565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610fc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb990611f42565b60405180910390fd5b610fcb81611070565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b610ffa611153565b73ffffffffffffffffffffffffffffffffffffffff16611018610756565b73ffffffffffffffffffffffffffffffffffffffff161461106e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106590611faa565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61118b82611162565b9050919050565b61119b81611181565b81146111a5575f80fd5b50565b5f813590506111b681611192565b92915050565b5f602082840312156111d1576111d061115a565b5b5f6111de848285016111a8565b91505092915050565b5f819050919050565b6111f9816111e7565b82525050565b5f6020820190506112125f8301846111f0565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61124a81611181565b82525050565b5f61125b8383611241565b60208301905092915050565b5f602082019050919050565b5f61127d82611218565b6112878185611222565b935061129283611232565b805f5b838110156112c25781516112a98882611250565b97506112b483611267565b925050600181019050611295565b5085935050505092915050565b5f6020820190508181035f8301526112e78184611273565b905092915050565b6112f8816111e7565b8114611302575f80fd5b50565b5f81359050611313816112ef565b92915050565b5f806040838503121561132f5761132e61115a565b5b5f61133c858286016111a8565b925050602061134d85828601611305565b9150509250929050565b61136081611181565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561139d578082015181840152602081019050611382565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6113c282611366565b6113cc8185611370565b93506113dc818560208601611380565b6113e5816113a8565b840191505092915050565b5f819050919050565b611402816113f0565b82525050565b5f60e08201905061141b5f83018a611357565b818103602083015261142d81896113b8565b905061143c60408301886111f0565b61144960608301876111f0565b61145660808301866111f0565b61146360a08301856113f9565b61147060c08301846111f0565b98975050505050505050565b5f60808201905061148f5f8301876113f9565b61149c60208301866111f0565b81810360408301526114ae81856113b8565b90506114bd60608301846111f0565b95945050505050565b5f6020820190506114d95f830184611357565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f840112611500576114ff6114df565b5b8235905067ffffffffffffffff81111561151d5761151c6114e3565b5b602083019150836001820283011115611539576115386114e7565b5b9250929050565b5f80602083850312156115565761155561115a565b5b5f83013567ffffffffffffffff8111156115735761157261115e565b5b61157f858286016114eb565b92509250509250929050565b5f805f604084860312156115a2576115a161115a565b5b5f6115af86828701611305565b935050602084013567ffffffffffffffff8111156115d0576115cf61115e565b5b6115dc868287016114eb565b92509250509250925092565b5f602082840312156115fd576115fc61115a565b5b5f61160a84828501611305565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61164a826111e7565b9150611655836111e7565b92508261166557611664611613565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116a7826111e7565b91506116b2836111e7565b92508282019050808211156116ca576116c9611670565b5b92915050565b5f6116da826111e7565b91506116e5836111e7565b92508282026116f3816111e7565b9150828204841483151761170a57611709611670565b5b5092915050565b5f61171b826111e7565b9150611726836111e7565b92508261173657611735611613565b5b828204905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061178557607f821691505b60208210810361179857611797611741565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6117f8602e83611370565b91506118038261179e565b604082019050919050565b5f6020820190508181035f830152611825816117ec565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61186461185f61185a8461182c565b611841565b611835565b9050919050565b6118748161184a565b82525050565b5f60208201905061188d5f83018461186b565b92915050565b5f819050919050565b6118ad6118a8826111e7565b611893565b82525050565b5f8160601b9050919050565b5f6118c9826118b3565b9050919050565b5f6118da826118bf565b9050919050565b6118f26118ed82611181565b6118d0565b82525050565b5f81905092915050565b828183375f83830152505050565b5f61191b83856118f8565b9350611928838584611902565b82840190509392505050565b5f61193f828861189c565b60208201915061194f828761189c565b60208201915061195f82866118e1565b601482019150611970828486611910565b91508190509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611a087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826119cd565b611a1286836119cd565b95508019841693508086168417925050509392505050565b5f611a44611a3f611a3a846111e7565b611841565b6111e7565b9050919050565b5f819050919050565b611a5d83611a2a565b611a71611a6982611a4b565b8484546119d9565b825550505050565b5f90565b611a85611a79565b611a90818484611a54565b505050565b5b81811015611ab357611aa85f82611a7d565b600181019050611a96565b5050565b601f821115611af857611ac9816119ac565b611ad2846119be565b81016020851015611ae1578190505b611af5611aed856119be565b830182611a95565b50505b505050565b5f82821c905092915050565b5f611b185f1984600802611afd565b1980831691505092915050565b5f611b308383611b09565b9150826002028217905092915050565b611b4982611366565b67ffffffffffffffff811115611b6257611b6161197f565b5b611b6c825461176e565b611b77828285611ab7565b5f60209050601f831160018114611ba8575f8415611b96578287015190505b611ba08582611b25565b865550611c07565b601f198416611bb6866119ac565b5f5b82811015611bdd57848901518255600182019150602085019450602081019050611bb8565b86831015611bfa5784890151611bf6601f891682611b09565b8355505b6001600288020188555050505b505050505050565b5f604082019050611c225f8301856113f9565b611c2f60208301846111f0565b9392505050565b7f70726f6f662077696e646f77732068617320636c6f73656400000000000000005f82015250565b5f611c6a601883611370565b9150611c7582611c36565b602082019050919050565b5f6020820190508181035f830152611c9781611c5e565b9050919050565b7f576f726b207375626d6974206e6f7420636f6d70617461626c652077697468205f8201527f6368616c6c656e67650000000000000000000000000000000000000000000000602082015250565b5f611cf8602983611370565b9150611d0382611c9e565b604082019050919050565b5f6020820190508181035f830152611d2581611cec565b9050919050565b5f819050919050565b611d46611d41826113f0565b611d2c565b82525050565b5f611d578285611d35565b602082019150611d67828461189c565b6020820191508190509392505050565b7f576f726b20646f6573206e6f74206d65657420646966666963756c74792074615f8201527f7267657400000000000000000000000000000000000000000000000000000000602082015250565b5f611dd1602483611370565b9150611ddc82611d77565b604082019050919050565b5f6020820190508181035f830152611dfe81611dc5565b9050919050565b5f611e0f826111e7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e4157611e40611670565b5b600182019050919050565b5f611e578385611370565b9350611e64838584611902565b611e6d836113a8565b840190509392505050565b5f60a0820190508181035f830152611e9181888a611e4c565b9050611ea060208301876111f0565b611ead60408301866111f0565b611eba60608301856113f9565b611ec760808301846111f0565b979650505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611f2c602683611370565b9150611f3782611ed2565b604082019050919050565b5f6020820190508181035f830152611f5981611f20565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611f94602083611370565b9150611f9f82611f60565b602082019050919050565b5f6020820190508181035f830152611fc181611f88565b905091905056fea2646970667358221220e66d961328264141eb34b4e4f3e4360ab4e09b23e6a36ea98882788925a681ae64736f6c63430008150033",
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

// TARGETDIFFICULTY is a free data retrieval call binding the contract method 0xf6f85168.
//
// Solidity: function TARGET_DIFFICULTY() view returns(uint256)
func (_Pow *PowCaller) TARGETDIFFICULTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "TARGET_DIFFICULTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETDIFFICULTY is a free data retrieval call binding the contract method 0xf6f85168.
//
// Solidity: function TARGET_DIFFICULTY() view returns(uint256)
func (_Pow *PowSession) TARGETDIFFICULTY() (*big.Int, error) {
	return _Pow.Contract.TARGETDIFFICULTY(&_Pow.CallOpts)
}

// TARGETDIFFICULTY is a free data retrieval call binding the contract method 0xf6f85168.
//
// Solidity: function TARGET_DIFFICULTY() view returns(uint256)
func (_Pow *PowCallerSession) TARGETDIFFICULTY() (*big.Int, error) {
	return _Pow.Contract.TARGETDIFFICULTY(&_Pow.CallOpts)
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

// MinerSubmissionCount is a free data retrieval call binding the contract method 0x148c684a.
//
// Solidity: function minerSubmissionCount(address ) view returns(uint256)
func (_Pow *PowCaller) MinerSubmissionCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "minerSubmissionCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerSubmissionCount is a free data retrieval call binding the contract method 0x148c684a.
//
// Solidity: function minerSubmissionCount(address ) view returns(uint256)
func (_Pow *PowSession) MinerSubmissionCount(arg0 common.Address) (*big.Int, error) {
	return _Pow.Contract.MinerSubmissionCount(&_Pow.CallOpts, arg0)
}

// MinerSubmissionCount is a free data retrieval call binding the contract method 0x148c684a.
//
// Solidity: function minerSubmissionCount(address ) view returns(uint256)
func (_Pow *PowCallerSession) MinerSubmissionCount(arg0 common.Address) (*big.Int, error) {
	return _Pow.Contract.MinerSubmissionCount(&_Pow.CallOpts, arg0)
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
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestap, uint256 complete_timestap, bytes32 challenge, uint256 difficulty)
func (_Pow *PowCaller) PowSubmissions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress    common.Address
	NodeId           string
	Nonce            *big.Int
	StartTimestap    *big.Int
	CompleteTimestap *big.Int
	Challenge        [32]byte
	Difficulty       *big.Int
}, error) {
	var out []interface{}
	err := _Pow.contract.Call(opts, &out, "powSubmissions", arg0, arg1)

	outstruct := new(struct {
		WalletAddress    common.Address
		NodeId           string
		Nonce            *big.Int
		StartTimestap    *big.Int
		CompleteTimestap *big.Int
		Challenge        [32]byte
		Difficulty       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NodeId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTimestap = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CompleteTimestap = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Challenge = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestap, uint256 complete_timestap, bytes32 challenge, uint256 difficulty)
func (_Pow *PowSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress    common.Address
	NodeId           string
	Nonce            *big.Int
	StartTimestap    *big.Int
	CompleteTimestap *big.Int
	Challenge        [32]byte
	Difficulty       *big.Int
}, error) {
	return _Pow.Contract.PowSubmissions(&_Pow.CallOpts, arg0, arg1)
}

// PowSubmissions is a free data retrieval call binding the contract method 0x4bbe05e4.
//
// Solidity: function powSubmissions(address , uint256 ) view returns(address walletAddress, string nodeId, uint256 nonce, uint256 start_timestap, uint256 complete_timestap, bytes32 challenge, uint256 difficulty)
func (_Pow *PowCallerSession) PowSubmissions(arg0 common.Address, arg1 *big.Int) (struct {
	WalletAddress    common.Address
	NodeId           string
	Nonce            *big.Int
	StartTimestap    *big.Int
	CompleteTimestap *big.Int
	Challenge        [32]byte
	Difficulty       *big.Int
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
func (_Pow *PowFilterer) FilterValidPOWSubmitted(opts *bind.FilterOpts, walletAddress []common.Address) (*PowValidPOWSubmittedIterator, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _Pow.contract.FilterLogs(opts, "ValidPOWSubmitted", walletAddressRule)
	if err != nil {
		return nil, err
	}
	return &PowValidPOWSubmittedIterator{contract: _Pow.contract, event: "ValidPOWSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidPOWSubmitted is a free log subscription operation binding the contract event 0x5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c860.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) WatchValidPOWSubmitted(opts *bind.WatchOpts, sink chan<- *PowValidPOWSubmitted, walletAddress []common.Address) (event.Subscription, error) {

	var walletAddressRule []interface{}
	for _, walletAddressItem := range walletAddress {
		walletAddressRule = append(walletAddressRule, walletAddressItem)
	}

	logs, sub, err := _Pow.contract.WatchLogs(opts, "ValidPOWSubmitted", walletAddressRule)
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

// ParseValidPOWSubmitted is a log parse operation binding the contract event 0x5ea61d67e3236a79ab8c3c29f810e2c408b0a6b1e19f4e0da2fcd6362129c860.
//
// Solidity: event ValidPOWSubmitted(address indexed walletAddress, string nodeId, uint256 nonce, uint256 timestamp, bytes32 challenge, uint256 difficulty)
func (_Pow *PowFilterer) ParseValidPOWSubmitted(log types.Log) (*PowValidPOWSubmitted, error) {
	event := new(PowValidPOWSubmitted)
	if err := _Pow.contract.UnpackLog(event, "ValidPOWSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
