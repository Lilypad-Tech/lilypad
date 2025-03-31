// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lilypadpaymentengine

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

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId           string
	JobCreator       common.Address
	ResourceProvider common.Address
	ModuleCreator    common.Address
	Solver           common.Address
	JobOfferCID      string
	ResourceOfferCID string
	Status           uint8
	Timestamp        *big.Int
	PaymentStructure SharedStructsDealPaymentStructure
}

// SharedStructsDealPaymentStructure is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPaymentStructure struct {
	JobCreatorSolverFee       *big.Int
	ResourceProviderSolverFee *big.Int
	NetworkCongestionFee      *big.Int
	ModuleCreatorFee          *big.Int
	PriceOfJobWithoutFees     *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	ResultId  string
	DealId    string
	ResultCID string
	Status    uint8
	Timestamp *big.Int
}

// SharedStructsValidationResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsValidationResult struct {
	ValidationResultId string
	ResultId           string
	ValidationCID      string
	Status             uint8
	Timestamp          *big.Int
	Validator          common.Address
}

// LilypadPaymentEngineMetaData contains all meta data concerning the LilypadPaymentEngine contract.
var LilypadPaymentEngineMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COLLATERAL_LOCK_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_RESOURCE_PROVIDER_DEPOSIT_AMOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeBurnTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeEscrow\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"activeEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canWithdrawEscrow\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositTimestamps\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"depositTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"escrowBalances\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleJobCompletion\",\"inputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleJobFailure\",\"inputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Result\",\"components\":[{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleValidationFailed\",\"inputs\":[{\"name\":\"_validationResult\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.ValidationResult\",\"components\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validationCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"_originalJobDeal\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.Deal\",\"components\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"moduleCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jobOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resourceOfferCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.DealStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentStructure\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.DealPaymentStructure\",\"components\":[{\"name\":\"jobCreatorSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderSolverFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"networkCongestionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"moduleCreatorFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceOfJobWithoutFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handleValidationPassed\",\"inputs\":[{\"name\":\"_validationResult\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.ValidationResult\",\"components\":[{\"name\":\"validationResultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"resultId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"validationCID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.ValidationResultStatusEnum\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_l2token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadUserAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lilypadTokenomicsAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasuryWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_valueBasedRewardsWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_validationPoolWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateLockupOfEscrowForJob\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderCollateralLockupAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payEscrow\",\"inputs\":[{\"name\":\"_payee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paymentReason\",\"type\":\"uint8\",\"internalType\":\"enumSharedStructs.PaymentReason\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL2Token\",\"inputs\":[{\"name\":\"_l2tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadStorage\",\"inputs\":[{\"name\":\"_lilypadStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadTokenomics\",\"inputs\":[{\"name\":\"_lilypadTokenomicsAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLilypadUser\",\"inputs\":[{\"name\":\"_lilypadUserAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTreasuryWallet\",\"inputs\":[{\"name\":\"_treasuryWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValidationPoolWallet\",\"inputs\":[{\"name\":\"_validationPoolWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setValueBasedRewardsWallet\",\"inputs\":[{\"name\":\"_valueBasedRewardsWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalActiveEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasuryWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateActiveBurnTokens\",\"inputs\":[{\"name\":\"_amountBurnt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validationPoolWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"valueBasedRewardsWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawEscrow\",\"inputs\":[{\"name\":\"_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ActiveEscrowLockedForJob\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ControllerRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ControllerRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__JobCompleted\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__JobFailed\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resultId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__TokensBurned\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockTimestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountBurnt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__TotalFeesGeneratedByJob\",\"inputs\":[{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dealId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__TreasuryWalletUpdated\",\"inputs\":[{\"name\":\"newTreasuryWallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ValidationFailed\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ValidationPassed\",\"inputs\":[{\"name\":\"jobCreator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"resourceProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ValidationPoolWalletUpdated\",\"inputs\":[{\"name\":\"newValidationPoolWallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ValueBasedRewardsWalletUpdated\",\"inputs\":[{\"name\":\"newValueBasedRewardsWallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__ZeroAmountPayout\",\"inputs\":[{\"name\":\"intended_recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__escrowPaid\",\"inputs\":[{\"name\":\"payee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentReason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumSharedStructs.PaymentReason\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__escrowPayout\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__escrowSlashed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"actor\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumSharedStructs.UserType\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LilypadPayment__escrowWithdrawn\",\"inputs\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__CannotRevokeOwnRole\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__HandleJobCompletion__InsufficientActiveEscrowToCompleteJob\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreatorActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCostOfJob\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderRequiredActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LilypadPayment__HandleJobCompletion__InvalidTreasuryAmounts\",\"inputs\":[{\"name\":\"pValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"p1Value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"p2Value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"p3Value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LilypadPayment__HandleJobFailure__InsufficientActiveEscrowToCompleteJob\",\"inputs\":[{\"name\":\"dealId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"jobCreatorActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCostOfJob\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resourceProviderRequiredActiveEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LilypadPayment__InsufficientActiveBurnTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__InvalidResultStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__InvalidValidationResultStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__RoleAlreadyAssigned\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__RoleNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroJobCreatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroPayeeAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroPayoutAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroResourceProviderAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroSlashAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroStorageAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroTokenAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroTokenomicsAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroTreasuryWallet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroUserAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroValidationPoolWallet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroValueBasedRewardsWallet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__ZeroWithdrawalAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__amountMustBeGreaterThanZero\",\"inputs\":[{\"name\":\"functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LilypadPayment__escrowNotWithdrawable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__escrowSlashAmountTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__insufficientActiveEscrowAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__insufficientEscrowAmount\",\"inputs\":[{\"name\":\"escrowAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LilypadPayment__insufficientEscrowBalanceForWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__minimumResourceProviderAndValidatorDepositAmountNotMet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__transferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LilypadPayment__unauthorizedWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// LilypadPaymentEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadPaymentEngineMetaData.ABI instead.
var LilypadPaymentEngineABI = LilypadPaymentEngineMetaData.ABI

// LilypadPaymentEngine is an auto generated Go binding around an Ethereum contract.
type LilypadPaymentEngine struct {
	LilypadPaymentEngineCaller     // Read-only binding to the contract
	LilypadPaymentEngineTransactor // Write-only binding to the contract
	LilypadPaymentEngineFilterer   // Log filterer for contract events
}

// LilypadPaymentEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadPaymentEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadPaymentEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadPaymentEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadPaymentEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadPaymentEngineSession struct {
	Contract     *LilypadPaymentEngine // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LilypadPaymentEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadPaymentEngineCallerSession struct {
	Contract *LilypadPaymentEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LilypadPaymentEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadPaymentEngineTransactorSession struct {
	Contract     *LilypadPaymentEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LilypadPaymentEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadPaymentEngineRaw struct {
	Contract *LilypadPaymentEngine // Generic contract binding to access the raw methods on
}

// LilypadPaymentEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadPaymentEngineCallerRaw struct {
	Contract *LilypadPaymentEngineCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadPaymentEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadPaymentEngineTransactorRaw struct {
	Contract *LilypadPaymentEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypadPaymentEngine creates a new instance of LilypadPaymentEngine, bound to a specific deployed contract.
func NewLilypadPaymentEngine(address common.Address, backend bind.ContractBackend) (*LilypadPaymentEngine, error) {
	contract, err := bindLilypadPaymentEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngine{LilypadPaymentEngineCaller: LilypadPaymentEngineCaller{contract: contract}, LilypadPaymentEngineTransactor: LilypadPaymentEngineTransactor{contract: contract}, LilypadPaymentEngineFilterer: LilypadPaymentEngineFilterer{contract: contract}}, nil
}

// NewLilypadPaymentEngineCaller creates a new read-only instance of LilypadPaymentEngine, bound to a specific deployed contract.
func NewLilypadPaymentEngineCaller(address common.Address, caller bind.ContractCaller) (*LilypadPaymentEngineCaller, error) {
	contract, err := bindLilypadPaymentEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineCaller{contract: contract}, nil
}

// NewLilypadPaymentEngineTransactor creates a new write-only instance of LilypadPaymentEngine, bound to a specific deployed contract.
func NewLilypadPaymentEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadPaymentEngineTransactor, error) {
	contract, err := bindLilypadPaymentEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineTransactor{contract: contract}, nil
}

// NewLilypadPaymentEngineFilterer creates a new log filterer instance of LilypadPaymentEngine, bound to a specific deployed contract.
func NewLilypadPaymentEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadPaymentEngineFilterer, error) {
	contract, err := bindLilypadPaymentEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineFilterer{contract: contract}, nil
}

// bindLilypadPaymentEngine binds a generic wrapper to an already deployed contract.
func bindLilypadPaymentEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadPaymentEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadPaymentEngine *LilypadPaymentEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadPaymentEngine.Contract.LilypadPaymentEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadPaymentEngine *LilypadPaymentEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.LilypadPaymentEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadPaymentEngine *LilypadPaymentEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.LilypadPaymentEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LilypadPaymentEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALLOCKDURATION is a free data retrieval call binding the contract method 0x48448e7e.
//
// Solidity: function COLLATERAL_LOCK_DURATION() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) COLLATERALLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "COLLATERAL_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COLLATERALLOCKDURATION is a free data retrieval call binding the contract method 0x48448e7e.
//
// Solidity: function COLLATERAL_LOCK_DURATION() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) COLLATERALLOCKDURATION() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.COLLATERALLOCKDURATION(&_LilypadPaymentEngine.CallOpts)
}

