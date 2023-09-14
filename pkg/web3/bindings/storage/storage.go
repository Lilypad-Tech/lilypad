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
	InstructionCount *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b614e7580620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101a7575f3560e01c8063822f80c5116100f7578063cdd82d1d11610095578063e850be371161006f578063e850be37146104d9578063ec95b967146104f5578063f2fde38b14610525578063f3d3d44814610541576101a7565b8063cdd82d1d1461045d578063e70791801461048d578063e7b957d1146104bd576101a7565b8063a4702958116100d1578063a4702958146103d5578063a6370b0e146103df578063b050e74b1461040f578063c57380a21461043f576101a7565b8063822f80c51461036b578063824518aa1461039b5780638da5cb5b146103b7576101a7565b8063511a9f6811610164578063795f9abf1161013e578063795f9abf146102f957806380ffdfe0146103155780638129fc1c146103315780638224ce5f1461033b576101a7565b8063511a9f68146102a3578063715018a6146102bf57806373db5c6a146102c9576101a7565b806311d5af33146101ab5780632244ad2b146101db578063297f9e551461020b5780633c4135da1461022757806346834d1e14610257578063498cc70d14610273575b5f80fd5b6101c560048036038101906101c09190612f78565b61055d565b6040516101d291906130e8565b60405180910390f35b6101f560048036038101906101f09190613234565b61066e565b6040516102029190613295565b60405180910390f35b61022560048036038101906102209190613234565b610685565b005b610241600480360381019061023c9190613234565b61070c565b60405161024e9190613457565b60405180910390f35b610271600480360381019061026c9190613234565b6108e3565b005b61028d60048036038101906102889190613234565b61096a565b60405161029a91906134c5565b60405180910390f35b6102bd60048036038101906102b89190613234565b610acb565b005b6102c7610b52565b005b6102e360048036038101906102de9190613234565b610b65565b6040516102f091906134f4565b60405180910390f35b610313600480360381019061030e9190613234565b610bbd565b005b61032f600480360381019061032a9190613234565b610c43565b005b610339610cca565b005b61035560048036038101906103509190613234565b610e01565b60405161036291906134f4565b60405180910390f35b61038560048036038101906103809190613537565b610e41565b60405161039291906134c5565b60405180910390f35b6103b560048036038101906103b09190613234565b611098565b005b6103bf61111f565b6040516103cc91906135ce565b60405180910390f35b6103dd611146565b005b6103f960048036038101906103f4919061387c565b61116a565b6040516104069190613b6e565b60405180910390f35b61042960048036038101906104249190613bb1565b611848565b6040516104369190613295565b60405180910390f35b6104476118e0565b60405161045491906135ce565b60405180910390f35b61047760048036038101906104729190613234565b611908565b6040516104849190613457565b60405180910390f35b6104a760048036038101906104a29190613234565b6119fc565b6040516104b49190613b6e565b60405180910390f35b6104d760048036038101906104d29190613234565b611d4a565b005b6104f360048036038101906104ee9190613234565b611dd1565b005b61050f600480360381019061050a9190613234565b611e58565b60405161051c9190613457565b60405180910390f35b61053f600480360381019061053a9190612f78565b61202f565b005b61055b60048036038101906105569190612f78565b6120b1565b005b606060045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610663578382905f5260205f200180546105d890613c38565b80601f016020809104026020016040519081016040528092919081815260200182805461060490613c38565b801561064f5780601f106106265761010080835404028352916020019161064f565b820191905f5260205f20905b81548152906001019060200180831161063257829003601f168201915b5050505050815260200190600101906105bb565b505050509050919050565b5f80610679836119fc565b5f015151119050919050565b61068d6121b9565b50610699816002611848565b6106d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cf90613cc2565b60405180910390fd5b426005826040516106e99190613d1a565b9081526020016040518091039020600601819055506107098160036122e6565b50565b610714612cbb565b61071c6121b9565b506107268261066e565b610765576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075c90613d7a565b60405180910390fd5b5f6005836040516107769190613d1a565b908152602001604051809103902060020154146107c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bf90613de2565b60405180910390fd5b426005836040516107d99190613d1a565b9081526020016040518091039020600201819055506107f78261236a565b6005826040516108079190613d1a565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115610844576108436132ae565b5b600a811115610856576108556132ae565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6108eb6121b9565b506108f7816002611848565b610936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092d90613cc2565b60405180910390fd5b426005826040516109479190613d1a565b9081526020016040518091039020600701819055506109678160046122e6565b50565b610972612d2e565b6006826040516109829190613d1a565b90815260200160405180910390206040518060600160405290815f820180546109aa90613c38565b80601f01602080910402602001604051908101604052809291908181526020018280546109d690613c38565b8015610a215780601f106109f857610100808354040283529160200191610a21565b820191905f5260205f20905b815481529060010190602001808311610a0457829003601f168201915b50505050508152602001600182018054610a3a90613c38565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6690613c38565b8015610ab15780601f10610a8857610100808354040283529160200191610ab1565b820191905f5260205f20905b815481529060010190602001808311610a9457829003601f168201915b505050505081526020016002820154815250509050919050565b610ad36121b9565b50610adf816001611848565b610b1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1590613e4a565b60405180910390fd5b42600582604051610b2f9190613d1a565b9081526020016040518091039020600b0181905550610b4f8160086122e6565b50565b610b5a612421565b610b635f61249f565b565b5f600682604051610b769190613d1a565b908152602001604051809103902060020154600383604051610b989190613d1a565b9081526020016040518091039020600d015f0154610bb69190613e95565b9050919050565b610bc56121b9565b50610bd0815f611848565b610c0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0690613f20565b60405180910390fd5b42600582604051610c209190613d1a565b9081526020016040518091039020600a0181905550610c408160076122e6565b50565b610c4b6121b9565b50610c57816004611848565b610c96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8d90613f88565b60405180910390fd5b42600582604051610ca79190613d1a565b908152602001604051809103902060090181905550610cc78160066122e6565b50565b5f600160169054906101000a900460ff16159050808015610cfc575060018060159054906101000a900460ff1660ff16105b80610d2a5750610d0b30612560565b158015610d29575060018060159054906101000a900460ff1660ff16145b5b610d69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6090614016565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610da55760018060166101000a81548160ff0219169083151502179055505b8015610dfe575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610df59190614082565b60405180910390a15b50565b5f610e0b82610b65565b600383604051610e1b9190613d1a565b9081526020016040518091039020600d0160020154610e3a9190613e95565b9050919050565b610e49612d2e565b610e516121b9565b50610e5d846001611848565b610e9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9390613e4a565b60405180910390fd5b42600585604051610ead9190613d1a565b908152602001604051809103902060050181905550610ecd8460026122e6565b604051806060016040528085815260200184815260200183815250600685604051610ef89190613d1a565b90815260200160405180910390205f820151815f019081610f19919061422f565b506020820151816001019081610f2f919061422f565b5060408201518160020155905050600684604051610f4d9190613d1a565b90815260200160405180910390206040518060600160405290815f82018054610f7590613c38565b80601f0160208091040260200160405190810160405280929190818152602001828054610fa190613c38565b8015610fec5780601f10610fc357610100808354040283529160200191610fec565b820191905f5260205f20905b815481529060010190602001808311610fcf57829003601f168201915b5050505050815260200160018201805461100590613c38565b80601f016020809104026020016040519081016040528092919081815260200182805461103190613c38565b801561107c5780601f106110535761010080835404028352916020019161107c565b820191905f5260205f20905b81548152906001019060200180831161105f57829003601f168201915b5050505050815260200160028201548152505090509392505050565b6110a06121b9565b506110ac816004611848565b6110eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e290613f88565b60405180910390fd5b426005826040516110fc9190613d1a565b90815260200160405180910390206008018190555061111c8160056122e6565b50565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61114e612421565b5f600160146101000a81548160ff021916908315150217905550565b611172612d4e565b61117a6121b9565b50611185855f611848565b6111c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bb90613f20565b60405180910390fd5b6111cd84612582565b6111d683612797565b6111df8561066e565b15611220575f6111ee866119fc565b90506111fe81602001518661282d565b61120c816040015185612aaa565b61121a816060015184612af4565b506114ff565b6040518060800160405280868152602001858152602001848152602001838152506003866040516112519190613d1a565b90815260200160405180910390205f820151815f019081611272919061422f565b506020820151816001015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019080519060200190611368929190612d88565b5050506040820151816005015f820151815f015f820151815f01556020820151816001015550506020820151816002015f820151815f01556020820151816001015550506040820151816004015f820151815f01556020820151816001015550506060820151816006015f820151815f01556020820151816001015550505050606082015181600d015f820151815f0155602082015181600101556040820151816002015560608201518160030155505090505060045f856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f90919091909150908161148c919061422f565b5060045f856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2085908060018154018082558091505060019003905f5260205f20015f9091909190915090816114fd919061422f565b505b60038560405161150f9190613d1a565b90815260200160405180910390206040518060800160405290815f8201805461153790613c38565b80601f016020809104026020016040519081016040528092919081815260200182805461156390613c38565b80156115ae5780601f10611585576101008083540402835291602001916115ae565b820191905f5260205f20905b81548152906001019060200180831161159157829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820180548060200260200160405190810160405280929190818152602001828054801561174757602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116116fe575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b5f6118528361066e565b611884575f600a811115611869576118686132ae565b5b82600a81111561187c5761187b6132ae565b5b1490506118da565b81600a811115611897576118966132ae565b5b6005846040516118a79190613d1a565b90815260200160405180910390205f015f9054906101000a900460ff16600a8111156118d6576118d56132ae565b5b1490505b92915050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611910612cbb565b6005826040516119209190613d1a565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a81111561195d5761195c6132ae565b5b600a81111561196f5761196e6132ae565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b611a04612d4e565b600382604051611a149190613d1a565b90815260200160405180910390206040518060800160405290815f82018054611a3c90613c38565b80601f0160208091040260200160405190810160405280929190818152602001828054611a6890613c38565b8015611ab35780601f10611a8a57610100808354040283529160200191611ab3565b820191905f5260205f20905b815481529060010190602001808311611a9657829003601f168201915b50505050508152602001600182016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611c4c57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611c03575b5050505050815250508152602001600582016040518060800160405290815f82016040518060400160405290815f82015481526020016001820154815250508152602001600282016040518060400160405290815f82015481526020016001820154815250508152602001600482016040518060400160405290815f82015481526020016001820154815250508152602001600682016040518060400160405290815f8201548152602001600182015481525050815250508152602001600d82016040518060800160405290815f82015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611d526121b9565b50611d5e816004611848565b611d9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d9490613f88565b60405180910390fd5b42600582604051611dae9190613d1a565b9081526020016040518091039020600d0181905550611dce81600a6122e6565b50565b611dd96121b9565b50611de5816002611848565b611e24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e1b90613cc2565b60405180910390fd5b42600582604051611e359190613d1a565b9081526020016040518091039020600c0181905550611e558160096122e6565b50565b611e60612cbb565b611e686121b9565b50611e728261066e565b611eb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ea890613d7a565b60405180910390fd5b5f600583604051611ec29190613d1a565b90815260200160405180910390206001015414611f14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f0b90614348565b60405180910390fd5b42600583604051611f259190613d1a565b908152602001604051809103902060010181905550611f438261236a565b600582604051611f539190613d1a565b9081526020016040518091039020604051806101c00160405290815f82015f9054906101000a900460ff16600a811115611f9057611f8f6132ae565b5b600a811115611fa257611fa16132ae565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b612037612421565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036120a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161209c906143d6565b60405180910390fd5b6120ae8161249f565b50565b6120b9612421565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161211e90614464565b60405180910390fd5b600160149054906101000a900460ff16612176576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161216d906144f2565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161224090614464565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16612289612c1e565b73ffffffffffffffffffffffffffffffffffffffff16146122df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d690614580565b60405180910390fd5b6001905090565b806005836040516122f79190613d1a565b90815260200160405180910390205f015f6101000a81548160ff0219169083600a811115612328576123276132ae565b5b02179055507f10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386828260405161235e9291906145e5565b60405180910390a15050565b5f60058260405161237b9190613d1a565b908152602001604051809103902060010154141580156123bc57505f6005826040516123a79190613d1a565b90815260200160405180910390206002015414155b156123f757426005826040516123d29190613d1a565b9081526020016040518091039020600401819055506123f28160016122e6565b61241e565b426005826040516124089190613d1a565b9081526020016040518091039020600301819055505b50565b612429612c1e565b73ffffffffffffffffffffffffffffffffffffffff1661244761111f565b73ffffffffffffffffffffffffffffffffffffffff161461249d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124949061465d565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16036125f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125eb906146c5565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161265d9061472d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff16036126d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126ce90614795565b60405180910390fd5b5f8160600151511161271e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612715906147fd565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1603612794576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161278b90614865565b60405180910390fd5b50565b5f815f015160200151146127e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127d7906148cd565b60405180910390fd5b5f8160600151602001511461282a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282190614935565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff16146128a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161289a9061499d565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612919576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161291090614a05565b60405180910390fd5b805f015173ffffffffffffffffffffffffffffffffffffffff16825f015173ffffffffffffffffffffffffffffffffffffffff161461298d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161298490614a6d565b60405180910390fd5b806060015151826060015151146129d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129d090614ad5565b60405180910390fd5b5f5b826060015151811015612aa557816060015181815181106129ff576129fe614af3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612a3457612a33614af3565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612a92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a8990614b6a565b60405180910390fd5b8080612a9d90614b88565b9150506129db565b505050565b612aba825f0151825f0151612c25565b612acc82602001518260200151612c25565b612ade82604001518260400151612c25565b612af082606001518260600151612c25565b5050565b805f0151825f015114612b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b3390614c19565b60405180910390fd5b8060200151826020015114612b86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b7d90614c81565b60405180910390fd5b8060400151826040015114612bd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bc790614ce9565b60405180910390fd5b8060600151826060015114612c1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1190614d51565b60405180910390fd5b5050565b5f33905090565b805f0151825f015114612c6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c6490614db9565b60405180910390fd5b8060200151826020015114612cb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cae90614e21565b60405180910390fd5b5050565b604051806101c001604052805f600a811115612cda57612cd96132ae565b5b81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b604051806060016040528060608152602001606081526020015f81525090565b604051806080016040528060608152602001612d68612e0f565b8152602001612d75612e76565b8152602001612d82612eb6565b81525090565b828054828255905f5260205f20908101928215612dfe579160200282015b82811115612dfd578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612da6565b5b509050612e0b9190612eda565b5090565b60405180608001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6040518060800160405280612e89612ef5565b8152602001612e96612ef5565b8152602001612ea3612ef5565b8152602001612eb0612ef5565b81525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5b80821115612ef1575f815f905550600101612edb565b5090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f612f4782612f1e565b9050919050565b612f5781612f3d565b8114612f61575f80fd5b50565b5f81359050612f7281612f4e565b92915050565b5f60208284031215612f8d57612f8c612f16565b5b5f612f9a84828501612f64565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015613003578082015181840152602081019050612fe8565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61302882612fcc565b6130328185612fd6565b9350613042818560208601612fe6565b61304b8161300e565b840191505092915050565b5f613061838361301e565b905092915050565b5f602082019050919050565b5f61307f82612fa3565b6130898185612fad565b93508360208202850161309b85612fbd565b805f5b858110156130d657848403895281516130b78582613056565b94506130c283613069565b925060208a0199505060018101905061309e565b50829750879550505050505092915050565b5f6020820190508181035f8301526131008184613075565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6131468261300e565b810181811067ffffffffffffffff8211171561316557613164613110565b5b80604052505050565b5f613177612f0d565b9050613183828261313d565b919050565b5f67ffffffffffffffff8211156131a2576131a1613110565b5b6131ab8261300e565b9050602081019050919050565b828183375f83830152505050565b5f6131d86131d384613188565b61316e565b9050828152602081018484840111156131f4576131f361310c565b5b6131ff8482856131b8565b509392505050565b5f82601f83011261321b5761321a613108565b5b813561322b8482602086016131c6565b91505092915050565b5f6020828403121561324957613248612f16565b5b5f82013567ffffffffffffffff81111561326657613265612f1a565b5b61327284828501613207565b91505092915050565b5f8115159050919050565b61328f8161327b565b82525050565b5f6020820190506132a85f830184613286565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600b81106132ec576132eb6132ae565b5b50565b5f8190506132fc826132db565b919050565b5f61330b826132ef565b9050919050565b61331b81613301565b82525050565b5f819050919050565b61333381613321565b82525050565b6101c082015f82015161334e5f850182613312565b506020820151613361602085018261332a565b506040820151613374604085018261332a565b506060820151613387606085018261332a565b50608082015161339a608085018261332a565b5060a08201516133ad60a085018261332a565b5060c08201516133c060c085018261332a565b5060e08201516133d360e085018261332a565b506101008201516133e861010085018261332a565b506101208201516133fd61012085018261332a565b5061014082015161341261014085018261332a565b5061016082015161342761016085018261332a565b5061018082015161343c61018085018261332a565b506101a08201516134516101a085018261332a565b50505050565b5f6101c08201905061346b5f830184613339565b92915050565b5f606083015f8301518482035f86015261348b828261301e565b915050602083015184820360208601526134a5828261301e565b91505060408301516134ba604086018261332a565b508091505092915050565b5f6020820190508181035f8301526134dd8184613471565b905092915050565b6134ee81613321565b82525050565b5f6020820190506135075f8301846134e5565b92915050565b61351681613321565b8114613520575f80fd5b50565b5f813590506135318161350d565b92915050565b5f805f6060848603121561354e5761354d612f16565b5b5f84013567ffffffffffffffff81111561356b5761356a612f1a565b5b61357786828701613207565b935050602084013567ffffffffffffffff81111561359857613597612f1a565b5b6135a486828701613207565b92505060406135b586828701613523565b9150509250925092565b6135c881612f3d565b82525050565b5f6020820190506135e15f8301846135bf565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561360957613608613110565b5b602082029050602081019050919050565b5f80fd5b5f61363061362b846135ef565b61316e565b905080838252602082019050602084028301858111156136535761365261361a565b5b835b8181101561367c57806136688882612f64565b845260208401935050602081019050613655565b5050509392505050565b5f82601f83011261369a57613699613108565b5b81356136aa84826020860161361e565b91505092915050565b5f608082840312156136c8576136c76135e7565b5b6136d2608061316e565b90505f6136e184828501612f64565b5f8301525060206136f484828501612f64565b602083015250604061370884828501612f64565b604083015250606082013567ffffffffffffffff81111561372c5761372b6135eb565b5b61373884828501613686565b60608301525092915050565b5f60408284031215613759576137586135e7565b5b613763604061316e565b90505f61377284828501613523565b5f83015250602061378584828501613523565b60208301525092915050565b5f61010082840312156137a7576137a66135e7565b5b6137b1608061316e565b90505f6137c084828501613744565b5f8301525060406137d384828501613744565b60208301525060806137e784828501613744565b60408301525060c06137fb84828501613744565b60608301525092915050565b5f6080828403121561381c5761381b6135e7565b5b613826608061316e565b90505f61383584828501613523565b5f83015250602061384884828501613523565b602083015250604061385c84828501613523565b604083015250606061387084828501613523565b60608301525092915050565b5f805f806101c0858703121561389557613894612f16565b5b5f85013567ffffffffffffffff8111156138b2576138b1612f1a565b5b6138be87828801613207565b945050602085013567ffffffffffffffff8111156138df576138de612f1a565b5b6138eb878288016136b3565b93505060406138fc87828801613791565b92505061014061390e87828801613807565b91505092959194509250565b61392381612f3d565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61395d838361391a565b60208301905092915050565b5f602082019050919050565b5f61397f82613929565b6139898185613933565b935061399483613943565b805f5b838110156139c45781516139ab8882613952565b97506139b683613969565b925050600181019050613997565b5085935050505092915050565b5f608083015f8301516139e65f86018261391a565b5060208301516139f9602086018261391a565b506040830151613a0c604086018261391a565b5060608301518482036060860152613a248282613975565b9150508091505092915050565b604082015f820151613a455f85018261332a565b506020820151613a58602085018261332a565b50505050565b61010082015f820151613a735f850182613a31565b506020820151613a866040850182613a31565b506040820151613a996080850182613a31565b506060820151613aac60c0850182613a31565b50505050565b608082015f820151613ac65f85018261332a565b506020820151613ad9602085018261332a565b506040820151613aec604085018261332a565b506060820151613aff606085018261332a565b50505050565b5f6101c083015f8301518482035f860152613b20828261301e565b91505060208301518482036020860152613b3a82826139d1565b9150506040830151613b4f6040860182613a5e565b506060830151613b63610140860182613ab2565b508091505092915050565b5f6020820190508181035f830152613b868184613b05565b905092915050565b600b8110613b9a575f80fd5b50565b5f81359050613bab81613b8e565b92915050565b5f8060408385031215613bc757613bc6612f16565b5b5f83013567ffffffffffffffff811115613be457613be3612f1a565b5b613bf085828601613207565b9250506020613c0185828601613b9d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680613c4f57607f821691505b602082108103613c6257613c61613c0b565b5b50919050565b5f82825260208201905092915050565b7f526573756c74735375626d6974746564000000000000000000000000000000005f82015250565b5f613cac601083613c68565b9150613cb782613c78565b602082019050919050565b5f6020820190508181035f830152613cd981613ca0565b9050919050565b5f81905092915050565b5f613cf482612fcc565b613cfe8185613ce0565b9350613d0e818560208601612fe6565b80840191505092915050565b5f613d258284613cea565b915081905092915050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f613d64601383613c68565b9150613d6f82613d30565b602082019050919050565b5f6020820190508181035f830152613d9181613d58565b9050919050565b7f4a432068617320616c72656164792061677265656400000000000000000000005f82015250565b5f613dcc601583613c68565b9150613dd782613d98565b602082019050919050565b5f6020820190508181035f830152613df981613dc0565b9050919050565b7f4465616c416772656564000000000000000000000000000000000000000000005f82015250565b5f613e34600a83613c68565b9150613e3f82613e00565b602082019050919050565b5f6020820190508181035f830152613e6181613e28565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613e9f82613321565b9150613eaa83613321565b9250828202613eb881613321565b91508282048414831517613ecf57613ece613e68565b5b5092915050565b7f4465616c4e65676f74696174696e6700000000000000000000000000000000005f82015250565b5f613f0a600f83613c68565b9150613f1582613ed6565b602082019050919050565b5f6020820190508181035f830152613f3781613efe565b9050919050565b7f526573756c7473436865636b65640000000000000000000000000000000000005f82015250565b5f613f72600e83613c68565b9150613f7d82613f3e565b602082019050919050565b5f6020820190508181035f830152613f9f81613f66565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f614000602e83613c68565b915061400b82613fa6565b604082019050919050565b5f6020820190508181035f83015261402d81613ff4565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61406c61406761406284614034565b614049565b61403d565b9050919050565b61407c81614052565b82525050565b5f6020820190506140955f830184614073565b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026140f77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826140bc565b61410186836140bc565b95508019841693508086168417925050509392505050565b5f61413361412e61412984613321565b614049565b613321565b9050919050565b5f819050919050565b61414c83614119565b6141606141588261413a565b8484546140c8565b825550505050565b5f90565b614174614168565b61417f818484614143565b505050565b5b818110156141a2576141975f8261416c565b600181019050614185565b5050565b601f8211156141e7576141b88161409b565b6141c1846140ad565b810160208510156141d0578190505b6141e46141dc856140ad565b830182614184565b50505b505050565b5f82821c905092915050565b5f6142075f19846008026141ec565b1980831691505092915050565b5f61421f83836141f8565b9150826002028217905092915050565b61423882612fcc565b67ffffffffffffffff81111561425157614250613110565b5b61425b8254613c38565b6142668282856141a6565b5f60209050601f831160018114614297575f8415614285578287015190505b61428f8582614214565b8655506142f6565b601f1984166142a58661409b565b5f5b828110156142cc578489015182556001820191506020850194506020810190506142a7565b868310156142e957848901516142e5601f8916826141f8565b8355505b6001600288020188555050505b505050505050565b7f52502068617320616c72656164792061677265656400000000000000000000005f82015250565b5f614332601583613c68565b915061433d826142fe565b602082019050919050565b5f6020820190508181035f83015261435f81614326565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f6143c0602683613c68565b91506143cb82614366565b604082019050919050565b5f6020820190508181035f8301526143ed816143b4565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f61444e603583613c68565b9150614459826143f4565b604082019050919050565b5f6020820190508181035f83015261447b81614442565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f6144dc603983613c68565b91506144e782614482565b604082019050919050565b5f6020820190508181035f830152614509816144d0565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f61456a603b83613c68565b915061457582614510565b604082019050919050565b5f6020820190508181035f8301526145978161455e565b9050919050565b5f6145a882612fcc565b6145b28185613c68565b93506145c2818560208601612fe6565b6145cb8161300e565b840191505092915050565b6145df81613301565b82525050565b5f6040820190508181035f8301526145fd818561459e565b905061460c60208301846145d6565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614647602083613c68565b915061465282614613565b602082019050919050565b5f6020820190508181035f8301526146748161463b565b9050919050565b7f5250206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f6146af600a83613c68565b91506146ba8261467b565b602082019050919050565b5f6020820190508181035f8301526146dc816146a3565b9050919050565b7f4a43206d697373696e67000000000000000000000000000000000000000000005f82015250565b5f614717600a83613c68565b9150614722826146e3565b602082019050919050565b5f6020820190508181035f8301526147448161470b565b9050919050565b7f536f6c766572206d697373696e670000000000000000000000000000000000005f82015250565b5f61477f600e83613c68565b915061478a8261474b565b602082019050919050565b5f6020820190508181035f8301526147ac81614773565b9050919050565b7f4d65646961746f7273203c3d20300000000000000000000000000000000000005f82015250565b5f6147e7600e83613c68565b91506147f2826147b3565b602082019050919050565b5f6020820190508181035f830152614814816147db565b9050919050565b7f5250202f204a432073616d6500000000000000000000000000000000000000005f82015250565b5f61484f600c83613c68565b915061485a8261481b565b602082019050919050565b5f6020820190508181035f83015261487c81614843565b9050919050565b7f4167726565206465706f736974206d75737420626520300000000000000000005f82015250565b5f6148b7601783613c68565b91506148c282614883565b602082019050919050565b5f6020820190508181035f8301526148e4816148ab565b9050919050565b7f4d656469617465206465706f736974206d7573742062652030000000000000005f82015250565b5f61491f601983613c68565b915061492a826148eb565b602082019050919050565b5f6020820190508181035f83015261494c81614913565b9050919050565b7f52500000000000000000000000000000000000000000000000000000000000005f82015250565b5f614987600283613c68565b915061499282614953565b602082019050919050565b5f6020820190508181035f8301526149b48161497b565b9050919050565b7f4a430000000000000000000000000000000000000000000000000000000000005f82015250565b5f6149ef600283613c68565b91506149fa826149bb565b602082019050919050565b5f6020820190508181035f830152614a1c816149e3565b9050919050565b7f536f6c76657200000000000000000000000000000000000000000000000000005f82015250565b5f614a57600683613c68565b9150614a6282614a23565b602082019050919050565b5f6020820190508181035f830152614a8481614a4b565b9050919050565b7f4d65646961746f727300000000000000000000000000000000000000000000005f82015250565b5f614abf600983613c68565b9150614aca82614a8b565b602082019050919050565b5f6020820190508181035f830152614aec81614ab3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4d65646961746f720000000000000000000000000000000000000000000000005f82015250565b5f614b54600883613c68565b9150614b5f82614b20565b602082019050919050565b5f6020820190508181035f830152614b8181614b48565b9050919050565b5f614b9282613321565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614bc457614bc3613e68565b5b600182019050919050565b7f50726963650000000000000000000000000000000000000000000000000000005f82015250565b5f614c03600583613c68565b9150614c0e82614bcf565b602082019050919050565b5f6020820190508181035f830152614c3081614bf7565b9050919050565b7f5061796d656e74000000000000000000000000000000000000000000000000005f82015250565b5f614c6b600783613c68565b9150614c7682614c37565b602082019050919050565b5f6020820190508181035f830152614c9881614c5f565b9050919050565b7f526573756c7473000000000000000000000000000000000000000000000000005f82015250565b5f614cd3600783613c68565b9150614cde82614c9f565b602082019050919050565b5f6020820190508181035f830152614d0081614cc7565b9050919050565b7f4d6564696174696f6e00000000000000000000000000000000000000000000005f82015250565b5f614d3b600983613c68565b9150614d4682614d07565b602082019050919050565b5f6020820190508181035f830152614d6881614d2f565b9050919050565b7f54696d656f7574000000000000000000000000000000000000000000000000005f82015250565b5f614da3600783613c68565b9150614dae82614d6f565b602082019050919050565b5f6020820190508181035f830152614dd081614d97565b9050919050565b7f436f6c6c61746572616c000000000000000000000000000000000000000000005f82015250565b5f614e0b600a83613c68565b9150614e1682614dd7565b602082019050919050565b5f6020820190508181035f830152614e3881614dff565b905091905056fea2646970667358221220ada7ece54df25add44f8dcf4c4ae4846f75f57229924b93ae69bed7236d1a78d64736f6c63430008150033",
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
