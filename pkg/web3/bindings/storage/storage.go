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
	Mediator                 common.Address
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
	InstructionCount *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ensureDeal2\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6153f080620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101c2575f3560e01c8063822f80c5116100f7578063cdd82d1d11610095578063e850be371161006f578063e850be3714610512578063ec95b9671461052e578063f2fde38b1461055e578063f3d3d4481461057a576101c2565b8063cdd82d1d14610496578063e7079180146104c6578063e7b957d1146104f6576101c2565b8063a4702958116100d1578063a47029581461040e578063a6370b0e14610418578063b050e74b14610448578063c57380a214610478576101c2565b8063822f80c5146103a4578063824518aa146103d45780638da5cb5b146103f0576101c2565b8063511a9f6811610164578063795f9abf1161013e578063795f9abf1461033257806380ffdfe01461034e5780638129fc1c1461036a5780638224ce5f14610374576101c2565b8063511a9f68146102dc578063715018a6146102f857806373db5c6a14610302576101c2565b8063297f9e55116101a0578063297f9e551461024457806336d9a17a146102605780633c4135da1461027c578063498cc70d146102ac576101c2565b806311d5af33146101c65780631f7b0b27146101f65780632244ad2b14610214575b5f80fd5b6101e060048036038101906101db919061349b565b610596565b6040516101ed919061360b565b60405180910390f35b6101fe6106a7565b60405161020b9190613897565b60405180910390f35b61022e600480360381019061022991906139e3565b6109fa565b60405161023b9190613a44565b60405180910390f35b61025e600480360381019061025991906139e3565b610a11565b005b61027a60048036038101906102759190613a5d565b610a98565b005b610296600480360381019061029191906139e3565b610b7f565b6040516102a39190613c5d565b60405180910390f35b6102c660048036038101906102c191906139e3565b610dab565b6040516102d39190613ccb565b60405180910390f35b6102f660048036038101906102f191906139e3565b610f0c565b005b610300610f93565b005b61031c600480360381019061031791906139e3565b610fa6565b6040516103299190613cfa565b60405180910390f35b61034c600480360381019061034791906139e3565b610ffe565b005b610368600480360381019061036391906139e3565b611084565b005b61037261110b565b005b61038e600480360381019061038991906139e3565b611242565b60405161039b9190613cfa565b60405180910390f35b6103be60048036038101906103b99190613d3d565b611282565b6040516103cb9190613ccb565b60405180910390f35b6103ee60048036038101906103e991906139e3565b6114d9565b005b6103f8611560565b6040516104059190613dd4565b60405180910390f35b610416611587565b005b610432600480360381019061042d9190614082565b6115ab565b60405161043f9190613897565b60405180910390f35b610462600480360381019061045d9190614143565b611c89565b60405161046f9190613a44565b60405180910390f35b610480611d21565b60405161048d9190613dd4565b60405180910390f35b6104b060048036038101906104ab91906139e3565b611d49565b6040516104bd9190613c5d565b60405180910390f35b6104e060048036038101906104db91906139e3565b611e92565b6040516104ed9190613897565b60405180910390f35b610510600480360381019061050b91906139e3565b6121e0565b005b61052c600480360381019061052791906139e3565b612267565b005b610548600480360381019061054391906139e3565b6122ee565b6040516105559190613c5d565b60405180910390f35b6105786004803603810190610573919061349b565b61251a565b005b610594600480360381019061058f919061349b565b61259c565b005b606060045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b8282101561069c578382905f5260205f20018054610611906141ca565b80601f016020809104026020016040519081016040528092919081815260200182805461063d906141ca565b80156106885780601f1061065f57610100808354040283529160200191610688565b820191905f5260205f20905b81548152906001019060200180831161066b57829003601f168201915b5050505050815260200190600101906105f4565b505050509050919050565b6106af6131c2565b6106b76126a4565b5060036040516106c69061424e565b90815260200160405180910390206040518060800160405290815f820180546106ee906141ca565b80601f016020809104026020016040519081016040528092919081815260200182805461071a906141ca565b80156107655780601f1061073c57610100808354040283529160200191610765565b820191905f5260205f20905b81548152906001019060200180831161074857829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054806020026020016040519081016040528092919081815260200182805480156108fe57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108b5575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f8201548152602001600182015481526020016002820154815260200160038201548152505081525050905090565b5f80610a0583611e92565b5f015151119050919050565b610a196126a4565b50610a25816002611c89565b610a64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5b906142bc565b60405180910390fd5b42600582604051610a75919061430a565b908152602001604051809103902060060181905550610a958160036127d1565b50565b610aa06126a4565b50610aac826002611c89565b610aeb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae2906142bc565b60405180910390fd5b42600583604051610afc919061430a565b90815260200160405180910390206007018190555080600583604051610b22919061430a565b90815260200160405180910390205f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b7b8260046127d1565b5050565b610b876131fc565b610b8f6126a4565b50610b99826109fa565b610bd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bcf9061436a565b60405180910390fd5b5f600583604051610be9919061430a565b90815260200160405180910390206002015414610c3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c32906143d2565b60405180910390fd5b42600583604051610c4c919061430a565b908152602001604051809103902060020181905550610c6a82612871565b600582604051610c7a919061430a565b9081526020016040518091039020604051806101e00160405290815f82015f9054906101000a900460ff16600a811115610cb757610cb6613ab7565b5b600a811115610cc957610cc8613ab7565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b610db361328b565b600682604051610dc3919061430a565b90815260200160405180910390206040518060600160405290815f82018054610deb906141ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610e17906141ca565b8015610e625780601f10610e3957610100808354040283529160200191610e62565b820191905f5260205f20905b815481529060010190602001808311610e4557829003601f168201915b50505050508152602001600182018054610e7b906141ca565b80601f0160208091040260200160405190810160405280929190818152602001828054610ea7906141ca565b8015610ef25780601f10610ec957610100808354040283529160200191610ef2565b820191905f5260205f20905b815481529060010190602001808311610ed557829003601f168201915b505050505081526020016002820154815250509050919050565b610f146126a4565b50610f20816001611c89565b610f5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f569061443a565b60405180910390fd5b42600582604051610f70919061430a565b9081526020016040518091039020600b0181905550610f908160086127d1565b50565b610f9b612928565b610fa45f6129a6565b565b5f600682604051610fb7919061430a565b908152602001604051809103902060020154600383604051610fd9919061430a565b9081526020016040518091039020600d015f0154610ff79190614485565b9050919050565b6110066126a4565b50611011815f611c89565b611050576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104790614510565b60405180910390fd5b42600582604051611061919061430a565b9081526020016040518091039020600a01819055506110818160076127d1565b50565b61108c6126a4565b50611098816004611c89565b6110d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ce90614578565b60405180910390fd5b426005826040516110e8919061430a565b9081526020016040518091039020600901819055506111088160066127d1565b50565b5f600160169054906101000a900460ff1615905080801561113d575060018060159054906101000a900460ff1660ff16105b8061116b575061114c30612a67565b15801561116a575060018060159054906101000a900460ff1660ff16145b5b6111aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111a190614606565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156111e65760018060166101000a81548160ff0219169083151502179055505b801561123f575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516112369190614672565b60405180910390a15b50565b5f61124c82610fa6565b60038360405161125c919061430a565b9081526020016040518091039020600d016002015461127b9190614485565b9050919050565b61128a61328b565b6112926126a4565b5061129e846001611c89565b6112dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d49061443a565b60405180910390fd5b426005856040516112ee919061430a565b90815260200160405180910390206005018190555061130e8460026127d1565b604051806060016040528085815260200184815260200183815250600685604051611339919061430a565b90815260200160405180910390205f820151815f01908161135a919061481f565b506020820151816001019081611370919061481f565b506040820151816002015590505060068460405161138e919061430a565b90815260200160405180910390206040518060600160405290815f820180546113b6906141ca565b80601f01602080910402602001604051908101604052809291908181526020018280546113e2906141ca565b801561142d5780601f106114045761010080835404028352916020019161142d565b820191905f5260205f20905b81548152906001019060200180831161141057829003601f168201915b50505050508152602001600182018054611446906141ca565b80601f0160208091040260200160405190810160405280929190818152602001828054611472906141ca565b80156114bd5780601f10611494576101008083540402835291602001916114bd565b820191905f5260205f20905b8154815290600101906020018083116114a057829003601f168201915b5050505050815260200160028201548152505090509392505050565b6114e16126a4565b506114ed816004611c89565b61152c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161152390614578565b60405180910390fd5b4260058260405161153d919061430a565b90815260200160405180910390206008018190555061155d8160056127d1565b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61158f612928565b5f600160146101000a81548160ff021916908315150217905550565b6115b36131c2565b6115bb6126a4565b506115c6855f611c89565b611605576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fc90614510565b60405180910390fd5b61160e84612a89565b61161783612c9e565b611620856109fa565b15611661575f61162f86611e92565b905061163f816020015186612d34565b61164d816040015185612fb1565b61165b816060015184612ffb565b50611940565b604051806080016040528086815260200185815260200184815260200183815250600386604051611692919061430a565b90815260200160405180910390205f820151815f0190816116b3919061481f565b506020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190805190602001906117a99291906132ab565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060045f856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f9091909190915090816118cd919061481f565b5060045f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150908161193e919061481f565b505b600385604051611950919061430a565b90815260200160405180910390206040518060800160405290815f82018054611978906141ca565b80601f01602080910402602001604051908101604052809291908181526020018280546119a4906141ca565b80156119ef5780601f106119c6576101008083540402835291602001916119ef565b820191905f5260205f20905b8154815290600101906020018083116119d257829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611b8857602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b3f575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f611c93836109fa565b611cc5575f600a811115611caa57611ca9613ab7565b5b82600a811115611cbd57611cbc613ab7565b5b149050611d1b565b81600a811115611cd857611cd7613ab7565b5b600584604051611ce8919061430a565b90815260200160405180910390205f015f9054906101000a900460ff16600a811115611d1757611d16613ab7565b5b1490505b92915050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611d516131fc565b600582604051611d61919061430a565b9081526020016040518091039020604051806101e00160405290815f82015f9054906101000a900460ff16600a811115611d9e57611d9d613ab7565b5b600a811115611db057611daf613ab7565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b611e9a6131c2565b600382604051611eaa919061430a565b90815260200160405180910390206040518060800160405290815f82018054611ed2906141ca565b80601f0160208091040260200160405190810160405280929190818152602001828054611efe906141ca565b8015611f495780601f10611f2057610100808354040283529160200191611f49565b820191905f5260205f20905b815481529060010190602001808311611f2c57829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054806020026020016040519081016040528092919081815260200182805480156120e257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612099575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b6121e86126a4565b506121f4816004611c89565b612233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222a90614578565b60405180910390fd5b42600582604051612244919061430a565b9081526020016040518091039020600d018190555061226481600a6127d1565b50565b61226f6126a4565b5061227b816002611c89565b6122ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b1906142bc565b60405180910390fd5b426005826040516122cb919061430a565b9081526020016040518091039020600c01819055506122eb8160096127d1565b50565b6122f66131fc565b6122fe6126a4565b50612308826109fa565b612347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233e9061436a565b60405180910390fd5b5f600583604051612358919061430a565b908152602001604051809103902060010154146123aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123a190614938565b60405180910390fd5b426005836040516123bb919061430a565b9081526020016040518091039020600101819055506123d982612871565b6005826040516123e9919061430a565b9081526020016040518091039020604051806101e00160405290815f82015f9054906101000a900460ff16600a81111561242657612425613ab7565b5b600a81111561243857612437613ab7565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b612522612928565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612590576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612587906149c6565b60405180910390fd5b612599816129a6565b50565b6125a4612928565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612612576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161260990614a54565b60405180910390fd5b600160149054906101000a900460ff16612661576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161265890614ae2565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161272b90614a54565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612774613125565b73ffffffffffffffffffffffffffffffffffffffff16146127ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127c190614b70565b60405180910390fd5b6001905090565b806005836040516127e2919061430a565b90815260200160405180910390205f015f6101000a81548160ff0219169083600a81111561281357612812613ab7565b5b021790555080600a81111561282b5761282a613ab7565b5b82604051612839919061430a565b60405180910390207f10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac38660405160405180910390a35050565b5f600582604051612882919061430a565b908152602001604051809103902060010154141580156128c357505f6005826040516128ae919061430a565b90815260200160405180910390206002015414155b156128fe57426005826040516128d9919061430a565b9081526020016040518091039020600401819055506128f98160016127d1565b612925565b4260058260405161290f919061430a565b9081526020016040518091039020600301819055505b50565b612930613125565b73ffffffffffffffffffffffffffffffffffffffff1661294e611560565b73ffffffffffffffffffffffffffffffffffffffff16146129a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161299b90614bd8565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612afb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612af290614c40565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612b6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b6490614ca8565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612bde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bd590614d10565b60405180910390fd5b5f81606001515111612c25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1c90614d78565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612c9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9290614de0565b60405180910390fd5b50565b5f815f01516020015114612ce7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cde90614e48565b60405180910390fd5b5f81606001516020015114612d31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d2890614eb0565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff1614612daa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da190614f18565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612e20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e1790614f80565b60405180910390fd5b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff1614612e94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e8b90614fe8565b60405180910390fd5b80606001515182606001515114612ee0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ed790615050565b60405180910390fd5b5f5b826060015151811015612fac5781606001518181518110612f0657612f0561506e565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612f3b57612f3a61506e565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612f99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f90906150e5565b60405180910390fd5b8080612fa490615103565b915050612ee2565b505050565b612fc1825f0151825f015161312c565b612fd38260200151826020015161312c565b612fe58260400151826040015161312c565b612ff78260600151826060015161312c565b5050565b805f0151825f015114613043576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161303a90615194565b60405180910390fd5b806020015182602001511461308d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613084906151fc565b60405180910390fd5b80604001518260400151146130d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130ce90615264565b60405180910390fd5b8060600151826060015114613121576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613118906152cc565b60405180910390fd5b5050565b5f33905090565b805f0151825f015114613174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161316b90615334565b60405180910390fd5b80602001518260200151146131be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131b59061539c565b60405180910390fd5b5050565b6040518060800160405280606081526020016131dc613332565b81526020016131e9613399565b81526020016131f66133d9565b81525090565b604051806101e001604052805f600a81111561321b5761321a613ab7565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b604051806060016040528060608152602001606081526020015f81525090565b828054828255905f5260205f20908101928215613321579160200282015b82811115613320578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906132c9565b5b50905061332e91906133fd565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b60405180608001604052806133ac613418565b81526020016133b9613418565b81526020016133c6613418565b81526020016133d3613418565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b80821115613414575f815f9055506001016133fe565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61346a82613441565b9050919050565b61347a81613460565b8114613484575f80fd5b50565b5f8135905061349581613471565b92915050565b5f602082840312156134b0576134af613439565b5b5f6134bd84828501613487565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561352657808201518184015260208101905061350b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61354b826134ef565b61355581856134f9565b9350613565818560208601613509565b61356e81613531565b840191505092915050565b5f6135848383613541565b905092915050565b5f602082019050919050565b5f6135a2826134c6565b6135ac81856134d0565b9350836020820285016135be856134e0565b805f5b858110156135f957848403895281516135da8582613579565b94506135e58361358c565b925060208a019950506001810190506135c1565b50829750879550505050505092915050565b5f6020820190508181035f8301526136238184613598565b905092915050565b61363481613460565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61366e838361362b565b60208301905092915050565b5f602082019050919050565b5f6136908261363a565b61369a8185613644565b93506136a583613654565b805f5b838110156136d55781516136bc8882613663565b97506136c78361367a565b9250506001810190506136a8565b5085935050505092915050565b5f608083015f8301516136f75f86018261362b565b50602083015161370a602086018261362b565b50604083015161371d604086018261362b565b50606083015184820360608601526137358282613686565b9150508091505092915050565b5f819050919050565b61375481613742565b82525050565b604082015f82015161376e5f85018261374b565b506020820151613781602085018261374b565b50505050565b61010082015f82015161379c5f85018261375a565b5060208201516137af604085018261375a565b5060408201516137c2608085018261375a565b5060608201516137d560c085018261375a565b50505050565b608082015f8201516137ef5f85018261374b565b506020820151613802602085018261374b565b506040820151613815604085018261374b565b506060820151613828606085018261374b565b50505050565b5f6101c083015f8301518482035f8601526138498282613541565b9150506020830151848203602086015261386382826136e2565b91505060408301516138786040860182613787565b50606083015161388c6101408601826137db565b508091505092915050565b5f6020820190508181035f8301526138af818461382e565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6138f582613531565b810181811067ffffffffffffffff82111715613914576139136138bf565b5b80604052505050565b5f613926613430565b905061393282826138ec565b919050565b5f67ffffffffffffffff821115613951576139506138bf565b5b61395a82613531565b9050602081019050919050565b828183375f83830152505050565b5f61398761398284613937565b61391d565b9050828152602081018484840111156139a3576139a26138bb565b5b6139ae848285613967565b509392505050565b5f82601f8301126139ca576139c96138b7565b5b81356139da848260208601613975565b91505092915050565b5f602082840312156139f8576139f7613439565b5b5f82013567ffffffffffffffff811115613a1557613a1461343d565b5b613a21848285016139b6565b91505092915050565b5f8115159050919050565b613a3e81613a2a565b82525050565b5f602082019050613a575f830184613a35565b92915050565b5f8060408385031215613a7357613a72613439565b5b5f83013567ffffffffffffffff811115613a9057613a8f61343d565b5b613a9c858286016139b6565b9250506020613aad85828601613487565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b8110613af557613af4613ab7565b5b50565b5f819050613b0582613ae4565b919050565b5f613b1482613af8565b9050919050565b613b2481613b0a565b82525050565b6101e082015f820151613b3f5f850182613b1b565b506020820151613b52602085018261362b565b506040820151613b65604085018261374b565b506060820151613b78606085018261374b565b506080820151613b8b608085018261374b565b5060a0820151613b9e60a085018261374b565b5060c0820151613bb160c085018261374b565b5060e0820151613bc460e085018261374b565b50610100820151613bd961010085018261374b565b50610120820151613bee61012085018261374b565b50610140820151613c0361014085018261374b565b50610160820151613c1861016085018261374b565b50610180820151613c2d61018085018261374b565b506101a0820151613c426101a085018261374b565b506101c0820151613c576101c085018261374b565b50505050565b5f6101e082019050613c715f830184613b2a565b92915050565b5f606083015f8301518482035f860152613c918282613541565b91505060208301518482036020860152613cab8282613541565b9150506040830151613cc0604086018261374b565b508091505092915050565b5f6020820190508181035f830152613ce38184613c77565b905092915050565b613cf481613742565b82525050565b5f602082019050613d0d5f830184613ceb565b92915050565b613d1c81613742565b8114613d26575f80fd5b50565b5f81359050613d3781613d13565b92915050565b5f805f60608486031215613d5457613d53613439565b5b5f84013567ffffffffffffffff811115613d7157613d7061343d565b5b613d7d868287016139b6565b935050602084013567ffffffffffffffff811115613d9e57613d9d61343d565b5b613daa868287016139b6565b9250506040613dbb86828701613d29565b9150509250925092565b613dce81613460565b82525050565b5f602082019050613de75f830184613dc5565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115613e0f57613e0e6138bf565b5b602082029050602081019050919050565b5f80fd5b5f613e36613e3184613df5565b61391d565b90508083825260208201905060208402830185811115613e5957613e58613e20565b5b835b81811015613e825780613e6e8882613487565b845260208401935050602081019050613e5b565b5050509392505050565b5f82601f830112613ea057613e9f6138b7565b5b8135613eb0848260208601613e24565b91505092915050565b5f60808284031215613ece57613ecd613ded565b5b613ed8608061391d565b90505f613ee784828501613487565b5f830152506020613efa84828501613487565b6020830152506040613f0e84828501613487565b604083015250606082013567ffffffffffffffff811115613f3257613f31613df1565b5b613f3e84828501613e8c565b60608301525092915050565b5f60408284031215613f5f57613f5e613ded565b5b613f69604061391d565b90505f613f7884828501613d29565b5f830152506020613f8b84828501613d29565b60208301525092915050565b5f6101008284031215613fad57613fac613ded565b5b613fb7608061391d565b90505f613fc684828501613f4a565b5f830152506040613fd984828501613f4a565b6020830152506080613fed84828501613f4a565b60408301525060c061400184828501613f4a565b60608301525092915050565b5f6080828403121561402257614021613ded565b5b61402c608061391d565b90505f61403b84828501613d29565b5f83015250602061404e84828501613d29565b602083015250604061406284828501613d29565b604083015250606061407684828501613d29565b60608301525092915050565b5f805f806101c0858703121561409b5761409a613439565b5b5f85013567ffffffffffffffff8111156140b8576140b761343d565b5b6140c4878288016139b6565b945050602085013567ffffffffffffffff8111156140e5576140e461343d565b5b6140f187828801613eb9565b935050604061410287828801613f97565b9250506101406141148782880161400d565b91505092959194509250565b600b811061412c575f80fd5b50565b5f8135905061413d81614120565b92915050565b5f806040838503121561415957614158613439565b5b5f83013567ffffffffffffffff8111156141765761417561343d565b5b614182858286016139b6565b92505060206141938582860161412f565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806141e157607f821691505b6020821081036141f4576141f361419d565b5b50919050565b5f81905092915050565b7f74657374000000000000000000000000000000000000000000000000000000005f82015250565b5f6142386004836141fa565b915061424382614204565b600482019050919050565b5f6142588261422c565b9150819050919050565b5f82825260208201905092915050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f6142a6601083614262565b91506142b182614272565b602082019050919050565b5f6020820190508181035f8301526142d38161429a565b9050919050565b5f6142e4826134ef565b6142ee81856141fa565b93506142fe818560208601613509565b80840191505092915050565b5f61431582846142da565b915081905092915050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f614354601383614262565b915061435f82614320565b602082019050919050565b5f6020820190508181035f83015261438181614348565b9050919050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f6143bc601583614262565b91506143c782614388565b602082019050919050565b5f6020820190508181035f8301526143e9816143b0565b9050919050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f614424600a83614262565b915061442f826143f0565b602082019050919050565b5f6020820190508181035f83015261445181614418565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61448f82613742565b915061449a83613742565b92508282026144a881613742565b915082820484148315176144bf576144be614458565b5b5092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f6144fa600f83614262565b9150614505826144c6565b602082019050919050565b5f6020820190508181035f830152614527816144ee565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f614562600e83614262565b915061456d8261452e565b602082019050919050565b5f6020820190508181035f83015261458f81614556565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f6145f0602e83614262565b91506145fb82614596565b604082019050919050565b5f6020820190508181035f83015261461d816145e4565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61465c61465761465284614624565b614639565b61462d565b9050919050565b61466c81614642565b82525050565b5f6020820190506146855f830184614663565b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026146e77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826146ac565b6146f186836146ac565b95508019841693508086168417925050509392505050565b5f61472361471e61471984613742565b614639565b613742565b9050919050565b5f819050919050565b61473c83614709565b6147506147488261472a565b8484546146b8565b825550505050565b5f90565b614764614758565b61476f818484614733565b505050565b5b81811015614792576147875f8261475c565b600181019050614775565b5050565b601f8211156147d7576147a88161468b565b6147b18461469d565b810160208510156147c0578190505b6147d46147cc8561469d565b830182614774565b50505b505050565b5f82821c905092915050565b5f6147f75f19846008026147dc565b1980831691505092915050565b5f61480f83836147e8565b9150826002028217905092915050565b614828826134ef565b67ffffffffffffffff811115614841576148406138bf565b5b61484b82546141ca565b614856828285614796565b5f60209050601f831160018114614887575f8415614875578287015190505b61487f8582614804565b8655506148e6565b601f1984166148958661468b565b5f5b828110156148bc57848901518255600182019150602085019450602081019050614897565b868310156148d957848901516148d5601f8916826147e8565b8355505b6001600288020188555050505b505050505050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614922601583614262565b915061492d826148ee565b602082019050919050565b5f6020820190508181035f83015261494f81614916565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6149b0602683614262565b91506149bb82614956565b604082019050919050565b5f6020820190508181035f8301526149dd816149a4565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f614a3e603583614262565b9150614a49826149e4565b604082019050919050565b5f6020820190508181035f830152614a6b81614a32565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f614acc603983614262565b9150614ad782614a72565b604082019050919050565b5f6020820190508181035f830152614af981614ac0565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f614b5a603b83614262565b9150614b6582614b00565b604082019050919050565b5f6020820190508181035f830152614b8781614b4e565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614bc2602083614262565b9150614bcd82614b8e565b602082019050919050565b5f6020820190508181035f830152614bef81614bb6565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f614c2a600a83614262565b9150614c3582614bf6565b602082019050919050565b5f6020820190508181035f830152614c5781614c1e565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f614c92600a83614262565b9150614c9d82614c5e565b602082019050919050565b5f6020820190508181035f830152614cbf81614c86565b9050919050565b7f536f6c766572206d697373696e670000000000000000000000000000000000005f82015250565b5f614cfa600e83614262565b9150614d0582614cc6565b602082019050919050565b5f6020820190508181035f830152614d2781614cee565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f614d62600e83614262565b9150614d6d82614d2e565b602082019050919050565b5f6020820190508181035f830152614d8f81614d56565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f614dca600c83614262565b9150614dd582614d96565b602082019050919050565b5f6020820190508181035f830152614df781614dbe565b9050919050565b7f4167726565206465706f736974203000000000000000000000000000000000005f82015250565b5f614e32600f83614262565b9150614e3d82614dfe565b602082019050919050565b5f6020820190508181035f830152614e5f81614e26565b9050919050565b7f4d656469617465206465706f73697420300000000000000000000000000000005f82015250565b5f614e9a601183614262565b9150614ea582614e66565b602082019050919050565b5f6020820190508181035f830152614ec781614e8e565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f614f02600283614262565b9150614f0d82614ece565b602082019050919050565b5f6020820190508181035f830152614f2f81614ef6565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f614f6a600283614262565b9150614f7582614f36565b602082019050919050565b5f6020820190508181035f830152614f9781614f5e565b9050919050565b7f536f6c76657200000000000000000000000000000000000000000000000000005f82015250565b5f614fd2600683614262565b9150614fdd82614f9e565b602082019050919050565b5f6020820190508181035f830152614fff81614fc6565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f61503a600983614262565b915061504582615006565b602082019050919050565b5f6020820190508181035f8301526150678161502e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f6150cf600883614262565b91506150da8261509b565b602082019050919050565b5f6020820190508181035f8301526150fc816150c3565b9050919050565b5f61510d82613742565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361513f5761513e614458565b5b600182019050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f61517e600583614262565b91506151898261514a565b602082019050919050565b5f6020820190508181035f8301526151ab81615172565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f6151e6600783614262565b91506151f1826151b2565b602082019050919050565b5f6020820190508181035f830152615213816151da565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f61524e600783614262565b91506152598261521a565b602082019050919050565b5f6020820190508181035f83015261527b81615242565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f6152b6600983614262565b91506152c182615282565b602082019050919050565b5f6020820190508181035f8301526152e3816152aa565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f61531e600783614262565b9150615329826152ea565b602082019050919050565b5f6020820190508181035f83015261534b81615312565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f615386600a83614262565b915061539182615352565b602082019050919050565b5f6020820190508181035f8301526153b38161537a565b905091905056fea2646970667358221220012c410fde28cd3e42dac7a2b79f9e06cca000a0317a30623b83157140b920a964736f6c63430008150033",
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