// COLLATERALLOCKDURATION is a free data retrieval call binding the contract method 0x48448e7e.
//
// Solidity: function COLLATERAL_LOCK_DURATION() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) COLLATERALLOCKDURATION() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.COLLATERALLOCKDURATION(&_LilypadPaymentEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadPaymentEngine.Contract.DEFAULTADMINROLE(&_LilypadPaymentEngine.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LilypadPaymentEngine.Contract.DEFAULTADMINROLE(&_LilypadPaymentEngine.CallOpts)
}

// MINRESOURCEPROVIDERDEPOSITAMOUNT is a free data retrieval call binding the contract method 0xe9b54422.
//
// Solidity: function MIN_RESOURCE_PROVIDER_DEPOSIT_AMOUNT() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) MINRESOURCEPROVIDERDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "MIN_RESOURCE_PROVIDER_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINRESOURCEPROVIDERDEPOSITAMOUNT is a free data retrieval call binding the contract method 0xe9b54422.
//
// Solidity: function MIN_RESOURCE_PROVIDER_DEPOSIT_AMOUNT() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) MINRESOURCEPROVIDERDEPOSITAMOUNT() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.MINRESOURCEPROVIDERDEPOSITAMOUNT(&_LilypadPaymentEngine.CallOpts)
}

// MINRESOURCEPROVIDERDEPOSITAMOUNT is a free data retrieval call binding the contract method 0xe9b54422.
//
// Solidity: function MIN_RESOURCE_PROVIDER_DEPOSIT_AMOUNT() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) MINRESOURCEPROVIDERDEPOSITAMOUNT() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.MINRESOURCEPROVIDERDEPOSITAMOUNT(&_LilypadPaymentEngine.CallOpts)
}

// ActiveBurnTokens is a free data retrieval call binding the contract method 0xfca98805.
//
// Solidity: function activeBurnTokens() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) ActiveBurnTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "activeBurnTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveBurnTokens is a free data retrieval call binding the contract method 0xfca98805.
//
// Solidity: function activeBurnTokens() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) ActiveBurnTokens() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.ActiveBurnTokens(&_LilypadPaymentEngine.CallOpts)
}

// ActiveBurnTokens is a free data retrieval call binding the contract method 0xfca98805.
//
// Solidity: function activeBurnTokens() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) ActiveBurnTokens() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.ActiveBurnTokens(&_LilypadPaymentEngine.CallOpts)
}

// ActiveEscrow is a free data retrieval call binding the contract method 0x7eb7ee66.
//
// Solidity: function activeEscrow(address account) view returns(uint256 activeEscrow)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) ActiveEscrow(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "activeEscrow", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveEscrow is a free data retrieval call binding the contract method 0x7eb7ee66.
//
// Solidity: function activeEscrow(address account) view returns(uint256 activeEscrow)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) ActiveEscrow(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.ActiveEscrow(&_LilypadPaymentEngine.CallOpts, account)
}

// ActiveEscrow is a free data retrieval call binding the contract method 0x7eb7ee66.
//
// Solidity: function activeEscrow(address account) view returns(uint256 activeEscrow)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) ActiveEscrow(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.ActiveEscrow(&_LilypadPaymentEngine.CallOpts, account)
}

// CanWithdrawEscrow is a free data retrieval call binding the contract method 0x71f647a0.
//
// Solidity: function canWithdrawEscrow(address _address) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) CanWithdrawEscrow(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "canWithdrawEscrow", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanWithdrawEscrow is a free data retrieval call binding the contract method 0x71f647a0.
//
// Solidity: function canWithdrawEscrow(address _address) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) CanWithdrawEscrow(_address common.Address) (bool, error) {
	return _LilypadPaymentEngine.Contract.CanWithdrawEscrow(&_LilypadPaymentEngine.CallOpts, _address)
}

// CanWithdrawEscrow is a free data retrieval call binding the contract method 0x71f647a0.
//
// Solidity: function canWithdrawEscrow(address _address) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) CanWithdrawEscrow(_address common.Address) (bool, error) {
	return _LilypadPaymentEngine.Contract.CanWithdrawEscrow(&_LilypadPaymentEngine.CallOpts, _address)
}

// DepositTimestamps is a free data retrieval call binding the contract method 0x240be944.
//
// Solidity: function depositTimestamps(address account) view returns(uint256 depositTimestamp)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) DepositTimestamps(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "depositTimestamps", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositTimestamps is a free data retrieval call binding the contract method 0x240be944.
//
// Solidity: function depositTimestamps(address account) view returns(uint256 depositTimestamp)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) DepositTimestamps(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.DepositTimestamps(&_LilypadPaymentEngine.CallOpts, account)
}

// DepositTimestamps is a free data retrieval call binding the contract method 0x240be944.
//
// Solidity: function depositTimestamps(address account) view returns(uint256 depositTimestamp)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) DepositTimestamps(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.DepositTimestamps(&_LilypadPaymentEngine.CallOpts, account)
}

// EscrowBalances is a free data retrieval call binding the contract method 0x9672ba57.
//
// Solidity: function escrowBalances(address account) view returns(uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) EscrowBalances(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "escrowBalances", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalances is a free data retrieval call binding the contract method 0x9672ba57.
//
// Solidity: function escrowBalances(address account) view returns(uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) EscrowBalances(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.EscrowBalances(&_LilypadPaymentEngine.CallOpts, account)
}

// EscrowBalances is a free data retrieval call binding the contract method 0x9672ba57.
//
// Solidity: function escrowBalances(address account) view returns(uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) EscrowBalances(account common.Address) (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.EscrowBalances(&_LilypadPaymentEngine.CallOpts, account)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadPaymentEngine.Contract.GetRoleAdmin(&_LilypadPaymentEngine.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LilypadPaymentEngine.Contract.GetRoleAdmin(&_LilypadPaymentEngine.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadPaymentEngine.Contract.HasRole(&_LilypadPaymentEngine.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LilypadPaymentEngine.Contract.HasRole(&_LilypadPaymentEngine.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadPaymentEngine.Contract.SupportsInterface(&_LilypadPaymentEngine.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LilypadPaymentEngine.Contract.SupportsInterface(&_LilypadPaymentEngine.CallOpts, interfaceId)
}

// TotalActiveEscrow is a free data retrieval call binding the contract method 0x8c848a73.
//
// Solidity: function totalActiveEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) TotalActiveEscrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "totalActiveEscrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveEscrow is a free data retrieval call binding the contract method 0x8c848a73.
//
// Solidity: function totalActiveEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) TotalActiveEscrow() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.TotalActiveEscrow(&_LilypadPaymentEngine.CallOpts)
}

// TotalActiveEscrow is a free data retrieval call binding the contract method 0x8c848a73.
//
// Solidity: function totalActiveEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) TotalActiveEscrow() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.TotalActiveEscrow(&_LilypadPaymentEngine.CallOpts)
}

// TotalEscrow is a free data retrieval call binding the contract method 0xa3d89844.
//
// Solidity: function totalEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) TotalEscrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "totalEscrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEscrow is a free data retrieval call binding the contract method 0xa3d89844.
//
// Solidity: function totalEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) TotalEscrow() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.TotalEscrow(&_LilypadPaymentEngine.CallOpts)
}

// TotalEscrow is a free data retrieval call binding the contract method 0xa3d89844.
//
// Solidity: function totalEscrow() view returns(uint256)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) TotalEscrow() (*big.Int, error) {
	return _LilypadPaymentEngine.Contract.TotalEscrow(&_LilypadPaymentEngine.CallOpts)
}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) TreasuryWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "treasuryWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) TreasuryWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.TreasuryWallet(&_LilypadPaymentEngine.CallOpts)
}

// TreasuryWallet is a free data retrieval call binding the contract method 0x4626402b.
//
// Solidity: function treasuryWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) TreasuryWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.TreasuryWallet(&_LilypadPaymentEngine.CallOpts)
}

// ValidationPoolWallet is a free data retrieval call binding the contract method 0xf7c3c341.
//
// Solidity: function validationPoolWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) ValidationPoolWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "validationPoolWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidationPoolWallet is a free data retrieval call binding the contract method 0xf7c3c341.
//
// Solidity: function validationPoolWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) ValidationPoolWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.ValidationPoolWallet(&_LilypadPaymentEngine.CallOpts)
}

// ValidationPoolWallet is a free data retrieval call binding the contract method 0xf7c3c341.
//
// Solidity: function validationPoolWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) ValidationPoolWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.ValidationPoolWallet(&_LilypadPaymentEngine.CallOpts)
}

// ValueBasedRewardsWallet is a free data retrieval call binding the contract method 0xfa7cd8bf.
//
// Solidity: function valueBasedRewardsWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) ValueBasedRewardsWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "valueBasedRewardsWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValueBasedRewardsWallet is a free data retrieval call binding the contract method 0xfa7cd8bf.
//
// Solidity: function valueBasedRewardsWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) ValueBasedRewardsWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.ValueBasedRewardsWallet(&_LilypadPaymentEngine.CallOpts)
}

