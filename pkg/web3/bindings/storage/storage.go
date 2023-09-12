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
	DealId   *big.Int
	Members  SharedStructsDealMembers
	Timeouts SharedStructsDealTimeouts
	Pricing  SharedStructsDealPricing
}

// SharedStructsDealMembers is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealMembers struct {
	Directory        common.Address
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
	DealId           *big.Int
	ResultsId        *big.Int
	InstructionCount *big.Int
}

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress        common.Address
	MetadataCID        *big.Int
	Url                string
	Roles              []uint8
	TrustedMediators   []common.Address
	TrustedDirectories []common.Address
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b615df380620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101ee575f3560e01c80637fb9c45f1161010d578063a7f96f06116100a0578063d74553011161006f578063d7455301146105d4578063ef816fd914610604578063f2fde38b14610634578063f3d3d44814610650576101ee565b8063a7f96f061461054e578063aeb5ec011461056a578063c1390a8914610586578063c57380a2146105b6576101ee565b80638da5cb5b116100dc5780638da5cb5b146104c6578063995e4339146104e4578063a15dcc8a14610514578063a470295814610544576101ee565b80637fb9c45f146104405780638129fc1c1461045c57806382fd5bac146104665780638462d54e14610496576101ee565b8063407c2d6211610185578063546cfd3411610154578063546cfd34146103ba5780635dd049fd146103ea5780636f77926b14610406578063715018a614610436576101ee565b8063407c2d621461030e57806342e4074e1461033e5780634f9f6fe61461036e578063542785671461039e576101ee565b806323a9a862116101c157806323a9a8621461028a57806332ebea04146102a657806335a7e268146102c25780633a83b3e4146102de576101ee565b80630396e3c1146101f257806311d5af331461022257806311d6adc414610252578063172257c71461026e575b5f80fd5b61020c6004803603810190610207919061397f565b61066c565b60405161021991906139b9565b60405180910390f35b61023c60048036038101906102379190613a2c565b6106aa565b6040516102499190613b0e565b60405180910390f35b61026c6004803603810190610267919061397f565b61073d565b005b6102886004803603810190610283919061397f565b6107b6565b005b6102a4600480360381019061029f919061397f565b610830565b005b6102c060048036038101906102bb919061397f565b6108aa565b005b6102dc60048036038101906102d7919061397f565b610924565b005b6102f860048036038101906102f39190613e4f565b61099e565b604051610305919061411e565b60405180910390f35b6103286004803603810190610323919061397f565b610fb8565b6040516103359190614158565b60405180910390f35b6103586004803603810190610353919061397f565b610fcf565b6040516103659190614317565b60405180910390f35b6103886004803603810190610383919061397f565b6111d4565b6040516103959190614317565b60405180910390f35b6103b860048036038101906103b3919061397f565b611310565b005b6103d460048036038101906103cf9190614331565b61138a565b6040516103e191906143c1565b60405180910390f35b61040460048036038101906103ff919061397f565b611496565b005b610420600480360381019061041b9190613a2c565b611510565b60405161042d91906145dd565b60405180910390f35b61043e6117f4565b005b61045a600480360381019061045591906145fd565b611807565b005b6104646118d4565b005b610480600480360381019061047b919061397f565b611a0b565b60405161048d919061411e565b60405180910390f35b6104b060048036038101906104ab919061397f565b611cc6565b6040516104bd9190614317565b60405180910390f35b6104ce611ecb565b6040516104db919061464a565b60405180910390f35b6104fe60048036038101906104f9919061397f565b611ef2565b60405161050b91906143c1565b60405180910390f35b61052e60048036038101906105299190614686565b611f3a565b60405161053b919061471d565b60405180910390f35b61054c611ffa565b005b61056860048036038101906105639190614686565b61201e565b005b610584600480360381019061057f9190614686565b612322565b005b6105a0600480360381019061059b91906148ad565b61251d565b6040516105ad91906145dd565b60405180910390f35b6105be612674565b6040516105cb919061464a565b60405180910390f35b6105ee60048036038101906105e9919061397f565b61269c565b6040516105fb91906139b9565b60405180910390f35b61061e600480360381019061061991906149b7565b6126cf565b60405161062b9190614158565b60405180910390f35b61064e60048036038101906106499190613a2c565b61275a565b005b61066a60048036038101906106659190613a2c565b6127dc565b005b5f60085f8381526020019081526020015f206002015460055f8481526020019081526020015f20600d015f01546106a39190614a22565b9050919050565b606060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561073157602002820191905f5260205f20905b81548152602001906001019080831161071d575b50505050509050919050565b6107456128e4565b50610750815f6126cf565b61078f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078690614abd565b60405180910390fd5b4260075f8381526020019081526020015f20600a01819055506107b3816007612a11565b50565b6107be6128e4565b506107ca8160026126cf565b610809576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080090614b25565b60405180910390fd5b4260075f8381526020019081526020015f206006018190555061082d816003612a11565b50565b6108386128e4565b506108448160046126cf565b610883576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087a90614b8d565b60405180910390fd5b4260075f8381526020019081526020015f20600801819055506108a7816005612a11565b50565b6108b26128e4565b506108be8160016126cf565b6108fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f490614bf5565b60405180910390fd5b4260075f8381526020019081526020015f20600b0181905550610921816008612a11565b50565b61092c6128e4565b506109388160046126cf565b610977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096e90614b8d565b60405180910390fd5b4260075f8381526020019081526020015f20600d018190555061099b81600a612a11565b50565b6109a66135d5565b6109ae6128e4565b506109b9855f6126cf565b6109f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ef90614abd565b60405180910390fd5b610a0184612a8f565b610a0a83612ca4565b610a1385610fb8565b15610a54575f610a2286611a0b565b9050610a32816020015186612d3a565b610a40816040015185612fb7565b610a4e816060015184613001565b50610d02565b60405180608001604052808681526020018581526020018481526020018381525060055f8781526020019081526020015f205f820151815f01556020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610b8392919061360e565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060065f856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f909190919091505560065f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150555b60055f8681526020019081526020015f206040518060800160405290815f8201548152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015610eb757602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e6e575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f80610fc383611a0b565b5f015114159050919050565b610fd7613695565b610fdf6128e4565b50610fe982610fb8565b611028576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101f90614c5d565b60405180910390fd5b5f60075f8481526020019081526020015f20600101541461107e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107590614cc5565b60405180910390fd5b4260075f8481526020019081526020015f20600101819055506110a08261312b565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a8111156110e0576110df614171565b5b600a8111156110f2576110f1614171565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6111dc613695565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a81111561121c5761121b614171565b5b600a81111561122e5761122d614171565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6113186128e4565b506113248160026126cf565b611363576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135a90614b25565b60405180910390fd5b4260075f8381526020019081526020015f20600c0181905550611387816009612a11565b50565b611392613724565b61139a6128e4565b506113a68460016126cf565b6113e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113dc90614bf5565b60405180910390fd5b4260075f8681526020019081526020015f2060050181905550611409846002612a11565b60405180606001604052808581526020018481526020018381525060085f8681526020019081526020015f205f820151815f0155602082015181600101556040820151816002015590505060085f8581526020019081526020015f206040518060600160405290815f82015481526020016001820154815260200160028201548152505090509392505050565b61149e6128e4565b506114aa8160046126cf565b6114e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114e090614b8d565b60405180910390fd5b4260075f8381526020019081526020015f206009018190555061150d816006612a11565b50565b611518613742565b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060c00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820180546115ce90614d10565b80601f01602080910402602001604051908101604052809291908181526020018280546115fa90614d10565b80156116455780601f1061161c57610100808354040283529160200191611645565b820191905f5260205f20905b81548152906001019060200180831161162857829003601f168201915b50505050508152602001600382018054806020026020016040519081016040528092919081815260200182805480156116ce57602002820191905f5260205f20905f905b82829054906101000a900460ff1660048111156116a9576116a8614171565b5b815260200190600101906020825f010492830192600103820291508084116116895790505b505050505081526020016004820180548060200260200160405190810160405280929190818152602001828054801561175957602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611710575b50505050508152602001600582018054806020026020016040519081016040528092919081815260200182805480156117e457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161179b575b5050505050815250509050919050565b6117fc6131ae565b6118055f61322c565b565b61180f6128e4565b5061181b8260026126cf565b61185a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161185190614b25565b60405180910390fd5b4260075f8481526020019081526020015f20600701819055508060075f8481526020019081526020015f205f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118d0826004612a11565b5050565b5f600160169054906101000a900460ff16159050808015611906575060018060159054906101000a900460ff1660ff16105b806119345750611915306132ed565b158015611933575060018060159054906101000a900460ff1660ff16145b5b611973576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196a90614db0565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156119af5760018060166101000a81548160ff0219169083151502179055505b8015611a08575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516119ff9190614e1c565b60405180910390a15b50565b611a136135d5565b60055f8381526020019081526020015f206040518060800160405290815f8201548152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611bc857602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b7f575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611cce613695565b611cd66128e4565b50611ce082610fb8565b611d1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1690614c5d565b60405180910390fd5b5f60075f8481526020019081526020015f206002015414611d75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d6c90614e7f565b60405180910390fd5b4260075f8481526020019081526020015f2060020181905550611d978261312b565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a811115611dd757611dd6614171565b5b600a811115611de957611de8614171565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611efa613724565b60085f8381526020019081526020015f206040518060600160405290815f8201548152602001600182015481526020016002820154815250509050919050565b606060045f836004811115611f5257611f51614171565b5b6004811115611f6457611f63614171565b5b81526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015611fee57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611fa5575b50505050509050919050565b6120026131ae565b5f600160146101000a81548160ff021916908315150217905550565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036120ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120e190614ee7565b60405180910390fd5b5f806120f6833261330f565b915091508061213a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161213190614f4f565b60405180910390fd5b5f8290505b600160045f86600481111561215757612156614171565b5b600481111561216957612168614171565b5b81526020019081526020015f20805490506121849190614f6d565b8110156122a55760045f8560048111156121a1576121a0614171565b5b60048111156121b3576121b2614171565b5b81526020019081526020015f206001826121cd9190614fa0565b815481106121de576121dd614fd3565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660045f86600481111561221c5761221b614171565b5b600481111561222e5761222d614171565b5b81526020019081526020015f20828154811061224d5761224c614fd3565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061229d90615000565b91505061213f565b5060045f8460048111156122bc576122bb614171565b5b60048111156122ce576122cd614171565b5b81526020019081526020015f208054806122eb576122ea615047565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036123ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123e590614ee7565b60405180910390fd5b5f6123f9823261330f565b915050801561243d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612434906150e4565b60405180910390fd5b6124478232613425565b612486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161247d9061514c565b60405180910390fd5b60045f83600481111561249c5761249b614171565b5b60048111156124ae576124ad614171565b5b81526020019081526020015f2032908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b612525613742565b5f6040518060c001604052803273ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018681526020018581526020018481525090508060035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201908161260c91906152fe565b50606082015181600301908051906020019061262992919061378c565b50608082015181600401908051906020019061264692919061360e565b5060a082015181600501908051906020019061266392919061360e565b509050508091505095945050505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f6126a68261066c565b60055f8481526020019081526020015f20600d01600201546126c89190614a22565b9050919050565b5f6126d983610fb8565b61270b575f600a8111156126f0576126ef614171565b5b82600a81111561270357612702614171565b5b149050612754565b81600a81111561271e5761271d614171565b5b60075f8581526020019081526020015f205f015f9054906101000a900460ff16600a8111156127505761274f614171565b5b1490505b92915050565b6127626131ae565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036127d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127c79061543d565b60405180910390fd5b6127d98161322c565b50565b6127e46131ae565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612852576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612849906154cb565b60405180910390fd5b600160149054906101000a900460ff166128a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161289890615559565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161296b906154cb565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166129b4613538565b73ffffffffffffffffffffffffffffffffffffffff1614612a0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a01906155e7565b60405180910390fd5b6001905090565b8060075f8481526020019081526020015f205f015f6101000a81548160ff0219169083600a811115612a4657612a45614171565b5b021790555080600a811115612a5e57612a5d614171565b5b827f17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d60405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612af89061564f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612b73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b6a906156b7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bdb9061571f565b60405180910390fd5b5f81606001515111612c2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c2290615787565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612ca1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c98906157ef565b60405180910390fd5b50565b5f815f01516020015114612ced576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ce490615857565b60405180910390fd5b5f81606001516020015114612d37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d2e906158bf565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff1614612db0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612da790615927565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612e26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e1d9061598f565b60405180910390fd5b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff1614612e9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e91906159f7565b60405180910390fd5b80606001515182606001515114612ee6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612edd90615a5f565b60405180910390fd5b5f5b826060015151811015612fb25781606001518181518110612f0c57612f0b614fd3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612f4157612f40614fd3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612f9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f9690615ac7565b60405180910390fd5b8080612faa90615000565b915050612ee8565b505050565b612fc7825f0151825f015161353f565b612fd98260200151826020015161353f565b612feb8260400151826040015161353f565b612ffd8260600151826060015161353f565b5050565b805f0151825f015114613049576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161304090615b2f565b60405180910390fd5b8060200151826020015114613093576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161308a90615b97565b60405180910390fd5b80604001518260400151146130dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130d490615bff565b60405180910390fd5b8060600151826060015114613127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161311e90615c67565b60405180910390fd5b5050565b5f60075f8381526020019081526020015f20600101541415801561316357505f60075f8381526020019081526020015f206002015414155b15613191574260075f8381526020019081526020015f206004018190555061318c816001612a11565b6131ab565b4260075f8381526020019081526020015f20600301819055505b50565b6131b6613538565b73ffffffffffffffffffffffffffffffffffffffff166131d4611ecb565b73ffffffffffffffffffffffffffffffffffffffff161461322a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161322190615ccf565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f805f805f5b60045f88600481111561332b5761332a614171565b5b600481111561333d5761333c614171565b5b81526020019081526020015f2080549050811015613415578573ffffffffffffffffffffffffffffffffffffffff1660045f89600481111561338257613381614171565b5b600481111561339457613393614171565b5b81526020019081526020015f2082815481106133b3576133b2614fd3565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036134025780925060019150613415565b808061340d90615000565b915050613315565b5081819350935050509250929050565b5f805f90505f5b60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018054905081101561352d5784600481111561348a57613489614171565b5b60035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030182815481106134dc576134db614fd3565b5b905f5260205f2090602091828204019190069054906101000a900460ff16600481111561350c5761350b614171565b5b0361351a576001915061352d565b808061352590615000565b91505061342c565b508091505092915050565b5f33905090565b805f0151825f015114613587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161357e90615d37565b60405180910390fd5b80602001518260200151146135d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135c890615d9f565b60405180910390fd5b5050565b60405180608001604052805f81526020016135ee61383d565b81526020016135fb6138a4565b81526020016136086138e4565b81525090565b828054828255905f5260205f20908101928215613684579160200282015b82811115613683578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061362c565b5b5090506136919190613908565b5090565b604051806101e001604052805f600a8111156136b4576136b3614171565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b60405180606001604052805f81526020015f81526020015f81525090565b6040518060c001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020016060815260200160608152602001606081525090565b828054828255905f5260205f2090601f0160209004810192821561382c579160200282015f5b838211156137fe57835183826101000a81548160ff021916908360048111156137de576137dd614171565b5b021790555092602001926001016020815f010492830192600103026137b2565b801561382a5782816101000a81549060ff02191690556001016020815f010492830192600103026137fe565b505b5090506138399190613908565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b60405180608001604052806138b7613923565b81526020016138c4613923565b81526020016138d1613923565b81526020016138de613923565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b8082111561391f575f815f905550600101613909565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61395e8161394c565b8114613968575f80fd5b50565b5f8135905061397981613955565b92915050565b5f6020828403121561399457613993613944565b5b5f6139a18482850161396b565b91505092915050565b6139b38161394c565b82525050565b5f6020820190506139cc5f8301846139aa565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6139fb826139d2565b9050919050565b613a0b816139f1565b8114613a15575f80fd5b50565b5f81359050613a2681613a02565b92915050565b5f60208284031215613a4157613a40613944565b5b5f613a4e84828501613a18565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b613a898161394c565b82525050565b5f613a9a8383613a80565b60208301905092915050565b5f602082019050919050565b5f613abc82613a57565b613ac68185613a61565b9350613ad183613a71565b805f5b83811015613b01578151613ae88882613a8f565b9750613af383613aa6565b925050600181019050613ad4565b5085935050505092915050565b5f6020820190508181035f830152613b268184613ab2565b905092915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613b7882613b32565b810181811067ffffffffffffffff82111715613b9757613b96613b42565b5b80604052505050565b5f613ba961393b565b9050613bb58282613b6f565b919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115613bdc57613bdb613b42565b5b602082029050602081019050919050565b5f80fd5b5f613c03613bfe84613bc2565b613ba0565b90508083825260208201905060208402830185811115613c2657613c25613bed565b5b835b81811015613c4f5780613c3b8882613a18565b845260208401935050602081019050613c28565b5050509392505050565b5f82601f830112613c6d57613c6c613bbe565b5b8135613c7d848260208601613bf1565b91505092915050565b5f60808284031215613c9b57613c9a613b2e565b5b613ca56080613ba0565b90505f613cb484828501613a18565b5f830152506020613cc784828501613a18565b6020830152506040613cdb84828501613a18565b604083015250606082013567ffffffffffffffff811115613cff57613cfe613bba565b5b613d0b84828501613c59565b60608301525092915050565b5f60408284031215613d2c57613d2b613b2e565b5b613d366040613ba0565b90505f613d458482850161396b565b5f830152506020613d588482850161396b565b60208301525092915050565b5f6101008284031215613d7a57613d79613b2e565b5b613d846080613ba0565b90505f613d9384828501613d17565b5f830152506040613da684828501613d17565b6020830152506080613dba84828501613d17565b60408301525060c0613dce84828501613d17565b60608301525092915050565b5f60808284031215613def57613dee613b2e565b5b613df96080613ba0565b90505f613e088482850161396b565b5f830152506020613e1b8482850161396b565b6020830152506040613e2f8482850161396b565b6040830152506060613e438482850161396b565b60608301525092915050565b5f805f806101c08587031215613e6857613e67613944565b5b5f613e758782880161396b565b945050602085013567ffffffffffffffff811115613e9657613e95613948565b5b613ea287828801613c86565b9350506040613eb387828801613d64565b925050610140613ec587828801613dda565b91505092959194509250565b613eda816139f1565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f613f148383613ed1565b60208301905092915050565b5f602082019050919050565b5f613f3682613ee0565b613f408185613eea565b9350613f4b83613efa565b805f5b83811015613f7b578151613f628882613f09565b9750613f6d83613f20565b925050600181019050613f4e565b5085935050505092915050565b5f608083015f830151613f9d5f860182613ed1565b506020830151613fb06020860182613ed1565b506040830151613fc36040860182613ed1565b5060608301518482036060860152613fdb8282613f2c565b9150508091505092915050565b604082015f820151613ffc5f850182613a80565b50602082015161400f6020850182613a80565b50505050565b61010082015f82015161402a5f850182613fe8565b50602082015161403d6040850182613fe8565b5060408201516140506080850182613fe8565b50606082015161406360c0850182613fe8565b50505050565b608082015f82015161407d5f850182613a80565b5060208201516140906020850182613a80565b5060408201516140a36040850182613a80565b5060608201516140b66060850182613a80565b50505050565b5f6101c083015f8301516140d25f860182613a80565b50602083015184820360208601526140ea8282613f88565b91505060408301516140ff6040860182614015565b506060830151614113610140860182614069565b508091505092915050565b5f6020820190508181035f83015261413681846140bc565b905092915050565b5f8115159050919050565b6141528161413e565b82525050565b5f60208201905061416b5f830184614149565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b81106141af576141ae614171565b5b50565b5f8190506141bf8261419e565b919050565b5f6141ce826141b2565b9050919050565b6141de816141c4565b82525050565b6101e082015f8201516141f95f8501826141d5565b50602082015161420c6020850182613ed1565b50604082015161421f6040850182613a80565b5060608201516142326060850182613a80565b5060808201516142456080850182613a80565b5060a082015161425860a0850182613a80565b5060c082015161426b60c0850182613a80565b5060e082015161427e60e0850182613a80565b50610100820151614293610100850182613a80565b506101208201516142a8610120850182613a80565b506101408201516142bd610140850182613a80565b506101608201516142d2610160850182613a80565b506101808201516142e7610180850182613a80565b506101a08201516142fc6101a0850182613a80565b506101c08201516143116101c0850182613a80565b50505050565b5f6101e08201905061432b5f8301846141e4565b92915050565b5f805f6060848603121561434857614347613944565b5b5f6143558682870161396b565b93505060206143668682870161396b565b92505060406143778682870161396b565b9150509250925092565b606082015f8201516143955f850182613a80565b5060208201516143a86020850182613a80565b5060408201516143bb6040850182613a80565b50505050565b5f6060820190506143d45f830184614381565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156144115780820151818401526020810190506143f6565b5f8484015250505050565b5f614426826143da565b61443081856143e4565b93506144408185602086016143f4565b61444981613b32565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6005811061448e5761448d614171565b5b50565b5f81905061449e8261447d565b919050565b5f6144ad82614491565b9050919050565b6144bd816144a3565b82525050565b5f6144ce83836144b4565b60208301905092915050565b5f602082019050919050565b5f6144f082614454565b6144fa818561445e565b93506145058361446e565b805f5b8381101561453557815161451c88826144c3565b9750614527836144da565b925050600181019050614508565b5085935050505092915050565b5f60c083015f8301516145575f860182613ed1565b50602083015161456a6020860182613a80565b5060408301518482036040860152614582828261441c565b9150506060830151848203606086015261459c82826144e6565b915050608083015184820360808601526145b68282613f2c565b91505060a083015184820360a08601526145d08282613f2c565b9150508091505092915050565b5f6020820190508181035f8301526145f58184614542565b905092915050565b5f806040838503121561461357614612613944565b5b5f6146208582860161396b565b925050602061463185828601613a18565b9150509250929050565b614644816139f1565b82525050565b5f60208201905061465d5f83018461463b565b92915050565b6005811061466f575f80fd5b50565b5f8135905061468081614663565b92915050565b5f6020828403121561469b5761469a613944565b5b5f6146a884828501614672565b91505092915050565b5f82825260208201905092915050565b5f6146cb82613ee0565b6146d581856146b1565b93506146e083613efa565b805f5b838110156147105781516146f78882613f09565b975061470283613f20565b9250506001810190506146e3565b5085935050505092915050565b5f6020820190508181035f83015261473581846146c1565b905092915050565b5f80fd5b5f67ffffffffffffffff82111561475b5761475a613b42565b5b61476482613b32565b9050602081019050919050565b828183375f83830152505050565b5f61479161478c84614741565b613ba0565b9050828152602081018484840111156147ad576147ac61473d565b5b6147b8848285614771565b509392505050565b5f82601f8301126147d4576147d3613bbe565b5b81356147e484826020860161477f565b91505092915050565b5f67ffffffffffffffff82111561480757614806613b42565b5b602082029050602081019050919050565b5f61482a614825846147ed565b613ba0565b9050808382526020820190506020840283018581111561484d5761484c613bed565b5b835b8181101561487657806148628882614672565b84526020840193505060208101905061484f565b5050509392505050565b5f82601f83011261489457614893613bbe565b5b81356148a4848260208601614818565b91505092915050565b5f805f805f60a086880312156148c6576148c5613944565b5b5f6148d38882890161396b565b955050602086013567ffffffffffffffff8111156148f4576148f3613948565b5b614900888289016147c0565b945050604086013567ffffffffffffffff81111561492157614920613948565b5b61492d88828901614880565b935050606086013567ffffffffffffffff81111561494e5761494d613948565b5b61495a88828901613c59565b925050608086013567ffffffffffffffff81111561497b5761497a613948565b5b61498788828901613c59565b9150509295509295909350565b600b81106149a0575f80fd5b50565b5f813590506149b181614994565b92915050565b5f80604083850312156149cd576149cc613944565b5b5f6149da8582860161396b565b92505060206149eb858286016149a3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614a2c8261394c565b9150614a378361394c565b9250828202614a458161394c565b91508282048414831517614a5c57614a5b6149f5565b5b5092915050565b5f82825260208201905092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f614aa7600f83614a63565b9150614ab282614a73565b602082019050919050565b5f6020820190508181035f830152614ad481614a9b565b9050919050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f614b0f601083614a63565b9150614b1a82614adb565b602082019050919050565b5f6020820190508181035f830152614b3c81614b03565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f614b77600e83614a63565b9150614b8282614b43565b602082019050919050565b5f6020820190508181035f830152614ba481614b6b565b9050919050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f614bdf600a83614a63565b9150614bea82614bab565b602082019050919050565b5f6020820190508181035f830152614c0c81614bd3565b9050919050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f614c47601383614a63565b9150614c5282614c13565b602082019050919050565b5f6020820190508181035f830152614c7481614c3b565b9050919050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614caf601583614a63565b9150614cba82614c7b565b602082019050919050565b5f6020820190508181035f830152614cdc81614ca3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680614d2757607f821691505b602082108103614d3a57614d39614ce3565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f614d9a602e83614a63565b9150614da582614d40565b604082019050919050565b5f6020820190508181035f830152614dc781614d8e565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f614e06614e01614dfc84614dce565b614de3565b614dd7565b9050919050565b614e1681614dec565b82525050565b5f602082019050614e2f5f830184614e0d565b92915050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614e69601583614a63565b9150614e7482614e35565b602082019050919050565b5f6020820190508181035f830152614e9681614e5d565b9050919050565b7f55736572206d75737420657869737400000000000000000000000000000000005f82015250565b5f614ed1600f83614a63565b9150614edc82614e9d565b602082019050919050565b5f6020820190508181035f830152614efe81614ec5565b9050919050565b7f55736572206973206e6f742070617274206f662074686174206c6973740000005f82015250565b5f614f39601d83614a63565b9150614f4482614f05565b602082019050919050565b5f6020820190508181035f830152614f6681614f2d565b9050919050565b5f614f778261394c565b9150614f828361394c565b9250828203905081811115614f9a57614f996149f5565b5b92915050565b5f614faa8261394c565b9150614fb58361394c565b9250828201905080821115614fcd57614fcc6149f5565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61500a8261394c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361503c5761503b6149f5565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f5573657220697320616c72656164792070617274206f662074686174206c69735f8201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b5f6150ce602183614a63565b91506150d982615074565b604082019050919050565b5f6020820190508181035f8301526150fb816150c2565b9050919050565b7f55736572206d7573742068617665207468617420726f6c6500000000000000005f82015250565b5f615136601883614a63565b915061514182615102565b602082019050919050565b5f6020820190508181035f8301526151638161512a565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026151c67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261518b565b6151d0868361518b565b95508019841693508086168417925050509392505050565b5f6152026151fd6151f88461394c565b614de3565b61394c565b9050919050565b5f819050919050565b61521b836151e8565b61522f61522782615209565b848454615197565b825550505050565b5f90565b615243615237565b61524e818484615212565b505050565b5b81811015615271576152665f8261523b565b600181019050615254565b5050565b601f8211156152b6576152878161516a565b6152908461517c565b8101602085101561529f578190505b6152b36152ab8561517c565b830182615253565b50505b505050565b5f82821c905092915050565b5f6152d65f19846008026152bb565b1980831691505092915050565b5f6152ee83836152c7565b9150826002028217905092915050565b615307826143da565b67ffffffffffffffff8111156153205761531f613b42565b5b61532a8254614d10565b615335828285615275565b5f60209050601f831160018114615366575f8415615354578287015190505b61535e85826152e3565b8655506153c5565b601f1984166153748661516a565b5f5b8281101561539b57848901518255600182019150602085019450602081019050615376565b868310156153b857848901516153b4601f8916826152c7565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f615427602683614a63565b9150615432826153cd565b604082019050919050565b5f6020820190508181035f8301526154548161541b565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6154b5603583614a63565b91506154c08261545b565b604082019050919050565b5f6020820190508181035f8301526154e2816154a9565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f615543603983614a63565b915061554e826154e9565b604082019050919050565b5f6020820190508181035f83015261557081615537565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f6155d1603b83614a63565b91506155dc82615577565b604082019050919050565b5f6020820190508181035f8301526155fe816155c5565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f615639600a83614a63565b915061564482615605565b602082019050919050565b5f6020820190508181035f8301526156668161562d565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f6156a1600a83614a63565b91506156ac8261566d565b602082019050919050565b5f6020820190508181035f8301526156ce81615695565b9050919050565b7f446972206d697373696e670000000000000000000000000000000000000000005f82015250565b5f615709600b83614a63565b9150615714826156d5565b602082019050919050565b5f6020820190508181035f830152615736816156fd565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f615771600e83614a63565b915061577c8261573d565b602082019050919050565b5f6020820190508181035f83015261579e81615765565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f6157d9600c83614a63565b91506157e4826157a5565b602082019050919050565b5f6020820190508181035f830152615806816157cd565b9050919050565b7f4167726565206465706f736974203000000000000000000000000000000000005f82015250565b5f615841600f83614a63565b915061584c8261580d565b602082019050919050565b5f6020820190508181035f83015261586e81615835565b9050919050565b7f4d656469617465206465706f73697420300000000000000000000000000000005f82015250565b5f6158a9601183614a63565b91506158b482615875565b602082019050919050565b5f6020820190508181035f8301526158d68161589d565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f615911600283614a63565b915061591c826158dd565b602082019050919050565b5f6020820190508181035f83015261593e81615905565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f615979600283614a63565b915061598482615945565b602082019050919050565b5f6020820190508181035f8301526159a68161596d565b9050919050565b7f44697200000000000000000000000000000000000000000000000000000000005f82015250565b5f6159e1600383614a63565b91506159ec826159ad565b602082019050919050565b5f6020820190508181035f830152615a0e816159d5565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f615a49600983614a63565b9150615a5482615a15565b602082019050919050565b5f6020820190508181035f830152615a7681615a3d565b9050919050565b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f615ab1600883614a63565b9150615abc82615a7d565b602082019050919050565b5f6020820190508181035f830152615ade81615aa5565b9050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f615b19600583614a63565b9150615b2482615ae5565b602082019050919050565b5f6020820190508181035f830152615b4681615b0d565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f615b81600783614a63565b9150615b8c82615b4d565b602082019050919050565b5f6020820190508181035f830152615bae81615b75565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f615be9600783614a63565b9150615bf482615bb5565b602082019050919050565b5f6020820190508181035f830152615c1681615bdd565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f615c51600983614a63565b9150615c5c82615c1d565b602082019050919050565b5f6020820190508181035f830152615c7e81615c45565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f615cb9602083614a63565b9150615cc482615c85565b602082019050919050565b5f6020820190508181035f830152615ce681615cad565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f615d21600783614a63565b9150615d2c82615ced565b602082019050919050565b5f6020820190508181035f830152615d4e81615d15565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f615d89600a83614a63565b9150615d9482615d55565b602082019050919050565b5f6020820190508181035f830152615db681615d7d565b905091905056fea2646970667358221220f29bdb438bb5e4946c6e8d64397ca382fc42ade33270ceefd2e0c10593a1d9dd64736f6c63430008150033",
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

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetAgreement(opts *bind.CallOpts, dealId *big.Int) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
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

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCaller) GetDeal(opts *bind.CallOpts, dealId *big.Int) (SharedStructsDeal, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCallerSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageCallerSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageCaller) GetJobCost(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageCaller) GetResult(opts *bind.CallOpts, dealId *big.Int) (SharedStructsResult, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageCaller) GetResultsCollateral(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Storage.Contract.GetUser(&_Storage.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Storage.Contract.GetUser(&_Storage.CallOpts, userAddress)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageCaller) HasDeal(opts *bind.CallOpts, dealId *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageSession) HasDeal(dealId *big.Int) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageCallerSession) HasDeal(dealId *big.Int) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageCaller) IsState(opts *bind.CallOpts, dealId *big.Int, state uint8) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageCallerSession) IsState(dealId *big.Int, state uint8) (bool, error) {
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

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Storage.Contract.ShowUsersInList(&_Storage.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Storage.Contract.ShowUsersInList(&_Storage.CallOpts, serviceType)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.AddUserToList(&_Storage.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.AddUserToList(&_Storage.TransactOpts, serviceType)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
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

// EnsureDeal is a paid mutator transaction binding the contract method 0x3a83b3e4.
//
// Solidity: function ensureDeal(uint256 dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId *big.Int, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "ensureDeal", dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x3a83b3e4.
//
// Solidity: function ensureDeal(uint256 dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) EnsureDeal(dealId *big.Int, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x3a83b3e4.
//
// Solidity: function ensureDeal(uint256 dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((uint256,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactorSession) EnsureDeal(dealId *big.Int, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
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

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.RemoveUserFromList(&_Storage.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.RemoveUserFromList(&_Storage.TransactOpts, serviceType)
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

// TimeoutAgree is a paid mutator transaction binding the contract method 0x11d6adc4.
//
// Solidity: function timeoutAgree(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutAgree(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutAgree", dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x11d6adc4.
//
// Solidity: function timeoutAgree(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutAgree(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x11d6adc4.
//
// Solidity: function timeoutAgree(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutAgree(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
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

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "updateUser", metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpdateUser(&_Storage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageTransactorSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpdateUser(&_Storage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
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
	DealId *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) FilterDealStateChange(opts *bind.FilterOpts, dealId []*big.Int, state []uint8) (*StorageDealStateChangeIterator, error) {

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

// WatchDealStateChange is a free log subscription operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *StorageDealStateChange, dealId []*big.Int, state []uint8) (event.Subscription, error) {

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

// ParseDealStateChange is a log parse operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
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
