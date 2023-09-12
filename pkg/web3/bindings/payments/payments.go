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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLilypadPayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600360146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61288080620001425f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c8063840e9d4e116100ab578063d934b4421161006f578063d934b4421461029e578063dd1e6201146102ba578063e75b2d5e146102d6578063f2fde38b146102f2578063f3d3d4481461030e5761012a565b8063840e9d4e146102225780638da5cb5b1461023e5780638eed1e9c1461025c578063a470295814610278578063c4d66de8146102825761012a565b80634bc28da1116100f25780634bc28da1146101ba5780635559adf2146101c4578063715018a6146101e0578063759a0f80146101ea5780637ec34959146102065761012a565b80630b01498d1461012e5780630d92aa591461014a578063245469fc1461016657806326a4e8d2146101825780634a5d593d1461019e575b5f80fd5b61014860048036038101906101439190611840565b61032a565b005b610164600480360381019061015f91906118dd565b6103d1565b005b610180600480360381019061017b91906118dd565b610467565b005b61019c60048036038101906101979190611941565b6104fb565b005b6101b860048036038101906101b391906118dd565b610643565b005b6101c26106d9565b005b6101de60048036038101906101d9919061196c565b6106fd565b005b6101e8610785565b005b61020460048036038101906101ff919061196c565b610798565b005b610220600480360381019061021b91906119bc565b610821565b005b61023c60048036038101906102379190611a59565b610904565b005b61024661099c565b6040516102539190611adf565b60405180910390f35b61027660048036038101906102719190611a59565b6109c3565b005b610280610a67565b005b61029c60048036038101906102979190611941565b610a8b565b005b6102b860048036038101906102b391906118dd565b610bcc565b005b6102d460048036038101906102cf9190611af8565b610c62565b005b6102f060048036038101906102eb9190611b81565b610d3d565b005b61030c60048036038101906103079190611941565b610e22565b005b61032860048036038101906103239190611941565b610ea4565b005b610332610fac565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146103a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039890611cb2565b60405180910390fd5b6103ad8786855f6110d9565b6103bb8786868460046111fe565b6103c88787846001611326565b50505050505050565b6103d9610fac565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610448576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043f90611d40565b60405180910390fd5b61045584848460026110d9565b6104618482600461144b565b50505050565b61046f610fac565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146104de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d590611d40565b60405180910390fd5b6104e984835f61144b565b6104f58482600261144b565b50505050565b610503611647565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610571576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056890611dce565b60405180910390fd5b600360149054906101000a900460ff166105c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b790611e5c565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61064b610fac565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146106ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b190611d40565b60405180910390fd5b6106c68484845f6110d9565b6106d384848360026110d9565b50505050565b6106e1611647565b5f600360146101000a81548160ff021916908315150217905550565b610705610fac565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076b90611eea565b60405180910390fd5b6107808382600261144b565b505050565b61078d611647565b6107965f6116c5565b565b6107a0610fac565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461080f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080690611eea565b60405180910390fd5b61081c83838360026110d9565b505050565b610829610fac565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610898576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088f90611d40565b60405180910390fd5b5f8490505f848611156108ad578491506108bc565b85856108b99190611f35565b90505b6108ca89888a8560036111fe565b5f8111156108df576108de8988835f6110d9565b5b6108ec89888560026110d9565b6108f989898660016110d9565b505050505050505050565b61090c610fac565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461097b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097290611eea565b60405180910390fd5b61098885858460016110d9565b6109958584836002611326565b5050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6109cb610fac565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610a3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3190611d40565b60405180910390fd5b610a468584845f6110d9565b610a5385848360026110d9565b610a608585836002611326565b5050505050565b610a6f611647565b5f600160146101000a81548160ff021916908315150217905550565b5f600160169054906101000a900460ff16159050808015610abd575060018060159054906101000a900460ff1660ff16105b80610aeb5750610acc30611786565b158015610aea575060018060159054906101000a900460ff1660ff16145b5b610b2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2190611fd8565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610b665760018060166101000a81548160ff0219169083151502179055505b610b6f826104fb565b8015610bc8575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610bbf9190612044565b60405180910390a15b5050565b610bd4610fac565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610c43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3a90611eea565b60405180910390fd5b610c4f8483600161144b565b610c5c84848360026110d9565b50505050565b610c6a610fac565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610cd057508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b610d0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d06906120cd565b60405180910390fd5b610d1c86868460016110d9565b610d288685855f6110d9565b610d3586858360046110d9565b505050505050565b610d45610fac565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610db4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dab90611cb2565b60405180910390fd5b5f8490505f84861115610dc957849150610dd8565b8585610dd59190611f35565b90505b610de68a898b8560036111fe565b610df48a89898660046111fe565b5f811115610e0957610e088a89835f6110d9565b5b610e168a8a8660016110d9565b50505050505050505050565b610e2a611647565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8f9061215b565b60405180910390fd5b610ea1816116c5565b50565b610eac611647565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f11906121e9565b60405180910390fd5b600160149054906101000a900460ff16610f69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6090612277565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361103c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611033906121e9565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661107c6117a8565b73ffffffffffffffffffffffffffffffffffffffff16146110d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c990612305565b60405180910390fd5b6001905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663599efa6b85856040518363ffffffff1660e01b8152600401611136929190612332565b6020604051808303815f875af1158015611152573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611176919061238e565b9050806111b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111af90612429565b60405180910390fd5b847fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c56385858560026040516111ef9493929190612500565b60405180910390a25050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663065e86c88686866040518463ffffffff1660e01b815260040161125d93929190612543565b6020604051808303815f875af1158015611279573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061129d919061238e565b9050806112df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d6906125c2565b60405180910390fd5b857fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c56385858560016040516113169493929190612500565b60405180910390a2505050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388c2bdfe85856040518363ffffffff1660e01b8152600401611383929190612332565b6020604051808303815f875af115801561139f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113c3919061238e565b905080611405576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fc90612650565b60405180910390fd5b847fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c563858585600360405161143c9493929190612500565b60405180910390a25050505050565b8160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231326040518263ffffffff1660e01b81526004016114a69190611adf565b602060405180830381865afa1580156114c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114e59190612682565b1015611526576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151d9061271d565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635407e93c846040518263ffffffff1660e01b8152600401611581919061273b565b6020604051808303815f875af115801561159d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115c1919061238e565b905080611603576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fa906127c4565b60405180910390fd5b837fb337867dbd6dc6eedc1de53ec25233a8b3fe6f6da8a468a83bf388537143c5633285855f6040516116399493929190612500565b60405180910390a250505050565b61164f6117a8565b73ffffffffffffffffffffffffffffffffffffffff1661166d61099c565b73ffffffffffffffffffffffffffffffffffffffff16146116c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116ba9061282c565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f80fd5b5f819050919050565b6117c5816117b3565b81146117cf575f80fd5b50565b5f813590506117e0816117bc565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61180f826117e6565b9050919050565b61181f81611805565b8114611829575f80fd5b50565b5f8135905061183a81611816565b92915050565b5f805f805f805f60e0888a03121561185b5761185a6117af565b5b5f6118688a828b016117d2565b97505060206118798a828b0161182c565b965050604061188a8a828b0161182c565b955050606061189b8a828b0161182c565b94505060806118ac8a828b016117d2565b93505060a06118bd8a828b016117d2565b92505060c06118ce8a828b016117d2565b91505092959891949750929550565b5f805f80608085870312156118f5576118f46117af565b5b5f611902878288016117d2565b94505060206119138782880161182c565b9350506040611924878288016117d2565b9250506060611935878288016117d2565b91505092959194509250565b5f60208284031215611956576119556117af565b5b5f6119638482850161182c565b91505092915050565b5f805f60608486031215611983576119826117af565b5b5f611990868287016117d2565b93505060206119a18682870161182c565b92505060406119b2868287016117d2565b9150509250925092565b5f805f805f805f60e0888a0312156119d7576119d66117af565b5b5f6119e48a828b016117d2565b97505060206119f58a828b0161182c565b9650506040611a068a828b0161182c565b9550506060611a178a828b016117d2565b9450506080611a288a828b016117d2565b93505060a0611a398a828b016117d2565b92505060c0611a4a8a828b016117d2565b91505092959891949750929550565b5f805f805f60a08688031215611a7257611a716117af565b5b5f611a7f888289016117d2565b9550506020611a908882890161182c565b9450506040611aa18882890161182c565b9350506060611ab2888289016117d2565b9250506080611ac3888289016117d2565b9150509295509295909350565b611ad981611805565b82525050565b5f602082019050611af25f830184611ad0565b92915050565b5f805f805f8060c08789031215611b1257611b116117af565b5b5f611b1f89828a016117d2565b9650506020611b3089828a0161182c565b9550506040611b4189828a0161182c565b9450506060611b5289828a016117d2565b9350506080611b6389828a016117d2565b92505060a0611b7489828a016117d2565b9150509295509295509295565b5f805f805f805f80610100898b031215611b9e57611b9d6117af565b5b5f611bab8b828c016117d2565b9850506020611bbc8b828c0161182c565b9750506040611bcd8b828c0161182c565b9650506060611bde8b828c0161182c565b9550506080611bef8b828c016117d2565b94505060a0611c008b828c016117d2565b93505060c0611c118b828c016117d2565b92505060e0611c228b828c016117d2565b9150509295985092959890939650565b5f82825260208201905092915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865206d65646961746f7200000000000000000000000000602082015250565b5f611c9c603383611c32565b9150611ca782611c42565b604082019050919050565b5f6020820190508181035f830152611cc981611c90565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865204a4300000000000000000000000000000000000000602082015250565b5f611d2a602d83611c32565b9150611d3582611cd0565b604082019050919050565b5f6020820190508181035f830152611d5781611d1e565b9050919050565b7f4c696c657061645061796d656e74733a20546f6b656e2061646472657373206d5f8201527f75737420626520646566696e6564000000000000000000000000000000000000602082015250565b5f611db8602e83611c32565b9150611dc382611d5e565b604082019050919050565b5f6020820190508181035f830152611de581611dac565b9050919050565b7f4c696c79706164546f6b656e3a2063616e4368616e6765546f6b656e416464725f8201527f6573732069732064697361626c65640000000000000000000000000000000000602082015250565b5f611e46602f83611c32565b9150611e5182611dec565b604082019050919050565b5f6020820190508181035f830152611e7381611e3a565b9050919050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c65642062792074686520525000000000000000000000000000000000000000602082015250565b5f611ed4602d83611c32565b9150611edf82611e7a565b604082019050919050565b5f6020820190508181035f830152611f0181611ec8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f3f826117b3565b9150611f4a836117b3565b9250828203905081811115611f6257611f61611f08565b5b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611fc2602e83611c32565b9150611fcd82611f68565b604082019050919050565b5f6020820190508181035f830152611fef81611fb6565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61202e61202961202484611ff6565b61200b565b611fff565b9050919050565b61203e81612014565b82525050565b5f6020820190506120575f830184612035565b92915050565b7f4c696c797061645061796d656e74733a2043616e206f6e6c792062652063616c5f8201527f6c656420627920746865205250206f72204a4300000000000000000000000000602082015250565b5f6120b7603383611c32565b91506120c28261205d565b604082019050919050565b5f6020820190508181035f8301526120e4816120ab565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612145602683611c32565b9150612150826120eb565b604082019050919050565b5f6020820190508181035f83015261217281612139565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6121d3603583611c32565b91506121de82612179565b604082019050919050565b5f6020820190508181035f830152612200816121c7565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f612261603983611c32565b915061226c82612207565b604082019050919050565b5f6020820190508181035f83015261228e81612255565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6122ef603b83611c32565b91506122fa82612295565b604082019050919050565b5f6020820190508181035f83015261231c816122e3565b9050919050565b61232c816117b3565b82525050565b5f6040820190506123455f830185611ad0565b6123526020830184612323565b9392505050565b5f8115159050919050565b61236d81612359565b8114612377575f80fd5b50565b5f8151905061238881612364565b92915050565b5f602082840312156123a3576123a26117af565b5b5f6123b08482850161237a565b91505092915050565b7f4c696c797061645061796d656e74733a20526566756e6420657363726f7720665f8201527f61696c6564000000000000000000000000000000000000000000000000000000602082015250565b5f612413602583611c32565b915061241e826123b9565b604082019050919050565b5f6020820190508181035f83015261244081612407565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6005811061248557612484612447565b5b50565b5f81905061249582612474565b919050565b5f6124a482612488565b9050919050565b6124b48161249a565b82525050565b600481106124cb576124ca612447565b5b50565b5f8190506124db826124ba565b919050565b5f6124ea826124ce565b9050919050565b6124fa816124e0565b82525050565b5f6080820190506125135f830187611ad0565b6125206020830186612323565b61252d60408301856124ab565b61253a60608301846124f1565b95945050505050565b5f6060820190506125565f830186611ad0565b6125636020830185611ad0565b6125706040830184612323565b949350505050565b7f4c696c797061645061796d656e74733a20506179206a6f62206661696c6564005f82015250565b5f6125ac601f83611c32565b91506125b782612578565b602082019050919050565b5f6020820190508181035f8301526125d9816125a0565b9050919050565b7f4c696c797061645061796d656e74733a20536c61736820657363726f772066615f8201527f696c656400000000000000000000000000000000000000000000000000000000602082015250565b5f61263a602483611c32565b9150612645826125e0565b604082019050919050565b5f6020820190508181035f8301526126678161262e565b9050919050565b5f8151905061267c816117bc565b92915050565b5f60208284031215612697576126966117af565b5b5f6126a48482850161266e565b91505092915050565b7f4c696c797061645061796d656e74733a20496e73756666696369656e742062615f8201527f6c616e6365000000000000000000000000000000000000000000000000000000602082015250565b5f612707602583611c32565b9150612712826126ad565b604082019050919050565b5f6020820190508181035f830152612734816126fb565b9050919050565b5f60208201905061274e5f830184612323565b92915050565b7f4c696c797061645061796d656e74733a2050617920657363726f77206661696c5f8201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b5f6127ae602283611c32565b91506127b982612754565b604082019050919050565b5f6020820190508181035f8301526127db816127a2565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f612816602083611c32565b9150612821826127e2565b602082019050919050565b5f6020820190508181035f8301526128438161280a565b905091905056fea2646970667358221220463ab79027ddbc70e947db6effe6d82c6620fc19f89a5de28dbace5a603f4cb764736f6c63430008150033",
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