// EnsureDeal2 is a free data retrieval call binding the contract method 0x1f7b0b27.
//
// Solidity: function ensureDeal2() view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCaller) EnsureDeal2(opts *bind.CallOpts) (SharedStructsDeal, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "ensureDeal2")

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// EnsureDeal2 is a free data retrieval call binding the contract method 0x1f7b0b27.
//
// Solidity: function ensureDeal2() view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) EnsureDeal2() (SharedStructsDeal, error) {
	return _Storage.Contract.EnsureDeal2(&_Storage.CallOpts)
}

// EnsureDeal2 is a free data retrieval call binding the contract method 0x1f7b0b27.
//
// Solidity: function ensureDeal2() view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCallerSession) EnsureDeal2() (SharedStructsDeal, error) {
	return _Storage.Contract.EnsureDeal2(&_Storage.CallOpts)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function getAgreement(string dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetAgreement(dealId string) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function getResult(string dealId) view returns((string,string,uint256))
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
// Solidity: function getResult(string dealId) view returns((string,string,uint256))
func (_Storage *StorageSession) GetResult(dealId string) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,uint256))
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

// AddResult is a paid mutator transaction binding the contract method 0x822f80c5.
//
// Solidity: function addResult(string dealId, string resultsId, uint256 instructionCount) returns((string,string,uint256))
func (_Storage *StorageTransactor) AddResult(opts *bind.TransactOpts, dealId string, resultsId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x822f80c5.
//
// Solidity: function addResult(string dealId, string resultsId, uint256 instructionCount) returns((string,string,uint256))
func (_Storage *StorageSession) AddResult(dealId string, resultsId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x822f80c5.
//
// Solidity: function addResult(string dealId, string resultsId, uint256 instructionCount) returns((string,string,uint256))
func (_Storage *StorageTransactorSession) AddResult(dealId string, resultsId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x36d9a17a.
//
// Solidity: function checkResult(string dealId, address mediator) returns()
func (_Storage *StorageTransactor) CheckResult(opts *bind.TransactOpts, dealId string, mediator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x36d9a17a.
//
// Solidity: function checkResult(string dealId, address mediator) returns()
func (_Storage *StorageSession) CheckResult(dealId string, mediator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x36d9a17a.
//
// Solidity: function checkResult(string dealId, address mediator) returns()
func (_Storage *StorageTransactorSession) CheckResult(dealId string, mediator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId, mediator)
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
	DealId common.Hash
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) FilterDealStateChange(opts *bind.FilterOpts, dealId []string, state []uint8) (*StorageDealStateChangeIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return &StorageDealStateChangeIterator{contract: _Storage.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *StorageDealStateChange, dealId []string, state []uint8) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DealStateChange", dealIdRule, stateRule)
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
// Solidity: event DealStateChange(string indexed dealId, uint8 indexed state)
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