// ValueBasedRewardsWallet is a free data retrieval call binding the contract method 0xfa7cd8bf.
//
// Solidity: function valueBasedRewardsWallet() view returns(address)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) ValueBasedRewardsWallet() (common.Address, error) {
	return _LilypadPaymentEngine.Contract.ValueBasedRewardsWallet(&_LilypadPaymentEngine.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadPaymentEngine *LilypadPaymentEngineCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LilypadPaymentEngine.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) Version() (string, error) {
	return _LilypadPaymentEngine.Contract.Version(&_LilypadPaymentEngine.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LilypadPaymentEngine *LilypadPaymentEngineCallerSession) Version() (string, error) {
	return _LilypadPaymentEngine.Contract.Version(&_LilypadPaymentEngine.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.GrantRole(&_LilypadPaymentEngine.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.GrantRole(&_LilypadPaymentEngine.TransactOpts, role, account)
}

// HandleJobCompletion is a paid mutator transaction binding the contract method 0x57ff4c40.
//
// Solidity: function handleJobCompletion((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) HandleJobCompletion(opts *bind.TransactOpts, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "handleJobCompletion", result)
}

// HandleJobCompletion is a paid mutator transaction binding the contract method 0x57ff4c40.
//
// Solidity: function handleJobCompletion((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) HandleJobCompletion(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleJobCompletion(&_LilypadPaymentEngine.TransactOpts, result)
}

// HandleJobCompletion is a paid mutator transaction binding the contract method 0x57ff4c40.
//
// Solidity: function handleJobCompletion((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) HandleJobCompletion(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleJobCompletion(&_LilypadPaymentEngine.TransactOpts, result)
}

// HandleJobFailure is a paid mutator transaction binding the contract method 0x90faf0ca.
//
// Solidity: function handleJobFailure((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) HandleJobFailure(opts *bind.TransactOpts, result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "handleJobFailure", result)
}

// HandleJobFailure is a paid mutator transaction binding the contract method 0x90faf0ca.
//
// Solidity: function handleJobFailure((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) HandleJobFailure(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleJobFailure(&_LilypadPaymentEngine.TransactOpts, result)
}

// HandleJobFailure is a paid mutator transaction binding the contract method 0x90faf0ca.
//
// Solidity: function handleJobFailure((string,string,string,uint8,uint256) result) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) HandleJobFailure(result SharedStructsResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleJobFailure(&_LilypadPaymentEngine.TransactOpts, result)
}

// HandleValidationFailed is a paid mutator transaction binding the contract method 0x60b5e939.
//
// Solidity: function handleValidationFailed((string,string,string,uint8,uint256,address) _validationResult, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) _originalJobDeal) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) HandleValidationFailed(opts *bind.TransactOpts, _validationResult SharedStructsValidationResult, _originalJobDeal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "handleValidationFailed", _validationResult, _originalJobDeal)
}

// HandleValidationFailed is a paid mutator transaction binding the contract method 0x60b5e939.
//
// Solidity: function handleValidationFailed((string,string,string,uint8,uint256,address) _validationResult, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) _originalJobDeal) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) HandleValidationFailed(_validationResult SharedStructsValidationResult, _originalJobDeal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleValidationFailed(&_LilypadPaymentEngine.TransactOpts, _validationResult, _originalJobDeal)
}

// HandleValidationFailed is a paid mutator transaction binding the contract method 0x60b5e939.
//
// Solidity: function handleValidationFailed((string,string,string,uint8,uint256,address) _validationResult, (string,address,address,address,address,string,string,uint8,uint256,(uint256,uint256,uint256,uint256,uint256)) _originalJobDeal) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) HandleValidationFailed(_validationResult SharedStructsValidationResult, _originalJobDeal SharedStructsDeal) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleValidationFailed(&_LilypadPaymentEngine.TransactOpts, _validationResult, _originalJobDeal)
}

// HandleValidationPassed is a paid mutator transaction binding the contract method 0x46e49520.
//
// Solidity: function handleValidationPassed((string,string,string,uint8,uint256,address) _validationResult) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) HandleValidationPassed(opts *bind.TransactOpts, _validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "handleValidationPassed", _validationResult)
}

// HandleValidationPassed is a paid mutator transaction binding the contract method 0x46e49520.
//
// Solidity: function handleValidationPassed((string,string,string,uint8,uint256,address) _validationResult) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) HandleValidationPassed(_validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleValidationPassed(&_LilypadPaymentEngine.TransactOpts, _validationResult)
}

// HandleValidationPassed is a paid mutator transaction binding the contract method 0x46e49520.
//
// Solidity: function handleValidationPassed((string,string,string,uint8,uint256,address) _validationResult) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) HandleValidationPassed(_validationResult SharedStructsValidationResult) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.HandleValidationPassed(&_LilypadPaymentEngine.TransactOpts, _validationResult)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _l2token, address _lilypadStorageAddress, address _lilypadUserAddress, address _lilypadTokenomicsAddress, address _treasuryWallet, address _valueBasedRewardsWallet, address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) Initialize(opts *bind.TransactOpts, _l2token common.Address, _lilypadStorageAddress common.Address, _lilypadUserAddress common.Address, _lilypadTokenomicsAddress common.Address, _treasuryWallet common.Address, _valueBasedRewardsWallet common.Address, _validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "initialize", _l2token, _lilypadStorageAddress, _lilypadUserAddress, _lilypadTokenomicsAddress, _treasuryWallet, _valueBasedRewardsWallet, _validationPoolWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _l2token, address _lilypadStorageAddress, address _lilypadUserAddress, address _lilypadTokenomicsAddress, address _treasuryWallet, address _valueBasedRewardsWallet, address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) Initialize(_l2token common.Address, _lilypadStorageAddress common.Address, _lilypadUserAddress common.Address, _lilypadTokenomicsAddress common.Address, _treasuryWallet common.Address, _valueBasedRewardsWallet common.Address, _validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.Initialize(&_LilypadPaymentEngine.TransactOpts, _l2token, _lilypadStorageAddress, _lilypadUserAddress, _lilypadTokenomicsAddress, _treasuryWallet, _valueBasedRewardsWallet, _validationPoolWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _l2token, address _lilypadStorageAddress, address _lilypadUserAddress, address _lilypadTokenomicsAddress, address _treasuryWallet, address _valueBasedRewardsWallet, address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) Initialize(_l2token common.Address, _lilypadStorageAddress common.Address, _lilypadUserAddress common.Address, _lilypadTokenomicsAddress common.Address, _treasuryWallet common.Address, _valueBasedRewardsWallet common.Address, _validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.Initialize(&_LilypadPaymentEngine.TransactOpts, _l2token, _lilypadStorageAddress, _lilypadUserAddress, _lilypadTokenomicsAddress, _treasuryWallet, _valueBasedRewardsWallet, _validationPoolWallet)
}

// InitiateLockupOfEscrowForJob is a paid mutator transaction binding the contract method 0x737414db.
//
// Solidity: function initiateLockupOfEscrowForJob(address jobCreator, address resourceProvider, string dealId, uint256 cost, uint256 resourceProviderCollateralLockupAmount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) InitiateLockupOfEscrowForJob(opts *bind.TransactOpts, jobCreator common.Address, resourceProvider common.Address, dealId string, cost *big.Int, resourceProviderCollateralLockupAmount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "initiateLockupOfEscrowForJob", jobCreator, resourceProvider, dealId, cost, resourceProviderCollateralLockupAmount)
}

// InitiateLockupOfEscrowForJob is a paid mutator transaction binding the contract method 0x737414db.
//
// Solidity: function initiateLockupOfEscrowForJob(address jobCreator, address resourceProvider, string dealId, uint256 cost, uint256 resourceProviderCollateralLockupAmount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) InitiateLockupOfEscrowForJob(jobCreator common.Address, resourceProvider common.Address, dealId string, cost *big.Int, resourceProviderCollateralLockupAmount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.InitiateLockupOfEscrowForJob(&_LilypadPaymentEngine.TransactOpts, jobCreator, resourceProvider, dealId, cost, resourceProviderCollateralLockupAmount)
}

// InitiateLockupOfEscrowForJob is a paid mutator transaction binding the contract method 0x737414db.
//
// Solidity: function initiateLockupOfEscrowForJob(address jobCreator, address resourceProvider, string dealId, uint256 cost, uint256 resourceProviderCollateralLockupAmount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) InitiateLockupOfEscrowForJob(jobCreator common.Address, resourceProvider common.Address, dealId string, cost *big.Int, resourceProviderCollateralLockupAmount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.InitiateLockupOfEscrowForJob(&_LilypadPaymentEngine.TransactOpts, jobCreator, resourceProvider, dealId, cost, resourceProviderCollateralLockupAmount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0xb2a93a2d.
//
// Solidity: function payEscrow(address _payee, uint8 _paymentReason, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) PayEscrow(opts *bind.TransactOpts, _payee common.Address, _paymentReason uint8, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "payEscrow", _payee, _paymentReason, _amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0xb2a93a2d.
//
// Solidity: function payEscrow(address _payee, uint8 _paymentReason, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) PayEscrow(_payee common.Address, _paymentReason uint8, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.PayEscrow(&_LilypadPaymentEngine.TransactOpts, _payee, _paymentReason, _amount)
}

