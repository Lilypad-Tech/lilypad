// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"escrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"refundEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashEscrow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055503480156200002a575f80fd5b5060405162002a0038038062002a008339818101604052810190620000509190620004a8565b82826200007262000066620000b360201b60201c565b620000ba60201b60201c565b81600590816200008391906200076d565b5080600690816200009591906200076d565b505050620000aa33826200017b60201b60201c565b50505062000962565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620001ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001e390620008af565b60405180910390fd5b620001ff5f8383620002e160201b60201c565b8060045f828254620002129190620008fc565b925050819055508060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002c2919062000947565b60405180910390a3620002dd5f8383620002e660201b60201c565b5050565b505050565b505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200034c8262000304565b810181811067ffffffffffffffff821117156200036e576200036d62000314565b5b80604052505050565b5f62000382620002eb565b905062000390828262000341565b919050565b5f67ffffffffffffffff821115620003b257620003b162000314565b5b620003bd8262000304565b9050602081019050919050565b5f5b83811015620003e9578082015181840152602081019050620003cc565b5f8484015250505050565b5f6200040a620004048462000395565b62000377565b90508281526020810184848401111562000429576200042862000300565b5b62000436848285620003ca565b509392505050565b5f82601f830112620004555762000454620002fc565b5b815162000467848260208601620003f4565b91505092915050565b5f819050919050565b620004848162000470565b81146200048f575f80fd5b50565b5f81519050620004a28162000479565b92915050565b5f805f60608486031215620004c257620004c1620002f4565b5b5f84015167ffffffffffffffff811115620004e257620004e1620002f8565b5b620004f0868287016200043e565b935050602084015167ffffffffffffffff811115620005145762000513620002f8565b5b62000522868287016200043e565b9250506040620005358682870162000492565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200058e57607f821691505b602082108103620005a457620005a362000549565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620006087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005cb565b620006148683620005cb565b95508019841693508086168417925050509392505050565b5f819050919050565b5f620006556200064f620006498462000470565b6200062c565b62000470565b9050919050565b5f819050919050565b620006708362000635565b620006886200067f826200065c565b848454620005d7565b825550505050565b5f90565b6200069e62000690565b620006ab81848462000665565b505050565b5b81811015620006d257620006c65f8262000694565b600181019050620006b1565b5050565b601f8211156200072157620006eb81620005aa565b620006f684620005bc565b8101602085101562000706578190505b6200071e6200071585620005bc565b830182620006b0565b50505b505050565b5f82821c905092915050565b5f620007435f198460080262000726565b1980831691505092915050565b5f6200075d838362000732565b9150826002028217905092915050565b62000778826200053f565b67ffffffffffffffff81111562000794576200079362000314565b5b620007a0825462000576565b620007ad828285620006d6565b5f60209050601f831160018114620007e3575f8415620007ce578287015190505b620007da858262000750565b86555062000849565b601f198416620007f386620005aa565b5f5b828110156200081c57848901518255600182019150602085019450602081019050620007f5565b868310156200083c578489015162000838601f89168262000732565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f62000897601f8362000851565b9150620008a48262000861565b602082019050919050565b5f6020820190508181035f830152620008c88162000889565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620009088262000470565b9150620009158362000470565b925082820190508082111562000930576200092f620008cf565b5b92915050565b620009418162000470565b82525050565b5f6020820190506200095c5f83018462000936565b92915050565b61209080620009705f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c806388c2bdfe116100b6578063a47029581161007a578063a4702958146103c4578063a9059cbb146103ce578063c57380a2146103fe578063dd62ed3e1461041c578063f2fde38b1461044c578063f3d3d4481461046857610140565b806388c2bdfe146102f85780638da5cb5b1461032857806395d89b4114610346578063987bf9a714610364578063a457c2d71461039457610140565b8063313ce56711610108578063313ce56714610210578063395093511461022e5780635407e93c1461025e578063599efa6b1461028e57806370a08231146102be578063715018a6146102ee57610140565b8063065e86c81461014457806306fdde0314610174578063095ea7b31461019257806318160ddd146101c257806323b872dd146101e0575b5f80fd5b61015e6004803603810190610159919061156f565b610484565b60405161016b91906115d9565b60405180910390f35b61017c610577565b604051610189919061167c565b60405180910390f35b6101ac60048036038101906101a7919061169c565b610607565b6040516101b991906115d9565b60405180910390f35b6101ca610629565b6040516101d791906116e9565b60405180910390f35b6101fa60048036038101906101f5919061156f565b610632565b60405161020791906115d9565b60405180910390f35b610218610660565b604051610225919061171d565b60405180910390f35b6102486004803603810190610243919061169c565b610668565b60405161025591906115d9565b60405180910390f35b61027860048036038101906102739190611736565b61069e565b60405161028591906115d9565b60405180910390f35b6102a860048036038101906102a3919061169c565b610706565b6040516102b591906115d9565b60405180910390f35b6102d860048036038101906102d39190611761565b610866565b6040516102e591906116e9565b60405180910390f35b6102f66108ac565b005b610312600480360381019061030d919061169c565b6108bf565b60405161031f91906115d9565b60405180910390f35b6103306109b8565b60405161033d919061179b565b60405180910390f35b61034e6109df565b60405161035b919061167c565b60405180910390f35b61037e60048036038101906103799190611761565b610a6f565b60405161038b91906116e9565b60405180910390f35b6103ae60048036038101906103a9919061169c565b610ab5565b6040516103bb91906115d9565b60405180910390f35b6103cc610b2a565b005b6103e860048036038101906103e3919061169c565b610b4e565b6040516103f591906115d9565b60405180910390f35b610406610b70565b604051610413919061179b565b60405180910390f35b610436600480360381019061043191906117b4565b610b98565b60405161044391906116e9565b60405180910390f35b61046660048036038101906104619190611761565b610c1a565b005b610482600480360381019061047d9190611761565b610c9c565b005b5f61048d610da4565b508160075f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561050e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050590611862565b60405180910390fd5b8160075f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461055a91906118ad565b9250508190555061056c308484610ed1565b600190509392505050565b6060600580546105869061190d565b80601f01602080910402602001604051908101604052809291908181526020018280546105b29061190d565b80156105fd5780601f106105d4576101008083540402835291602001916105fd565b820191905f5260205f20905b8154815290600101906020018083116105e057829003601f168201915b5050505050905090565b5f80610611611140565b905061061e818585611147565b600191505092915050565b5f600454905090565b5f8061063c611140565b905061064985828561130a565b610654858585610ed1565b60019150509392505050565b5f6012905090565b5f80610672611140565b90506106938185856106848589610b98565b61068e919061193d565b611147565b600191505092915050565b5f6106aa323084610ed1565b8160075f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106f6919061193d565b9250508190555060019050919050565b5f61070f610da4565b505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361077e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610775906119e0565b60405180910390fd5b8160075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156107fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f590611862565b60405180910390fd5b8160075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461084a91906118ad565b9250508190555061085c308484610ed1565b6001905092915050565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6108b4611395565b6108bd5f611413565b565b5f6108c8610da4565b508160075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610949576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094090611862565b60405180910390fd5b8160075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461099591906118ad565b925050819055506109ae306109a86109b8565b84610ed1565b6001905092915050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600680546109ee9061190d565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1a9061190d565b8015610a655780601f10610a3c57610100808354040283529160200191610a65565b820191905f5260205f20905b815481529060010190602001808311610a4857829003601f168201915b5050505050905090565b5f60075f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f80610abf611140565b90505f610acc8286610b98565b905083811015610b11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0890611a6e565b60405180910390fd5b610b1e8286868403611147565b60019250505092915050565b610b32611395565b5f600160146101000a81548160ff021916908315150217905550565b5f80610b58611140565b9050610b65818585610ed1565b600191505092915050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b610c22611395565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c90576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8790611afc565b60405180910390fd5b610c9981611413565b50565b610ca4611395565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0990611b8a565b60405180910390fd5b600160149054906101000a900460ff16610d61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5890611c18565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2b90611b8a565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610e74611140565b73ffffffffffffffffffffffffffffffffffffffff1614610eca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec190611ca6565b60405180910390fd5b6001905090565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610f3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3690611d34565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610fad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa490611dc2565b60405180910390fd5b610fb88383836114d4565b5f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508181101561103c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103390611e50565b60405180910390fd5b81810360025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161112791906116e9565b60405180910390a361113a8484846114d9565b50505050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036111b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ac90611ede565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611223576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121a90611f6c565b60405180910390fd5b8060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516112fd91906116e9565b60405180910390a3505050565b5f6113158484610b98565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461138f5781811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137890611fd4565b60405180910390fd5b61138e8484848403611147565b5b50505050565b61139d611140565b73ffffffffffffffffffffffffffffffffffffffff166113bb6109b8565b73ffffffffffffffffffffffffffffffffffffffff1614611411576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114089061203c565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61150b826114e2565b9050919050565b61151b81611501565b8114611525575f80fd5b50565b5f8135905061153681611512565b92915050565b5f819050919050565b61154e8161153c565b8114611558575f80fd5b50565b5f8135905061156981611545565b92915050565b5f805f60608486031215611586576115856114de565b5b5f61159386828701611528565b93505060206115a486828701611528565b92505060406115b58682870161155b565b9150509250925092565b5f8115159050919050565b6115d3816115bf565b82525050565b5f6020820190506115ec5f8301846115ca565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561162957808201518184015260208101905061160e565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61164e826115f2565b61165881856115fc565b935061166881856020860161160c565b61167181611634565b840191505092915050565b5f6020820190508181035f8301526116948184611644565b905092915050565b5f80604083850312156116b2576116b16114de565b5b5f6116bf85828601611528565b92505060206116d08582860161155b565b9150509250929050565b6116e38161153c565b82525050565b5f6020820190506116fc5f8301846116da565b92915050565b5f60ff82169050919050565b61171781611702565b82525050565b5f6020820190506117305f83018461170e565b92915050565b5f6020828403121561174b5761174a6114de565b5b5f6117588482850161155b565b91505092915050565b5f60208284031215611776576117756114de565b5b5f61178384828501611528565b91505092915050565b61179581611501565b82525050565b5f6020820190506117ae5f83018461178c565b92915050565b5f80604083850312156117ca576117c96114de565b5b5f6117d785828601611528565b92505060206117e885828601611528565b9150509250929050565b7f4c696c79706164546f6b656e3a206e6f7420656e6f7567682066756e647320695f8201527f6e20657363726f77000000000000000000000000000000000000000000000000602082015250565b5f61184c6028836115fc565b9150611857826117f2565b604082019050919050565b5f6020820190508181035f83015261187981611840565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118b78261153c565b91506118c28361153c565b92508282039050818111156118da576118d9611880565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061192457607f821691505b602082108103611937576119366118e0565b5b50919050565b5f6119478261153c565b91506119528361153c565b925082820190508082111561196a57611969611880565b5b92915050565b7f4c696c79706164546f6b656e3a20746f416464726573732063616e6e6f7420625f8201527f65207a65726f2061646472657373000000000000000000000000000000000000602082015250565b5f6119ca602e836115fc565b91506119d582611970565b604082019050919050565b5f6020820190508181035f8301526119f7816119be565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f611a586025836115fc565b9150611a63826119fe565b604082019050919050565b5f6020820190508181035f830152611a8581611a4c565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611ae66026836115fc565b9150611af182611a8c565b604082019050919050565b5f6020820190508181035f830152611b1381611ada565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f611b746035836115fc565b9150611b7f82611b1a565b604082019050919050565b5f6020820190508181035f830152611ba181611b68565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611c026039836115fc565b9150611c0d82611ba8565b604082019050919050565b5f6020820190508181035f830152611c2f81611bf6565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f611c90603b836115fc565b9150611c9b82611c36565b604082019050919050565b5f6020820190508181035f830152611cbd81611c84565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f611d1e6025836115fc565b9150611d2982611cc4565b604082019050919050565b5f6020820190508181035f830152611d4b81611d12565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f611dac6023836115fc565b9150611db782611d52565b604082019050919050565b5f6020820190508181035f830152611dd981611da0565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611e3a6026836115fc565b9150611e4582611de0565b604082019050919050565b5f6020820190508181035f830152611e6781611e2e565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f611ec86024836115fc565b9150611ed382611e6e565b604082019050919050565b5f6020820190508181035f830152611ef581611ebc565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f611f566022836115fc565b9150611f6182611efc565b604082019050919050565b5f6020820190508181035f830152611f8381611f4a565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000005f82015250565b5f611fbe601d836115fc565b9150611fc982611f8a565b602082019050919050565b5f6020820190508181035f830152611feb81611fb2565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6120266020836115fc565b915061203182611ff2565b602082019050919050565b5f6020820190508181035f8301526120538161201a565b905091905056fea2646970667358221220745731f5d3ff8ec6d43a885cc63ce6c9fccc7680b8949610c14c83cc4681c03064736f6c63430008150033",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMetaData.Bin instead.
var TokenBin = TokenMetaData.Bin

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, initialSupply *big.Int) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBin), backend, name, symbol, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_Token *TokenCaller) EscrowBalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "escrowBalanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_Token *TokenSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _Token.Contract.EscrowBalanceOf(&_Token.CallOpts, _address)
}

