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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61500380620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101a7575f3560e01c80638224ce5f116100f7578063cdd82d1d11610095578063e850be371161006f578063e850be37146104d9578063ec95b967146104f5578063f2fde38b14610525578063f3d3d44814610541576101a7565b8063cdd82d1d1461045d578063e70791801461048d578063e7b957d1146104bd576101a7565b8063a4702958116100d1578063a4702958146103d5578063a6370b0e146103df578063b050e74b1461040f578063c57380a21461043f576101a7565b80638224ce5f1461036b578063824518aa1461039b5780638da5cb5b146103b7576101a7565b8063498cc70d1161016457806373db5c6a1161013e57806373db5c6a146102f9578063795f9abf1461032957806380ffdfe0146103455780638129fc1c14610361576101a7565b8063498cc70d146102a3578063511a9f68146102d3578063715018a6146102ef576101a7565b806311d5af33146101ab5780632244ad2b146101db578063297f9e551461020b5780633955548e146102275780633c4135da1461025757806346834d1e14610287575b5f80fd5b6101c560048036038101906101c091906130bc565b61055d565b6040516101d2919061322c565b60405180910390f35b6101f560048036038101906101f09190613378565b61066e565b60405161020291906133d9565b60405180910390f35b61022560048036038101906102209190613378565b610685565b005b610241600480360381019061023c9190613425565b61070c565b60405161024e919061355a565b60405180910390f35b610271600480360381019061026c9190613378565b610a10565b60405161027e919061370b565b60405180910390f35b6102a1600480360381019061029c9190613378565b610be7565b005b6102bd60048036038101906102b89190613378565b610c6e565b6040516102ca919061355a565b60405180910390f35b6102ed60048036038101906102e89190613378565b610e5f565b005b6102f7610ee6565b005b610313600480360381019061030e9190613378565b610ef9565b6040516103209190613734565b60405180910390f35b610343600480360381019061033e9190613378565b610f51565b005b61035f600480360381019061035a9190613378565b610fd7565b005b61036961105e565b005b61038560048036038101906103809190613378565b611195565b6040516103929190613734565b60405180910390f35b6103b560048036038101906103b09190613378565b6111d5565b005b6103bf61125c565b6040516103cc919061375c565b60405180910390f35b6103dd611283565b005b6103f960048036038101906103f49190613a0a565b6112a7565b6040516104069190613cfc565b60405180910390f35b61042960048036038101906104249190613d3f565b611985565b60405161043691906133d9565b60405180910390f35b610447611a1d565b604051610454919061375c565b60405180910390f35b61047760048036038101906104729190613378565b611a45565b604051610484919061370b565b60405180910390f35b6104a760048036038101906104a29190613378565b611b39565b6040516104b49190613cfc565b60405180910390f35b6104d760048036038101906104d29190613378565b611e87565b005b6104f360048036038101906104ee9190613378565b611f0e565b005b61050f600480360381019061050a9190613378565b611f95565b60405161051c919061370b565b60405180910390f35b61053f600480360381019061053a91906130bc565b61216c565b005b61055b600480360381019061055691906130bc565b6121ee565b005b606060045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610663578382905f5260205f200180546105d890613dc6565b80601f016020809104026020016040519081016040528092919081815260200182805461060490613dc6565b801561064f5780601f106106265761010080835404028352916020019161064f565b820191905f5260205f20905b81548152906001019060200180831161063257829003601f168201915b5050505050815260200190600101906105bb565b505050509050919050565b5f8061067983611b39565b5f015151119050919050565b61068d6122f6565b50610699816002611985565b6106d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cf90613e50565b60405180910390fd5b426005826040516106e99190613ea8565b908152602001604051809103902060060181905550610709816003612423565b50565b610714612df8565b61071c6122f6565b50610728856001611985565b610767576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075e90613f08565b60405180910390fd5b426005866040516107789190613ea8565b908152602001604051809103902060050181905550610798856002612423565b6040518060800160405280868152602001858152602001848152602001838152506006866040516107c99190613ea8565b90815260200160405180910390205f820151815f0190816107ea91906140c3565b50602082015181600101908161080091906140c3565b50604082015181600201908161081691906140c3565b50606082015181600301559050506006856040516108349190613ea8565b90815260200160405180910390206040518060800160405290815f8201805461085c90613dc6565b80601f016020809104026020016040519081016040528092919081815260200182805461088890613dc6565b80156108d35780601f106108aa576101008083540402835291602001916108d3565b820191905f5260205f20905b8154815290600101906020018083116108b657829003601f168201915b505050505081526020016001820180546108ec90613dc6565b80601f016020809104026020016040519081016040528092919081815260200182805461091890613dc6565b80156109635780601f1061093a57610100808354040283529160200191610963565b820191905f5260205f20905b81548152906001019060200180831161094657829003601f168201915b5050505050815260200160028201805461097c90613dc6565b80601f01602080910402602001604051908101604052809291908181526020018280546109a890613dc6565b80156109f35780601f106109ca576101008083540402835291602001916109f3565b820191905f5260205f20905b8154815290600101906020018083116109d657829003601f168201915b505050505081526020016003820154815250509050949350505050565b610a18612e1f565b610a206122f6565b50610a2a8261066e565b610a69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a60906141dc565b60405180910390fd5b5f600583604051610a7a9190613ea8565b90815260200160405180910390206002015414610acc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac390614244565b60405180910390fd5b42600583604051610add9190613ea8565b908152602001604051809103902060020181905550610afb826124a7565b600582604051610b0b9190613ea8565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115610b4857610b4761357a565b5b600a811115610b5a57610b5961357a565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b610bef6122f6565b50610bfb816002611985565b610c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3190613e50565b60405180910390fd5b42600582604051610c4b9190613ea8565b908152602001604051809103902060070181905550610c6b816004612423565b50565b610c76612df8565b600682604051610c869190613ea8565b90815260200160405180910390206040518060800160405290815f82018054610cae90613dc6565b80601f0160208091040260200160405190810160405280929190818152602001828054610cda90613dc6565b8015610d255780601f10610cfc57610100808354040283529160200191610d25565b820191905f5260205f20905b815481529060010190602001808311610d0857829003601f168201915b50505050508152602001600182018054610d3e90613dc6565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6a90613dc6565b8015610db55780601f10610d8c57610100808354040283529160200191610db5565b820191905f5260205f20905b815481529060010190602001808311610d9857829003601f168201915b50505050508152602001600282018054610dce90613dc6565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfa90613dc6565b8015610e455780601f10610e1c57610100808354040283529160200191610e45565b820191905f5260205f20905b815481529060010190602001808311610e2857829003601f168201915b505050505081526020016003820154815250509050919050565b610e676122f6565b50610e73816001611985565b610eb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea990613f08565b60405180910390fd5b42600582604051610ec39190613ea8565b9081526020016040518091039020600b0181905550610ee3816008612423565b50565b610eee61255e565b610ef75f6125dc565b565b5f600682604051610f0a9190613ea8565b908152602001604051809103902060030154600383604051610f2c9190613ea8565b9081526020016040518091039020600d015f0154610f4a919061428f565b9050919050565b610f596122f6565b50610f64815f611985565b610fa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9a9061431a565b60405180910390fd5b42600582604051610fb49190613ea8565b9081526020016040518091039020600a0181905550610fd4816007612423565b50565b610fdf6122f6565b50610feb816004611985565b61102a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102190614382565b60405180910390fd5b4260058260405161103b9190613ea8565b90815260200160405180910390206009018190555061105b816006612423565b50565b5f600160169054906101000a900460ff16159050808015611090575060018060159054906101000a900460ff1660ff16105b806110be575061109f3061269d565b1580156110bd575060018060159054906101000a900460ff1660ff16145b5b6110fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f490614410565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156111395760018060166101000a81548160ff0219169083151502179055505b8015611192575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516111899190614473565b60405180910390a15b50565b5f61119f82610ef9565b6003836040516111af9190613ea8565b9081526020016040518091039020600d01600201546111ce919061428f565b9050919050565b6111dd6122f6565b506111e9816004611985565b611228576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121f90614382565b60405180910390fd5b426005826040516112399190613ea8565b908152602001604051809103902060080181905550611259816005612423565b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61128b61255e565b5f600160146101000a81548160ff021916908315150217905550565b6112af612e92565b6112b76122f6565b506112c2855f611985565b611301576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112f89061431a565b60405180910390fd5b61130a846126bf565b611313836128d4565b61131c8561066e565b1561135d575f61132b86611b39565b905061133b81602001518661296a565b611349816040015185612be7565b611357816060015184612c31565b5061163c565b60405180608001604052808681526020018581526020018481526020018381525060038660405161138e9190613ea8565b90815260200160405180910390205f820151815f0190816113af91906140c3565b506020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190805190602001906114a5929190612ecc565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060045f856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f9091909190915090816115c991906140c3565b5060045f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150908161163a91906140c3565b505b60038560405161164c9190613ea8565b90815260200160405180910390206040518060800160405290815f8201805461167490613dc6565b80601f01602080910402602001604051908101604052809291908181526020018280546116a090613dc6565b80156116eb5780601f106116c2576101008083540402835291602001916116eb565b820191905f5260205f20905b8154815290600101906020018083116116ce57829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820180548060200260200160405190810160405280929190818152602001828054801561188457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161183b575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f61198f8361066e565b6119c1575f600a8111156119a6576119a561357a565b5b82600a8111156119b9576119b861357a565b5b149050611a17565b81600a8111156119d4576119d361357a565b5b6005846040516119e49190613ea8565b90815260200160405180910390205f015f9054906101000a900460ff16600a811115611a1357611a1261357a565b5b1490505b92915050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611a4d612e1f565b600582604051611a5d9190613ea8565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115611a9a57611a9961357a565b5b600a811115611aac57611aab61357a565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b611b41612e92565b600382604051611b519190613ea8565b90815260200160405180910390206040518060800160405290815f82018054611b7990613dc6565b80601f0160208091040260200160405190810160405280929190818152602001828054611ba590613dc6565b8015611bf05780601f10611bc757610100808354040283529160200191611bf0565b820191905f5260205f20905b815481529060010190602001808311611bd357829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611d8957602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d40575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611e8f6122f6565b50611e9b816004611985565b611eda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ed190614382565b60405180910390fd5b42600582604051611eeb9190613ea8565b9081526020016040518091039020600d0181905550611f0b81600a612423565b50565b611f166122f6565b50611f22816002611985565b611f61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5890613e50565b60405180910390fd5b42600582604051611f729190613ea8565b9081526020016040518091039020600c0181905550611f92816009612423565b50565b611f9d612e1f565b611fa56122f6565b50611faf8261066e565b611fee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fe5906141dc565b60405180910390fd5b5f600583604051611fff9190613ea8565b90815260200160405180910390206001015414612051576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612048906144d6565b60405180910390fd5b426005836040516120629190613ea8565b908152602001604051809103902060010181905550612080826124a7565b6005826040516120909190613ea8565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a8111156120cd576120cc61357a565b5b600a8111156120df576120de61357a565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b61217461255e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036121e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121d990614564565b60405180910390fd5b6121eb816125dc565b50565b6121f661255e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161225b906145f2565b60405180910390fd5b600160149054906101000a900460ff166122b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122aa90614680565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612386576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161237d906145f2565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166123c6612d5b565b73ffffffffffffffffffffffffffffffffffffffff161461241c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124139061470e565b60405180910390fd5b6001905090565b806005836040516124349190613ea8565b90815260200160405180910390205f015f6101000a81548160ff0219169083600a8111156124655761246461357a565b5b02179055507f10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386828260405161249b929190614773565b60405180910390a15050565b5f6005826040516124b89190613ea8565b908152602001604051809103902060010154141580156124f957505f6005826040516124e49190613ea8565b90815260200160405180910390206002015414155b15612534574260058260405161250f9190613ea8565b90815260200160405180910390206004018190555061252f816001612423565b61255b565b426005826040516125459190613ea8565b9081526020016040518091039020600301819055505b50565b612566612d5b565b73ffffffffffffffffffffffffffffffffffffffff1661258461125c565b73ffffffffffffffffffffffffffffffffffffffff16146125da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125d1906147eb565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612731576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161272890614853565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16036127a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161279a906148bb565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612814576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161280b90614923565b60405180910390fd5b5f8160600151511161285b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128529061498b565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16036128d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128c8906149f3565b60405180910390fd5b50565b5f815f0151602001511461291d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161291490614a5b565b60405180910390fd5b5f81606001516020015114612967576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161295e90614ac3565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff16146129e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129d790614b2b565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a4d90614b93565b60405180910390fd5b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff1614612aca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ac190614bfb565b60405180910390fd5b80606001515182606001515114612b16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0d90614c63565b60405180910390fd5b5f5b826060015151811015612be25781606001518181518110612b3c57612b3b614c81565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612b7157612b70614c81565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612bcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bc690614cf8565b60405180910390fd5b8080612bda90614d16565b915050612b18565b505050565b612bf7825f0151825f0151612d62565b612c0982602001518260200151612d62565b612c1b82604001518260400151612d62565b612c2d82606001518260600151612d62565b5050565b805f0151825f015114612c79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c7090614da7565b60405180910390fd5b8060200151826020015114612cc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cba90614e0f565b60405180910390fd5b8060400151826040015114612d0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d0490614e77565b60405180910390fd5b8060600151826060015114612d57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d4e90614edf565b60405180910390fd5b5050565b5f33905090565b805f0151825f015114612daa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da190614f47565b60405180910390fd5b8060200151826020015114612df4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612deb90614faf565b60405180910390fd5b5050565b60405180608001604052806060815260200160608152602001606081526020015f81525090565b604051806101c001604052805f600a811115612e3e57612e3d61357a565b5b81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b604051806080016040528060608152602001612eac612f53565b8152602001612eb9612fba565b8152602001612ec6612ffa565b81525090565b828054828255905f5260205f20908101928215612f42579160200282015b82811115612f41578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612eea565b5b509050612f4f919061301e565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6040518060800160405280612fcd613039565b8152602001612fda613039565b8152602001612fe7613039565b8152602001612ff4613039565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b80821115613035575f815f90555060010161301f565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61308b82613062565b9050919050565b61309b81613081565b81146130a5575f80fd5b50565b5f813590506130b681613092565b92915050565b5f602082840312156130d1576130d061305a565b5b5f6130de848285016130a8565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561314757808201518184015260208101905061312c565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61316c82613110565b613176818561311a565b935061318681856020860161312a565b61318f81613152565b840191505092915050565b5f6131a58383613162565b905092915050565b5f602082019050919050565b5f6131c3826130e7565b6131cd81856130f1565b9350836020820285016131df85613101565b805f5b8581101561321a57848403895281516131fb858261319a565b9450613206836131ad565b925060208a019950506001810190506131e2565b50829750879550505050505092915050565b5f6020820190508181035f83015261324481846131b9565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61328a82613152565b810181811067ffffffffffffffff821117156132a9576132a8613254565b5b80604052505050565b5f6132bb613051565b90506132c78282613281565b919050565b5f67ffffffffffffffff8211156132e6576132e5613254565b5b6132ef82613152565b9050602081019050919050565b828183375f83830152505050565b5f61331c613317846132cc565b6132b2565b90508281526020810184848401111561333857613337613250565b5b6133438482856132fc565b509392505050565b5f82601f83011261335f5761335e61324c565b5b813561336f84826020860161330a565b91505092915050565b5f6020828403121561338d5761338c61305a565b5b5f82013567ffffffffffffffff8111156133aa576133a961305e565b5b6133b68482850161334b565b91505092915050565b5f8115159050919050565b6133d3816133bf565b82525050565b5f6020820190506133ec5f8301846133ca565b92915050565b5f819050919050565b613404816133f2565b811461340e575f80fd5b50565b5f8135905061341f816133fb565b92915050565b5f805f806080858703121561343d5761343c61305a565b5b5f85013567ffffffffffffffff81111561345a5761345961305e565b5b6134668782880161334b565b945050602085013567ffffffffffffffff8111156134875761348661305e565b5b6134938782880161334b565b935050604085013567ffffffffffffffff8111156134b4576134b361305e565b5b6134c08782880161334b565b92505060606134d187828801613411565b91505092959194509250565b6134e6816133f2565b82525050565b5f608083015f8301518482035f8601526135068282613162565b915050602083015184820360208601526135208282613162565b9150506040830151848203604086015261353a8282613162565b915050606083015161354f60608601826134dd565b508091505092915050565b5f6020820190508181035f83015261357281846134ec565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b81106135b8576135b761357a565b5b50565b5f8190506135c8826135a7565b919050565b5f6135d7826135bb565b9050919050565b6135e7816135cd565b82525050565b6101c082015f8201516136025f8501826135de565b50602082015161361560208501826134dd565b50604082015161362860408501826134dd565b50606082015161363b60608501826134dd565b50608082015161364e60808501826134dd565b5060a082015161366160a08501826134dd565b5060c082015161367460c08501826134dd565b5060e082015161368760e08501826134dd565b5061010082015161369c6101008501826134dd565b506101208201516136b16101208501826134dd565b506101408201516136c66101408501826134dd565b506101608201516136db6101608501826134dd565b506101808201516136f06101808501826134dd565b506101a08201516137056101a08501826134dd565b50505050565b5f6101c08201905061371f5f8301846135ed565b92915050565b61372e816133f2565b82525050565b5f6020820190506137475f830184613725565b92915050565b61375681613081565b82525050565b5f60208201905061376f5f83018461374d565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561379757613796613254565b5b602082029050602081019050919050565b5f80fd5b5f6137be6137b98461377d565b6132b2565b905080838252602082019050602084028301858111156137e1576137e06137a8565b5b835b8181101561380a57806137f688826130a8565b8452602084019350506020810190506137e3565b5050509392505050565b5f82601f8301126138285761382761324c565b5b81356138388482602086016137ac565b91505092915050565b5f6080828403121561385657613855613775565b5b61386060806132b2565b90505f61386f848285016130a8565b5f830152506020613882848285016130a8565b6020830152506040613896848285016130a8565b604083015250606082013567ffffffffffffffff8111156138ba576138b9613779565b5b6138c684828501613814565b60608301525092915050565b5f604082840312156138e7576138e6613775565b5b6138f160406132b2565b90505f61390084828501613411565b5f83015250602061391384828501613411565b60208301525092915050565b5f610100828403121561393557613934613775565b5b61393f60806132b2565b90505f61394e848285016138d2565b5f830152506040613961848285016138d2565b6020830152506080613975848285016138d2565b60408301525060c0613989848285016138d2565b60608301525092915050565b5f608082840312156139aa576139a9613775565b5b6139b460806132b2565b90505f6139c384828501613411565b5f8301525060206139d684828501613411565b60208301525060406139ea84828501613411565b60408301525060606139fe84828501613411565b60608301525092915050565b5f805f806101c08587031215613a2357613a2261305a565b5b5f85013567ffffffffffffffff811115613a4057613a3f61305e565b5b613a4c8782880161334b565b945050602085013567ffffffffffffffff811115613a6d57613a6c61305e565b5b613a7987828801613841565b9350506040613a8a8782880161391f565b925050610140613a9c87828801613995565b91505092959194509250565b613ab181613081565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f613aeb8383613aa8565b60208301905092915050565b5f602082019050919050565b5f613b0d82613ab7565b613b178185613ac1565b9350613b2283613ad1565b805f5b83811015613b52578151613b398882613ae0565b9750613b4483613af7565b925050600181019050613b25565b5085935050505092915050565b5f608083015f830151613b745f860182613aa8565b506020830151613b876020860182613aa8565b506040830151613b9a6040860182613aa8565b5060608301518482036060860152613bb28282613b03565b9150508091505092915050565b604082015f820151613bd35f8501826134dd565b506020820151613be660208501826134dd565b50505050565b61010082015f820151613c015f850182613bbf565b506020820151613c146040850182613bbf565b506040820151613c276080850182613bbf565b506060820151613c3a60c0850182613bbf565b50505050565b608082015f820151613c545f8501826134dd565b506020820151613c6760208501826134dd565b506040820151613c7a60408501826134dd565b506060820151613c8d60608501826134dd565b50505050565b5f6101c083015f8301518482035f860152613cae8282613162565b91505060208301518482036020860152613cc88282613b5f565b9150506040830151613cdd6040860182613bec565b506060830151613cf1610140860182613c40565b508091505092915050565b5f6020820190508181035f830152613d148184613c93565b905092915050565b600b8110613d28575f80fd5b50565b5f81359050613d3981613d1c565b92915050565b5f8060408385031215613d5557613d5461305a565b5b5f83013567ffffffffffffffff811115613d7257613d7161305e565b5b613d7e8582860161334b565b9250506020613d8f85828601613d2b565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680613ddd57607f821691505b602082108103613df057613def613d99565b5b50919050565b5f82825260208201905092915050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f613e3a601083613df6565b9150613e4582613e06565b602082019050919050565b5f6020820190508181035f830152613e6781613e2e565b9050919050565b5f81905092915050565b5f613e8282613110565b613e8c8185613e6e565b9350613e9c81856020860161312a565b80840191505092915050565b5f613eb38284613e78565b915081905092915050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f613ef2600a83613df6565b9150613efd82613ebe565b602082019050919050565b5f6020820190508181035f830152613f1f81613ee6565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302613f827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82613f47565b613f8c8683613f47565b95508019841693508086168417925050509392505050565b5f819050919050565b5f613fc7613fc2613fbd846133f2565b613fa4565b6133f2565b9050919050565b5f819050919050565b613fe083613fad565b613ff4613fec82613fce565b848454613f53565b825550505050565b5f90565b614008613ffc565b614013818484613fd7565b505050565b5b818110156140365761402b5f82614000565b600181019050614019565b5050565b601f82111561407b5761404c81613f26565b61405584613f38565b81016020851015614064578190505b61407861407085613f38565b830182614018565b50505b505050565b5f82821c905092915050565b5f61409b5f1984600802614080565b1980831691505092915050565b5f6140b3838361408c565b9150826002028217905092915050565b6140cc82613110565b67ffffffffffffffff8111156140e5576140e4613254565b5b6140ef8254613dc6565b6140fa82828561403a565b5f60209050601f83116001811461412b575f8415614119578287015190505b61412385826140a8565b86555061418a565b601f19841661413986613f26565b5f5b828110156141605784890151825560018201915060208501945060208101905061413b565b8683101561417d5784890151614179601f89168261408c565b8355505b6001600288020188555050505b505050505050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f6141c6601383613df6565b91506141d182614192565b602082019050919050565b5f6020820190508181035f8301526141f3816141ba565b9050919050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f61422e601583613df6565b9150614239826141fa565b602082019050919050565b5f6020820190508181035f83015261425b81614222565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614299826133f2565b91506142a4836133f2565b92508282026142b2816133f2565b915082820484148315176142c9576142c8614262565b5b5092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f614304600f83613df6565b915061430f826142d0565b602082019050919050565b5f6020820190508181035f830152614331816142f8565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f61436c600e83613df6565b915061437782614338565b602082019050919050565b5f6020820190508181035f83015261439981614360565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6143fa602e83613df6565b9150614405826143a0565b604082019050919050565b5f6020820190508181035f830152614427816143ee565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f61445d6144586144538461442e565b613fa4565b614437565b9050919050565b61446d81614443565b82525050565b5f6020820190506144865f830184614464565b92915050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f6144c0601583613df6565b91506144cb8261448c565b602082019050919050565b5f6020820190508181035f8301526144ed816144b4565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61454e602683613df6565b9150614559826144f4565b604082019050919050565b5f6020820190508181035f83015261457b81614542565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6145dc603583613df6565b91506145e782614582565b604082019050919050565b5f6020820190508181035f830152614609816145d0565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f61466a603983613df6565b915061467582614610565b604082019050919050565b5f6020820190508181035f8301526146978161465e565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6146f8603b83613df6565b91506147038261469e565b604082019050919050565b5f6020820190508181035f830152614725816146ec565b9050919050565b5f61473682613110565b6147408185613df6565b935061475081856020860161312a565b61475981613152565b840191505092915050565b61476d816135cd565b82525050565b5f6040820190508181035f83015261478b818561472c565b905061479a6020830184614764565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6147d5602083613df6565b91506147e0826147a1565b602082019050919050565b5f6020820190508181035f830152614802816147c9565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f61483d600a83613df6565b915061484882614809565b602082019050919050565b5f6020820190508181035f83015261486a81614831565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f6148a5600a83613df6565b91506148b082614871565b602082019050919050565b5f6020820190508181035f8301526148d281614899565b9050919050565b7f536f6c766572206d697373696e670000000000000000000000000000000000005f82015250565b5f61490d600e83613df6565b9150614918826148d9565b602082019050919050565b5f6020820190508181035f83015261493a81614901565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f614975600e83613df6565b915061498082614941565b602082019050919050565b5f6020820190508181035f8301526149a281614969565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f6149dd600c83613df6565b91506149e8826149a9565b602082019050919050565b5f6020820190508181035f830152614a0a816149d1565b9050919050565b7f4167726565206465706f736974206d75737420626520300000000000000000005f82015250565b5f614a45601783613df6565b9150614a5082614a11565b602082019050919050565b5f6020820190508181035f830152614a7281614a39565b9050919050565b7f4d656469617465206465706f736974206d7573742062652030000000000000005f82015250565b5f614aad601983613df6565b9150614ab882614a79565b602082019050919050565b5f6020820190508181035f830152614ada81614aa1565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f614b15600283613df6565b9150614b2082614ae1565b602082019050919050565b5f6020820190508181035f830152614b4281614b09565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f614b7d600283613df6565b9150614b8882614b49565b602082019050919050565b5f6020820190508181035f830152614baa81614b71565b9050919050565b7f536f6c76657200000000000000000000000000000000000000000000000000005f82015250565b5f614be5600683613df6565b9150614bf082614bb1565b602082019050919050565b5f6020820190508181035f830152614c1281614bd9565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f614c4d600983613df6565b9150614c5882614c19565b602082019050919050565b5f6020820190508181035f830152614c7a81614c41565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f614ce2600883613df6565b9150614ced82614cae565b602082019050919050565b5f6020820190508181035f830152614d0f81614cd6565b9050919050565b5f614d20826133f2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614d5257614d51614262565b5b600182019050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f614d91600583613df6565b9150614d9c82614d5d565b602082019050919050565b5f6020820190508181035f830152614dbe81614d85565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f614df9600783613df6565b9150614e0482614dc5565b602082019050919050565b5f6020820190508181035f830152614e2681614ded565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f614e61600783613df6565b9150614e6c82614e2d565b602082019050919050565b5f6020820190508181035f830152614e8e81614e55565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f614ec9600983613df6565b9150614ed482614e95565b602082019050919050565b5f6020820190508181035f830152614ef681614ebd565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f614f31600783613df6565b9150614f3c82614efd565b602082019050919050565b5f6020820190508181035f830152614f5e81614f25565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f614f99600a83613df6565b9150614fa482614f65565b602082019050919050565b5f6020820190508181035f830152614fc681614f8d565b905091905056fea2646970667358221220b85216dfd5e2f27ee71c0e2a49a6b80d6ee0e4d9d689bca5626ba67ab837a77964736f6c63430008150033",
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
