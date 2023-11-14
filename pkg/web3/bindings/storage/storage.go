// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealCreatedAt            *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutAgreeAt           *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId   string
	Members  SharedStructsDealMembers
	Timeouts SharedStructsDealTimeouts
	Pricing  SharedStructsDealPricing
}

// SharedStructsDealMembers is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealMembers struct {
	Solver           common.Address
	JobCreator       common.Address
	ResourceProvider common.Address
	Mediators        []common.Address
}

// SharedStructsDealPricing is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPricing struct {
	InstructionPrice          *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsDealTimeout is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeout struct {
	Timeout    *big.Int
	Collateral *big.Int
}

// SharedStructsDealTimeouts is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeouts struct {
	Agree          SharedStructsDealTimeout
	SubmitResults  SharedStructsDealTimeout
	JudgeResults   SharedStructsDealTimeout
	MediateResults SharedStructsDealTimeout
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	DealId           string
	ResultsId        string
	DataId           string
	InstructionCount *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"name\":\"Deal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61511c80620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101a7575f3560e01c80638224ce5f116100f7578063cdd82d1d11610095578063e850be371161006f578063e850be37146104d9578063ec95b967146104f5578063f2fde38b14610525578063f3d3d44814610541576101a7565b8063cdd82d1d1461045d578063e70791801461048d578063e7b957d1146104bd576101a7565b8063a4702958116100d1578063a4702958146103d5578063a6370b0e146103df578063b050e74b1461040f578063c57380a21461043f576101a7565b80638224ce5f1461036b578063824518aa1461039b5780638da5cb5b146103b7576101a7565b8063498cc70d1161016457806373db5c6a1161013e57806373db5c6a146102f9578063795f9abf1461032957806380ffdfe0146103455780638129fc1c14610361576101a7565b8063498cc70d146102a3578063511a9f68146102d3578063715018a6146102ef576101a7565b806311d5af33146101ab5780632244ad2b146101db578063297f9e551461020b5780633955548e146102275780633c4135da1461025757806346834d1e14610287575b5f80fd5b6101c560048036038101906101c0919061310a565b61055d565b6040516101d2919061327a565b60405180910390f35b6101f560048036038101906101f091906133c6565b61066e565b6040516102029190613427565b60405180910390f35b610225600480360381019061022091906133c6565b610685565b005b610241600480360381019061023c9190613473565b61070c565b60405161024e91906135a8565b60405180910390f35b610271600480360381019061026c91906133c6565b610a10565b60405161027e9190613759565b60405180910390f35b6102a1600480360381019061029c91906133c6565b610be7565b005b6102bd60048036038101906102b891906133c6565b610c6e565b6040516102ca91906135a8565b60405180910390f35b6102ed60048036038101906102e891906133c6565b610e5f565b005b6102f7610ee6565b005b610313600480360381019061030e91906133c6565b610ef9565b6040516103209190613782565b60405180910390f35b610343600480360381019061033e91906133c6565b610f51565b005b61035f600480360381019061035a91906133c6565b610fd7565b005b61036961105e565b005b610385600480360381019061038091906133c6565b611195565b6040516103929190613782565b60405180910390f35b6103b560048036038101906103b091906133c6565b6111d5565b005b6103bf61125c565b6040516103cc91906137aa565b60405180910390f35b6103dd611283565b005b6103f960048036038101906103f49190613a58565b6112a7565b6040516104069190613d4a565b60405180910390f35b61042960048036038101906104249190613d8d565b6119d3565b6040516104369190613427565b60405180910390f35b610447611a6b565b60405161045491906137aa565b60405180910390f35b610477600480360381019061047291906133c6565b611a93565b6040516104849190613759565b60405180910390f35b6104a760048036038101906104a291906133c6565b611b87565b6040516104b49190613d4a565b60405180910390f35b6104d760048036038101906104d291906133c6565b611ed5565b005b6104f360048036038101906104ee91906133c6565b611f5c565b005b61050f600480360381019061050a91906133c6565b611fe3565b60405161051c9190613759565b60405180910390f35b61053f600480360381019061053a919061310a565b6121ba565b005b61055b6004803603810190610556919061310a565b61223c565b005b606060045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610663578382905f5260205f200180546105d890613e14565b80601f016020809104026020016040519081016040528092919081815260200182805461060490613e14565b801561064f5780601f106106265761010080835404028352916020019161064f565b820191905f5260205f20905b81548152906001019060200180831161063257829003601f168201915b5050505050815260200190600101906105bb565b505050509050919050565b5f8061067983611b87565b5f015151119050919050565b61068d612344565b506106998160026119d3565b6106d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cf90613e9e565b60405180910390fd5b426005826040516106e99190613ef6565b908152602001604051809103902060060181905550610709816003612471565b50565b610714612e46565b61071c612344565b506107288560016119d3565b610767576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e90613f56565b60405180910390fd5b426005866040516107789190613ef6565b908152602001604051809103902060050181905550610798856002612471565b6040518060800160405280868152602001858152602001848152602001838152506006866040516107c99190613ef6565b90815260200160405180910390205f820151815f0190816107ea9190614111565b5060208201518160010190816108009190614111565b5060408201518160020190816108169190614111565b50606082015181600301559050506006856040516108349190613ef6565b90815260200160405180910390206040518060800160405290815f8201805461085c90613e14565b80601f016020809104026020016040519081016040528092919081815260200182805461088890613e14565b80156108d35780601f106108aa576101008083540402835291602001916108d3565b820191905f5260205f20905b8154815290600101906020018083116108b657829003601f168201915b505050505081526020016001820180546108ec90613e14565b80601f016020809104026020016040519081016040528092919081815260200182805461091890613e14565b80156109635780601f1061093a57610100808354040283529160200191610963565b820191905f5260205f20905b81548152906001019060200180831161094657829003601f168201915b5050505050815260200160028201805461097c90613e14565b80601f01602080910402602001604051908101604052809291908181526020018280546109a890613e14565b80156109f35780601f106109ca576101008083540402835291602001916109f3565b820191905f5260205f20905b8154815290600101906020018083116109d657829003601f168201915b505050505081526020016003820154815250509050949350505050565b610a18612e6d565b610a20612344565b50610a2a8261066e565b610a69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a609061422a565b60405180910390fd5b5f600583604051610a7a9190613ef6565b90815260200160405180910390206002015414610acc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac390614292565b60405180910390fd5b42600583604051610add9190613ef6565b908152602001604051809103902060020181905550610afb826124f5565b600582604051610b0b9190613ef6565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115610b4857610b476135c8565b5b600a811115610b5a57610b596135c8565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b610bef612344565b50610bfb8160026119d3565b610c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3190613e9e565b60405180910390fd5b42600582604051610c4b9190613ef6565b908152602001604051809103902060070181905550610c6b816004612471565b50565b610c76612e46565b600682604051610c869190613ef6565b90815260200160405180910390206040518060800160405290815f82018054610cae90613e14565b80601f0160208091040260200160405190810160405280929190818152602001828054610cda90613e14565b8015610d255780601f10610cfc57610100808354040283529160200191610d25565b820191905f5260205f20905b815481529060010190602001808311610d0857829003601f168201915b50505050508152602001600182018054610d3e90613e14565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a90613e14565b8015610db55780601f10610d8c57610100808354040283529160200191610db5565b820191905f5260205f20905b815481529060010190602001808311610d9857829003601f168201915b50505050508152602001600282018054610dce90613e14565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfa90613e14565b8015610e455780601f10610e1c57610100808354040283529160200191610e45565b820191905f5260205f20905b815481529060010190602001808311610e2857829003601f168201915b505050505081526020016003820154815250509050919050565b610e67612344565b50610e738160016119d3565b610eb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea990613f56565b60405180910390fd5b42600582604051610ec39190613ef6565b9081526020016040518091039020600b0181905550610ee3816008612471565b50565b610eee6125ac565b610ef75f61262a565b565b5f600682604051610f0a9190613ef6565b908152602001604051809103902060030154600383604051610f2c9190613ef6565b9081526020016040518091039020600d015f0154610f4a91906142dd565b9050919050565b610f59612344565b50610f64815f6119d3565b610fa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9a90614368565b60405180910390fd5b42600582604051610fb49190613ef6565b9081526020016040518091039020600a0181905550610fd4816007612471565b50565b610fdf612344565b50610feb8160046119d3565b61102a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611021906143d0565b60405180910390fd5b4260058260405161103b9190613ef6565b90815260200160405180910390206009018190555061105b816006612471565b50565b5f600160169054906101000a900460ff16159050808015611090575060018060159054906101000a900460ff1660ff16105b806110be575061109f306126eb565b1580156110bd575060018060159054906101000a900460ff1660ff16145b5b6110fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f49061445e565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156111395760018060166101000a81548160ff0219169083151502179055505b8015611192575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161118991906144c1565b60405180910390a15b50565b5f61119f82610ef9565b6003836040516111af9190613ef6565b9081526020016040518091039020600d01600201546111ce91906142dd565b9050919050565b6111dd612344565b506111e98160046119d3565b611228576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121f906143d0565b60405180910390fd5b426005826040516112399190613ef6565b908152602001604051809103902060080181905550611259816005612471565b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61128b6125ac565b5f600160146101000a81548160ff021916908315150217905550565b6112af612ee0565b6112b7612344565b506112c2855f6119d3565b611301576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112f890614368565b60405180910390fd5b61130a8461270d565b61131383612922565b61131c8561066e565b1561135d575f61132b86611b87565b905061133b8160200151866129b8565b611349816040015185612c35565b611357816060015184612c7f565b5061168a565b60405180608001604052808681526020018581526020018481526020018381525060038660405161138e9190613ef6565b90815260200160405180910390205f820151815f0190816113af9190614111565b506020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190805190602001906114a5929190612f1a565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060045f856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f9091909190915090816115c99190614111565b5060045f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150908161163a9190614111565b507f542b0c7ab03fea46499267e26e574794a6ebce8e10cca32c61c289eacca5576985855f015186602001518760400151886060015160405161168195949392919061457e565b60405180910390a15b60038560405161169a9190613ef6565b90815260200160405180910390206040518060800160405290815f820180546116c290613e14565b80601f01602080910402602001604051908101604052809291908181526020018280546116ee90613e14565b80156117395780601f1061171057610100808354040283529160200191611739565b820191905f5260205f20905b81548152906001019060200180831161171c57829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054806020026020016040519081016040528092919081815260200182805480156118d257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611889575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f6119dd8361066e565b611a0f575f600a8111156119f4576119f36135c8565b5b82600a811115611a0757611a066135c8565b5b149050611a65565b81600a811115611a2257611a216135c8565b5b600584604051611a329190613ef6565b90815260200160405180910390205f015f9054906101000a900460ff16600a811115611a6157611a606135c8565b5b1490505b92915050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611a9b612e6d565b600582604051611aab9190613ef6565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115611ae857611ae76135c8565b5b600a811115611afa57611af96135c8565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b611b8f612ee0565b600382604051611b9f9190613ef6565b90815260200160405180910390206040518060800160405290815f82018054611bc790613e14565b80601f0160208091040260200160405190810160405280929190818152602001828054611bf390613e14565b8015611c3e5780601f10611c1557610100808354040283529160200191611c3e565b820191905f5260205f20905b815481529060010190602001808311611c2157829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611dd757602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d8e575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611edd612344565b50611ee98160046119d3565b611f28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f1f906143d0565b60405180910390fd5b42600582604051611f399190613ef6565b9081526020016040518091039020600d0181905550611f5981600a612471565b50565b611f64612344565b50611f708160026119d3565b611faf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fa690613e9e565b60405180910390fd5b42600582604051611fc09190613ef6565b9081526020016040518091039020600c0181905550611fe0816009612471565b50565b611feb612e6d565b611ff3612344565b50611ffd8261066e565b61203c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120339061422a565b60405180910390fd5b5f60058360405161204d9190613ef6565b9081526020016040518091039020600101541461209f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161209690614627565b60405180910390fd5b426005836040516120b09190613ef6565b9081526020016040518091039020600101819055506120ce826124f5565b6005826040516120de9190613ef6565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a81111561211b5761211a6135c8565b5b600a81111561212d5761212c6135c8565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6121c26125ac565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612230576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612227906146b5565b60405180910390fd5b6122398161262a565b50565b6122446125ac565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036122b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122a990614743565b60405180910390fd5b600160149054906101000a900460ff16612301576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122f8906147d1565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036123d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123cb90614743565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612414612da9565b73ffffffffffffffffffffffffffffffffffffffff161461246a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124619061485f565b60405180910390fd5b6001905090565b806005836040516124829190613ef6565b90815260200160405180910390205f015f6101000a81548160ff0219169083600a8111156124b3576124b26135c8565b5b02179055507f10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac38682826040516124e992919061488c565b60405180910390a15050565b5f6005826040516125069190613ef6565b9081526020016040518091039020600101541415801561254757505f6005826040516125329190613ef6565b90815260200160405180910390206002015414155b15612582574260058260405161255d9190613ef6565b90815260200160405180910390206004018190555061257d816001612471565b6125a9565b426005826040516125939190613ef6565b9081526020016040518091039020600301819055505b50565b6125b4612da9565b73ffffffffffffffffffffffffffffffffffffffff166125d261125c565b73ffffffffffffffffffffffffffffffffffffffff1614612628576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161261f90614904565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff160361277f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127769061496c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16036127f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127e8906149d4565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612862576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161285990614a3c565b60405180910390fd5b5f816060015151116128a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128a090614aa4565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff160361291f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161291690614b0c565b60405180910390fd5b50565b5f815f0151602001511461296b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161296290614b74565b60405180910390fd5b5f816060015160200151146129b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ac90614bdc565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff1614612a2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a2590614c44565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612aa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a9b90614cac565b60405180910390fd5b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff1614612b18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0f90614d14565b60405180910390fd5b80606001515182606001515114612b64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b5b90614d7c565b60405180910390fd5b5f5b826060015151811015612c305781606001518181518110612b8a57612b89614d9a565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612bbf57612bbe614d9a565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612c1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1490614e11565b60405180910390fd5b8080612c2890614e2f565b915050612b66565b505050565b612c45825f0151825f0151612db0565b612c5782602001518260200151612db0565b612c6982604001518260400151612db0565b612c7b82606001518260600151612db0565b5050565b805f0151825f015114612cc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cbe90614ec0565b60405180910390fd5b8060200151826020015114612d11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d0890614f28565b60405180910390fd5b8060400151826040015114612d5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d5290614f90565b60405180910390fd5b8060600151826060015114612da5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d9c90614ff8565b60405180910390fd5b5050565b5f33905090565b805f0151825f015114612df8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612def90615060565b60405180910390fd5b8060200151826020015114612e42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e39906150c8565b60405180910390fd5b5050565b60405180608001604052806060815260200160608152602001606081526020015f81525090565b604051806101c001604052805f600a811115612e8c57612e8b6135c8565b5b81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b604051806080016040528060608152602001612efa612fa1565b8152602001612f07613008565b8152602001612f14613048565b81525090565b828054828255905f5260205f20908101928215612f90579160200282015b82811115612f8f578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612f38565b5b509050612f9d919061306c565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b604051806080016040528061301b613087565b8152602001613028613087565b8152602001613035613087565b8152602001613042613087565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b80821115613083575f815f90555060010161306d565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6130d9826130b0565b9050919050565b6130e9816130cf565b81146130f3575f80fd5b50565b5f81359050613104816130e0565b92915050565b5f6020828403121561311f5761311e6130a8565b5b5f61312c848285016130f6565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561319557808201518184015260208101905061317a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6131ba8261315e565b6131c48185613168565b93506131d4818560208601613178565b6131dd816131a0565b840191505092915050565b5f6131f383836131b0565b905092915050565b5f602082019050919050565b5f61321182613135565b61321b818561313f565b93508360208202850161322d8561314f565b805f5b85811015613268578484038952815161324985826131e8565b9450613254836131fb565b925060208a01995050600181019050613230565b50829750879550505050505092915050565b5f6020820190508181035f8301526132928184613207565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6132d8826131a0565b810181811067ffffffffffffffff821117156132f7576132f66132a2565b5b80604052505050565b5f61330961309f565b905061331582826132cf565b919050565b5f67ffffffffffffffff821115613334576133336132a2565b5b61333d826131a0565b9050602081019050919050565b828183375f83830152505050565b5f61336a6133658461331a565b613300565b9050828152602081018484840111156133865761338561329e565b5b61339184828561334a565b509392505050565b5f82601f8301126133ad576133ac61329a565b5b81356133bd848260208601613358565b91505092915050565b5f602082840312156133db576133da6130a8565b5b5f82013567ffffffffffffffff8111156133f8576133f76130ac565b5b61340484828501613399565b91505092915050565b5f8115159050919050565b6134218161340d565b82525050565b5f60208201905061343a5f830184613418565b92915050565b5f819050919050565b61345281613440565b811461345c575f80fd5b50565b5f8135905061346d81613449565b92915050565b5f805f806080858703121561348b5761348a6130a8565b5b5f85013567ffffffffffffffff8111156134a8576134a76130ac565b5b6134b487828801613399565b945050602085013567ffffffffffffffff8111156134d5576134d46130ac565b5b6134e187828801613399565b935050604085013567ffffffffffffffff811115613502576135016130ac565b5b61350e87828801613399565b925050606061351f8782880161345f565b91505092959194509250565b61353481613440565b82525050565b5f608083015f8301518482035f86015261355482826131b0565b9150506020830151848203602086015261356e82826131b0565b9150506040830151848203604086015261358882826131b0565b915050606083015161359d606086018261352b565b508091505092915050565b5f6020820190508181035f8301526135c0818461353a565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b8110613606576136056135c8565b5b50565b5f819050613616826135f5565b919050565b5f61362582613609565b9050919050565b6136358161361b565b82525050565b6101c082015f8201516136505f85018261362c565b506020820151613663602085018261352b565b506040820151613676604085018261352b565b506060820151613689606085018261352b565b50608082015161369c608085018261352b565b5060a08201516136af60a085018261352b565b5060c08201516136c260c085018261352b565b5060e08201516136d560e085018261352b565b506101008201516136ea61010085018261352b565b506101208201516136ff61012085018261352b565b5061014082015161371461014085018261352b565b5061016082015161372961016085018261352b565b5061018082015161373e61018085018261352b565b506101a08201516137536101a085018261352b565b50505050565b5f6101c08201905061376d5f83018461363b565b92915050565b61377c81613440565b82525050565b5f6020820190506137955f830184613773565b92915050565b6137a4816130cf565b82525050565b5f6020820190506137bd5f83018461379b565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156137e5576137e46132a2565b5b602082029050602081019050919050565b5f80fd5b5f61380c613807846137cb565b613300565b9050808382526020820190506020840283018581111561382f5761382e6137f6565b5b835b81811015613858578061384488826130f6565b845260208401935050602081019050613831565b5050509392505050565b5f82601f8301126138765761387561329a565b5b81356138868482602086016137fa565b91505092915050565b5f608082840312156138a4576138a36137c3565b5b6138ae6080613300565b90505f6138bd848285016130f6565b5f8301525060206138d0848285016130f6565b60208301525060406138e4848285016130f6565b604083015250606082013567ffffffffffffffff811115613908576139076137c7565b5b61391484828501613862565b60608301525092915050565b5f60408284031215613935576139346137c3565b5b61393f6040613300565b90505f61394e8482850161345f565b5f8301525060206139618482850161345f565b60208301525092915050565b5f6101008284031215613983576139826137c3565b5b61398d6080613300565b90505f61399c84828501613920565b5f8301525060406139af84828501613920565b60208301525060806139c384828501613920565b60408301525060c06139d784828501613920565b60608301525092915050565b5f608082840312156139f8576139f76137c3565b5b613a026080613300565b90505f613a118482850161345f565b5f830152506020613a248482850161345f565b6020830152506040613a388482850161345f565b6040830152506060613a4c8482850161345f565b60608301525092915050565b5f805f806101c08587031215613a7157613a706130a8565b5b5f85013567ffffffffffffffff811115613a8e57613a8d6130ac565b5b613a9a87828801613399565b945050602085013567ffffffffffffffff811115613abb57613aba6130ac565b5b613ac78782880161388f565b9350506040613ad88782880161396d565b925050610140613aea878288016139e3565b91505092959194509250565b613aff816130cf565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f613b398383613af6565b60208301905092915050565b5f602082019050919050565b5f613b5b82613b05565b613b658185613b0f565b9350613b7083613b1f565b805f5b83811015613ba0578151613b878882613b2e565b9750613b9283613b45565b925050600181019050613b73565b5085935050505092915050565b5f608083015f830151613bc25f860182613af6565b506020830151613bd56020860182613af6565b506040830151613be86040860182613af6565b5060608301518482036060860152613c008282613b51565b9150508091505092915050565b604082015f820151613c215f85018261352b565b506020820151613c34602085018261352b565b50505050565b61010082015f820151613c4f5f850182613c0d565b506020820151613c626040850182613c0d565b506040820151613c756080850182613c0d565b506060820151613c8860c0850182613c0d565b50505050565b608082015f820151613ca25f85018261352b565b506020820151613cb5602085018261352b565b506040820151613cc8604085018261352b565b506060820151613cdb606085018261352b565b50505050565b5f6101c083015f8301518482035f860152613cfc82826131b0565b91505060208301518482036020860152613d168282613bad565b9150506040830151613d2b6040860182613c3a565b506060830151613d3f610140860182613c8e565b508091505092915050565b5f6020820190508181035f830152613d628184613ce1565b905092915050565b600b8110613d76575f80fd5b50565b5f81359050613d8781613d6a565b92915050565b5f8060408385031215613da357613da26130a8565b5b5f83013567ffffffffffffffff811115613dc057613dbf6130ac565b5b613dcc85828601613399565b9250506020613ddd85828601613d79565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680613e2b57607f821691505b602082108103613e3e57613e3d613de7565b5b50919050565b5f82825260208201905092915050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f613e88601083613e44565b9150613e9382613e54565b602082019050919050565b5f6020820190508181035f830152613eb581613e7c565b9050919050565b5f81905092915050565b5f613ed08261315e565b613eda8185613ebc565b9350613eea818560208601613178565b80840191505092915050565b5f613f018284613ec6565b915081905092915050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f613f40600a83613e44565b9150613f4b82613f0c565b602082019050919050565b5f6020820190508181035f830152613f6d81613f34565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302613fd07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82613f95565b613fda8683613f95565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61401561401061400b84613440565b613ff2565b613440565b9050919050565b5f819050919050565b61402e83613ffb565b61404261403a8261401c565b848454613fa1565b825550505050565b5f90565b61405661404a565b614061818484614025565b505050565b5b81811015614084576140795f8261404e565b600181019050614067565b5050565b601f8211156140c95761409a81613f74565b6140a384613f86565b810160208510156140b2578190505b6140c66140be85613f86565b830182614066565b50505b505050565b5f82821c905092915050565b5f6140e95f19846008026140ce565b1980831691505092915050565b5f61410183836140da565b9150826002028217905092915050565b61411a8261315e565b67ffffffffffffffff811115614133576141326132a2565b5b61413d8254613e14565b614148828285614088565b5f60209050601f831160018114614179575f8415614167578287015190505b61417185826140f6565b8655506141d8565b601f19841661418786613f74565b5f5b828110156141ae57848901518255600182019150602085019450602081019050614189565b868310156141cb57848901516141c7601f8916826140da565b8355505b6001600288020188555050505b505050505050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f614214601383613e44565b915061421f826141e0565b602082019050919050565b5f6020820190508181035f83015261424181614208565b9050919050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f61427c601583613e44565b915061428782614248565b602082019050919050565b5f6020820190508181035f8301526142a981614270565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6142e782613440565b91506142f283613440565b925082820261430081613440565b91508282048414831517614317576143166142b0565b5b5092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f614352600f83613e44565b915061435d8261431e565b602082019050919050565b5f6020820190508181035f83015261437f81614346565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f6143ba600e83613e44565b91506143c582614386565b602082019050919050565b5f6020820190508181035f8301526143e7816143ae565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f614448602e83613e44565b9150614453826143ee565b604082019050919050565b5f6020820190508181035f8301526144758161443c565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f6144ab6144a66144a18461447c565b613ff2565b614485565b9050919050565b6144bb81614491565b82525050565b5f6020820190506144d45f8301846144b2565b92915050565b5f6144e48261315e565b6144ee8185613e44565b93506144fe818560208601613178565b614507816131a0565b840191505092915050565b5f82825260208201905092915050565b5f61452c82613b05565b6145368185614512565b935061454183613b1f565b805f5b838110156145715781516145588882613b2e565b975061456383613b45565b925050600181019050614544565b5085935050505092915050565b5f60a0820190508181035f83015261459681886144da565b90506145a5602083018761379b565b6145b2604083018661379b565b6145bf606083018561379b565b81810360808301526145d18184614522565b90509695505050505050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614611601583613e44565b915061461c826145dd565b602082019050919050565b5f6020820190508181035f83015261463e81614605565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61469f602683613e44565b91506146aa82614645565b604082019050919050565b5f6020820190508181035f8301526146cc81614693565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f61472d603583613e44565b9150614738826146d3565b604082019050919050565b5f6020820190508181035f83015261475a81614721565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f6147bb603983613e44565b91506147c682614761565b604082019050919050565b5f6020820190508181035f8301526147e8816147af565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f614849603b83613e44565b9150614854826147ef565b604082019050919050565b5f6020820190508181035f8301526148768161483d565b9050919050565b6148868161361b565b82525050565b5f6040820190508181035f8301526148a481856144da565b90506148b3602083018461487d565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6148ee602083613e44565b91506148f9826148ba565b602082019050919050565b5f6020820190508181035f83015261491b816148e2565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f614956600a83613e44565b915061496182614922565b602082019050919050565b5f6020820190508181035f8301526149838161494a565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f6149be600a83613e44565b91506149c98261498a565b602082019050919050565b5f6020820190508181035f8301526149eb816149b2565b9050919050565b7f536f6c766572206d697373696e670000000000000000000000000000000000005f82015250565b5f614a26600e83613e44565b9150614a31826149f2565b602082019050919050565b5f6020820190508181035f830152614a5381614a1a565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f614a8e600e83613e44565b9150614a9982614a5a565b602082019050919050565b5f6020820190508181035f830152614abb81614a82565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f614af6600c83613e44565b9150614b0182614ac2565b602082019050919050565b5f6020820190508181035f830152614b2381614aea565b9050919050565b7f4167726565206465706f736974206d75737420626520300000000000000000005f82015250565b5f614b5e601783613e44565b9150614b6982614b2a565b602082019050919050565b5f6020820190508181035f830152614b8b81614b52565b9050919050565b7f4d656469617465206465706f736974206d7573742062652030000000000000005f82015250565b5f614bc6601983613e44565b9150614bd182614b92565b602082019050919050565b5f6020820190508181035f830152614bf381614bba565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f614c2e600283613e44565b9150614c3982614bfa565b602082019050919050565b5f6020820190508181035f830152614c5b81614c22565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f614c96600283613e44565b9150614ca182614c62565b602082019050919050565b5f6020820190508181035f830152614cc381614c8a565b9050919050565b7f536f6c76657200000000000000000000000000000000000000000000000000005f82015250565b5f614cfe600683613e44565b9150614d0982614cca565b602082019050919050565b5f6020820190508181035f830152614d2b81614cf2565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f614d66600983613e44565b9150614d7182614d32565b602082019050919050565b5f6020820190508181035f830152614d9381614d5a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f614dfb600883613e44565b9150614e0682614dc7565b602082019050919050565b5f6020820190508181035f830152614e2881614def565b9050919050565b5f614e3982613440565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614e6b57614e6a6142b0565b5b600182019050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f614eaa600583613e44565b9150614eb582614e76565b602082019050919050565b5f6020820190508181035f830152614ed781614e9e565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f614f12600783613e44565b9150614f1d82614ede565b602082019050919050565b5f6020820190508181035f830152614f3f81614f06565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f614f7a600783613e44565b9150614f8582614f46565b602082019050919050565b5f6020820190508181035f830152614fa781614f6e565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f614fe2600983613e44565b9150614fed82614fae565b602082019050919050565b5f6020820190508181035f83015261500f81614fd6565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f61504a600783613e44565b915061505582615016565b602082019050919050565b5f6020820190508181035f8301526150778161503e565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f6150b2600a83613e44565b91506150bd8261507e565b602082019050919050565b5f6020820190508181035f8301526150df816150a6565b905091905056fea2646970667358221220652fa501d90c3c55f341461f1e2aeec4c070a39c2e58b9dae25432d472439d6864736f6c63430008150033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetAgreement(opts *bind.CallOpts, dealId string) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetAgreement(dealId string) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetAgreement(dealId string) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageSession) GetControllerAddress() (common.Address, error) {
	return _Storage.Contract.GetControllerAddress(&_Storage.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageCallerSession) GetControllerAddress() (common.Address, error) {
	return _Storage.Contract.GetControllerAddress(&_Storage.CallOpts)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCaller) GetDeal(opts *bind.CallOpts, dealId string) (SharedStructsDeal, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCallerSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageSession) GetDealsForParty(party common.Address) ([]string, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageCallerSession) GetDealsForParty(party common.Address) ([]string, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageCaller) GetJobCost(opts *bind.CallOpts, dealId string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageSession) GetJobCost(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetJobCost(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageCaller) GetResult(opts *bind.CallOpts, dealId string) (SharedStructsResult, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageSession) GetResult(dealId string) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageCallerSession) GetResult(dealId string) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageCaller) GetResultsCollateral(opts *bind.CallOpts, dealId string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageSession) GetResultsCollateral(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetResultsCollateral(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageCaller) HasDeal(opts *bind.CallOpts, dealId string) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageSession) HasDeal(dealId string) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageCallerSession) HasDeal(dealId string) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageCaller) IsState(opts *bind.CallOpts, dealId string, state uint8) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageSession) IsState(dealId string, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageCallerSession) IsState(dealId string, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageTransactorSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageTransactor) AddResult(opts *bind.TransactOpts, dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addResult", dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageTransactorSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageTransactor) CheckResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "checkResult", dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageTransactorSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "ensureDeal", dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) EnsureDeal(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactorSession) EnsureDeal(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, members, timeouts, pricing)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageTransactorSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageTransactorSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutAgree(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutAgree", dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// StorageDealIterator is returned from FilterDeal and is used to iterate over the raw logs and unpacked data for Deal events raised by the Storage contract.
type StorageDealIterator struct {
	Event *StorageDeal // Event containing the contract specifics and raw log

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
func (it *StorageDealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDeal)
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
		it.Event = new(StorageDeal)
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
func (it *StorageDealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDeal represents a Deal event raised by the Storage contract.
type StorageDeal struct {
	DealId           string
	Solver           common.Address
	JobCreator       common.Address
	ResourceProvider common.Address
	Mediators        []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeal is a free log retrieval operation binding the contract event 0x542b0c7ab03fea46499267e26e574794a6ebce8e10cca32c61c289eacca55769.
//
// Solidity: event Deal(string dealId, address solver, address jobCreator, address resourceProvider, address[] mediators)
func (_Storage *StorageFilterer) FilterDeal(opts *bind.FilterOpts) (*StorageDealIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Deal")
	if err != nil {
		return nil, err
	}
	return &StorageDealIterator{contract: _Storage.contract, event: "Deal", logs: logs, sub: sub}, nil
}

// WatchDeal is a free log subscription operation binding the contract event 0x542b0c7ab03fea46499267e26e574794a6ebce8e10cca32c61c289eacca55769.
//
// Solidity: event Deal(string dealId, address solver, address jobCreator, address resourceProvider, address[] mediators)
func (_Storage *StorageFilterer) WatchDeal(opts *bind.WatchOpts, sink chan<- *StorageDeal) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Deal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDeal)
				if err := _Storage.contract.UnpackLog(event, "Deal", log); err != nil {
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

// ParseDeal is a log parse operation binding the contract event 0x542b0c7ab03fea46499267e26e574794a6ebce8e10cca32c61c289eacca55769.
//
// Solidity: event Deal(string dealId, address solver, address jobCreator, address resourceProvider, address[] mediators)
func (_Storage *StorageFilterer) ParseDeal(log types.Log) (*StorageDeal, error) {
	event := new(StorageDeal)
	if err := _Storage.contract.UnpackLog(event, "Deal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDealStateChangeIterator is returned from FilterDealStateChange and is used to iterate over the raw logs and unpacked data for DealStateChange events raised by the Storage contract.
type StorageDealStateChangeIterator struct {
	Event *StorageDealStateChange // Event containing the contract specifics and raw log

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
func (it *StorageDealStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDealStateChange)
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
		it.Event = new(StorageDealStateChange)
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
func (it *StorageDealStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDealStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDealStateChange represents a DealStateChange event raised by the Storage contract.
type StorageDealStateChange struct {
	DealId string
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) FilterDealStateChange(opts *bind.FilterOpts) (*StorageDealStateChangeIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DealStateChange")
	if err != nil {
		return nil, err
	}
	return &StorageDealStateChangeIterator{contract: _Storage.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *StorageDealStateChange) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DealStateChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDealStateChange)
				if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
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

// ParseDealStateChange is a log parse operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) ParseDealStateChange(log types.Log) (*StorageDealStateChange, error) {
	event := new(StorageDealStateChange)
	if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
