// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package users

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

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress common.Address
	MetadataCID string
	Url         string
	Roles       []uint8
}

// UsersMetaData contains all meta data concerning the Users contract.
var UsersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612002806101065f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063a15dcc8a11610064578063a15dcc8a146100f7578063a7f96f0614610127578063aeb5ec0114610143578063ebbbca271461015f578063f2fde38b1461018f57610091565b80636f77926b14610095578063715018a6146100c55780638129fc1c146100cf5780638da5cb5b146100d9575b5f80fd5b6100af60048036038101906100aa91906111cf565b6101ab565b6040516100bc919061141c565b60405180910390f35b6100cd6103ff565b005b6100d7610412565b005b6100e1610547565b6040516100ee919061144b565b60405180910390f35b610111600480360381019061010c9190611487565b61056e565b60405161011e919061155a565b60405180910390f35b610141600480360381019061013c9190611487565b61062e565b005b61015d60048036038101906101589190611487565b610932565b005b6101796004803603810190610174919061176a565b610b2d565b604051610186919061141c565b60405180910390f35b6101a960048036038101906101a491906111cf565b610c48565b005b6101b361105b565b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461025f9061183b565b80601f016020809104026020016040519081016040528092919081815260200182805461028b9061183b565b80156102d65780601f106102ad576101008083540402835291602001916102d6565b820191905f5260205f20905b8154815290600101906020018083116102b957829003601f168201915b505050505081526020016002820180546102ef9061183b565b80601f016020809104026020016040519081016040528092919081815260200182805461031b9061183b565b80156103665780601f1061033d57610100808354040283529160200191610366565b820191905f5260205f20905b81548152906001019060200180831161034957829003601f168201915b50505050508152602001600382018054806020026020016040519081016040528092919081815260200182805480156103ef57602002820191905f5260205f20905f905b82829054906101000a900460ff1660038111156103ca576103c96112bc565b5b815260200190600101906020825f010492830192600103820291508084116103aa5790505b5050505050815250509050919050565b610407610cca565b6104105f610d48565b565b5f8060159054906101000a900460ff16159050808015610443575060015f60149054906101000a900460ff1660ff16105b80610471575061045230610e09565b158015610470575060015f60149054906101000a900460ff1660ff16145b5b6104b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a7906118eb565b60405180910390fd5b60015f60146101000a81548160ff021916908360ff16021790555080156104ec5760015f60156101000a81548160ff0219169083151502179055505b8015610544575f8060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161053b9190611957565b60405180910390a15b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060025f836003811115610586576105856112bc565b5b6003811115610598576105976112bc565b5b81526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561062257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116105d9575b50505050509050919050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036106fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f1906119ba565b60405180910390fd5b5f806107068332610e2b565b915091508061074a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074190611a22565b60405180910390fd5b5f8290505b600160025f866003811115610767576107666112bc565b5b6003811115610779576107786112bc565b5b81526020019081526020015f20805490506107949190611a76565b8110156108b55760025f8560038111156107b1576107b06112bc565b5b60038111156107c3576107c26112bc565b5b81526020019081526020015f206001826107dd9190611aa9565b815481106107ee576107ed611adc565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660025f86600381111561082c5761082b6112bc565b5b600381111561083e5761083d6112bc565b5b81526020019081526020015f20828154811061085d5761085c611adc565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806108ad90611b09565b91505061074f565b5060025f8460038111156108cc576108cb6112bc565b5b60038111156108de576108dd6112bc565b5b81526020019081526020015f208054806108fb576108fa611b50565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f5906119ba565b60405180910390fd5b5f610a098232610e2b565b9150508015610a4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4490611bed565b60405180910390fd5b610a578232610f41565b610a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8d90611c55565b60405180910390fd5b60025f836003811115610aac57610aab6112bc565b5b6003811115610abe57610abd6112bc565b5b81526020019081526020015f2032908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610b3561105b565b5f60405180608001604052803273ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018481525090508060015f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610c069190611e07565b506040820151816002019081610c1c9190611e07565b506060820151816003019080519060200190610c39929190611098565b50905050809150509392505050565b610c50610cca565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb590611f46565b60405180910390fd5b610cc781610d48565b50565b610cd2611054565b73ffffffffffffffffffffffffffffffffffffffff16610cf0610547565b73ffffffffffffffffffffffffffffffffffffffff1614610d46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3d90611fae565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f805f805f5b60025f886003811115610e4757610e466112bc565b5b6003811115610e5957610e586112bc565b5b81526020019081526020015f2080549050811015610f31578573ffffffffffffffffffffffffffffffffffffffff1660025f896003811115610e9e57610e9d6112bc565b5b6003811115610eb057610eaf6112bc565b5b81526020019081526020015f208281548110610ecf57610ece611adc565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610f1e5780925060019150610f31565b8080610f2990611b09565b915050610e31565b5081819350935050509250929050565b5f805f90505f5b60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018054905081101561104957846003811115610fa657610fa56112bc565b5b60015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018281548110610ff857610ff7611adc565b5b905f5260205f2090602091828204019190069054906101000a900460ff166003811115611028576110276112bc565b5b036110365760019150611049565b808061104190611b09565b915050610f48565b508091505092915050565b5f33905090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001606081525090565b828054828255905f5260205f2090601f01602090048101928215611138579160200282015f5b8382111561110a57835183826101000a81548160ff021916908360038111156110ea576110e96112bc565b5b021790555092602001926001016020815f010492830192600103026110be565b80156111365782816101000a81549060ff02191690556001016020815f0104928301926001030261110a565b505b5090506111459190611149565b5090565b5b80821115611160575f815f90555060010161114a565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61119e82611175565b9050919050565b6111ae81611194565b81146111b8575f80fd5b50565b5f813590506111c9816111a5565b92915050565b5f602082840312156111e4576111e361116d565b5b5f6111f1848285016111bb565b91505092915050565b61120381611194565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611240578082015181840152602081019050611225565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61126582611209565b61126f8185611213565b935061127f818560208601611223565b6112888161124b565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600481106112fa576112f96112bc565b5b50565b5f81905061130a826112e9565b919050565b5f611319826112fd565b9050919050565b6113298161130f565b82525050565b5f61133a8383611320565b60208301905092915050565b5f602082019050919050565b5f61135c82611293565b611366818561129d565b9350611371836112ad565b805f5b838110156113a1578151611388888261132f565b975061139383611346565b925050600181019050611374565b5085935050505092915050565b5f608083015f8301516113c35f8601826111fa565b50602083015184820360208601526113db828261125b565b915050604083015184820360408601526113f5828261125b565b9150506060830151848203606086015261140f8282611352565b9150508091505092915050565b5f6020820190508181035f83015261143481846113ae565b905092915050565b61144581611194565b82525050565b5f60208201905061145e5f83018461143c565b92915050565b60048110611470575f80fd5b50565b5f8135905061148181611464565b92915050565b5f6020828403121561149c5761149b61116d565b5b5f6114a984828501611473565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6114e683836111fa565b60208301905092915050565b5f602082019050919050565b5f611508826114b2565b61151281856114bc565b935061151d836114cc565b805f5b8381101561154d57815161153488826114db565b975061153f836114f2565b925050600181019050611520565b5085935050505092915050565b5f6020820190508181035f83015261157281846114fe565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6115b88261124b565b810181811067ffffffffffffffff821117156115d7576115d6611582565b5b80604052505050565b5f6115e9611164565b90506115f582826115af565b919050565b5f67ffffffffffffffff82111561161457611613611582565b5b61161d8261124b565b9050602081019050919050565b828183375f83830152505050565b5f61164a611645846115fa565b6115e0565b9050828152602081018484840111156116665761166561157e565b5b61167184828561162a565b509392505050565b5f82601f83011261168d5761168c61157a565b5b813561169d848260208601611638565b91505092915050565b5f67ffffffffffffffff8211156116c0576116bf611582565b5b602082029050602081019050919050565b5f80fd5b5f6116e76116e2846116a6565b6115e0565b9050808382526020820190506020840283018581111561170a576117096116d1565b5b835b81811015611733578061171f8882611473565b84526020840193505060208101905061170c565b5050509392505050565b5f82601f8301126117515761175061157a565b5b81356117618482602086016116d5565b91505092915050565b5f805f606084860312156117815761178061116d565b5b5f84013567ffffffffffffffff81111561179e5761179d611171565b5b6117aa86828701611679565b935050602084013567ffffffffffffffff8111156117cb576117ca611171565b5b6117d786828701611679565b925050604084013567ffffffffffffffff8111156117f8576117f7611171565b5b6118048682870161173d565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061185257607f821691505b6020821081036118655761186461180e565b5b50919050565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6118d5602e8361186b565b91506118e08261187b565b604082019050919050565b5f6020820190508181035f830152611902816118c9565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61194161193c61193784611909565b61191e565b611912565b9050919050565b61195181611927565b82525050565b5f60208201905061196a5f830184611948565b92915050565b7f55736572206d75737420657869737400000000000000000000000000000000005f82015250565b5f6119a4600f8361186b565b91506119af82611970565b602082019050919050565b5f6020820190508181035f8301526119d181611998565b9050919050565b7f55736572206973206e6f742070617274206f662074686174206c6973740000005f82015250565b5f611a0c601d8361186b565b9150611a17826119d8565b602082019050919050565b5f6020820190508181035f830152611a3981611a00565b9050919050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a8082611a40565b9150611a8b83611a40565b9250828203905081811115611aa357611aa2611a49565b5b92915050565b5f611ab382611a40565b9150611abe83611a40565b9250828201905080821115611ad657611ad5611a49565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611b1382611a40565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b4557611b44611a49565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f5573657220697320616c72656164792070617274206f662074686174206c69735f8201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b5f611bd760218361186b565b9150611be282611b7d565b604082019050919050565b5f6020820190508181035f830152611c0481611bcb565b9050919050565b7f55736572206d7573742068617665207468617420726f6c6500000000000000005f82015250565b5f611c3f60188361186b565b9150611c4a82611c0b565b602082019050919050565b5f6020820190508181035f830152611c6c81611c33565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611ccf7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c94565b611cd98683611c94565b95508019841693508086168417925050509392505050565b5f611d0b611d06611d0184611a40565b61191e565b611a40565b9050919050565b5f819050919050565b611d2483611cf1565b611d38611d3082611d12565b848454611ca0565b825550505050565b5f90565b611d4c611d40565b611d57818484611d1b565b505050565b5b81811015611d7a57611d6f5f82611d44565b600181019050611d5d565b5050565b601f821115611dbf57611d9081611c73565b611d9984611c85565b81016020851015611da8578190505b611dbc611db485611c85565b830182611d5c565b50505b505050565b5f82821c905092915050565b5f611ddf5f1984600802611dc4565b1980831691505092915050565b5f611df78383611dd0565b9150826002028217905092915050565b611e1082611209565b67ffffffffffffffff811115611e2957611e28611582565b5b611e33825461183b565b611e3e828285611d7e565b5f60209050601f831160018114611e6f575f8415611e5d578287015190505b611e678582611dec565b865550611ece565b601f198416611e7d86611c73565b5f5b82811015611ea457848901518255600182019150602085019450602081019050611e7f565b86831015611ec15784890151611ebd601f891682611dd0565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611f3060268361186b565b9150611f3b82611ed6565b604082019050919050565b5f6020820190508181035f830152611f5d81611f24565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611f9860208361186b565b9150611fa382611f64565b602082019050919050565b5f6020820190508181035f830152611fc581611f8c565b905091905056fea2646970667358221220c0135d551ce553a21ef28ca9ba3ec44dece3cb17a8e77fba94ef717c0b09f59c64736f6c63430008150033",
}