// EscrowBalanceOf is a free data retrieval call binding the contract method 0x987bf9a7.
//
// Solidity: function escrowBalanceOf(address _address) view returns(uint256)
func (_Token *TokenCallerSession) EscrowBalanceOf(_address common.Address) (*big.Int, error) {
	return _Token.Contract.EscrowBalanceOf(&_Token.CallOpts, _address)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Token *TokenCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Token *TokenSession) GetControllerAddress() (common.Address, error) {
	return _Token.Contract.GetControllerAddress(&_Token.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Token *TokenCallerSession) GetControllerAddress() (common.Address, error) {
	return _Token.Contract.GetControllerAddress(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Token *TokenTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Token *TokenSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Token.Contract.DisableChangeControllerAddress(&_Token.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Token *TokenTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Token.Contract.DisableChangeControllerAddress(&_Token.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_Token *TokenTransactor) PayEscrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "payEscrow", amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_Token *TokenSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.PayEscrow(&_Token.TransactOpts, amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0x5407e93c.
//
// Solidity: function payEscrow(uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) PayEscrow(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.PayEscrow(&_Token.TransactOpts, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactor) PayJob(opts *bind.TransactOpts, fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "payJob", fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_Token *TokenSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.PayJob(&_Token.TransactOpts, fromAddress, toAddress, amount)
}

// PayJob is a paid mutator transaction binding the contract method 0x065e86c8.
//
// Solidity: function payJob(address fromAddress, address toAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) PayJob(fromAddress common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.PayJob(&_Token.TransactOpts, fromAddress, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactor) RefundEscrow(opts *bind.TransactOpts, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "refundEscrow", toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_Token *TokenSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RefundEscrow(&_Token.TransactOpts, toAddress, amount)
}

// RefundEscrow is a paid mutator transaction binding the contract method 0x599efa6b.
//
// Solidity: function refundEscrow(address toAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) RefundEscrow(toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.RefundEscrow(&_Token.TransactOpts, toAddress, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Token *TokenTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Token *TokenSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetControllerAddress(&_Token.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Token *TokenTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetControllerAddress(&_Token.TransactOpts, _controllerAddress)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactor) SlashEscrow(opts *bind.TransactOpts, slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "slashEscrow", slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_Token *TokenSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SlashEscrow(&_Token.TransactOpts, slashedAddress, amount)
}

// SlashEscrow is a paid mutator transaction binding the contract method 0x88c2bdfe.
//
// Solidity: function slashEscrow(address slashedAddress, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) SlashEscrow(slashedAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SlashEscrow(&_Token.TransactOpts, slashedAddress, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