// PayEscrow is a paid mutator transaction binding the contract method 0xb2a93a2d.
//
// Solidity: function payEscrow(address _payee, uint8 _paymentReason, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) PayEscrow(_payee common.Address, _paymentReason uint8, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.PayEscrow(&_LilypadPaymentEngine.TransactOpts, _payee, _paymentReason, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.RenounceRole(&_LilypadPaymentEngine.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.RenounceRole(&_LilypadPaymentEngine.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.RevokeRole(&_LilypadPaymentEngine.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.RevokeRole(&_LilypadPaymentEngine.TransactOpts, role, account)
}

// SetL2Token is a paid mutator transaction binding the contract method 0xfc883a82.
//
// Solidity: function setL2Token(address _l2tokenAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetL2Token(opts *bind.TransactOpts, _l2tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setL2Token", _l2tokenAddress)
}

// SetL2Token is a paid mutator transaction binding the contract method 0xfc883a82.
//
// Solidity: function setL2Token(address _l2tokenAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetL2Token(_l2tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetL2Token(&_LilypadPaymentEngine.TransactOpts, _l2tokenAddress)
}

// SetL2Token is a paid mutator transaction binding the contract method 0xfc883a82.
//
// Solidity: function setL2Token(address _l2tokenAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetL2Token(_l2tokenAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetL2Token(&_LilypadPaymentEngine.TransactOpts, _l2tokenAddress)
}

// SetLilypadStorage is a paid mutator transaction binding the contract method 0x393f7fe9.
//
// Solidity: function setLilypadStorage(address _lilypadStorageAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetLilypadStorage(opts *bind.TransactOpts, _lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setLilypadStorage", _lilypadStorageAddress)
}

// SetLilypadStorage is a paid mutator transaction binding the contract method 0x393f7fe9.
//
// Solidity: function setLilypadStorage(address _lilypadStorageAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetLilypadStorage(_lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadStorage(&_LilypadPaymentEngine.TransactOpts, _lilypadStorageAddress)
}

// SetLilypadStorage is a paid mutator transaction binding the contract method 0x393f7fe9.
//
// Solidity: function setLilypadStorage(address _lilypadStorageAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetLilypadStorage(_lilypadStorageAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadStorage(&_LilypadPaymentEngine.TransactOpts, _lilypadStorageAddress)
}

// SetLilypadTokenomics is a paid mutator transaction binding the contract method 0x045256f5.
//
// Solidity: function setLilypadTokenomics(address _lilypadTokenomicsAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetLilypadTokenomics(opts *bind.TransactOpts, _lilypadTokenomicsAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setLilypadTokenomics", _lilypadTokenomicsAddress)
}

// SetLilypadTokenomics is a paid mutator transaction binding the contract method 0x045256f5.
//
// Solidity: function setLilypadTokenomics(address _lilypadTokenomicsAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetLilypadTokenomics(_lilypadTokenomicsAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadTokenomics(&_LilypadPaymentEngine.TransactOpts, _lilypadTokenomicsAddress)
}

// SetLilypadTokenomics is a paid mutator transaction binding the contract method 0x045256f5.
//
// Solidity: function setLilypadTokenomics(address _lilypadTokenomicsAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetLilypadTokenomics(_lilypadTokenomicsAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadTokenomics(&_LilypadPaymentEngine.TransactOpts, _lilypadTokenomicsAddress)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUserAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetLilypadUser(opts *bind.TransactOpts, _lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setLilypadUser", _lilypadUserAddress)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUserAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetLilypadUser(_lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadUser(&_LilypadPaymentEngine.TransactOpts, _lilypadUserAddress)
}

// SetLilypadUser is a paid mutator transaction binding the contract method 0xf30e1383.
//
// Solidity: function setLilypadUser(address _lilypadUserAddress) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetLilypadUser(_lilypadUserAddress common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetLilypadUser(&_LilypadPaymentEngine.TransactOpts, _lilypadUserAddress)
}

// SetTreasuryWallet is a paid mutator transaction binding the contract method 0xa8602fea.
//
// Solidity: function setTreasuryWallet(address _treasuryWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetTreasuryWallet(opts *bind.TransactOpts, _treasuryWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setTreasuryWallet", _treasuryWallet)
}

// SetTreasuryWallet is a paid mutator transaction binding the contract method 0xa8602fea.
//
// Solidity: function setTreasuryWallet(address _treasuryWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetTreasuryWallet(_treasuryWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetTreasuryWallet(&_LilypadPaymentEngine.TransactOpts, _treasuryWallet)
}

// SetTreasuryWallet is a paid mutator transaction binding the contract method 0xa8602fea.
//
// Solidity: function setTreasuryWallet(address _treasuryWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetTreasuryWallet(_treasuryWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetTreasuryWallet(&_LilypadPaymentEngine.TransactOpts, _treasuryWallet)
}

// SetValidationPoolWallet is a paid mutator transaction binding the contract method 0x9ec72297.
//
// Solidity: function setValidationPoolWallet(address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetValidationPoolWallet(opts *bind.TransactOpts, _validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setValidationPoolWallet", _validationPoolWallet)
}

// SetValidationPoolWallet is a paid mutator transaction binding the contract method 0x9ec72297.
//
// Solidity: function setValidationPoolWallet(address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetValidationPoolWallet(_validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetValidationPoolWallet(&_LilypadPaymentEngine.TransactOpts, _validationPoolWallet)
}

// SetValidationPoolWallet is a paid mutator transaction binding the contract method 0x9ec72297.
//
// Solidity: function setValidationPoolWallet(address _validationPoolWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetValidationPoolWallet(_validationPoolWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetValidationPoolWallet(&_LilypadPaymentEngine.TransactOpts, _validationPoolWallet)
}

// SetValueBasedRewardsWallet is a paid mutator transaction binding the contract method 0x2d039603.
//
// Solidity: function setValueBasedRewardsWallet(address _valueBasedRewardsWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) SetValueBasedRewardsWallet(opts *bind.TransactOpts, _valueBasedRewardsWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "setValueBasedRewardsWallet", _valueBasedRewardsWallet)
}

// SetValueBasedRewardsWallet is a paid mutator transaction binding the contract method 0x2d039603.
//
// Solidity: function setValueBasedRewardsWallet(address _valueBasedRewardsWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) SetValueBasedRewardsWallet(_valueBasedRewardsWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetValueBasedRewardsWallet(&_LilypadPaymentEngine.TransactOpts, _valueBasedRewardsWallet)
}

// SetValueBasedRewardsWallet is a paid mutator transaction binding the contract method 0x2d039603.
//
// Solidity: function setValueBasedRewardsWallet(address _valueBasedRewardsWallet) returns()
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) SetValueBasedRewardsWallet(_valueBasedRewardsWallet common.Address) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.SetValueBasedRewardsWallet(&_LilypadPaymentEngine.TransactOpts, _valueBasedRewardsWallet)
}

// UpdateActiveBurnTokens is a paid mutator transaction binding the contract method 0x19665bfe.
//
// Solidity: function updateActiveBurnTokens(uint256 _amountBurnt) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) UpdateActiveBurnTokens(opts *bind.TransactOpts, _amountBurnt *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "updateActiveBurnTokens", _amountBurnt)
}

// UpdateActiveBurnTokens is a paid mutator transaction binding the contract method 0x19665bfe.
//
// Solidity: function updateActiveBurnTokens(uint256 _amountBurnt) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) UpdateActiveBurnTokens(_amountBurnt *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.UpdateActiveBurnTokens(&_LilypadPaymentEngine.TransactOpts, _amountBurnt)
}

// UpdateActiveBurnTokens is a paid mutator transaction binding the contract method 0x19665bfe.
//
// Solidity: function updateActiveBurnTokens(uint256 _amountBurnt) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) UpdateActiveBurnTokens(_amountBurnt *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.UpdateActiveBurnTokens(&_LilypadPaymentEngine.TransactOpts, _amountBurnt)
}

// WithdrawEscrow is a paid mutator transaction binding the contract method 0x3946b978.
//
// Solidity: function withdrawEscrow(address _withdrawer, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactor) WithdrawEscrow(opts *bind.TransactOpts, _withdrawer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.contract.Transact(opts, "withdrawEscrow", _withdrawer, _amount)
}

// WithdrawEscrow is a paid mutator transaction binding the contract method 0x3946b978.
//
// Solidity: function withdrawEscrow(address _withdrawer, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineSession) WithdrawEscrow(_withdrawer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.WithdrawEscrow(&_LilypadPaymentEngine.TransactOpts, _withdrawer, _amount)
}

// WithdrawEscrow is a paid mutator transaction binding the contract method 0x3946b978.
//
// Solidity: function withdrawEscrow(address _withdrawer, uint256 _amount) returns(bool)
func (_LilypadPaymentEngine *LilypadPaymentEngineTransactorSession) WithdrawEscrow(_withdrawer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LilypadPaymentEngine.Contract.WithdrawEscrow(&_LilypadPaymentEngine.TransactOpts, _withdrawer, _amount)
}

// LilypadPaymentEngineInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineInitializedIterator struct {
	Event *LilypadPaymentEngineInitialized // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineInitialized)
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
		it.Event = new(LilypadPaymentEngineInitialized)
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
func (it *LilypadPaymentEngineInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineInitialized represents a Initialized event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterInitialized(opts *bind.FilterOpts) (*LilypadPaymentEngineInitializedIterator, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineInitializedIterator{contract: _LilypadPaymentEngine.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineInitialized) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineInitialized)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseInitialized(log types.Log) (*LilypadPaymentEngineInitialized, error) {
	event := new(LilypadPaymentEngineInitialized)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator is returned from FilterLilypadPaymentActiveEscrowLockedForJob and is used to iterate over the raw logs and unpacked data for LilypadPaymentActiveEscrowLockedForJob events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob)
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
func (it *LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob represents a LilypadPaymentActiveEscrowLockedForJob event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob struct {
	JobCreator       common.Address
	ResourceProvider common.Address
	DealId           common.Hash
	Cost             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentActiveEscrowLockedForJob is a free log retrieval operation binding the contract event 0x10fed0cb4234c2a2a041280f5077a798af26fd9bdf2243900568ed2f8626d0a3.
//
// Solidity: event LilypadPayment__ActiveEscrowLockedForJob(address indexed jobCreator, address indexed resourceProvider, string indexed dealId, uint256 cost)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentActiveEscrowLockedForJob(opts *bind.FilterOpts, jobCreator []common.Address, resourceProvider []common.Address, dealId []string) (*LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ActiveEscrowLockedForJob", jobCreatorRule, resourceProviderRule, dealIdRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJobIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ActiveEscrowLockedForJob", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentActiveEscrowLockedForJob is a free log subscription operation binding the contract event 0x10fed0cb4234c2a2a041280f5077a798af26fd9bdf2243900568ed2f8626d0a3.
//
// Solidity: event LilypadPayment__ActiveEscrowLockedForJob(address indexed jobCreator, address indexed resourceProvider, string indexed dealId, uint256 cost)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentActiveEscrowLockedForJob(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob, jobCreator []common.Address, resourceProvider []common.Address, dealId []string) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ActiveEscrowLockedForJob", jobCreatorRule, resourceProviderRule, dealIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ActiveEscrowLockedForJob", log); err != nil {
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

// ParseLilypadPaymentActiveEscrowLockedForJob is a log parse operation binding the contract event 0x10fed0cb4234c2a2a041280f5077a798af26fd9bdf2243900568ed2f8626d0a3.
//
// Solidity: event LilypadPayment__ActiveEscrowLockedForJob(address indexed jobCreator, address indexed resourceProvider, string indexed dealId, uint256 cost)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentActiveEscrowLockedForJob(log types.Log) (*LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob, error) {
	event := new(LilypadPaymentEngineLilypadPaymentActiveEscrowLockedForJob)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ActiveEscrowLockedForJob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator is returned from FilterLilypadPaymentControllerRoleGranted and is used to iterate over the raw logs and unpacked data for LilypadPaymentControllerRoleGranted events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentControllerRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentControllerRoleGranted)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentControllerRoleGranted)
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
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentControllerRoleGranted represents a LilypadPaymentControllerRoleGranted event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentControllerRoleGranted struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentControllerRoleGranted is a free log retrieval operation binding the contract event 0x5650aa8ebdf2e949b482b8da38e6030fa52f2bed02d86a4a21b333b8e7a2e660.
//
// Solidity: event LilypadPayment__ControllerRoleGranted(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentControllerRoleGranted(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ControllerRoleGranted", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentControllerRoleGrantedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ControllerRoleGranted", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentControllerRoleGranted is a free log subscription operation binding the contract event 0x5650aa8ebdf2e949b482b8da38e6030fa52f2bed02d86a4a21b333b8e7a2e660.
//
// Solidity: event LilypadPayment__ControllerRoleGranted(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentControllerRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentControllerRoleGranted, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ControllerRoleGranted", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentControllerRoleGranted)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ControllerRoleGranted", log); err != nil {
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

// ParseLilypadPaymentControllerRoleGranted is a log parse operation binding the contract event 0x5650aa8ebdf2e949b482b8da38e6030fa52f2bed02d86a4a21b333b8e7a2e660.
//
// Solidity: event LilypadPayment__ControllerRoleGranted(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentControllerRoleGranted(log types.Log) (*LilypadPaymentEngineLilypadPaymentControllerRoleGranted, error) {
	event := new(LilypadPaymentEngineLilypadPaymentControllerRoleGranted)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ControllerRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator is returned from FilterLilypadPaymentControllerRoleRevoked and is used to iterate over the raw logs and unpacked data for LilypadPaymentControllerRoleRevoked events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentControllerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentControllerRoleRevoked)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentControllerRoleRevoked)
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
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentControllerRoleRevoked represents a LilypadPaymentControllerRoleRevoked event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentControllerRoleRevoked struct {
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentControllerRoleRevoked is a free log retrieval operation binding the contract event 0xe52618db293e6a02cdcd452821ab48bddc960db0d0e6410e5bf5c0d8b9d950b9.
//
// Solidity: event LilypadPayment__ControllerRoleRevoked(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentControllerRoleRevoked(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ControllerRoleRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentControllerRoleRevokedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ControllerRoleRevoked", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentControllerRoleRevoked is a free log subscription operation binding the contract event 0xe52618db293e6a02cdcd452821ab48bddc960db0d0e6410e5bf5c0d8b9d950b9.
//
// Solidity: event LilypadPayment__ControllerRoleRevoked(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentControllerRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentControllerRoleRevoked, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ControllerRoleRevoked", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentControllerRoleRevoked)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ControllerRoleRevoked", log); err != nil {
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

// ParseLilypadPaymentControllerRoleRevoked is a log parse operation binding the contract event 0xe52618db293e6a02cdcd452821ab48bddc960db0d0e6410e5bf5c0d8b9d950b9.
//
// Solidity: event LilypadPayment__ControllerRoleRevoked(address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentControllerRoleRevoked(log types.Log) (*LilypadPaymentEngineLilypadPaymentControllerRoleRevoked, error) {
	event := new(LilypadPaymentEngineLilypadPaymentControllerRoleRevoked)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ControllerRoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentJobCompletedIterator is returned from FilterLilypadPaymentJobCompleted and is used to iterate over the raw logs and unpacked data for LilypadPaymentJobCompleted events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentJobCompletedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentJobCompleted // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentJobCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentJobCompleted)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentJobCompleted)
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
func (it *LilypadPaymentEngineLilypadPaymentJobCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentJobCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentJobCompleted represents a LilypadPaymentJobCompleted event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentJobCompleted struct {
	JobCreator       common.Address
	ResourceProvider common.Address
	DealId           string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentJobCompleted is a free log retrieval operation binding the contract event 0x26a52488c42dc8f6e51d5530e4f72539eeab0574e3073beac42c0db339123f0a.
//
// Solidity: event LilypadPayment__JobCompleted(address indexed jobCreator, address indexed resourceProvider, string dealId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentJobCompleted(opts *bind.FilterOpts, jobCreator []common.Address, resourceProvider []common.Address) (*LilypadPaymentEngineLilypadPaymentJobCompletedIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__JobCompleted", jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentJobCompletedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__JobCompleted", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentJobCompleted is a free log subscription operation binding the contract event 0x26a52488c42dc8f6e51d5530e4f72539eeab0574e3073beac42c0db339123f0a.
//
// Solidity: event LilypadPayment__JobCompleted(address indexed jobCreator, address indexed resourceProvider, string dealId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentJobCompleted(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentJobCompleted, jobCreator []common.Address, resourceProvider []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__JobCompleted", jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentJobCompleted)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__JobCompleted", log); err != nil {
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

// ParseLilypadPaymentJobCompleted is a log parse operation binding the contract event 0x26a52488c42dc8f6e51d5530e4f72539eeab0574e3073beac42c0db339123f0a.
//
// Solidity: event LilypadPayment__JobCompleted(address indexed jobCreator, address indexed resourceProvider, string dealId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentJobCompleted(log types.Log) (*LilypadPaymentEngineLilypadPaymentJobCompleted, error) {
	event := new(LilypadPaymentEngineLilypadPaymentJobCompleted)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__JobCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentJobFailedIterator is returned from FilterLilypadPaymentJobFailed and is used to iterate over the raw logs and unpacked data for LilypadPaymentJobFailed events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentJobFailedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentJobFailed // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentJobFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentJobFailed)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentJobFailed)
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
func (it *LilypadPaymentEngineLilypadPaymentJobFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentJobFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentJobFailed represents a LilypadPaymentJobFailed event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentJobFailed struct {
	JobCreator       common.Address
	ResourceProvider common.Address
	ResultId         string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentJobFailed is a free log retrieval operation binding the contract event 0x7323ccc9f5da99e8e56b37ff9bbb92c99d50bf62504c9cbe5f076200fef7ceeb.
//
// Solidity: event LilypadPayment__JobFailed(address indexed jobCreator, address indexed resourceProvider, string resultId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentJobFailed(opts *bind.FilterOpts, jobCreator []common.Address, resourceProvider []common.Address) (*LilypadPaymentEngineLilypadPaymentJobFailedIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__JobFailed", jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentJobFailedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__JobFailed", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentJobFailed is a free log subscription operation binding the contract event 0x7323ccc9f5da99e8e56b37ff9bbb92c99d50bf62504c9cbe5f076200fef7ceeb.
//
// Solidity: event LilypadPayment__JobFailed(address indexed jobCreator, address indexed resourceProvider, string resultId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentJobFailed(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentJobFailed, jobCreator []common.Address, resourceProvider []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__JobFailed", jobCreatorRule, resourceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentJobFailed)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__JobFailed", log); err != nil {
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

// ParseLilypadPaymentJobFailed is a log parse operation binding the contract event 0x7323ccc9f5da99e8e56b37ff9bbb92c99d50bf62504c9cbe5f076200fef7ceeb.
//
// Solidity: event LilypadPayment__JobFailed(address indexed jobCreator, address indexed resourceProvider, string resultId)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentJobFailed(log types.Log) (*LilypadPaymentEngineLilypadPaymentJobFailed, error) {
	event := new(LilypadPaymentEngineLilypadPaymentJobFailed)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__JobFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentTokensBurnedIterator is returned from FilterLilypadPaymentTokensBurned and is used to iterate over the raw logs and unpacked data for LilypadPaymentTokensBurned events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTokensBurnedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentTokensBurned // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentTokensBurned)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentTokensBurned)
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
func (it *LilypadPaymentEngineLilypadPaymentTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentTokensBurned represents a LilypadPaymentTokensBurned event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTokensBurned struct {
	BlockNumber    *big.Int
	BlockTimestamp *big.Int
	AmountBurnt    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentTokensBurned is a free log retrieval operation binding the contract event 0xcea021bb9e5c2a7580571d7229b24a9992ee0ca565ec3cf5073269ec9c1af73f.
//
// Solidity: event LilypadPayment__TokensBurned(uint256 blockNumber, uint256 blockTimestamp, uint256 amountBurnt)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentTokensBurned(opts *bind.FilterOpts) (*LilypadPaymentEngineLilypadPaymentTokensBurnedIterator, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__TokensBurned")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentTokensBurnedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__TokensBurned", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentTokensBurned is a free log subscription operation binding the contract event 0xcea021bb9e5c2a7580571d7229b24a9992ee0ca565ec3cf5073269ec9c1af73f.
//
// Solidity: event LilypadPayment__TokensBurned(uint256 blockNumber, uint256 blockTimestamp, uint256 amountBurnt)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentTokensBurned(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentTokensBurned) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__TokensBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentTokensBurned)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TokensBurned", log); err != nil {
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

// ParseLilypadPaymentTokensBurned is a log parse operation binding the contract event 0xcea021bb9e5c2a7580571d7229b24a9992ee0ca565ec3cf5073269ec9c1af73f.
//
// Solidity: event LilypadPayment__TokensBurned(uint256 blockNumber, uint256 blockTimestamp, uint256 amountBurnt)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentTokensBurned(log types.Log) (*LilypadPaymentEngineLilypadPaymentTokensBurned, error) {
	event := new(LilypadPaymentEngineLilypadPaymentTokensBurned)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator is returned from FilterLilypadPaymentTotalFeesGeneratedByJob and is used to iterate over the raw logs and unpacked data for LilypadPaymentTotalFeesGeneratedByJob events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob)
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
func (it *LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob represents a LilypadPaymentTotalFeesGeneratedByJob event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob struct {
	ResourceProvider common.Address
	JobCreator       common.Address
	DealId           string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentTotalFeesGeneratedByJob is a free log retrieval operation binding the contract event 0xdaffa0fb3470125863041fdb4c4ca6c41001b9a63a54cea87c0fbe3931706bde.
//
// Solidity: event LilypadPayment__TotalFeesGeneratedByJob(address indexed resourceProvider, address indexed jobCreator, string dealId, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentTotalFeesGeneratedByJob(opts *bind.FilterOpts, resourceProvider []common.Address, jobCreator []common.Address) (*LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__TotalFeesGeneratedByJob", resourceProviderRule, jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJobIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__TotalFeesGeneratedByJob", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentTotalFeesGeneratedByJob is a free log subscription operation binding the contract event 0xdaffa0fb3470125863041fdb4c4ca6c41001b9a63a54cea87c0fbe3931706bde.
//
// Solidity: event LilypadPayment__TotalFeesGeneratedByJob(address indexed resourceProvider, address indexed jobCreator, string dealId, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentTotalFeesGeneratedByJob(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob, resourceProvider []common.Address, jobCreator []common.Address) (event.Subscription, error) {

	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__TotalFeesGeneratedByJob", resourceProviderRule, jobCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TotalFeesGeneratedByJob", log); err != nil {
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

// ParseLilypadPaymentTotalFeesGeneratedByJob is a log parse operation binding the contract event 0xdaffa0fb3470125863041fdb4c4ca6c41001b9a63a54cea87c0fbe3931706bde.
//
// Solidity: event LilypadPayment__TotalFeesGeneratedByJob(address indexed resourceProvider, address indexed jobCreator, string dealId, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentTotalFeesGeneratedByJob(log types.Log) (*LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob, error) {
	event := new(LilypadPaymentEngineLilypadPaymentTotalFeesGeneratedByJob)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TotalFeesGeneratedByJob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator is returned from FilterLilypadPaymentTreasuryWalletUpdated and is used to iterate over the raw logs and unpacked data for LilypadPaymentTreasuryWalletUpdated events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated)
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
func (it *LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated represents a LilypadPaymentTreasuryWalletUpdated event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated struct {
	NewTreasuryWallet common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentTreasuryWalletUpdated is a free log retrieval operation binding the contract event 0xac99775b354bce99af2298c1e7d58da60e1dfb9962c14149e21a03e781965212.
//
// Solidity: event LilypadPayment__TreasuryWalletUpdated(address newTreasuryWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentTreasuryWalletUpdated(opts *bind.FilterOpts) (*LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__TreasuryWalletUpdated")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdatedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__TreasuryWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentTreasuryWalletUpdated is a free log subscription operation binding the contract event 0xac99775b354bce99af2298c1e7d58da60e1dfb9962c14149e21a03e781965212.
//
// Solidity: event LilypadPayment__TreasuryWalletUpdated(address newTreasuryWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentTreasuryWalletUpdated(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__TreasuryWalletUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TreasuryWalletUpdated", log); err != nil {
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

// ParseLilypadPaymentTreasuryWalletUpdated is a log parse operation binding the contract event 0xac99775b354bce99af2298c1e7d58da60e1dfb9962c14149e21a03e781965212.
//
// Solidity: event LilypadPayment__TreasuryWalletUpdated(address newTreasuryWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentTreasuryWalletUpdated(log types.Log) (*LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated, error) {
	event := new(LilypadPaymentEngineLilypadPaymentTreasuryWalletUpdated)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__TreasuryWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentValidationFailedIterator is returned from FilterLilypadPaymentValidationFailed and is used to iterate over the raw logs and unpacked data for LilypadPaymentValidationFailed events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationFailedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentValidationFailed // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentValidationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentValidationFailed)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentValidationFailed)
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
func (it *LilypadPaymentEngineLilypadPaymentValidationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentValidationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentValidationFailed represents a LilypadPaymentValidationFailed event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationFailed struct {
	JobCreator       common.Address
	ResourceProvider common.Address
	Validator        common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentValidationFailed is a free log retrieval operation binding the contract event 0x100daa80ae82a115325870a0c921a9b23bdf6d087a4469411b3acc386aa25675.
//
// Solidity: event LilypadPayment__ValidationFailed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentValidationFailed(opts *bind.FilterOpts, jobCreator []common.Address, resourceProvider []common.Address, validator []common.Address) (*LilypadPaymentEngineLilypadPaymentValidationFailedIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ValidationFailed", jobCreatorRule, resourceProviderRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentValidationFailedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ValidationFailed", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentValidationFailed is a free log subscription operation binding the contract event 0x100daa80ae82a115325870a0c921a9b23bdf6d087a4469411b3acc386aa25675.
//
// Solidity: event LilypadPayment__ValidationFailed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentValidationFailed(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentValidationFailed, jobCreator []common.Address, resourceProvider []common.Address, validator []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ValidationFailed", jobCreatorRule, resourceProviderRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentValidationFailed)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationFailed", log); err != nil {
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

// ParseLilypadPaymentValidationFailed is a log parse operation binding the contract event 0x100daa80ae82a115325870a0c921a9b23bdf6d087a4469411b3acc386aa25675.
//
// Solidity: event LilypadPayment__ValidationFailed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentValidationFailed(log types.Log) (*LilypadPaymentEngineLilypadPaymentValidationFailed, error) {
	event := new(LilypadPaymentEngineLilypadPaymentValidationFailed)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentValidationPassedIterator is returned from FilterLilypadPaymentValidationPassed and is used to iterate over the raw logs and unpacked data for LilypadPaymentValidationPassed events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationPassedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentValidationPassed // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentValidationPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentValidationPassed)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentValidationPassed)
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
func (it *LilypadPaymentEngineLilypadPaymentValidationPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentValidationPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentValidationPassed represents a LilypadPaymentValidationPassed event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationPassed struct {
	JobCreator       common.Address
	ResourceProvider common.Address
	Validator        common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentValidationPassed is a free log retrieval operation binding the contract event 0x9bf286a59ab4f968cbd1e7ae3520598bc9938fed1f355ba8cca8b5e5dd4e4abf.
//
// Solidity: event LilypadPayment__ValidationPassed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentValidationPassed(opts *bind.FilterOpts, jobCreator []common.Address, resourceProvider []common.Address, validator []common.Address) (*LilypadPaymentEngineLilypadPaymentValidationPassedIterator, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ValidationPassed", jobCreatorRule, resourceProviderRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentValidationPassedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ValidationPassed", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentValidationPassed is a free log subscription operation binding the contract event 0x9bf286a59ab4f968cbd1e7ae3520598bc9938fed1f355ba8cca8b5e5dd4e4abf.
//
// Solidity: event LilypadPayment__ValidationPassed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentValidationPassed(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentValidationPassed, jobCreator []common.Address, resourceProvider []common.Address, validator []common.Address) (event.Subscription, error) {

	var jobCreatorRule []interface{}
	for _, jobCreatorItem := range jobCreator {
		jobCreatorRule = append(jobCreatorRule, jobCreatorItem)
	}
	var resourceProviderRule []interface{}
	for _, resourceProviderItem := range resourceProvider {
		resourceProviderRule = append(resourceProviderRule, resourceProviderItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ValidationPassed", jobCreatorRule, resourceProviderRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentValidationPassed)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationPassed", log); err != nil {
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

// ParseLilypadPaymentValidationPassed is a log parse operation binding the contract event 0x9bf286a59ab4f968cbd1e7ae3520598bc9938fed1f355ba8cca8b5e5dd4e4abf.
//
// Solidity: event LilypadPayment__ValidationPassed(address indexed jobCreator, address indexed resourceProvider, address indexed validator, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentValidationPassed(log types.Log) (*LilypadPaymentEngineLilypadPaymentValidationPassed, error) {
	event := new(LilypadPaymentEngineLilypadPaymentValidationPassed)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationPassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator is returned from FilterLilypadPaymentValidationPoolWalletUpdated and is used to iterate over the raw logs and unpacked data for LilypadPaymentValidationPoolWalletUpdated events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated)
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
func (it *LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated represents a LilypadPaymentValidationPoolWalletUpdated event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated struct {
	NewValidationPoolWallet common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentValidationPoolWalletUpdated is a free log retrieval operation binding the contract event 0xcc86aa3e89fa6b1716afd46c4fda5ac91cfafb53c0fc828016b605e58243fc99.
//
// Solidity: event LilypadPayment__ValidationPoolWalletUpdated(address newValidationPoolWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentValidationPoolWalletUpdated(opts *bind.FilterOpts) (*LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ValidationPoolWalletUpdated")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdatedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ValidationPoolWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentValidationPoolWalletUpdated is a free log subscription operation binding the contract event 0xcc86aa3e89fa6b1716afd46c4fda5ac91cfafb53c0fc828016b605e58243fc99.
//
// Solidity: event LilypadPayment__ValidationPoolWalletUpdated(address newValidationPoolWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentValidationPoolWalletUpdated(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ValidationPoolWalletUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationPoolWalletUpdated", log); err != nil {
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

// ParseLilypadPaymentValidationPoolWalletUpdated is a log parse operation binding the contract event 0xcc86aa3e89fa6b1716afd46c4fda5ac91cfafb53c0fc828016b605e58243fc99.
//
// Solidity: event LilypadPayment__ValidationPoolWalletUpdated(address newValidationPoolWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentValidationPoolWalletUpdated(log types.Log) (*LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated, error) {
	event := new(LilypadPaymentEngineLilypadPaymentValidationPoolWalletUpdated)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValidationPoolWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator is returned from FilterLilypadPaymentValueBasedRewardsWalletUpdated and is used to iterate over the raw logs and unpacked data for LilypadPaymentValueBasedRewardsWalletUpdated events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated)
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
func (it *LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated represents a LilypadPaymentValueBasedRewardsWalletUpdated event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated struct {
	NewValueBasedRewardsWallet common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentValueBasedRewardsWalletUpdated is a free log retrieval operation binding the contract event 0x15d760d36b9389dde13705a91ff4518b98034adfc40167541e152a572ba9ae49.
//
// Solidity: event LilypadPayment__ValueBasedRewardsWalletUpdated(address newValueBasedRewardsWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentValueBasedRewardsWalletUpdated(opts *bind.FilterOpts) (*LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ValueBasedRewardsWalletUpdated")
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdatedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ValueBasedRewardsWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentValueBasedRewardsWalletUpdated is a free log subscription operation binding the contract event 0x15d760d36b9389dde13705a91ff4518b98034adfc40167541e152a572ba9ae49.
//
// Solidity: event LilypadPayment__ValueBasedRewardsWalletUpdated(address newValueBasedRewardsWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentValueBasedRewardsWalletUpdated(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated) (event.Subscription, error) {

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ValueBasedRewardsWalletUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValueBasedRewardsWalletUpdated", log); err != nil {
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

// ParseLilypadPaymentValueBasedRewardsWalletUpdated is a log parse operation binding the contract event 0x15d760d36b9389dde13705a91ff4518b98034adfc40167541e152a572ba9ae49.
//
// Solidity: event LilypadPayment__ValueBasedRewardsWalletUpdated(address newValueBasedRewardsWallet)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentValueBasedRewardsWalletUpdated(log types.Log) (*LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated, error) {
	event := new(LilypadPaymentEngineLilypadPaymentValueBasedRewardsWalletUpdated)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ValueBasedRewardsWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator is returned from FilterLilypadPaymentZeroAmountPayout and is used to iterate over the raw logs and unpacked data for LilypadPaymentZeroAmountPayout events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentZeroAmountPayout // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentZeroAmountPayout)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentZeroAmountPayout)
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
func (it *LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentZeroAmountPayout represents a LilypadPaymentZeroAmountPayout event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentZeroAmountPayout struct {
	IntendedRecipient common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentZeroAmountPayout is a free log retrieval operation binding the contract event 0xafb7b3ffde7639ad01789621f9018abdb0094b6dc43ccf889dae5f343b2f07f1.
//
// Solidity: event LilypadPayment__ZeroAmountPayout(address indexed intended_recipient)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentZeroAmountPayout(opts *bind.FilterOpts, intended_recipient []common.Address) (*LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator, error) {

	var intended_recipientRule []interface{}
	for _, intended_recipientItem := range intended_recipient {
		intended_recipientRule = append(intended_recipientRule, intended_recipientItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__ZeroAmountPayout", intended_recipientRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentZeroAmountPayoutIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__ZeroAmountPayout", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentZeroAmountPayout is a free log subscription operation binding the contract event 0xafb7b3ffde7639ad01789621f9018abdb0094b6dc43ccf889dae5f343b2f07f1.
//
// Solidity: event LilypadPayment__ZeroAmountPayout(address indexed intended_recipient)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentZeroAmountPayout(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentZeroAmountPayout, intended_recipient []common.Address) (event.Subscription, error) {

	var intended_recipientRule []interface{}
	for _, intended_recipientItem := range intended_recipient {
		intended_recipientRule = append(intended_recipientRule, intended_recipientItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__ZeroAmountPayout", intended_recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentZeroAmountPayout)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ZeroAmountPayout", log); err != nil {
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

// ParseLilypadPaymentZeroAmountPayout is a log parse operation binding the contract event 0xafb7b3ffde7639ad01789621f9018abdb0094b6dc43ccf889dae5f343b2f07f1.
//
// Solidity: event LilypadPayment__ZeroAmountPayout(address indexed intended_recipient)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentZeroAmountPayout(log types.Log) (*LilypadPaymentEngineLilypadPaymentZeroAmountPayout, error) {
	event := new(LilypadPaymentEngineLilypadPaymentZeroAmountPayout)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__ZeroAmountPayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentEscrowPaidIterator is returned from FilterLilypadPaymentEscrowPaid and is used to iterate over the raw logs and unpacked data for LilypadPaymentEscrowPaid events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowPaidIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentEscrowPaid // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentEscrowPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowPaid)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowPaid)
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
func (it *LilypadPaymentEngineLilypadPaymentEscrowPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentEscrowPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentEscrowPaid represents a LilypadPaymentEscrowPaid event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowPaid struct {
	Payee         common.Address
	PaymentReason uint8
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentEscrowPaid is a free log retrieval operation binding the contract event 0x716de7e83ac36520a9d7a838ed00b942c36ce99d099834ee5d63ff58e976f2c6.
//
// Solidity: event LilypadPayment__escrowPaid(address indexed payee, uint8 indexed paymentReason, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentEscrowPaid(opts *bind.FilterOpts, payee []common.Address, paymentReason []uint8) (*LilypadPaymentEngineLilypadPaymentEscrowPaidIterator, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var paymentReasonRule []interface{}
	for _, paymentReasonItem := range paymentReason {
		paymentReasonRule = append(paymentReasonRule, paymentReasonItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__escrowPaid", payeeRule, paymentReasonRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentEscrowPaidIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__escrowPaid", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentEscrowPaid is a free log subscription operation binding the contract event 0x716de7e83ac36520a9d7a838ed00b942c36ce99d099834ee5d63ff58e976f2c6.
//
// Solidity: event LilypadPayment__escrowPaid(address indexed payee, uint8 indexed paymentReason, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentEscrowPaid(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentEscrowPaid, payee []common.Address, paymentReason []uint8) (event.Subscription, error) {

	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}
	var paymentReasonRule []interface{}
	for _, paymentReasonItem := range paymentReason {
		paymentReasonRule = append(paymentReasonRule, paymentReasonItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__escrowPaid", payeeRule, paymentReasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentEscrowPaid)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowPaid", log); err != nil {
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

// ParseLilypadPaymentEscrowPaid is a log parse operation binding the contract event 0x716de7e83ac36520a9d7a838ed00b942c36ce99d099834ee5d63ff58e976f2c6.
//
// Solidity: event LilypadPayment__escrowPaid(address indexed payee, uint8 indexed paymentReason, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentEscrowPaid(log types.Log) (*LilypadPaymentEngineLilypadPaymentEscrowPaid, error) {
	event := new(LilypadPaymentEngineLilypadPaymentEscrowPaid)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator is returned from FilterLilypadPaymentEscrowPayout and is used to iterate over the raw logs and unpacked data for LilypadPaymentEscrowPayout events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentEscrowPayout // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowPayout)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowPayout)
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
func (it *LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentEscrowPayout represents a LilypadPaymentEscrowPayout event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowPayout struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentEscrowPayout is a free log retrieval operation binding the contract event 0x132543e11dad0c907f02cf7b2a5e8fa1c827b82112e4534976f0e641f7ec64e1.
//
// Solidity: event LilypadPayment__escrowPayout(address indexed to, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentEscrowPayout(opts *bind.FilterOpts, to []common.Address) (*LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__escrowPayout", toRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentEscrowPayoutIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__escrowPayout", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentEscrowPayout is a free log subscription operation binding the contract event 0x132543e11dad0c907f02cf7b2a5e8fa1c827b82112e4534976f0e641f7ec64e1.
//
// Solidity: event LilypadPayment__escrowPayout(address indexed to, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentEscrowPayout(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentEscrowPayout, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__escrowPayout", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentEscrowPayout)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowPayout", log); err != nil {
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

// ParseLilypadPaymentEscrowPayout is a log parse operation binding the contract event 0x132543e11dad0c907f02cf7b2a5e8fa1c827b82112e4534976f0e641f7ec64e1.
//
// Solidity: event LilypadPayment__escrowPayout(address indexed to, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentEscrowPayout(log types.Log) (*LilypadPaymentEngineLilypadPaymentEscrowPayout, error) {
	event := new(LilypadPaymentEngineLilypadPaymentEscrowPayout)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowPayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator is returned from FilterLilypadPaymentEscrowSlashed and is used to iterate over the raw logs and unpacked data for LilypadPaymentEscrowSlashed events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentEscrowSlashed // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowSlashed)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowSlashed)
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
func (it *LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentEscrowSlashed represents a LilypadPaymentEscrowSlashed event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowSlashed struct {
	Account common.Address
	Actor   uint8
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentEscrowSlashed is a free log retrieval operation binding the contract event 0x5395458998742163fa2bf7601281eab0d4a1332e05b0c514f3762e7049b76cb0.
//
// Solidity: event LilypadPayment__escrowSlashed(address indexed account, uint8 indexed actor, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentEscrowSlashed(opts *bind.FilterOpts, account []common.Address, actor []uint8) (*LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__escrowSlashed", accountRule, actorRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentEscrowSlashedIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__escrowSlashed", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentEscrowSlashed is a free log subscription operation binding the contract event 0x5395458998742163fa2bf7601281eab0d4a1332e05b0c514f3762e7049b76cb0.
//
// Solidity: event LilypadPayment__escrowSlashed(address indexed account, uint8 indexed actor, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentEscrowSlashed(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentEscrowSlashed, account []common.Address, actor []uint8) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var actorRule []interface{}
	for _, actorItem := range actor {
		actorRule = append(actorRule, actorItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__escrowSlashed", accountRule, actorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentEscrowSlashed)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowSlashed", log); err != nil {
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

// ParseLilypadPaymentEscrowSlashed is a log parse operation binding the contract event 0x5395458998742163fa2bf7601281eab0d4a1332e05b0c514f3762e7049b76cb0.
//
// Solidity: event LilypadPayment__escrowSlashed(address indexed account, uint8 indexed actor, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentEscrowSlashed(log types.Log) (*LilypadPaymentEngineLilypadPaymentEscrowSlashed, error) {
	event := new(LilypadPaymentEngineLilypadPaymentEscrowSlashed)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator is returned from FilterLilypadPaymentEscrowWithdrawn and is used to iterate over the raw logs and unpacked data for LilypadPaymentEscrowWithdrawn events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator struct {
	Event *LilypadPaymentEngineLilypadPaymentEscrowWithdrawn // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowWithdrawn)
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
		it.Event = new(LilypadPaymentEngineLilypadPaymentEscrowWithdrawn)
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
func (it *LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineLilypadPaymentEscrowWithdrawn represents a LilypadPaymentEscrowWithdrawn event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineLilypadPaymentEscrowWithdrawn struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLilypadPaymentEscrowWithdrawn is a free log retrieval operation binding the contract event 0x304648227986fb683486b3c3592863900205c472b40b44a9ae754f643b85703e.
//
// Solidity: event LilypadPayment__escrowWithdrawn(address indexed withdrawer, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterLilypadPaymentEscrowWithdrawn(opts *bind.FilterOpts, withdrawer []common.Address) (*LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "LilypadPayment__escrowWithdrawn", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineLilypadPaymentEscrowWithdrawnIterator{contract: _LilypadPaymentEngine.contract, event: "LilypadPayment__escrowWithdrawn", logs: logs, sub: sub}, nil
}

// WatchLilypadPaymentEscrowWithdrawn is a free log subscription operation binding the contract event 0x304648227986fb683486b3c3592863900205c472b40b44a9ae754f643b85703e.
//
// Solidity: event LilypadPayment__escrowWithdrawn(address indexed withdrawer, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchLilypadPaymentEscrowWithdrawn(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineLilypadPaymentEscrowWithdrawn, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "LilypadPayment__escrowWithdrawn", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineLilypadPaymentEscrowWithdrawn)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowWithdrawn", log); err != nil {
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

// ParseLilypadPaymentEscrowWithdrawn is a log parse operation binding the contract event 0x304648227986fb683486b3c3592863900205c472b40b44a9ae754f643b85703e.
//
// Solidity: event LilypadPayment__escrowWithdrawn(address indexed withdrawer, uint256 amount)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseLilypadPaymentEscrowWithdrawn(log types.Log) (*LilypadPaymentEngineLilypadPaymentEscrowWithdrawn, error) {
	event := new(LilypadPaymentEngineLilypadPaymentEscrowWithdrawn)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "LilypadPayment__escrowWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleAdminChangedIterator struct {
	Event *LilypadPaymentEngineRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineRoleAdminChanged)
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
		it.Event = new(LilypadPaymentEngineRoleAdminChanged)
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
func (it *LilypadPaymentEngineRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineRoleAdminChanged represents a RoleAdminChanged event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LilypadPaymentEngineRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineRoleAdminChangedIterator{contract: _LilypadPaymentEngine.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineRoleAdminChanged)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseRoleAdminChanged(log types.Log) (*LilypadPaymentEngineRoleAdminChanged, error) {
	event := new(LilypadPaymentEngineRoleAdminChanged)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleGrantedIterator struct {
	Event *LilypadPaymentEngineRoleGranted // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineRoleGranted)
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
		it.Event = new(LilypadPaymentEngineRoleGranted)
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
func (it *LilypadPaymentEngineRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineRoleGranted represents a RoleGranted event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadPaymentEngineRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineRoleGrantedIterator{contract: _LilypadPaymentEngine.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineRoleGranted)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseRoleGranted(log types.Log) (*LilypadPaymentEngineRoleGranted, error) {
	event := new(LilypadPaymentEngineRoleGranted)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadPaymentEngineRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleRevokedIterator struct {
	Event *LilypadPaymentEngineRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LilypadPaymentEngineRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadPaymentEngineRoleRevoked)
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
		it.Event = new(LilypadPaymentEngineRoleRevoked)
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
func (it *LilypadPaymentEngineRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadPaymentEngineRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadPaymentEngineRoleRevoked represents a RoleRevoked event raised by the LilypadPaymentEngine contract.
type LilypadPaymentEngineRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LilypadPaymentEngineRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LilypadPaymentEngineRoleRevokedIterator{contract: _LilypadPaymentEngine.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LilypadPaymentEngineRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LilypadPaymentEngine.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadPaymentEngineRoleRevoked)
				if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LilypadPaymentEngine *LilypadPaymentEngineFilterer) ParseRoleRevoked(log types.Log) (*LilypadPaymentEngineRoleRevoked, error) {
	event := new(LilypadPaymentEngineRoleRevoked)
	if err := _LilypadPaymentEngine.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
