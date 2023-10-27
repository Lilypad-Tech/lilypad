// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

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

// PaymentsMetaData contains all meta data concerning the Payments contract.
var PaymentsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600360146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6128c080620001425f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c80639e3868dc116100b6578063b91880351161007a578063b9188035146102d2578063c4d66de8146102ee578063c57380a21461030a578063d2485cce14610328578063f2fde38b14610344578063f3d3d4481461036057610140565b80639e3868dc14610258578063a470295814610274578063aea382511461027e578063afe1dff71461029a578063b1356714146102b657610140565b80632a1f9072116101085780632a1f9072146101d257806338698529146101ee5780634bc28da11461020a578063715018a614610214578063823f3de11461021e5780638da5cb5b1461023a57610140565b806302fd8f801461014457806309cab510146101605780630ef0d89e1461017c57806310fe9ae81461019857806326a4e8d2146101b6575b5f80fd5b61015e60048036038101906101599190611951565b61037c565b005b61017a600480360381019061017591906119e4565b610420565b005b61019660048036038101906101919190611a64565b6104b6565b005b6101a061053f565b6040516101ad9190611adf565b60405180910390f35b6101d060048036038101906101cb9190611af8565b610567565b005b6101ec60048036038101906101e79190611b23565b6106af565b005b61020860048036038101906102039190611951565b610725565b005b6102126107bd565b005b61021c6107e1565b005b61023860048036038101906102339190611bdc565b6107f4565b005b6102426108cf565b60405161024f9190611adf565b60405180910390f35b610272600480360381019061026d9190611a64565b6108f6565b005b61027c61097e565b005b610298600480360381019061029391906119e4565b6109a2565b005b6102b460048036038101906102af91906119e4565b610a38565b005b6102d060048036038101906102cb9190611b23565b610ace565b005b6102ec60048036038101906102e791906119e4565b610bb1565b005b61030860048036038101906103039190611af8565b610c45565b005b610312610d86565b60405161031f9190611adf565b60405180910390f35b610342600480360381019061033d9190611bdc565b610dae565b005b61035e60048036038101906103599190611af8565b610de6565b005b61037a60048036038101906103759190611af8565b610e68565b005b610384610f70565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146103f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea90611d01565b60405180910390fd5b6103ff8584845f61109d565b61040c858483600261109d565b61041985858360026111c3565b5050505050565b610428610f70565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90611d8f565b60405180910390fd5b6104a3848360016112e9565b6104b0848483600261109d565b50505050565b6104be610f70565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461052d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052490611d8f565b60405180910390fd5b61053a838383600261109d565b505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61056f6114e6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d490611e1d565b60405180910390fd5b600360149054906101000a900460ff1661062c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062390611eab565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6106b7610f70565b505f8490505f848611156106cd578491506106dc565b85856106d99190611ef6565b90505b6106ea89888a856003611564565b6106f8898832866004611564565b5f81111561070d5761070c8988835f61109d565b5b61071a898986600161109d565b505050505050505050565b61072d610f70565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079390611d8f565b60405180910390fd5b6107a9858584600161109d565b6107b685848360026111c3565b5050505050565b6107c56114e6565b5f600360146101000a81548160ff021916908315150217905550565b6107e96114e6565b6107f25f61168d565b565b6107fc610f70565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16148061086257508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b6108a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089890611f99565b60405180910390fd5b6108ae868684600161109d565b6108ba8685855f61109d565b6108c7868583600461109d565b505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6108fe610f70565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461096d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096490611d8f565b60405180910390fd5b610979838260026112e9565b505050565b6109866114e6565b5f600160146101000a81548160ff021916908315150217905550565b6109aa610f70565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610a19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1090611d01565b60405180910390fd5b610a26848484600261109d565b610a32848260046112e9565b50505050565b610a40610f70565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610aaf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa690611d01565b60405180910390fd5b610abb8484845f61109d565b610ac8848483600261109d565b50505050565b610ad6610f70565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610b45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3c90611d01565b60405180910390fd5b5f8490505f84861115610b5a57849150610b69565b8585610b669190611ef6565b90505b610b7789888a856003611564565b5f811115610b8c57610b8b8988835f61109d565b5b610b99898885600261109d565b610ba6898986600161109d565b505050505050505050565b610bb9610f70565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610c28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1f90611d01565b60405180910390fd5b610c3384835f6112e9565b610c3f848260026112e9565b50505050565b5f600160169054906101000a900460ff16159050808015610c77575060018060159054906101000a900460ff1660ff16105b80610ca55750610c863061174e565b158015610ca4575060018060159054906101000a900460ff1660ff16145b5b610ce4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cdb90612027565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610d205760018060166101000a81548160ff0219169083151502179055505b610d2982610567565b8015610d82575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610d799190612093565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610db6610f70565b50610dc38685855f61109d565b610dd1868532846004611564565b610dde86868460016111c3565b505050505050565b610dee6114e6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e539061211c565b60405180910390fd5b610e658161168d565b50565b610e706114e6565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ede576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed5906121aa565b60405180910390fd5b600160149054906101000a900460ff16610f2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2490612238565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611000576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff7906121aa565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611040611770565b73ffffffffffffffffffffffffffffffffffffffff1614611096576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108d906122c6565b60405180910390fd5b6001905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663599efa6b85856040518363ffffffff1660e01b81526004016110fa9291906122f3565b6020604051808303815f875af1158015611116573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061113a919061234f565b90508061117c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611173906123ea565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a98428585858560026040516111b495949392919061252b565b60405180910390a15050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388c2bdfe85856040518363ffffffff1660e01b81526004016112209291906122f3565b6020604051808303815f875af115801561123c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611260919061234f565b9050806112a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611299906125f3565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a98428585858560036040516112da95949392919061252b565b60405180910390a15050505050565b8160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231326040518263ffffffff1660e01b81526004016113449190611adf565b602060405180830381865afa15801561135f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113839190612625565b10156113c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113bb906126c0565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635407e93c846040518263ffffffff1660e01b815260040161141f91906126de565b6020604051808303815f875af115801561143b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061145f919061234f565b9050806114a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149890612767565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842843285855f6040516114d895949392919061252b565b60405180910390a150505050565b6114ee611770565b73ffffffffffffffffffffffffffffffffffffffff1661150c6108cf565b73ffffffffffffffffffffffffffffffffffffffff1614611562576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611559906127cf565b60405180910390fd5b565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663065e86c88686866040518463ffffffff1660e01b81526004016115c3939291906127ed565b6020604051808303815f875af11580156115df573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611603919061234f565b905080611645576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163c9061286c565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a984286858585600160405161167d95949392919061252b565b60405180910390a1505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6117d682611790565b810181811067ffffffffffffffff821117156117f5576117f46117a0565b5b80604052505050565b5f611807611777565b905061181382826117cd565b919050565b5f67ffffffffffffffff821115611832576118316117a0565b5b61183b82611790565b9050602081019050919050565b828183375f83830152505050565b5f61186861186384611818565b6117fe565b9050828152602081018484840111156118845761188361178c565b5b61188f848285611848565b509392505050565b5f82601f8301126118ab576118aa611788565b5b81356118bb848260208601611856565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6118ed826118c4565b9050919050565b6118fd816118e3565b8114611907575f80fd5b50565b5f81359050611918816118f4565b92915050565b5f819050919050565b6119308161191e565b811461193a575f80fd5b50565b5f8135905061194b81611927565b92915050565b5f805f805f60a0868803121561196a57611969611780565b5b5f86013567ffffffffffffffff81111561198757611986611784565b5b61199388828901611897565b95505060206119a48882890161190a565b94505060406119b58882890161190a565b93505060606119c68882890161193d565b92505060806119d78882890161193d565b9150509295509295909350565b5f805f80608085870312156119fc576119fb611780565b5b5f85013567ffffffffffffffff811115611a1957611a18611784565b5b611a2587828801611897565b9450506020611a368782880161190a565b9350506040611a478782880161193d565b9250506060611a588782880161193d565b91505092959194509250565b5f805f60608486031215611a7b57611a7a611780565b5b5f84013567ffffffffffffffff811115611a9857611a97611784565b5b611aa486828701611897565b9350506020611ab58682870161190a565b9250506040611ac68682870161193d565b9150509250925092565b611ad9816118e3565b82525050565b5f602082019050611af25f830184611ad0565b92915050565b5f60208284031215611b0d57611b0c611780565b5b5f611b1a8482850161190a565b91505092915050565b5f805f805f805f60e0888a031215611b3e57611b3d611780565b5b5f88013567ffffffffffffffff811115611b5b57611b5a611784565b5b611b678a828b01611897565b9750506020611b788a828b0161190a565b9650506040611b898a828b0161190a565b9550506060611b9a8a828b0161193d565b9450506080611bab8a828b0161193d565b93505060a0611bbc8a828b0161193d565b92505060c0611bcd8a828b0161193d565b91505092959891949750929550565b5f805f805f8060c08789031215611bf657611bf5611780565b5b5f87013567ffffffffffffffff811115611c1357611c12611784565b5b611c1f89828a01611897565b9650506020611c3089828a0161190a565b9550506040611c4189828a0161190a565b9450506060611c5289828a0161193d565b9350506080611c6389828a0161193d565b92505060a0611c7489828a0161193d565b9150509295509295509295565b5f82825260208201905092915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865204a4300000000000000000000000000000000000000602082015250565b5f611ceb602d83611c81565b9150611cf682611c91565b604082019050919050565b5f6020820190508181035f830152611d1881611cdf565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c65642062792074686520525000000000000000000000000000000000000000602082015250565b5f611d79602d83611c81565b9150611d8482611d1f565b604082019050919050565b5f6020820190508181035f830152611da681611d6d565b9050919050565b7f4c696c657061645061796d656e74733a20546f6b656e2061646472657373206d5f8201527f75737420626520646566696e6564000000000000000000000000000000000000602082015250565b5f611e07602e83611c81565b9150611e1282611dad565b604082019050919050565b5f6020820190508181035f830152611e3481611dfb565b9050919050565b7f4c696c79706164546f6b656e3a2063616e4368616e6765546f6b656e416464725f8201527f6573732069732064697361626c65640000000000000000000000000000000000602082015250565b5f611e95602f83611c81565b9150611ea082611e3b565b604082019050919050565b5f6020820190508181035f830152611ec281611e89565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f008261191e565b9150611f0b8361191e565b9250828203905081811115611f2357611f22611ec9565b5b92915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865205250206f72204a4300000000000000000000000000602082015250565b5f611f83603383611c81565b9150611f8e82611f29565b604082019050919050565b5f6020820190508181035f830152611fb081611f77565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f612011602e83611c81565b915061201c82611fb7565b604082019050919050565b5f6020820190508181035f83015261203e81612005565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61207d61207861207384612045565b61205a565b61204e565b9050919050565b61208d81612063565b82525050565b5f6020820190506120a65f830184612084565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612106602683611c81565b9150612111826120ac565b604082019050919050565b5f6020820190508181035f830152612133816120fa565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f612194603583611c81565b915061219f8261213a565b604082019050919050565b5f6020820190508181035f8301526121c181612188565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f612222603983611c81565b915061222d826121c8565b604082019050919050565b5f6020820190508181035f83015261224f81612216565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6122b0603b83611c81565b91506122bb82612256565b604082019050919050565b5f6020820190508181035f8301526122dd816122a4565b9050919050565b6122ed8161191e565b82525050565b5f6040820190506123065f830185611ad0565b61231360208301846122e4565b9392505050565b5f8115159050919050565b61232e8161231a565b8114612338575f80fd5b50565b5f8151905061234981612325565b92915050565b5f6020828403121561236457612363611780565b5b5f6123718482850161233b565b91505092915050565b7f4c696c797061645061796d656e74733a20526566756e6420657363726f7720665f8201527f61696c6564000000000000000000000000000000000000000000000000000000602082015250565b5f6123d4602583611c81565b91506123df8261237a565b604082019050919050565b5f6020820190508181035f830152612401816123c8565b9050919050565b5f81519050919050565b5f5b8381101561242f578082015181840152602081019050612414565b5f8484015250505050565b5f61244482612408565b61244e8185611c81565b935061245e818560208601612412565b61246781611790565b840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600581106124b0576124af612472565b5b50565b5f8190506124c08261249f565b919050565b5f6124cf826124b3565b9050919050565b6124df816124c5565b82525050565b600481106124f6576124f5612472565b5b50565b5f819050612506826124e5565b919050565b5f612515826124f9565b9050919050565b6125258161250b565b82525050565b5f60a0820190508181035f830152612543818861243a565b90506125526020830187611ad0565b61255f60408301866122e4565b61256c60608301856124d6565b612579608083018461251c565b9695505050505050565b7f4c696c797061645061796d656e74733a20536c61736820657363726f772066615f8201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b5f6125dd602483611c81565b91506125e882612583565b604082019050919050565b5f6020820190508181035f83015261260a816125d1565b9050919050565b5f8151905061261f81611927565b92915050565b5f6020828403121561263a57612639611780565b5b5f61264784828501612611565b91505092915050565b7f4c696c797061645061796d656e74733a20496e73756666696369656e742062615f8201527f6c616e6365000000000000000000000000000000000000000000000000000000602082015250565b5f6126aa602583611c81565b91506126b582612650565b604082019050919050565b5f6020820190508181035f8301526126d78161269e565b9050919050565b5f6020820190506126f15f8301846122e4565b92915050565b7f4c696c797061645061796d656e74733a2050617920657363726f77206661696c5f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f612751602283611c81565b915061275c826126f7565b604082019050919050565b5f6020820190508181035f83015261277e81612745565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6127b9602083611c81565b91506127c482612785565b602082019050919050565b5f6020820190508181035f8301526127e6816127ad565b9050919050565b5f6060820190506128005f830186611ad0565b61280d6020830185611ad0565b61281a60408301846122e4565b949350505050565b7f4c696c797061645061796d656e74733a20506179206a6f62206661696c6564005f82015250565b5f612856601f83611c81565b915061286182612822565b602082019050919050565b5f6020820190508181035f8301526128838161284a565b905091905056fea264697066735822122055ff335a3901ce1bb55b485e62f915fd5eec008c00175337ff438c6d2738e4f564736f6c63430008150033",
}

// PaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentsMetaData.ABI instead.
var PaymentsABI = PaymentsMetaData.ABI

// PaymentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentsMetaData.Bin instead.
var PaymentsBin = PaymentsMetaData.Bin

// DeployPayments deploys a new Ethereum contract, binding an instance of Payments to it.
func DeployPayments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Payments, error) {
	parsed, err := PaymentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// Payments is an auto generated Go binding around an Ethereum contract.
type Payments struct {
	PaymentsCaller     // Read-only binding to the contract
	PaymentsTransactor // Write-only binding to the contract
	PaymentsFilterer   // Log filterer for contract events
}

// PaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentsSession struct {
	Contract     *Payments         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentsCallerSession struct {
	Contract *PaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentsTransactorSession struct {
	Contract     *PaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentsRaw struct {
	Contract *Payments // Generic contract binding to access the raw methods on
}

// PaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentsCallerRaw struct {
	Contract *PaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentsTransactorRaw struct {
	Contract *PaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayments creates a new instance of Payments, bound to a specific deployed contract.
func NewPayments(address common.Address, backend bind.ContractBackend) (*Payments, error) {
	contract, err := bindPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// NewPaymentsCaller creates a new read-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PaymentsCaller, error) {
	contract, err := bindPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsCaller{contract: contract}, nil
}

// NewPaymentsTransactor creates a new write-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentsTransactor, error) {
	contract, err := bindPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransactor{contract: contract}, nil
}

// NewPaymentsFilterer creates a new log filterer instance of Payments, bound to a specific deployed contract.
func NewPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentsFilterer, error) {
	contract, err := bindPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentsFilterer{contract: contract}, nil
}

// bindPayments binds a generic wrapper to an already deployed contract.
func bindPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.PaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsSession) GetControllerAddress() (common.Address, error) {
	return _Payments.Contract.GetControllerAddress(&_Payments.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsCallerSession) GetControllerAddress() (common.Address, error) {
	return _Payments.Contract.GetControllerAddress(&_Payments.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsSession) GetTokenAddress() (common.Address, error) {
	return _Payments.Contract.GetTokenAddress(&_Payments.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsCallerSession) GetTokenAddress() (common.Address, error) {
	return _Payments.Contract.GetTokenAddress(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsCallerSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AcceptResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "acceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AddResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "addResult", dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AddResult(dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AddResult(dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) CheckResult(opts *bind.TransactOpts, dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "checkResult", dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) CheckResult(dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckResult(&_Payments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) CheckResult(dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckResult(&_Payments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeControllerAddress(&_Payments.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeControllerAddress(&_Payments.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsTransactor) DisableChangeTokenAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "disableChangeTokenAddress")
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeTokenAddress(&_Payments.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsTransactorSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeTokenAddress(&_Payments.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.Initialize(&_Payments.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.Initialize(&_Payments.TransactOpts, _tokenAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payments.Contract.RenounceOwnership(&_Payments.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payments.Contract.RenounceOwnership(&_Payments.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetControllerAddress(&_Payments.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetControllerAddress(&_Payments.TransactOpts, _controllerAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetTokenAddress(&_Payments.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetTokenAddress(&_Payments.TransactOpts, _tokenAddress)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundJobCreator(opts *bind.TransactOpts, dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundResourceProvider(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutJudgeResults(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutJudgeResults", dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutJudgeResults(dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutJudgeResults(dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutMediateResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) TimeoutMediateResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) TimeoutMediateResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutSubmitResults(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutSubmitResults", dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutSubmitResults(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutSubmitResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutSubmitResults(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutSubmitResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// PaymentsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Payments contract.
type PaymentsInitializedIterator struct {
	Event *PaymentsInitialized // Event containing the contract specifics and raw log

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
func (it *PaymentsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsInitialized)
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
		it.Event = new(PaymentsInitialized)
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
func (it *PaymentsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsInitialized represents a Initialized event raised by the Payments contract.
type PaymentsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Payments *PaymentsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PaymentsInitializedIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PaymentsInitializedIterator{contract: _Payments.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Payments *PaymentsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PaymentsInitialized) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsInitialized)
				if err := _Payments.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Payments *PaymentsFilterer) ParseInitialized(log types.Log) (*PaymentsInitialized, error) {
	event := new(PaymentsInitialized)
	if err := _Payments.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Payments contract.
type PaymentsOwnershipTransferredIterator struct {
	Event *PaymentsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsOwnershipTransferred)
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
		it.Event = new(PaymentsOwnershipTransferred)
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
func (it *PaymentsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsOwnershipTransferred represents a OwnershipTransferred event raised by the Payments contract.
type PaymentsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payments *PaymentsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsOwnershipTransferredIterator{contract: _Payments.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payments *PaymentsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsOwnershipTransferred)
				if err := _Payments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Payments *PaymentsFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentsOwnershipTransferred, error) {
	event := new(PaymentsOwnershipTransferred)
	if err := _Payments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the Payments contract.
type PaymentsPaymentIterator struct {
	Event *PaymentsPayment // Event containing the contract specifics and raw log

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
func (it *PaymentsPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsPayment)
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
		it.Event = new(PaymentsPayment)
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
func (it *PaymentsPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsPayment represents a Payment event raised by the Payments contract.
type PaymentsPayment struct {
	DealId    string
	Payee     common.Address
	Amount    *big.Int
	Reason    uint8
	Direction uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) FilterPayment(opts *bind.FilterOpts) (*PaymentsPaymentIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &PaymentsPaymentIterator{contract: _Payments.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PaymentsPayment) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsPayment)
				if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) ParsePayment(log types.Log) (*PaymentsPayment, error) {
	event := new(PaymentsPayment)
	if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
