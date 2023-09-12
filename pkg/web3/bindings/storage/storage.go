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
	ResourceProvider common.Address
	JobCreator       common.Address
	Directory        common.Address
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"directory\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b615da080620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101e3575f3560e01c8063715018a61161010d578063a4702958116100a0578063d74553011161006f578063d7455301146105ab578063ef816fd9146105db578063f2fde38b1461060b578063f3d3d44814610627576101e3565b8063a470295814610539578063a7f96f0614610543578063aeb5ec011461055f578063c1390a891461057b576101e3565b80638462d54e116100dc5780638462d54e1461048b5780638da5cb5b146104bb578063995e4339146104d9578063a15dcc8a14610509576101e3565b8063715018a61461042b5780637fb9c45f146104355780638129fc1c1461045157806382fd5bac1461045b576101e3565b80633a83b3e41161018557806354278567116101545780635427856714610393578063546cfd34146103af5780635dd049fd146103df5780636f77926b146103fb576101e3565b80633a83b3e4146102d3578063407c2d621461030357806342e4074e146103335780634f9f6fe614610363576101e3565b8063172257c7116101c1578063172257c71461026357806323a9a8621461027f57806332ebea041461029b57806335a7e268146102b7576101e3565b80630396e3c1146101e757806311d5af331461021757806311d6adc414610247575b5f80fd5b61020160048036038101906101fc919061392c565b610643565b60405161020e9190613966565b60405180910390f35b610231600480360381019061022c91906139d9565b610681565b60405161023e9190613abb565b60405180910390f35b610261600480360381019061025c919061392c565b610714565b005b61027d6004803603810190610278919061392c565b61078d565b005b6102996004803603810190610294919061392c565b610807565b005b6102b560048036038101906102b0919061392c565b610881565b005b6102d160048036038101906102cc919061392c565b6108fb565b005b6102ed60048036038101906102e89190613dfc565b610975565b6040516102fa91906140cb565b60405180910390f35b61031d6004803603810190610318919061392c565b610f8e565b60405161032a9190614105565b60405180910390f35b61034d6004803603810190610348919061392c565b610fa5565b60405161035a91906142c4565b60405180910390f35b61037d6004803603810190610378919061392c565b6111aa565b60405161038a91906142c4565b60405180910390f35b6103ad60048036038101906103a8919061392c565b6112e6565b005b6103c960048036038101906103c491906142de565b611360565b6040516103d6919061436e565b60405180910390f35b6103f960048036038101906103f4919061392c565b61146c565b005b610415600480360381019061041091906139d9565b6114e6565b604051610422919061458a565b60405180910390f35b6104336117ca565b005b61044f600480360381019061044a91906145aa565b6117dd565b005b6104596118aa565b005b6104756004803603810190610470919061392c565b6119e1565b60405161048291906140cb565b60405180910390f35b6104a560048036038101906104a0919061392c565b611c9c565b6040516104b291906142c4565b60405180910390f35b6104c3611ea1565b6040516104d091906145f7565b60405180910390f35b6104f360048036038101906104ee919061392c565b611ec8565b604051610500919061436e565b60405180910390f35b610523600480360381019061051e9190614633565b611f10565b60405161053091906146ca565b60405180910390f35b610541611fd0565b005b61055d60048036038101906105589190614633565b611ff4565b005b61057960048036038101906105749190614633565b6122f8565b005b6105956004803603810190610590919061485a565b6124f3565b6040516105a2919061458a565b60405180910390f35b6105c560048036038101906105c0919061392c565b61264a565b6040516105d29190613966565b60405180910390f35b6105f560048036038101906105f09190614964565b61267d565b6040516106029190614105565b60405180910390f35b610625600480360381019061062091906139d9565b612708565b005b610641600480360381019061063c91906139d9565b61278a565b005b5f60085f8381526020019081526020015f206002015460055f8481526020019081526020015f20600d015f015461067a91906149cf565b9050919050565b606060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561070857602002820191905f5260205f20905b8154815260200190600101908083116106f4575b50505050509050919050565b61071c612892565b50610727815f61267d565b610766576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075d90614a6a565b60405180910390fd5b4260075f8381526020019081526020015f20600a018190555061078a8160076129bf565b50565b610795612892565b506107a181600261267d565b6107e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d790614ad2565b60405180910390fd5b4260075f8381526020019081526020015f20600601819055506108048160036129bf565b50565b61080f612892565b5061081b81600461267d565b61085a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085190614b3a565b60405180910390fd5b4260075f8381526020019081526020015f206008018190555061087e8160056129bf565b50565b610889612892565b5061089581600161267d565b6108d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108cb90614ba2565b60405180910390fd5b4260075f8381526020019081526020015f20600b01819055506108f88160086129bf565b50565b610903612892565b5061090f81600461267d565b61094e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094590614b3a565b60405180910390fd5b4260075f8381526020019081526020015f20600d018190555061097281600a6129bf565b50565b61097d613582565b610985612892565b50610990855f61267d565b6109cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c690614a6a565b60405180910390fd5b6109d884612a3d565b6109e183612c51565b6109ea85610f8e565b15610a2b575f6109f9866119e1565b9050610a09816020015186612ce7565b610a17816040015185612f64565b610a25816060015184612fae565b50610cd8565b60405180608001604052808681526020018581526020018481526020018381525060055f8781526020019081526020015f205f820151815f01556020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190610b5a9291906135bb565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060065f855f015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f909190919091505560065f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150555b60055f8681526020019081526020015f206040518060800160405290815f8201548152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015610e8d57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e44575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f80610f99836119e1565b5f015114159050919050565b610fad613642565b610fb5612892565b50610fbf82610f8e565b610ffe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff590614c0a565b60405180910390fd5b5f60075f8481526020019081526020015f206001015414611054576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104b90614c72565b60405180910390fd5b4260075f8481526020019081526020015f2060010181905550611076826130d8565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a8111156110b6576110b561411e565b5b600a8111156110c8576110c761411e565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6111b2613642565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a8111156111f2576111f161411e565b5b600a8111156112045761120361411e565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6112ee612892565b506112fa81600261267d565b611339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133090614ad2565b60405180910390fd5b4260075f8381526020019081526020015f20600c018190555061135d8160096129bf565b50565b6113686136d1565b611370612892565b5061137c84600161267d565b6113bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b290614ba2565b60405180910390fd5b4260075f8681526020019081526020015f20600501819055506113df8460026129bf565b60405180606001604052808581526020018481526020018381525060085f8681526020019081526020015f205f820151815f0155602082015181600101556040820151816002015590505060085f8581526020019081526020015f206040518060600160405290815f82015481526020016001820154815260200160028201548152505090509392505050565b611474612892565b5061148081600461267d565b6114bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b690614b3a565b60405180910390fd5b4260075f8381526020019081526020015f20600901819055506114e38160066129bf565b50565b6114ee6136ef565b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060c00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820180546115a490614cbd565b80601f01602080910402602001604051908101604052809291908181526020018280546115d090614cbd565b801561161b5780601f106115f25761010080835404028352916020019161161b565b820191905f5260205f20905b8154815290600101906020018083116115fe57829003601f168201915b50505050508152602001600382018054806020026020016040519081016040528092919081815260200182805480156116a457602002820191905f5260205f20905f905b82829054906101000a900460ff16600481111561167f5761167e61411e565b5b815260200190600101906020825f0104928301926001038202915080841161165f5790505b505050505081526020016004820180548060200260200160405190810160405280929190818152602001828054801561172f57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116116e6575b50505050508152602001600582018054806020026020016040519081016040528092919081815260200182805480156117ba57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611771575b5050505050815250509050919050565b6117d261315b565b6117db5f6131d9565b565b6117e5612892565b506117f182600261267d565b611830576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161182790614ad2565b60405180910390fd5b4260075f8481526020019081526020015f20600701819055508060075f8481526020019081526020015f205f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118a68260046129bf565b5050565b5f600160169054906101000a900460ff161590508080156118dc575060018060159054906101000a900460ff1660ff16105b8061190a57506118eb3061329a565b158015611909575060018060159054906101000a900460ff1660ff16145b5b611949576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161194090614d5d565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156119855760018060166101000a81548160ff0219169083151502179055505b80156119de575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516119d59190614dc9565b60405180910390a15b50565b6119e9613582565b60055f8381526020019081526020015f206040518060800160405290815f8201548152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611b9e57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611b55575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611ca4613642565b611cac612892565b50611cb682610f8e565b611cf5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cec90614c0a565b60405180910390fd5b5f60075f8481526020019081526020015f206002015414611d4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d4290614e2c565b60405180910390fd5b4260075f8481526020019081526020015f2060020181905550611d6d826130d8565b60075f8381526020019081526020015f20604051806101e00160405290815f82015f9054906101000a900460ff16600a811115611dad57611dac61411e565b5b600a811115611dbf57611dbe61411e565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611ed06136d1565b60085f8381526020019081526020015f206040518060600160405290815f8201548152602001600182015481526020016002820154815250509050919050565b606060045f836004811115611f2857611f2761411e565b5b6004811115611f3a57611f3961411e565b5b81526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015611fc457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611f7b575b50505050509050919050565b611fd861315b565b5f600160146101000a81548160ff021916908315150217905550565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036120c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b790614e94565b60405180910390fd5b5f806120cc83326132bc565b9150915080612110576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161210790614efc565b60405180910390fd5b5f8290505b600160045f86600481111561212d5761212c61411e565b5b600481111561213f5761213e61411e565b5b81526020019081526020015f208054905061215a9190614f1a565b81101561227b5760045f8560048111156121775761217661411e565b5b60048111156121895761218861411e565b5b81526020019081526020015f206001826121a39190614f4d565b815481106121b4576121b3614f80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660045f8660048111156121f2576121f161411e565b5b60048111156122045761220361411e565b5b81526020019081526020015f20828154811061222357612222614f80565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061227390614fad565b915050612115565b5060045f8460048111156122925761229161411e565b5b60048111156122a4576122a361411e565b5b81526020019081526020015f208054806122c1576122c0614ff4565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036123c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123bb90614e94565b60405180910390fd5b5f6123cf82326132bc565b9150508015612413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161240a90615091565b60405180910390fd5b61241d82326133d2565b61245c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612453906150f9565b60405180910390fd5b60045f8360048111156124725761247161411e565b5b60048111156124845761248361411e565b5b81526020019081526020015f2032908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6124fb6136ef565b5f6040518060c001604052803273ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018681526020018581526020018481525090508060035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190816125e291906152ab565b5060608201518160030190805190602001906125ff929190613739565b50608082015181600401908051906020019061261c9291906135bb565b5060a08201518160050190805190602001906126399291906135bb565b509050508091505095945050505050565b5f61265482610643565b60055f8481526020019081526020015f20600d016002015461267691906149cf565b9050919050565b5f61268783610f8e565b6126b9575f600a81111561269e5761269d61411e565b5b82600a8111156126b1576126b061411e565b5b149050612702565b81600a8111156126cc576126cb61411e565b5b60075f8581526020019081526020015f205f015f9054906101000a900460ff16600a8111156126fe576126fd61411e565b5b1490505b92915050565b61271061315b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361277e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612775906153ea565b60405180910390fd5b612787816131d9565b50565b61279261315b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f790615478565b60405180910390fd5b600160149054906101000a900460ff1661284f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284690615506565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612922576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161291990615478565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166129626134e5565b73ffffffffffffffffffffffffffffffffffffffff16146129b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129af90615594565b60405180910390fd5b6001905090565b8060075f8481526020019081526020015f205f015f6101000a81548160ff0219169083600a8111156129f4576129f361411e565b5b021790555080600a811115612a0c57612a0b61411e565b5b827f17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d60405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612aae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612aa5906155fc565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612b20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b1790615664565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612b92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b89906156cc565b60405180910390fd5b5f81606001515111612bd9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bd090615734565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff1603612c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c459061579c565b60405180910390fd5b50565b5f815f01516020015114612c9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c9190615804565b60405180910390fd5b5f81606001516020015114612ce4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cdb9061586c565b60405180910390fd5b50565b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff1614612d5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d52906158d4565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612dd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dc89061593c565b60405180910390fd5b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff1614612e47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e3e906159a4565b60405180910390fd5b80606001515182606001515114612e93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e8a90615a0c565b60405180910390fd5b5f5b826060015151811015612f5f5781606001518181518110612eb957612eb8614f80565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612eee57612eed614f80565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f4390615a74565b60405180910390fd5b8080612f5790614fad565b915050612e95565b505050565b612f74825f0151825f01516134ec565b612f86826020015182602001516134ec565b612f98826040015182604001516134ec565b612faa826060015182606001516134ec565b5050565b805f0151825f015114612ff6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fed90615adc565b60405180910390fd5b8060200151826020015114613040576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161303790615b44565b60405180910390fd5b806040015182604001511461308a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161308190615bac565b60405180910390fd5b80606001518260600151146130d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130cb90615c14565b60405180910390fd5b5050565b5f60075f8381526020019081526020015f20600101541415801561311057505f60075f8381526020019081526020015f206002015414155b1561313e574260075f8381526020019081526020015f20600401819055506131398160016129bf565b613158565b4260075f8381526020019081526020015f20600301819055505b50565b6131636134e5565b73ffffffffffffffffffffffffffffffffffffffff16613181611ea1565b73ffffffffffffffffffffffffffffffffffffffff16146131d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131ce90615c7c565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f805f805f5b60045f8860048111156132d8576132d761411e565b5b60048111156132ea576132e961411e565b5b81526020019081526020015f20805490508110156133c2578573ffffffffffffffffffffffffffffffffffffffff1660045f89600481111561332f5761332e61411e565b5b60048111156133415761334061411e565b5b81526020019081526020015f2082815481106133605761335f614f80565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036133af57809250600191506133c2565b80806133ba90614fad565b9150506132c2565b5081819350935050509250929050565b5f805f90505f5b60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301805490508110156134da578460048111156134375761343661411e565b5b60035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600301828154811061348957613488614f80565b5b905f5260205f2090602091828204019190069054906101000a900460ff1660048111156134b9576134b861411e565b5b036134c757600191506134da565b80806134d290614fad565b9150506133d9565b508091505092915050565b5f33905090565b805f0151825f015114613534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161352b90615ce4565b60405180910390fd5b806020015182602001511461357e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161357590615d4c565b60405180910390fd5b5050565b60405180608001604052805f815260200161359b6137ea565b81526020016135a8613851565b81526020016135b5613891565b81525090565b828054828255905f5260205f20908101928215613631579160200282015b82811115613630578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906135d9565b5b50905061363e91906138b5565b5090565b604051806101e001604052805f600a8111156136615761366061411e565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b60405180606001604052805f81526020015f81526020015f81525090565b6040518060c001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020016060815260200160608152602001606081525090565b828054828255905f5260205f2090601f016020900481019282156137d9579160200282015f5b838211156137ab57835183826101000a81548160ff0219169083600481111561378b5761378a61411e565b5b021790555092602001926001016020815f0104928301926001030261375f565b80156137d75782816101000a81549060ff02191690556001016020815f010492830192600103026137ab565b505b5090506137e691906138b5565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b60405180608001604052806138646138d0565b81526020016138716138d0565b815260200161387e6138d0565b815260200161388b6138d0565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b808211156138cc575f815f9055506001016138b6565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61390b816138f9565b8114613915575f80fd5b50565b5f8135905061392681613902565b92915050565b5f60208284031215613941576139406138f1565b5b5f61394e84828501613918565b91505092915050565b613960816138f9565b82525050565b5f6020820190506139795f830184613957565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6139a88261397f565b9050919050565b6139b88161399e565b81146139c2575f80fd5b50565b5f813590506139d3816139af565b92915050565b5f602082840312156139ee576139ed6138f1565b5b5f6139fb848285016139c5565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b613a36816138f9565b82525050565b5f613a478383613a2d565b60208301905092915050565b5f602082019050919050565b5f613a6982613a04565b613a738185613a0e565b9350613a7e83613a1e565b805f5b83811015613aae578151613a958882613a3c565b9750613aa083613a53565b925050600181019050613a81565b5085935050505092915050565b5f6020820190508181035f830152613ad38184613a5f565b905092915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613b2582613adf565b810181811067ffffffffffffffff82111715613b4457613b43613aef565b5b80604052505050565b5f613b566138e8565b9050613b628282613b1c565b919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115613b8957613b88613aef565b5b602082029050602081019050919050565b5f80fd5b5f613bb0613bab84613b6f565b613b4d565b90508083825260208201905060208402830185811115613bd357613bd2613b9a565b5b835b81811015613bfc5780613be888826139c5565b845260208401935050602081019050613bd5565b5050509392505050565b5f82601f830112613c1a57613c19613b6b565b5b8135613c2a848260208601613b9e565b91505092915050565b5f60808284031215613c4857613c47613adb565b5b613c526080613b4d565b90505f613c61848285016139c5565b5f830152506020613c74848285016139c5565b6020830152506040613c88848285016139c5565b604083015250606082013567ffffffffffffffff811115613cac57613cab613b67565b5b613cb884828501613c06565b60608301525092915050565b5f60408284031215613cd957613cd8613adb565b5b613ce36040613b4d565b90505f613cf284828501613918565b5f830152506020613d0584828501613918565b60208301525092915050565b5f6101008284031215613d2757613d26613adb565b5b613d316080613b4d565b90505f613d4084828501613cc4565b5f830152506040613d5384828501613cc4565b6020830152506080613d6784828501613cc4565b60408301525060c0613d7b84828501613cc4565b60608301525092915050565b5f60808284031215613d9c57613d9b613adb565b5b613da66080613b4d565b90505f613db584828501613918565b5f830152506020613dc884828501613918565b6020830152506040613ddc84828501613918565b6040830152506060613df084828501613918565b60608301525092915050565b5f805f806101c08587031215613e1557613e146138f1565b5b5f613e2287828801613918565b945050602085013567ffffffffffffffff811115613e4357613e426138f5565b5b613e4f87828801613c33565b9350506040613e6087828801613d11565b925050610140613e7287828801613d87565b91505092959194509250565b613e878161399e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f613ec18383613e7e565b60208301905092915050565b5f602082019050919050565b5f613ee382613e8d565b613eed8185613e97565b9350613ef883613ea7565b805f5b83811015613f28578151613f0f8882613eb6565b9750613f1a83613ecd565b925050600181019050613efb565b5085935050505092915050565b5f608083015f830151613f4a5f860182613e7e565b506020830151613f5d6020860182613e7e565b506040830151613f706040860182613e7e565b5060608301518482036060860152613f888282613ed9565b9150508091505092915050565b604082015f820151613fa95f850182613a2d565b506020820151613fbc6020850182613a2d565b50505050565b61010082015f820151613fd75f850182613f95565b506020820151613fea6040850182613f95565b506040820151613ffd6080850182613f95565b50606082015161401060c0850182613f95565b50505050565b608082015f82015161402a5f850182613a2d565b50602082015161403d6020850182613a2d565b5060408201516140506040850182613a2d565b5060608201516140636060850182613a2d565b50505050565b5f6101c083015f83015161407f5f860182613a2d565b50602083015184820360208601526140978282613f35565b91505060408301516140ac6040860182613fc2565b5060608301516140c0610140860182614016565b508091505092915050565b5f6020820190508181035f8301526140e38184614069565b905092915050565b5f8115159050919050565b6140ff816140eb565b82525050565b5f6020820190506141185f8301846140f6565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b811061415c5761415b61411e565b5b50565b5f81905061416c8261414b565b919050565b5f61417b8261415f565b9050919050565b61418b81614171565b82525050565b6101e082015f8201516141a65f850182614182565b5060208201516141b96020850182613e7e565b5060408201516141cc6040850182613a2d565b5060608201516141df6060850182613a2d565b5060808201516141f26080850182613a2d565b5060a082015161420560a0850182613a2d565b5060c082015161421860c0850182613a2d565b5060e082015161422b60e0850182613a2d565b50610100820151614240610100850182613a2d565b50610120820151614255610120850182613a2d565b5061014082015161426a610140850182613a2d565b5061016082015161427f610160850182613a2d565b50610180820151614294610180850182613a2d565b506101a08201516142a96101a0850182613a2d565b506101c08201516142be6101c0850182613a2d565b50505050565b5f6101e0820190506142d85f830184614191565b92915050565b5f805f606084860312156142f5576142f46138f1565b5b5f61430286828701613918565b935050602061431386828701613918565b925050604061432486828701613918565b9150509250925092565b606082015f8201516143425f850182613a2d565b5060208201516143556020850182613a2d565b5060408201516143686040850182613a2d565b50505050565b5f6060820190506143815f83018461432e565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156143be5780820151818401526020810190506143a3565b5f8484015250505050565b5f6143d382614387565b6143dd8185614391565b93506143ed8185602086016143a1565b6143f681613adf565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6005811061443b5761443a61411e565b5b50565b5f81905061444b8261442a565b919050565b5f61445a8261443e565b9050919050565b61446a81614450565b82525050565b5f61447b8383614461565b60208301905092915050565b5f602082019050919050565b5f61449d82614401565b6144a7818561440b565b93506144b28361441b565b805f5b838110156144e25781516144c98882614470565b97506144d483614487565b9250506001810190506144b5565b5085935050505092915050565b5f60c083015f8301516145045f860182613e7e565b5060208301516145176020860182613a2d565b506040830151848203604086015261452f82826143c9565b915050606083015184820360608601526145498282614493565b915050608083015184820360808601526145638282613ed9565b91505060a083015184820360a086015261457d8282613ed9565b9150508091505092915050565b5f6020820190508181035f8301526145a281846144ef565b905092915050565b5f80604083850312156145c0576145bf6138f1565b5b5f6145cd85828601613918565b92505060206145de858286016139c5565b9150509250929050565b6145f18161399e565b82525050565b5f60208201905061460a5f8301846145e8565b92915050565b6005811061461c575f80fd5b50565b5f8135905061462d81614610565b92915050565b5f60208284031215614648576146476138f1565b5b5f6146558482850161461f565b91505092915050565b5f82825260208201905092915050565b5f61467882613e8d565b614682818561465e565b935061468d83613ea7565b805f5b838110156146bd5781516146a48882613eb6565b97506146af83613ecd565b925050600181019050614690565b5085935050505092915050565b5f6020820190508181035f8301526146e2818461466e565b905092915050565b5f80fd5b5f67ffffffffffffffff82111561470857614707613aef565b5b61471182613adf565b9050602081019050919050565b828183375f83830152505050565b5f61473e614739846146ee565b613b4d565b90508281526020810184848401111561475a576147596146ea565b5b61476584828561471e565b509392505050565b5f82601f83011261478157614780613b6b565b5b813561479184826020860161472c565b91505092915050565b5f67ffffffffffffffff8211156147b4576147b3613aef565b5b602082029050602081019050919050565b5f6147d76147d28461479a565b613b4d565b905080838252602082019050602084028301858111156147fa576147f9613b9a565b5b835b81811015614823578061480f888261461f565b8452602084019350506020810190506147fc565b5050509392505050565b5f82601f83011261484157614840613b6b565b5b81356148518482602086016147c5565b91505092915050565b5f805f805f60a08688031215614873576148726138f1565b5b5f61488088828901613918565b955050602086013567ffffffffffffffff8111156148a1576148a06138f5565b5b6148ad8882890161476d565b945050604086013567ffffffffffffffff8111156148ce576148cd6138f5565b5b6148da8882890161482d565b935050606086013567ffffffffffffffff8111156148fb576148fa6138f5565b5b61490788828901613c06565b925050608086013567ffffffffffffffff811115614928576149276138f5565b5b61493488828901613c06565b9150509295509295909350565b600b811061494d575f80fd5b50565b5f8135905061495e81614941565b92915050565b5f806040838503121561497a576149796138f1565b5b5f61498785828601613918565b925050602061499885828601614950565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6149d9826138f9565b91506149e4836138f9565b92508282026149f2816138f9565b91508282048414831517614a0957614a086149a2565b5b5092915050565b5f82825260208201905092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f614a54600f83614a10565b9150614a5f82614a20565b602082019050919050565b5f6020820190508181035f830152614a8181614a48565b9050919050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f614abc601083614a10565b9150614ac782614a88565b602082019050919050565b5f6020820190508181035f830152614ae981614ab0565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f614b24600e83614a10565b9150614b2f82614af0565b602082019050919050565b5f6020820190508181035f830152614b5181614b18565b9050919050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f614b8c600a83614a10565b9150614b9782614b58565b602082019050919050565b5f6020820190508181035f830152614bb981614b80565b9050919050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f614bf4601383614a10565b9150614bff82614bc0565b602082019050919050565b5f6020820190508181035f830152614c2181614be8565b9050919050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614c5c601583614a10565b9150614c6782614c28565b602082019050919050565b5f6020820190508181035f830152614c8981614c50565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680614cd457607f821691505b602082108103614ce757614ce6614c90565b5b50919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f614d47602e83614a10565b9150614d5282614ced565b604082019050919050565b5f6020820190508181035f830152614d7481614d3b565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f614db3614dae614da984614d7b565b614d90565b614d84565b9050919050565b614dc381614d99565b82525050565b5f602082019050614ddc5f830184614dba565b92915050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614e16601583614a10565b9150614e2182614de2565b602082019050919050565b5f6020820190508181035f830152614e4381614e0a565b9050919050565b7f55736572206d75737420657869737400000000000000000000000000000000005f82015250565b5f614e7e600f83614a10565b9150614e8982614e4a565b602082019050919050565b5f6020820190508181035f830152614eab81614e72565b9050919050565b7f55736572206973206e6f742070617274206f662074686174206c6973740000005f82015250565b5f614ee6601d83614a10565b9150614ef182614eb2565b602082019050919050565b5f6020820190508181035f830152614f1381614eda565b9050919050565b5f614f24826138f9565b9150614f2f836138f9565b9250828203905081811115614f4757614f466149a2565b5b92915050565b5f614f57826138f9565b9150614f62836138f9565b9250828201905080821115614f7a57614f796149a2565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f614fb7826138f9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614fe957614fe86149a2565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f5573657220697320616c72656164792070617274206f662074686174206c69735f8201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b5f61507b602183614a10565b915061508682615021565b604082019050919050565b5f6020820190508181035f8301526150a88161506f565b9050919050565b7f55736572206d7573742068617665207468617420726f6c6500000000000000005f82015250565b5f6150e3601883614a10565b91506150ee826150af565b602082019050919050565b5f6020820190508181035f830152615110816150d7565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026151737fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82615138565b61517d8683615138565b95508019841693508086168417925050509392505050565b5f6151af6151aa6151a5846138f9565b614d90565b6138f9565b9050919050565b5f819050919050565b6151c883615195565b6151dc6151d4826151b6565b848454615144565b825550505050565b5f90565b6151f06151e4565b6151fb8184846151bf565b505050565b5b8181101561521e576152135f826151e8565b600181019050615201565b5050565b601f8211156152635761523481615117565b61523d84615129565b8101602085101561524c578190505b61526061525885615129565b830182615200565b50505b505050565b5f82821c905092915050565b5f6152835f1984600802615268565b1980831691505092915050565b5f61529b8383615274565b9150826002028217905092915050565b6152b482614387565b67ffffffffffffffff8111156152cd576152cc613aef565b5b6152d78254614cbd565b6152e2828285615222565b5f60209050601f831160018114615313575f8415615301578287015190505b61530b8582615290565b865550615372565b601f19841661532186615117565b5f5b8281101561534857848901518255600182019150602085019450602081019050615323565b868310156153655784890151615361601f891682615274565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6153d4602683614a10565b91506153df8261537a565b604082019050919050565b5f6020820190508181035f830152615401816153c8565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f615462603583614a10565b915061546d82615408565b604082019050919050565b5f6020820190508181035f83015261548f81615456565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f6154f0603983614a10565b91506154fb82615496565b604082019050919050565b5f6020820190508181035f83015261551d816154e4565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f61557e603b83614a10565b915061558982615524565b604082019050919050565b5f6020820190508181035f8301526155ab81615572565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f6155e6600a83614a10565b91506155f1826155b2565b602082019050919050565b5f6020820190508181035f830152615613816155da565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f61564e600a83614a10565b91506156598261561a565b602082019050919050565b5f6020820190508181035f83015261567b81615642565b9050919050565b7f446972206d697373696e670000000000000000000000000000000000000000005f82015250565b5f6156b6600b83614a10565b91506156c182615682565b602082019050919050565b5f6020820190508181035f8301526156e3816156aa565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f61571e600e83614a10565b9150615729826156ea565b602082019050919050565b5f6020820190508181035f83015261574b81615712565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f615786600c83614a10565b915061579182615752565b602082019050919050565b5f6020820190508181035f8301526157b38161577a565b9050919050565b7f4167726565206465706f736974203000000000000000000000000000000000005f82015250565b5f6157ee600f83614a10565b91506157f9826157ba565b602082019050919050565b5f6020820190508181035f83015261581b816157e2565b9050919050565b7f4d656469617465206465706f73697420300000000000000000000000000000005f82015250565b5f615856601183614a10565b915061586182615822565b602082019050919050565b5f6020820190508181035f8301526158838161584a565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f6158be600283614a10565b91506158c98261588a565b602082019050919050565b5f6020820190508181035f8301526158eb816158b2565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f615926600283614a10565b9150615931826158f2565b602082019050919050565b5f6020820190508181035f8301526159538161591a565b9050919050565b7f44697200000000000000000000000000000000000000000000000000000000005f82015250565b5f61598e600383614a10565b91506159998261595a565b602082019050919050565b5f6020820190508181035f8301526159bb81615982565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f6159f6600983614a10565b9150615a01826159c2565b602082019050919050565b5f6020820190508181035f830152615a23816159ea565b9050919050565b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f615a5e600883614a10565b9150615a6982615a2a565b602082019050919050565b5f6020820190508181035f830152615a8b81615a52565b9050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f615ac6600583614a10565b9150615ad182615a92565b602082019050919050565b5f6020820190508181035f830152615af381615aba565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f615b2e600783614a10565b9150615b3982615afa565b602082019050919050565b5f6020820190508181035f830152615b5b81615b22565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f615b96600783614a10565b9150615ba182615b62565b602082019050919050565b5f6020820190508181035f830152615bc381615b8a565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f615bfe600983614a10565b9150615c0982615bca565b602082019050919050565b5f6020820190508181035f830152615c2b81615bf2565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f615c66602083614a10565b9150615c7182615c32565b602082019050919050565b5f6020820190508181035f830152615c9381615c5a565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f615cce600783614a10565b9150615cd982615c9a565b602082019050919050565b5f6020820190508181035f830152615cfb81615cc2565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f615d36600a83614a10565b9150615d4182615d02565b602082019050919050565b5f6020820190508181035f830152615d6381615d2a565b905091905056fea2646970667358221220e15b8905f4ec2db176a37e5e952623801df187abe1dd8e6ea0a6cb05f747fd7964736f6c63430008150033",
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
