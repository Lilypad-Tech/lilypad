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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600360146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6128d180620001425f395ff3fe608060405234801561000f575f80fd5b5060043610610135575f3560e01c8063840e9d4e116100b6578063c57380a21161007a578063c57380a2146102a9578063d934b442146102c7578063dd1e6201146102e3578063e75b2d5e146102ff578063f2fde38b1461031b578063f3d3d4481461033757610135565b8063840e9d4e1461022d5780638da5cb5b146102495780638eed1e9c14610267578063a470295814610283578063c4d66de81461028d57610135565b80634bc28da1116100fd5780634bc28da1146101c55780635559adf2146101cf578063715018a6146101eb578063759a0f80146101f55780637ec349591461021157610135565b80630b01498d146101395780630d92aa5914610155578063245469fc1461017157806326a4e8d21461018d5780634a5d593d146101a9575b5f80fd5b610153600480360381019061014e9190611891565b610353565b005b61016f600480360381019061016a919061192e565b6103fa565b005b61018b6004803603810190610186919061192e565b610490565b005b6101a760048036038101906101a29190611992565b610524565b005b6101c360048036038101906101be919061192e565b61066c565b005b6101cd610702565b005b6101e960048036038101906101e491906119bd565b610726565b005b6101f36107ae565b005b61020f600480360381019061020a91906119bd565b6107c1565b005b61022b60048036038101906102269190611a0d565b61084a565b005b61024760048036038101906102429190611aaa565b61092d565b005b6102516109c5565b60405161025e9190611b30565b60405180910390f35b610281600480360381019061027c9190611aaa565b6109ec565b005b61028b610a90565b005b6102a760048036038101906102a29190611992565b610ab4565b005b6102b1610bf5565b6040516102be9190611b30565b60405180910390f35b6102e160048036038101906102dc919061192e565b610c1d565b005b6102fd60048036038101906102f89190611b49565b610cb3565b005b61031960048036038101906103149190611bd2565b610d8e565b005b61033560048036038101906103309190611992565b610e73565b005b610351600480360381019061034c9190611992565b610ef5565b005b61035b610ffd565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146103ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c190611d03565b60405180910390fd5b6103d68786855f61112a565b6103e487868684600461124f565b6103f18787846001611377565b50505050505050565b610402610ffd565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046890611d91565b60405180910390fd5b61047e848484600261112a565b61048a8482600461149c565b50505050565b610498610ffd565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fe90611d91565b60405180910390fd5b61051284835f61149c565b61051e8482600261149c565b50505050565b61052c611698565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361059a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059190611e1f565b60405180910390fd5b600360149054906101000a900460ff166105e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e090611ead565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610674610ffd565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146106e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106da90611d91565b60405180910390fd5b6106ef8484845f61112a565b6106fc848483600261112a565b50505050565b61070a611698565b5f600360146101000a81548160ff021916908315150217905550565b61072e610ffd565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461079d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079490611f3b565b60405180910390fd5b6107a98382600261149c565b505050565b6107b6611698565b6107bf5f611716565b565b6107c9610ffd565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f90611f3b565b60405180910390fd5b610845838383600261112a565b505050565b610852610ffd565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146108c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b890611d91565b60405180910390fd5b5f8490505f848611156108d6578491506108e5565b85856108e29190611f86565b90505b6108f389888a85600361124f565b5f811115610908576109078988835f61112a565b5b610915898885600261112a565b610922898986600161112a565b505050505050505050565b610935610ffd565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b90611f3b565b60405180910390fd5b6109b1858584600161112a565b6109be8584836002611377565b5050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6109f4610ffd565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610a63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5a90611d91565b60405180910390fd5b610a6f8584845f61112a565b610a7c858483600261112a565b610a898585836002611377565b5050505050565b610a98611698565b5f600160146101000a81548160ff021916908315150217905550565b5f600160169054906101000a900460ff16159050808015610ae6575060018060159054906101000a900460ff1660ff16105b80610b145750610af5306117d7565b158015610b13575060018060159054906101000a900460ff1660ff16145b5b610b53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4a90612029565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610b8f5760018060166101000a81548160ff0219169083151502179055505b610b9882610524565b8015610bf1575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610be89190612095565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610c25610ffd565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610c94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8b90611f3b565b60405180910390fd5b610ca08483600161149c565b610cad848483600261112a565b50505050565b610cbb610ffd565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610d2157508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b610d60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d579061211e565b60405180910390fd5b610d6d868684600161112a565b610d798685855f61112a565b610d86868583600461112a565b505050505050565b610d96610ffd565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610e05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfc90611d03565b60405180910390fd5b5f8490505f84861115610e1a57849150610e29565b8585610e269190611f86565b90505b610e378a898b85600361124f565b610e458a898986600461124f565b5f811115610e5a57610e598a89835f61112a565b5b610e678a8a86600161112a565b50505050505050505050565b610e7b611698565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ee9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee0906121ac565b60405180910390fd5b610ef281611716565b50565b610efd611698565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f629061223a565b60405180910390fd5b600160149054906101000a900460ff16610fba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb1906122c8565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361108d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110849061223a565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166110cd6117f9565b73ffffffffffffffffffffffffffffffffffffffff1614611123576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111a90612356565b60405180910390fd5b6001905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663599efa6b85856040518363ffffffff1660e01b8152600401611187929190612383565b6020604051808303815f875af11580156111a3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c791906123df565b905080611209576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112009061247a565b60405180910390fd5b847fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c56385858560026040516112409493929190612551565b60405180910390a25050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663065e86c88686866040518463ffffffff1660e01b81526004016112ae93929190612594565b6020604051808303815f875af11580156112ca573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112ee91906123df565b905080611330576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132790612613565b60405180910390fd5b857fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c56385858560016040516113679493929190612551565b60405180910390a2505050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388c2bdfe85856040518363ffffffff1660e01b81526004016113d4929190612383565b6020604051808303815f875af11580156113f0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061141491906123df565b905080611456576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144d906126a1565b60405180910390fd5b847fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563858585600360405161148d9493929190612551565b60405180910390a25050505050565b8160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231326040518263ffffffff1660e01b81526004016114f79190611b30565b602060405180830381865afa158015611512573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061153691906126d3565b1015611577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156e9061276e565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635407e93c846040518263ffffffff1660e01b81526004016115d2919061278c565b6020604051808303815f875af11580156115ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061161291906123df565b905080611654576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164b90612815565b60405180910390fd5b837fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c5633285855f60405161168a9493929190612551565b60405180910390a250505050565b6116a06117f9565b73ffffffffffffffffffffffffffffffffffffffff166116be6109c5565b73ffffffffffffffffffffffffffffffffffffffff1614611714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161170b9061287d565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f80fd5b5f819050919050565b61181681611804565b8114611820575f80fd5b50565b5f813590506118318161180d565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61186082611837565b9050919050565b61187081611856565b811461187a575f80fd5b50565b5f8135905061188b81611867565b92915050565b5f805f805f805f60e0888a0312156118ac576118ab611800565b5b5f6118b98a828b01611823565b97505060206118ca8a828b0161187d565b96505060406118db8a828b0161187d565b95505060606118ec8a828b0161187d565b94505060806118fd8a828b01611823565b93505060a061190e8a828b01611823565b92505060c061191f8a828b01611823565b91505092959891949750929550565b5f805f806080858703121561194657611945611800565b5b5f61195387828801611823565b94505060206119648782880161187d565b935050604061197587828801611823565b925050606061198687828801611823565b91505092959194509250565b5f602082840312156119a7576119a6611800565b5b5f6119b48482850161187d565b91505092915050565b5f805f606084860312156119d4576119d3611800565b5b5f6119e186828701611823565b93505060206119f28682870161187d565b9250506040611a0386828701611823565b9150509250925092565b5f805f805f805f60e0888a031215611a2857611a27611800565b5b5f611a358a828b01611823565b9750506020611a468a828b0161187d565b9650506040611a578a828b0161187d565b9550506060611a688a828b01611823565b9450506080611a798a828b01611823565b93505060a0611a8a8a828b01611823565b92505060c0611a9b8a828b01611823565b91505092959891949750929550565b5f805f805f60a08688031215611ac357611ac2611800565b5b5f611ad088828901611823565b9550506020611ae18882890161187d565b9450506040611af28882890161187d565b9350506060611b0388828901611823565b9250506080611b1488828901611823565b9150509295509295909350565b611b2a81611856565b82525050565b5f602082019050611b435f830184611b21565b92915050565b5f805f805f8060c08789031215611b6357611b62611800565b5b5f611b7089828a01611823565b9650506020611b8189828a0161187d565b9550506040611b9289828a0161187d565b9450506060611ba389828a01611823565b9350506080611bb489828a01611823565b92505060a0611bc589828a01611823565b9150509295509295509295565b5f805f805f805f80610100898b031215611bef57611bee611800565b5b5f611bfc8b828c01611823565b9850506020611c0d8b828c0161187d565b9750506040611c1e8b828c0161187d565b9650506060611c2f8b828c0161187d565b9550506080611c408b828c01611823565b94505060a0611c518b828c01611823565b93505060c0611c628b828c01611823565b92505060e0611c738b828c01611823565b9150509295985092959890939650565b5f82825260208201905092915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865206d65646961746f7200000000000000000000000000602082015250565b5f611ced603383611c83565b9150611cf882611c93565b604082019050919050565b5f6020820190508181035f830152611d1a81611ce1565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865204a4300000000000000000000000000000000000000602082015250565b5f611d7b602d83611c83565b9150611d8682611d21565b604082019050919050565b5f6020820190508181035f830152611da881611d6f565b9050919050565b7f4c696c657061645061796d656e74733a20546f6b656e2061646472657373206d5f8201527f75737420626520646566696e6564000000000000000000000000000000000000602082015250565b5f611e09602e83611c83565b9150611e1482611daf565b604082019050919050565b5f6020820190508181035f830152611e3681611dfd565b9050919050565b7f4c696c79706164546f6b656e3a2063616e4368616e6765546f6b656e416464725f8201527f6573732069732064697361626c65640000000000000000000000000000000000602082015250565b5f611e97602f83611c83565b9150611ea282611e3d565b604082019050919050565b5f6020820190508181035f830152611ec481611e8b565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c65642062792074686520525000000000000000000000000000000000000000602082015250565b5f611f25602d83611c83565b9150611f3082611ecb565b604082019050919050565b5f6020820190508181035f830152611f5281611f19565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f9082611804565b9150611f9b83611804565b9250828203905081811115611fb357611fb2611f59565b5b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f612013602e83611c83565b915061201e82611fb9565b604082019050919050565b5f6020820190508181035f83015261204081612007565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61207f61207a61207584612047565b61205c565b612050565b9050919050565b61208f81612065565b82525050565b5f6020820190506120a85f830184612086565b92915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865205250206f72204a4300000000000000000000000000602082015250565b5f612108603383611c83565b9150612113826120ae565b604082019050919050565b5f6020820190508181035f830152612135816120fc565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612196602683611c83565b91506121a18261213c565b604082019050919050565b5f6020820190508181035f8301526121c38161218a565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f612224603583611c83565b915061222f826121ca565b604082019050919050565b5f6020820190508181035f83015261225181612218565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f6122b2603983611c83565b91506122bd82612258565b604082019050919050565b5f6020820190508181035f8301526122df816122a6565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f612340603b83611c83565b915061234b826122e6565b604082019050919050565b5f6020820190508181035f83015261236d81612334565b9050919050565b61237d81611804565b82525050565b5f6040820190506123965f830185611b21565b6123a36020830184612374565b9392505050565b5f8115159050919050565b6123be816123aa565b81146123c8575f80fd5b50565b5f815190506123d9816123b5565b92915050565b5f602082840312156123f4576123f3611800565b5b5f612401848285016123cb565b91505092915050565b7f4c696c797061645061796d656e74733a20526566756e6420657363726f7720665f8201527f61696c6564000000000000000000000000000000000000000000000000000000602082015250565b5f612464602583611c83565b915061246f8261240a565b604082019050919050565b5f6020820190508181035f83015261249181612458565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600581106124d6576124d5612498565b5b50565b5f8190506124e6826124c5565b919050565b5f6124f5826124d9565b9050919050565b612505816124eb565b82525050565b6004811061251c5761251b612498565b5b50565b5f81905061252c8261250b565b919050565b5f61253b8261251f565b9050919050565b61254b81612531565b82525050565b5f6080820190506125645f830187611b21565b6125716020830186612374565b61257e60408301856124fc565b61258b6060830184612542565b95945050505050565b5f6060820190506125a75f830186611b21565b6125b46020830185611b21565b6125c16040830184612374565b949350505050565b7f4c696c797061645061796d656e74733a20506179206a6f62206661696c6564005f82015250565b5f6125fd601f83611c83565b9150612608826125c9565b602082019050919050565b5f6020820190508181035f83015261262a816125f1565b9050919050565b7f4c696c797061645061796d656e74733a20536c61736820657363726f772066615f8201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b5f61268b602483611c83565b915061269682612631565b604082019050919050565b5f6020820190508181035f8301526126b88161267f565b9050919050565b5f815190506126cd8161180d565b92915050565b5f602082840312156126e8576126e7611800565b5b5f6126f5848285016126bf565b91505092915050565b7f4c696c797061645061796d656e74733a20496e73756666696369656e742062615f8201527f6c616e6365000000000000000000000000000000000000000000000000000000602082015250565b5f612758602583611c83565b9150612763826126fe565b604082019050919050565b5f6020820190508181035f8301526127858161274c565b9050919050565b5f60208201905061279f5f830184612374565b92915050565b7f4c696c797061645061796d656e74733a2050617920657363726f77206661696c5f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f6127ff602283611c83565b915061280a826127a5565b604082019050919050565b5f6020820190508181035f83015261282c816127f3565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f612867602083611c83565b915061287282612833565b602082019050919050565b5f6020820190508181035f8301526128948161285b565b905091905056fea2646970667358221220bbce75231f2d3085801f382d7cb7a37c7732d0eaf57811207bec793c9a0c358164736f6c63430008150033",
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

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "acceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x7ec34959.
//
// Solidity: function acceptResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "addResult", dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0xd934b442.
//
// Solidity: function addResult(uint256 dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AddResult(dealId *big.Int, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x245469fc.
//
// Solidity: function agreeJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x5559adf2.
//
// Solidity: function agreeResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "checkResult", dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckResult(&_Payments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0x0d92aa59.
//
// Solidity: function checkResult(uint256 dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) CheckResult(dealId *big.Int, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
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

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0xe75b2d5e.
//
// Solidity: function mediationAcceptResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationAcceptResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, mediator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x0b01498d.
//
// Solidity: function mediationRejectResult(uint256 dealId, address resourceProvider, address jobCreator, address mediator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationRejectResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, mediator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
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

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0x4a5d593d.
//
// Solidity: function timeoutAgreeRefundJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundJobCreator(opts *bind.TransactOpts, dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0x4a5d593d.
//
// Solidity: function timeoutAgreeRefundJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0x4a5d593d.
//
// Solidity: function timeoutAgreeRefundJobCreator(uint256 dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundJobCreator(dealId *big.Int, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x759a0f80.
//
// Solidity: function timeoutAgreeRefundResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundResourceProvider(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x759a0f80.
//
// Solidity: function timeoutAgreeRefundResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x759a0f80.
//
// Solidity: function timeoutAgreeRefundResourceProvider(uint256 dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundResourceProvider(dealId *big.Int, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutJudgeResults(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutJudgeResults", dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutJudgeResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x840e9d4e.
//
// Solidity: function timeoutJudgeResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutJudgeResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutMediateResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xdd1e6201.
//
// Solidity: function timeoutMediateResult(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) TimeoutMediateResult(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutSubmitResults(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutSubmitResults", dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutSubmitResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutSubmitResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x8eed1e9c.
//
// Solidity: function timeoutSubmitResults(uint256 dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutSubmitResults(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
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
	DealId    *big.Int
	Payee     common.Address
	Amount    *big.Int
	Reason    uint8
	Direction uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) FilterPayment(opts *bind.FilterOpts, dealId []*big.Int) (*PaymentsPaymentIterator, error) {

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

// WatchPayment is a free log subscription operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PaymentsPayment, dealId []*big.Int) (event.Subscription, error) {

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

// ParsePayment is a log parse operation binding the contract event 0xb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563.
//
// Solidity: event Payment(uint256 indexed dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) ParsePayment(log types.Log) (*PaymentsPayment, error) {
	event := new(PaymentsPayment)
	if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