// UsersABI is the input ABI used to generate the binding from.
// Deprecated: Use UsersMetaData.ABI instead.
var UsersABI = UsersMetaData.ABI

// UsersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UsersMetaData.Bin instead.
var UsersBin = UsersMetaData.Bin

// DeployUsers deploys a new Ethereum contract, binding an instance of Users to it.
func DeployUsers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Users, error) {
	parsed, err := UsersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UsersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// Users is an auto generated Go binding around an Ethereum contract.
type Users struct {
	UsersCaller     // Read-only binding to the contract
	UsersTransactor // Write-only binding to the contract
	UsersFilterer   // Log filterer for contract events
}

// UsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersSession struct {
	Contract     *Users            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersCallerSession struct {
	Contract *UsersCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersTransactorSession struct {
	Contract     *UsersTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersRaw struct {
	Contract *Users // Generic contract binding to access the raw methods on
}

// UsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersCallerRaw struct {
	Contract *UsersCaller // Generic read-only contract binding to access the raw methods on
}

// UsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersTransactorRaw struct {
	Contract *UsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsers creates a new instance of Users, bound to a specific deployed contract.
func NewUsers(address common.Address, backend bind.ContractBackend) (*Users, error) {
	contract, err := bindUsers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// NewUsersCaller creates a new read-only instance of Users, bound to a specific deployed contract.
func NewUsersCaller(address common.Address, caller bind.ContractCaller) (*UsersCaller, error) {
	contract, err := bindUsers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsersCaller{contract: contract}, nil
}

// NewUsersTransactor creates a new write-only instance of Users, bound to a specific deployed contract.
func NewUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersTransactor, error) {
	contract, err := bindUsers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsersTransactor{contract: contract}, nil
}

// NewUsersFilterer creates a new log filterer instance of Users, bound to a specific deployed contract.
func NewUsersFilterer(address common.Address, filterer bind.ContractFilterer) (*UsersFilterer, error) {
	contract, err := bindUsers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsersFilterer{contract: contract}, nil
}

// bindUsers binds a generic wrapper to an already deployed contract.
func bindUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.UsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Users.Contract.GetUser(&_Users.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Users.Contract.GetUser(&_Users.CallOpts, userAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersSession) Owner() (common.Address, error) {
	return _Users.Contract.Owner(&_Users.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersCallerSession) Owner() (common.Address, error) {
	return _Users.Contract.Owner(&_Users.CallOpts)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Users.Contract.ShowUsersInList(&_Users.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Users.Contract.ShowUsersInList(&_Users.CallOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.AddUserToList(&_Users.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.AddUserToList(&_Users.TransactOpts, serviceType)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersSession) Initialize() (*types.Transaction, error) {
	return _Users.Contract.Initialize(&_Users.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersTransactorSession) Initialize() (*types.Transaction, error) {
	return _Users.Contract.Initialize(&_Users.TransactOpts)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.RemoveUserFromList(&_Users.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.RemoveUserFromList(&_Users.TransactOpts, serviceType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersSession) RenounceOwnership() (*types.Transaction, error) {
	return _Users.Contract.RenounceOwnership(&_Users.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Users.Contract.RenounceOwnership(&_Users.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Users.Contract.TransferOwnership(&_Users.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Users.Contract.TransferOwnership(&_Users.TransactOpts, newOwner)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "updateUser", metadataCID, url, roles)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersSession) UpdateUser(metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.Contract.UpdateUser(&_Users.TransactOpts, metadataCID, url, roles)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersTransactorSession) UpdateUser(metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.Contract.UpdateUser(&_Users.TransactOpts, metadataCID, url, roles)
}

// UsersInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Users contract.
type UsersInitializedIterator struct {
	Event *UsersInitialized // Event containing the contract specifics and raw log

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
func (it *UsersInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersInitialized)
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
		it.Event = new(UsersInitialized)
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
func (it *UsersInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersInitialized represents a Initialized event raised by the Users contract.
type UsersInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Users *UsersFilterer) FilterInitialized(opts *bind.FilterOpts) (*UsersInitializedIterator, error) {

	logs, sub, err := _Users.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UsersInitializedIterator{contract: _Users.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Users *UsersFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UsersInitialized) (event.Subscription, error) {

	logs, sub, err := _Users.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersInitialized)
				if err := _Users.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Users *UsersFilterer) ParseInitialized(log types.Log) (*UsersInitialized, error) {
	event := new(UsersInitialized)
	if err := _Users.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Users contract.
type UsersOwnershipTransferredIterator struct {
	Event *UsersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UsersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersOwnershipTransferred)
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
		it.Event = new(UsersOwnershipTransferred)
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
func (it *UsersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersOwnershipTransferred represents a OwnershipTransferred event raised by the Users contract.
type UsersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Users *UsersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsersOwnershipTransferredIterator{contract: _Users.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Users *UsersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersOwnershipTransferred)
				if err := _Users.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Users *UsersFilterer) ParseOwnershipTransferred(log types.Log) (*UsersOwnershipTransferred, error) {
	event := new(UsersOwnershipTransferred)
	if err := _Users.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
