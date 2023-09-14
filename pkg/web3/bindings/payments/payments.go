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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600360146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612c0580620001425f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c80639e3868dc116100b6578063b91880351161007a578063b9188035146102d2578063c4d66de8146102ee578063c57380a21461030a578063ddce511b14610328578063f2fde38b14610344578063f3d3d4481461036057610140565b80639e3868dc14610258578063a470295814610274578063aea382511461027e578063afe1dff71461029a578063b1356714146102b657610140565b8063386985291161010857806338698529146101d25780634bc28da1146101ee578063715018a6146101f8578063823f3de114610202578063888aa14c1461021e5780638da5cb5b1461023a57610140565b806302fd8f801461014457806309cab510146101605780630ef0d89e1461017c57806310fe9ae81461019857806326a4e8d2146101b6575b5f80fd5b61015e60048036038101906101599190611a7f565b61037c565b005b61017a60048036038101906101759190611b12565b610420565b005b61019660048036038101906101919190611b92565b6104b6565b005b6101a061053f565b6040516101ad9190611c0d565b60405180910390f35b6101d060048036038101906101cb9190611c26565b610567565b005b6101ec60048036038101906101e79190611a7f565b6106af565b005b6101f6610747565b005b61020061076b565b005b61021c60048036038101906102179190611c51565b61077e565b005b61023860048036038101906102339190611cf6565b610859565b005b610242610900565b60405161024f9190611c0d565b60405180910390f35b610272600480360381019061026d9190611b92565b610927565b005b61027c6109af565b005b61029860048036038101906102939190611b12565b6109d3565b005b6102b460048036038101906102af9190611b12565b610a69565b005b6102d060048036038101906102cb9190611daf565b610aff565b005b6102ec60048036038101906102e79190611b12565b610be2565b005b61030860048036038101906103039190611c26565b610c76565b005b610312610db7565b60405161031f9190611c0d565b60405180910390f35b610342600480360381019061033d9190611e68565b610ddf565b005b61035e60048036038101906103599190611c26565b610ec4565b005b61037a60048036038101906103759190611c26565b610f46565b005b61038461104e565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146103f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea90611fb5565b60405180910390fd5b6103ff8584845f61117b565b61040c858483600261117b565b61041985858360026112b5565b5050505050565b61042861104e565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e90612043565b60405180910390fd5b6104a3848360016113ef565b6104b0848483600261117b565b50505050565b6104be61104e565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461052d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052490612043565b60405180910390fd5b61053a838383600261117b565b505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61056f611600565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d4906120d1565b60405180910390fd5b600360149054906101000a900460ff1661062c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106239061215f565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6106b761104e565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071d90612043565b60405180910390fd5b610733858584600161117b565b61074085848360026112b5565b5050505050565b61074f611600565b5f600360146101000a81548160ff021916908315150217905550565b610773611600565b61077c5f61167e565b565b61078661104e565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614806107ec57508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b61082b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610822906121ed565b60405180910390fd5b610838868684600161117b565b6108448685855f61117b565b610851868583600461117b565b505050505050565b61086161104e565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146108d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c79061227b565b60405180910390fd5b6108dc8786855f61117b565b6108ea87868684600461173f565b6108f787878460016112b5565b50505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61092f61104e565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461099e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099590612043565b60405180910390fd5b6109aa838260026113ef565b505050565b6109b7611600565b5f600160146101000a81548160ff021916908315150217905550565b6109db61104e565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610a4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4190611fb5565b60405180910390fd5b610a57848484600261117b565b610a63848260046113ef565b50505050565b610a7161104e565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610ae0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad790611fb5565b60405180910390fd5b610aec8484845f61117b565b610af9848483600261117b565b50505050565b610b0761104e565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610b76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6d90611fb5565b60405180910390fd5b5f8490505f84861115610b8b57849150610b9a565b8585610b9791906122c6565b90505b610ba889888a85600361173f565b5f811115610bbd57610bbc8988835f61117b565b5b610bca898885600261117b565b610bd7898986600161117b565b505050505050505050565b610bea61104e565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610c59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5090611fb5565b60405180910390fd5b610c6484835f6113ef565b610c70848260026113ef565b50505050565b5f600160169054906101000a900460ff16159050808015610ca8575060018060159054906101000a900460ff1660ff16105b80610cd65750610cb73061187c565b158015610cd5575060018060159054906101000a900460ff1660ff16145b5b610d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0c90612369565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610d515760018060166101000a81548160ff0219169083151502179055505b610d5a82610567565b8015610db3575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610daa91906123d5565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610de761104e565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610e56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4d9061227b565b60405180910390fd5b5f8490505f84861115610e6b57849150610e7a565b8585610e7791906122c6565b90505b610e888a898b85600361173f565b610e968a898986600461173f565b5f811115610eab57610eaa8a89835f61117b565b5b610eb88a8a86600161117b565b50505050505050505050565b610ecc611600565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f319061245e565b60405180910390fd5b610f438161167e565b50565b610f4e611600565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610fbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb3906124ec565b60405180910390fd5b600160149054906101000a900460ff1661100b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110029061257a565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036110de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d5906124ec565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661111e61189e565b73ffffffffffffffffffffffffffffffffffffffff1614611174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116b90612608565b60405180910390fd5b6001905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663599efa6b85856040518363ffffffff1660e01b81526004016111d8929190612635565b6020604051808303815f875af11580156111f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112189190612691565b90508061125a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112519061272c565b60405180910390fd5b8460405161126891906127b6565b60405180910390207f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a984285858560026040516112a69493929190612885565b60405180910390a25050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388c2bdfe85856040518363ffffffff1660e01b8152600401611312929190612635565b6020604051808303815f875af115801561132e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113529190612691565b905080611394576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138b90612938565b60405180910390fd5b846040516113a291906127b6565b60405180910390207f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a984285858560036040516113e09493929190612885565b60405180910390a25050505050565b8160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231326040518263ffffffff1660e01b815260040161144a9190611c0d565b602060405180830381865afa158015611465573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611489919061296a565b10156114ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c190612a05565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635407e93c846040518263ffffffff1660e01b81526004016115259190612a23565b6020604051808303815f875af1158015611541573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115659190612691565b9050806115a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161159e90612aac565b60405180910390fd5b836040516115b591906127b6565b60405180910390207f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a98423285855f6040516115f29493929190612885565b60405180910390a250505050565b61160861189e565b73ffffffffffffffffffffffffffffffffffffffff16611626610900565b73ffffffffffffffffffffffffffffffffffffffff161461167c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167390612b14565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663065e86c88686866040518463ffffffff1660e01b815260040161179e93929190612b32565b6020604051808303815f875af11580156117ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117de9190612691565b905080611820576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161181790612bb1565b60405180910390fd5b8560405161182e91906127b6565b60405180910390207f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842858585600160405161186c9493929190612885565b60405180910390a2505050505050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611904826118be565b810181811067ffffffffffffffff82111715611923576119226118ce565b5b80604052505050565b5f6119356118a5565b905061194182826118fb565b919050565b5f67ffffffffffffffff8211156119605761195f6118ce565b5b611969826118be565b9050602081019050919050565b828183375f83830152505050565b5f61199661199184611946565b61192c565b9050828152602081018484840111156119b2576119b16118ba565b5b6119bd848285611976565b509392505050565b5f82601f8301126119d9576119d86118b6565b5b81356119e9848260208601611984565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611a1b826119f2565b9050919050565b611a2b81611a11565b8114611a35575f80fd5b50565b5f81359050611a4681611a22565b92915050565b5f819050919050565b611a5e81611a4c565b8114611a68575f80fd5b50565b5f81359050611a7981611a55565b92915050565b5f805f805f60a08688031215611a9857611a976118ae565b5b5f86013567ffffffffffffffff811115611ab557611ab46118b2565b5b611ac1888289016119c5565b9550506020611ad288828901611a38565b9450506040611ae388828901611a38565b9350506060611af488828901611a6b565b9250506080611b0588828901611a6b565b9150509295509295909350565b5f805f8060808587031215611b2a57611b296118ae565b5b5f85013567ffffffffffffffff811115611b4757611b466118b2565b5b611b53878288016119c5565b9450506020611b6487828801611a38565b9350506040611b7587828801611a6b565b9250506060611b8687828801611a6b565b91505092959194509250565b5f805f60608486031215611ba957611ba86118ae565b5b5f84013567ffffffffffffffff811115611bc657611bc56118b2565b5b611bd2868287016119c5565b9350506020611be386828701611a38565b9250506040611bf486828701611a6b565b9150509250925092565b611c0781611a11565b82525050565b5f602082019050611c205f830184611bfe565b92915050565b5f60208284031215611c3b57611c3a6118ae565b5b5f611c4884828501611a38565b91505092915050565b5f805f805f8060c08789031215611c6b57611c6a6118ae565b5b5f87013567ffffffffffffffff811115611c8857611c876118b2565b5b611c9489828a016119c5565b9650506020611ca589828a01611a38565b9550506040611cb689828a01611a38565b9450506060611cc789828a01611a6b565b9350506080611cd889828a01611a6b565b92505060a0611ce989828a01611a6b565b9150509295509295509295565b5f805f805f805f60e0888a031215611d1157611d106118ae565b5b5f88013567ffffffffffffffff811115611d2e57611d2d6118b2565b5b611d3a8a828b016119c5565b9750506020611d4b8a828b01611a38565b9650506040611d5c8a828b01611a38565b9550506060611d6d8a828b01611a38565b9450506080611d7e8a828b01611a6b565b93505060a0611d8f8a828b01611a6b565b92505060c0611da08a828b01611a6b565b91505092959891949750929550565b5f805f805f805f60e0888a031215611dca57611dc96118ae565b5b5f88013567ffffffffffffffff811115611de757611de66118b2565b5b611df38a828b016119c5565b9750506020611e048a828b01611a38565b9650506040611e158a828b01611a38565b9550506060611e268a828b01611a6b565b9450506080611e378a828b01611a6b565b93505060a0611e488a828b01611a6b565b92505060c0611e598a828b01611a6b565b91505092959891949750929550565b5f805f805f805f80610100898b031215611e8557611e846118ae565b5b5f89013567ffffffffffffffff811115611ea257611ea16118b2565b5b611eae8b828c016119c5565b9850506020611ebf8b828c01611a38565b9750506040611ed08b828c01611a38565b9650506060611ee18b828c01611a38565b9550506080611ef28b828c01611a6b565b94505060a0611f038b828c01611a6b565b93505060c0611f148b828c01611a6b565b92505060e0611f258b828c01611a6b565b9150509295985092959890939650565b5f82825260208201905092915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865204a4300000000000000000000000000000000000000602082015250565b5f611f9f602d83611f35565b9150611faa82611f45565b604082019050919050565b5f6020820190508181035f830152611fcc81611f93565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c65642062792074686520525000000000000000000000000000000000000000602082015250565b5f61202d602d83611f35565b915061203882611fd3565b604082019050919050565b5f6020820190508181035f83015261205a81612021565b9050919050565b7f4c696c657061645061796d656e74733a20546f6b656e2061646472657373206d5f8201527f75737420626520646566696e6564000000000000000000000000000000000000602082015250565b5f6120bb602e83611f35565b91506120c682612061565b604082019050919050565b5f6020820190508181035f8301526120e8816120af565b9050919050565b7f4c696c79706164546f6b656e3a2063616e4368616e6765546f6b656e416464725f8201527f6573732069732064697361626c65640000000000000000000000000000000000602082015250565b5f612149602f83611f35565b9150612154826120ef565b604082019050919050565b5f6020820190508181035f8301526121768161213d565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865205250206f72204a4300000000000000000000000000602082015250565b5f6121d7603383611f35565b91506121e28261217d565b604082019050919050565b5f6020820190508181035f830152612204816121cb565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865206d65646961746f7200000000000000000000000000602082015250565b5f612265603383611f35565b91506122708261220b565b604082019050919050565b5f6020820190508181035f83015261229281612259565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6122d082611a4c565b91506122db83611a4c565b92508282039050818111156122f3576122f2612299565b5b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f612353602e83611f35565b915061235e826122f9565b604082019050919050565b5f6020820190508181035f83015261238081612347565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6123bf6123ba6123b584612387565b61239c565b612390565b9050919050565b6123cf816123a5565b82525050565b5f6020820190506123e85f8301846123c6565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612448602683611f35565b9150612453826123ee565b604082019050919050565b5f6020820190508181035f8301526124758161243c565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6124d6603583611f35565b91506124e18261247c565b604082019050919050565b5f6020820190508181035f830152612503816124ca565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f612564603983611f35565b915061256f8261250a565b604082019050919050565b5f6020820190508181035f83015261259181612558565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6125f2603b83611f35565b91506125fd82612598565b604082019050919050565b5f6020820190508181035f83015261261f816125e6565b9050919050565b61262f81611a4c565b82525050565b5f6040820190506126485f830185611bfe565b6126556020830184612626565b9392505050565b5f8115159050919050565b6126708161265c565b811461267a575f80fd5b50565b5f8151905061268b81612667565b92915050565b5f602082840312156126a6576126a56118ae565b5b5f6126b38482850161267d565b91505092915050565b7f4c696c797061645061796d656e74733a20526566756e6420657363726f7720665f8201527f61696c6564000000000000000000000000000000000000000000000000000000602082015250565b5f612716602583611f35565b9150612721826126bc565b604082019050919050565b5f6020820190508181035f8301526127438161270a565b9050919050565b5f81519050919050565b5f81905092915050565b5f5b8381101561277b578082015181840152602081019050612760565b5f8484015250505050565b5f6127908261274a565b61279a8185612754565b93506127aa81856020860161275e565b80840191505092915050565b5f6127c18284612786565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061280a576128096127cc565b5b50565b5f81905061281a826127f9565b919050565b5f6128298261280d565b9050919050565b6128398161281f565b82525050565b600481106128505761284f6127cc565b5b50565b5f8190506128608261283f565b919050565b5f61286f82612853565b9050919050565b61287f81612865565b82525050565b5f6080820190506128985f830187611bfe565b6128a56020830186612626565b6128b26040830185612830565b6128bf6060830184612876565b95945050505050565b7f4c696c797061645061796d656e74733a20536c61736820657363726f772066615f8201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b5f612922602483611f35565b915061292d826128c8565b604082019050919050565b5f6020820190508181035f83015261294f81612916565b9050919050565b5f8151905061296481611a55565b92915050565b5f6020828403121561297f5761297e6118ae565b5b5f61298c84828501612956565b91505092915050565b7f4c696c797061645061796d656e74733a20496e73756666696369656e742062615f8201527f6c616e6365000000000000000000000000000000000000000000000000000000602082015250565b5f6129ef602583611f35565b91506129fa82612995565b604082019050919050565b5f6020820190508181035f830152612a1c816129e3565b9050919050565b5f602082019050612a365f830184612626565b92915050565b7f4c696c797061645061796d656e74733a2050617920657363726f77206661696c5f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f612a96602283611f35565b9150612aa182612a3c565b604082019050919050565b5f6020820190508181035f830152612ac381612a8a565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f612afe602083611f35565b9150612b0982612aca565b602082019050919050565b5f6020820190508181035f830152612b2b81612af2565b9050919050565b5f606082019050612b455f830186611bfe565b612b526020830185611bfe565b612b5f6040830184612626565b949350505050565b7f4c696c797061645061796d656e74733a20506179206a6f62206661696c6564005f82015250565b5f612b9b601f83611f35565b9150612ba682612b67565b602082019050919050565b5f6020820190508181035f830152612bc881612b8f565b905091905056fea2646970667358221220a0953ce3c7aaa4051f6b58fc0faa1e4a7bc77bdcee8f969c809c00ebcfd4f50e64736f6c63430008150033",
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

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xddce511b.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xddce511b.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xddce511b.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x888aa14c.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x888aa14c.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x888aa14c.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
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
	DealId    common.Hash
	Payee     common.Address
	Amount    *big.Int
	Reason    uint8
	Direction uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) FilterPayment(opts *bind.FilterOpts, dealId []string) (*PaymentsPaymentIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Payment", dealIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsPaymentIterator{contract: _Payments.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PaymentsPayment, dealId []string) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Payment", dealIdRule)
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
// Solidity: event Payment(string indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) ParsePayment(log types.Log) (*PaymentsPayment, error) {
	event := new(PaymentsPayment)
	if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
